
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_release"
sidebar_current: "docs-intersight-resource-softwarerepository-release"
description: |-
  A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
---

# Resource: intersight_softwarerepository._release
A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `catalog`:(Array with Maximum of one item) - A reference to a softwarerepositoryCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `release_notes_url`:(string) The URL for the release notes of this image. 
* `supported_models`:
                (Array of schema.TypeString) -
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `type`:(string) The platform type for which the images are released. This can be a fabric interconnect or compute server hardware.* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.* `ComputeSystem` - The images in a release that correspond to servers. 
* `version`:(string) Cisco provided release version. 
