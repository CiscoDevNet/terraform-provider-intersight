
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `local_configuration_locked`:(bool) Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled. 
* `log_level`:(string) The log level of the device connector service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
