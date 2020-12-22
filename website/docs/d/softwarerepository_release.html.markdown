---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_release"
description: |-
  A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
---

# Data Source: intersight_softwarerepository_release
A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `release_notes_url`:(string) The URL for the release notes of this image. 
* `type`:(string) The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware.* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.* `ComputeSystem` - The images in a release that correspond to servers. 
* `nr_version`:(string) Cisco provided release version. 
