---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_port"
description: |-
  Port entity in Hitachi storage array.
---

# Data Source: intersight_storage_hitachi_port
Port entity in Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `fabric_mode`:(bool) Fabric mode of the port. true, Set. false, Not set. 
* `ipv4_address`:(string) IPv4 address of Hitachi Port. 
* `ipv6_global_address`:(string) IPv6 global address value. 
* `ipv6_link_local_address`:(string) IPv6 link local address value. 
* `iqn`:(string) ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'. 
* `is_ipv6_enable`:(bool) IPv6 mode. 
* `loop_id`:(string) The value set for the port loop ID (AL_PA). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the physical port available in storage array. 
* `port_connection`:(string) Topology setting for the port. 
* `port_lun_security`:(bool) LUN security setting for the port. 
* `shortport_id`:(string) Port ID (short) of the port. 
* `speed`:(int) Operational speed of physical port measured in Gbps. 
* `status`:(string) Status of storage array port.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `tcp_mtu`:(int) Value of MTU for iSCSI communication. 
* `type`:(string) Port type - possible values are FC, FCoE and iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `wwn`:(string) World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon.Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'. 
* `wwnn`:(string) World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'. 
* `wwpn`:(string) World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'. 
