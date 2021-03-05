---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_config_policy"
description: |-
  An Adapter Configuration Policy configures the Ethernet and Fibre-Channel settings for the VIC adapter.
---

# Data Source: intersight_adapter_config_policy
An Adapter Configuration Policy configures the Ethernet and Fibre-Channel settings for the VIC adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_config_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 