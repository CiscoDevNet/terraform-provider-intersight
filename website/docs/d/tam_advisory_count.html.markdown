
---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_count"
sidebar_current: "docs-intersight-data-source-tam-advisory-count"
description: |-
Total number of advisories currently affecting a given Account.
---

# Data Source: intersight_tam._advisory_count
Total number of advisories currently affecting a given Account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_count`:(int) Total number of advisories affecting the account. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
