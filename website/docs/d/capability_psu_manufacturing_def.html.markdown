---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_psu_manufacturing_def"
description: |-
        The capability.PsuManufacturingDef object is a capability definition that stores manufacturing-specific details for a Power Supply Unit (PSU). This provides a comprehensive set of identifiers for product and inventory management.
        #### Purpose
        This object's primary function is to act as a reference for the manufacturing details of a PSU. It aggregates key product identifiers such as pid (Product Identifier), vid (Vendor ID), sku (Stock Keeping Unit), and productName. This information is essential for accurate inventory, asset tracking, and integration with support and ordering systems.
        #### Key Concepts
        - **Manufacturing Reference:** Centralizes key manufacturing and product data for a specific PSU model.
        - **Product Identification:** Provides standard identifiers (pid, sku, productName) that allow for easy lookup and correlation with external product databases.
        - **Descriptive Data:** Includes a caption and description to offer human-readable context about the PSU.
        - **Capability Catalog Entry:** Serves as a definition within the capability catalog, which can be referenced by other objects to retrieve detailed manufacturing information.

---

# Data Source: intersight_capability_psu_manufacturing_def
The capability.PsuManufacturingDef object is a capability definition that stores manufacturing-specific details for a Power Supply Unit (PSU). This provides a comprehensive set of identifiers for product and inventory management.
#### Purpose
This object's primary function is to act as a reference for the manufacturing details of a PSU. It aggregates key product identifiers such as pid (Product Identifier), vid (Vendor ID), sku (Stock Keeping Unit), and productName. This information is essential for accurate inventory, asset tracking, and integration with support and ordering systems.
#### Key Concepts
- **Manufacturing Reference:** Centralizes key manufacturing and product data for a specific PSU model.
- **Product Identification:** Provides standard identifiers (pid, sku, productName) that allow for easy lookup and correlation with external product databases.
- **Descriptive Data:** Includes a caption and description to offer human-readable context about the PSU.
- **Capability Catalog Entry:** Serves as a definition within the capability catalog, which can be referenced by other objects to retrieve detailed manufacturing information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_psu_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for a power supply unit. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for a power supply unit. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a power supply unit. 
* `product_name`:(string) Product Name for Power Supplu Unit. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for a power supply unit. 
* `vid`:(string) VID information for a power supply unit. 
 
