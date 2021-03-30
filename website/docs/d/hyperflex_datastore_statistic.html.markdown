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
To access the ith object of the results obtained, use `data.intersight_hyperflex_datastore_statistic.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accessibility_summary`:(string) HyperFlex datastore accessibility summary.* `ACCESSIBLE` - The HyperFlex Accessibility Summary is Accessible.* `NOT_ACCESSIBLE` - The HyperFlex Accessibility Summary is Not Accessible.* `PARTIALLY_ACCESSIBLE` - The HyperFlex Accessibility Summary is Partially Accessible. 
* `creation_time`:(string) Timestamp the datastore object was created. 
* `datastore_status`:(string) HyperFlex datastore status.* `NORMAL` - The HyperFlex datastore status is normal.* `ALERT` - The HyperFlex datastore status is alert.* `FAILED` - The HyperFlex datastore status is failed. 
* `free_capacity_in_bytes`:(int) Free capacity of the datastore in bytes. 
* `last_modified_time`:(string) Timestamp the datastore object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mount_summary`:(string) HyperFlex datastore mount summary.* `MOUNTED` - The HyperFlex mount summary is mounted.* `UNMOUNTED` - The HyperFlex mount summary is unmounted.* `MOUNT_FAILURE` - The HyperFlex mount summary is mount failure.* `UNMOUNT_FAILURE` - The HyperFlex mount summary is unmount failure. 
* `parent_uuid`:(string) UUID of the parent datastore object. 
* `total_capacity_in_bytes`:(int) Total capacity of the datastore object. 
* `un_compressed_used_bytes`:(int) Number of uncompressed used bytes in the datastore. 
* `unshared_used_bytes`:(int) Unshared used capacity of the datastore in bytes. 
* `uuid`:(string) UUID for the datastore object. 
 