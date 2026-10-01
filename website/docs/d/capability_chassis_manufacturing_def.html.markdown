---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_chassis_manufacturing_def"
description: |-
        The ChassisManufacturingDef object represents manufacturing and product-identification metadata for a chassis platform in the capability catalog.
        #### Purpose
        This standardizes chassis product identity fields (such as PID/SKU and descriptive metadata) so management and UI layers can consistently render chassis identity and correlate discovered inventory to catalog-backed platform definitions.
        #### Key Concepts
        - **Catalog-backed identity:** Links physical chassis inventory to a known definition in the capability catalog.
        - **Consistent labeling:** Provides canonical naming/description fields used across views and workflows.
        - **Platform classification:** Enables downstream logic that depends on chassis platform characteristics.
        - **Lifecycle portability:** Keeps “what this chassis is” independent from “where it is deployed.

---

# Data Source: intersight_capability_chassis_manufacturing_def
The ChassisManufacturingDef object represents manufacturing and product-identification metadata for a chassis platform in the capability catalog.
#### Purpose
This standardizes chassis product identity fields (such as PID/SKU and descriptive metadata) so management and UI layers can consistently render chassis identity and correlate discovered inventory to catalog-backed platform definitions.
#### Key Concepts
- **Catalog-backed identity:** Links physical chassis inventory to a known definition in the capability catalog.
- **Consistent labeling:** Provides canonical naming/description fields used across views and workflows.
- **Platform classification:** Enables downstream logic that depends on chassis platform characteristics.
- **Lifecycle portability:** Keeps “what this chassis is” independent from “where it is deployed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_chassis_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for Chassis enclosure. 
* `chassis_code_name`:(string) Chassis Code Name for Chassis enclosure. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for Chassis enclosure. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a Chassis enclosure. 
* `product_name`:(string) Product Name for Chassis enclosure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Chassis enclosure. 
* `vid`:(string) VID information for Chassis enclosure. 
 
