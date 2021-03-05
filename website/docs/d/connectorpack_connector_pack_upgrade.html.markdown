---
subcategory: "connectorpack"
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_connector_pack_upgrade"
description: |-
  Used to download or install connector packs on the target device.
---

# Data Source: intersight_connectorpack_connector_pack_upgrade
Used to download or install connector packs on the target device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_connectorpack_connector_pack_upgrade.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 