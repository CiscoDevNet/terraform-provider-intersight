---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hyper_flex_storage_container"
description: |-
        The HyperFlexStorageContainer object is a critical component of the HyperFlex storage architecture, representing a Storage Container (or Datastore) entity. It plays a pivotal role in managing and organizing the storage resources within a HyperFlex cluster.
        #### Purpose
        The primary purpose of the HyperFlexStorageContainer is to provide a structured and efficient way to manage storage allocations and utilization within a HyperFlex environment. It acts as a container for data storage, supporting various storage types and ensuring optimal data management.
        #### Key Concepts
        - **Storage Management:** Facilitates the organization and management of storage resources, allowing for efficient allocation and utilization of space within a cluster.
        - **Data Integrity:** Ensures the reliability and consistency of data storage, supporting various storage types such as SMB, NFS, and iSCSI.
        - **Access Control:** Offers a robust access control system with privilege sets, ensuring that only authorized personnel can manage storage arrays.
        - **Integration:** Seamlessly integrates with the HyperFlex cluster's infrastructure, providing essential storage services and capabilities.

---

# Data Source: intersight_storage_hyper_flex_storage_container
The HyperFlexStorageContainer object is a critical component of the HyperFlex storage architecture, representing a Storage Container (or Datastore) entity. It plays a pivotal role in managing and organizing the storage resources within a HyperFlex cluster.
#### Purpose
The primary purpose of the HyperFlexStorageContainer is to provide a structured and efficient way to manage storage allocations and utilization within a HyperFlex environment. It acts as a container for data storage, supporting various storage types and ensuring optimal data management.
#### Key Concepts
- **Storage Management:** Facilitates the organization and management of storage resources, allowing for efficient allocation and utilization of space within a cluster.
- **Data Integrity:** Ensures the reliability and consistency of data storage, supporting various storage types such as SMB, NFS, and iSCSI.
- **Access Control:** Offers a robust access control system with privilege sets, ensuring that only authorized personnel can manage storage arrays.
- **Integration:** Seamlessly integrates with the HyperFlex cluster's infrastructure, providing essential storage services and capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hyper_flex_storage_container.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Storage container's creation time. 
* `data_block_size`:(int) Storage Container data block size 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `in_use`:(bool) Indicates whether the Storage Container has Volumes. 
* `last_access_time`:(string) Storage container's last access time. 
* `last_modified_time`:(string) Storage container's last modified time. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `provisioned_capacity`:(int) Provisioned Capacity of the Storage container. 
* `provisioned_volume_capacity_utilization`:(float) Provisioned Capacity Utilization of All Volumes associated with the Storage Container. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Storage Container type (SMB/NFS/iSCSI).* `NFS` - Storage container created/accesed through NFS protocol.* `SMB` - Storage container created/accessed through SMB protocol.* `iSCSI` - Storage container created/accessed through iSCSI protocol. 
* `un_compressed_used_bytes`:(int) Uncompressed bytes on Storage Container. 
* `uuid`:(string) UUID of the Datastore/Storage Containter. 
* `volume_count`:(int) Number of Volumes associated with the Storage Container. 
 
