---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_adapter_policy"
description: |-
  Set of iSCSI properties that govern the host-side behavior of the adapter.
---

# Data Source: intersight_vnic_iscsi_adapter_policy
Set of iSCSI properties that govern the host-side behavior of the adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_iscsi_adapter_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connection_time_out`:(int) The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `dhcp_timeout`:(int) The number of seconds to wait before the initiator assumes that the DHCP server is unavailable. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `lun_busy_retry_count`:(int) The number of times to retry the connection in case of a failure during iSCSI LUN discovery. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 