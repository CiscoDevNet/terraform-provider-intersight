---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_fex_support_meta"
description: |-
        The FexSupportMeta object provides internal metadata used to manage Fabric Extender (FEX) firmware upgrade compatibility.
        #### Purpose
        It enables the system to identify which FEX models are supported or unsupported during firmware operations, helping to block or allow domain upgrades based on connected FEX hardware.
        #### Key Concepts
        - **Compatibility Rules:** Defines series and models of FEX hardware that are subject to specific firmware upgrade constraints.
        - **Operational Safety:** Prevents incompatible firmware operations when specific FEX models are detected.

---

# Data Source: intersight_capability_fex_support_meta
The FexSupportMeta object provides internal metadata used to manage Fabric Extender (FEX) firmware upgrade compatibility.
#### Purpose
It enables the system to identify which FEX models are supported or unsupported during firmware operations, helping to block or allow domain upgrades based on connected FEX hardware.
#### Key Concepts
- **Compatibility Rules:** Defines series and models of FEX hardware that are subject to specific firmware upgrade constraints.
- **Operational Safety:** Prevents incompatible firmware operations when specific FEX models are detected.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_fex_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Information related to the list of FEXs. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series_id`:(string) Series names of FEXs which will be unsupported in the firmware operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
