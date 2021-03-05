---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_lun"
description: |-
  A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
---

# Data Source: intersight_storage_pure_host_lun
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_host_lun.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_group_name`:(string) Name of the host group associated with LUN. 
* `host_name`:(string) Name of the host associated with LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared`:(bool) Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 