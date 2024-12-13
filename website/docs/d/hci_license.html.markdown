---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_license"
description: |-
        A license instance reported by a Prism Central. A license can be consumed by multiple clusters.

---

# Data Source: intersight_hci_license
A license instance reported by a Prism Central. A license can be consumed by multiple clusters.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_license.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) The category of a license instance. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiry_date`:(string) The expiry date of a license instance. 
* `ext_id`:(string) The unique identifier of a license. 
* `license_id`:(string) The identifier of a license, usually in LIC-xxxx format. 
* `meter`:(string) A license capacity is expressed in terms of meter. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of a license instance. 
* `pc_ext_id`:(string) The unique identifier of the domain manager (Prism Central) instance which manages this cluster. 
* `quantity`:(float) The scope defines where licenses can be applied. 
* `scope`:(string) The scope defines where licenses can be applied. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sub_category`:(string) The subCategory of a license instance, such as if it is an add-on orwith unlimited capacity. 
* `type`:(string) The type of a license instance. 
 
