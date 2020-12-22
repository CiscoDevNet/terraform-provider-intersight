---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_ip_lease"
description: |-
  IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
---

# Data Source: intersight_ippool_ip_lease
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address given as a lease to an external entity like server profiles. 
* `ip_v6_address`:(string) IPv6 Address given as a lease to an external entity like server profiles. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
