---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_provider"
description: |-
        ### Overview
        The LDAP Provider object represents a single LDAP server that an endpoint will use for authentication. It is a component of an LdapPolicy.
        #### Purpose
        The purpose is to define the connection details for one LDAP server. An LdapPolicy can contain a list of LdapProvider objects, creating an ordered list of servers that the endpoint will attempt to contact for authentication.
        #### Key Concepts
        - **Server Endpoint:** - Specifies the IP address or hostname and the port of the LDAP server.
        - **Component of a Policy:** - It does not function independently but is always part of an LdapPolicy.
        - **Ordered List:** - The order of LdapProvider objects within the policy determines the failover sequence for authentication attempts.

---

# Data Source: intersight_iam_ldap_provider
### Overview
The LDAP Provider object represents a single LDAP server that an endpoint will use for authentication. It is a component of an LdapPolicy.
#### Purpose
The purpose is to define the connection details for one LDAP server. An LdapPolicy can contain a list of LdapProvider objects, creating an ordered list of servers that the endpoint will attempt to contact for authentication.
#### Key Concepts
- **Server Endpoint:** - Specifies the IP address or hostname and the port of the LDAP server.
- **Component of a Policy:** - It does not function independently but is always part of an LdapPolicy.
- **Ordered List:** - The order of LdapProvider objects within the policy determines the failover sequence for authentication attempts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_ldap_provider.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port`:(int) LDAP Server Port for connection establishment. 
* `server`:(string) LDAP Server Address, can be IP address or hostname. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) LDAP server vendor type used for authentication.* `OpenLDAP` - Open source LDAP server for remote authentication.* `MSAD` - Microsoft active directory for remote authentication. 
 
