---
subcategory: "scheduler"
layout: "intersight"
page_title: "Intersight: intersight_scheduler_task_schedule"
description: |-
        TaskSchedules are the runtime scheduling objects that drive one-time or recurring executions. A TaskSchedule can run directly using explicit schedule parameters or can be “policy-based” by referencing a SchedulePolicy. It also defines what to execute (a REST-stim task request and/or an associated workflow definition) and tracks current scheduling status.
        #### Purpose
        Create and manage scheduled executions including lifecycle actions like suspend or resume, while exposing current and historical execution status.
        #### Key Concepts
        - **Schedule orchestration:** Combines *when to run* (schedule params or referenced policy) with *what to run* (task request or workflow definition).
        - **Policy-backed vs direct schedules:** `usePolicy` determines whether execution timing comes from a SchedulePolicy or inline schedule params.
        - **Lifecycle actions:** Supports operational actions like suspend or resume (and captures the last applied action).
        - **Status and execution telemetry:** Exposes next or previous run timestamps, counters (completed, failed, or skipped), and suspension reasons.
        - **Permission inheritance from target:** Inherits permissions from an `associatedObject`, ensuring schedule management aligns with the governed resource.

---

# Data Source: intersight_scheduler_task_schedule
TaskSchedules are the runtime scheduling objects that drive one-time or recurring executions. A TaskSchedule can run directly using explicit schedule parameters or can be “policy-based” by referencing a SchedulePolicy. It also defines what to execute (a REST-stim task request and/or an associated workflow definition) and tracks current scheduling status.
#### Purpose
Create and manage scheduled executions including lifecycle actions like suspend or resume, while exposing current and historical execution status.
#### Key Concepts
- **Schedule orchestration:** Combines *when to run* (schedule params or referenced policy) with *what to run* (task request or workflow definition).
- **Policy-backed vs direct schedules:** `usePolicy` determines whether execution timing comes from a SchedulePolicy or inline schedule params.
- **Lifecycle actions:** Supports operational actions like suspend or resume (and captures the last applied action).
- **Status and execution telemetry:** Exposes next or previous run timestamps, counters (completed, failed, or skipped), and suspension reasons.
- **Permission inheritance from target:** Inherits permissions from an `associatedObject`, ensuring schedule management aligns with the governed resource.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_scheduler_task_schedule.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action of the scheduled task such as suspend or resume.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. Not supported in this release. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) A description to describe the schedule for easier identification. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_action`:(string) The last action for the scheduled task is saved in this field. Set to none if there was no action.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. Not supported in this release. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A schedule name for easier identification (not required to be unique). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `submitted_by`:(string) Email address of the user who created the schedule. 
* `suspend_end_time`:(string) Suspend a task until an end date. this applies only to the action suspendTill. 
* `type`:(string) An Enum describing the type of scheduler to use.* `None` - No value was set for the schedule type (Enum value None).* `OneTime` - Define a one-time task execution time that will not automatically repeat.* `Recurring` - Specify a recurring task cadence based on a predefined pattern, such as daily, weekly, monthly, or every <interval> pattern. 
* `use_policy`:(bool) Indicates if the schedule is policy based or not. 
 
