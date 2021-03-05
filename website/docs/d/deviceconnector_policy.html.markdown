---
subcategory: "deviceconnector"
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
description: |-
  Policy to control configuration changes allowed from Cisco IMC.
---

# Data Source: intersight_deviceconnector_policy
Policy to control configuration changes allowed from Cisco IMC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_deviceconnector_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `lockout_enabled`:(bool) Enables configuration lockout on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 