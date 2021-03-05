---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_smartlicense_token"
description: |-
  SmartlicenseToken collection stores license registration tokens.
---

# Data Source: intersight_license_smartlicense_token
SmartlicenseToken collection stores license registration tokens.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_smartlicense_token.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `token`:(string) Smart license registration token. 
 