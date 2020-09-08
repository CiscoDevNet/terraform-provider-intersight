
---
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_slot_array"
sidebar_current: "docs-intersight-resource-capability-equipment-slot-array"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability._equipment_slot_array
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `first_index`:(float) First Index information for a Switch/Fabric-Interconnect hardware. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect hardware. 
* `horizontal_start_offset`:(float) Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_separation`:(float) Inline Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_size`:(float) Inline Group Size information for a Switch/Fabric-Interconnect hardware. 
* `inline_offset`:(float) Inline Offset information for a Switch/Fabric-Interconnect hardware. 
* `location`:(string) Location information for a Switch/Fabric-Interconnect hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `number_of_slots`:(int) Number of Slots information for a Switch/Fabric-Interconnect hardware. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `orientation`:(string) Orientation information for a Switch/Fabric-Interconnect hardware. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect. 
* `section`:(Array with Maximum of one item) - A reference to a capabilitySection resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `selector`:(string) Selector information for a Switch/Fabric-Interconnect hardware. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `slots_per_line`:(int) Slots per line information for a Switch/Fabric-Interconnect hardware. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `transverse_group_separation`:(float) Transverse Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `transverse_group_size`:(float) Transverse Group Size information for a Switch/Fabric-Interconnect hardware. 
* `transverse_offset`:(float) Transverse Offset information for a Switch/Fabric-Interconnect hardware. 
* `vertical_start_offset`:(float) Vertical Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect hardware. 
