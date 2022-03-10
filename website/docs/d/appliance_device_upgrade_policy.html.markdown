---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_upgrade_policy"
description: |-
  DeviceUpgradePolicy stores the current upgrade policy of the Intersight Appliances.
---

# Data Source: intersight_appliance_device_upgrade_policy
DeviceUpgradePolicy stores the current upgrade policy of the Intersight Appliances.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_device_upgrade_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_upgrade`:(bool) Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored. 
* `blackout_dates_enabled`:(bool) If enabled, allows the user to define a blackout period during which the appliance will not be upgraded. 
* `blackout_end_date`:(string) End date of the black out period. 
* `blackout_start_date`:(string) Start date of the black out period. The appliance will not be upgraded during this period. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_meta_data_sync`:(bool) Indicates if the updated metadata files should be synced immediately or at the next upgrade. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `serial_id`:(string) SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `software_download_type`:(string) UpgradeType is used to indicate the kink of software upload to upgrade.* `unknown` - Indicates user setting of upgrade service to unknown.* `connected` - Indicates if the upgrade service is set to upload software to latest version automatically.* `manual` - Indicates if the upgrade service is set to upload software to user picked verison manually. 
 