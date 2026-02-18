---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_schema_descriptor"
description: |-
        The capability.ServerSchemaDescriptor is a hardware descriptor that identifies a server's Redfish capability for specific actions, such as controlling the locator LED, based on its CIMC firmware and Redfish schema version.
        #### Purpose
        The primary purpose of this object is to create a mapping between a server's firmware version and the specific Redfish properties used to control certain hardware features. For example, it defines the correct locatorLedName property to use for toggling the locator LED, which can vary between different Redfish schema versions. This allows for consistent management actions across a diverse range of server firmware.
        #### Key Concepts
        - **Redfish Capability Mapping:** Links server hardware actions to specific Redfish schema properties (redfishSchema, locatorLedName).
        - **Firmware-Dependent Logic:** The descriptor is identified by vendor, model, and version, making the mapping dependent on the server's firmware.
        - **Action Abstraction:** Allows the management system to use a consistent action (e.g., turn on locator LED) while the descriptor provides the specific implementation detail for the target server.
        - **Extensible Catalog:** As a HardwareDescriptor, it is part of a catalog that can be updated to support new server models and Redfish schema versions without changing core management logic.

---

# Data Source: intersight_capability_server_schema_descriptor
The capability.ServerSchemaDescriptor is a hardware descriptor that identifies a server's Redfish capability for specific actions, such as controlling the locator LED, based on its CIMC firmware and Redfish schema version.
#### Purpose
The primary purpose of this object is to create a mapping between a server's firmware version and the specific Redfish properties used to control certain hardware features. For example, it defines the correct locatorLedName property to use for toggling the locator LED, which can vary between different Redfish schema versions. This allows for consistent management actions across a diverse range of server firmware.
#### Key Concepts
- **Redfish Capability Mapping:** Links server hardware actions to specific Redfish schema properties (redfishSchema, locatorLedName).
- **Firmware-Dependent Logic:** The descriptor is identified by vendor, model, and version, making the mapping dependent on the server's firmware.
- **Action Abstraction:** Allows the management system to use a consistent action (e.g., "turn on locator LED") while the descriptor provides the specific implementation detail for the target server.
- **Extensible Catalog:** As a HardwareDescriptor, it is part of a catalog that can be updated to support new server models and Redfish schema versions without changing core management logic.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_schema_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `locator_led_name`:(string) Redfish property name for the server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `redfish_schema`:(string) Redfish version information for the server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
