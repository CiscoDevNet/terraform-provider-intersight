
---
layout: "intersight"
page_title: "Intersight: intersight_capability_catalog"
sidebar_current: "docs-intersight-data-source-capability-catalog"
description: |-
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
---

# Data Source: intersight_capability._catalog
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
