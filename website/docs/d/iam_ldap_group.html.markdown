
---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_group"
sidebar_current: "docs-intersight-data-source-iam-ldap-group"
description: |-
Mapping of LDAP Group to EndPointRoles.
---

# Data Source: intersight_iam._ldap_group
Mapping of LDAP Group to EndPointRoles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `domain`:(string) LDAP server domain the Group resides in. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) LDAP Group name in the LDAP server database. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
