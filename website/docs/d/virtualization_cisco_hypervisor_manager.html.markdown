---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_cisco_hypervisor_manager"
description: |-
        A hypervisor manager to manage Cisco Intersight Workload Engine compute clusters and is available per user account.

---

# Data Source: intersight_virtualization_cisco_hypervisor_manager
A hypervisor manager to manage Cisco Intersight Workload Engine compute clusters and is available per user account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_cisco_hypervisor_manager.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `build`:(string) The build number of the Hypervisor Manger (e.g., 4541947, 6.3.9600.18692). The build number may indicate some feature support that applications might rely on. The build number may not always be an integer. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). 
 
