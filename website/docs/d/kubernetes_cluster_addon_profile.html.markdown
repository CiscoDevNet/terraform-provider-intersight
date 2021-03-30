---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_cluster_addon_profile"
description: |-
  A profile that describes a collection of Addons for a particular cluster.
---

# Data Source: intersight_kubernetes_cluster_addon_profile
A profile that describes a collection of Addons for a particular cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_cluster_addon_profile.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the cluster addon profile. 
 