
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_pool_member"
sidebar_current: "docs-intersight-data-source-ippool-pool-member"
description: |-
PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
---

# Data Source: intersight_ippool._pool_member
PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address of this pool member. 
* `ip_v6_address`:(string) IPv6 Address of this pool member. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
