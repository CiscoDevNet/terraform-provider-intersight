---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_endpoint"
description: |-
        PCIe endpoints that are connected to a PCIe switch.

---

# Data Source: intersight_pci_endpoint
PCIe endpoints that are connected to a PCIe switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_endpoint.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_enclosure_id`:(int) The identifier of the enclosure device of the actual physical device to which the logical PCIe endpoint is pointing to. 
* `device_id`:(int) The identifier of the actual physical device to which the logical PCIe endpoint is pointing to. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_type`:(string) The type of the actual physical device to which the logical PCIe endpoint is pointing to.* `Unknown` - The device type of the physical device is unknown.* `NetworkAdapter` - The device type of the physical device is a NIC adapter.* `CPU` - The device type of the physical device is CPU.* `GPU` - The device type of the physical device is GPU. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `endpoint_id`:(string) The identifier of the PCIe endpoint within a X-Fabric module PCIe switch. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the PCIe endpoint, as reported by the XFM platform software (BMC). 
* `oper_state`:(string) Operational state of the PCIe endpoint. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
