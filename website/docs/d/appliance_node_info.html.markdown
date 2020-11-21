
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_info"
sidebar_current: "docs-intersight-data-source-appliance-node-info"
description: |-
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
---

# Data Source: intersight_appliance._node_info
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `hostname`:(string) Cluster node's FQDN or IP address. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
