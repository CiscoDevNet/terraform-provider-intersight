
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper"
sidebar_current: "docs-intersight-data-source-softwarerepository-category-mapper"
description: |-
Maps a Cisco software repository image category identifier to its applicable hardware models.
---

# Data Source: intersight_softwarerepository._category_mapper
Maps a Cisco software repository image category identifier to its applicable hardware models.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) The category of the model series. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `file_type`:(string) The type of distributable image, example huu, scu, driver, os.* `Distributable` - Stores firmware host utility images and fabric images.* `DriverDistributable` - Stores driver distributable images.* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.* `OperatingSystemFile` - Stores operating system iso images.* `HyperflexDistributable` - It stores HyperFlex images. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `regex_pattern`:(string) The regex that all images of this category follow. 
* `source`:(string) The image can be downloaded from cisco.com or external cloud store.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `sw_id`:(string) The software type id provided by cisco.com. 
* `version`:(string) The version from which user can download images from amazon store, if source is external cloud store. 
