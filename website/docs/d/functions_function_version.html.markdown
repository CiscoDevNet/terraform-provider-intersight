---
subcategory: "functions"
layout: "intersight"
page_title: "Intersight: intersight_functions_function_version"
description: |-
        The managed object which has info about a specific version of custom function.

---

# Data Source: intersight_functions_function_version
The managed object which has info about a specific version of custom function.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_functions_function_version.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `code`:(string) Custom function code for Function MO. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the Function. 
* `default_version`:(bool) When true this function version will be used in functions table. The very first function created with a name will be set as the default version. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the Function. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current representation of the Function MO state.* `Saved` - Function is saved, yet to be built and deployed.* `Building` - Function is currently being built.* `Built` - The Function has been built and can now be deployed.* `Deploying` - The built Function is currently being deployed.* `Deployed` - The Function has been deployed.* `Undeploying` - The deployed function is being Undeployed.* `Deleting` - The Function is being deleted. 
* `nr_version`:(int) The version of the function to support multiple versions. 
 
