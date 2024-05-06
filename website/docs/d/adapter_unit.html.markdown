---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_unit"
description: |-
        The physical adapter present on a server.

---

# Data Source: intersight_adapter_unit
The physical adapter present on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_id`:(string) Unique Identifier of an adapter Unit within a Rack Interface. 
* `base_mac_address`:(string) Original Base Mac address of an adapter unit. 
* `connection_status`:(string) Connectivity Status of adapter - A or B or AB. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `integrated`:(string) Cisco Integrated adapter or other type. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of an adapter unit. 
* `operability`:(string) Operability state of an adapter unit. 
* `part_number`:(string) Part number of an adapter unit. 
* `pci_slot`:(string) PCIe slot of the adapter in the server. 
* `power`:(string) Power state of an adapter unit. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `thermal`:(string) Thermal state of an adapter unit. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vic_communicable`:(string) Records the current state of communication between the Virtual Interface Card (VIC) and the Cisco Integrated Management Controller (CIMC) on the server.* `Not Applicable` - Set the state of VIC communication to Not Applicable for other Platforms.* `Yes` - VIC is reachable from CIMC.* `No` - VIC is not reachable from CIMC. 
* `vid`:(string) Virtual Id of the adapter in the server. 
 
