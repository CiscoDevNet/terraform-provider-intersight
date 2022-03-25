---
subcategory: "search"
layout: "intersight"
page_title: "Intersight: intersight_search_tag_item"
description: |-
        The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
        See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.

---

# Data Source: intersight_search_tag_item
The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_search_tag_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `nr_count`:(int) The number of times this tag key has been set across all resources. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(string) Key of the Tag from all the resources in Intersight. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
