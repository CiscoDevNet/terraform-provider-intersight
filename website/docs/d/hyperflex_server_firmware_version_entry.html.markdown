---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_server_firmware_version_entry"
description: |-
  An entry specifying supported server firmware version in regex format.
---

# Data Source: intersight_hyperflex_server_firmware_version_entry
An entry specifying supported server firmware version in regex format.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_server_firmware_version_entry.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_platform`:(string) The server platform type that is applicable for the server firmware bundle version.* `M5` - M5 generation of UCS server.* `M4` - M4 generation of UCS server. 
* `nr_version`:(string) The server firmware bundle version. 
 