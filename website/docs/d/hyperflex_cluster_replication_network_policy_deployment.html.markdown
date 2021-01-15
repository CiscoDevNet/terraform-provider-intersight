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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cluster_uuid`:(string) Uuid of the HyperFlex cluster. 
* `description`:(string) Description from corresponding ClusterReplicationNetworkPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name from corresponding ClusterReplicationNetworkPolicy. 
* `policy_moid`:(string) Deployed network policy moid. 
* `profile_moid`:(string) Deployed cluster profile moid. 
* `replication_bandwidth_mbps`:(int) Bandwidth for the Replication network in Mbps. 
* `replication_mtu`:(int) MTU for the Replication network. 
* `request_id`:(string) Unique request ID allowing retry of the same logical request following a transient communication failure. 
