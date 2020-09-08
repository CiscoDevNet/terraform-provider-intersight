
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The catalog name. There will be one for system and one for each user account. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
