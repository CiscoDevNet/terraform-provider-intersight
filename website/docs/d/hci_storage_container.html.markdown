---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_storage_container"
description: |-
        Storage containers in Nutanix serve as logical abstraction layers for virtual disk storage which can be backed by HCI (local) storage or external storage provided by vendors such as Pure Storage.
        #### Purpose
        The StorageContainer object provides a view of both HCI and external storage containers in a Nutanix cluster, enabling administrators to monitor storage configuration and capacity.
        #### Key Concepts
        - **Storage Abstraction:** Logical containers for virtual disk storage that can use local or external storage backends.
        - **Capacity Management:** Tracks physical and logical capacity, reservations, and utilization metrics.
        - **Data Efficiency:** Supports compression, deduplication, and erasure coding for optimized storage usage.

---

# Data Source: intersight_hci_storage_container
Storage containers in Nutanix serve as logical abstraction layers for virtual disk storage which can be backed by HCI (local) storage or external storage provided by vendors such as Pure Storage.
#### Purpose
The StorageContainer object provides a view of both HCI and external storage containers in a Nutanix cluster, enabling administrators to monitor storage configuration and capacity.
#### Key Concepts
- **Storage Abstraction:** Logical containers for virtual disk storage that can use local or external storage backends.
- **Capacity Management:** Tracks physical and logical capacity, reservations, and utilization metrics.
- **Data Efficiency:** Supports compression, deduplication, and erasure coding for optimized storage usage.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_storage_container.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cache_deduplication`:(string) Current status of Cache Deduplication for the Storage Container. Possible values: ON, OFF. Does not apply to external storage container case. 
* `cluster_ext_id`:(string) The external identifier of the cluster owning the Storage Container. 
* `cluster_name`:(string) The corresponding name of the cluster owning the Storage Container instance. 
* `compression_delay_secs`:(int) The compression delay in seconds. Does not apply to external storage container case. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `erasure_code`:(string) Current status value of Erasure Coding for the Storage Container. Possible values: ON, OFF.Note: Erasure coding is available with certain licensing levels (e.g., Pro licensing or above). Does not apply to external storage container case. 
* `erasure_code_delay_secs`:(int) Delay in performing Erasure Code for the current Storage Container instance. Does not apply to external storage container case. 
* `external_storage_ext_id`:(string) External Storage extId. Presence of externalStorageExtId indicates this storage-container is an external storage container instead of a HCI storage container. 
* `has_higher_ec_fault_domain_preference`:(bool) Indicates whether to prefer a higher Erasure Code fault domain. Does not apply to external storage container case. 
* `is_compression_enabled`:(bool) Indicates whether the compression is enabled for the Storage Container. Does not apply to external storage container case. 
* `is_encrypted`:(bool) Indicates whether the Storage Container is encrypted or not. Does not apply to external storage container case. 
* `is_external_storage`:(bool) Derived value to indicate if the container is backed by an external storage or not. The value is true when the externalStorageExtId is not empty. 
* `is_inline_ec_enabled`:(bool) Indicates whether data written to this Storage Container should be inline erasure-coded or not. This field is only considered if ErasureCoding is enabled. Inline Erasure-Coded refers to a data protection method where erasure coding is applied directly and immediately to data as it is written (in-line) to the storage system. Does not apply to external storage container case. 
* `is_internal`:(bool) Indicates whether the Storage Container is internal and is managed by Nutanix. 
* `is_marked_for_removal`:(bool) Indicates whether the Storage Container is marked for removal. This field is set when the Storage Container is about to be destroyed. 
* `is_nfs_allowlist_inherited`:(bool) Indicates whether the NFS allowlist is inherited from the global configuration. Does not apply to external storage container case. 
* `is_shared`:(bool) Indicates whether the Storage Container is shared. When shared, all PEs registered under the same PC can access and utilize its storage. Once the field is set during the creation, it is immutable except in the case of SelfServiceContainer. Even for SelfServiceContainer only FALSE to TRUE is allowed. Does not apply to external storage container case. 
* `is_software_encryption_enabled`:(bool) Indicates whether the Storage Container instance has software encryption enabled. Does not apply to external storage container case. 
* `logical_advertised_capacity_bytes`:(int) Maximum capacity of the Storage Container as defined by the user. Does not apply to external storage container case. 
* `logical_explicit_reserved_capacity_bytes`:(int) Total reserved size (in bytes) of the Storage Container (set by Admin). This also includes the replication factor of the Storage Container. The actual reserved capacity of the Storage Container will be the maximum of explicitReservedCapacity and implicitReservedCapacity. Does not apply to external storage container case. 
* `logical_implicit_reserved_capacity_bytes`:(int) Sum of the reservations provisioned on all vDisks in the Storage Container. The actual reserved capacity of the Storage Container will be the maximum of explicitReservedCapacity and implicitReservedCapacity. Does not apply to external storage container case. 
* `max_capacity_bytes`:(int) Maximum physical capacity of the Storage Container in bytes. Does not apply to external storage container case. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Storage Container. Prism Central requires that the name of Storage Container should be unique in every registered cluster. 
* `on_disk_dedup`:(string) Current status of Disk deduplication for the Storage Container. Possible values:- POST_PROCESS: Deduplication is enabled for the Storage Container instance.- OFF: Deduplication is disabled for the Storage Container instance.Does not apply to external storage container case. 
* `replication_factor`:(int) Replication factor of the Storage Container. Does not apply to external storage container case. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_container_ext_id`:(string) The external identifier of the Storage Container. 
* `storage_pool_ext_id`:(string) The external identifier of the Storage Pool owning the Storage Container instance. Does not apply to external storage container case. 
 
