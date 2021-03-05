---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_replication_network_policy"
description: |-
  Specifies all replication network parameters for the cluster.
---

# Data Source: intersight_hyperflex_cluster_replication_network_policy
Specifies all replication network parameters for the cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_replication_network_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `replication_bandwidth_mbps`:(int) Bandwidth for the Replication network in Mbps. 
* `replication_mtu`:(int) MTU for the Replication network. 
 