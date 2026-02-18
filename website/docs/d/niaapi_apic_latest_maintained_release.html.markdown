---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_latest_maintained_release"
description: |-
        The ApicLatestMaintainedRelease object is crucial in software lifecycle management, focusing on maintained releases for the APIC platform. It ensures users have access to stable and supported software versions.
        #### Purpose
        ApicLatestMaintainedRelease provides information on the latest maintained software releases for the APIC platform, ensuring users are informed about stable and supported versions.
        #### Key Concepts
        - **Platform-Specific Maintenance:** Focuses on maintained releases for the APIC platform, supporting system stability.
        - **Version Continuity:** Ensures users can rely on maintained releases for consistent operation.
        - **Release Updates:** Keeps users informed about updates to maintained software versions, promoting reliability.

---

# Data Source: intersight_niaapi_apic_latest_maintained_release
The ApicLatestMaintainedRelease object is crucial in software lifecycle management, focusing on maintained releases for the APIC platform. It ensures users have access to stable and supported software versions.
#### Purpose
ApicLatestMaintainedRelease provides information on the latest maintained software releases for the APIC platform, ensuring users are informed about stable and supported versions.
#### Key Concepts
- **Platform-Specific Maintenance:** Focuses on maintained releases for the APIC platform, supporting system stability.
- **Version Continuity:** Ensures users can rely on maintained releases for consistent operation.
- **Release Updates:** Keeps users informed about updates to maintained software versions, promoting reliability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_apic_latest_maintained_release.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `maintained_release`:(string) Lastest maintained release. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `software_release`:(string) Software release version string. 
* `version_tag`:(string) Long lived version or short lived version. 
 
