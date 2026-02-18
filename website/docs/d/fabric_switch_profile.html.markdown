---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_profile"
description: |-
        The SwitchProfile object is a critical component in network switch management, designed to specify configuration policies for individual managed network switches. This provides detailed configuration options to ensure optimal switch operation and network performance.
        #### Purpose
        The SwitchProfile serves as a detailed configuration entity for individual network switches, allowing administrators to define and manage switch settings comprehensively. It supports the deployment of configuration changes and validates compliance with established policies.
        #### Key Concepts
        - **Policy Specification:** Outlines specific configuration policies for managed switches, ensuring adherence to network requirements and standards.
        - **Configuration Management:** Supports detailed configuration change tracking and validation, promoting efficient management of switch configurations.
        - **Access Control:** Utilizes privilege sets to secure configuration operations, maintaining integrity and security of switch management.
        - **Snapshot and Compliance:** Provides mechanisms for taking configuration snapshots and checking compliance, ensuring consistent network performance.

---

# Data Source: intersight_fabric_switch_profile
The SwitchProfile object is a critical component in network switch management, designed to specify configuration policies for individual managed network switches. This provides detailed configuration options to ensure optimal switch operation and network performance. 
#### Purpose
The SwitchProfile serves as a detailed configuration entity for individual network switches, allowing administrators to define and manage switch settings comprehensively. It supports the deployment of configuration changes and validates compliance with established policies. 
#### Key Concepts
- **Policy Specification:** Outlines specific configuration policies for managed switches, ensuring adherence to network requirements and standards.
- **Configuration Management:** Supports detailed configuration change tracking and validation, promoting efficient management of switch configurations.
- **Access Control:** Utilizes privilege sets to secure configuration operations, maintaining integrity and security of switch management.
- **Snapshot and Compliance:** Provides mechanisms for taking configuration snapshots and checking compliance, ensuring consistent network performance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) Value indicating the switch side on which the switch profile or template has to be deployed.* `None` - Switch side not defined for the policy configurations in the switch profile or template.* `A` - Policy configurations in the switch profile or template to be deployed on fabric interconnect A.* `B` - Policy configurations in the switch profile or template to be deployed on fabric interconnect B. 
* `target_platform`:(string) Type of the profile. 'UcsDomain' profile for network and management configuration on UCS Fabric Interconnect. 'UnifiedEdge' profile for network, management and chassis configuration on Unified Edge.* `UCS Domain` - Profile/policy type for network and management configuration on UCS Fabric Interconnect.* `Unified Edge` - Profile/policy type for network, management and chassis configuration on Unified Edge. 
* `template_sync_status`:(string) The sync status of the current MO wrt the attached Template MO.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
