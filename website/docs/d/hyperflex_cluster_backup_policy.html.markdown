---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_backup_policy"
description: |-
        The ClusterBackupPolicy object is a critical component in the HyperFlex system, designed to specify backup configurations for a HyperFlex Cluster. This provides a structured approach to managing data backup operations, ensuring data integrity and availability across various deployment scenarios.
        #### Purpose
        The ClusterBackupPolicy serves as the blueprint for configuring backup operations within a HyperFlex Cluster. It facilitates the definition and management of backup settings, allowing administrators to customize and optimize backup strategies according to their specific needs.
        #### Key Concepts
        - **Policy-Based Management:** Enables centralized control and management of backup configurations, providing a consistent framework for data protection across the cluster.
        - **Flexibility:** Supports various customization options, including datastore settings, snapshot retention, and encryption, allowing tailored backup strategies.
        - **Integration:** Works seamlessly with HyperFlex Cluster profiles, ensuring that backup configurations are aligned with cluster operations and requirements.
        - **Security:** Incorporates encryption options to safeguard data, ensuring compliance with security standards and protecting against unauthorized access.

---

# Data Source: intersight_hyperflex_cluster_backup_policy
The ClusterBackupPolicy object is a critical component in the HyperFlex system, designed to specify backup configurations for a HyperFlex Cluster. This provides a structured approach to managing data backup operations, ensuring data integrity and availability across various deployment scenarios.  
#### Purpose  
The ClusterBackupPolicy serves as the blueprint for configuring backup operations within a HyperFlex Cluster. It facilitates the definition and management of backup settings, allowing administrators to customize and optimize backup strategies according to their specific needs.  
#### Key Concepts  
- **Policy-Based Management:** Enables centralized control and management of backup configurations, providing a consistent framework for data protection across the cluster. 
- **Flexibility:** Supports various customization options, including datastore settings, snapshot retention, and encryption, allowing tailored backup strategies. 
- **Integration:** Works seamlessly with HyperFlex Cluster profiles, ensuring that backup configurations are aligned with cluster operations and requirements. 
- **Security:** Incorporates encryption options to safeguard data, ensuring compliance with security standards and protecting against unauthorized access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_backup_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_data_store_name`:(string) Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up. 
* `backup_data_store_size`:(int) Replication data store size in backupDataStoreSizeUnit. 
* `backup_data_store_size_unit`:(string) Replication data store size. 
* `create_time`:(string) The time when this managed object was created. 
* `data_store_encryption_enabled`:(bool) Whether the datastore is encrypted or not. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `local_snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `replication_pair_name_prefix`:(string) Replication cluster pairing name prefix. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
 
