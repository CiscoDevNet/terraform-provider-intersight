---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_precision_policy"
description: |-
        The Boot PrecisionPolicy object is a reusable policy that defines a precise and ordered list of bootable devices for a server. It supports both UEFI and Legacy boot modes and provides granular control over the boot sequence.
        #### Purpose
        The primary purpose of this policy is to automate and enforce a specific boot order on servers. It is critical for ensuring that servers boot from the correct media, whether it is a local disk, a SAN LUN, a PXE server, or a virtual media device. By using a policy, administrators can guarantee consistent boot behavior for servers assigned to a Server Profile.
        #### Key Concepts
        - **UEFI and Legacy Support:** The policy allows administrators to specify the desired boot mode (Uefi or Legacy) and whether to enforce UEFI Secure Boot.
        - **Ordered Device List:** It contains an ordered list of boot devices, where each entry specifies the device type (e.g., LocalDisk, Pxe, San, Nvme) and its specific configuration.
        - **Profile-Based Application:** The policy is attached to a Server Profile to apply the defined boot order to a physical server.
        - **Reboot Required:** Changes to the boot order policy are flagged with ActivationRequiresReboot, as a server reboot is necessary to apply the new boot sequence.
        - **Dependent on Connectivity Policies:** The policy's effectiveness can depend on other policies, such as LAN and SAN connectivity policies, which configure the network interfaces used for PXE or SAN booting.

---

# Data Source: intersight_boot_precision_policy
The Boot PrecisionPolicy object is a reusable policy that defines a precise and ordered list of bootable devices for a server. It supports both UEFI and Legacy boot modes and provides granular control over the boot sequence.
#### Purpose
The primary purpose of this policy is to automate and enforce a specific boot order on servers. It is critical for ensuring that servers boot from the correct media, whether it is a local disk, a SAN LUN, a PXE server, or a virtual media device. By using a policy, administrators can guarantee consistent boot behavior for servers assigned to a Server Profile.
#### Key Concepts
- **UEFI and Legacy Support:** The policy allows administrators to specify the desired boot mode (Uefi or Legacy) and whether to enforce UEFI Secure Boot.
- **Ordered Device List:** It contains an ordered list of boot devices, where each entry specifies the device type (e.g., LocalDisk, Pxe, San, Nvme) and its specific configuration.
- **Profile-Based Application:** The policy is attached to a Server Profile to apply the defined boot order to a physical server.
- **Reboot Required:** Changes to the boot order policy are flagged with ActivationRequiresReboot, as a server reboot is necessary to apply the new boot sequence.
- **Dependent on Connectivity Policies:** The policy's effectiveness can depend on other policies, such as LAN and SAN connectivity policies, which configure the network interfaces used for PXE or SAN booting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_precision_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `configured_boot_mode`:(string) Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the MBR partitioning scheme. To apply this setting, Please reboot the server.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the MBR to locate the bootloader. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enforce_uefi_secure_boot`:(bool) If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
