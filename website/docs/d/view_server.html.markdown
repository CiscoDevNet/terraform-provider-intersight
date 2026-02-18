---
subcategory: "view"
layout: "intersight"
page_title: "Intersight: intersight_view_server"
description: |-
        The Server object facilitates querying a list of UCS servers, including API resource types compute.Blade and compute.RackUnit, along with associated managed object information.
        #### Purpose
        The Server object is intended to provide a REST endpoint for accessing and managing server-related data within the UCS ecosystem. It helps users query comprehensive server information, supporting various operational and management tasks.
        #### Key Concepts
        - **REST Endpoint:** Offers a standardized interface for querying server data, supporting OData operators for filtering and ordering results.
        - **Integration with MO Information:** Allows access to related managed object information such as device registration and health status.

---

# Data Source: intersight_view_server
The Server object facilitates querying a list of UCS servers, including API resource types compute.Blade and compute.RackUnit, along with associated managed object information.
#### Purpose
The Server object is intended to provide a REST endpoint for accessing and managing server-related data within the UCS ecosystem. It helps users query comprehensive server information, supporting various operational and management tasks.
#### Key Concepts
- **REST Endpoint:** Offers a standardized interface for querying server data, supporting OData operators for filtering and ordering results.
- **Integration with MO Information:** Allows access to related managed object information such as device registration and health status.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_view_server.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
