---
subcategory: "scheduler"
layout: "intersight"
page_title: "Intersight: intersight_scheduler_schedule_policy"
description: |-
        SchedulePolicies define reusable, policy-based scheduling definitions for running tasks either one-time or on a recurring cadence. They encapsulate one or more schedule definitions (each with its own start time, timezone, duration, and cadence parameters) and can optionally enforce “block dates” during which schedules must not run.
        #### Purpose
        Provide a centrally managed scheduling policy that can be attached to schedules or deployments so recurring operational actions (for example, workflow executions) run predictably and consistently across an organization.
        #### Key Concepts
        - **Policy-based scheduling:** A single policy can contain multiple schedule definitions, making scheduling reusable and standardized.
        - **One-time and recurring support:** Policies can describe schedules that execute once or on cadences like daily/weekly/monthly via typed schedule params.
        - **Block dates (blackout windows):** Optional block date intervals prevent execution during sensitive periods (maintenance freezes, change windows).
        - **Usage governance:** Tracks usage count so administrators can understand whether a policy is safe to delete (cannot delete when in use).
        - **Attachable association:** Maintains references to associated objects (e.g., task schedules and workload deployments) using the policy.

---

# Data Source: intersight_scheduler_schedule_policy
SchedulePolicies define reusable, policy-based scheduling definitions for running tasks either one-time or on a recurring cadence. They encapsulate one or more schedule definitions (each with its own start time, timezone, duration, and cadence parameters) and can optionally enforce “block dates” during which schedules must not run.
#### Purpose
Provide a centrally managed scheduling policy that can be attached to schedules or deployments so recurring operational actions (for example, workflow executions) run predictably and consistently across an organization.
#### Key Concepts
- **Policy-based scheduling:** A single policy can contain multiple schedule definitions, making scheduling reusable and standardized.
- **One-time and recurring support:** Policies can describe schedules that execute once or on cadences like daily/weekly/monthly via typed schedule params.
- **Block dates (blackout windows):** Optional block date intervals prevent execution during sensitive periods (maintenance freezes, change windows).
- **Usage governance:** Tracks usage count so administrators can understand whether a policy is safe to delete (cannot delete when in use).
- **Attachable association:** Maintains references to associated objects (e.g., task schedules and workload deployments) using the policy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_scheduler_schedule_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_block_dates`:(bool) Enable or disable block dates. If set to true, the schedule will not run during the block date interval. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage_count`:(int) The number of profiles, templates and deployments that are using this policy. This is used to determine if the policy can be deleted.If the usageCount is greater than 0, the policy cannot be deleted. 
 
