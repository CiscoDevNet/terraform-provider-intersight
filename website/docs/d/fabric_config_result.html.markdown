---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_config_result"
description: |-
        The ConfigResult object is an integral part of the configuration management framework, designed to provide comprehensive state and detailed information for deploy and validation profile configuration results. It facilitates the tracking and reporting of configuration outcomes.
        #### Purpose
        The ConfigResult serves as a consolidated entity for capturing the results of configuration deployments and validations, ensuring transparent reporting and management of configuration states. It aids in identifying errors, warnings, and informational messages during configuration processes.
        #### Key Concepts
        - **Result Tracking:** Captures detailed configuration results, including errors and warnings, to facilitate informed decision-making and troubleshooting.
        - **Validation and Deployment:** Supports validation and deployment processes, ensuring accurate and reliable configuration management.
        - **Access and Security:** Leverages privilege sets to manage access to configuration results, maintaining security and integrity of configuration data.
        - **Relationship Management:** Integrates with related objects to provide context for configuration results, enhancing understanding and management of configuration states.

---

# Data Source: intersight_fabric_config_result
The ConfigResult object is an integral part of the configuration management framework, designed to provide comprehensive state and detailed information for deploy and validation profile configuration results. It facilitates the tracking and reporting of configuration outcomes. 
#### Purpose
The ConfigResult serves as a consolidated entity for capturing the results of configuration deployments and validations, ensuring transparent reporting and management of configuration states. It aids in identifying errors, warnings, and informational messages during configuration processes. 
#### Key Concepts
- **Result Tracking:** Captures detailed configuration results, including errors and warnings, to facilitate informed decision-making and troubleshooting.
- **Validation and Deployment:** Supports validation and deployment processes, ensuring accurate and reliable configuration management.
- **Access and Security:** Leverages privilege sets to manage access to configuration results, maintaining security and integrity of configuration data.
- **Relationship Management:** Integrates with related objects to provide context for configuration results, enhancing understanding and management of configuration states.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_config_result.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
 
