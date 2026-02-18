---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_vmedia_device"
description: |-
        The boot.VmediaDevice object represents a virtual media device as a configurable boot option within a server's boot policy. Virtual media allows mounting of a remote image (like an ISO or IMG file) as if it were a local physical device (like a CD/DVD or USB drive).
        #### Purpose
        The VmediaDevice object is used to enable booting from a remotely mounted image file. It is crucial for OS installation, recovery operations, and diagnostics without needing physical access to the server. It allows administrators to define virtual media as a boot source and set its priority within a boot policy.
        #### Key Concepts
        - **Remote Media Boot:** Defines a virtual CD/DVD or USB drive, sourced from a remote image file, as a bootable device.
        - **Flexible OS Installation:** Facilitates the installation of operating systems and software from ISO images without physical media.
        - **Ordered Boot Sequence:** The order property allows its position in the boot process to be controlled relative to other physical or network devices.
        - **Boot Policy Component:** Acts as a key element within a boot policy that is applied to a compute.Physical server object.

---

# Data Source: intersight_boot_vmedia_device
The boot.VmediaDevice object represents a virtual media device as a configurable boot option within a server's boot policy. Virtual media allows mounting of a remote image (like an ISO or IMG file) as if it were a local physical device (like a CD/DVD or USB drive).
#### Purpose
The VmediaDevice object is used to enable booting from a remotely mounted image file. It is crucial for OS installation, recovery operations, and diagnostics without needing physical access to the server. It allows administrators to define virtual media as a boot source and set its priority within a boot policy.
#### Key Concepts
- **Remote Media Boot:** Defines a virtual CD/DVD or USB drive, sourced from a remote image file, as a bootable device.
- **Flexible OS Installation:** Facilitates the installation of operating systems and software from ISO images without physical media.
- **Ordered Boot Sequence:** The order property allows its position in the boot process to be controlled relative to other physical or network devices.
- **Boot Policy Component:** Acts as a key element within a boot policy that is applied to a compute.Physical server object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_vmedia_device.<custom_name>.results[i].<propertyname>`.
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
 
