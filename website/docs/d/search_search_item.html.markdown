
---
layout: "intersight"
page_title: "Intersight: intersight_search_search_item"
sidebar_current: "docs-intersight-data-source-search-search-item"
description: |-
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.
---

# Data Source: intersight_search._search_item
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
