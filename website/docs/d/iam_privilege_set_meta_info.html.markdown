---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege_set_meta_info"
description: |-
        PrivilegeSetMetaInfo centralizes metadata to validate, relate, and maintain privilege sets. It captures the context needed to evaluate dependencies and associations, ensuring consistency with access policies.
        #### Purpose
        PrivilegeSetMetaInfo validates privilege sets before persistence by tracking upstream and downstream relationships. It identifies missing dependencies and surfaces insights to ensure secure, confident design.
        #### Key Concepts
        - **Dependency Intelligence:** Continuously analyzes privilege-set relationships to reveal required, computed, and missing dependencies, helping teams remediate gaps before deployment.
        - **Association Awareness:** Maintains links between related privilege sets, providing a holistic view of how access models interact within broader governance frameworks.
        - **Validation-First Workflow:** Operates as a non-persistent validation layer, allowing privilege configurations to be assessed, iterated, and approved before becoming part of the active system.
        - **Access Governance Alignment:** Integrates with defined privilege sets and management roles to ensure only authorized stakeholders can inspect or adjust metadata, supporting least-privilege best practices.

---

# Data Source: intersight_iam_privilege_set_meta_info
PrivilegeSetMetaInfo centralizes metadata to validate, relate, and maintain privilege sets. It captures the context needed to evaluate dependencies and associations, ensuring consistency with access policies.
#### Purpose
PrivilegeSetMetaInfo validates privilege sets before persistence by tracking upstream and downstream relationships. It identifies missing dependencies and surfaces insights to ensure secure, confident design.
#### Key Concepts
- **Dependency Intelligence:** Continuously analyzes privilege-set relationships to reveal required, computed, and missing dependencies, helping teams remediate gaps before deployment.
- **Association Awareness:** Maintains links between related privilege sets, providing a holistic view of how access models interact within broader governance frameworks.
- **Validation-First Workflow:** Operates as a non-persistent validation layer, allowing privilege configurations to be assessed, iterated, and approved before becoming part of the active system.
- **Access Governance Alignment:** Integrates with defined privilege sets and management roles to ensure only authorized stakeholders can inspect or adjust metadata, supporting least-privilege best practices.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_privilege_set_meta_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the privilege set. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the privilege set. 
 
