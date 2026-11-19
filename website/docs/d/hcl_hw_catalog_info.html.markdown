---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_hw_catalog_info"
description: |-
        The HwCatalogInfos object provides the essential hardware catalog metadata for the HyperFlex Data Platform, which is required for performing Bill of Materials (BOM) validation.
        #### Purpose
        This serves as the foundation for hardware compatibility checks, ensuring that the server models and types used in a HyperFlex cluster are supported by the platform.
        #### Key Concepts
        - **BOM Validation:** Acts as the reference data for validating hardware configurations during cluster deployment.
        - **Hardware Identification:** Tracks specific server models and types, such as All-Flash (AF) configurations.
        - **Catalog Integration:** Maintains relationships with detailed server hardware catalog information to provide a comprehensive validation check.

---

# Data Source: intersight_hcl_hw_catalog_info
The HwCatalogInfos object provides the essential hardware catalog metadata for the HyperFlex Data Platform, which is required for performing Bill of Materials (BOM) validation.
#### Purpose
This serves as the foundation for hardware compatibility checks, ensuring that the server models and types used in a HyperFlex cluster are supported by the platform.
#### Key Concepts
- **BOM Validation:** Acts as the reference data for validating hardware configurations during cluster deployment.
- **Hardware Identification:** Tracks specific server models and types, such as All-Flash (AF) configurations.
- **Catalog Integration:** Maintains relationships with detailed server hardware catalog information to provide a comprehensive validation check.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_hw_catalog_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_model`:(string) Server model information for HyperFlex servers. 
* `server_type`:(string) Server type of the server hardware. For example, server type AF is for an all-flash server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
