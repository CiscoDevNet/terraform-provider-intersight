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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_ip_lease.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `has_duplicate`:(bool) HasDuplicate represents if there are other pools in which this id exists. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address given as a lease to an external entity like server profiles. 
* `ip_v6_address`:(string) IPv6 Address given as a lease to an external entity like server profiles. 
* `migrate`:(bool) The migration capability is applicable only for dynamic lease requests and it works in conjunction with  preferred ID. If there is an existing dynamic or static lease that matches the preferred ID, that existing  lease will be migrated to the current pool. That means the existing lease will be deleted and a new lease  will be created in the pool. If there is a reservation exists that matches with preferred ID, that  reservation will be kept as is and next available ID from the pool will be leased. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `preferred_ip_v4_address`:(string) The preferred IPv4 address can be specified only for dynamic lease requests. Intersight will make its best effort to allocate that IPv4 address if it is available in the pool. If the specified preferred IPv4 address is not in the range of the pool or if it is already leased or reserved, then the next available IPv4 address from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease request will fail if it specifies the preferred IPv4 address property. When the preferred IPv4 address property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be replaced by the new lease. Migration also supported only for dynamic lease requests. 
* `preferred_ip_v6_address`:(string) The preferred IPv6 address can be specified only for dynamic lease requests. Intersight will make its best effort to allocate that IPv6 address if it is available in the pool. If the specified preferred IPv6 address is not in the range of the pool or if it is already leased or reserved, then the next available IPv6 address from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease request will fail if it specifies the preferred IPv6 address property. When the preferred IPv6 address property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be replaced by the new lease. Migration also supported only for dynamic lease requests. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
