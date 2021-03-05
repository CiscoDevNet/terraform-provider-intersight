---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_version"
description: |-
  A policy capturing software versions for a Kubernetes cluster.
---

# Data Source: intersight_kubernetes_version
A policy capturing software versions for a Kubernetes cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_version.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `kubernetes_version`:(string) Desired Kubernetes version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of this IKS kubernetes version. 
 