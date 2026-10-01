---
subcategory: "port"
layout: "intersight"
page_title: "Intersight: intersight_port_mac_binding"
description: |-
        MacBindings establish a discovered relationship between local switch ports and connected endpoints using LLDP TLVs. They record both the local port identity (slot/port/aggregate/switch) and the remote device identity (device MAC and module/chassis descriptors) to enable topology correlation.
        #### Purpose
        Provide a topology-mapping mechanism that ties a switch/FEX interface to the connected endpoint based on LLDP-advertised identifiers, supporting inventory correlation and troubleshooting of cabling/connectivity.
        #### Key Concepts
        - **LLDP-based correlation:** Uses LLDP TLVs to bind a local port to a remote endpoint identity.
        - **Local port addressing:** Captures slotId/portId/aggregatePortId and switchId to identify the local interface.
        - **Remote endpoint identity:** Captures `deviceMac` and additional module/chassis metadata (serial/model/vendor/slot/side) to identify what is connected.
        - **Chassis/module context:** Includes chassis identifiers/serial/model and module identifiers/serial/model to support mapping through intermediate components (IOM/SIOC/Adapter).
        - **Stable identity per switch context:** Identified by ` (portMac, networkElement)`, ensuring uniqueness within the owning network element domain.

---

# Data Source: intersight_port_mac_binding
MacBindings establish a discovered relationship between local switch ports and connected endpoints using LLDP TLVs. They record both the local port identity (slot/port/aggregate/switch) and the remote device identity (device MAC and module/chassis descriptors) to enable topology correlation.
#### Purpose
Provide a topology-mapping mechanism that ties a switch/FEX interface to the connected endpoint based on LLDP-advertised identifiers, supporting inventory correlation and troubleshooting of cabling/connectivity.
#### Key Concepts
- **LLDP-based correlation:** Uses LLDP TLVs to bind a local port to a remote endpoint identity.
- **Local port addressing:** Captures slotId/portId/aggregatePortId and switchId to identify the local interface.
- **Remote endpoint identity:** Captures `deviceMac` and additional module/chassis metadata (serial/model/vendor/slot/side) to identify what is connected.
- **Chassis/module context:** Includes chassis identifiers/serial/model and module identifiers/serial/model to support mapping through intermediate components (IOM/SIOC/Adapter).
- **Stable identity per switch context:** Identified by ` (portMac, networkElement)`, ensuring uniqueness within the owning network element domain.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_port_mac_binding.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Aggregate Port ID of the local Switch Interface. 
* `chassis_id`:(int) Chassis/FEX device idetifier that is local to a cluster. 
* `chassis_model`:(string) Chassis/Rack Model that is associated with the Switch/FEX interface. 
* `chassis_serial`:(string) Chassis/Rack Serial that is associated with the Switch/FEX interface. 
* `chassis_vendor`:(string) Chassis/Rack Vendor that is associated with the Switch/FEX interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mac`:(string) Device ID value that is advertised and available as a part of LLDP TLV. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `module_mode`:(int) IOM/SIOC/Adapter Mode that is associated with the Switch/FEX interface. 
* `module_model`:(string) IOM/SIOC/Adapter Model that is associated with the Switch/FEX interface. 
* `module_port_id`:(int) Uplink port identifier of the VIC that is associated with the Switch/FEX interface. 
* `module_serial`:(string) IOM/SIOC/Adapter Serial that is associated with the Switch/FEX interface. 
* `module_side`:(int) IOM/SIOC/Adapter Side that is associated with the Switch/FEX interface. 
* `module_slot`:(int) IOM/SIOC/Adapter Slot that is associated with the Switch/FEX interface. 
* `module_vendor`:(string) IOM/SIOC/Adapter Vendor that is associated with the Switch/FEX interface. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port ID of the local Switch Interface. 
* `port_mac`:(string) Port ID value that is advertised and available as a part of LLDP TLV. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot ID of the local Switch slot Interface. 
* `switch_id`:(int) Switch Identifier that is local to a cluster. 
 
