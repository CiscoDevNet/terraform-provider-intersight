---
subcategory: "management"
layout: "intersight"
page_title: "Intersight: intersight_management_controller"
description: |-
        Controllers (management) represent a management controller (service processor) that monitors server physical state via sensors and communicates out-of-band. It may also report replay-config status for Fabric Interconnect reboot replay behavior.
        #### Purpose
        Provide inventory and operational visibility for the management controller, including certificates and management interface relationships.
        #### Key Concepts
        - **Out-of-band control-plane:** Represents the service processor used for monitoring and management.
        - **Security artifacts:** Can reference certificates used for authorization and trust establishment.
        - **Replay status (FI):** Includes replay config status/timestamps for post-reboot config replay visibility.
        - **Interface inventory linkage:** Related to one or more management Interfaces that expose IP/MAC/VLAN settings.

---

# Data Source: intersight_management_controller
Controllers (management) represent a management controller (service processor) that monitors server physical state via sensors and communicates out-of-band. It may also report replay-config status for Fabric Interconnect reboot replay behavior.
#### Purpose
Provide inventory and operational visibility for the management controller, including certificates and management interface relationships.
#### Key Concepts
- **Out-of-band control-plane:** Represents the service processor used for monitoring and management.
- **Security artifacts:** Can reference certificates used for authorization and trust establishment.
- **Replay status (FI):** Includes replay config status/timestamps for post-reboot config replay visibility.
- **Interface inventory linkage:** Related to one or more management Interfaces that expose IP/MAC/VLAN settings.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_management_controller.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the endpoint that houses the management controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uem_stream_admin_state`:(string) Desired state of the UEM stream.* `Disabled` - The UEM event channel is disabled.* `Enabled` - The UEM event channel is enabled. 
 
