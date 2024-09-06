---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_vedge_device"
description: |-
        Details for the Catalyst SDWAN Vedge entities.

---

# Data Source: intersight_catalystsdwan_vedge_device
Details for the Catalyst SDWAN Vedge entities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_vedge_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_status_message`:(string) The Catalyst SDWAN device config status message. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) The Catalyst SDWAN device id. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_state`:(string) The Catalyst SDWAN device state. 
* `device_type`:(string) The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hardware_version`:(string) The hardware version of the device. 
* `host_name`:(string) The Catalyst SDWAN device host name. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `platform_family`:(string) The Catalyst SDWAN device platform family. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `reachability`:(string) The Catalyst SDWAN device reachability. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(string) The Catalyst SDWAN device site id. 
* `site_name`:(string) The Catalyst SDWAN device site name. 
* `sp_organization_name`:(string) The Catalyst SDWAN device sp organization name. 
* `system_ip`:(string) The Catalyst SDWAN device system IP. 
* `template_status`:(string) The Catalyst SDWAN device template status. 
* `uuid`:(string) Unique identity of the device. 
* `validity`:(string) The Catalyst SDWAN device validity. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Current running software version of the device. 
 
