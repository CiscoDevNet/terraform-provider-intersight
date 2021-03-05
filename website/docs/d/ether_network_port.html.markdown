---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_network_port"
description: |-
  Model contains the details of the ethernet port connected to the FI side.
---

# Data Source: intersight_ether_network_port
Model contains the details of the ethernet port connected to the FI side.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ether_network_port.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `module_id`:(int) Febric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of an Interface. 
* `peer_dn`:(string) Peer DN for network host port of fabric extender. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Network Port Speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
 