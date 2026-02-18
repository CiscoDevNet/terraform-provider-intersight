---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_cui_integration"
description: |-
        Stores CUI and CCC integration details and synchronization status for an account.

---

# Data Source: intersight_iam_cui_integration
Stores CUI and CCC integration details and synchronization status for an account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_cui_integration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cisco_cloud_control_inventory_sync`:(string) CCC inventory synchronization setting.* `enable` - Cisco Cloud Control inventory synchronization is enabled.* `disable` - Cisco Cloud Control inventory synchronization is disabled. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sync_status`:(string) Current synchronization status of CUI integration data.* `Needs-Sync` - CUI integration data needs to be synchronized.* `InProgress` - CUI integration synchronization is in progress.* `Done` - CUI integration synchronization is completed. 
* `tenant_id`:(string) Primary tenant ID for CUI integration. 
 
