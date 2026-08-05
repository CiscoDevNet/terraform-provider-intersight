---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_category_summary"
description: |-
        The ConfigCategorySummary object provides a categorized summary of what was captured in a backup instance. It is structured to help customers understand coverage and contents without inspecting raw backup archives.
        #### Purpose
        ConfigCategorySummary exists to describe the backed-up configuration inventory in a structured way, grouped by scope (account vs organization) and by category (such as profiles, policies, templates, and similar groupings).
        #### Key Concepts
        - **Human-Readable Backup Inventory:** Offers an at-a-glance view of what the backup contains, improving transparency and operational confidence.
        - **Scope-Aware Categorization:** Separates account-owned and organization-owned configuration, reflecting how configuration is organized and governed.
        - **Category-Based Organization:** Groups items by meaningful configuration domains to support quick validation and reporting.
        - **Backup Instance Association:** Tightly bound to a specific backup artifact, ensuring summaries remain consistent with the exact backup content.

---

# Data Source: intersight_mgmt_config_category_summary
The ConfigCategorySummary object provides a categorized summary of what was captured in a backup instance. It is structured to help customers understand coverage and contents without inspecting raw backup archives.
#### Purpose
ConfigCategorySummary exists to describe the backed-up configuration inventory in a structured way, grouped by scope (account vs organization) and by category (such as profiles, policies, templates, and similar groupings).
#### Key Concepts
- **Human-Readable Backup Inventory:** Offers an at-a-glance view of what the backup contains, improving transparency and operational confidence.
- **Scope-Aware Categorization:** Separates account-owned and organization-owned configuration, reflecting how configuration is organized and governed.
- **Category-Based Organization:** Groups items by meaningful configuration domains to support quick validation and reporting.
- **Backup Instance Association:** Tightly bound to a specific backup artifact, ensuring summaries remain consistent with the exact backup content.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_category_summary.<custom_name>.results[i].<propertyname>`.
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
 
