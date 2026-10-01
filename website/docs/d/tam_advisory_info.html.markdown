---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_info"
description: |-
        Per-account acknowledgement bookkeeping for advisories; it records which advisories a user has acknowledged, and does not itself list the servers or components an advisory impacts (that is tam.AdvisoryInstance). A record is generally present only for advisories that have been acknowledged, so an advisory with no AdvisoryInfo record is simply unacknowledged. The reliable way to find unacknowledged advisories is therefore by exclusion - take the impacting advisories from tam.AdvisoryInstance and remove those whose Advisory is acknowledged here - rather than by querying this object for active records. Acknowledgement records can also persist after an advisory has stopped impacting the account, so a record here does not by itself mean the advisory still applies; when reporting acknowledged advisories that currently impact the account, intersect these acknowledged advisories with the ones actually present in tam.AdvisoryInstance. When asked how many advisories have been acknowledged, answer with that intersected count - the acknowledged advisories that currently impact the account - rather than the raw record total, unless the user explicitly asks for all acknowledgement records including ones no longer applicable.
        The AdvisoryInfo object captures the state of Security Advisories and Field Notices, including whether an advisory is visible in the Intersight account, as depicted by its acknowledgement status.
        #### Purpose
        AdvisoryInfo serves as a personalized repository of advisory states for each account, allowing users to manage and track advisories relevant to their environment.
        #### Key Concepts
        - **State Management:** Tracks whether advisories are active or acknowledged, reflecting user preferences for updates.
        - **Account Association:** Directly links advisory states to specific accounts for tailored management.
        - **Shared Resource:** Supports collaborative management of advisories within shared organizational contexts.
        - **Controlled Permissions:** Restricts creation, updating, and deletion of advisory states to Account Administrators.

---

# Data Source: intersight_tam_advisory_info
Per-account acknowledgement bookkeeping for advisories; it records which advisories a user has acknowledged, and does not itself list the servers or components an advisory impacts (that is tam.AdvisoryInstance). A record is generally present only for advisories that have been acknowledged, so an advisory with no AdvisoryInfo record is simply unacknowledged. The reliable way to find unacknowledged advisories is therefore by exclusion - take the impacting advisories from tam.AdvisoryInstance and remove those whose Advisory is acknowledged here - rather than by querying this object for active records. Acknowledgement records can also persist after an advisory has stopped impacting the account, so a record here does not by itself mean the advisory still applies; when reporting acknowledged advisories that currently impact the account, intersect these acknowledged advisories with the ones actually present in tam.AdvisoryInstance. When asked how many advisories have been acknowledged, answer with that intersected count - the acknowledged advisories that currently impact the account - rather than the raw record total, unless the user explicitly asks for all acknowledgement records including ones no longer applicable.
The AdvisoryInfo object captures the state of Security Advisories and Field Notices, including whether an advisory is visible in the Intersight account, as depicted by its acknowledgement status.
#### Purpose
AdvisoryInfo serves as a personalized repository of advisory states for each account, allowing users to manage and track advisories relevant to their environment.
#### Key Concepts
- **State Management:** Tracks whether advisories are active or acknowledged, reflecting user preferences for updates.
- **Account Association:** Directly links advisory states to specific accounts for tailored management.
- **Shared Resource:** Supports collaborative management of advisories within shared organizational contexts.
- **Controlled Permissions:** Restricts creation, updating, and deletion of advisory states to Account Administrators.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory. Because a record is usually created only when an advisory is acknowledged, the absence of an AdvisoryInfo record for an advisory means it is unacknowledged; this value is most reliably used to collect the acknowledged advisories ('acknowledged'), which can then be excluded from the impacting advisories in tam.AdvisoryInstance. An 'acknowledged' record may also be stale - it can remain after the advisory no longer impacts the account - so to report acknowledged advisories that still apply, intersect this acknowledged set with tam.AdvisoryInstance.* `active` - Advisory is currently active and the user wants to receive updates for this advisory.* `acknowledged` - Advisory is seen and acknowledged by the user and she no longer wants to recieve updates. 
 
