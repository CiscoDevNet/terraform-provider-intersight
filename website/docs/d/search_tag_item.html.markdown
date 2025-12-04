---
subcategory: "search"
layout: "intersight"
page_title: "Intersight: intersight_search_tag_item"
description: |-
        ### Overview
        The TagItem object serves as an entry point for searching tags associated with Intersight REST resources. It allows users to query and discover tags across the system using OData Query syntax, providing a mechanism for organizing and categorizing resources.
        #### Purpose
        The purpose of a TagItem is to enable efficient discovery and management of tags applied to various Intersight resources. This provides a centralized way to search for specific tags and understand their prevalence and associated values across the system, facilitating resource organization and retrieval based on custom metadata.
        #### Key Concepts
        - **Tag-Based Search:** - Specializes in querying resources based on their assigned tags, offering an alternative or supplementary search method to direct resource queries.
        - **Metadata Organization:** - Supports the systematic organization and classification of Intersight resources through tagging.

---

# Data Source: intersight_search_tag_item
### Overview
The TagItem object serves as an entry point for searching tags associated with Intersight REST resources. It allows users to query and discover tags across the system using OData Query syntax, providing a mechanism for organizing and categorizing resources.
#### Purpose
The purpose of a TagItem is to enable efficient discovery and management of tags applied to various Intersight resources. This provides a centralized way to search for specific tags and understand their prevalence and associated values across the system, facilitating resource organization and retrieval based on custom metadata.
#### Key Concepts
- **Tag-Based Search:** - Specializes in querying resources based on their assigned tags, offering an alternative or supplementary search method to direct resource queries.
- **Metadata Organization:** - Supports the systematic organization and classification of Intersight resources through tagging.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_search_tag_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `nr_count`:(int) The number of times this tag key has been set across all resources. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(string) Key of the tag from all the resources in Intersight. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of the tag definition. Refer to comm.TagDefinitionType API. 
 
