
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_ip_lease"
sidebar_current: "docs-intersight-data-source-ippool-ip-lease"
description: |-
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
---

# Data Source: intersight_ippool._ip_lease
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address given as a lease to an external entity like server profiles. 
* `ip_v6_address`:(string) IPv6 Address given as a lease to an external entity like server profiles. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
