---
subcategory: "access"
layout: "intersight"
page_title: "Intersight: intersight_access_ip_address"
description: |-
        The Access IpAddress object is a read-only inventory object that represents the IP address lease information for a specific server profile. This provides a consolidated view of the management IP addresses assigned to a server's management controller.
        #### Purpose
        This serves as a status and reporting entity, not a configurable one. Its purpose is to display the IPv4 and IPv6 addresses that have been leased to a server profile from the IP pools specified in the corresponding Access Policy. This allows administrators and automation systems to easily query the assigned management IP addresses for a given server.
        #### Key Concepts
        - **Inventory Object:** It reflects the current state of IP allocation and is not directly configurable by users. Its data is populated by the system after an Access policy is successfully deployed.
        - **Read-Only Access:** The API methods for this are limited to READ, reinforcing its role as an informational object.
        - **Lease Association:** It maintains relationships to the underlying IP lease objects (ippool.IpLease), providing a clear link between the assigned address and its source pool.
        - **Profile-Specific:** Each IpAddress instance is tied to a single server profile, showing the specific addresses in use by that server.

---

# Data Source: intersight_access_ip_address
The Access IpAddress object is a read-only inventory object that represents the IP address lease information for a specific server profile. This provides a consolidated view of the management IP addresses assigned to a server's management controller.
#### Purpose
This serves as a status and reporting entity, not a configurable one. Its purpose is to display the IPv4 and IPv6 addresses that have been leased to a server profile from the IP pools specified in the corresponding Access Policy. This allows administrators and automation systems to easily query the assigned management IP addresses for a given server.
#### Key Concepts
- **Inventory Object:** It reflects the current state of IP allocation and is not directly configurable by users. Its data is populated by the system after an Access policy is successfully deployed.
- **Read-Only Access:** The API methods for this are limited to READ, reinforcing its role as an informational object.
- **Lease Association:** It maintains relationships to the underlying IP lease objects (ippool.IpLease), providing a clear link between the assigned address and its source pool.
- **Profile-Specific:** Each IpAddress instance is tied to a single server profile, showing the specific addresses in use by that server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_access_ip_address.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ipv4_address`:(string) IPv4 Address leased to this Server Profile for In-Band Deployment. 
* `ipv6_address`:(string) IPv4 Address leased to this Server Profile for In-Band Deployment. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `oob_ipv4_address`:(string) IPv4 Address leased to this Server Profile for Out-Of-Band deployment. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
