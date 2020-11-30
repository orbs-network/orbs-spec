# Rollout to Staging

[release-to-staging](https://github.com/orbs-network/release-to-staging) repository contains scripts and instructions for launching a staging environment with selected image versions.

This repo will help you create your own staging network on [AWS](https://aws.amazon.com/).

Staging environment is a personal test network that is owned and operated in full by the developer. It is deployed using standard tooling like Polygon, normally as a private network (since all nodes are named explicitly).

In terms of simulating production, staging is much more accurate than automated end-to-end tests because automated end-to-end tests rarely include **all** subsystems from Polygon to Boyar to all services and virtual chains. Please note that this accuracy comes at a price and running a staging test can take more than 30 minutes.

It is recommended to run staging tests before deploying significant changes to production to avoid discovering bugs in production.
