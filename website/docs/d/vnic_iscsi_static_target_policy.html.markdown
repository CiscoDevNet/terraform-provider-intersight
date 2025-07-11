---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_static_target_policy"
description: |-
        Configuration parameters that defines the reachability of iSCSI Target portal.

---

# Data Source: intersight_vnic_iscsi_static_target_policy
Configuration parameters that defines the reachability of iSCSI Target portal.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_iscsi_static_target_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_address`:(string) The IP address assigned to the iSCSI target. 
* `iscsi_ip_type`:(string) Type of the IP address requested for iSCSI vNIC - IPv4/IPv6.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `port`:(int) The port associated with the iSCSI target. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_name`:(string) Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target. 
 
