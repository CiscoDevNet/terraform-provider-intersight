---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_server_setting"
description: |-
        Models the configurable properties of a server in Intersight.

---

# Data Source: intersight_compute_server_setting
Models the configurable properties of a server in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_server_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_locator_led_state`:(string) User configured state of the locator LED for the server.* `None` - No operation property for locator led.* `On` - The Locator Led is turned on.* `Off` - The Locator Led is turned off. 
* `admin_power_state`:(string) User configured power state of the server.* `Policy` - Power state is set to the default value in the policy.* `PowerOn` - Power state of the server is set to On.* `PowerOff` - Power state is the server set to Off.* `PowerCycle` - Power state the server is reset.* `HardReset` - Power state the server is hard reset.* `Shutdown` - Operating system on the server is shut down.* `Reboot` - Power state of IMC is rebooted. 
* `cmos_reset`:(string) The allowed actions on the CMOS Reset.* `Ready` - CMOS Reset operation is allowed to be done on the server in this state.* `Pending` - The identifier to state that the previous CMOS Reset operation on this server has not completed due to a pending power cycle. CMOS Reset operation cannot be done on the server when in this state.* `Reset` - The value that the UI/API needs to provide to trigger a CMOS Reset operation on a server. 
* `config_state`:(string) The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `front_panel_lock_state`:(string) The allowed actions on the Front Panel Lock.* `Unlock` - Front Panel of the server is set to Unlocked state.* `Lock` - Front Panel of the server is set to Locked state. 
* `kvm_reset`:(string) The allowed actions on the vKVM Reset.* `Ready` - Reset vKVM operation is allowed to be done on the server in this state.* `Reset` - The value that the UI/API needs to provide to trigger a Reset vKVM operation on a server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The property used to identify the name of the server it is associated with. 
* `one_time_boot_device`:(string) The name of the device chosen by user for configuring One-Time Boot device. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tunneled_kvm_state`:(string) By default, the tunneled vKVM property appears in Ready state. The property can be configured by performing allowed actions. Once the property is configured, it reverts to Ready state.* `Ready` - Tunneled vKVM is ready to be configured on the server.* `Enable` - Tunneled vKVM is enabled for the server.* `Disable` - Tunneled vKVM is disabled for the server. 
 
