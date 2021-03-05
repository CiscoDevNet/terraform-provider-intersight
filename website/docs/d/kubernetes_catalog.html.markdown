---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_catalog"
description: |-
  A catalog to hold the Kubernetes related items such as versions and addons for Intersight Kubernetes Services.
---

# Data Source: intersight_kubernetes_catalog
A catalog to hold the Kubernetes related items such as versions and addons for Intersight Kubernetes Services.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the catalog. The names are populated and predefined during MO creation. 
 