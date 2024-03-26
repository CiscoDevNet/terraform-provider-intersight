---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile"
description: |-
        This specifies the configuration policies for a cluster of switches.

---

# Data Source: intersight_fabric_switch_cluster_profile
This specifies the configuration policies for a cluster of switches.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_cluster_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
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
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
