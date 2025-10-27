---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_instance"
description: |-
        # Overview
        The AdvisoryInstance object represents instances of advisories applicable to Intersight managed objects,
        maintaining advisory applicability and state changes.
        ## Purpose
        AdvisoryInstance tracks the lifecycle and applicability of advisories for specific managed objects,
        ensuring users have visibility into which devices or components are impacted and require resolution.
        ## Key Concepts
        - **Lifecycle Tracking** – Maintains records of advisory applicability, including state changes over time.
        - **Managed Object Integration** – Directly links advisories to affected managed objects for precise applicability.
        - **Controlled Access** – Provides access and management capabilities based on administrative privileges.

---

# Data Source: intersight_tam_advisory_instance
# Overview
The AdvisoryInstance object represents instances of advisories applicable to Intersight managed objects,  
maintaining advisory applicability and state changes.
## Purpose
AdvisoryInstance tracks the lifecycle and applicability of advisories for specific managed objects,  
ensuring users have visibility into which devices or components are impacted and require resolution.
## Key Concepts
- **Lifecycle Tracking** – Maintains records of advisory applicability, including state changes over time.
 - **Managed Object Integration** – Directly links advisories to affected managed objects for precise applicability.
- **Controlled Access** – Provides access and management capabilities based on administrative privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `affected_object_moid`:(string) Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `affected_object_type`:(string) Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_state_change_time`:(string) Timestamp when a state change was observed on this advisory instnace. 
* `last_verified_time`:(string) Timestamp when this advisory was last evaluated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory instance (Active/Cleared/Unknown etc.).* `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object.* `active` - Advisory instance is currently active and applicable for the affected managed object.* `cleared` - Advisory instance is no longer applicable for the affected managed object. 
 
