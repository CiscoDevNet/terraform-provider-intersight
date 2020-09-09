
---
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_physical_def"
sidebar_current: "docs-intersight-resource-capability-equipment-physical-def"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability._equipment_physical_def
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `depth`:(float) Depth information for a Switch/Fabric-Interconnect. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect. 
* `max_power`:(int) Max Power information for a Switch/Fabric-Interconnect. 
* `min_power`:(int) Min Power information for a Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `nominal_power`:(int) Nominal Power information for a Switch/Fabric-Interconnect. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect. 
* `section`:(Array with Maximum of one item) - A reference to a capabilitySection resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `weight`:(float) Weight information for a Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect. 
