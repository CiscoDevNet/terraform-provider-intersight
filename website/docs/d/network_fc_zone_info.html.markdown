
---
layout: "intersight"
page_title: "Intersight: intersight_network_fc_zone_info"
sidebar_current: "docs-intersight-data-source-network-fc-zone-info"
description: |-
FC Zone information of a Fabric Interconnect.
---

# Data Source: intersight_network._fc_zone_info
FC Zone information of a Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `user_zone_count`:(int) The number of Fibre Channel user zones defined on a Fabric Interconnect. 
* `user_zone_limit`:(int) The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect. 
* `zone_count`:(int) The number of Fibre Channel zones defined on a Fabric Interconnect. 
* `zone_limit`:(int) The maximum number of Fibre Channel zones allowed on a Fabric Interconnect. 
