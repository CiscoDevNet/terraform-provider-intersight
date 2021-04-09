---
subcategory: "forecast"
layout: "intersight"
page_title: "Intersight: intersight_forecast_instance"
description: |-
  Entity representing forecast result for instance of managed object, ie, data source.
---

# Data Source: intersight_forecast_instance
Entity representing forecast result for instance of managed object, ie, data source.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_forecast_instance.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) The Moid of the Intersight managed device instance for which regression model is derived. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `full_cap_days`:(int) The number of days remaining before the device reaches its full functional capacity. 
* `last_model_update_time`:(string) The time when the forecast model was last updated. 
* `metric_name`:(string) The name of the metric for which regression model is generated. 
* `min_days_for_forecast`:(int) The minimum number of days the HyperFlex cluster should be up for computing forecast. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `threshold_days`:(int) The number of days remaining before the device reaches the specified threshold for the metric as defined in definition. 
 