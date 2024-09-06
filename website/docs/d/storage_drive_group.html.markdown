---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_drive_group"
description: |-
        A reusable RAID drive group configuration that specifies a pool of drives and a set of virtual drives that are to be created using this pool of drives.

---

# Data Source: intersight_storage_drive_group
A reusable RAID drive group configuration that specifies a pool of drives and a set of virtual drives that are to be created using this pool of drives.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_drive_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the drive group. The name can be between 1 and 64 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. 
* `raid_level`:(string) The supported RAID level for the disk group.* `Raid0` - RAID 0 Stripe Raid Level.* `Raid1` - RAID 1 Mirror Raid Level.* `Raid5` - RAID 5 Mirror Raid Level.* `Raid6` - RAID 6 Mirror Raid Level.* `Raid10` - RAID 10 Mirror Raid Level.* `Raid50` - RAID 50 Mirror Raid Level.* `Raid60` - RAID 60 Mirror Raid Level. 
* `secure_drive_group`:(bool) Enables/disables the drive security on all the drives used in this policy. This flag just enables the drive security and only after Remote/Manual key setting configured, the actual security will be applied. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(int) Type of drive selection to be used for this drive group.* `0` - Drives are selected manually by the user.* `1` - Drives are selected automatically based on the RAID and virtual drive configuration. 
 
