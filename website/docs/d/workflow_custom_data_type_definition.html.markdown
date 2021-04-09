---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_custom_data_type_definition"
description: |-
  Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
---

# Data Source: intersight_workflow_custom_data_type_definition
Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_custom_data_type_definition.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `composite_type`:(bool) When true this data type definition is a collection of type definitions to represent composite data like JSON. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) A human-friendly description of this custom data type indicating it's domain and usage. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 