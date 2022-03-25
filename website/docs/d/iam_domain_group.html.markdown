---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_domain_group"
description: |-
        Intersight services are mapped to three different categories of services for scaling purpose.
        Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with
        a specific number of partitions. For each cloud environment these numbers will be different.

---

# Data Source: intersight_iam_domain_group
Intersight services are mapped to three different categories of services for scaling purpose.
Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with
a specific number of partitions. For each cloud environment these numbers will be different.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_domain_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the domain-group. 
* `partition1`:(int) The partition number domain group related messages are produced for 'Partition1' category service topics. 
* `partition2`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics. 
* `partition3`:(int) In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics. 
* `partition_key`:(string) Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as partition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage`:(int) The number of devices in the domain-group.Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. 
 
