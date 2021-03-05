---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_container_runtime_policy"
description: |-
  A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
---

# Data Source: intersight_kubernetes_container_runtime_policy
A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_container_runtime_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `docker_bridge_network_cidr`:(string) The DNS Search Domain Name. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 