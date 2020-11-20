
---
layout: "intersight"
page_title: "Intersight: intersight_networkconfig_policy"
sidebar_current: "docs-intersight-data-source-networkconfig-policy"
description: |-
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
---

# Data Source: intersight_networkconfig._policy
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `alternate_ipv4dns_server`:(string) IP address of the secondary DNS server. 
* `alternate_ipv6dns_server`:(string) IP address of the secondary DNS server. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `dynamic_dns_domain`:(string) The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request. 
* `enable_dynamic_dns`:(bool) If enabled, updates the resource records to the DNS from Cisco IMC. 
* `enable_ipv4dns_from_dhcp`:(bool) If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it. 
* `enable_ipv6`:(bool) If enabled, allows to configure IPv6 properties. 
* `enable_ipv6dns_from_dhcp`:(bool) If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `preferred_ipv4dns_server`:(string) IP address of the primary DNS server. 
* `preferred_ipv6dns_server`:(string) IP address of the primary DNS server. 
