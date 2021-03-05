---
subcategory: "search"
layout: "intersight"
page_title: "Intersight: intersight_search_search_item"
description: |-
  The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.
---

# Data Source: intersight_search_search_item
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_search_search_item.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 