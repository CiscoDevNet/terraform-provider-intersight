---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_storage_policy"
description: |-
        The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.

---

# Data Source: intersight_storage_storage_policy
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_storage_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `global_hot_spares`:(string) A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `unused_disks_state`:(string) State to which disks, not used in this policy, are to be moved. NoChange will not change the drive state.* `NoChange` - Drive state will not be modified by Storage Policy.* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.* `Jbod` - JBOD state where the disks start showing up to Host OS. 
* `use_jbod_for_vd_creation`:(bool) Disks in JBOD State are used to create virtual drives. 
 
