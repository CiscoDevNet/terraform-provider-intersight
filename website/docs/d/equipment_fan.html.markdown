---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan"
description: |-
        The equipment.Fan object represents an individual cooling fan within a larger piece of equipment, such as a fan module, chassis, or Fabric Extender (FEX).
        #### Purpose
        The Fan object is designed to provide granular inventory and monitoring for a single fan. It captures the fan's identity (fanId), its operational state (operState), and any specific health issues (operReason). This level of detail is crucial for precise thermal management and fault detection within a system.
        #### Key Concepts
        - **Granular Component Monitoring:** Allows for the tracking of the status of each individual fan in a system.
        - **Health and State Reporting:** Provides the current operState (e.g., ok, fail) and detailed health reasons, such as speed or temperature threshold conditions.
        - **Unique Identification:** Identified by a fanId and moduleId, which specify its position within its parent fan module.
        - **Hierarchical Context:** Exists as a child of an equipment.FanModule or equipment.FEX, linking it to its physical housing and the larger system.

---

# Data Source: intersight_equipment_fan
The equipment.Fan object represents an individual cooling fan within a larger piece of equipment, such as a fan module, chassis, or Fabric Extender (FEX).
#### Purpose
The Fan object is designed to provide granular inventory and monitoring for a single fan. It captures the fan's identity (fanId), its operational state (operState), and any specific health issues (operReason). This level of detail is crucial for precise thermal management and fault detection within a system.
#### Key Concepts
- **Granular Component Monitoring:** Allows for the tracking of the status of each individual fan in a system.
- **Health and State Reporting:** Provides the current operState (e.g., "ok", "fail") and detailed health reasons, such as speed or temperature threshold conditions.
- **Unique Identification:** Identified by a fanId and moduleId, which specify its position within its parent fan module.
- **Hierarchical Context:** Exists as a child of an equipment.FanModule or equipment.FEX, linking it to its physical housing and the larger system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_fan.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description for the fan. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fan_id`:(int) This field acts as the identifier for this particular Fan, within the Fabric Interconnect. 
* `fan_module_id`:(int) This field is used to identify the Fan Module to which this Fan belongs. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `module_id`:(int) Fan module Identifier for the fan. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the pluggable Fan. 
* `oper_state`:(string) This field is used to indicate this fan unit's operational state. 
* `part_number`:(string) This field identifies the Part Number for this Fan Unit. 
* `pid`:(string) This field identifies the Product ID for the fans. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Fan Unit. 
* `tray_id`:(int) Tray identifier for the fan module. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vid`:(string) This field identifies the Vendor ID for this Fan Unit. 
 
