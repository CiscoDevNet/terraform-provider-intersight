---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_policy"
description: |-
        ### Overview
        The LDAP Policy object is a reusable policy for configuring Lightweight Directory Access Protocol (LDAP) settings, enabling remote user authentication for endpoints like servers and chassis.
        #### Purpose
        The primary purpose of an LDAP Policy is to integrate endpoints with a central directory service (like Microsoft Active Directory or OpenLDAP) for user authentication. This centralizes user management and allows administrators to grant access to Intersight-managed devices based on existing corporate directory groups and users, rather than creating separate local accounts on each device.
        #### Key Concepts
        - **Centralized Authentication:** - Enables endpoints to authenticate users against one or more external LDAP servers.
        - **Reusable and Profile-Based:** - As a policy object, this provides a standardized LDAP configuration that can be applied to multiple server or chassis profiles.
        - **Flexible Configuration:** - Supports various settings including the Base DN for searches, bind methods (Anonymous, Configured Credentials, Login Credentials), and connection settings like encryption and timeout.
        - **Group-to-Role Mapping:** - Works with LdapGroup objects to map directory groups to local endpoint roles, enabling granular, role-based access control for remote users.
        - **Provider and DNS Settings:** - Manages a list of LDAP servers (LdapProvider objects) and can be configured to use DNS for server discovery.

---

# Data Source: intersight_iam_ldap_policy
### Overview
The LDAP Policy object is a reusable policy for configuring Lightweight Directory Access Protocol (LDAP) settings, enabling remote user authentication for endpoints like servers and chassis.
#### Purpose
The primary purpose of an LDAP Policy is to integrate endpoints with a central directory service (like Microsoft Active Directory or OpenLDAP) for user authentication. This centralizes user management and allows administrators to grant access to Intersight-managed devices based on existing corporate directory groups and users, rather than creating separate local accounts on each device.
#### Key Concepts
- **Centralized Authentication:** - Enables endpoints to authenticate users against one or more external LDAP servers.
- **Reusable and Profile-Based:** - As a policy object, this provides a standardized LDAP configuration that can be applied to multiple server or chassis profiles.
- **Flexible Configuration:** - Supports various settings including the Base DN for searches, bind methods (Anonymous, Configured Credentials, Login Credentials), and connection settings like encryption and timeout.
- **Group-to-Role Mapping:** - Works with LdapGroup objects to map directory groups to local endpoint roles, enabling granular, role-based access control for remote users.
- **Provider and DNS Settings:** - Manages a list of LDAP servers (LdapProvider objects) and can be configured to use DNS for server discovery.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_ldap_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_dns`:(bool) Enables DNS to access LDAP servers. 
* `enabled`:(bool) LDAP server performs authentication. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_search_precedence`:(string) Search precedence between local user database and LDAP user database.* `LocalUserDb` - Precedence is given to local user database while searching.* `LDAPUserDb` - Precedence is given to LADP user database while searching. 
 
