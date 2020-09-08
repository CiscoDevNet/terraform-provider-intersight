
---
layout: "intersight"
page_title: "Intersight: intersight_inventory_device_info"
sidebar_current: "docs-intersight-data-source-inventory-device-info"
description: |-
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
---

# Data Source: intersight_inventory._device_info
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_state`:(string) Configuration state of server profile config context. 
* `control_action`:(string) Control action of server profile config context. 
* `error_state`:(string) Error state of server profile config context. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Operational state of server profile config context. 
* `profile_mo_id`:(string) Server profile MO ID of the server. 
