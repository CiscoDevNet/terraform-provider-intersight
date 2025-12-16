---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_blueprint"
description: |-
        ### Overview
        The Blueprint object defines a reusable, versioned template for deploying infrastructure configurations and services. It encapsulates input schemas, managed object generation rules, service item associations, and resource constraints into a cohesive deployment unit that can be referenced by workload definitions.
        #### Purpose
        A Blueprint serves as a modular building block for constructing complex workload definitions. Blueprints provides a standardized way to define infrastructure patterns, application configurations, and service deployments that can be reused across multiple workloads. Blueprints support versioning, platform targeting, and dependency management to ensure reliable and consistent deployments.
        #### Key Concepts
        - **Versioning and Default Versions:** - Each blueprint supports multiple versions with the ability to designate a default version, enabling safe evolution of blueprint specifications while maintaining compatibility with existing workloads.
        - **Input Definition Schema:** - Defines the schema for all inputs required by the blueprint, supporting complex data types including primitives, collections, and managed object references.
        - **Generated Object Management:** - Specifies which Intersight managed objects (policies, profiles, etc.) should be automatically generated from blueprint inputs and how those objects should be configured.
        - **Service Item Integration:** - Associates service items with the blueprint, defining which service workflows are executed during deployment and lifecycle operations.
        - **Resource Constraints:** - Defines infrastructure resource requirements and constraints that must be satisfied for successful blueprint deployment.
        - **Blueprint Dependencies:** - Supports declaring dependencies on other blueprints, enabling compositional patterns and ensuring all required components are available during deployment.
        - **Platform Targeting:** - Specifies which Intersight platform types the blueprint supports, ensuring deployments only target compatible infrastructure.

---

# Data Source: intersight_workload_blueprint
### Overview
The Blueprint object defines a reusable, versioned template for deploying infrastructure configurations and services. It encapsulates input schemas, managed object generation rules, service item associations, and resource constraints into a cohesive deployment unit that can be referenced by workload definitions.
#### Purpose
A Blueprint serves as a modular building block for constructing complex workload definitions. Blueprints provides a standardized way to define infrastructure patterns, application configurations, and service deployments that can be reused across multiple workloads. Blueprints support versioning, platform targeting, and dependency management to ensure reliable and consistent deployments.
#### Key Concepts
- **Versioning and Default Versions:** - Each blueprint supports multiple versions with the ability to designate a default version, enabling safe evolution of blueprint specifications while maintaining compatibility with existing workloads.
- **Input Definition Schema:** - Defines the schema for all inputs required by the blueprint, supporting complex data types including primitives, collections, and managed object references.
- **Generated Object Management:** - Specifies which Intersight managed objects (policies, profiles, etc.) should be automatically generated from blueprint inputs and how those objects should be configured.
- **Service Item Integration:** - Associates service items with the blueprint, defining which service workflows are executed during deployment and lifecycle operations.
- **Resource Constraints:** - Defines infrastructure resource requirements and constraints that must be satisfied for successful blueprint deployment.
- **Blueprint Dependencies:** - Supports declaring dependencies on other blueprints, enabling compositional patterns and ensuring all required components are available during deployment.
- **Platform Targeting:** - Specifies which Intersight platform types the blueprint supports, ensuring deployments only target compatible infrastructure.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_blueprint.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_version`:(bool) The flag to indicate that this is the default version of the blueprint. 
* `description`:(string) The description for this blueprint which provides information on what are the pre-requisites to deploy the blueprint and what features are supported on the blueprint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_meta`:(bool) When set to false the blueprint is created for use within internal services. 
* `label`:(string) A user friendly short name to identify the blueprint. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this blueprint. You can have multiple versions of the blueprint with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), or an underscore (_). 
* `platform_type`:(string) The Intersight platforms supported by this blueprint.* `None` - Default value is none, platform is not mentioned.* `UnifiedEdge` - The blueprint is created for Unified Edge platform. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version of the blueprint to support multiple versions. 
 
