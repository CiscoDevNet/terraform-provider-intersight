
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `transport`:(string) Type of port group. Values are Eth or Fc. 
