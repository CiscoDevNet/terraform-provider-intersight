---
subcategory: "rack"
layout: "intersight"
page_title: "Intersight: intersight_rack_unit_personality"
description: |-
        UnitPersonalities represent “rack unit personality” records that model a server as if it had a particular defined personality, without requiring the server’s product identifier (PID) to be reprogrammed. This enables internal workflows and inventory consumers to treat a rack server as a specific personality profile for capability, compatibility, or behavior modeling.
        #### Purpose
        Provide a read-only (and selectively updatable) inventory object that captures an assigned personality identity and descriptive metadata for a rack unit, allowing the platform to model and reason about the server’s effective personality independent of its physical PID programming.
        #### Key Concepts
        - **Personality-based modeling**: Enables representing a server according to a defined personality (`personalityId`/`name`) without changing hardware PID.
        - **Rack unit scope**: Labeled “RackUnit Personality” and permission-inherited from `computeRackUnit`, tying access and applicability to the underlying rack unit.
        - **Inventory object pattern**: Extends `inventory.Base`, indicating it is an observed/recorded entity used for inventory and correlation.
        - **Identity and metadata**:
        - `personalityId` uniquely identifies the personality applied/recognized.
        - `name` provides a human-readable personality name.
        - `additionalInfo` carries supporting descriptive context.
        - **Source integration**: Includes handlers for UCSM (`computePersonality`) and CIMC (`rackUnitPersonality`) queries, showing it can be populated from multiple management endpoints.
        - **Governed updates**: UPDATE is limited to administrators who manage servers, reflecting controlled modification of the modeled personality context.

---

# Data Source: intersight_rack_unit_personality
UnitPersonalities represent “rack unit personality” records that model a server as if it had a particular defined personality, without requiring the server’s product identifier (PID) to be reprogrammed. This enables internal workflows and inventory consumers to treat a rack server as a specific personality profile for capability, compatibility, or behavior modeling.
#### Purpose
Provide a read-only (and selectively updatable) inventory object that captures an assigned personality identity and descriptive metadata for a rack unit, allowing the platform to model and reason about the server’s effective personality independent of its physical PID programming.
#### Key Concepts
- **Personality-based modeling**: Enables representing a server according to a defined personality (`personalityId`/`name`) without changing hardware PID.
- **Rack unit scope**: Labeled “RackUnit Personality” and permission-inherited from `computeRackUnit`, tying access and applicability to the underlying rack unit.
- **Inventory object pattern**: Extends `inventory.Base`, indicating it is an observed/recorded entity used for inventory and correlation.
- **Identity and metadata**:
  - `personalityId` uniquely identifies the personality applied/recognized.
  - `name` provides a human-readable personality name.
  - `additionalInfo` carries supporting descriptive context.
- **Source integration**: Includes handlers for UCSM (`computePersonality`) and CIMC (`rackUnitPersonality`) queries, showing it can be populated from multiple management endpoints.
- **Governed updates**: UPDATE is limited to administrators who manage servers, reflecting controlled modification of the modeled personality context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_rack_unit_personality.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `additional_info`:(string) Additional info about the added software personality. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the software personality. 
* `personality_id`:(int) Unique identity of added software personality. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
