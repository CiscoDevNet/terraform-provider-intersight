---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_instance"
description: |-
        Service item instance is one instance of a service item based on a service item definition.

---

# Data Source: intersight_workflow_service_item_instance
Service item instance is one instance of a service item based on a service item definition.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_service_item_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for this service item instance. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the resource. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `last_status`:(string) Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted.* `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state.* `Okay` - The last action on the service item instance completed and the service item instance is in Okay state.* `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the service item instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `resourcelifecycle_status`:(string) Lifecycle state of service item instance.* `Creating` - The service item is not yet created and creation action is in progress.* `Created` - The service item is created.* `Decommissioning` - The service item is not yet decommissioned and decommission action is in progress.* `Decommissioned` - The service item is decommisioned.* `Deleting` - The service item is not yet deleted and deletion action is in progress.* `Deleted` - The service item is deleted.* `Failed` - The service item action is failed to perform the operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the service item instance which controls the actions that can be performed on this instance.* `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state.* `Okay` - The last action on the service item instance completed and the service item instance is in Okay state.* `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. 
* `user_id_or_email`:(string) The user identifier which indicates the user that started this workflow. 
 
