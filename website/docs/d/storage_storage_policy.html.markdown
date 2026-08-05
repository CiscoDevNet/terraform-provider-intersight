---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_storage_policy"
description: |-
        The StoragePolicy object models reusable storage configurations applicable to multiple servers. It enables RAID group creation and offers flexibility in drive state management.
        #### Purpose
        The StoragePolicy object provides a unified framework for configuring and managing storage solutions. It enables the creation of RAID groups using predefined disk group policies and offers flexible drive state management, allowing unused disks to be moved to JBOD or Unconfigured Good states while supporting hybrid slot configurations for transitions between direct and controller modes.
        #### Key Concepts
        - **Drive Security:** Enhances data protection through remote or manual key settings for drive encryption.
        - **Global Hot Spares:** Configures disks as hot spares across RAID groups to improve storage reliability.
        - **Nvme Drive Management:** Includes settings for NVMe drive attachment modes, supporting both direct and controller attached configurations.

---

# Data Source: intersight_storage_storage_policy
The StoragePolicy object models reusable storage configurations applicable to multiple servers. It enables RAID group creation and offers flexibility in drive state management.
#### Purpose
The StoragePolicy object provides a unified framework for configuring and managing storage solutions. It enables the creation of RAID groups using predefined disk group policies and offers flexible drive state management, allowing unused disks to be moved to JBOD or Unconfigured Good states while supporting hybrid slot configurations for transitions between direct and controller modes.
#### Key Concepts
- **Drive Security:** Enhances data protection through remote or manual key settings for drive encryption.
- **Global Hot Spares:** Configures disks as hot spares across RAID groups to improve storage reliability.
- **Nvme Drive Management:** Includes settings for NVMe drive attachment modes, supporting both direct and controller attached configurations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_storage_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to controller attached mode. Allowed slots are 1-9, 21-24, 101-104. Allowed value is a comma or hyphen separated number ranges. 
* `create_time`:(string) The time when this managed object was created. 
* `default_drive_mode`:(string) All unconfigured drives will move to the selected state on deployment. Newly inserted drives will move to the selected state. Select Unconfigured Good option to retain the existing configuration. Select JBOD to move the unconfigured drives to JBOD state. Select RAID0 to create a RAID0 virtual drive on each of the unconfigured drives. If JBOD is selected, unconfigured drives will move to JBOD state on host reboot. This setting is applicable only to selected set of controllers on FI attached servers.* `UnconfiguredGood` - Newly inserted drives or on reboot, drives will remain the same state.* `Jbod` - Newly inserted drives or on reboot, drives will automatically move to JBOD state if drive state was UnconfiguredGood.* `RAID0` - Newly inserted drives or on reboot, virtual drives will be created, respective drives will move to Online state. 
* `description`:(string) Description of the policy. 
* `direct_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to Direct attached mode. Allowed slots are 1-9, 21-24, 101-104. Allowed value is a comma or hyphen separated number ranges. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `global_hot_spares`:(string) A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `raid_attached_nvme_slots`:(string) Only U.3 NVMe drives need to be specified, entered slots will be moved to RAID attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number ranges. Deprecated in favor of controllerAttachedNvmeSlots. 
* `secure_jbods`:(string) JBOD drives specified in this slot range will be encrypted. Allowed values are 'ALL', or a comma or hyphen separated number range. Sample format is ALL or 1, 3 or 4-6, 8. Setting the value to 'ALL' will encrypt all the unused UnconfigureGood/JBOD disks. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `unused_disks_state`:(string) State to which drives, not used in this policy, are to be moved. NoChange will not change the drive state. No Change must be selected if Default Drive State is set to JBOD or RAID0.* `NoChange` - Drive state will not be modified by Storage Policy.* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.* `Jbod` - JBOD state where the disks start showing up to Host OS. 
* `use_jbod_for_vd_creation`:(bool) Disks in JBOD State are used to create virtual drives. This setting must be disabled if Default Drive State is set to JBOD. 
 
