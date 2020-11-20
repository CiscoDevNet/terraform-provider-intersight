
---
layout: "intersight"
page_title: "Intersight: intersight_iam_idp_reference"
sidebar_current: "docs-intersight-data-source-iam-idp-reference"
description: |-
Default Cisco IdP for authentication.
---

# Data Source: intersight_iam._idp_reference
Default Cisco IdP for authentication.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `domain_name`:(string) The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `idp_entity_id`:(string) Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multi_factor_authentication`:(bool) The flag represents if the second factor of authentication is required for Cisco IdP users. 
* `name`:(string) Cisco IdP reference in an account. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
