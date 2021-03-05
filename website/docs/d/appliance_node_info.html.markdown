---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_info"
description: |-
  NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
---

# Data Source: intersight_appliance_node_info
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_node_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hostname`:(string) Cluster node's FQDN or IP address. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
 