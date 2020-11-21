
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
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `release_notes_url`:(string) The URL for the release notes of this image. 
* `supported_models`:
                (Array of schema.TypeString) -
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `type`:(string) The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware.* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.* `ComputeSystem` - The images in a release that correspond to servers. 
* `version`:(string) Cisco provided release version. 
