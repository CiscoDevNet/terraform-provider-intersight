---
subcategory: "fc"
layout: "intersight"
page_title: "Intersight: intersight_fc_physical_port"
description: |-
  Physical fibre channel port present on a FI.
---

# Data Source: intersight_fc_physical_port
Physical fibre channel port present on a FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fc_physical_port.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Administrator configured Speed applied on the port. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port. 
* `b2b_credit`:(int) Buffer to Buffer credits of FC port. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `max_speed`:(string) Maximum Speed with which the port operates. 
* `mode`:(string) Mode information N_proxy, F or E associated to the Fibre Channel port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_speed`:(string) Operational Speed with which the port operates. 
* `oper_state`:(string) Operational state of this port (enabled/disabled). 
* `oper_state_qual`:(string) Reason for this port's Operational state. 
* `peer_dn`:(string) PeerDn for fibre channel physical port. 
* `port_channel_id`:(int) Port channel id of FC port channel created on FI switch. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The role assigned to this port. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `transceiver_type`:(string) Transceiver type of a Fibre Channel port. 
* `vsan`:(int) Virtual San that is associated to the port. 
* `wwn`:(string) World Wide Name of a Fibre Channel port. 
 