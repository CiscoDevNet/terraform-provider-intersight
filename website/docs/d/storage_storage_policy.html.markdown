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
To access the ith object of the results obtained, use `data.intersight_storage_storage_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `retain_policy_virtual_drives`:(bool) Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy. 
* `unused_disks_state`:(string) Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved.* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.* `Jbod` - JBOD state where the disks start showing up to host os. 
 