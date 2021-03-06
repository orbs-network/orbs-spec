# Boyar

> V2 release

This daemon service is the main orchestrator of the node. This is the first service to run and its purpose is to execute all the other services, provide their launch configuration, upgrade their software version, etc.

## Input Configuration

Boyar is a stateless service that receives all operation instructions as an external JSON configuration. There are two sources of configs that Boyar receives:

* Static bootstrap configuration - Provided once during launch as command-line arguments. In a live cloud-based Orbs node, this is normally set up by [Polygon](https://github.com/orbs-network/polygon) when the node is first created and originates from Polygon configuration.

    Format: [boyarin/README.md](https://github.com/orbs-network/boyarin/blob/master/README.md#Configuration)

* Dynamic configuration - Polled repeatedly by Boyar from an HTTP endpoint given in static configuration. This configuration reflects dynamic changes in the network such as which ONG instances should be running (based on which virtual chain subscriptions are currently active). In a public Orbs network, the dynamic configuration is served by [management service](MGMT-SERVICE.md).

    Format: [NODE-MGMT-CONFIG.md](NODE-MGMT-CONFIG.md)

&nbsp;

## Lifecycle

On machine boot, when Boyar is started, it retrieves from persistent storage the bootstrap config JSON which contains the static bootstrap configuration. In a public Orbs node for example, Boyar uses this bootstrap configuration to launch the initial management service.

As long as the management service is alive and serving valid dynamic configuration, Boyar polls this configuration repeatedly every minute and provisions this configuration (makes sure infrastructure like Docker Swarm reflects these requirements).

If the management service fails to provide valid dynamic configuration for a timeout period (30 minutes), Boyar falls back to the bootstrap flow and restarts itself according to the bootstrap config JSON. The boostrap flow does not know anything about previously applied configuration, and as a result virtual chains (ONGs) will be restarted, as well as other services.

If the Boyar process itself crashes and terminates, it is restarted automatically by the operating system and starts its bootstrap flow. This is currently implemented using [supervisord](http://supervisord.org/).

&nbsp;

## Docker Orchestration

Boyar relies on [Docker Swarm](https://docs.docker.com/engine/swarm/) for running all docker instances across the node cluster (note that a single Orbs node can span several machines). Docker Swarm monitors the instances, restarts them on errors and moves them across machines as needed (when a machine disappears or based on workload). Since Boyar instructs Docker Swarm, it is running outside of docker as a regular system process.

Since there’s no affinity between a docker instance and the physical machine it runs on, persistent storage must be machine agnostic. This is achieved using [NFS](https://en.wikipedia.org/wiki/Network_File_System) (eg. [EFS](https://aws.amazon.com/efs/) on AWS).

Boyar supports configuring the services it launches as docker instances via a static configuration file in a predefined path. This method supports sensitive information such as private keys as it is secured through [Docker secrets](https://docs.docker.com/engine/swarm/secrets/). In addition, Boyar relies on the service executable to have a predefined path. If a service has difficulty conforming to Boyar’s explicit launch contract, the recommended approach is to wrap the service with a [Dockerfile](https://docs.docker.com/engine/reference/builder/) that makes the appropriate adjustments. This Dockerfile is part of the service release cycle and not part of Boyar, since Boyar does not have intimate knowledge about the services it runs.

### Explicit Contract with Services

* Execution paths:

    * Working directory:

        ```
        /opt/orbs
        ```

    * Executable being ran: 

        ```
        /opt/orbs/service
        ```

        No command-line arguments given.
        
    * Directory for persistent data:
    
        ```
        /opt/orbs/cache
        ```

* Config files:

    * Config section of the service from the node management config:

        ```
        /run/secrets/config.json
        ```

        Note: in some Linux distros the secrets folder is `/var/run/secrets`

    * `network.json`
    
        ```json
        {
          "node-address": "a328846cd5b4979d68a8c58a9bdfeee657b34de7"
        }
        ```
    
        *todo* *- temporary, will eventually be killed*
    
    * `keys.json`
    
        *todo* *- temporary, will eventually be killed*

* Health check:

    * Implemented as an executable located at:

        ```
        /opt/orbs/healthcheck
        ```
        
        Note: The executable can also be a script / shell command. A common combination is to have a simple script that queries a status endpoint and triggers restart if it doesn't respond.

    * Returns exit 0 if the service is stable or non-zero if needs restart.

        * By default, the healthcheck is polled every 30 seconds by Docker Swarm. The service is restarted by Docker Swarm if healthcheck fails 3 consecutive checks. This can be monitored in a live system using Boyar's `status.json`.

        * To assist with debugging, health check should log important information to stdout. The output is available in a live system using Boyar's `status.json`.

* Status JSON:

    * A service should write JSON output containing non-sensitive public information about the current service status to:

        ```
        /opt/orbs/status/status.json
        ```

    * The format of the JSON is:

        ```json
        {
            "Error": "Human readable explanation of current error, field exists only if the status is erroneous.",
            "Status": "Human readable explanation of current status, field always exists.",
            "Timestamp": "2020-03-19T11:50:21.0846185Z",
            "Payload": {
                "Version": {
                    "Semantic": "v1.3.1"
                },
                "CustomFieldsGoHere": 17,
                "MoreCustomFields": "any data type"
            }
        }
        ```

        The timestamp is the last time the status was updated. The error field must appear if and only if the service is currently in an erroneous state and reports that it does not function properly. The status page for example will display the service in red if the error field exists, otherwise in green.

    * The status JSON is accessible to all via the following HTTP endpoints on the node gateway (Nginx):

        ```
        /services/{SERVICE-NAME}/status
        /vchains/{VCHAIN-ID}/status
        ```
        
* Logs:

    * Logs are accessible via the following HTTP endpoint on the node gateway (Nginx):

        ```        
        /services/{SERVICE-NAME}/logs
        /vchains/{VCHAIN-ID}/logs
        ```

    * More flexible access to logs (including live streaming log output and access to old rotated logs) is available through a dedicated service (see Logs Service). This service acquires access to the log files of all services using the `MountNodeLogs: true` node management configuration.

* Volumes:

    * Logs are located at:

        ```
        /opt/orbs/logs
        ```
    * Persistent cache (file storage for any service that remains between restarts) is located at:

        ```
        /opt/orbs/cache
        ```

 ### Internal Network Endpoints

* Internal network endpoints for node services are:

    ```
    http://{SERVICE-NAME}:{INTERNAL-PORT}
    ```

    For example: `http://management-service:8080`

    Service names are defined under [Image naming and tagging](../version-release/NAMING.md).

* Internal network endpoints for virtual chains (Public API) are:

    ```
    http://chain-{VCHAIN-ID}:{INTERNAL-PORT}
    ```

    For example: `http://chain-42:8080`

&nbsp;

## Private Networks

The Orbs architecture is primarily designed to accommodate the public Orbs network where Ethereum mainnet is the source of truth for network state. Nevertheless, the Orbs codebase supports running an isolated collection of nodes in a private network mode.

When running in a private network mode, Boyar does not instantiate the management service. Instead, it receives via static configuration a URL for the dynamic configuration JSON and polls it directly. This is indicated by the bootstrap management config JSON given via the `--config-url` command-line argument.

&nbsp;

## Partial Updates

Boyar supports intelligent partial updates in its JSON configuration and checks for diffs between a new dynamic configuration and the last configuration it acted upon. For example, if the launch configuration for a single node service changed (and all other parts of the JSON, like other node services, remain the same), Boyar will only restart the changed node service with the new launch configuration and keep the other services running.

In addition, if Boyar’s provided JSON configuration is partial (eg. some sections missing), Boyar will only act upon the provided parameters and ignore changes in the missing parts. For example, if a new dynamic configuration JSON does not contain the section describing virtual chains, Boyar will ignore changes to virtual chains and keep them running undisturbed.

&nbsp;

## Upgrading Boyar

Validator can opt in to allow Boyar to upgrade on new releases by supplying the Polygon configuration `boyarAutoUpdate: true`. When this is the case, the [node management configuration](NODE-MGMT-CONFIG.md) may specify the required Boyar binary version using the fields:

```
"orchestrator": {
    "ExecutableImage": {
        "Url": "https://github.com/orbs-network/boyarin/releases/download/v1.4.0/boyar-v1.4.0.bin",
        "Sha256": "1998cc1f7721acfe1954ab2878cc0ad8062cd6d919cd61fa22401c6750e195fe"
    },
}
```

If the SHA256 over its own binary is different than specified, Boyar will download the new binary, verify the SHA256 over the binary and if it matches the configuration, will terminate itself and the process will start again using the new binary (currently implemented with an external process manager, Supervisord).
