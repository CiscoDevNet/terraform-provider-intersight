---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_initiator_group"
description: |-
  NetApp Initiator Group specifies host access to LUNs on the storage system.
---

# Data Source: intersight_storage_net_app_initiator_group
NetApp Initiator Group specifies host access to LUNs on the storage system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_initiator_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the host. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `os_type`:(string) Operating system running on the host. 
* `protocol`:(string) Initiator group protocol.* `FCP` - Fibre channel initiator type which contains WWN of an HBA on the host.* `iSCSI` - An iSCSI initiator type used by the host.* `mixed` - For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the LUN. 
 