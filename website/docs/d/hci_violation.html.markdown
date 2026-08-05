---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_violation"
description: |-
        The Violation object identifies instances of license violations within Nutanix Prism Central. This provides information for tracking license violation issues.
        #### Purpose
        The Violation object provides a view of license violations on the cluster. It reports features which exceed their licensed capacity as well as expired licenses.
        #### Key Concepts
        - **Compliance Monitoring:** Tracks instances of license violations, supporting timely resolution.
        - **Issue Management:** Provides insights into feature and capacity violations, aiding in corrective actions.

---

# Data Source: intersight_hci_violation
The Violation object identifies instances of license violations within Nutanix Prism Central. This provides information for tracking license violation issues.
#### Purpose 
The Violation object provides a view of license violations on the cluster. It reports features which exceed their licensed capacity as well as expired licenses.
#### Key Concepts  
- **Compliance Monitoring:** Tracks instances of license violations, supporting timely resolution. 
- **Issue Management:** Provides insights into feature and capacity violations, aiding in corrective actions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_violation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The unique identifier of the cluster that has the violations. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_multicluster`:(bool) Indicates if the violation is on multi-cluster which is PrismCentral in Nutanix's term.Otherwise, it is on a Prism Element cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
