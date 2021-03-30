---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_infra_config_policy"
description: |-
  A policy specifying compute, storage and network infrastructure configuration for a Virtual Machine.
---

# Data Source: intersight_kubernetes_virtual_machine_infra_config_policy
A policy specifying compute, storage and network infrastructure configuration for a Virtual Machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_virtual_machine_infra_config_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 