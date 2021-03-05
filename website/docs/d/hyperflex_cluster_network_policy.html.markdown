---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_network_policy"
description: |-
  A policy specifying VLANs for management, VM migration, and VM traffic.
---

# Data Source: intersight_hyperflex_cluster_network_policy
A policy specifying VLANs for management, VM migration, and VM traffic.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_network_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `jumbo_frame`:(bool) Enable or disable jumbo frames. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `uplink_speed`:(string) Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.* `default` - Current default value set on the hardware platform.* `1G` - A link speed of 1 gigabit per second.* `10G` - A link speed of 10 gigabits per second or above. 
 