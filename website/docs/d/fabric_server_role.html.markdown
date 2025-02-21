---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_server_role"
description: |-
        Configuration object sent by user to create a server port.

---

# Data Source: intersight_fabric_server_role
Configuration object sent by user to create a server port.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_server_role.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `auto_negotiation_disabled`:(bool) Auto negotiation configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 100G speed port on UCS-FI-6536 and should be set as True. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fec`:(string) Forward error correction configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 25G speed ports on UCS-FI-6454/UCS-FI-64108 and should be set as Cl74.* `Auto` - Forward error correction option 'Auto'.* `Cl91` - Forward error correction option 'cl91'.* `Cl74` - Forward error correction option 'cl74'.* `rs-cons16` - Forward Error Correction option \ rs-cons16\ .* `rs-ieee` - Forward Error Correction option \ rs-ieee\ .* `Off` - Turn off Forward Error Correction. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `preferred_device_id`:(int) Preferred device ID to be configured by user for the connected device. This ID must be specified together with the 'PreferredDeviceType' property. This ID will only takes effect if the actual connected device matches the 'PreferredDeviceType'. If the preferred ID is not available, the ID is automatically allocated and assigned by the system. If different preferred IDs are specified for the ports connected to the same device, only the preferred ID (if specified) of the port that is discovered first will be considered. 
* `preferred_device_type`:(string) Device type for which preferred ID to be configured. If the actual connected device does not match the specified device type, the system ignores the 'PreferredDeviceId' property.* `Auto` - Preferred Id will be ignored if specified with this type.* `RackServer` - Connected device type is Rack Unit Server.* `Chassis` - Connected device type is Chassis. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
 
