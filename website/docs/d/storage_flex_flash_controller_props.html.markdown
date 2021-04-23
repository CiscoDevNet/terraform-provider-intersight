---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_flex_flash_controller_props"
description: |-
  Flex flash controller properties.
---

# Data Source: intersight_storage_flex_flash_controller_props
Flex flash controller properties.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_flex_flash_controller_props.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cards_manageable`:(string) Manageable card on the flex flash controller. 
* `configured_mode`:(string) Mode configured on the flex flash controller. 
* `controller_name`:(string) The current name of the flex flash controller. 
* `controller_status`:(string) The current status of the flex flash controller. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fw_version`:(string) Firmware version of the flex flash controller. 
* `internal_state`:(string) Internal state of the flex flash controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operating_mode`:(string) Operating mode of flex flash controller. 
* `physical_drive_count`:(string) Number of connected physical drives to a specific Flex flash controller. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `product_name`:(string) Product name of the flex flash controller. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `startup_fw_version`:(string) Startup firmware version of the Flex flash controller. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `virtual_drive_count`:(string) Number of virtual drives for a specific Flex flash controller. 
 