---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_update_constraint_meta"
description: |-
        Internal meta-data to enable adapter unit update related constraints.

---

# Data Source: intersight_capability_adapter_update_constraint_meta
Internal meta-data to enable adapter unit update related constraints.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_adapter_update_constraint_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_secure_boot_supported`:(bool) Flag to indicate support for secure boot. 
* `min_supported_version`:(string) Firmware version below which firmware update is not supported for this inventory unit. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the inventory unit which will be supported in firmware operation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `platform_type`:(string) Platform type for which the constraint is to be enforced. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `supported_platform`:(string) Platform for which the constraint is to be enforced. 
 
