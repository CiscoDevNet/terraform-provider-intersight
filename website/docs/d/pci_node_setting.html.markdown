---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_node_setting"
description: |-
        Models the configurable properties of a PCIe Node in Intersight.

---

# Data Source: intersight_pci_node_setting
Models the configurable properties of a PCIe Node in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_node_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_action`:(string) User configured action on the PCIe Node.* `None` - Placeholder default value for PCIe Node admin state property.* `Reboot` - PCIe node reboot state property value. 
* `admin_locator_led_state`:(string) User configured state of the locator LED for the PCIe Node.* `None` - No operation property for locator led.* `On` - The Locator Led is turned on.* `Off` - The Locator Led is turned off. 
* `config_state`:(string) The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Scheduled` - User configured settings are scheduled to be applied.* `Failed` - User configured settings could not be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The property used to identify the PCIe node it is associated with. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
