---
subcategory: "sdcard"
layout: "intersight"
page_title: "Intersight: intersight_sdcard_policy_inventory"
description: |-
        The Policy object is a framework for configuring SD Card settings on endpoints, with a focus on managing partition types and virtual drives.
        #### Purpose
        The Policy object simplifies the configuration and deployment of SD Card settings. It enables partition management by defining partition types and associated virtual drives, and facilitates drive configuration within these partitions, enhancing overall endpoint resource management.
        #### Key Concepts
        - **Virtual Drives:** Supports the creation and management of virtual drives within OS and utility partitions.
        - **Inventory Integration:** Generates inventory objects to track and manage SD Card settings across endpoints.
        - **Deprecated Relationships:** While maintaining backward compatibility, certain relationship settings are marked deprecated.

---

# Data Source: intersight_sdcard_policy_inventory
The Policy object is a framework for configuring SD Card settings on endpoints, with a focus on managing partition types and virtual drives.
#### Purpose
The Policy object simplifies the configuration and deployment of SD Card settings. It enables partition management by defining partition types and associated virtual drives, and facilitates drive configuration within these partitions, enhancing overall endpoint resource management.
#### Key Concepts
- **Virtual Drives:** Supports the creation and management of virtual drives within OS and utility partitions.
- **Inventory Integration:** Generates inventory objects to track and manage SD Card settings across endpoints.
- **Deprecated Relationships:** While maintaining backward compatibility, certain relationship settings are marked deprecated.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdcard_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
