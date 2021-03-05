---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_data_export_policy"
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

# Data Source: intersight_appliance_data_export_policy
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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_data_export_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Data Export Policy. 
 