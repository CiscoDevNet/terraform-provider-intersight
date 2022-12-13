---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_qtree"
description: |-
        NetApp qtree is a logically defined file system that can exist as a special subdirectory of the root directory within a volume.

---

# Data Source: intersight_storage_net_app_qtree
NetApp qtree is a logically defined file system that can exist as a special subdirectory of the root directory within a volume.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_qtree.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `export_policy_id`:(string) Unique identifier of NetApp export policy. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NetApp Qtree. 
* `path`:(string) Client visible path to the qtree. 
* `permission`:(string) Identifies the UNIX permissions for the qtree. 
* `qtree_id`:(int) NetApp Qtree ID, unique within the qtree's volume. 
* `security_style`:(string) Identifies the security style for the qtree, it determines how access to the qtree is controlled.* `UNIX` - Security style for UNIX uid, gid and mode bits.* `NTFS` - Security style for CIFS ACLs.* `Mixed` - Security style for NFS and CIFS access. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `volume_uuid`:(string) NetApp Volume uuid, unique identifier for the NetApp volume. 
 
