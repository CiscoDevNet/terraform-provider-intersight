---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_lun_map"
description: |-
  NetApp LUN mapping is the process of associating a LUN with an igroup. When a LUN is mapped to an igroup, initiators in the igroup are granted. access to the LUN.
---

# Data Source: intersight_storage_net_app_lun_map
NetApp LUN mapping is the process of associating a LUN with an igroup. When a LUN is mapped to an igroup, initiators in the igroup are granted. access to the LUN.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_lun_map.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `uuid`:(string) UUID of the LUN. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 