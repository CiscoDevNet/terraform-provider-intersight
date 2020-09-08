
---
layout: "intersight"
page_title: "Intersight: intersight_iam_domain_group"
sidebar_current: "docs-intersight-data-source-iam-domain-group"
description: |-
Intersight services are mapped to three different categories of services for scaling purpose.
Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with
a specific number of partitions. For each cloud environment these numbers will be different.
---

# Data Source: intersight_iam._domain_group
Intersight services are mapped to three different categories of services for scaling purpose.
Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with
a specific number of partitions. For each cloud environment these numbers will be different.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the domain-group. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `partition1`:(int) The partition number domain group related messages are produced for 'Partition1' category service topics. 
* `partition2`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics. 
* `partition3`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics. 
* `partition_key`:(string) Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as parition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. 
* `usage`:(int) The number of devices in the domain-group.Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. 
