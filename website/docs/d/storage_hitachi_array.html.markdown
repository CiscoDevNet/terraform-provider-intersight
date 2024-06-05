---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_array"
description: |-
        The details of the Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_array
The details of the Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_array.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `ctl1_ip`:(string) IP address of controller 1 of the storage system. 
* `ctl1_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 1. 
* `ctl2_ip`:(string) IP address of controller 2 of the storage system. 
* `ctl2_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 2. 
* `device_id`:(string) ID of the Storage device. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_type`:(string) The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hardware_version`:(string) The hardware version of the device. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svp_ip`:(string) IP address of the SVP (Service Processor). The SVP provides out-of-band configuration and management of the storage system, and collects performance data for key components to enable diagnostic testing and analysis. 
* `target_ctl`:(string) Controller operated by the REST API. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Current running software version of the device. 
 
