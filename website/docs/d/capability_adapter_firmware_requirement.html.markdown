---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_firmware_requirement"
description: |-
        Firmware requirements for enabling Intersight based management for an adaptor.

---

# Data Source: intersight_capability_adapter_firmware_requirement
Firmware requirements for enabling Intersight based management for an adaptor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_adapter_firmware_requirement.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_series`:(string) Series of adapter. Example - Cruz, Bodega. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ignore_empty_current_version`:(bool) Do not update if the current version is reported as empty. 
* `minimum_adapter_version`:(string) The minimum adapter version that supports this adapter upgrade. 
* `minimum_bmc_version`:(string) The minimum BMC version that supports this adapter upgrade. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `recommended_bmc_version`:(string) The recommended BMC version that supports this adapter upgrade. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
