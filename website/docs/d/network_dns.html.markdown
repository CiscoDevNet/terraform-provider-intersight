---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_dns"
description: |-
        Dns represents the set of DNS settings configured on a Nexus endpoint, scoped to a specific VRF. It captures which domains are used for name resolution and which DNS servers are consulted when resolving names from that VRF context.
        #### Purpose
        Provide read-only visibility into per-VRF DNS configuration (default domain, additional search domains, and name server addresses) for troubleshooting, audit, and operational validation.
        #### Key Concepts
        - **VRF-scoped DNS**: DNS configuration is associated with a particular VRF (`vrfName`), reflecting segmented routing contexts.
        - **Search domain behavior**: `defaultDomain` and `additionalDomains` indicate the domains appended during hostname resolution.
        - **Resolver targets**: `nameServers` lists the DNS server addresses used for queries.
        - **Device association**: `registeredDevice` links the DNS configuration to the specific onboarded Nexus device.

---

# Data Source: intersight_network_dns
Dns represents the set of DNS settings configured on a Nexus endpoint, scoped to a specific VRF. It captures which domains are used for name resolution and which DNS servers are consulted when resolving names from that VRF context.
#### Purpose
Provide read-only visibility into per-VRF DNS configuration (default domain, additional search domains, and name server addresses) for troubleshooting, audit, and operational validation.
#### Key Concepts
- **VRF-scoped DNS**: DNS configuration is associated with a particular VRF (`vrfName`), reflecting segmented routing contexts.
- **Search domain behavior**: `defaultDomain` and `additionalDomains` indicate the domains appended during hostname resolution.
- **Resolver targets**: `nameServers` lists the DNS server addresses used for queries.
- **Device association**: `registeredDevice` links the DNS configuration to the specific onboarded Nexus device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_dns.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_domain`:(string) Default domain configured for VRF. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vrf_name`:(string) Name of the VRF configured for the DNS. 
 
