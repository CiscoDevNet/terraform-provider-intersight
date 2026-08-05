---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_restore_category_summary"
description: |-
        The ConfigRestoreCategorySummary object provides a categorized view of what is being restored (and what was restored) as part of a restore operation. It supports restore validation and reporting through structured summaries.
        #### Purpose
        ConfigRestoreCategorySummary exists to summarize restore activity by scope and category, allowing customers to understand which configuration domains were targeted and their restored outcomes without relying on low-level import records.
        #### Key Concepts
        - **Restore Coverage Transparency:** Enables quick review of restore scope and category-level composition.
        - **Scope-Aware Reporting:** Reflects both account and organization configuration domains, aligning with how configuration ownership is structured.
        - **Category-Oriented Organization:** Groups restored items into meaningful operational categories to simplify validation by users and post-restore checks.
        - **Restore Operation Association:** Bound to a specific restore run, ensuring summaries correspond exactly to the initiated restore process.

---

# Data Source: intersight_mgmt_config_restore_category_summary
The ConfigRestoreCategorySummary object provides a categorized view of what is being restored (and what was restored) as part of a restore operation. It supports restore validation and reporting through structured summaries.
#### Purpose
ConfigRestoreCategorySummary exists to summarize restore activity by scope and category, allowing customers to understand which configuration domains were targeted and their restored outcomes without relying on low-level import records.
#### Key Concepts
- **Restore Coverage Transparency:** Enables quick review of restore scope and category-level composition.
- **Scope-Aware Reporting:** Reflects both account and organization configuration domains, aligning with how configuration ownership is structured.
- **Category-Oriented Organization:** Groups restored items into meaningful operational categories to simplify validation by users and post-restore checks.
- **Restore Operation Association:** Bound to a specific restore run, ensuring summaries correspond exactly to the initiated restore process.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_restore_category_summary.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category_name`:(string) The name of the sub category. The objects are categorized based on the type of the object. For organization - it can be pools, policies, profiles, templates. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the organization or account. 
* `scope`:(string) The scope of the category, the objects belong to an organization or an account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
