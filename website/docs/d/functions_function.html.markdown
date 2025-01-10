---
subcategory: "functions"
layout: "intersight"
page_title: "Intersight: intersight_functions_function"
description: |-
        The managed object which has info about custom function to be built and deployed.

---

# Data Source: intersight_functions_function
The managed object which has info about custom function to be built and deployed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_functions_function.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action of the function such as build, deploy, undeploy.* `None` - No action is set, this is the default value for action field.* `Publish` - Publish a Function that was saved or built. 
* `code`:(string) Custom function code to create the first function version, mandatory in function creation payload. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the Function. 
* `description`:(string) Description of the function. 
* `display_name`:(string) The display name of the function. Display name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the Function. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the function. Name can only contain letters (a-z), numbers (0-9), hyphen (-). 
* `runtime_moid`:(string) Moid of runtime which is used to create the first function version, mandatory in function creation payload. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The target version of the function, which is needed by action. 
 
