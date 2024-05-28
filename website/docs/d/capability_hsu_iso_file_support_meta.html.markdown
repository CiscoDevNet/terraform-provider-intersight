---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_hsu_iso_file_support_meta"
description: |-
        Internal meta-data to check the HSU support to accept an iso file path in the upgrade request.

---

# Data Source: intersight_capability_hsu_iso_file_support_meta
Internal meta-data to check the HSU support to accept an iso file path in the upgrade request.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_hsu_iso_file_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_file_name`:(string) Name of the symbolic link to the actual iso file in the default case. 
* `default_min_version`:(string) Firmware version from which the HSU capability is present in the default case. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series`:(string) Series name for the capability group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
