---
subcategory: "search"
layout: "intersight"
page_title: "Intersight: intersight_search_tag_item"
description: |-
  The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
---

# Data Source: intersight_search_tag_item
The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_search_tag_item.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `nr_count`:(int) The number of times this tag key has been set across all resources. 
* `key`:(string) Key of the Tag from all the resources in Intersight. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 