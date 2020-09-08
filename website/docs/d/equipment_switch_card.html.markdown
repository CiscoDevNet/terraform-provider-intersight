
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_switch_card"
sidebar_current: "docs-intersight-data-source-equipment-switch-card"
description: |-
Fixed / Removable module on a Fabric Interconnect / Switch.
---

# Data Source: intersight_equipment._switch_card
Fixed / Removable module on a Fabric Interconnect / Switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Detailed description of this switch hardware. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_ports`:(int) Number of ports present in this switch hardware. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `out_of_band_ip_address`:(string) Field specifies this Switch's Out-of-band IP address. 
* `out_of_band_ip_gateway`:(string) Field specifies this Switch's default gateway for the out-of-band management interface. 
* `presence`:(string) Presence for this switch hardware. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `slot_id`:(int) Slot identifier of the local Switch slot Interface. 
* `state`:(string) Operational state of the switch hardware. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `vendor`:(string) This field identifies the vendor of the given component. 
