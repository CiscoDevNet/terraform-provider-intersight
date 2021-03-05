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
To access the ith object of the results obtained, use `data.intersight_firmware_distributable_meta.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bucket_name`:(string) The S3 bucket name where the images are present, if source is external cloud store. 
* `file_type`:(string) The type of distributable image, example huu, scu, driver, os.* `Distributable` - Stores firmware host utility images and fabric images.* `DriverDistributable` - Stores driver distributable images.* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.* `OperatingSystemFile` - Stores operating system iso images.* `HyperflexDistributable` - It stores HyperFlex images. 
* `from_version`:(string) The version from which user can download images from amazon store, if source is external cloud store. 
* `mdfid`:(string) The mdfid of the image provided by cisco.com. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `software_type_id`:(string) The software type id provided by cisco.com. 
* `nr_source`:(string) The image can be downloaded from cisco.com or external cloud store.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `to_version`:(string) The version till which user can download images from amazon store, if source is external cloud store. 
 