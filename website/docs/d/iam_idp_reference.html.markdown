---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_idp_reference"
description: |-
        The IdpReference object represents the system-provided Cisco IdP association within an account, acting as the default reference for user and group assignments.
        #### Purpose
        IdpReference simplifies managing the default IdP relationship for user onboarding, preferences, and access.
        #### Key Concepts
        - **Default IdP linkage:** Provides a stable, system-backed IdP reference.
        - **User/group Association:** Central to managing users and groups tied to the Cisco IdP.
        - **Preference Anchor:** Links user UI preferences and settings to the IdP context.
        - **Account Scope:** Managed within the owning account's boundaries.

---

# Data Source: intersight_iam_idp_reference
The IdpReference object represents the system-provided Cisco IdP association within an account, acting as the default reference for user and group assignments.
#### Purpose
IdpReference simplifies managing the default IdP relationship for user onboarding, preferences, and access.
#### Key Concepts
- **Default IdP linkage:** Provides a stable, system-backed IdP reference.
- **User/group Association:** Central to managing users and groups tied to the Cisco IdP.
- **Preference Anchor:** Links user UI preferences and settings to the IdP context.
- **Account Scope:** Managed within the owning account's boundaries.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_idp_reference.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `domain_name`:(string) The DomainNameInfo object represents an organization’s email domain used during login mapping and ownership verification. It supports verifying domain control via DNS to enhance trust.#### PurposeDomainNameInfo ensures that only verified domains are used to route users to the correct identity provider, improving login integrity.#### Key Concepts- **Domain Routing:** Associates email domains with IdPs for seamless login.- **Ownership Verification:** Supports DNS-based verification for trust establishment.- **Status Tracking:** Maintains domain verification lifecycle state.- **Account Scoped:** Managed within the context of the owning account. 
* `idp_entity_id`:(string) Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multi_factor_authentication`:(bool) The flag represents if the second factor of authentication is required for Cisco IdP users. 
* `name`:(string) Cisco IdP reference in an account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
