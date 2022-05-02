---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_health_check_definition"
description: |-
        Service item health check definition metadata.

---

# Data Source: intersight_workflow_service_item_health_check_definition
Service item health check definition metadata.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_service_item_health_check_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) Category that the health check belongs to. 
* `common_cause_and_resolution`:(string) Static information detailing the common cause for the health check failure.It also gives possible remediation actions that can be taken to remedy the health check failure. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the health check definition. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execution_mode`:(string) Execution mode of the health check on service item.* `OnDemand` - Execute the health check on-demand.* `Periodic` - Execute the health check on a periodic basis. 
* `label`:(string) Label for the health check definition that is displayed on UI. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check definition. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
