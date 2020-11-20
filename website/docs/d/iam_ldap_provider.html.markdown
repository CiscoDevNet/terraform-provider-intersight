
---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_provider"
sidebar_current: "docs-intersight-data-source-iam-ldap-provider"
description: |-
LDAP Provider or LDAP Server for user authentication.
---

# Data Source: intersight_iam._ldap_provider
LDAP Provider or LDAP Server for user authentication.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `port`:(int) LDAP Server Port for connection establishment. 
* `server`:(string) LDAP Server Address, can be IP address or hostname. 
