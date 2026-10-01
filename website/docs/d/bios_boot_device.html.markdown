---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_boot_device"
description: |-
        BootDevices represent the actual bootable devices enumerated by a server’s BIOS. Each BootDevice captures the device identity and type as the BIOS currently sees it, forming the per-entry list that makes up the system’s effective boot sequence.
        #### Purpose
        Provide a read-only inventory of BIOS-enumerated boot devices so administrators can verify what the system can boot from and troubleshoot boot-order or boot-device discovery issues.
        #### Key Concepts
        - **Observed (BIOS) reality vs policy intent**: This is the *actual* device list reported by BIOS, not a desired boot policy configuration.
        - **Device identity and classification**: `deviceName` identifies the boot entry; `deviceType` indicates what kind of boot device it is.
        - **System association**: `registeredDevice` links the boot device inventory to the specific managed server/device in Intersight.
        - **Lifecycle handling**: `onpeerdelete: unset` on `registeredDevice` preserves the BootDevice record semantics when the registration relationship is removed.
        - **Collection workflow support**: `GetActualBootOrderTask` is the async task responsible for retrieving the actual boot order information from the endpoint (with retries/timeouts for resilience).

---

# Data Source: intersight_bios_boot_device
BootDevices represent the actual bootable devices enumerated by a server’s BIOS. Each BootDevice captures the device identity and type as the BIOS currently sees it, forming the per-entry list that makes up the system’s effective boot sequence.
#### Purpose
Provide a read-only inventory of BIOS-enumerated boot devices so administrators can verify what the system can boot from and troubleshoot boot-order or boot-device discovery issues.
#### Key Concepts
- **Observed (BIOS) reality vs policy intent**: This is the *actual* device list reported by BIOS, not a desired boot policy configuration.
- **Device identity and classification**: `deviceName` identifies the boot entry; `deviceType` indicates what kind of boot device it is.
- **System association**: `registeredDevice` links the boot device inventory to the specific managed server/device in Intersight.
- **Lifecycle handling**: `onpeerdelete: unset` on `registeredDevice` preserves the BootDevice record semantics when the registration relationship is removed.
- **Collection workflow support**: `GetActualBootOrderTask` is the async task responsible for retrieving the actual boot order information from the endpoint (with retries/timeouts for resilience).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_boot_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_name`:(string) Name of the Configured Boot Device. 
* `device_type`:(string) Type of the Configured Boot Device. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
