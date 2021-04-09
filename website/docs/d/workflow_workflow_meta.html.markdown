---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_meta"
description: |-
  Contains a workflow definition which is a sequence of tasks to execute.
---

# Data Source: intersight_workflow_workflow_meta
Contains a workflow definition which is a sequence of tasks to execute.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_workflow_meta.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for the workflow. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name given to the workflow. 
* `retryable`:(bool) When true, this workflow can be retried for 2 weeks since the last modification of the workflow. 
* `schema_version`:(int) The Conductor schema version that decides what attribute can be supported. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src`:(string) The src is workflow owner service. 
* `type`:(string) The type of workflow definition.* `SystemDefined` - System defined workflow definition.* `UserDefined` - User defined workflow definition.* `Dynamic` - Dynamically defined workflow definition. 
* `nr_version`:(int) The version for the workflow so we can support multiple versions for the same workflow name. 
* `wait_on_duplicate`:(bool) Parameter decides if workflows will wait for a duplicate to finish before starting a new one. 
 