---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_device_configuration"
description: |-
  The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
---

# Data Source: intersight_asset_device_configuration
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_device_configuration.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `local_configuration_locked`:(bool) Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled. 
* `log_level`:(string) The log level of the device connector service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 