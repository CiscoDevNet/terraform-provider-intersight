---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_driver_image"
description: |-
        Collection used to store the driver ISO urls for each server based on how it is managed.

---

# Data Source: intersight_hcl_driver_image
Collection used to store the driver ISO urls for each server based on how it is managed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_driver_image.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `driver_iso_url`:(string) URL of the driver ISO images. 
* `management_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.* `UCSM` - The server is managed by UCS Manager.* `IMC` - The server is standalone managed by CIMC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
