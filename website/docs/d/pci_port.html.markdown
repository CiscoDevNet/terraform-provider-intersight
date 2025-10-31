---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_port"
description: |-
        PCIe Switch port on the PCIe Switch.

---

# Data Source: intersight_pci_port
PCIe Switch port on the PCIe Switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the PCIe Switch Port. Name will be of the format 'PCIe Port <$id>'. 
* `oper_state`:(string) The operational status of the PCIe switch port. 
* `port_id`:(int) The identifier of the PCIe switch port. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The current role of the PCIe switch port.* `Unconfigured` - Indicates that the PCIe switch port role is currently unconfigured.* `Downstream` - Indicates that the PCIe switch port role is currently downstream. A downstream port connects to a PCIe peripheral device such as a GPU or a network adapter.* `Upstream` - Indicates that the PCIe switch port role is currently upstream. An upstream port connects to a PCIe root complex such as a CPU.* `Unknown` - Indicates that the PCIe switch port role is currently unknown. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uri`:(string) The unique identifier of the PCIe switch port as reported by chassis expander module management controller. 
* `width`:(int) The link width of the PCIe switch port. 
 
