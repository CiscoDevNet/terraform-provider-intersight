
---
layout: "intersight"
page_title: "Intersight: intersight_capability_port_group_aggregation_def"
sidebar_current: "docs-intersight-data-source-capability-port-group-aggregation-def"
description: |-
FEX/IOCARD module port group aggregation capabilities.
---

# Data Source: intersight_capability._port_group_aggregation_def
FEX/IOCARD module port group aggregation capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aggregation_cap`:(string) Aggregation capability for port group. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `hw40_g_port_group_cap`:(bool) Indicates support for 40G port group capability. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pgtype`:(string) The type of port group for which this capability is defined. 
