---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_io_card_operation"
description: |-
        Models the configurable properties of a iomodule in Intersight.

---

# Data Source: intersight_equipment_io_card_operation
Models the configurable properties of a iomodule in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_io_card_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_peer_power_state`:(string) User configured power state of the peer IO module.* `None` - Placeholder default value for iom power state property.* `Reboot` - IO Module reboot state property value. 
* `admin_power_state`:(string) User configured power state of the IO module.* `None` - Placeholder default value for iom power state property.* `Reboot` - IO Module reboot state property value. 
* `affected_obj_name`:(string) Placeholder for affected object name which is a combination of chassis and IOM ID. Used to display affected object in audit log. 
* `config_state`:(string) The configured state of these settings in the target IO module. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target IO module. Applying - This state denotes that the settings are being applied in the target IO module. Failed - This state denotes that the settings could not be applied in the target IO module.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
