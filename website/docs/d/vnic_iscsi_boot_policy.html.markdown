---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_boot_policy"
description: |-
        The IscsiBootPolicy object encapsulates the configuration required to enable servers to boot from remote iSCSI targets over the network.
        #### Purpose
        IscsiBootPolicy automates and standardizes the process of configuring iSCSI boot, including authentication, target selection, and IP address assignment, to ensure reliable and secure remote booting.
        #### Key Concepts
        - **Remote Boot Enablement:** Supports network-based operating system boot via iSCSI.
        - **Comprehensive Configuration:** Manages authentication, IP assignment, and target discovery.
        - **Policy Integration:** Links to adapter and IP pool policies for end-to-end automation.

---

# Data Source: intersight_vnic_iscsi_boot_policy
The IscsiBootPolicy object encapsulates the configuration required to enable servers to boot from remote iSCSI targets over the network. 
#### Purpose
IscsiBootPolicy automates and standardizes the process of configuring iSCSI boot, including authentication, target selection, and IP address assignment, to ensure reliable and secure remote booting.
#### Key Concepts
- **Remote Boot Enablement:** Supports network-based operating system boot via iSCSI.
- **Comprehensive Configuration:** Manages authentication, IP assignment, and target discovery.
- **Policy Integration:** Links to adapter and IP pool policies for end-to-end automation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_iscsi_boot_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_targetvendor_name`:(string) Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 64 characters. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `initiator_ip_source`:(string) Source Type of Initiator IP Address - Auto/Static/Pool.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `initiator_static_ip_v4_address`:(string) Static IPv4 address provided for iSCSI Initiator. 
* `initiator_static_ip_v6_address`:(string) Static IPv6 address provided for iSCSI Initiator. 
* `iscsi_ip_type`:(string) Type of the IP address requested for iSCSI vNIC - IPv4/IPv6.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_source_type`:(string) Source Type of Targets - Auto/Static.* `Static` - Type indicates that static target interface is assigned to iSCSI boot.* `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot. 
 
