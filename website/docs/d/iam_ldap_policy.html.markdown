---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_policy"
description: |-
  LDAP Policy configurations.
---

# Data Source: intersight_iam_ldap_policy
LDAP Policy configurations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `enable_dns`:(bool) Enables DNS to access LDAP servers. 
* `enabled`:(bool) LDAP server performs authentication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `user_search_precedence`:(string) Search precedence between local user database and LDAP user database.* `LocalUserDb` - Precedence is given to local user database while searching.* `LDAPUserDb` - Precedence is given to LADP user database while searching. 
