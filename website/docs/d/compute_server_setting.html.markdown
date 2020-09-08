
---
layout: "intersight"
page_title: "Intersight: intersight_compute_server_setting"
sidebar_current: "docs-intersight-data-source-compute-server-setting"
description: |-
Models the configurable properties of a server in Intersight.
---

# Data Source: intersight_compute._server_setting
Models the configurable properties of a server in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_state`:(string) User configured state of the locator LED for the server.* `None` - No operation property for locator led.* `On` - The Locator Led is turned on.* `Off` - The Locator Led is turned off. 
* `admin_power_state`:(string) User configured power state of the server.* `Policy` - Power state is set to the default value in the policy.* `PowerOn` - Power state of the server is set to On.* `PowerOff` - Power state is the server set to Off.* `PowerCycle` - Power state the server is reset.* `HardReset` - Power state the server is hard reset.* `Shutdown` - Operating system on the server is shut down.* `Reboot` - Power state of IMC is rebooted. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_state`:(string) The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `one_time_boot_device`:(string) The name of the device chosen by user for configuring One-Time Boot device. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
