---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_puv_script_downloader"
description: |-
        The PuvScriptDownloader object is essential for script management, providing presigned URLs for downloading metadata scripts from servers. It supports efficient access to necessary scripts.
        #### Purpose
        PuvScriptDownloader offers presigned URLs for secure and efficient script downloads, aiding users in accessing required scripts for their systems.
        #### Key Concepts
        - **Secure Script Access:** Provides presigned URLs for script downloads, ensuring secure and reliable access.
        - **Efficient Retrieval:** Facilitates quick and efficient access to scripts, supporting user needs and system requirements.
        - **System Compatibility:** Ensures seamless integration with systems for efficient script management and download.

---

# Data Source: intersight_niaapi_puv_script_downloader
The PuvScriptDownloader object is essential for script management, providing presigned URLs for downloading metadata scripts from servers. It supports efficient access to necessary scripts.
#### Purpose
PuvScriptDownloader offers presigned URLs for secure and efficient script downloads, aiding users in accessing required scripts for their systems.
#### Key Concepts
- **Secure Script Access:** Provides presigned URLs for script downloads, ensuring secure and reliable access.
- **Efficient Retrieval:** Facilitates quick and efficient access to scripts, supporting user needs and system requirements.
- **System Compatibility:** Ensures seamless integration with systems for efficient script management and download.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_puv_script_downloader.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) Filename of this Metadata script file, folder will be handled by api. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presigned_url`:(string) The presigned URL from server to download this script. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
