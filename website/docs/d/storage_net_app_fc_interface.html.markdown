---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_fc_interface"
description: |-
        The NetAppFcInterfaces object represents a logical interface used for Fibre Channel (SAN) traffic.
        ####  Purpose
        It manages the SAN connectivity for hosts, ensuring that LUNs are accessible over the Fibre Channel fabric.
        ####  Key Concepts
        - **SAN Connectivity:** Maps logical FC interfaces to physical FC ports.
        - **Operational State:** Monitors the status of the FC interface.
        - **SVM Association:** Links the interface to the appropriate Storage VM.

---

# Data Source: intersight_storage_net_app_fc_interface
The NetAppFcInterfaces object represents a logical interface used for Fibre Channel (SAN) traffic.
####  Purpose
It manages the SAN connectivity for hosts, ensuring that LUNs are accessible over the Fibre Channel fabric.
####  Key Concepts
- **SAN Connectivity:** Maps logical FC interfaces to physical FC ports.
- **Operational State:** Monitors the status of the FC interface.
- **SVM Association:** Links the interface to the appropriate Storage VM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_fc_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(string) FC interface is enabled or not. 
* `interface_state`:(string) The state of the FC interface.* `Down` - The state is set to down if the interface is not enabled.* `Up` - The state is set to up if the interface is enabled. 
* `iqn`:(string) ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the physical port available in storage array. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(int) Operational speed of physical port measured in Gbps. 
* `state`:(string) The state of the FC interface.* `down` - An inactive port is listed as Down.* `up` - An active port is listed as Up.* `present` - An active port is listed as present. 
* `status`:(string) Status of storage array port.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `svm_name`:(string) The storage virtual machine name for the interface. 
* `type`:(string) Port type - possible values are FC, FCoE and iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `uuid`:(string) Uuid of NetApp FC Interface. 
* `volume_name`:(string) The parent volume name for the interface. 
* `wwn`:(string) World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon.Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'. 
* `wwnn`:(string) World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'. 
* `wwpn`:(string) World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'. 
 
