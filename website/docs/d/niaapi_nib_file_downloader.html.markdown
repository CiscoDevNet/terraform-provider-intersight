---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nib_file_downloader"
description: |-
        The NibFileDownloader object is central to file management, offering presigned URLs for downloading metadata files from servers. It supports secure and efficient data retrieval.
        #### Purpose
        NibFileDownloader provides presigned URLs for metadata file downloads, ensuring users have reliable access to necessary data for their systems.
        #### Key Concepts
        - **Secure Retrieval:** Offers presigned URLs for safe and efficient metadata file downloads.
        - **Data Accessibility:** Facilitates quick and reliable access to metadata files, supporting user needs.
        - **System Integration:** Ensures compatibility with systems for efficient file management and download.

---

# Data Source: intersight_niaapi_nib_file_downloader
The NibFileDownloader object is central to file management, offering presigned URLs for downloading metadata files from servers. It supports secure and efficient data retrieval.
#### Purpose
NibFileDownloader provides presigned URLs for metadata file downloads, ensuring users have reliable access to necessary data for their systems.
#### Key Concepts
- **Secure Retrieval:** Offers presigned URLs for safe and efficient metadata file downloads.
- **Data Accessibility:** Facilitates quick and reliable access to metadata files, supporting user needs.
- **System Integration:** Ensures compatibility with systems for efficient file management and download.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_nib_file_downloader.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) Filename of this Metadata package file, folder will be handled by api. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presigned_url`:(string) The presigned URL from server to download this file. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
