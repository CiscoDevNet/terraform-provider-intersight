---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_ncc_check_policy"
description: |-
        A Nutanix Cluster Check (NCC) policy configured on the Prism Central.

---

# Data Source: intersight_hci_ncc_check_policy
A Nutanix Cluster Check (NCC) policy configured on the Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_ncc_check_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) NCC policy description text. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `entity_type`:(string) Entity type against which the alert is raised. 
* `impact_type`:(string) Impact type to which this rule applies. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the NCC policy. 
* `ncc_policy_ext_id`:(string) Unique ID of the NCC policy. 
* `pc_ext_id`:(string) Unique identifier of the domain manager (Prism Central) instance which owns this policy. 
* `policy_id`:(string) Unique ID associated with the policy. 
* `publisher`:(string) Publisher of the NCC policy. 
* `scope`:(string) Scope for the polcy execution. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sub_type`:(string) Sub-type classification of the NCC policy. 
* `title`:(string) The title of the NCC policy. 
* `type`:(string) Defines the type of the NCC policy. 
 
