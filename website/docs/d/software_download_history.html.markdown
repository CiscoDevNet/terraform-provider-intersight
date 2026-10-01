---
subcategory: "software"
layout: "intersight"
page_title: "Intersight: intersight_software_download_history"
description: |-
        The DownloadHistories object tracks software download operations from the Private Appliance portal.
        #### Purpose
        It provides an audit trail of software downloads, including product details, versions, and user information.
        #### Key Concepts
        - **Audit Trail:** Records all software download activities for accountability.
        - **Usage Tracking:** Tracks which software versions were downloaded by which users.

---

# Data Source: intersight_software_download_history
The DownloadHistories object tracks software download operations from the Private Appliance portal.
#### Purpose
It provides an audit trail of software downloads, including product details, versions, and user information.
#### Key Concepts
- **Audit Trail:** Records all software download activities for accountability.
- **Usage Tracking:** Tracks which software versions were downloaded by which users.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_software_download_history.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of software which was downloaded. 
* `product`:(string) The product type of the downloaded software. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timestamp`:(string) The download time of the software image. 
* `user_id_or_email`:(string) The email id of the user who initiated the software download. 
* `nr_version`:(string) The version of software which was downloaded. 
 
