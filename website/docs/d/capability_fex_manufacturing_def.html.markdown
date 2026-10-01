---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_fex_manufacturing_def"
description: |-
        The FexManufacturingDef object provides manufacturing and descriptive metadata for a FEX platform.
        #### Purpose
        This standardizes product identity and descriptive fields for FEX platforms, supporting consistent UI rendering and enabling catalog-backed identification across inventory and workflows.
        #### Key Concepts
        - **Catalog-driven product identity:** Provides canonical PID/SKU and descriptive fields.
        - **Consistent platform labeling:** Enables uniform naming across systems and reports.
        - **Model differentiation:** Supports categorizing FEX platforms beyond raw inventory strings.
        - **Operational usability:** Improves readability and traceability in management workflows.

---

# Data Source: intersight_capability_fex_manufacturing_def
The FexManufacturingDef object provides manufacturing and descriptive metadata for a FEX platform.
#### Purpose
This standardizes product identity and descriptive fields for FEX platforms, supporting consistent UI rendering and enabling catalog-backed identification across inventory and workflows.
#### Key Concepts
- **Catalog-driven product identity:** Provides canonical PID/SKU and descriptive fields.
- **Consistent platform labeling:** Enables uniform naming across systems and reports.
- **Model differentiation:** Supports categorizing FEX platforms beyond raw inventory strings.
- **Operational usability:** Improves readability and traceability in management workflows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_fex_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for Fabric extender. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for Fabric extender. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fex_code_name`:(string) Code Name for Fabric extender. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a Fabric extender. 
* `product_name`:(string) Product Name for Fabric extender. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Fabric extender. 
* `vid`:(string) VID information for Fabric extender. 
 
