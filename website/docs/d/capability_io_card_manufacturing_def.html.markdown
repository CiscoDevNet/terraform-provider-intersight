---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_manufacturing_def"
description: |-
        The IoCardCapabilityDef object describes capabilities for chassis I/O modules (IOM/IO cards) in the capability catalog.
        #### Purpose
        IoCardCapabilityDef provides platform-specific capability flags (for example, connector support characteristics) that drive validation, deployment logic, and conditional behavior for chassis networking components.
        #### Key Concepts
        - **IOM feature declaration:** Encodes what an IOM model supports so configurations are constrained accordingly.
        - **Platform-aware validation:** Enables policy validation and workflow selection based on IOM capabilities.
        - **Catalog-backed behavior:** Prevents clients and services from relying on hard-coded per-model rules.
        - **Consistent chassis operations:** Ensures multi-chassis or mixed-platform environments behave predictably.

---

# Data Source: intersight_capability_io_card_manufacturing_def
The IoCardCapabilityDef object describes capabilities for chassis I/O modules (IOM/IO cards) in the capability catalog.
#### Purpose
IoCardCapabilityDef provides platform-specific capability flags (for example, connector support characteristics) that drive validation, deployment logic, and conditional behavior for chassis networking components.
#### Key Concepts
- **IOM feature declaration:** Encodes what an IOM model supports so configurations are constrained accordingly.
- **Platform-aware validation:** Enables policy validation and workflow selection based on IOM capabilities.
- **Catalog-backed behavior:** Prevents clients and services from relying on hard-coded per-model rules.
- **Consistent chassis operations:** Ensures multi-chassis or mixed-platform environments behave predictably.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_io_card_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for a chassis Iocard module. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for a chassis Iocard module. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a chassis Iocard module. 
* `product_name`:(string) Product Name for IO Card Module. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for a chassis Iocard module. 
* `vid`:(string) VID information for a chassis Iocard module. 
 
