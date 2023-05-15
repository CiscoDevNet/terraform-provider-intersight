---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_snap_mirror_relationship"
description: |-
        NetApp SnapMirror relationship.

---

# Data Source: intersight_storage_net_app_snap_mirror_relationship
NetApp SnapMirror relationship.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_snap_mirror_relationship.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `destination_path`:(string) Path to the destination endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `healthy`:(string) Whether the relationship is healthy or not. 
* `lag_time`:(string) Time since the exported Snapshot copy was created. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `policy_name`:(string) Name of the NetApp SnapMirror policy. 
* `policy_type`:(string) SnapMirror policy type can be async, sync, or continuous. 
* `policy_uuid`:(string) Uuid of the NetApp SnapMirror policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_path`:(string) Path to the source endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. 
* `state`:(string) State of the relationship. 
* `uuid`:(string) Uuid of the NetApp SnapMirror relationship. 
 
