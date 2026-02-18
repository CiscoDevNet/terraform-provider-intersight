---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_release_recommend"
description: |-
        The DcnmReleaseRecommend object is pivotal in software recommendation, offering guidance on optimal software versions for the DCNM platform. It helps users ensure compatibility and maximize performance.
        #### Purpose
        DcnmReleaseRecommend provides recommendations for software versions based on hardware identifiers for the DCNM platform, offering tailored advice for users.
        #### Key Concepts
        - **Platform-Specific Guidance:** Offers recommendations on software releases for the DCNM platform.
        - **Compatibility Assurance:** Ensures users have access to compatible versions, reducing system issues.
        - **Performance Optimization:** Guides users towards releases that enhance system efficiency.

---

# Data Source: intersight_niaapi_dcnm_release_recommend
The DcnmReleaseRecommend object is pivotal in software recommendation, offering guidance on optimal software versions for the DCNM platform. It helps users ensure compatibility and maximize performance.
#### Purpose
DcnmReleaseRecommend provides recommendations for software versions based on hardware identifiers for the DCNM platform, offering tailored advice for users.
#### Key Concepts
- **Platform-Specific Guidance:** Offers recommendations on software releases for the DCNM platform.
- **Compatibility Assurance:** Ensures users have access to compatible versions, reducing system issues.
- **Performance Optimization:** Guides users towards releases that enhance system efficiency.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_release_recommend.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cll`:(string) Current long-lived release. 
* `create_time`:(string) The time when this managed object was created. 
* `crr`:(string) Customer recommended releases. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pid`:(string) Hardware model identificator. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `ull`:(string) Upcoming long-lived release. 
 
