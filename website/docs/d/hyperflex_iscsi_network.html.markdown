---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_iscsi_network"
description: |-
        The HyperFlex iSCSI Network configuration.
        Contains detailed information about the iSCSI network which includes subnet, gateway, Virtual local area network (VLAN) name,
        VLAN identity, maximum transmission unit (MTU), ranges of the IP addresses belonging to the HyperFlex iSCSI network.

---

# Data Source: intersight_hyperflex_iscsi_network
The HyperFlex iSCSI Network configuration.
Contains detailed information about the iSCSI network which includes subnet, gateway, Virtual local area network (VLAN) name,
VLAN identity, maximum transmission unit (MTU), ranges of the IP addresses belonging to the HyperFlex iSCSI network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_iscsi_network.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gateway`:(string) The gateway of the HyperFlex iSCSI network. 
* `inventory_source`:(string) Source of this inventory object.* `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable.* `ONLINE` - The source of the HyperFlex inventory is online.* `OFFLINE` - The source of the HyperFlex inventory is offline. 
* `iscsi_cip`:(string) An IP within the iSCSI IP Address Range which is CIP for iSCSI network. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(string) The maximum transmission unit of the HyperFlex iSCSI network.* `UNKNOWN` - The maximum transmission unit of the HyperFlex iSCSI network is unknown.* `MTU_1500` - The maximum transmission unit of the HyperFlex iSCSI network is 1500 bytes.* `MTU_9000` - The maximum transmission unit of the HyperFlex iSCSI network is 9000 bytes. 
* `name`:(string) Name of the HyperFlex iSCSI network configuration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subnet`:(string) Subnet of the HyperFlex iSCSI network. Subnet is in a.b.c.d/e notation. 
* `ucs_host`:(string) UCS Manager Host IP or FQDN. 
* `uuid`:(string) UUID of the HyperFlex iSCSI network configuration. 
* `nr_version`:(int) Version of Network configuration in Inventory. 
* `vlan_name`:(string) The Virtual local area network (VLAN) name of the HyperFlex iSCSI network. 
* `vlanid`:(int) The VLAN ID of the HyperFlex iSCSI network. 
 
