---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_count"
description: |-
        The AdvisoryCount object provides a high-level summary of the total number of advisories currently affecting a specific Intersight account.
        #### Purpose
        This object serves as a quick-reference metric, allowing administrators to gauge the overall security and lifecycle posture of their infrastructure. By providing a consolidated count of active advisories—such as Field Notices, PSIRT, or EOL milestones—it enables organizations to prioritize their maintenance and remediation efforts effectively.
        #### Key Concepts
        - **Account-Level Aggregation:** Summarizes the total impact of all active advisories across the entire account, providing a single point of visibility for infrastructure health.
        - **Operational Visibility:** Offers a high-level dashboard metric that helps administrators quickly identify if there is a significant number of advisories requiring attention.
        - **License-Restricted Access:** Access to this object is governed by the 'Essentials' license entitlement, ensuring that capacity planning and maintenance monitoring are available to licensed users.
        - **Account Association:** Directly links the advisory count to a specific Intersight account, ensuring that the data is scoped correctly to the user's organizational context.

---

# Data Source: intersight_tam_advisory_count
The AdvisoryCount object provides a high-level summary of the total number of advisories currently affecting a specific Intersight account.
#### Purpose
This object serves as a quick-reference metric, allowing administrators to gauge the overall security and lifecycle posture of their infrastructure. By providing a consolidated count of active advisories—such as Field Notices, PSIRT, or EOL milestones—it enables organizations to prioritize their maintenance and remediation efforts effectively.
#### Key Concepts
- **Account-Level Aggregation:** Summarizes the total impact of all active advisories across the entire account, providing a single point of visibility for infrastructure health.
- **Operational Visibility:** Offers a high-level dashboard metric that helps administrators quickly identify if there is a significant number of advisories requiring attention.
- **License-Restricted Access:** Access to this object is governed by the 'Essentials' license entitlement, ensuring that capacity planning and maintenance monitoring are available to licensed users.
- **Account Association:** Directly links the advisory count to a specific Intersight account, ensuring that the data is scoped correctly to the user's organizational context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_count.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advisory_count`:(int) Total number of advisories affecting the account. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
