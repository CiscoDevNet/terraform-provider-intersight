
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `data_source`:(string) Data source from where we get the data for the metrics to compute regression model. For example Druid. 
* `metric_name`:(string) Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage. 
* `min_num_of_days_of_data`:(int) Minimum number of days of data required for computing forecast model. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_days_of_historical_data`:(int) Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `platform_type`:(string) The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement. 
