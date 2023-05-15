---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cifs_share"
description: |-
        NetApp CIFS share is a named access point in a volume which is tied to the CIFS server on the SVM.

---

# Data Source: intersight_storage_net_app_cifs_share
NetApp CIFS share is a named access point in a volume which is tied to the CIFS server on the SVM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cifs_share.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `comment`:(string) Description of the CIFS share. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `encryption`:(string) Indicates that SMB encryption must be used when accessing the share. 
* `home_directory`:(string) Indicates that the share is a home directory share, where the share and path names are dynamic. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NetApp CIFS share. 
* `path`:(string) The fully-qualified pathname in the owning SVM namespace that is shared through the share. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_name`:(string) The storage virtual machine name for the CIFS share. 
* `svm_uuid`:(string) Unique identifier for the NetApp Storage Virtual Machine. 
 
