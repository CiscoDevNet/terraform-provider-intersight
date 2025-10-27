---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_config_result_entry"
description: |-
        # Overview
        The ConfigResultEntry object describes the result of a cluster validation or deployment operation, providing detailed result entries and context.
        ## Purpose
        ConfigResultEntry is designed to offer detailed feedback on validation and deployment operations, ensuring that administrators have comprehensive insight into configuration processes.
        ## Key Concepts
        - **Detailed Feedback** – Provides comprehensive result entries, supporting efficient troubleshooting and understanding of validation and deployment  operations.
        - **Result Management** –  Facilitates structured management of configuration outcomes, ensuring efficient adjustments and improvements.
        - **Access Control** –  Managed through privilege sets, ensuring secure access and management of configuration result entries.
        - **Integration** –  Supports integration with configuration results, ensuring cohesive feedback and management of validation and deployment operations.

---

# Data Source: intersight_hyperflex_config_result_entry
# Overview
The ConfigResultEntry object describes the result of a cluster validation or deployment operation, providing detailed result entries and context.
## Purpose
ConfigResultEntry is designed to offer detailed feedback on validation and deployment operations, ensuring that administrators have comprehensive insight into configuration processes.
## Key Concepts
- **Detailed Feedback** – Provides comprehensive result entries, supporting efficient troubleshooting and understanding of validation and deployment  operations.
- **Result Management** –  Facilitates structured management of configuration outcomes, ensuring efficient adjustments and improvements.
- **Access Control** –  Managed through privilege sets, ensuring secure access and management of configuration result entries.
- **Integration** –  Supports integration with configuration results, ensuring cohesive feedback and management of validation and deployment operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_config_result_entry.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `completed_time`:(string) The completed time of the task in installer. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `message`:(string) Localized message based on the locale setting of the user's context. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owner_id`:(string) The identifier of the object that owns the result message.The owner ID is used to correlate a given result entry to a task or entity. For example, a config resultentry that describes the result of a workflow task may have the task's instance ID as the owner. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Values  -- Ok, Ok-with-warning, Errored. 
* `type`:(string) Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config. 
 
