---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_lun"
description: |-
        A HyperFlex iSCSI logical unit number (LUN) entity.
        Contains detailed information about the iSCSI LUN which includes the identity and capacity information, and the
        iSCSI target to which it is associated.

---

# Data Source: intersight_hyperflex_lun
A HyperFlex iSCSI logical unit number (LUN) entity.
Contains detailed information about the iSCSI LUN which includes the identity and capacity information, and the
iSCSI target to which it is associated.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_lun.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of iSCSI LUN. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ds_capacity_in_bytes`:(int) The datastore capacity in bytes. 
* `ds_name`:(string) The HyperFlex datastore name. 
* `ds_uuid`:(string) The HyperFlex datastore UUID. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `inventory_source`:(string) Source of the lun inventory.* `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable.* `ONLINE` - The source of the HyperFlex inventory is online.* `OFFLINE` - The source of the HyperFlex inventory is offline. 
* `is_encrypted`:(bool) Indicates if the datastore on which LUN resides is encrypted or un-encrypted. 
* `lun_id`:(string) The identity of HyperFlex iSCSI LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the HyperFlex iSCSI LUN. 
* `serial_no`:(string) Serial number of HyperFlex iSCSI LUN. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_capacity_in_bytes`:(int) Total capacity of iSCSI LUN in bytes. 
* `tuuid`:(string) HyperFlex iSCSI Target associated with the HyperFlex iSCSI LUN. 
* `used_capacity_in_bytes`:(int) Used capacity of iSCSI LUN in bytes. 
* `uuid`:(string) UUID of the HyperFlex iSCSI LUN. 
* `nr_version`:(int) Version of the HyperFlex iSCSI lun. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 
