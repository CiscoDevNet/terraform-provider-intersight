---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host"
description: |-
  A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
---

# Data Source: intersight_storage_pure_host
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Short description about the host. 
* `host_group_name`:(string) Name of host group where the host belongs to. Empty if host is not part of any HostGroup. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
