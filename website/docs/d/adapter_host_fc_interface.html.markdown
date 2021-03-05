---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_host_fc_interface"
description: |-
  Host facing fibre channel interface on a server adapter.
---

# Data Source: intersight_adapter_host_fc_interface
Host facing fibre channel interface on a server adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_host_fc_interface.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin Configured State of Host Fibre Channel Interface. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ep_dn`:(string) The Endpoint Config Dn of the Host Fibre Channel Interface. 
* `host_fc_interface_id`:(int) Identifier of Host Fibre Channel Interface. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of Host Fibre Channel Interface. 
* `oper_state`:(string) Operational State of Host Fibre Channel Interface. 
* `operability`:(string) Operability status of Host Fibre Channel Interface. 
* `original_wwnn`:(string) The uniquely distinguishable factory default  World Wide Node Name of the Host. 
* `original_wwpn`:(string) The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface. 
* `peer_dn`:(string) PeerPort Dn of Host Fibre Channel Interface. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `wwnn`:(string) The uniquely distinguishable user configured World Wide Node Name of the Host. 
* `wwpn`:(string) The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface. 
 