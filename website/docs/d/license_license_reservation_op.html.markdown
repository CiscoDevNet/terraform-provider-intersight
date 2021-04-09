---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_license_reservation_op"
description: |-
  Customer operation object to request reservation code.
---

# Data Source: intersight_license_license_reservation_op
Customer operation object to request reservation code.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_license_reservation_op.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auth_code`:(string) Revervation code used to install the license. 
* `auth_code_installed`:(bool) Flag to indicate whether authorization code is installed. 
* `confirm_code`:(string) Confirm code used to complete the license update on smart license account. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `generate_request_code`:(bool) Trigger the generation of request code for specific license reservation. 
* `generate_return_code`:(bool) Trigger the generation of return code for specific license reservation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `request_code`:(string) Revervation code used to generate authorization code from CSSM. 
* `return_code`:(string) Return code used to return the reserved license to smart license account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 