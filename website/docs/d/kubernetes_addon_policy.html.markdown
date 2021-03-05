---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon_policy"
description: |-
  A policy that defines which AddonDefinitions to use.
---

# Data Source: intersight_kubernetes_addon_policy
A policy that defines which AddonDefinitions to use.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_addon_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `system_managed`:(bool) To determine if Addon Policy is automatically managed by the system. 
 