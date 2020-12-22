---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper"
description: |-
  Maps a Cisco software repository image category identifier to its applicable hardware models.
---

# Data Source: intersight_softwarerepository_category_mapper
Maps a Cisco software repository image category identifier to its applicable hardware models.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) The category of the model series. 
* `file_type`:(string) The type of distributable image, example huu, scu, driver, os.* `Distributable` - Stores firmware host utility images and fabric images.* `DriverDistributable` - Stores driver distributable images.* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.* `OperatingSystemFile` - Stores operating system iso images.* `HyperflexDistributable` - It stores HyperFlex images. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `regex_pattern`:(string) The regex that all images of this category follow. 
* `nr_source`:(string) The image can be downloaded from cisco.com or external cloud store.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `sw_id`:(string) The software type id provided by cisco.com. 
* `nr_version`:(string) The version from which user can download images from amazon store, if source is external cloud store. 
