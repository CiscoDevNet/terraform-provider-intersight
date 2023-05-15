---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_external_storage_port"
description: |-
        Port entity in Hitachi storage array. A port for an external storage system that is connected to the local storage system.

---

# Data Source: intersight_storage_hitachi_external_storage_port
Port entity in Hitachi storage array. A port for an external storage system that is connected to the local storage system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_external_storage_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_is_used`:(bool) Is Used of the external storage. 
* `external_path_mode`:(string) Path Mode of the external storage. 
* `external_port_id`:(string) Object ID of ports on an external storage system. 
* `external_serial_number`:(string) Serial Number of the external storage. 
* `external_storage_info`:(string) Storage Information of the external storage. 
* `external_wwn`:(string) WWN of the external storage port. 
* `iscsi_ip_address`:(string) The iSCSI IP Address of the external storage port. 
* `iscsi_name`:(string) The iSCSI Name of the external storage port. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(string) Port ID of the local storage. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `virtual_port_id`:(int) Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled. 
 
