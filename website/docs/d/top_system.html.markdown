
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ipv4_address`:(string) The IPv4 address of system. 
* `ipv6_address`:(string) The IPv6 address of system. 
* `mode`:(string) The current mode of the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The admin configured name of the system. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `time_zone`:(string) The operational timezone of the system, empty indicates no timezone has been set specifically. 
