---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_storage_vm"
description: |-
        ### Overview
        The NetAppStorageVm object encapsulates a storage virtual machine, a fundamental entity within the NetAppCluster environment, responsible for serving data to clients through logical interfaces. It is integral to managing virtualized storage resources and ensuring efficient data delivery.
        #### Purpose
        A NetAppStorageVm provides a virtual framework for storage operations, enabling flexible data management and service provisioning. It acts as a virtual entity within the cluster, orchestrating data access and supporting diverse protocol configurations.
        #### Key Concepts
        - **Virtualization:** - Offers a virtualized approach to storage management, facilitating dynamic resource allocation and provisioning.
        - **Service Integration:** - Manages multiple logical interfaces (LIFs) to deliver data services to clients, ensuring seamless connectivity and interaction.
        - **Protocol Support:** - Supports various storage protocols, including NFS, CIFS, iSCSI, and NVMe, to accommodate diverse client requirements.
        - **Data Protection:** - Implements robust data protection strategies, utilizing SnapMirror and other technologies to safeguard data integrity and availability.

---

# Data Source: intersight_storage_net_app_storage_vm
### Overview
The NetAppStorageVm object encapsulates a storage virtual machine, a fundamental entity within the NetAppCluster environment, responsible for serving data to clients through logical interfaces. It is integral to managing virtualized storage resources and ensuring efficient data delivery.
#### Purpose
A NetAppStorageVm provides a virtual framework for storage operations, enabling flexible data management and service provisioning. It acts as a virtual entity within the cluster, orchestrating data access and supporting diverse protocol configurations.
#### Key Concepts
- **Virtualization:** - Offers a virtualized approach to storage management, facilitating dynamic resource allocation and provisioning.
- **Service Integration:** - Manages multiple logical interfaces (LIFs) to deliver data services to clients, ensuring seamless connectivity and interaction.
- **Protocol Support:** - Supports various storage protocols, including NFS, CIFS, iSCSI, and NVMe, to accommodate diverse client requirements.
- **Data Protection:** - Implements robust data protection strategies, utilizing SnapMirror and other technologies to safeguard data integrity and availability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_storage_vm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cifs_enabled`:(bool) Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fcp_enabled`:(bool) Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers. 
* `ipspace`:(string) IPspace name. IPspaces are distinct IP address spaces in which storage virtual machines (SVMs) reside. 
* `is_protected`:(string) Specifies whether the Storage VM is a SnapMirror source Storage VM, using SnapMirror to protect its data. 
* `iscsi_enabled`:(bool) Status for iSCSI protocol allowed to run on Vservers. 
* `key`:(string) Unique identifier of VServer across data center. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the tenant in storage array. 
* `native_fpolicy_count`:(int) The number of native FPolicy engines enabled on this SVM. 
* `nfs_enabled`:(bool) Status for Network File System Protocol ( NFS ) allowed to run on  Vservers. 
* `nvme_enabled`:(bool) Status for NVME protocol allowed to run on Vservers. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of this tenant.* `Unknown` - Component state is not available.* `Starting` - Component is being started.* `Running` - Component is currently running.* `Stopping` - Component is being stopped.* `Stopped` - Component has been stopped.* `Deleting` - Component deletion is in progress. 
* `subtype`:(string) SVM subtype (default, dp_destination, sync_source, or sync_destination). The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster. 
* `uuid`:(string) The uuid of this tenant in storage array. 
 
