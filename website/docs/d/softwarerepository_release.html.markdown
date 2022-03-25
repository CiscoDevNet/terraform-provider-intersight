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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_release.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `release_date`:(string) The date when the file was released or distributed by its vendor. 
* `release_notes_url`:(string) The URL for the release notes of this image. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware.* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.* `ComputeSystem` - The images in a release that correspond to servers. 
* `nr_version`:(string) Cisco provided release version. 
 
