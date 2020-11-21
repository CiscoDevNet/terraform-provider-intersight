
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_template`:(string) Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the router node object. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `uuid`:(string) Uniquely identifies the router by its chassis number. 
