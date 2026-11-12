---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_domain_name_info"
description: |-
        DomainNameInfos represent verified (or pending verification) email domain ownership records within an account. They support DNS TXT-record validation and maintain verification lifecycle state.
        #### Purpose
        Establish and track domain ownership so login routing and IdP association based on email domain can be trusted and governed.
        #### Key Concepts
        - **DNS-based verification:** Uses TXT records to prove control of the domain.
        - **Lifecycle status:** Tracks Pending/Verified/Failed/Expired verification states and expiry timestamps.
        - **Action-driven updates:** Supports regenerate vs verify flows via an action field with conditional privileges.
        - **Account-scoped trust signal:** Anchored to the account and used to strengthen identity-routing integrity.

---

# Data Source: intersight_iam_domain_name_info
DomainNameInfos represent verified (or pending verification) email domain ownership records within an account. They support DNS TXT-record validation and maintain verification lifecycle state.
#### Purpose
Establish and track domain ownership so login routing and IdP association based on email domain can be trusted and governed.
#### Key Concepts
- **DNS-based verification:** Uses TXT records to prove control of the domain.
- **Lifecycle status:** Tracks Pending/Verified/Failed/Expired verification states and expiry timestamps.
- **Action-driven updates:** Supports regenerate vs verify flows via an action field with conditional privileges.
- **Account-scoped trust signal:** Anchored to the account and used to strengthen identity-routing integrity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_domain_name_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Regenerate TXT record and validate TXT record.* `generate` - Generate TXT record for domain name ownership validation.* `verify` - Verify TXT record for domain name ownership validation. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `domain_name`:(string) Email domain name. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `record_expiry_time`:(string) Expiration time of TXT Record. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of Domain Ownership Verification.* `Pending` - Domain verification is pending.* `Failed` - Domain verification failed. Re-generate token and verify.* `Verified` - Domain verification succeeded.* `Expired` - TXT Record for Domain verification expired. 
* `txt_record`:(string) Resource record used to verify Domain Ownership. Users need to verify the domain by adding the TXT Record in their DNS Host. 
 
