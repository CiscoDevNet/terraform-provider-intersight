
---
layout: "intersight"
page_title: "Intersight: intersight_iam_idp"
sidebar_current: "docs-intersight-data-source-iam-idp"
description: |-
The SAML identity provider such as Cisco, that has been used to log in to Intersight.
---

# Data Source: intersight_iam._idp
The SAML identity provider such as Cisco, that has been used to log in to Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `domain_name`:(string) Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `idp_entity_id`:(string) The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider. 
* `metadata`:(string) SAML metadata of the IdP. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Identity Provider, for example Cisco, Okta, or OneID. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `type`:(string) Authentication protocol used by the IdP.* `saml` - Use SAML as the authentication protocol for sign-on.* `oidc` - Open ID connect to be used as an authentication protocol for sign-on.* `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP. 
