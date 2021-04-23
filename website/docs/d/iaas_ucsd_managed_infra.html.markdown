---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_managed_infra"
description: |-
  Describes about UCSD Managed infrastructure statistics.
---

# Data Source: intersight_iaas_ucsd_managed_infra
Describes about UCSD Managed infrastructure statistics.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_ucsd_managed_infra.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advanced_catalog_count`:(int) Total advanced catalogs in UCSD. 
* `bm_catalog_count`:(int) Total bare metal catalogs in UCSD. 
* `container_catalog_count`:(int) Total service container catalogs in UCSD. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `esxi_host_count`:(int) Total ESXi hosts in UCSD. 
* `external_group_count`:(int) Total external (Ldap) groups in UCSD. 
* `hyperv_host_count`:(int) Total HyperV hosts in UCSD. 
* `local_group_count`:(int) Total local groups in UCSD. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `standard_catalog_count`:(int) Total standard catalogs in UCSD. 
* `user_count`:(int) Total user accounts in UCSD. 
* `vdc_count`:(int) Total virtual datacenters in UCSD. 
* `vm_count`:(int) Total Virtual machines in UCSD. 
 