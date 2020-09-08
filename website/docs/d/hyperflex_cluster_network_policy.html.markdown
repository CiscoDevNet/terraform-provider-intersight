
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_network_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-cluster-network-policy"
description: |-
A policy specifying VLANs for management, VM migration, and VM traffic.
---

# Data Source: intersight_hyperflex._cluster_network_policy
A policy specifying VLANs for management, VM migration, and VM traffic.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `jumbo_frame`:(bool) Enable or disable jumbo frames. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `uplink_speed`:(string) Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.* `default` - Current default value set on the hardware platform.* `1G` - A link speed of 1 gigabit per second.* `10G` - A link speed of 10 gigabits per second or above. 
