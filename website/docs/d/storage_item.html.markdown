---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_item"
description: |-
        Items (storage) represent local storage information, including name, size, utilization, operational state, and alarm type, and can relate to file inventory objects.
        #### Purpose
        Provide high-level visibility into local storage items for monitoring capacity/health and enumerating associated file items.
        #### Key Concepts
        - **Local storage summary:** Tracks size and used percentage.
        - **Operational signals:** Includes oper state and alarm type.
        - **File relationship:** Can relate to file item objects representing stored content.

---

# Data Source: intersight_storage_item
Items (storage) represent local storage information, including name, size, utilization, operational state, and alarm type, and can relate to file inventory objects.
#### Purpose
Provide high-level visibility into local storage items for monitoring capacity/health and enumerating associated file items.
#### Key Concepts
- **Local storage summary:** Tracks size and used percentage.
- **Operational signals:** Includes oper state and alarm type.
- **File relationship:** Can relate to file item objects representing stored content.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alarm_type`:(string) The alarmType of the Local storage. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Local storage. 
* `oper_state`:(string) The operState of the Local storage. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(string) The size (MiB) of the Local storage. 
* `used`:(string) The used percent of the Local storage. 
* `used_val`:(float) The used value (MiB) of the Local storage. 
 
