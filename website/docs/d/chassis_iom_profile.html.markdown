---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_iom_profile"
description: |-
        ### Overview
        The IomProfile object is a specialized component within the chassis configuration framework, designed to manage settings specific to Input/Output Module (IOM) configurations. It facilitates streamlined management of IOMs within chassis, ensuring consistent and effective configuration across networked systems.
        #### Purpose
        IomProfile serves as the configuration container for IOM settings within a chassis, supporting efficient deployment and validation of IOM configurations.
        #### Key Concepts
        - **Targeted Configuration:** - Focuses on IOM-specific settings, ensuring precise control over these critical components.
        - **Deployment Insight:** - Provides validation and deployment results, aiding in efficient management of IOM configurations.
        - **Integrated Relationships:** - Maintains connections with related objects for comprehensive configuration management.
        - **Access Control:** - Governed by privilege sets to ensure secure handling of IOM configuration processes.

---

# Data Source: intersight_chassis_iom_profile
### Overview
The IomProfile object is a specialized component within the chassis configuration framework, designed to manage settings specific to Input/Output Module (IOM) configurations. It facilitates streamlined management of IOMs within chassis, ensuring consistent and effective configuration across networked systems.   
#### Purpose  
IomProfile serves as the configuration container for IOM settings within a chassis, supporting efficient deployment and validation of IOM configurations.   
#### Key Concepts  
- **Targeted Configuration:** - Focuses on IOM-specific settings, ensuring precise control over these critical components. 
- **Deployment Insight:** - Provides validation and deployment results, aiding in efficient management of IOM configurations. 
- **Integrated Relationships:** - Maintains connections with related objects for comprehensive configuration management. 
- **Access Control:** - Governed by privilege sets to ensure secure handling of IOM configuration processes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_iom_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `iom_entity`:(string) IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight.* `IOMA` - IOM on left side of chassis.* `IOMB` - IOM on right side of chassis. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
