---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_config_result"
description: |-
        # Overview
        ConfigResult is designed to provide feedback on profile configuration processes, offering insights into the state and outcomes of validation and deployment operations.
        
        ## Purpose
        ConfigResult is designed to provide feedback on profile configuration processes, offering insights into the state and outcomes of validation and deployment operations.
        ## Key Concepts
        - **Result Management** –  Provides structured information on configuration outcomes, supporting efficient troubleshooting and adjustments.
        - **Feedback Support** –  Offers detailed result messages, facilitating understanding of configuration processes and outcomes.
        - **Access Control** –  Managed through privilege sets, ensuring secure access and management of configuration results.
        - **Integration** –  Supports integration with configuration profiles, ensuring cohesive feedback and management of configuration operations.

---

# Data Source: intersight_hyperflex_config_result
# Overview
 ConfigResult is designed to provide feedback on profile configuration processes, offering insights into the state and outcomes of validation and deployment operations.
 
## Purpose
ConfigResult is designed to provide feedback on profile configuration processes, offering insights into the state and outcomes of validation and deployment operations.
## Key Concepts
- **Result Management** –  Provides structured information on configuration outcomes, supporting efficient troubleshooting and adjustments.
- **Feedback Support** –  Offers detailed result messages, facilitating understanding of configuration processes and outcomes.
- **Access Control** –  Managed through privilege sets, ensuring secure access and management of configuration results.
- **Integration** –  Supports integration with configuration profiles, ensuring cohesive feedback and management of configuration operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_config_result.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_progress`:(string) The progress percentage of the running configuration or workflow. 
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `duration`:(string) The duration of the running configuration or workflow. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) The start time of the configuration or workflow. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
 
