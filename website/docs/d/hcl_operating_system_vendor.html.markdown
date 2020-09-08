
---
layout: "intersight"
page_title: "Intersight: intersight_hcl_operating_system_vendor"
sidebar_current: "docs-intersight-data-source-hcl-operating-system-vendor"
description: |-
Collection used to store operating system vendors details.
---

# Data Source: intersight_hcl._operating_system_vendor
Collection used to store operating system vendors details.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the vendor of the operating system. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
