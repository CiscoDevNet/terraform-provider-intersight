---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_battery_backup_unit"
description: |-
        Information of Battery Backup Unit in the storage controller.

---

# Data Source: intersight_storage_battery_backup_unit
Information of Battery Backup Unit in the storage controller.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_battery_backup_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacitance_in_percent`:(int) This holds the capacitance (in percent) of the battery backup unit of the storage controller. 
* `charging_state`:(string) This holds the charging state of the battery backup unit of the storage controller. 
* `create_time`:(string) The time when this managed object was created. 
* `current_in_amps`:(float) This holds the current (in Amps) of the battery backup unit of the storage controller. 
* `design_capacity_in_joules`:(string) This holds the design Capacity (in joules) of the battery backup unit of the storage controller. 
* `design_voltage_in_volts`:(float) This holds the design volatage (in Volts) of the battery backup unit of the storage controller. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_name`:(string) This refers to the device name of the battery backup unit of the storage controller. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_battery_present`:(bool) This indicates whether the battery is present for the battery backup unit of the storage controller. 
* `is_capacitor`:(bool) This indicates the capacitor for the battery backup unit of the storage controller. 
* `is_learn_cycle_requested`:(bool) This indicates learn cycle request of the battery backup unit of the storage controller. 
* `is_learn_cycle_transparent`:(bool) This indicates the learn cycle transparent for the battery backup unit of the storage controller. 
* `is_temperature_high`:(bool) This indicates the temperature is high for the battery backup unit of the storage controller. 
* `is_voltage_low`:(bool) This indicates the voltage is Low for the battery backup unit of the storage controller. 
* `learn_cycle_progress_end_time_stamp`:(string) This refers to learn cycle progress end time of the battery backup unit of the storage controller. 
* `learn_cycle_progress_start_time_stamp`:(string) This refers to learn cycle progress start time of the battery backup unit of the storage controller. 
* `learn_cycle_progress_status`:(string) This refers to learn cycle progress status of the battery backup unit of the storage controller. 
* `learn_mode`:(string) This refers to the learn mode of the battery backup unit of the storage controller. 
* `manufacturing_date`:(string) This refers to the manufacture date of the battery backup unit of the storage controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `module_version`:(string) This refers to the current module version of the battery backup unit of the storage controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_learn_cycle_time_stamp`:(string) This refers to next learn cycle timestamp of the battery backup unit of the storage controller. 
* `pack_energy_in_joules`:(string) This holds the pack energy (in joules) of the battery backup unit of the storage controller. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `remaining_pool_space_in_percent`:(int) This holds the remaining pool space (in percent) of the battery backup unit of the storage controller. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) This holds the current status of the battery backup unit of the storage controller. 
* `temperature_in_cel`:(int) This holds the temperature (in Celsius) of the battery backup unit of the storage controller. 
* `type`:(string) This refers to the type of the battery backup unit of the storage controller. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `voltage_in_volts`:(string) This holds the volatage (in Volts) of the battery backup unit of the storage controller. 
 
