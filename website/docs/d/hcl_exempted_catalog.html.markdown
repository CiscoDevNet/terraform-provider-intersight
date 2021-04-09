---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_exempted_catalog"
description: |-
  Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
---

# Data Source: intersight_hcl_exempted_catalog
Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_exempted_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `comments`:(string) Reason for the exemption. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique descriptive name of the exemption. 
* `os_vendor`:(string) Vendor of the Operating System. 
* `os_version`:(string) Version of the Operating system. 
* `processor_name`:(string) Name of the processor supported for the server. 
* `product_type`:(string) Type of the product/adapter say GPU for graphic cards.* `` - Default type of the Product.* `Adapter` - Represents network adapter cards.* `StorageController` - Represents storage controllers.* `GPU` - Represents graphics cards. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `ucs_version`:(string) Version of the UCS software. 
* `version_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. 
 