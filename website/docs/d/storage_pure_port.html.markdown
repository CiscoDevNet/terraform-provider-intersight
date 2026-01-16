---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_port"
description: |-
        The PurePort object represents a port entity within a PureStorage FlashArray, facilitating network connectivity and data flow.
        #### Purpose
        PurePort facilitates the management of storage ports, ensuring efficient network connectivity and data operations within the FlashArray.
        #### Key Concepts
        - **Port Management:** Offers structured management of ports, optimizing network connectivity and data flow operations.
        - **Integration:** Supports integration with controller and array objects, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage port operations within the storage network.

---

# Data Source: intersight_storage_pure_port
The PurePort object represents a port entity within a PureStorage FlashArray, facilitating network connectivity and data flow.  
#### Purpose  
PurePort facilitates the management of storage ports, ensuring efficient network connectivity and data operations within the FlashArray. 
#### Key Concepts  
- **Port Management:** Offers structured management of ports, optimizing network connectivity and data flow operations. 
- **Integration:** Supports integration with controller and array objects, enhancing storage management capabilities. 
 - **Secure Access:** Utilizes privilege sets to secure and manage port operations within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `failover`:(string) Name of the port to which this port has failed over. 
* `iqn`:(string) ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the physical port available in storage array. 
* `nqn`:(string) The NVMe Qualified Name (NQN) associated with the host for ethernet port. 
* `portal`:(string) Ip address of iSCSI portal configured on the port. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(int) Operational speed of physical port measured in Gbps. 
* `status`:(string) Status of storage array port.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `type`:(string) Port type - possible values are FC, FCoE and iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `wwn`:(string) World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon.Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'. 
* `wwnn`:(string) World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'. 
* `wwpn`:(string) World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'. 
 
