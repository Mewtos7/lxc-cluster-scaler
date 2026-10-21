# lxc-cluster-scaler
Scaler for lxc-cluster on cloud provider environment.

## Project description
lxc-cluster-scaler is an idea to introduce automatic scaling of incus cluster on cloud provider infrastructure.
This scaler should be a free pendant to exisiting and known virtual machine scaler or kubernetes cluster scaler, to host virtual machines and container.

### Benefits
Through up- and downscaling mechanism in combination with live migration, the scaler will optimize any costs of used cloud ressources.
If more ressources are necessary, cluster will be up-scaled with new nodes/cluster members.
If ressources are not used, machines and container on nodes/cluster members are consolidated to non-empty ones and empty members are removed.
Additionally, the scaler works with multiple cloud provider, bringing in the possibility to scale server on different provider infrastructure and getting the best price.


### Use-cases
Beneficial use-case is the necessity for hosted virtual machines, that should run on an cloud infrastructure but save cost when not used or less used (e.g. temporary sandbox machines, game servers or load system).
Detailed use-case would be to provide game-servers, that could be throttled when not used na dtransferred with live-migration to another cluster member.

## Current state
In development