---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_workload_instance"
description: |-
        A workload instance that can be deployed, modified, or managed.

---

# Data Source: intersight_workload_workload_instance
A workload instance that can be deployed, modified, or managed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_workload_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action to be performed on the workload instance.* `None` - Absence of any action on the workload instance.* `Suspend` - Pauses the execution of the workload instance, temporarily stopping its operations without permanently removing it.* `Resume` - Restarts a suspended workload instance, allowing it to continue operations from where it left off.* `Deploy` - Initiates the deployment of the workload instance, provisioning the necessary resources and starting its execution.* `Retry` - Attempts to re-deploy the workload instance, either due to a previous failure or to apply changes made to the instance.* `RetryAll` - Attempts to re-deploy all workload instances associated with the same deployment, either due to a previous failure or to apply changes made to the instances.* `Attach` - Associates the workload instance with its assigned resources, allowing it to utilize the resources for its operations.* `Detach` - Disassociates the workload instance from its assigned resources, preventing it from using the resources for its operations.* `UnAssign` - Detaches assigned resources from the workload instance while keeping the instance active. 
* `conformance`:(string) The conformance status of the deployment.* `Ok` - The deployment conforms to the preferred version of the workload.* `NonConformant` - The deployment does not conform to the preferred version of the workload. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_action`:(string) The last action performed on the workload instance.* `None` - Absence of any action on the workload instance.* `Suspend` - Pauses the execution of the workload instance, temporarily stopping its operations without permanently removing it.* `Resume` - Restarts a suspended workload instance, allowing it to continue operations from where it left off.* `Deploy` - Initiates the deployment of the workload instance, provisioning the necessary resources and starting its execution.* `Retry` - Attempts to re-deploy the workload instance, either due to a previous failure or to apply changes made to the instance.* `RetryAll` - Attempts to re-deploy all workload instances associated with the same deployment, either due to a previous failure or to apply changes made to the instances.* `Attach` - Associates the workload instance with its assigned resources, allowing it to utilize the resources for its operations.* `Detach` - Disassociates the workload instance from its assigned resources, preventing it from using the resources for its operations.* `UnAssign` - Detaches assigned resources from the workload instance while keeping the instance active. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this Workload instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The current status of the workload instance.* `Staging` - The instance is in the staging phase, awaiting further actions.* `ReadyToDeploy` - The instance is fully configured and ready for deployment.* `InProgress` - Deployment or modification of the instance is currently in progress.* `Ok` - The instance is running successfully without issues.* `Failed` - The instance has encountered an error or failure preventing normal operation.* `Suspended` - The instance has been temporarily paused and is inactive.* `ChangesScheduled` - There is a change in the configuration that needs to be pushed to the instance.* `InSufficientResource` - The instance lacks the necessary resources to operate.* `OutOfService` - The instance is no longer available or operational.* `UnAssigning` - The instance is being unassigned or removed from service. 
* `status_change_reason`:(string) The context or justification for the status transition.* `None` - No changes have been made.* `ResourceDisqualified` - The change in resource status triggered due to the resource being disqualified. 
 
