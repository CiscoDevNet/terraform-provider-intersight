---
subcategory: "sdwan"
layout: "intersight"
page_title: "Intersight: intersight_sdwan_router_node"
description: |-
  Configuration settings for a SDWAN vEdge router.
---

# Data Source: intersight_sdwan_router_node
Configuration settings for a SDWAN vEdge router.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdwan_router_node.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_template`:(string) Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the router node object. 
* `uuid`:(string) Uniquely identifies the router by its chassis number. 
 