---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_valid_remote_target"
description: |-
        ValidRemoteTargets provide validated remote boot/install targets for a server based on its server profile, focusing on SAN/iSCSI paths. The results enumerate reachable Fibre Channel LUNs and iSCSI LUNs that can be used as remote install targets.
        #### Purpose
        Enable OS installation workflows to discover and present valid remote installation targets (SAN boot style) derived from the server’s configured profile connectivity.
        #### Key Concepts
        - **Read-based retrieval**: Exposed via READ to fetch current valid remote target options.
        - **Remote target types**: Returns `fibreChannelLuns` (with initiator/target WWPNs and LUN ID constraints) and `iscsiLuns` (with initiator vNIC MAC and target IQN validation).
        - **Profile-derived paths**: Targets are determined from the server profile configuration associated with the server.
        - **Server association**: `server` identifies the compute.Physical for which remote targets are being reported.
        - **UI-oriented response shaping**: `src` can indicate orchestration UI usage patterns for response formatting.

---

# Data Source: intersight_os_valid_remote_target
ValidRemoteTargets provide validated remote boot/install targets for a server based on its server profile, focusing on SAN/iSCSI paths. The results enumerate reachable Fibre Channel LUNs and iSCSI LUNs that can be used as remote install targets.
#### Purpose
Enable OS installation workflows to discover and present valid remote installation targets (SAN boot style) derived from the server’s configured profile connectivity.
#### Key Concepts
- **Read-based retrieval**: Exposed via READ to fetch current valid remote target options.
- **Remote target types**: Returns `fibreChannelLuns` (with initiator/target WWPNs and LUN ID constraints) and `iscsiLuns` (with initiator vNIC MAC and target IQN validation).
- **Profile-derived paths**: Targets are determined from the server profile configuration associated with the server.
- **Server association**: `server` identifies the compute.Physical for which remote targets are being reported.
- **UI-oriented response shaping**: `src` can indicate orchestration UI usage patterns for response formatting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_valid_remote_target.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src`:(string) Flag to denote the source of the request.If the call is from Orchestration UI, only the flat list of Install targets can be sent as response. 
 
