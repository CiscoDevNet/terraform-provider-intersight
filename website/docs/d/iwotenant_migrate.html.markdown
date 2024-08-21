---
subcategory: "iwotenant"
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_migrate"
description: |-
        Api used to transfer customer's IWO data.

---

# Data Source: intersight_iwotenant_migrate
Api used to transfer customer's IWO data.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iwotenant_migrate.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password that satisfies Cisco's Password Complexity security requirement.Minimum password length, which MUST be nonzero and up to the current maximum lengthContain both upper-case and lower-case lettersContain at least one number (for example, 0-9)Contain at least one special character. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `url`:(string) S3 bucket URL for uploading the migration data. 
 
