
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_network_policy"
sidebar_current: "docs-intersight-data-source-fabric-fc-network-policy"
description: |-
A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.
---

# Data Source: intersight_fabric._fc_network_policy
A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `enable_trunking`:(bool) Enable or Disable Trunking on all of configured FC uplink ports. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
