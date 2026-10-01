---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cluster_snapshot_policy"
description: |-
        The NetAppClusterSnapshotPolicies object controls the creation and retention of volume snapshots at the cluster level.
        #### Purpose
        It automates data protection by defining when and how often snapshots are taken, ensuring that recovery points are consistently maintained.
        ####  Key Concepts
        - **Snapshot Automation:** Defines intervals and schedules for snapshot creation.
        - **Retention Management:** Specifies the number of copies to maintain for each schedule.
        - **Policy Scope:** Manages policies globally across the cluster.

---

# Data Source: intersight_storage_net_app_cluster_snapshot_policy
The NetAppClusterSnapshotPolicies object controls the creation and retention of volume snapshots at the cluster level.
#### Purpose
It automates data protection by defining when and how often snapshots are taken, ensuring that recovery points are consistently maintained.
####  Key Concepts
- **Snapshot Automation:** Defines intervals and schedules for snapshot creation.
- **Retention Management:** Specifies the number of copies to maintain for each schedule.
- **Policy Scope:** Manages policies globally across the cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cluster_snapshot_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(string) Snapshot policy is enabled or not. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NetApp Snapshot Policy. 
* `scope`:(string) Identifies whether the Snapshot Policy is owned by the Storage Virtual Machine or the cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) Uuid of the NetApp Snapshot Policy. 
 
