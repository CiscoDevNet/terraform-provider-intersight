---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_aws_subnet"
description: |-
  Subnet object in AWS inventory. It is a range of IP addresses in a VPC that can be used to isolate
different EC2 resources from each other or from the Internet.
---

# Data Source: intersight_cloud_aws_subnet
Subnet object in AWS inventory. It is a range of IP addresses in a VPC that can be used to isolate
different EC2 resources from each other or from the Internet.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_aws_subnet.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_assign_private_ip_v6`:(bool) If true, Ipv4 address is assigned automatically. 
* `auto_assign_public_ip_v4`:(bool) If true, Ipv4 address is assigned automatically. 
* `availability_zone_name`:(string) The Availability Zone name of the subnet. 
* `cidr`:(string) CIDR scheme for defining an IP block. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The internally generated identity of network. This entity cannot manipulated by users. It aids in uniquely identifying the network object. For VMware, this is MOR (managed object reference). 
* `ipv4_cidr`:(string) The IPv4 CIDR block assigned to the subnet.. 
* `ipv6_cidr`:(string) The IPv6 CIDR block assigned to the subnet.. 
* `is_default`:(bool) If true, indicates that this is default subnet. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the portgroup. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the subnet (pending | available). 
 