
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `config_state`:(string) Configuration state of server profile config context. 
* `control_action`:(string) Control action of server profile config context. 
* `error_state`:(string) Error state of server profile config context. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `oper_state`:(string) Operational state of server profile config context. 
* `profile_mo_id`:(string) Server profile MO ID of the server. 
