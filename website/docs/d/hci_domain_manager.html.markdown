---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_domain_manager"
description: |-
        The current HCI Prism Central instance reported by Prism Central.

---

# Data Source: intersight_hci_domain_manager
The current HCI Prism Central instance reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_domain_manager.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the domain manager. 
* `pc_ext_id`:(string) The unique identifier of the domain manager (Prism Central) instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(string) The size of the domain manager such as STARTER, SMALL, LARGE, EXTRALARGE.It determines the resources used by the domain manager. 
 
