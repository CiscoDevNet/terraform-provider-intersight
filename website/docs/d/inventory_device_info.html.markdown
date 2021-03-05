---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_device_info"
description: |-
  Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
---

# Data Source: intersight_inventory_device_info
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_device_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_state`:(string) Configuration state of server profile config context. 
* `control_action`:(string) Control action of server profile config context. 
* `error_state`:(string) Error state of server profile config context. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of server profile config context. 
* `profile_mo_id`:(string) Server profile MO ID of the server. 
 