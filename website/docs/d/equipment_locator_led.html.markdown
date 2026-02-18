---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_locator_led"
description: |-
        The equipment.LocatorLed object represents the physical locator LED (Light Emitting Diode) on a piece of equipment, such as a server, chassis, or Fabric Interconnect.
        #### Purpose
        The main function of this object is to inventory and control the state of a locator LED. The locator LED is a visual indicator used by data center technicians to physically identify a specific piece of hardware among many others in a rack. It allows an administrator to remotely turn the LED on or off.
        #### Key Concepts
        - **Physical Identification:** Provides a mechanism to visually identify hardware in a data center.
        - **State Control:** While the operState property is read-only in the inventory model, it is the target of actions that change the LED's state (typically on or off).
        - **Visual Feedback:** The color property reports the color of the LED, which is often blue.
        - **Broad Applicability:** Can be associated with various types of equipment, including compute.RackUnit, compute.Blade, equipment.Chassis, and network.Element.

---

# Data Source: intersight_equipment_locator_led
The equipment.LocatorLed object represents the physical locator LED (Light Emitting Diode) on a piece of equipment, such as a server, chassis, or Fabric Interconnect.
#### Purpose
The main function of this object is to inventory and control the state of a locator LED. The locator LED is a visual indicator used by data center technicians to physically identify a specific piece of hardware among many others in a rack. It allows an administrator to remotely turn the LED on or off.
#### Key Concepts
- **Physical Identification:** Provides a mechanism to visually identify hardware in a data center.
- **State Control:** While the operState property is read-only in the inventory model, it is the target of actions that change the LED's state (typically "on" or "off").
- **Visual Feedback:** The color property reports the color of the LED, which is often blue.
- **Broad Applicability:** Can be associated with various types of equipment, including compute.RackUnit, compute.Blade, equipment.Chassis, and network.Element.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_locator_led.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `color`:(string) Color of the locatorled available on an equipment. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Identifies the operational state of locatorled. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
