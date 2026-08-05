---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hyper_flex_volume"
description: |-
        The HyperFlexVolume object represents a HyperFlex Volume entity, which is integral to the storage capabilities of the HyperFlex system. It encapsulates the essential attributes and functionalities of a volume within the HyperFlex storage architecture.
        #### Purpose
        HyperFlexVolume serves as the primary entity for managing volume storage within a HyperFlex cluster. It provides essential capabilities for volume management, ensuring efficient storage operations and data handling.
        #### Key Concepts
        - **Volume Management:** Supports the creation, configuration, and maintenance of volumes, essential for organizing storage resources in a HyperFlex cluster.
        - **Data Handling:** Facilitates data storage operations, ensuring high performance and reliability within the storage environment.
        - **Access Control:** Implements comprehensive access control measures with privilege sets, allowing secure management of volume resources.
        - **Integration:** Works in conjunction with other storage components and systems, enabling a cohesive and efficient storage infrastructure.

---

# Data Source: intersight_storage_hyper_flex_volume
The HyperFlexVolume object represents a HyperFlex Volume entity, which is integral to the storage capabilities of the HyperFlex system. It encapsulates the essential attributes and functionalities of a volume within the HyperFlex storage architecture.
#### Purpose
HyperFlexVolume serves as the primary entity for managing volume storage within a HyperFlex cluster. It provides essential capabilities for volume management, ensuring efficient storage operations and data handling.
#### Key Concepts
- **Volume Management:** Supports the creation, configuration, and maintenance of volumes, essential for organizing storage resources in a HyperFlex cluster.
- **Data Handling:** Facilitates data storage operations, ensuring high performance and reliability within the storage environment.
- **Access Control:** Implements comprehensive access control measures with privilege sets, allowing secure management of volume resources.
- **Integration:** Works in conjunction with other storage components and systems, enabling a cohesive and efficient storage infrastructure.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hyper_flex_volume.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(int) Provisioned Capacity of the Storage container in Bytes. 
* `client_id`:(string) Client ID to which the volume belongs. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the volume. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_modified_time`:(string) Last modified time as UTC of the volume. 
* `lun_uuid`:(string) UUID of Lun associated with the volume. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `serial_number`:(string) Serial number of the volume. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `uuid`:(string) UUID of the Datastore/Storage Containter. 
* `volume_access_mode`:(string) Access Mode of the volume.* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.* `` - Unknown disk access mode. 
* `volume_mode`:(string) Mode of the volume.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk.* `` - Disk mode is either unknown or not supported. 
* `volume_type`:(string) The Type of the volume. 
 
