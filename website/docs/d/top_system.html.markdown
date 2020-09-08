
---
layout: "intersight"
page_title: "Intersight: intersight_top_system"
sidebar_current: "docs-intersight-data-source-top-system"
description: |-
Root container for all UCSM / CIMC MOs.
---

# Data Source: intersight_top._system
Root container for all UCSM / CIMC MOs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ipv4_address`:(string) The IPv4 address of system. 
* `ipv6_address`:(string) The IPv6 address of system. 
* `mode`:(string) The current mode of the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The admin configured name of the system. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `time_zone`:(string) The operational timezone of the system, empty indicates no timezone has been set specifically. 
