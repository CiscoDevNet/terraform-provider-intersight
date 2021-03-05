---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_host_lun"
description: |-
  A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
---

# Data Source: intersight_storage_hitachi_host_lun
A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_host_lun.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(string) Port ID of the Hitachi host lun. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 