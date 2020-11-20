
---
layout: "intersight"
page_title: "Intersight: intersight_os_catalog"
sidebar_current: "docs-intersight-data-source-os-catalog"
description: |-
A catalog of operating systems related objects such as configuration files and post install scripts. Each user account will have a local OS catalog where account users can store their private configuration files and post install scripts.
Cisco provides validated answer files and post install scripts to Intersight users via shared catalogs. Intersight users will be able to read, use these files and scripts in OS installation within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
---

# Data Source: intersight_os._catalog
A catalog of operating systems related objects such as configuration files and post install scripts. Each user account will have a local OS catalog where account users can store their private configuration files and post install scripts.
Cisco provides validated answer files and post install scripts to Intersight users via shared catalogs. Intersight users will be able to read, use these files and scripts in OS installation within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The catalog name. There will be one for system and one for each user account. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
