
---
layout: "intersight"
page_title: "Intersight: intersight_forecast_definition"
sidebar_current: "docs-intersight-data-source-forecast-definition"
description: |-
Definition for forecast metric settings.
---

# Data Source: intersight_forecast._definition
Definition for forecast metric settings.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `alert_threshold_in_percentage`:(int) Threshold above which user needs to be indicated through alarm/alert. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `data_source`:(string) Data source from where we get the data for the metrics to compute regression model. For example Druid. 
* `metric_name`:(string) Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage. 
* `min_num_of_days_of_data`:(int) Minimum number of days of data required for computing forecast model. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_days_of_historical_data`:(int) Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `platform_type`:(string) The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement. 
