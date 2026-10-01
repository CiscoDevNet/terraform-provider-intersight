---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_chassis_descriptor"
description: |-
        ChassisDescriptors are capability-catalog hardware descriptors that uniquely identify a chassis enclosure by its vendor, model, version, revision, and catalog section context. They serve as the catalog-level identity record used to recognize and reason about chassis hardware in a standardized way.
        #### Purpose
        Provide a canonical catalog entry for chassis enclosure hardware so capabilities, compatibility logic, and platform behavior can be tied to a uniquely identified chassis model/revision.
        #### Key Concepts
        - **Catalog-based hardware identity**: Extends `HardwareDescriptor`, indicating it participates in a broader capability-catalog hardware description framework.
        - **Unique enclosure fingerprint**: Identity is composed of `vendor`, `model`, `version`, `revision`, and `section`, ensuring chassis descriptors are uniquely distinguished in the catalog.
        - **Revision-aware modeling**: `revision` captures enclosure revision differences that may affect supportability or behavior.
        - **Section-scoped governance**: Inherits permissions from `section`, tying access and organization to the containing catalog section.
        - **Catalog-admin lifecycle**: Creation and modification are restricted to `CapabilityCatalog Administrator`, while READ is available to broader operational roles.

---

# Data Source: intersight_capability_chassis_descriptor
ChassisDescriptors are capability-catalog hardware descriptors that uniquely identify a chassis enclosure by its vendor, model, version, revision, and catalog section context. They serve as the catalog-level identity record used to recognize and reason about chassis hardware in a standardized way.
#### Purpose
Provide a canonical catalog entry for chassis enclosure hardware so capabilities, compatibility logic, and platform behavior can be tied to a uniquely identified chassis model/revision.
#### Key Concepts
- **Catalog-based hardware identity**: Extends `HardwareDescriptor`, indicating it participates in a broader capability-catalog hardware description framework.
- **Unique enclosure fingerprint**: Identity is composed of `vendor`, `model`, `version`, `revision`, and `section`, ensuring chassis descriptors are uniquely distinguished in the catalog.
- **Revision-aware modeling**: `revision` captures enclosure revision differences that may affect supportability or behavior.
- **Section-scoped governance**: Inherits permissions from `section`, tying access and organization to the containing catalog section.
- **Catalog-admin lifecycle**: Creation and modification are restricted to `CapabilityCatalog Administrator`, while READ is available to broader operational roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_chassis_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision for the chassis enclosure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
