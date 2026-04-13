---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_backup_cluster"
description: |-
        The BackupCluster object is central to managing backup-related operations for HyperFlex clusters, providing comprehensive information about the backup source and target clusters.
        #### Purpose
        BackupCluster is designed to offer insights into the backup processes associated with HyperFlex clusters. It outlines the relationships and UUIDs pertinent to backup operations, facilitating seamless tracking and management of backup activities.
        #### Key Concepts
        - **Cluster Associations:** Establishes connections between source and target clusters, ensuring clarity in backup operations.
        - **Protection Details:** Describes backup sources and references, aiding in the management and validation of backup integrity.
        - **Access Control:** Utilizes privilege sets to regulate access to backup information, maintaining security and data protection standards.

---

# Data Source: intersight_hyperflex_backup_cluster
The BackupCluster object is central to managing backup-related operations for HyperFlex clusters, providing comprehensive information about the backup source and target clusters.
#### Purpose
BackupCluster is designed to offer insights into the backup processes associated with HyperFlex clusters. It outlines the relationships and UUIDs pertinent to backup operations, facilitating seamless tracking and management of backup activities.
#### Key Concepts
- **Cluster Associations:** Establishes connections between source and target clusters, ensuring clarity in backup operations.
- **Protection Details:** Describes backup sources and references, aiding in the management and validation of backup integrity.
- **Access Control:** Utilizes privilege sets to regulate access to backup information, maintaining security and data protection standards.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_backup_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_data_store`:(string) Defines the backup source cluster and its references. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_cluster_uuid`:(string) UUID for the cluster to allow lookups across unclaim/reclaim. 
 
