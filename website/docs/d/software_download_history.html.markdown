---
subcategory: "software"
layout: "intersight"
page_title: "Intersight: intersight_software_download_history"
description: |-
  An object to keep track of software downloads from the Private Appliance portal in SaaS.
---

# Data Source: intersight_software_download_history
An object to keep track of software downloads from the Private Appliance portal in SaaS.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_software_download_history.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of software which was downloaded. 
* `product`:(string) The product type of the downloaded software. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timestamp`:(string) The download time of the software image. 
* `nr_version`:(string) The version of software which was downloaded. 
 