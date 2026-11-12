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
To access the ith object of the results obtained, use `data.intersight_iam_idp.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `certificate_expiry_time`:(string) Expiration time of the certificate in the SP metadata for the IdP. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `domain_name`:(string) Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `enable_single_logout`:(bool) Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP. 
* `idp_entity_id`:(string) The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider. 
* `idp_metadata_expiry_state`:(string) The expiration state of the certificates in the IdP metadata.* `expiration-beyond-warning-period` - The expiry date is not within the warning period.* `expiration-within-info-period` - The expiry date is within the info period.* `expiration-within-warning-period` - The expiry date is within the warning period.* `expiration-within-critical-period` - The expiry date is within the critical period.* `expired` - The expiry date has passed. 
* `metadata`:(string) SAML metadata of the IdP. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Identity Provider, for example Cisco, Okta, or OneID. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_warning`:(bool) When users attempt the Account URL login with an unverified Domain Name, they get a warning stating that they are logging in using an unverified Domain Name. Enable the slider if you do not wish to see the warning message. 
* `sp_cert_rotated`:(bool) Status of the SP certificate rotation for the IdP. 
* `sp_metadata_expiry_state`:(string) The expiration state of the certificate in the SP metadata used by this IdP. The state is not-applicable when SAML assertion encryption and AuthNRequest signing are both disabled.* `expiration-within-info-period` - The expiry date is beyond 270 days.* `expiration-within-270-day-info-period` - The expiry date is within 270 days but beyond 180 days.* `expiration-within-180-day-info-period` - The expiry date is within 180 days but beyond 90 days.* `expiration-within-90-day-warning-period` - The expiry date is within 90 days but beyond 60 days.* `expiration-within-60-day-warning-period` - The expiry date is within 60 days but beyond 30 days.* `expiration-within-30-day-critical-period` - The expiry date is within 30 days but beyond 7 days.* `expiration-within-critical-period` - The expiry date is within 7 days.* `expired` - The expiry date has passed.* `not-applicable` - The SP certificate is not used because SAML assertion encryption and AuthNRequest signing are both disabled. 
* `sp_metadata_hash_method`:(string) Determines if the AuthNRequest sent to the Idp should be signed by the new (SHA-2) or the old (SHA-1) SP metadata during the grace period of SP SAML metadata rotation.* `Sha2` - The SAML SP metadata hash is generated using SHA-256 algorithm.* `Sha1` - The SAML SP metadata hash is generated using SHA-1 algorithm. 
* `type`:(string) Authentication protocol used by the IdP.* `saml` - Use SAML as the authentication protocol for sign-on.* `oidc` - Open ID connect to be used as an authentication protocol for sign-on.* `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.* `guest` - Guest users accessing Intersight are assigned to a guest Identity Provider (IdP). They do not have single sign-on options available and must log in  using their email address and a one-time password (OTP). Upon successful authentication, all guest users associated with the account are grouped under  the guest IdP. Guest IdP is primarily used in cloud environment. 
 
