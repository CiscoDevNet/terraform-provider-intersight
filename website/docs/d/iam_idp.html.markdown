---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_idp"
description: |-
  The SAML identity provider such as Cisco, that has been used to log in to Intersight.
---

# Data Source: intersight_iam_idp
The SAML identity provider such as Cisco, that has been used to log in to Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_idp.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `domain_name`:(string) Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `enable_single_logout`:(bool) Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP. 
* `idp_entity_id`:(string) The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider. 
* `metadata`:(string) SAML metadata of the IdP. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Identity Provider, for example Cisco, Okta, or OneID. 
* `type`:(string) Authentication protocol used by the IdP.* `saml` - Use SAML as the authentication protocol for sign-on.* `oidc` - Open ID connect to be used as an authentication protocol for sign-on.* `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP. 
 