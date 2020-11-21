
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_pool"
sidebar_current: "docs-intersight-data-source-ippool-pool"
description: |-
Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.
---

# Data Source: intersight_ippool._pool
Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `size`:(int) Total number of identifiers in this pool. 
* `v4_assigned`:(int) Number of IPv4 addresses currently in use. 
* `v4_size`:(int) Number of IPv4 addresses in this pool. 
* `v6_assigned`:(int) Number of IPv6 addresses currently in use. 
* `v6_size`:(int) Number of IPv6 addresses in this pool. 
