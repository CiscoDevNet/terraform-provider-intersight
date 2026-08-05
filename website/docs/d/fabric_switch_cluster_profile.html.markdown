---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile"
description: |-
        The SwitchClusterProfile object is an essential element in the network management domain, designed to specify configuration policies for a cluster of switches. It consolidates the configuration context of all the switch profiles referred by the Switch Cluster Profile, ensuring seamless integration and management of network resources.
        #### Purpose
        The SwitchClusterProfile serves as a comprehensive configuration entity for managing multiple switches as a unified cluster. It allows administrators to define, deploy, and manage network switch profiles efficiently, providing a streamlined approach to switch configuration and deployment.
        #### Key Concepts
        - **Configuration Consolidation:** Aggregates configuration policies across multiple switch profiles, providing a single point of management for cluster configurations.
        - **Deployment Management:** Facilitates the deployment of switch profiles across different switches in the cluster, ensuring coordinated configuration changes and updates.
        - **Access Control:** Incorporates privilege sets to manage access and operations, ensuring secure configuration management.
        - **Template Utilization:** Supports the creation and management of switch profiles from templates, promoting consistency and reducing setup time.

---

# Data Source: intersight_fabric_switch_cluster_profile
The SwitchClusterProfile object is an essential element in the network management domain, designed to specify configuration policies for a cluster of switches. It consolidates the configuration context of all the switch profiles referred by the Switch Cluster Profile, ensuring seamless integration and management of network resources.
#### Purpose
The SwitchClusterProfile serves as a comprehensive configuration entity for managing multiple switches as a unified cluster. It allows administrators to define, deploy, and manage network switch profiles efficiently, providing a streamlined approach to switch configuration and deployment.
#### Key Concepts
- **Configuration Consolidation:** Aggregates configuration policies across multiple switch profiles, providing a single point of management for cluster configurations.
- **Deployment Management:** Facilitates the deployment of switch profiles across different switches in the cluster, ensuring coordinated configuration changes and updates.
- **Access Control:** Incorporates privilege sets to manage access and operations, ensuring secure configuration management.
- **Template Utilization:** Supports the creation and management of switch profiles from templates, promoting consistency and reducing setup time.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_cluster_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The support actions are -- Deploy, Unassign. 
* `chassis_assignment_mode`:(string) Source of the chassis assigned to the Switch Cluster Profile. Values can be Static or None. Static is used if a chassis is attached directly to a Switch Cluster Profile. None is used if no chassis is attached to a Switch Cluster Profile. Serial pre-assignment is also considered None.* `Static` - Chassis is directly assigned to switch cluster profile.* `None` - No chassis is assigned to the switch cluster profile. 
* `chassis_pre_assign_by_serial`:(string) Serial number of the chassis that would be assigned to this pre-assigned switch cluster profile. It can be any string that adheres to the following constraints:It should start and end with an alphanumeric character.It cannot be more than 20 characters. 
* `create_time`:(string) The time when this managed object was created. 
* `deploy_status`:(string) Deploy status of the switch cluster profile indicating if deployment has been initiated on all the members of the cluster profile.* `None` - Switch profiles not deployed on either of the switches.* `Complete` - Both switch profiles of the cluster profile are deployed.* `Partial` - Only one of the switch profiles of the cluster profile is deployed. 
* `deployed_switches`:(string) Values indicating the switches on which the cluster profile has been deployed. 0 indicates that the profile has not been deployed on any switch, 1 indicates that the profile has been deployed on A, 2 indicates that it is deployed on B and 3 indicates that it is deployed on both.* `None` - Switch profiles not deployed on either of the fabric interconnects.* `A` - Switch profiles deployed only on fabric interconnect A.* `B` - Switch profiles deployed only on fabric interconnect B.* `AB` - Switch profiles deployed on both fabric interconnect A and B. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_profiles_count`:(int) Number of switch profiles that are part of this cluster profile. 
* `target_platform`:(string) Type of the profile. 'UcsDomain' profile for network and management configuration on UCS Fabric Interconnect. 'UnifiedEdge' profile for network, management and chassis configuration on Unified Edge.* `UCS Domain` - Profile/policy type for network and management configuration on UCS Fabric Interconnect.* `Unified Edge` - Profile/policy type for network, management and chassis configuration on Unified Edge. 
* `template_sync_status`:(string) The sync status of the current MO wrt the attached Template MO.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `user_label`:(string) The user defined label assigned to the switch profile. 
 
