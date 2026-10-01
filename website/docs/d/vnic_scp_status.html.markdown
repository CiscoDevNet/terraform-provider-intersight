---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_scp_status"
description: |-
        The ScpStatus object checks if a SAN Connectivity Policy (SCP) can be deployed on a specific server profile.
        #### Purpose
        It validates the deployment readiness of an SCP, reporting any errors before the deployment is executed.
        #### Key Concepts
        - **Deployment Validation:** Assesses if an SCP is ready for deployment.
        - **Status Reporting:** Provides detailed reasons for validation failures.

---

# Data Source: intersight_vnic_scp_status
The ScpStatus object checks if a SAN Connectivity Policy (SCP) can be deployed on a specific server profile.
#### Purpose
It validates the deployment readiness of an SCP, reporting any errors before the deployment is executed.
#### Key Concepts
- **Deployment Validation:** Assesses if an SCP is ready for deployment.
- **Status Reporting:** Provides detailed reasons for validation failures.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_scp_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reason`:(string) The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Indicates if the LCP is ready for Deploy or not.* `ok` - No issues with the LCP/SCP/VIF.* `error` - The LCP/SCP/VIF cannot be deployed due to error.* `validating` - Validation in progress for the LCP. 
 
