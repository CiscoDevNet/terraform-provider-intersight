
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS distribution such as ESXi, CentOS. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
