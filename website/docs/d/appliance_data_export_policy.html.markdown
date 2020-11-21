
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_data_export_policy"
sidebar_current: "docs-intersight-data-source-appliance-data-export-policy"
description: |-
Data Export Policy is a category-based data collection policy that enables or disables
data export (data collection) from the Intersight Appliance to the Intersight. The Data
Export Policy configuration is organized hierarchically as follows.
  Global:
     Inventory:
        Network
        Storage
     TechSupport
When the DataExportPolicy for a category is enabled/disabled, all the sub-category configurations
are enabled/disabled as well. For example, if you enable/disable Inventory, all its sub-category
configurations (ie. Network and Storage) are also enabled/disabled.
---

# Data Source: intersight_appliance._data_export_policy
Data Export Policy is a category-based data collection policy that enables or disables
data export (data collection) from the Intersight Appliance to the Intersight. The Data
Export Policy configuration is organized hierarchically as follows.
  Global:
     Inventory:
        Network
        Storage
     TechSupport
When the DataExportPolicy for a category is enabled/disabled, all the sub-category configurations
are enabled/disabled as well. For example, if you enable/disable Inventory, all its sub-category
configurations (ie. Network and Storage) are also enabled/disabled.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Data Export Policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
