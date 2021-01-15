---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_mode"
description: |-
  Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
---

# Resource: intersight_fabric_port_mode
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `custom_mode`:(string) Custom Port Mode specified for the port range.* `FibreChannel` - Fibre Channel Port Types.* `BreakoutEthernet10G` - Breakout Ethernet 10G Port Type.* `BreakoutEthernet25G` - Breakout Ethernet 25G Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `port_id_end`:(int) Ending range of the Port Identifier. 
* `port_id_start`:(int) Starting range of the Port Identifier. 
* `port_policy`:(Array with Maximum of one item) - A reference to a fabricPortPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `slot_id`:(int) Slot Identifier of the switch. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_fabric_port_mode` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_port_mode.example 1234567890987654321abcde
```