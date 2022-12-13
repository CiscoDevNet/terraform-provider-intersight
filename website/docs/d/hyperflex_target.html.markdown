---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_target"
description: |-
        A HyperFlex iSCSI target entity.
        Contains detailed information about the iSCSI target which includes a list of linked iSCSI initiator group objects,
        list of iSCSI lun objects associated with the iSCSI target, authorization method and iSCSI qualified name.

---

# Data Source: intersight_hyperflex_target
A HyperFlex iSCSI target entity.
Contains detailed information about the iSCSI target which includes a list of linked iSCSI initiator group objects,
list of iSCSI lun objects associated with the iSCSI target, authorization method and iSCSI qualified name.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_target.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auth_method`:(string) Auth method of the target inventory.* `NOT_APPLICABLE` - Authorization method of the HyperFlex iSCSI target is not applicable.* `CHAP` - Authorization method of the HyperFlex iSCSI target is CHAP.* `NONE` - Authorization method of the HyperFlex iSCSI target is none. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the target. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `inventory_source`:(string) Source of the target inventory.* `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable.* `ONLINE` - The source of the HyperFlex inventory is online.* `OFFLINE` - The source of the HyperFlex inventory is offline. 
* `iqn`:(string) The iSCSI qualified name (IQN) of target. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the target in storage array. 
* `num_active_initiators`:(int) Number of active initiators in the initiator group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the HyperFlex iSCSI target. 
* `nr_version`:(int) Version of the Initiator Group. 
 
