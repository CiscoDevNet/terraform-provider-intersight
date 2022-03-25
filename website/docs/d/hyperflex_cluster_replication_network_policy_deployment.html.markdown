---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_replication_network_policy_deployment"
description: |-
        Record of HyperFlex Cluster replication network policy deployment.

---

# Data Source: intersight_hyperflex_cluster_replication_network_policy_deployment
Record of HyperFlex Cluster replication network policy deployment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_replication_network_policy_deployment.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_uuid`:(string) Uuid of the HyperFlex cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description from corresponding ClusterReplicationNetworkPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name from corresponding ClusterReplicationNetworkPolicy. 
* `policy_moid`:(string) Deployed network policy moid. 
* `profile_moid`:(string) Deployed cluster profile moid. 
* `replication_bandwidth_mbps`:(int) Bandwidth for the Replication network in Mbps. 
* `replication_mtu`:(int) MTU for the Replication network. 
* `request_id`:(string) Unique request ID allowing retry of the same logical request following a transient communication failure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
