---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_device_info"
description: |-
        Information pertaining to a Registered Device in Intersight which is pertinent to Inventory Microservice.

---

# Data Source: intersight_inventory_device_info
Information pertaining to a Registered Device in Intersight which is pertinent to Inventory Microservice.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_device_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `disable_events`:(bool) Subscribe/Unsubscribe events for the device. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `event_counter`:(int) Information regarding the number of events recorded for this device. 
* `event_counter_enabled`:(bool) Indicates whether the event counter enabled for the device. 
* `interval`:(int) The time interval (in hours) between job runs. 
* `last_re_inventory_time`:(string) Last Reinventory time of the device. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `re_inventory_count`:(int) Number of full inventory within the given time. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
