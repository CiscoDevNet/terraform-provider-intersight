---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_storage_policy"
description: |-
        The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The security of drives can be enabled through this policy using remote keys from a KMIP server or Manually configured keys.

---

# Resource: intersight_storage_storage_policy
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The security of drives can be enabled through this policy using remote keys from a KMIP server or Manually configured keys.
## Usage Example
### Resource Creation

```hcl
resource "intersight_storage_storage_policy" "tf_storage_policy" {
  name                     = "tf_storage_policy"
  use_jbod_for_vd_creation = true
  description              = "storage policy test"
  unused_disks_state       = "UnconfiguredGood"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  global_hot_spares = "3"
  m2_virtual_drive {
    enable      = false
    object_type = "storage.M2VirtualDriveConfig"
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `controller_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to controller attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number ranges. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `default_drive_mode`:(string) All unconfigured drives will move to the selected state on deployment. Newly inserted drives will move to the selected state. Select Unconfigured Good option to retain the existing configuration. Select JBOD to move the unconfigured drives to JBOD state. Select RAID0 to create a RAID0 virtual drive on each of the unconfigured drives. If JBOD is selected, unconfigured drives will move to JBOD state on host reboot. This setting is applicable only to selected set of controllers on FI attached servers.* `UnconfiguredGood` - Newly inserted drives or on reboot, drives will remain the same state.* `Jbod` - Newly inserted drives or on reboot, drives will automatically move to JBOD state if drive state was UnconfiguredGood.* `RAID0` - Newly inserted drives or on reboot, virtual drives will be created, respective drives will move to Online state. 
* `description`:(string) Description of the policy. 
* `direct_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to Direct attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number ranges. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `drive_group`:(Array) An array of relationships to storageDriveGroup resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `global_hot_spares`:(string) A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen. 
* `m2_virtual_drive`:(HashMap) - Virtual Drive configuration for M.2 RAID controller. 
This complex property has following sub-properties:
  + `controller_slot`:(string) Select the M.2 RAID controller slot on which the virtual drive is to be created. Select 'MSTOR-RAID-1' to create virtual drive on the M.2 RAID controller in the first slot or in the MSTOR-RAID slot, 'MSTOR-RAID-2' for second slot, 'MSTOR-RAID-1, MSTOR-RAID-2' for both slots or either slot.* `MSTOR-RAID-1` - Virtual drive  will be created on the M.2 RAID controller in the first slot.* `MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available.* `MSTOR-RAID-1,MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available. 
  + `enable`:(bool) If enabled, this will create a virtual drive on the M.2 RAID controller. 
  + `name`:(string) The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed. This field will be pre-populated with the default or user configured value which can be edited. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `raid0_drive`:(HashMap) - The list of disks where RAID0 virtual drives must be created on each individual drive. 
This complex property has following sub-properties:
  + `drive_slots`:(string) The set of drive slots where RAID0 virtual drives must be created. 
  + `drive_slots_list`:(string)(ReadOnly) The list of drive slots where RAID0 virtual drives must be created (comma separated). 
  + `enable`:(bool) If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `virtual_drive_policy`:(HashMap) - This defines the characteristics of a specific virtual drive. 
This complex property has following sub-properties:
    + `access_policy`:(string) Access policy that host has on this virtual drive.* `Default` - Use platform default access mode.* `ReadWrite` - Enables host to perform read-write on the VD.* `ReadOnly` - Host can only read from the VD.* `Blocked` - Host can neither read nor write to the VD. 
    + `drive_cache`:(string) Disk cache policy for the virtual drive.* `Default` - Use platform default drive cache mode.* `NoChange` - Drive cache policy is unchanged.* `Enable` - Enables IO caching on the drive.* `Disable` - Disables IO caching on the drive. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `read_policy`:(string) Read ahead mode to be used to read data from this virtual drive.* `Default` - Use platform default read ahead mode.* `ReadAhead` - Use read ahead mode for the policy.* `NoReadAhead` - Do not use read ahead mode for the policy. 
    + `strip_size`:(int) Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB.* `64` - Number of bytes in a strip is 64 Kibibytes.* `128` - Number of bytes in a strip is 128 Kibibytes.* `256` - Number of bytes in a strip is 256 Kibibytes.* `512` - Number of bytes in a strip is 512 Kibibytes.* `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte. 
    + `write_policy`:(string) Write mode to be used to write data to this virtual drive.* `Default` - Use platform default write mode.* `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache.* `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure.* `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. 
* `raid_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to RAID attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number ranges. Deprecated in favor of controllerAttachedNvmeSlots. 
* `secure_jbods`:(string) JBOD drives specified in this slot range will be encrypted. Allowed values are 'ALL', or a comma or hyphen separated number range. Sample format is ALL or 1, 3 or 4-6, 8. Setting the value to 'ALL' will encrypt all the unused UnconfigureGood/JBOD disks. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `unused_disks_state`:(string) State to which drives, not used in this policy, are to be moved. NoChange will not change the drive state. No Change must be selected if Default Drive State is set to JBOD or RAID0.* `NoChange` - Drive state will not be modified by Storage Policy.* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.* `Jbod` - JBOD state where the disks start showing up to Host OS. 
* `use_jbod_for_vd_creation`:(bool) Disks in JBOD State are used to create virtual drives. This setting must be disabled if Default Drive State is set to JBOD. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_storage_storage_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_storage_storage_policy.example 1234567890987654321abcde
``` 
