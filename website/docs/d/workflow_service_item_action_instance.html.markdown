---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_action_instance"
description: |-
        Service item action instance which represents one action on a service item instance.

---

# Data Source: intersight_workflow_service_item_action_instance
Service item action instance which represents one action on a service item instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_service_item_action_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Name of the action that needs to be performed on the service item instance.* `None` - No action is set, this is the default value for action field.* `Validate` - Validate the action instance inputs and run the validation workflows.* `Start` - Start a new execution of the action instance.* `Rerun` - Rerun the service item action instance from the beginning.* `Retry` - Retry the workflow that has failed from the failure point.* `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing.* `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time when the action was stopped or completed execution last time. 
* `last_action`:(string) The last action that was issued on the action definition workflows is saved in this property.* `None` - No action is set, this is the default value for action field.* `Validate` - Validate the action instance inputs and run the validation workflows.* `Start` - Start a new execution of the action instance.* `Rerun` - Rerun the service item action instance from the beginning.* `Retry` - Retry the workflow that has failed from the failure point.* `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing.* `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the action instance is created in the system by appending name of the service item instance to the name of the action definition. 
* `resourcelifecycle_status`:(string) Lifecycle state of service item instance.* `Creating` - The service item is not yet created and creation action is in progress.* `Created` - The service item is created.* `Decommissioning` - The service item is not yet decommissioned and decommission action is in progress.* `Decommissioned` - The service item is decommisioned.* `Deleting` - The service item is not yet deleted and deletion action is in progress.* `Deleted` - The service item is deleted.* `Failed` - The service item action is failed to perform the operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) The time when the action was started for execution last time. 
* `status`:(string) State of the service item action instance.* `NotStarted` - An action on the service item is not yet started and it is in a draft mode. A service item action instance can be deleted in this state.* `Validating` - A validate action has been triggered on the action and until it completes the start action cannot be issued.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The action on the service item instance failed and can be retried.* `Completed` - The action on the service item instance completed successfully.* `Stopping` - The stop action is running on the action instance.* `Stopped` - The action on the service item instance has stopped. 
* `user_id_or_email`:(string) The user identifier who invoked the request to create the service item instance. 
 
