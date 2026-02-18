---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_file_downloader"
description: |-
        The FileDownloader object is pivotal in file management, providing presigned URLs for downloading metadata files from servers. It ensures users have efficient access to required data.
        #### Purpose
        FileDownloader offers presigned URLs for secure and efficient download of metadata files, supporting users in accessing necessary information for their systems.
        #### Key Concepts
        - **Secure Access:** Provides presigned URLs for metadata file downloads, ensuring secure and reliable access.
        - **Efficient Download:** Facilitates the retrieval of metadata files, supporting users in accessing required data quickly.
        - **System Integration:** Ensures seamless integration with systems for efficient data management and access.

---

# Data Source: intersight_niaapi_file_downloader
The FileDownloader object is pivotal in file management, providing presigned URLs for downloading metadata files from servers. It ensures users have efficient access to required data.
#### Purpose
FileDownloader offers presigned URLs for secure and efficient download of metadata files, supporting users in accessing necessary information for their systems.
#### Key Concepts
- **Secure Access:** Provides presigned URLs for metadata file downloads, ensuring secure and reliable access.
- **Efficient Download:** Facilitates the retrieval of metadata files, supporting users in accessing required data quickly.
- **System Integration:** Ensures seamless integration with systems for efficient data management and access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_file_downloader.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) Filename of this Metadata package file, folder will be handled by api. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presigned_url`:(string) The presigned URL from server to download this file. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
