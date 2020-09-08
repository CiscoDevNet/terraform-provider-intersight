
---
layout: "intersight"
page_title: "Intersight: intersight_sdwan_router_node"
sidebar_current: "docs-intersight-data-source-sdwan-router-node"
description: |-
Configuration settings for a SDWAN vEdge router.
---

# Data Source: intersight_sdwan._router_node
Configuration settings for a SDWAN vEdge router.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_template`:(string) Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the router node object. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `uuid`:(string) Uniquely identifies the router by its chassis number. 
