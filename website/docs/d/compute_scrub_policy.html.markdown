---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_scrub_policy"
description: |-
        The scrub policy models the data cleanup on the host. If scrub policy is associated, then its configuration will be used to initiate data erase during the server profile undeploy or disassociation.

---

# Data Source: intersight_compute_scrub_policy
The scrub policy models the data cleanup on the host. If scrub policy is associated, then its configuration will be used to initiate data erase during the server profile undeploy or disassociation.
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
 
