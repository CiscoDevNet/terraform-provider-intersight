
---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_configuration"
sidebar_current: "docs-intersight-data-source-asset-device-configuration"
description: |-
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
---

# Data Source: intersight_asset._device_configuration
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `local_configuration_locked`:(bool) Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled. 
* `log_level`:(string) The log level of the device connector service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
