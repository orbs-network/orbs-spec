# Gradual Rollout

One of the requirements of the consensus algorithm is that the network maintains a running quorum of nodes at any given time. This presents a challenge with software upgrades and node restarts since we don’t want all nodes to go offline at the same time. To accommodate this, a gradual rollout scheme is used where every node randomizes their upgrade timeslot (and downtime). For simplicity, the gradual rollout scheme should be applied on releases to all node services and all virtual chains.

The [deploy](README.md#deploy-tool) script allows selecting the desired rollout plan for each deployment.

## Normal vs. Hotfix vs. Immediate

* **Normal** - Standard rollout events should take place across the entire network (all nodes) over the course of **24 hours**.
 
* **Hotfix** - In extreme conditions, like an urgent bugfix or an urgent recovery of a failing network, shorter time may be required. Hotfix roll out events should take place across the entire network (all nodes) over the course of **1 hour**.

* **Immediate** - In even more extreme conditions, like security vulnerabilities, immediate deployment is also supported. Please refrain from using this mode since virtual chain consensus may fail with forks if all virtual chain nodes are offline at the same time.

## Randomized Timeslots

Once the **rollout window** has been determined (24 hours or 1 hour), every node chooses its own random rollout time based on its own address and the current time:

```
rollout_time = local_time + [hash(own_address, current_time) mod rollout_window]
```

If the **local time** is greater than **rollout time**, the node should commence the upgrade.

Since the rollout mechanism is stateful (the node realizes a new version was just published), restarts of the management service may lose all state. In this edge case, it is ok for the node to upgrade immediately to all released versions without gradual rollout. Note that on a manual redeploy of the node, the newest revision is deployed regardless of the timeslot.
