---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_storage_policy"
description: |-
  A policy specifying HyperFlex cluster storage settings (optional).
---

# Data Source: intersight_hyperflex_cluster_storage_policy
A policy specifying HyperFlex cluster storage settings (optional).
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `disk_partition_cleanup`:(bool) If enabled, formats existing disk partitions (destroys all user data). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `vdi_optimization`:(bool) Enable or disable VDI optimization (hybrid HyperFlex systems only). 
