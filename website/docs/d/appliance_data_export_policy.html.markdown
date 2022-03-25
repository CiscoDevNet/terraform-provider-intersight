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
To access the ith object of the results obtained, use `data.intersight_appliance_data_export_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Data Export Policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
