---
subcategory: "networkconfig"
layout: "intersight"
page_title: "Intersight: intersight_networkconfig_policy"
description: |-
  Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
---

# Data Source: intersight_networkconfig_policy
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_networkconfig_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alternate_ipv4dns_server`:(string) IP address of the secondary DNS server. 
* `alternate_ipv6dns_server`:(string) IP address of the secondary DNS server. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dynamic_dns_domain`:(string) The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request. 
* `enable_dynamic_dns`:(bool) If enabled, updates the resource records to the DNS from Cisco IMC. 
* `enable_ipv4dns_from_dhcp`:(bool) If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it. 
* `enable_ipv6`:(bool) If enabled, allows to configure IPv6 properties. 
* `enable_ipv6dns_from_dhcp`:(bool) If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `preferred_ipv4dns_server`:(string) IP address of the primary DNS server. 
* `preferred_ipv6dns_server`:(string) IP address of the primary DNS server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 