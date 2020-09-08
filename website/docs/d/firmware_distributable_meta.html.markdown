
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_distributable_meta"
sidebar_current: "docs-intersight-data-source-firmware-distributable-meta"
description: |-
Meta information for various firmware images stored in the database. Gives information on the particular category for a product.
---

# Data Source: intersight_firmware._distributable_meta
Meta information for various firmware images stored in the database. Gives information on the particular category for a product.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bucket_name`:(string) The S3 bucket name where the images are present, if source is external cloud store. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `file_type`:(string) The type of distributable image, example huu, scu, driver, os.* `Distributable` - Stores firmware host utility images and fabric images.* `DriverDistributable` - Stores driver distributable images.* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.* `OperatingSystemFile` - Stores operating system iso images.* `HyperflexDistributable` - It stores HyperFlex images. 
* `from_version`:(string) The version from which user can download images from amazon store, if source is external cloud store. 
* `mdfid`:(string) The mdfid of the image provided by cisco.com. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `software_type_id`:(string) The software type id provided by cisco.com. 
* `source`:(string) The image can be downloaded from cisco.com or external cloud store.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `to_version`:(string) The version till which user can download images from amazon store, if source is external cloud store. 
