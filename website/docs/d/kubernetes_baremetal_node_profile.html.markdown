---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_baremetal_node_profile"
description: |-
        A profile specifying configuration settings for a baremetal node. Users can do operations like Drain, Cordon, Rebuild on a node.

---

# Data Source: intersight_kubernetes_baremetal_node_profile
A profile specifying configuration settings for a baremetal node. Users can do operations like Drain, Cordon, Rebuild on a node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_baremetal_node_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `cloud_provider`:(string) Cloud provider for this node profile.* `noProvider` - Enables the use of no cloud provider.* `external` - Out of tree cloud provider, e.g. CPI for vsphere. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `kubernetes_nic`:(string) Network interface from NetworkInfo (by name) to use for kubernetes VIP. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
