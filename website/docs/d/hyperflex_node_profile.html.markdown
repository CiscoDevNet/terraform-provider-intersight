---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_profile"
description: |-
        The NodeProfiles object defines the configuration for an individual node within a HyperFlex cluster.
        #### Purpose
        It acts as the configuration profile for a specific node, linking physical server hardware to its required network and storage IP configurations.
        #### Key Concepts
        - **Node Configuration:** Manages IP settings for hypervisors and controller VMs.
        - **Role Assignment:** Defines the role of the node (e.g., converged, compute) within the cluster.
        - **Server Mapping:** Associates the profile with a specific physical server.

---

# Data Source: intersight_hyperflex_node_profile
The NodeProfiles object defines the configuration for an individual node within a HyperFlex cluster.
#### Purpose
It acts as the configuration profile for a specific node, linking physical server hardware to its required network and storage IP configurations.
#### Key Concepts
- **Node Configuration:** Manages IP settings for hypervisors and controller VMs.
- **Role Assignment:** Defines the role of the node (e.g., converged, compute) within the cluster.
- **Server Mapping:** Associates the profile with a specific physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_node_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hxdp_data_ip`:(string) IP address for storage data network (Controller VM interface). 
* `hxdp_mgmt_ip`:(string) IP address for HyperFlex management network. 
* `hxdp_storage_client_ip`:(string) IP address for storage client network (Controller VM interface). 
* `hypervisor_control_ip`:(string) IP address for hypervisor control such as VM migration or pod management. 
* `hypervisor_data_ip`:(string) IP address for storage data network (Hypervisor interface). 
* `hypervisor_mgmt_ip`:(string) IP address for Hypervisor management network. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `node_role`:(string) The role that this node performs in the HyperFlex cluster.* `Unknown` - The node role is not available.* `Storage` - The node persists data and contributes to the storage capacity of a cluster.* `Compute` - The node contributes to the compute capacity of a cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
