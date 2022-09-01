---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_license"
description: |-
        NetApp licenses for NetApp Ontap.

---

# Data Source: intersight_storage_net_app_license
NetApp licenses for NetApp Ontap.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_license.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_uuid`:(string) Unique identity of the device. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the license package. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Summary state of license package based on all installed licenses.* `Unknown` - The summary state of the license package is unknown.* `Compliant` - The summary state of the license package is compliant.* `Noncompliant` - The summary state of the license package is noncompliant.* `Unlicensed` - The summary state of the license package is unlicensed. 
 
