---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_drive_group"
description: |-
        A reusable RAID drive group configuration that specifies a pool of drives and a set of virtual drives that are to be created using this pool of drives.

---

# Resource: intersight_storage_drive_group
A reusable RAID drive group configuration that specifies a pool of drives and a set of virtual drives that are to be created using this pool of drives.
## Usage Example
### Resource Creation

```hcl
resource "intersight_storage_drive_group" "tf_drive_gp" {
  name       = "tf_drive_gp"
  raid_level = "Raid0"
  manual_drive_group {
    object_type = "storage.ManualDriveGroup"
    span_groups {
      slots       = "2"
      object_type = "storage.SpanDrives"
    }
  }
  virtual_drives {
    name                = "tf_drive_gp-vd"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    object_type         = "storage.VirtualDriveConfiguration"
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
  virtual_drives {
    name                = "drive_gp-vd-01"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    object_type         = "storage.VirtualDriveConfiguration"
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
  storage_policy {
    moid        = var.tf_storage_policy
    object_type = "storage.StoragePolicy"
  }
}

variable "tf_storage_policy" {
  type        = string
  description = "value for tf_storage_policy"
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
* `automatic_drive_group`:(HashMap) - This drive group is created using automatic drive selection. 
This complex property has following sub-properties:
  + `drive_type`:(string) Type of drive that should be used for this RAID group.* `Any` - Any type of drive can be used for virtual drive creation.* `HDD` - Hard disk drives should be used for virtual drive creation.* `SSD` - Solid state drives should be used for virtual drive creation. 
  + `drives_per_span`:(int) Number of drives within this span group. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk. RAID1 and RAID10 requires at least 2 and in multiples of . RAID5 and RAID50 require at least 3 disks in a span group. RAID6 and RAID60 require atleast 4 disks in a span. 
  + `minimum_drive_size`:(int) Minimum size of the drive to be used for creating this RAID group. 
  + `num_dedicated_hot_spares`:(string) Number of dedicated hot spare disks for this RAID group. Allowed value is a comma or hyphen separated number range. 
  + `number_of_spans`:(int) Number of span groups to be created for this RAID group. Non-nested RAID levels have a single span. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `use_remaining_drives`:(bool) This flag enables the drive group to use all the remaining drives on the server. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `manual_drive_group`:(HashMap) - This drive group is created by specifying the drive slots to be used. 
This complex property has following sub-properties:
  + `dedicated_hot_spares`:(string) A collection of drives to be used as hot spares for this Drive Group. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `span_groups`:(Array)
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `slots`:(string) Collection of local disks that are part of this span group. Allowed value is a comma or hyphen separated number range. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk, RAID1 and RAID10 requires at least 2 and in multiples of 2, RAID5 RAID50 RAID6 and RAID60 require at least 3 disks in a span group. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the drive group. The name can be between 1 and 64 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. 
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
* `raid_level`:(string) The supported RAID level for the disk group.* `Raid0` - RAID 0 Stripe Raid Level.* `Raid1` - RAID 1 Mirror Raid Level.* `Raid5` - RAID 5 Mirror Raid Level.* `Raid6` - RAID 6 Mirror Raid Level.* `Raid10` - RAID 10 Mirror Raid Level.* `Raid50` - RAID 50 Mirror Raid Level.* `Raid60` - RAID 60 Mirror Raid Level. 
* `secure_drive_group`:(bool) Enables/disables the drive encryption on all the drives used in this policy. This flag just enables the drive security and only after remote key setting configured, the actual encryption will be done. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_policy`:(HashMap) - A reference to a storageStoragePolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `type`:(int)(ReadOnly) Type of drive selection to be used for this drive group.* `0` - Drives are selected manually by the user.* `1` - Drives are selected automatically based on the RAID and virtual drive configuration. 
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
* `virtual_drives`:(Array)
This complex property has following sub-properties:
  + `boot_drive`:(bool) This flag enables this virtual drive to be used as a boot drive. 
  + `expand_to_available`:(bool) This flag enables the virtual drive to use all the space available in the disk group. When this flag is enabled, the size property is ignored. 
  + `name`:(string) The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `size`:(int) Virtual drive size in MebiBytes. Size is mandatory field except when the Expand to Available option is enabled. 
  + `virtual_drive_policy`:(HashMap) - This defines the characteristics of a specific virtual drive. 
This complex property has following sub-properties:
    + `access_policy`:(string) Access policy that host has on this virtual drive.* `Default` - Use platform default access mode.* `ReadWrite` - Enables host to perform read-write on the VD.* `ReadOnly` - Host can only read from the VD.* `Blocked` - Host can neither read nor write to the VD. 
    + `drive_cache`:(string) Disk cache policy for the virtual drive.* `Default` - Use platform default drive cache mode.* `NoChange` - Drive cache policy is unchanged.* `Enable` - Enables IO caching on the drive.* `Disable` - Disables IO caching on the drive. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `read_policy`:(string) Read ahead mode to be used to read data from this virtual drive.* `Default` - Use platform default read ahead mode.* `ReadAhead` - Use read ahead mode for the policy.* `NoReadAhead` - Do not use read ahead mode for the policy. 
    + `strip_size`:(int) Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB.* `64` - Number of bytes in a strip is 64 Kibibytes.* `128` - Number of bytes in a strip is 128 Kibibytes.* `256` - Number of bytes in a strip is 256 Kibibytes.* `512` - Number of bytes in a strip is 512 Kibibytes.* `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte. 
    + `write_policy`:(string) Write mode to be used to write data to this virtual drive.* `Default` - Use platform default write mode.* `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache.* `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure.* `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. 


## Import
`intersight_storage_drive_group` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_storage_drive_group.example 1234567890987654321abcde
``` 
