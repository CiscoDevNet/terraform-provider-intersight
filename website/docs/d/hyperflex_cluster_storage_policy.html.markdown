
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `disk_partition_cleanup`:(bool) If enabled, formats existing disk partitions (destroys all user data). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `vdi_optimization`:(bool) Enable or disable VDI optimization (hybrid HyperFlex systems only). 
