---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_initiator_group"
description: |-
        A HyperFlex iSCSI initiator group entity.
        Contains detailed information about the initaitor group which includes a list of iSCSI initiators and iSCSI target objects.

---

# Data Source: intersight_hyperflex_initiator_group
A HyperFlex iSCSI initiator group entity.
Contains detailed information about the initaitor group which includes a list of iSCSI initiators and iSCSI target objects.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_initiator_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the host group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `initiator_count`:(int) Count of initiators in the iSCSI initiator group. 
* `inventory_source`:(string) The source of the iSCSI initiator group inventory.* `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable.* `ONLINE` - The source of the HyperFlex inventory is online.* `OFFLINE` - The source of the HyperFlex inventory is offline. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host group in storage array. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the HyperFlex iSCSI initiator group. 
* `nr_version`:(int) Version of the iSCSI initiator group. 
 
