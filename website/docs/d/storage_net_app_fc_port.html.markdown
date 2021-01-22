---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_fc_port"
description: |-
  Fibre Channel (FC) port is a port on a node in a storage array.
---

# Data Source: intersight_storage_net_app_fc_port
Fibre Channel (FC) port is a port on a node in a storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `iqn`:(string) ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the physical port available in storage array. 
* `port_status`:(string) Status of storage array port. 
* `speed`:(int) Operational speed of physical port measured in Gbps. 
* `state`:(string) State of the port available in storage array.* `Unknown` - Default unknown port state.* `StartUp` - The port in the storage array is booting up.* `LinkNotConnected` - The port has finished initialization, but a link with the fabric is not established.* `Online` - The port is initialized and a link with the fabric has been established.* `LinkDisconnected` - The link on this port is currently not established.* `OfflineUser` - The port is administratively disabled.* `OfflineSystem` - The port is set to offline by the system. This happens when the port encounters too many errors.* `NodeOffline` - The state information for the port cannot be retrieved. The node is offline or inaccessible. 
* `status`:(string) Status of storage array port.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `type`:(string) Port type - possible values are FC, FCoE and iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `uuid`:(string) UUID of physical port. 
* `wwn`:(string) World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon.Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'. 
* `wwnn`:(string) World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'. 
* `wwpn`:(string) World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'. 
