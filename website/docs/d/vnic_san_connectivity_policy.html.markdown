
---
layout: "intersight"
page_title: "Intersight: intersight_vnic_san_connectivity_policy"
sidebar_current: "docs-intersight-data-source-vnic-san-connectivity-policy"
description: |-
SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
---

# Data Source: intersight_vnic._san_connectivity_policy
SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `placement_mode`:(string) The mode used for placement of vNICs on network adapters. It can either be Auto or Custom.* `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.* `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
