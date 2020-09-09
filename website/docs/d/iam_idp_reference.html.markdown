
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `domain_name`:(string) The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `idp_entity_id`:(string) Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multi_factor_authentication`:(bool) The flag represents if the second factor of authentication is required for Cisco IdP users. 
* `name`:(string) Cisco IdP reference in an account. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
