
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `hw40_g_port_group_cap`:(bool) Indicates support for 40G port group capability. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pgtype`:(string) The type of port group for which this capability is defined. 
