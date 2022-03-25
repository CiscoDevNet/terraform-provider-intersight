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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_host.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `authentication_mode`:(string) Authentication mode for the iSCSI target.* `N/A` - Authentication Mode is not available.* `CHAP` - CHAP-authentication mode.* `NONE` - Authentication mode is not set.* `BOTH` - Comply with Host Setting. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the host. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_group_id`:(string) ID of the host group for this host. 
* `host_group_number`:(int) Host group number for this host. 
* `is_chap_mutual`:(bool) Mutual CHAP setting that is Enabled or Disabled. 
* `iscsi_name`:(string) IQN (iSCSI qualified name). Can be up to 255 characters long. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
* `port_id`:(string) Port ID of the host group. 
* `port_lun_security`:(bool) LUN security setting for the port. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Host Group type, it can be FC or iSCSI.* `FC` - Port supports fibre channel protocol.* `iSCSI` - Port supports iSCSI protocol.* `FCoE` - Port supports fibre channel over ethernet protocol. 
* `wwn`:(string) World wide port name, 64 bit address represented in hexa decimal notation. 
 
