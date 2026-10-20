---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_wan_edge_device_maintenance_check"
description: |-
        An on-demand user-scoped maintenance check for a set of Unified Edge chassis. Created under a user and populated asynchronously by an Intersight workflow that queries Catalyst SD-WAN Manager and inlines the latest result.

---

# Data Source: intersight_catalystsdwan_wan_edge_device_maintenance_check
An on-demand user-scoped maintenance check for a set of Unified Edge chassis. Created under a user and populated asynchronously by an Intersight workflow that queries Catalyst SD-WAN Manager and inlines the latest result.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_wan_edge_device_maintenance_check.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action requested for this maintenance check. CREATE auto-triggers a refresh, and the server resets it to None when the check reaches a terminal state.* `None` - No refresh is currently requested.* `Refresh` - Refresh the maintenance activity data for the maintenance check. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Current status of the maintenance check.* `None` - Initial state before processing has started.* `InProgress` - The check workflow is running.* `Completed` - The check completed successfully for all entries in targetChassis.* `PartiallyCompleted` - The check completed for a strict subset of targetChassis; the remainder were skipped due to validation or resolution failures. Per-entry failure details are available on the linked workflowInfo. Consumers may treat this as success or failure based on their own strict/lenient policy.* `Failed` - The maintenance check failed before it could produce a usable final activity snapshot for the requested targetChassis set.* `ValidationFailed` - No targetChassis entry resolved to a Secure Router known to Catalyst SD-WAN Manager, so the check could not run. If only some entries failed to resolve, status is PartiallyCompleted instead.* `NoOp` - No maintenance activity was queried because none of the targetChassis had a Secure Router, or no Secure Router was claimed in Catalyst SD-WAN Manager. 
 
