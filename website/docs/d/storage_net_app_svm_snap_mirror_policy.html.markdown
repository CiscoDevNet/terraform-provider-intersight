---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_svm_snap_mirror_policy"
description: |-
        NetApp SnapMirror policy owned by a storage virtual machine. NetApp SnapMirror policy when applied to a SnapMirror relationship, controls the behavior of the relationship and specifies the configuration attributes for that relationship.

---

# Data Source: intersight_storage_net_app_svm_snap_mirror_policy
NetApp SnapMirror policy owned by a storage virtual machine. NetApp SnapMirror policy when applied to a SnapMirror relationship, controls the behavior of the relationship and specifies the configuration attributes for that relationship.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_svm_snap_mirror_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `comment`:(string) Comment associated with the policy. 
* `copy_all_source_snapshots`:(bool) Specifies whether all source Snapshot copies should be copied to the destination on a transfer rather than specifying specific retentions. It is applicable only to async policies. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NetApp SnapMirror policy. 
* `scope`:(string) Identifies whether the SnapMirror policy is owned by the storage virtual machine or the cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_name`:(string) The storage virtual machine name for the policy. 
* `sync_type`:(string) SnapMirror policy sync_type is either sync, strict_sync, or automated_failover. Property is applicable only to the policies of type \ sync\ . 
* `transfer_schedule_name`:(string) Name of the schedule used to update asynchronous relationships. 
* `transfer_schedule_uuid`:(string) Uuid of the schedule used to update asynchronous relationships. 
* `type`:(string) SnapMirror policy type can be async, sync, or continuous. 
* `uuid`:(string) Uuid of the NetApp SnapMirror policy. 
 
