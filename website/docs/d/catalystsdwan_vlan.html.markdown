---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_vlan"
description: |-
        Details for the WAN Edge Device VLAN entities.

---

# Data Source: intersight_catalystsdwan_vlan
Details for the WAN Edge Device VLAN entities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_vlan.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) The administrative state of the WAN Edge Device.* `Unknown` - Administrative state is unknown.* `Up` - Administrative state is up.* `Down` - Administrative state is down. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) The distinguished name of the VLAN. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `system_ip`:(string) The system IP address of the WAN Edge Device. 
* `uuid`:(string) The UUID of the WAN Edge Device to which this VLAN belongs. 
* `vlan_id`:(int) The VLAN number configured on the WAN Edge Device. 
* `vlan_state`:(string) The state of the VLAN on the WAN Edge Device.* `Unknown` - VLAN state is unknown on the Catalyst SDWAN device.* `Active` - VLAN state is active on the Catalyst SDWAN device.* `Inactive` - VLAN state is inactive on the Catalyst SDWAN device.* `Suspended` - VLAN state is suspended on the Catalyst SDWAN device. 
 
