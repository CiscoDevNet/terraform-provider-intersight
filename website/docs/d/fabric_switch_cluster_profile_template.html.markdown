---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile_template"
description: |-
        The SwitchClusterProfileTemplate object is a pivotal element in network configuration management, designed to provide reusable switch profile configurations that can be applied across multiple profiles. It enables administrators to efficiently derive switch cluster profiles from predefined templates.
        #### Purpose
        The SwitchClusterProfileTemplate acts as a repository of common switch configurations that can be reused across different switch profiles. It streamlines the creation process, allowing for quick deployment and management of switch configurations.
        #### Key Concepts
        - **Reusability:** Serves as a template for deriving multiple switch cluster profiles, promoting uniformity and reducing redundancy.
        - **Synchronization:** Facilitates updates and synchronization of derived profiles through API calls, ensuring consistency across configurations.
        - **Operational Flexibility:** Supports various operations including read, create, update, and delete, granting flexibility in configuration management.
        - **Integration:** Enables seamless integration with existing network setups, aiding in efficient configuration deployment.

---

# Data Source: intersight_fabric_switch_cluster_profile_template
The SwitchClusterProfileTemplate object is a pivotal element in network configuration management, designed to provide reusable switch profile configurations that can be applied across multiple profiles. It enables administrators to efficiently derive switch cluster profiles from predefined templates. 
#### Purpose
The SwitchClusterProfileTemplate acts as a repository of common switch configurations that can be reused across different switch profiles. It streamlines the creation process, allowing for quick deployment and management of switch configurations.
#### Key Concepts
- **Reusability:** Serves as a template for deriving multiple switch cluster profiles, promoting uniformity and reducing redundancy.
- **Synchronization:** Facilitates updates and synchronization of derived profiles through API calls, ensuring consistency across configurations.
- **Operational Flexibility:** Supports various operations including read, create, update, and delete, granting flexibility in configuration management.
- **Integration:** Enables seamless integration with existing network setups, aiding in efficient configuration deployment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_cluster_profile_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_profiles_count`:(int) Number of switch profiles that are part of this cluster profile. 
* `target_platform`:(string) Type of the profile. 'UcsDomain' profile for network and management configuration on UCS Fabric Interconnect. 'UnifiedEdge' profile for network, management and chassis configuration on Unified Edge.* `UCS Domain` - Profile/policy type for network and management configuration on UCS Fabric Interconnect.* `Unified Edge` - Profile/policy type for network, management and chassis configuration on Unified Edge. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `usage`:(int) The count of switch cluster profiles derived from the template. 
 
