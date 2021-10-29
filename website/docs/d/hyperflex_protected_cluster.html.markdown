---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_protected_cluster"
description: |-
  Object for the protected clusters view.
---

# Data Source: intersight_hyperflex_protected_cluster
Object for the protected clusters view.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_protected_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hx_version`:(string) Version of the Hyperflex cluster. 
* `hypervisor_version`:(string) The version of hypervisor running on this cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protected_datastore_name`:(string) Name of the protected datastore. 
* `protected_vms_count`:(int) Number of VMs protected on this cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_cluster_name`:(string) Name of the source cluster. 
* `target_cluster_name`:(string) Name of the target cluster. 
* `target_datastore_name`:(string) Name of the target datastore. 
* `target_datastore_utilization`:(float) Percent usage of the datastore. 
 