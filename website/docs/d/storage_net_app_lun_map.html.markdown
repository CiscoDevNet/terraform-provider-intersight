---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_lun_map"
description: |-
        NetApp LUN mapping is the process of associating a LUN with an initiator group. When a LUN is mapped to an initiator group, initiators in the initiator group are granted access to the LUN.

---

# Data Source: intersight_storage_net_app_lun_map
NetApp LUN mapping is the process of associating a LUN with an initiator group. When a LUN is mapped to an initiator group, initiators in the initiator group are granted access to the LUN.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_lun_map.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `igroup_uuid`:(string) UUID of the initiator group. 
* `lun_uuid`:(string) Universally unique identifier of the LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 
