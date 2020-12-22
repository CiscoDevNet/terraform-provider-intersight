---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_storage_policy"
description: |-
  The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
---

# Resource: intersight_storage_storage_policy
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `disk_group_policies`:(Array) An array of relationships to storageDiskGroupPolicy resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `global_hot_spares`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `slot_number`:(int) The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `retain_policy_virtual_drives`:(bool) Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `unused_disks_state`:(string) Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved.* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.* `Jbod` - JBOD state where the disks start showing up to host os. 
* `virtual_drives`:(Array)
This complex property has following sub-properties:
  + `access_policy`:(string) Access policy that host has on this virtual drive.* `Default` - Use platform default access mode.* `ReadWrite` - Enables host to perform read-write on the VD.* `ReadOnly` - Host can only read from the VD.* `Blocked` - Host can neither read nor write to the VD. 
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `boot_drive`:(bool) The flag enables the use of this virtual drive as a boot drive. 
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `disk_group_name`:(string)(Computed) Disk group policy that has the disk group in which this virtual drive needs to be created. 
  + `disk_group_policy`:(string) Disk group policy that has the disk group in which this virtual drive needs to be created. 
  + `drive_cache`:(string) Drive Cache property expect disk cache policy.* `Default` - Use platform default drive cache mode.* `NoChange` - Drive cache policy is unchanged.* `Enable` - Enables IO caching on the drive.* `Disable` - Disables IO caching on the drive. 
  + `expand_to_available`:(bool) The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored. 
  + `io_policy`:(string) Desired IO mode - direct IO or cached IO.* `Default` - Use platform default IO mode.* `Direct` - Use direct IO for writing directly into the disk.* `Cached` - Use cached IO for writing into cache and then to disk. 
  + `name`:(string) The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `read_policy`:(string) Read ahead mode to be used to read data from this virtual drive.* `Default` - Use platform default read ahead mode.* `ReadAhead` - Use read ahead mode for the policy.* `NoReadAhead` - Do not use read ahead mode for the policy. 
  + `size`:(int) Virtual drive size in MB. Size is mandatory field unless the 'Expand to Available' option is enabled. 
  + `strip_size`:(string) The strip size is the portion of a stripe that resides on a single drive in the drive group. The stripe consists of the data segments that the RAID controller writes across multiple drives, not including parity drives.* `Default` - Use platform default strip size for a virtual drive.* `32k` - Enables a strip size of 32k for a virtual drive.* `64k` - Enables a strip size of 64k for a virtual drive.* `128k` - Enables a strip size of 128k for a virtual drive.* `256k` - Enables a strip size of 256k for a virtual drive.* `512k` - Enables a strip size of 512k for a virtual drive.* `1024k` - Enables a strip size of 1024k for a virtual drive. 
  + `vdid`:(string)(Computed) Unique Id of the Virtual Drive to be used to identify Virtual Drive when name is empty. 
  + `write_policy`:(string) Write mode to be used to write data to this virtual drive.* `Default` - Use platform default write mode.* `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache.* `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure.* `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. 


## Import
`intersight_storage_storage_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_storage_storage_policy.example 1234567890987654321abcde
```