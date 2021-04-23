---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_eula"
description: |-
  End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
---

# Data Source: intersight_firmware_eula
End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_eula.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accepted`:(bool) EULA acceptance status for the account. 
* `account_moid`:(string) The Account ID for this managed object. 
* `content`:(string) EULA acceptance form content provided by cisco.com. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 