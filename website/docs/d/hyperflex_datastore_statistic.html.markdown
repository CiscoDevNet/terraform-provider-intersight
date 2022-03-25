---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_datastore_statistic"
description: |-
        Datastore Statistic describing more detailed information about the Data Protection Peer.

---

# Data Source: intersight_hyperflex_datastore_statistic
Datastore Statistic describing more detailed information about the Data Protection Peer.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_datastore_statistic.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accessibility_summary`:(string) HyperFlex datastore accessibility summary.* `ACCESSIBLE` - The HyperFlex Accessibility Summary is Accessible.* `NOT_ACCESSIBLE` - The HyperFlex Accessibility Summary is Not Accessible.* `PARTIALLY_ACCESSIBLE` - The HyperFlex Accessibility Summary is Partially Accessible. 
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `creation_time`:(string) Timestamp the datastore object was created. 
* `datastore_kind`:(string) HyperFlex Datastore Kind.* `UNKNOWN` - HyperFlex datastore kind is unknown.* `USER_CREATED` - HyperFlex datastore kind is user created.* `INTERNAL` - HyperFlex datastore kind is internal. 
* `datastore_status`:(string) HyperFlex datastore status.* `NORMAL` - The HyperFlex datastore status is normal.* `ALERT` - The HyperFlex datastore status is alert.* `FAILED` - The HyperFlex datastore status is failed. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_capacity_in_bytes`:(int) Free capacity of the datastore in bytes. 
* `is_encrypted`:(bool) Indicates if the datastore is encrypted or un-encrypted. 
* `last_access_time`:(string) Timestamp the datastore object was last accessed. 
* `last_modified_time`:(string) Timestamp the datastore object was last modified. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mount_summary`:(string) HyperFlex datastore mount summary.* `MOUNTED` - The HyperFlex mount summary is mounted.* `UNMOUNTED` - The HyperFlex mount summary is unmounted.* `MOUNT_FAILURE` - The HyperFlex mount summary is mount failure.* `UNMOUNT_FAILURE` - The HyperFlex mount summary is unmount failure. 
* `parent_uuid`:(string) UUID of the parent datastore object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_capacity_in_bytes`:(int) Total capacity of the datastore object. 
* `un_compressed_used_bytes`:(int) Number of uncompressed used bytes in the datastore. 
* `unshared_used_bytes`:(int) Unshared used capacity of the datastore in bytes. 
* `uuid`:(string) UUID for the datastore object. 
 
