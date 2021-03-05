---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_device_package_download_state"
description: |-
  HyperFlex Device Package Download State.
---

# Data Source: intersight_hyperflex_device_package_download_state
HyperFlex Device Package Download State.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_device_package_download_state.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `checksum`:(string) Checksum of HyperFlex health check Debian package installed on the HyperFlex Device. 
* `hx_device_name`:(string) HyperFlex Device Name for which the package download state is tracked. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `timestamp`:(string) Timestamp of the last health check Debian package installation on the HyperFlex Device. 
 