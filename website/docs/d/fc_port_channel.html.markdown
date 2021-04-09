---
subcategory: "fc"
layout: "intersight"
page_title: "Intersight: intersight_fc_port_channel"
description: |-
  Model contains the details of the ethernet port-channels configured on the FI.
---

# Data Source: intersight_fc_port_channel
Model contains the details of the ethernet port-channels configured on the FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fc_port_channel.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_speed`:(string) Administrator configured Speed applied on the port channel. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this portchannel. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Mode information N_proxy, F or E associated to the Fibre Channel portchannel. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_speed`:(string) Operational speed of this port-channel. 
* `oper_state`:(string) Operational state of this port-channel. 
* `oper_state_qual`:(string) Reason for this port-channel's Operational state. 
* `port_channel_id`:(int) Unique identifier for this port-channel on the FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) This port-channel's configured role (fcUplink, fcStorage, etc.). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `vsan`:(int) Virtual San that is associated to the port-channel. 
 