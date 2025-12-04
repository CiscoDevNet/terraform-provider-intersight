---
subcategory: "search"
layout: "intersight"
page_title: "Intersight: intersight_search_search_item"
description: |-
        ### Overview
        The SearchItem object represents an entry point for searching Intersight REST resources. It enables users to query and retrieve data using the OData query syntax, providing a standardized and flexible method for resource discovery within the system.
        #### Purpose
        A SearchItem serves as the primary interface for performing structured searches across various Intersight resources. Its purpose is to facilitate efficient data retrieval by allowing users to construct complex queries based on OData syntax, thereby enabling precise filtering, ordering, and selection of desired information.
        #### Key Concepts
        - **Resource Discovery:** - Designed to help users locate and access specific REST resources within the Intersight environment.
        - **API Entry Point:** - Functions as a dedicated endpoint for initiating search operations, streamlining the process of data lookup.

---

# Data Source: intersight_search_search_item
### Overview
The SearchItem object represents an entry point for searching Intersight REST resources. It enables users to query and retrieve data using the OData query syntax, providing a standardized and flexible method for resource discovery within the system.
#### Purpose
A SearchItem serves as the primary interface for performing structured searches across various Intersight resources. Its purpose is to facilitate efficient data retrieval by allowing users to construct complex queries based on OData syntax, thereby enabling precise filtering, ordering, and selection of desired information.
#### Key Concepts
- **Resource Discovery:** - Designed to help users locate and access specific REST resources within the Intersight environment.
- **API Entry Point:** - Functions as a dedicated endpoint for initiating search operations, streamlining the process of data lookup.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_search_search_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
