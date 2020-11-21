
---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_policy"
sidebar_current: "docs-intersight-data-source-iam-ldap-policy"
description: |-
LDAP Policy configurations.
---

# Data Source: intersight_iam._ldap_policy
LDAP Policy configurations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `enable_dns`:(bool) Enables DNS to access LDAP servers. 
* `enabled`:(bool) LDAP server performs authentication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `user_search_precedence`:(string) Search precedence between local user database and LDAP user database.* `LocalUserDb` - Precedence is given to local user database while searching.* `LDAPUserDb` - Precedence is given to LADP user database while searching. 
