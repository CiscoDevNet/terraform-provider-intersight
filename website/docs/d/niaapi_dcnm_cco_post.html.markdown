---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_cco_post"
description: |-
        The DcnmCcoPost object is integral to software release management, conveying information about new software versions specific to the DCNM platform. It ensures users are informed about the latest releases.
        #### Purpose
        DcnmCcoPost provides comprehensive details about new software releases on the DCNM platform, including the title, description, and links to download the release.
        #### Key Concepts
        - **Platform-Specific Communication:** Serves as a notice board for DCNM software releases.
        - **Structured Information:** Contains detailed content about DCNM releases.
        - **System Integration:** Supports integration with systems for DCNM updates.

---

# Data Source: intersight_niaapi_dcnm_cco_post
The DcnmCcoPost object is integral to software release management, conveying information about new software versions specific to the DCNM platform. It ensures users are informed about the latest releases.
#### Purpose
DcnmCcoPost provides comprehensive details about new software releases on the DCNM platform, including the title, description, and links to download the release.
#### Key Concepts 
- **Platform-Specific Communication:** Serves as a notice board for DCNM software releases.
- **Structured Information:** Contains detailed content about DCNM releases.
- **System Integration:** Supports integration with systems for DCNM updates.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_cco_post.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `post_date`:(string) The date when this new release notice is posted. 
* `post_type`:(string) The document type of this post. 
* `postid`:(string) Identificator of this inbox post. 
* `revision`:(string) Revision number of this notice. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
