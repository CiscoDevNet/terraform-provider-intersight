---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_server_hw_catalog_info"
description: |-
        The ServerHwCatalogInfos object provides granular, server-specific hardware catalog information used for detailed BOM validation.
        #### Purpose
        This stores the technical specifications and constraints for individual server models, allowing the system to perform deep validation of server hardware components.
        #### Key Concepts
        - **Detailed Specifications:** Contains JSON-formatted values that define the technical specifications of server components, such as processors, memory, and disks.
        - **Display Customization:** Provides user-friendly display models that map technical Cisco PIDs to readable strings for Intersight users.
        - **Constraint Enforcement:** Implements constraints to handle hardware-specific rules, such as maximum data disk counts based on the management platform.

---

# Data Source: intersight_hcl_server_hw_catalog_info
The ServerHwCatalogInfos object provides granular, server-specific hardware catalog information used for detailed BOM validation.
#### Purpose
This stores the technical specifications and constraints for individual server models, allowing the system to perform deep validation of server hardware components.
#### Key Concepts
- **Detailed Specifications:** Contains JSON-formatted values that define the technical specifications of server components, such as processors, memory, and disks.
- **Display Customization:** Provides user-friendly display models that map technical Cisco PIDs to readable strings for Intersight users.
- **Constraint Enforcement:** Implements constraints to handle hardware-specific rules, such as maximum data disk counts based on the management platform.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_server_hw_catalog_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `display_model`:(string) Display model of the server hardware. In many cases, the model string used in the catalog will be the hardware Cisco PID and the Intersight model string is a more user-friendly string with vendor information in it. This will be the user-friendly modal string to be used in Intersight. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the server hardware from the catalog file. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of specific tag, required to choose the correct datatype while reading the value. 
 
