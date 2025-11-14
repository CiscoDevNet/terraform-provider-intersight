---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_hcl_status_job"
description: |-
        ### Overview
        The HclStatusJob object is utilized to batch inventory notifications and process evaluations of HclStatus. It ensures timely and efficient handling of compatibility checks within Cisco's hardware ecosystem.
        #### Purpose
        HclStatusJob is designed to manage and execute batch jobs related to hardware compatibility evaluations, ensuring that managed objects are consistently assessed against the HCL. It automates the processing of inventory updates to maintain compliance.
        #### Key Concepts
        - **Batch Processing:** - Handles inventory notifications in batches, optimizing the evaluation of HclStatus.
        - **Automation:** - Automates the assessment of compatibility, reducing manual intervention and enhancing efficiency.
        - **Relationship with Managed Objects:** - Links directly with managed objects to facilitate seamless compatibility evaluations.
        - **Security:** - Ensures evaluations are performed securely by inheriting permissions from related device registrations.

---

# Data Source: intersight_cond_hcl_status_job
### Overview
The HclStatusJob object is utilized to batch inventory notifications and process evaluations of HclStatus. It ensures timely and efficient handling of compatibility checks within Cisco's hardware ecosystem.
#### Purpose
HclStatusJob is designed to manage and execute batch jobs related to hardware compatibility evaluations, ensuring that managed objects are consistently assessed against the HCL. It automates the processing of inventory updates to maintain compliance.
#### Key Concepts
- **Batch Processing:** - Handles inventory notifications in batches, optimizing the evaluation of HclStatus.
- **Automation:** - Automates the assessment of compatibility, reducing manual intervention and enhancing efficiency.
- **Relationship with Managed Objects:** - Links directly with managed objects to facilitate seamless compatibility evaluations.
- **Security:** - Ensures evaluations are performed securely by inheriting permissions from related device registrations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_hcl_status_job.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
