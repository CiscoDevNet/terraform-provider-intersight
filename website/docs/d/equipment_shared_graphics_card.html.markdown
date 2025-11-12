---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_shared_graphics_card"
description: |-
        Graphics card within a PCIe node that can be shared by one or more servers.

---

# Data Source: intersight_equipment_shared_graphics_card
Graphics card within a PCIe node that can be shared by one or more servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_shared_graphics_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of the GPU card. 
* `device_id`:(int) The unique device identifier assigned by the vendor to a specific model of GPU. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) The version of the GPU firmware. 
* `gpu_id`:(string) The identifier of the graphics card. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of graphics card. 
* `part_number`:(string) Part number identifier for the graphics card. 
* `pci_slot`:(string) PCIe slot of the GPU in the PCIe node. 
* `pid`:(string) The unique product ID associated with the GPU card. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sub_device_id`:(int) The subsystem device identifier assigned by the subsystem vendor to a specific model of GPU. 
* `sub_vendor_id`:(int) The unique vendor identifier assigned to the organization which integrates the GPU. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vendor_id`:(int) The unique vendor identifier assigned to the manufacturer of the GPU. 
 
