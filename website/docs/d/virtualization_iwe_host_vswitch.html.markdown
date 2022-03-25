---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_iwe_host_vswitch"
description: |-
        A Intersight Workload Engine vSwitch entity that is part of a cluster wide dvSwitch.

---

# Data Source: intersight_virtualization_iwe_host_vswitch
A Intersight Workload Engine vSwitch entity that is part of a cluster wide dvSwitch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_iwe_host_vswitch.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name`:(string) The name of the host to which this vSwitch belongs to. 
* `identity`:(string) The internally generated identity of this switch. This entity is not manipulated by users. It aids in uniquely identifying the switch object. For VMware, this is MOR (managed object reference). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the switch. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
