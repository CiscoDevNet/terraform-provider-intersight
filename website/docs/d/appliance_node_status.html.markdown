
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_status"
sidebar_current: "docs-intersight-data-source-appliance-node-status"
description: |-
Status of an Intersight Appliance node.
---

# Data Source: intersight_appliance._node_status
Status of an Intersight Appliance node.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cpu_usage`:(float) Percentage of CPU currently in use. 
* `mem_usage`:(float) Percentage of memory currently in use. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node.The system incrementally assigns identifiers to each node inthe Intersight Appliance cluster starting with a value of 1. 
* `node_state`:(string) State of the node in terms of its readiness to host Kubernetes pods.* `Down` - The node is yet to come up and join as a member of theKubernetes cluster.* `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods.* `Ready` - The node is ready to host Kubernetes pods. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
