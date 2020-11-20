
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the domain-group. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `partition1`:(int) The partition number domain group related messages are produced for 'Partition1' category service topics. 
* `partition2`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics. 
* `partition3`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics. 
* `partition_key`:(string) Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as parition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. 
* `usage`:(int) The number of devices in the domain-group.Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. 
