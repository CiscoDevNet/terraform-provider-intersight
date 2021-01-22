---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_initiator_group"
description: |-
  NetApp Initiator Group specifies host access to LUNs on the storage system.
---

# Data Source: intersight_storage_net_app_initiator_group
NetApp Initiator Group specifies host access to LUNs on the storage system.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Short description about the host. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
* `protocol`:(string) Initiator group protocol.* `FCP` - Fibre channel initiator type which contains WWN of an HBA on the host.* `iSCSI` - An iSCSI initiator type used by the host.* `mixed` - For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups. 
* `uuid`:(string) UUID of the LUN. 
