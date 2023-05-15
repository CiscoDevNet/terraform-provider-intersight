---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_node"
description: |-
        NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.

---

# Data Source: intersight_storage_net_app_node
NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cdpd_enabled`:(string) Storage node option for cdpd state.* `unknown` - The cdpd option is unknown on the node.* `on` - The cdpd option is enabled on the node.* `off` - The cdpd option is disabled on the node. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health`:(bool) The health of the NetApp Node. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `key`:(string) Unique identifier of NetApp Node across data center. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Storage array controller name. 
* `operational_mode`:(string) Controller running mode, Primary or Secondary.* `Unknown` - Component operational state is unknown.* `Primary` - Component operates in primary mode and accepts workloads.* `Secondary` - Component is running as a secondary or standby mode.* `Maintenance` - Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the NetApp Node. 
* `status`:(string) Status of the storage controller.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `systemid`:(string) The system id of the NetApp Node. 
* `uuid`:(string) Universally unique identifier of NetApp Node. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Software version running on a storage controller. 
 
