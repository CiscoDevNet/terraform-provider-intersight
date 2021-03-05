---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_license_info"
description: |-
  Describes about license info currently available in UCSD.
---

# Data Source: intersight_iaas_license_info
Describes about license info currently available in UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_license_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `license_expiration_date`:(string) UCS Director license expiration date. 
* `license_type`:(string) License type of UCSD whether it is EVAL/Permanent/Subscription.. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 