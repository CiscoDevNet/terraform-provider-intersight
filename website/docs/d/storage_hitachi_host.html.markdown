---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_host"
description: |-
  A host group entity in Hitachi storage array. It is an abstraction used by Hitachi storage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
---

# Data Source: intersight_storage_hitachi_host
A host group entity in Hitachi storage array. It is an abstraction used by Hitachi storage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `authentication_mode`:(string) Authentication mode for the iSCSI target.* `N/A` - Not available.* `CHAP` - CHAP-authentication mode.* `NONE` - No-authentication mode.* `BOTH` - Comply with Host Setting. 
* `description`:(string) Short description about the host. 
* `host_group_id`:(string) ID of the host group. 
* `host_group_number`:(int) Host group number for this host. 
* `is_chap_mutual`:(bool) Mutual CHAP setting that is Enabled or Disabled. 
* `iscsi_name`:(string) IQN (iSCSI qualified name). Can be up to 255 characters long. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
* `port_id`:(string) Port ID of the host group. 
* `port_lun_security`:(bool) LUN security setting for the port. 
* `type`:(string) Host Group type, it can be FC or iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `wwn`:(string) World wide port name, 64 bit address represented in hexa decimal notation. 
