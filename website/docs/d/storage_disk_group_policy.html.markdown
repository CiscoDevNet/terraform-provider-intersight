
---
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group_policy"
sidebar_current: "docs-intersight-data-source-storage-disk-group-policy"
description: |-
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
---

# Data Source: intersight_storage._disk_group_policy
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `raid_level`:(string) The supported RAID level for the disk group.* `Raid0` - RAID 0 Stripe Raid Level.* `Raid1` - RAID 1 Mirror Raid Level.* `Raid5` - RAID 5 Mirror Raid Level.* `Raid6` - RAID 6 Mirror Raid Level.* `Raid10` - RAID 10 Mirror Raid Level.* `Raid50` - RAID 50 Mirror Raid Level.* `Raid60` - RAID 60 Mirror Raid Level. 
* `use_jbods`:(bool) Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation. 
