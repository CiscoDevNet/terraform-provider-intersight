---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vsan"
description: |-
        The Vsan object represents a VSAN configuration definition used in Fibre Channel network policy contexts. It models SAN segmentation intent and associated VSAN identity.
        #### Purpose
        Vsan provides a policy-driven way to define VSANs that should exist and be configured in the SAN fabric. It enables consistent SAN segmentation and supports validation/deployment workflows for FC-related configuration.
        #### Key Concepts
        - **SAN segmentation intent:** Encodes VSAN identity and intended behavior for FC fabrics.
        - **Policy association:** Exists under an FC network policy that governs SAN configuration.
        - **Reusable configuration unit:** Supports repeating the same VSAN definitions across profiles/domains.
        - **Lifecycle through workflows:** Designed for validation and deployment through fabric workflows.

---

# Data Source: intersight_fabric_vsan
The Vsan object represents a VSAN configuration definition used in Fibre Channel network policy contexts. It models SAN segmentation intent and associated VSAN identity.
#### Purpose
Vsan provides a policy-driven way to define VSANs that should exist and be configured in the SAN fabric. It enables consistent SAN segmentation and supports validation/deployment workflows for FC-related configuration.
#### Key Concepts
- **SAN segmentation intent:** Encodes VSAN identity and intended behavior for FC fabrics.
- **Policy association:** Exists under an FC network policy that governs SAN configuration.
- **Reusable configuration unit:** Supports repeating the same VSAN definitions across profiles/domains.
- **Lifecycle through workflows:** Designed for validation and deployment through fabric workflows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_vsan.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_zoning`:(string) Enables or Disables the default zoning state.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fc_zone_sharing_mode`:(string) Logical grouping mode for fc ports. 
* `fcoe_vlan`:(int) FCOE Vlan associated to the VSAN configuration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User given name for the VSAN configuration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vsan_id`:(int) Virtual San Identifier in the switch. 
* `vsan_scope`:(string) Used to indicate whether the VSAN Id is defined for storage or uplink or both traffics in FI.* `Uplink` - Vsan associated with uplink network.* `Storage` - Vsan associated with storage network.* `Common` - Vsan that is common for uplink and storage network. 
 
