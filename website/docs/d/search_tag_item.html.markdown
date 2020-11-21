
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `count`:(int) The number of times this tag key has been set across all resources. 
* `key`:(string) Key of the Tag from all the resources in Intersight. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
