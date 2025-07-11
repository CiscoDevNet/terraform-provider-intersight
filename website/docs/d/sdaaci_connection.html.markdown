---
subcategory: "sdaaci"
layout: "intersight"
page_title: "Intersight: intersight_sdaaci_connection"
description: |-
        SDA-ACI direct-connect connection.

---

# Data Source: intersight_sdaaci_connection
SDA-ACI direct-connect connection.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdaaci_connection.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aci_l3_out`:(string) User input for ACI L3Out name. 
* `aci_match_rule_name`:(string) Match rule name in Cisco APIC. 
* `aci_tenant`:(string) ACI tenant name for the selected APIC target. 
* `campus_fabric_site`:(string) Campus fabric site Id where the border node is configured. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epg`:(string) Application EPG name for this connection. 
* `epg_subnet`:(string) EPG Subnet Ipv4Cidr configured on APIC. 
* `firewall_device`:(string) Device within the selected domain used for firewall configuration. 
* `firewall_domain`:(string) Domain used for firewall configuration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_profile`:(string) L3Out node profile in Cisco APIC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Connection status between SDA and ACI.* `NotConnected` - Status of the connection:Not connected.* `Connected` - Status of the connection:Connected. 
* `transit`:(string) Transit Id for the given border node. 
* `virtual_network`:(string) Virtual Network for this connection. 
* `vn_epg`:(string) VN and EPG information for this connection. 
* `vrf`:(string) Tenant VRF in Cisco APIC. 
 
