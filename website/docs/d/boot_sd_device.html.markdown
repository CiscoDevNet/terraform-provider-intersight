---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_sd_device"
description: |-
        The boot.SdDevice object represents an SD (Secure Digital) card as a configurable boot option within a server's boot policy. This typically refers to an internal SD card slot, often used for booting hypervisors.
        #### Purpose
        The SdDevice object provides a standardized way to include an SD card in a server's boot order. It allows administrators to define the SD card as a bootable source, set its priority, and manage its state (enabled/disabled) within a boot policy. It is particularly useful for embedded operating systems or hypervisor installations.
        #### Key Concepts
        - **Embedded OS Boot:** Defines an internal SD card as a bootable source, commonly used for lightweight hypervisors like VMware ESXi.
        - **Ordered Boot Sequence:** The order property allows its position in the boot process to be controlled relative to other devices.
        - **State Management:** The state property allows the SD card to be easily enabled or disabled as a boot option within a policy.
        - **Boot Policy Component:** Acts as an element within a boot policy that is applied to a compute.Physical server object.

---

# Data Source: intersight_boot_sd_device
The boot.SdDevice object represents an SD (Secure Digital) card as a configurable boot option within a server's boot policy. This typically refers to an internal SD card slot, often used for booting hypervisors.
#### Purpose
The SdDevice object provides a standardized way to include an SD card in a server's boot order. It allows administrators to define the SD card as a bootable source, set its priority, and manage its state (enabled/disabled) within a boot policy. It is particularly useful for embedded operating systems or hypervisor installations.
#### Key Concepts
- **Embedded OS Boot:** Defines an internal SD card as a bootable source, commonly used for lightweight hypervisors like VMware ESXi. 
- **Ordered Boot Sequence:** The order property allows its position in the boot process to be controlled relative to other devices.
- **State Management:** The state property allows the SD card to be easily enabled or disabled as a boot option within a policy.
- **Boot Policy Component:** Acts as an element within a boot policy that is applied to a compute.Physical server object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_sd_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the boot device configured in the boot policy. 
* `order`:(int) The order of the boot device configured in the boot policy. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the boot device configured in the boot policy. 
* `type`:(string) The type of the boot device configured in the boot policy. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
