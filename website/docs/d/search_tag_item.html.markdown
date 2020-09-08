
---
layout: "intersight"
page_title: "Intersight: intersight_search_tag_item"
sidebar_current: "docs-intersight-data-source-search-tag-item"
description: |-
The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
---

# Data Source: intersight_search._tag_item
The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `count`:(int) The number of times this tag key has been set across all resources. 
* `key`:(string) Key of the Tag from all the resources in Intersight. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
