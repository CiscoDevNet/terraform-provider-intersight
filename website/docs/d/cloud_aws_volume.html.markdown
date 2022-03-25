---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_aws_volume"
description: |-
        AWS volume object is represented here.It depicts configuration used to create volume and its identity.

---

# Data Source: intersight_cloud_aws_volume
AWS volume object is represented here.It depicts configuration used to create volume and its identity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_aws_volume.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `encryption_state`:(string) Encryption state of this volume.For example ENCRYPTED, NOT ENCRYPTED, etc.* `Automatic` - Volume encryption state is automatic.Cloud provider takes the decision of when to encrypt the data.* `Encrypted` - Volume data is encrypted. Can be decrypted only by authorized users.* `Not_Encrypted` - Volume data is not encrypted. 
* `identity`:(string) The internally generated identity of this VM. It aids in uniquely identifying the volume object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Disk size represented in bytes. 
* `source_image_id`:(string) Source Image Id used for the volume. 
* `state`:(string) The volume state.For example AVAILABLE, IN_USE, DELETED, etc.* `UnRecognized` - Volume is in unrecognized state.* `Pending` - Volume is  in pending state, due to errors encountered during volume creation.* `Bound` - Volume is in bound state.* `Available` - Volume is in available state.* `Error` - Volume is in error state.* `Released` - Volume is in released state.* `in-use` - Volume is in in-use state.* `Creating` - Volume is in creating state.* `Deleting` - Volume is in deleting state.* `Deleted` - Volume is in deleted state. 
* `uuid`:(string) UUID (Universally Unique IDentifier) is a 128-bit value used for unique identification of Volume. 
* `volume_create_time`:(string) Time when this volume is created. 
 
