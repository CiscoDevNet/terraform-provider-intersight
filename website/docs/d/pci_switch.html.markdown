---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_switch"
description: |-
        PCIe switch present in the system which connects PCIe devices (CPU, GPU and the PCIe adapter).

---

# Data Source: intersight_pci_switch
PCIe switch present in the system which connects PCIe devices (CPU, GPU and the PCIe adapter).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_switch.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) The device id of the switch. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health`:(string) The composite health of the switch. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_adaptors`:(string) The number of GPUs and PCIe adapters connected the switch. 
* `oper_state`:(string) Operational state of the PCIe switch that connects the CPU GPU and the PCIe adapter. 
* `pci_address`:(string) The PCIe address of the switch. 
* `pci_slot`:(string) The PCIe slot name of the switch. 
* `pci_switch_id`:(string) Unique Identifier of a PCIe switch within a chassis expander module. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `product_name`:(string) The model information for the switch. 
* `product_revision`:(string) The product revision of the switch. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sub_device_id`:(string) The sub device id of the switch. 
* `sub_vendor_id`:(string) The sub vendor id of the switch. 
* `temperature`:(string) The current temperature of the switch. 
* `type`:(string) The type information of the switch. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vendor_id`:(string) The vendor id of the switch. 
 
