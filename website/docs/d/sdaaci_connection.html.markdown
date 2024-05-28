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
* `aci_l3_out`:(string) ACI L3Out Name User Input. 
* `aci_match_rule_name`:(string) Name of the Match Rule in Cisco APIC. 
* `aci_tenant`:(string) ACI tenant Name for Selected APIC Target. 
* `campus_fabric_site`:(string) Campus fabric site id in which the border node has configured. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epg`:(string) Application EPG Name of this connection. 
* `epg_subnet`:(string) EPG Subnet Ipv4Cidr which is configured on APIC. 
* `firewall_device`:(string) Device within the selected domain used to configure Firewall. 
* `firewall_domain`:(string) Domain used to configure Firewall. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_profile`:(string) L3Out Node Profile in Cisco APIC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Connection status between SDA and ACI.* `NotConnected` - Connection Status NotConnected.* `Connected` - Connection Status Connected. 
* `transit`:(string) Transit id for given border node. 
* `virtual_network`:(string) Virtual Network of this connection. 
* `vn_epg`:(string) Contains both VN and EPG of this connection. 
* `vrf`:(string) APIC Tenant VRF from APIC. 
 
