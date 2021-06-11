---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_precision_policy"
description: |-
  Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
---

# Resource: intersight_boot_precision_policy
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
## Usage Example
### Resource Creation

```hcl
resource "intersight_boot_precision_policy" "boot_precision1" {
  name                     = "boot_precision1"
  description              = "test policy"
  configured_boot_mode     = "Legacy"
  enforce_uefi_secure_boot = false
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  boot_devices {
    enabled     = true
    name        = "scu-device-hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  boot_devices {
    enabled     = true
    name        = "NIIODCIMCDVD"
    object_type = "boot.VirtualMedia"
    additional_properties = jsonencode({
      Subtype = "cimc-mapped-dvd"
    })
  }
  boot_devices {
    enabled     = true
    name        = "hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `boot_devices`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [boot.Iscsi](#bootIscsi)
[boot.LocalCdd](#bootLocalCdd)
[boot.LocalDisk](#bootLocalDisk)
[boot.Nvme](#bootNvme)
[boot.PchStorage](#bootPchStorage)
[boot.Pxe](#bootPxe)
[boot.San](#bootSan)
[boot.SdCard](#bootSdCard)
[boot.UefiShell](#bootUefiShell)
[boot.Usb](#bootUsb)
[boot.VirtualMedia](#bootVirtualMedia)
  + `enabled`:(bool) Specifies if the boot device is enabled or disabled. 
  + `name`:(string) A name that helps identify a boot device. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `configured_boot_mode`:(string) Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `enforce_uefi_secure_boot`:(bool) If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_boot_precision_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_boot_precision_policy.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [boot.Iscsi](#argument-reference)
Device type used when booting from iSCSI boot device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during iSCSI boot. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 
* `interface_name`:(string) The name of the underlying virtual ethernet interface used by the iSCSI boot device. 
* `port`:(int) Port ID of the ISCSI boot device. 
* `slot`:(string) The slot id of the device. Supported values are (1 - 255, \ MLOM\ , \ L\ , \ L1\ , \ L2\ , \ OCP\ ). 

### [boot.LocalCdd](#argument-reference)
Device type used when booting from local CD drive.

### [boot.LocalDisk](#argument-reference)
Device type used when booting from a local disk device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during boot from local disk. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 
* `slot`:(string) The slot id of the local disk device. Supported values are (1-255, \ M\ , \ HBA\ , \ SAS\ , \ RAID\ , \ MRAID\ , \ MRAID1\ , \ MRAID2\ , \ MSTOR-RAID\ ). 

### [boot.Nvme](#argument-reference)
Device type used when booting from a NVMe device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during NVME boot. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 

### [boot.PchStorage](#argument-reference)
Device type used when booting from a PCH Storage device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during PCH Storage boot. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 
* `lun`:(int) The Logical Unit Number (LUN) of the device. It is the Virtual Drive number for Cisco UCS C-Series Servers. Virtual Drive number is found in storage inventory. 

### [boot.Pxe](#argument-reference)
Device type used when booting from a PXE boot device.
* `interface_name`:(string) The name of the underlying virtual ethernet interface used by the PXE boot device. 
* `interface_source`:(string) Lists the supported Interface Source for PXE device. Supported values are \ name\  and \ mac\ .* `name` - Use interface name to select virtual ethernet interface.* `mac` - Use MAC address to select virtual ethernet interface.* `port` - Use port to select virtual ethernet interface. 
* `ip_type`:(string) The IP Address family type to use during the PXE Boot process.* `None` - Default value if IpType is not specified.* `IPv4` - The IPv4 address family type.* `IPv6` - The IPv6 address family type. 
* `mac_address`:(string) The MAC Address of the underlying virtual ethernet interface used by the PXE boot device. 
* `port`:(int) The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are -1 to 255. 
* `slot`:(string) The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \ MLOM\ , \ L\ , \ L1\ , \ L2\ , \ OCP\ ). 

### [boot.San](#argument-reference)
Device type used when booting from SAN Boot device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during SAN boot. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 
* `interface_name`:(string) The name of the underlying vHBA interface to be used by the SAN boot device. 
* `lun`:(int) The Logical Unit Number (LUN) of the device. 
* `slot`:(string) Slot ID of the device. Supported values are ( 1 - 255, \ MLOM\ , \ L1\ , \ L2\  ). 
* `wwpn`:(string) The WWPN Address of the underlying fiber channel interface used by the SAN boot device. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. 

### [boot.SdCard](#argument-reference)
Device type used when booting from SD Card device.
* `bootloader`:(HashMap) - Details of the bootloader to be used during SD card boot. 
This complex property has following sub-properties:
  + `description`:(string) Carries more information about the bootloader. 
  + `name`:(string) Name of the bootloader image. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `path`:(string) Path to the bootloader image. 
* `lun`:(int) The Logical Unit Number (LUN) of the device. 
* `subtype`:(string) The subtype for the selected device type.* `None` - No sub type for SD card boot device.* `flex-util` - Use of FlexUtil (microSD) card as sub type for SD card boot device.* `flex-flash` - Use of FlexFlash (SD) card as sub type for SD card boot device.* `SDCARD` - Use of SD card as sub type for the SD Card boot device. 

### [boot.UefiShell](#argument-reference)
Device type used when booting from UEFI shell.

### [boot.Usb](#argument-reference)
Device type used when booting from USB device.
* `subtype`:(string) The subtype for the selected device type.* `None` - No sub type for USB boot device.* `usb-cd` - Use of Compact Disk (CD) as sub-type for the USB boot device.* `usb-fdd` - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device.* `usb-hdd` - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device. 

### [boot.VirtualMedia](#argument-reference)
Device type used when booting from Virtual Media device.
* `subtype`:(string) The subtype for the selected device type.* `None` - No sub type for virtual media.* `cimc-mapped-dvd` - The virtual media device is mapped to a virtual DVD device.* `cimc-mapped-hdd` - The virtual media device is mapped to a virtual HDD device.* `kvm-mapped-dvd` - A KVM mapped DVD virtual media device.* `kvm-mapped-hdd` - A KVM mapped HDD virtual media device.* `kvm-mapped-fdd` - A KVM mapped FDD virtual media device. 
  