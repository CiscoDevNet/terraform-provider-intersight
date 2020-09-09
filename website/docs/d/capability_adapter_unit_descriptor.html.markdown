
---
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_unit_descriptor"
sidebar_current: "docs-intersight-data-source-capability-adapter-unit-descriptor"
description: |-
Descriptor that uniquely identifies an adaptor.
---

# Data Source: intersight_capability._adapter_unit_descriptor
Descriptor that uniquely identifies an adaptor.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connectivity_order`:(string) Order in which the ports are connected; sequential or cyclic. 
* `description`:(string) Detailed information about the endpoint. 
* `ethernet_port_speed`:(int) The port speed for ethernet ports in Mbps. 
* `fibre_channel_port_speed`:(int) The port speed for fibre channel ports in Mbps. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_dce_ports`:(int) Number of Dce Ports for the adaptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `prom_card_type`:(string) Prom card type for the adaptor. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
