
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_release"
sidebar_current: "docs-intersight-data-source-softwarerepository-release"
description: |-
A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
---

# Data Source: intersight_softwarerepository._release
A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `release_notes_url`:(string) The URL for the release notes of this image. 
* `type`:(string) The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware.* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.* `ComputeSystem` - The images in a release that correspond to servers. 
* `version`:(string) Cisco provided release version. 
