---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_universe"
description: |-
        Universes are bookkeeping containers used to track all allocated and reserved identities for a given Intersight account and pool type (for example, WWNN/WWPN identity pools). They provide the account-scoped boundary in which identity uniqueness and allocation tracking can be maintained consistently.
        #### Purpose
        Maintain an authoritative, account-level container for identity tracking so pool operations (allocation, reservation, membership) can reference a consistent universe for correlation, audit, and lifecycle management.
        #### Key Concepts
        - **Account-scoped identity domain**: The `account` relationship anchors the Universe to a specific `iam.Account`, establishing the scope for identity tracking.
        - **Pool-type bookkeeping container**: Intended as the shared reference point for “all IDs” of a particular pool type within an account (e.g., WWNN/WWPN pools).
        - **Read-focused consumption**: Exposed via READ to support inspection and correlation rather than direct user-driven manipulation.
        - **Lifecycle coupling to account**: `onpeerdelete: cascade` on `account` ties Universe lifecycle to the account—when the account is removed, the Universe is removed as well.
        - **Broad operational visibility**: READ privileges include both general read roles and WWNN/WWPN pool read roles, supporting inventory and troubleshooting workflows.

---

# Data Source: intersight_fcpool_universe
Universes are bookkeeping containers used to track all allocated and reserved identities for a given Intersight account and pool type (for example, WWNN/WWPN identity pools). They provide the account-scoped boundary in which identity uniqueness and allocation tracking can be maintained consistently.
#### Purpose
Maintain an authoritative, account-level container for identity tracking so pool operations (allocation, reservation, membership) can reference a consistent universe for correlation, audit, and lifecycle management.
#### Key Concepts
- **Account-scoped identity domain**: The `account` relationship anchors the Universe to a specific `iam.Account`, establishing the scope for identity tracking.
- **Pool-type bookkeeping container**: Intended as the shared reference point for “all IDs” of a particular pool type within an account (e.g., WWNN/WWPN pools).
- **Read-focused consumption**: Exposed via READ to support inspection and correlation rather than direct user-driven manipulation.
- **Lifecycle coupling to account**: `onpeerdelete: cascade` on `account` ties Universe lifecycle to the account—when the account is removed, the Universe is removed as well.
- **Broad operational visibility**: READ privileges include both general read roles and WWNN/WWPN pool read roles, supporting inventory and troubleshooting workflows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_universe.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
