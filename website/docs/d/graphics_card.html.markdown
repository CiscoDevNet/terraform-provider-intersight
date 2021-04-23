---
subcategory: "graphics"
layout: "intersight"
page_title: "Intersight: intersight_graphics_card"
description: |-
  Graphics Card present in a server.
---

# Data Source: intersight_graphics_card
Graphics Card present in a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_graphics_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `card_id`:(int) The id of the graphics card. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(int) The device id of the graphics card. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expander_slot`:(string) The expander slot information of the card. 
* `firmware_version`:(string) The firmware version of the graphics card. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) The current mode of the graphics card. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_gpus`:(string) The number of controllers under each card. 
* `oper_state`:(string) The current operational state of the graphics card. 
* `pci_address`:(string) The PCI address of the graphics card. 
* `pci_address_list`:(string) This list contains the PCI address of all controllers for corresponding card. 
* `pci_slot`:(string) The PCI slot name of the graphics card. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 