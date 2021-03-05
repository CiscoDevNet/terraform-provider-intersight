---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_infrastructure_provider"
description: |-
  Infrastructure backend for providing virtual machines from a target.
---

# Data Source: intersight_kubernetes_virtual_machine_infrastructure_provider
Infrastructure backend for providing virtual machines from a target.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_virtual_machine_infrastructure_provider.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description for the infrastructure provider. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of an infrastructure provider. 
 