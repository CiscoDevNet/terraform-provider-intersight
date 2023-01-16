---
subcategory: "openapi"
layout: "intersight"
page_title: "Intersight: intersight_openapi_api_method_meta"
description: |-
        Contains metadata about APIs in the previously uploaded OpenAPI specification file.

---

# Data Source: intersight_openapi_api_method_meta
Contains metadata about APIs in the previously uploaded OpenAPI specification file.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_openapi_api_method_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of the given API. 
* `display_label`:(string) The display label of the given API. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `method`:(string) The method type for the given API.* `GET` - Method type which indicates it is a GET API call* `POST` - Method type which indicates it is a POST API call* `PUT` - Method type which indicates it is a PUT API call* `PATCH` - Method type which indicates it is a PATCH API call* `DELETE` - Method type which indicates it is a DELETE API call 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The description of the given API. 
* `path`:(string) Path of the selected API endpoint. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
