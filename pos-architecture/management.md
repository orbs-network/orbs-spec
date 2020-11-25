# Contracts Management and Governance Architecture

Orbs PoS contacts architecture is designed for future migration, including migration of individual contracts. Contracts in the PoS system that interact with other contracts store their current address in the contact storage. Upon a migration of a contract, the registry sends a push notification to all its managed contacts to update their contract list accordingly.

In addition to the system contracts, the contract registry stores the up-to-date managers list. To ensure synchronization, the contacts query the registry on the current relevant manager before every governance operation. Managers are appointed by the registryAdmin that can appoint or revoke the role of a manager.

The contract registry emits an update event for any contact address update as well as upon registry migration. By knowing the initial contract registry address, an application may track the up to date contracts list as well as construct an integrated event list from all the contract versions.

An exception to contacts management by the contract registry are the protocol wallet contracts. The protocol wallet contracts have separate admins that manage their functionality and migration.  

### Roles
**registryAdmin** - the registryAdmin manages the contract registry and is the most privileged governance entity. The registryAdmin can update the registry properties and contacts and assign managers. The registryAdmin admin is stored on each of the contracts managed by the registry allowing to modify an individual contact appointed registry.

**initializationAdmin** - the initializationAdmin is responsible to initialize the contracts state, from governance parameters to network migrated state such as delegations, or guardians list. The initializationAdmin is set to the contract deployer and his actions may be considered as an extension to the contact constructor. The initializationAdmin is a highly privileged role, targeting migration periods, it can not be transferred and the role is revoked once initialization is completed.

**migrationManager** - the migrationManager is appointed by the contract registry, and controls all the migration actions. The migration manager can set contacts in the registry and migrate state and balances of contacts. The migrationManager has high privileges therefore it is recommended to appoint a migrationManager for a specific mission, such as a migration to a new contracts version and then revoke its privileges.

**functionalManager** - the functionalManager is appointed by the contract registry, and controls contracts governance parameters and functionality. 

**certificationManager** - the certificationManager manages the guardians certification. The certificationManager can set and clear a guardian certification status. The guardian certification is determined according to the guardian’s identification data as advertised in the registration contract using the ID url metadata.