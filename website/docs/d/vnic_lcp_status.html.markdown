---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_lcp_status"
description: |-
        An internal MO to check if a LCP can be deployed or not on a specific Server Profile.

---

# Data Source: intersight_vnic_lcp_status
An internal MO to check if a LCP can be deployed or not on a specific Server Profile.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_lcp_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reason`:(string) The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Indicates if the LCP is ready for Deploy or not.* `ok` - No issues with the LCP/SCP/VIF.* `error` - The LCP/SCP/VIF cannot be deployed due to error.* `validating` - Validation in progress for the LCP. 
 
