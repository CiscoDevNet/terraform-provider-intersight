
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_universe"
sidebar_current: "docs-intersight-data-source-ippool-universe"
description: |-
Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.
---

# Data Source: intersight_ippool._universe
Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
