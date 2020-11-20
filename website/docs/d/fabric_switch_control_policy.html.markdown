
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_control_policy"
sidebar_current: "docs-intersight-data-source-fabric-switch-control-policy"
description: |-
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
---

# Data Source: intersight_fabric._switch_control_policy
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `vlan_port_optimization_enabled`:(bool) To enable or disable the VLAN port count optimization. 
