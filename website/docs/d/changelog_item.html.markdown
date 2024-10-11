---
subcategory: "changelog"
layout: "intersight"
page_title: "Intersight: intersight_changelog_item"
description: |-
        An API contract changelog item. It represents an item of contract changes between the version indicated by the attribute semanticVersion and the previous version.

---

# Data Source: intersight_changelog_item
An API contract changelog item. It represents an item of contract changes between the version indicated by the attribute semanticVersion and the previous version.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_changelog_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `date_version`:(string) The date version for the API contract changelog item in the format rfc3339 with no fraction seconds set.  Note that there can be more than one item per DateVersion. Example: 2023-12-19T00:00:00Z . 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `entity`:(string) The operationId of the endpoint for which changelog item is being generated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `semantic_version`:(string) The semantic version for the API contract changelog item. Note that there can be more than one item per SemanticVersion. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `value`:(string) The value of the API contract changelog item. 
 
