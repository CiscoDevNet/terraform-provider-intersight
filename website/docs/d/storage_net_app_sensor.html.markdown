---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_sensor"
description: |-
        Information for a particular sensor on a NetApp storage array controller.

---

# Data Source: intersight_storage_net_app_sensor
Information for a particular sensor on a NetApp storage array controller.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_sensor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_name`:(string) The name of the storage array controller that the sensor applies to. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of a specific sensor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the specified sensor. 
* `type`:(string) The type of the specified sensor. 
* `units`:(string) The units that correspond to the value of the sensor, if applicable. 
* `value`:(string) The value of the specified sensor. 
 
