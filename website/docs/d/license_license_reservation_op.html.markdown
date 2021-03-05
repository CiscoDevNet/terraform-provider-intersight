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
* `auth_code`:(string) Revervation code used to install the license. 
* `auth_code_installed`:(bool) Flag to indicate whether authorization code is installed. 
* `confirm_code`:(string) Confirm code used to complete the license update on smart license account. 
* `generate_request_code`:(bool) Trigger the generation of request code for specific license reservation. 
* `generate_return_code`:(bool) Trigger the generation of return code for specific license reservation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `request_code`:(string) Revervation code used to generate authorization code from CSSM. 
* `return_code`:(string) Return code used to return the reserved license to smart license account. 
 