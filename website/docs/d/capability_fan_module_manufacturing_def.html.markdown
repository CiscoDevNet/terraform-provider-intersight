---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_fan_module_manufacturing_def"
description: |-
        The capability.FanModuleManufacturingDef object is a capability definition that stores manufacturing-specific details for a fan module. This provides a comprehensive set of identifiers used for inventory and product identification.
        #### Purpose
        This object's primary function is to serve as a reference for the manufacturing details of a fan module unit, which may contain multiple individual fans. It aggregates key product identifiers like pid (Product Identifier), vid (Vendor ID), sku (Stock Keeping Unit), and productName. This information is crucial for accurate inventory management, asset tracking, and support.
        #### Key Concepts
        - **Manufacturing Data:** Centralizes key manufacturing and product identifiers for a fan module.
        - **Product Identification:** Provides standard identifiers (pid, sku, productName) for easy lookup and integration with other systems like ordering and support tools.
        - **Descriptive Information:** Includes a caption and description to provide human-readable context about the fan module.
        - **Capability Reference:** Acts as a definition within the capability catalog, referenced by other objects to retrieve detailed manufacturing information.

---

# Data Source: intersight_capability_fan_module_manufacturing_def
The capability.FanModuleManufacturingDef object is a capability definition that stores manufacturing-specific details for a fan module. This provides a comprehensive set of identifiers used for inventory and product identification.
#### Purpose
This object's primary function is to serve as a reference for the manufacturing details of a fan module unit, which may contain multiple individual fans. It aggregates key product identifiers like pid (Product Identifier), vid (Vendor ID), sku (Stock Keeping Unit), and productName. This information is crucial for accurate inventory management, asset tracking, and support.
#### Key Concepts
- **Manufacturing Data:** Centralizes key manufacturing and product identifiers for a fan module.
- **Product Identification:** Provides standard identifiers (pid, sku, productName) for easy lookup and integration with other systems like ordering and support tools.
- **Descriptive Information:** Includes a caption and description to provide human-readable context about the fan module.
- **Capability Reference:** Acts as a definition within the capability catalog, referenced by other objects to retrieve detailed manufacturing information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_fan_module_manufacturing_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `caption`:(string) Caption for a fan module. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for a fan module. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a fan module. 
* `product_name`:(string) Product Name for Fan Module Unit. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for a fan module. 
* `vid`:(string) VID information for a fan module. 
 
