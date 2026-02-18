---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_scrub_policy"
description: |-
        The ScrubPolicy object defines the parameters for data cleanup on servers, particularly during server profile undeployments. It ensures that server components are effectively scrubbed as per configuration.
        #### Purpose
        ScrubPolicy provides a framework for executing data cleanup tasks during server profile operations, safeguarding data integrity and compliance with policies. It plays a critical role in managing server disassociations.
        #### Key Concepts
        - **Policy Definition:** Sets clear guidelines for scrub operations, detailing which components should be cleaned during undeployments.
        - **Integration:** Links with server profiles to ensure scrub operations are executed as part of broader server management tasks.
        - **Security and Compliance:** Supports compliance by ensuring data cleanup is performed consistently and securely.
        - **Access Control:** Establishes privilege sets for policy management, ensuring controlled access to scrub configurations.

---

# Data Source: intersight_compute_scrub_policy
The ScrubPolicy object defines the parameters for data cleanup on servers, particularly during server profile undeployments. It ensures that server components are effectively scrubbed as per configuration.
#### Purpose
ScrubPolicy provides a framework for executing data cleanup tasks during server profile operations, safeguarding data integrity and compliance with policies. It plays a critical role in managing server disassociations.
#### Key Concepts
- **Policy Definition:** Sets clear guidelines for scrub operations, detailing which components should be cleaned during undeployments.
- **Integration:** Links with server profiles to ensure scrub operations are executed as part of broader server management tasks.
- **Security and Compliance:** Supports compliance by ensuring data cleanup is performed consistently and securely.
- **Access Control:** Establishes privilege sets for policy management, ensuring controlled access to scrub configurations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_scrub_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
