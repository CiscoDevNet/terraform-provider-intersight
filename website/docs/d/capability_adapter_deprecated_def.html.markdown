---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_deprecated_def"
description: |-
        The capability.AdapterDeprecatedDef object is a capability definition used to identify network adapters that are unsupported or have been deprecated in the current system release. It is typically used within a server descriptor to flag incompatible hardware.
        #### Purpose
        The primary purpose of this object is to maintain a catalog of unsupported adapters. When a server is discovered, its adapters can be checked against this catalog. If a match is found, the system can raise an alarm or prevent certain operations, ensuring that only supported hardware configurations are used.
        #### Key Concepts
        - **Hardware Compatibility Catalog:** Acts as a list of known unsupported or deprecated adapters.
        - **Identification by Vendor and Model:** Uses vendor and model properties to uniquely identify an adapter type.
        - **Proactive Issue Detection:** Enables the system to flag potential support and functionality issues by identifying deprecated hardware during server inventory.
        - **Nested Definition:** Designed to be used as a sub-object within a capability.ServerDescriptor, linking unsupported adapters to specific server models where they are not applicable.

---

# Data Source: intersight_capability_adapter_deprecated_def
The capability.AdapterDeprecatedDef object is a capability definition used to identify network adapters that are unsupported or have been deprecated in the current system release. It is typically used within a server descriptor to flag incompatible hardware.
#### Purpose
The primary purpose of this object is to maintain a catalog of unsupported adapters. When a server is discovered, its adapters can be checked against this catalog. If a match is found, the system can raise an alarm or prevent certain operations, ensuring that only supported hardware configurations are used.
#### Key Concepts
- **Hardware Compatibility Catalog:** Acts as a list of known unsupported or deprecated adapters.
- **Identification by Vendor and Model:** Uses vendor and model properties to uniquely identify an adapter type.
- **Proactive Issue Detection:** Enables the system to flag potential support and functionality issues by identifying deprecated hardware during server inventory.
- **Nested Definition:** Designed to be used as a sub-object within a capability.ServerDescriptor, linking unsupported adapters to specific server models where they are not applicable.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_adapter_deprecated_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the unsupported adapter. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) Vendor of the unsupported adapter. 
 
