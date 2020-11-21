
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `sched_time`:(string) The time at which the regression model needs to run for all the metrics specified in catalog. 
* `version`:(string) The catalog version used in forecast configuration service. 
