---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_file_item"
description: |-
        FileItems represent files stored in a server’s local storage repository. They provide an inventory view of file artifacts (for example ISO images or CSV files), including basic metadata such as size, type, visibility/mapping to the host, and the time the file was uploaded or last updated.
        #### Purpose
        Expose a read-only inventory of local storage files on servers so administrators can discover available artifacts, validate what is present on local storage, and support workflows that depend on locally stored files (such as mounting media).
        #### Key Concepts
        - **Local storage file inventory**: Models individual file artifacts present on server-local storage.
        - **File identity and metadata**: Captures identifiers and descriptive attributes such as `fileId`, `name`, `description`, `type`, and `size`.
        - **Host mapping visibility**: `hostVisible` indicates whether the file is mapped/visible for host-side use.
        - **Recency tracking**: `updateTime` helps determine when the file was uploaded/updated.
        - **Device association**: `registeredDevice` links the file inventory to the specific managed server/device in Intersight.
        - **Permission inheritance via storage item**: Inherits permissions from `storageItem`, aligning access with the parent local storage context.

---

# Data Source: intersight_storage_file_item
FileItems represent files stored in a server’s local storage repository. They provide an inventory view of file artifacts (for example ISO images or CSV files), including basic metadata such as size, type, visibility/mapping to the host, and the time the file was uploaded or last updated.
#### Purpose
Expose a read-only inventory of local storage files on servers so administrators can discover available artifacts, validate what is present on local storage, and support workflows that depend on locally stored files (such as mounting media).
#### Key Concepts
- **Local storage file inventory**: Models individual file artifacts present on server-local storage.
- **File identity and metadata**: Captures identifiers and descriptive attributes such as `fileId`, `name`, `description`, `type`, and `size`.
- **Host mapping visibility**: `hostVisible` indicates whether the file is mapped/visible for host-side use.
- **Recency tracking**: `updateTime` helps determine when the file was uploaded/updated.
- **Device association**: `registeredDevice` links the file inventory to the specific managed server/device in Intersight.
- **Permission inheritance via storage item**: Inherits permissions from `storageItem`, aligning access with the parent local storage context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_file_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the local Storage File. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_id`:(int) The Id of the local Storage File. 
* `host_visible`:(bool) The mapping visibility of the local Storage File. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the local Storage. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(string) Total size of the local Storage File. 
* `type`:(string) File type like CSV, ISO image. 
* `update_time`:(string) Timestamp to indicate the uploaded time for this file. 
 
