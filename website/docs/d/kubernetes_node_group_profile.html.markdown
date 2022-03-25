---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_node_group_profile"
description: |-
        A configuration profile for a node group in a Kubernetes cluster.

---

# Data Source: intersight_kubernetes_node_group_profile
A configuration profile for a node group in a Kubernetes cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_node_group_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `currentsize`:(int) Current number of nodes in this node group at any given point in time. 
* `description`:(string) Description of the profile. 
* `desiredsize`:(int) Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `maxsize`:(int) Maximum number of nodes this node group can scale up to during repair, replacement or upgrade operations. 
* `minsize`:(int) Minimum number of available nodes this node group can scale down to during repair, replacement or upgrade operations. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `node_type`:(string) The node type ControlPlane, Worker or ControlPlaneWorker.* `Worker` - Node will be marked as a worker node.* `ControlPlane` - Node will be marked as a control plane node.* `ControlPlaneWorker` - Node will be both a controle plane and a worker. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
