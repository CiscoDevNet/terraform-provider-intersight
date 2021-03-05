---
subcategory: "access"
layout: "intersight"
page_title: "Intersight: intersight_access_policy"
description: |-
  Policy to configure server or chassis management options.
---

# Data Source: intersight_access_policy
Policy to configure server or chassis management options.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_access_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `inband_vlan`:(int) VLAN to be used for server access over Inband network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 