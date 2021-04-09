---
subcategory: "forecast"
layout: "intersight"
page_title: "Intersight: intersight_forecast_definition"
description: |-
  Definition for forecast metric settings.
---

# Data Source: intersight_forecast_definition
Definition for forecast metric settings.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_forecast_definition.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alert_threshold_in_percentage`:(int) Threshold above which user needs to be indicated through alarm/alert. 
* `create_time`:(string) The time when this managed object was created. 
* `data_source`:(string) Data source from where we get the data for the metrics to compute regression model. For example Druid. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `metric_name`:(string) Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage. 
* `min_num_of_days_of_data`:(int) Minimum number of days of data required for computing forecast model. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_days_of_historical_data`:(int) Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model. 
* `platform_type`:(string) The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 