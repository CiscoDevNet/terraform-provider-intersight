
---
layout: "intersight"
page_title: "Intersight: intersight_forecast_catalog"
sidebar_current: "docs-intersight-data-source-forecast-catalog"
description: |-
A catalog for managing forecast settings.
---

# Data Source: intersight_forecast._catalog
A catalog for managing forecast settings.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `sched_time`:(string) The time at which the regression model needs to run for all the metrics specified in catalog. 
* `version`:(string) The catalog version used in forecast configuration service. 
