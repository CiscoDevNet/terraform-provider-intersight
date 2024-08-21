---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_source_vnic_eth_if"
description: |-
        Configures a VNIC as the SPAN Source for a given SPAN session. For failover VNICs, VNIC must be added as a
        source to a SPAN session on the peer Fabric Interconnect to ensure that traffic is mirrored on both
        Fabric Interconnects.

---

# Data Source: intersight_fabric_span_source_vnic_eth_if
Configures a VNIC as the SPAN Source for a given SPAN session. For failover VNICs, VNIC must be added as a
source to a SPAN session on the peer Fabric Interconnect to ensure that traffic is mirrored on both
Fabric Interconnects.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_source_vnic_eth_if.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `direction`:(string) Direction of the source SPAN.* `Receive` - SPAN incoming traffic on the SPAN source interface.* `Transmit` - SPAN outgoing traffic on the SPAN source interface.* `Both` - SPAN incoming and outgoing traffic on the SPAN source interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the VNIC referenced by vnic relationship. Vnic name is not updated if it gets updated by the user. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vif_id`:(int) VIF ID of the VNIC placed on Fabric Interconnect associated to the SPAN session. 
 
