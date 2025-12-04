---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_config_result"
description: |-
        ### Overview
        The ConfigResult object plays a vital role in server profile management by providing detailed insights into the configuration process. This captures the results of configuration actions, including deployments and validations, offering transparency and accountability in configuration management.
        #### Purpose
        ConfigResult is essential for tracking the outcomes of server profile configuration actions, ensuring that deployments and validations are accurately reported and issues are promptly identified.
        #### Key Concepts
        - **Result Tracking:** - Captures detailed information about configuration results, including errors, warnings, and successful actions.
        - **Validation and Deployment:** - Provides insights into the validation and deployment stages, supporting troubleshooting and optimization of server configurations.
        - ** Relationship Management:** - Linked with other objects to offer comprehensive views of configuration impacts.
        - **Read-Only Access:** - Primarily designed for consumption, ensuring safe integration with systems requiring configuration insights.

---

# Data Source: intersight_server_config_result
### Overview
The ConfigResult object plays a vital role in server profile management by providing detailed insights into the configuration process. This captures the results of configuration actions, including deployments and validations, offering transparency and accountability in configuration management.   
#### Purpose  
ConfigResult is essential for tracking the outcomes of server profile configuration actions, ensuring that deployments and validations are accurately reported and issues are promptly identified.   
#### Key Concepts  
- **Result Tracking:** - Captures detailed information about configuration results, including errors, warnings, and successful actions. 
- **Validation and Deployment:** - Provides insights into the validation and deployment stages, supporting troubleshooting and optimization of server configurations.
- ** Relationship Management:** - Linked with other objects to offer comprehensive views of configuration impacts. 
- **Read-Only Access:** - Primarily designed for consumption, ensuring safe integration with systems requiring configuration insights.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_config_result.<custom_name>.results[i].<propertyname>`.
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
 
