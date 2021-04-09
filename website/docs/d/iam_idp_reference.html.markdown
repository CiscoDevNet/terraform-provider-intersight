---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_idp_reference"
description: |-
  Default Cisco IdP for authentication.
---

# Data Source: intersight_iam_idp_reference
Default Cisco IdP for authentication.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_idp_reference.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `domain_name`:(string) The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `idp_entity_id`:(string) Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multi_factor_authentication`:(bool) The flag represents if the second factor of authentication is required for Cisco IdP users. 
* `name`:(string) Cisco IdP reference in an account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 