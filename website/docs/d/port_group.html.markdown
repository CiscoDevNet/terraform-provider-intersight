
---
layout: "intersight"
page_title: "Intersight: intersight_port_group"
sidebar_current: "docs-intersight-data-source-port-group"
description: |-
Holder for multiple ports. A switch card will have one or more port groups.
---

# Data Source: intersight_port._group
Holder for multiple ports. A switch card will have one or more port groups.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `transport`:(string) Type of port group. Values are Eth or Fc. 
