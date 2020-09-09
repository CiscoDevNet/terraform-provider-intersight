
---
layout: "intersight"
page_title: "Intersight: intersight_os_distribution"
sidebar_current: "docs-intersight-data-source-os-distribution"
description: |-
Intersight has the distribution details for all the Intersight supported OS
distributions. There will be a Distribution object for each supported OS.
---

# Data Source: intersight_os._distribution
Intersight has the distribution details for all the Intersight supported OS
distributions. There will be a Distribution object for each supported OS.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS distribution such as ESXi, CentOS. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
