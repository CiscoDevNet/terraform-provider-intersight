
---
layout: "intersight"
page_title: "Intersight: intersight_forecast_instance"
sidebar_current: "docs-intersight-data-source-forecast-instance"
description: |-
Entity representing forecast result for instance of managed object, ie, data source.
---

# Data Source: intersight_forecast._instance
Entity representing forecast result for instance of managed object, ie, data source.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_id`:(string) The Moid of the Intersight managed device instance for which regression model is derived. 
* `full_cap_days`:(int) The number of days remaining before the device reaches its full functional capacity. 
* `metric_name`:(string) The name of the metric for which regression model is generated. 
* `min_days_for_forecast`:(int) The minimum number of days the HyperFlex cluster should be up for computing forecast. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `threshold_days`:(int) The number of days remaining before the device reaches the specified threshold for the metric as defined in definition. 
