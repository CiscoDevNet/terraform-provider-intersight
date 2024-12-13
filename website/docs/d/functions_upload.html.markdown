---
subcategory: "functions"
layout: "intersight"
page_title: "Intersight: intersight_functions_upload"
description: |-
        The managed object which has info about uploaded file.

---

# Data Source: intersight_functions_upload
The managed object which has info about uploaded file.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_functions_upload.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action against the Upload.* `None` - No action is set, this is the default value for action field.* `CompleteUploading` - Mark the instance of a Upload as uploaded. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the Upload. 
* `description`:(string) Description of the Upload. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) The file name of the Upload. File name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `file_size`:(int) The size (in bytes) of the file. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the Upload. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Upload. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `part_size`:(int) The chunk size (in bytes) for each part of the file to be uploaded. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current representation of the state of Upload.* `Uploading` - File uploading is in progress.* `Uploaded` - File uploading is completed.* `Failed` - File uploading is failed. 
 
