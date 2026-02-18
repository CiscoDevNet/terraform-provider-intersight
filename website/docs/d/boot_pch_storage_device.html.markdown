---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_pch_storage_device"
description: |-
        The boot.PchStorageDevice object represents a storage device connected to the Platform Controller Hub (PCH) as a configurable boot option within a server's boot policy. This typically refers to onboard SATA or M.2 devices managed by the chipset.
        #### Purpose
        The PchStorageDevice object provides a standardized way to include PCH-attached storage, such as an M.2 SATA drive, in a server's boot order. It allows administrators to define the device's boot priority and enabled state within a managed boot policy.
        #### Key Concepts
        - **Onboard Storage Boot:** Defines a device connected directly to the motherboard's PCH as a bootable source.
        - **Configurable Priority:** The order property determines its position in the boot sequence relative to other devices like RAID controllers or PCIe drives.
        - **State Management:** The state property allows the PCH storage device to be enabled or disabled as a boot option within a policy.
        - **Boot Policy Component:** Functions as an element within a boot policy that is associated with a compute.Physical server object.

---

# Data Source: intersight_boot_pch_storage_device
The boot.PchStorageDevice object represents a storage device connected to the Platform Controller Hub (PCH) as a configurable boot option within a server's boot policy. This typically refers to onboard SATA or M.2 devices managed by the chipset.
#### Purpose
The PchStorageDevice object provides a standardized way to include PCH-attached storage, such as an M.2 SATA drive, in a server's boot order. It allows administrators to define the device's boot priority and enabled state within a managed boot policy.
#### Key Concepts
- **Onboard Storage Boot:** Defines a device connected directly to the motherboard's PCH as a bootable source.
- **Configurable Priority:** The order property determines its position in the boot sequence relative to other devices like RAID controllers or PCIe drives.
- **State Management:** The state property allows the PCH storage device to be enabled or disabled as a boot option within a policy.
- **Boot Policy Component:** Functions as an element within a boot policy that is associated with a compute.Physical server object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_pch_storage_device.<custom_name>.results[i].<propertyname>`.
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
 
