
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_storage_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-cluster-storage-policy"
description: |-
A policy specifying HyperFlex cluster storage settings (optional).
---

# Data Source: intersight_hyperflex._cluster_storage_policy
A policy specifying HyperFlex cluster storage settings (optional).
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `disk_partition_cleanup`:(bool) If enabled, formats existing disk partitions (destroys all user data). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `vdi_optimization`:(bool) Enable or disable VDI optimization (hybrid HyperFlex systems only). 
