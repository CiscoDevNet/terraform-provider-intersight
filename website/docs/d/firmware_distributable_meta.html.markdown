---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_distributable_meta"
description: |-
        Meta information for various firmware images stored in the database. Gives information on the particular category for a product.

---

# Data Source: intersight_firmware_distributable_meta
Meta information for various firmware images stored in the database. Gives information on the particular category for a product.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_distributable_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bucket_name`:(string) The S3 bucket name where the images are present, if source is external cloud store. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_type`:(string) The type of distributable image, example huu, scu, driver, os.* `Distributable` - Stores firmware host utility images and fabric images.* `DriverDistributable` - Stores driver distributable images.* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.* `OperatingSystemFile` - Stores operating system iso images.* `HyperflexDistributable` - It stores HyperFlex images.* `HciDistributable` - It stores HCI images, such as bootstrap iso images. 
* `from_version`:(string) The version from which user can download images from amazon store, if source is external cloud store. 
* `mdfid`:(string) The mdfid of the image provided by cisco.com. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `software_type_id`:(string) The software type id provided by cisco.com. 
* `nr_source`:(string) The image can be downloaded from cisco.com or external cloud store.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `to_version`:(string) The version till which user can download images from amazon store, if source is external cloud store. 
 
