---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_host_interface"
description: |-
  A HyperFlex Application Platform compute host interface entity that can be connected by HxapHostVswitch.
---

# Data Source: intersight_hyperflex_hxap_host_interface
A HyperFlex Application Platform compute host interface entity that can be connected by HxapHostVswitch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_host_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name`:(string) The UUID of the host to which this interface belongs to. 
* `host_uuid`:(string) The UUID of the host to which this interface belongs to. 
* `if_type`:(string) A hint of the interface type, such as \ ovs-bond\ , \ device\ , \ openvswitch\ . 
* `link_state`:(string) Link state information such as \ up\ , \ down\ .* `unknown` - The interface line is unknown.* `up` - The interface line is up.* `down` - The interface line is down.* `degraded` - For a bond/team interface, not all member interface is up. 
* `mac`:(string) The MAC address of the interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The MTU size of the interface. 
* `name`:(string) The name of the host to which this interface belongs to. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vlans`:(string) A list of vlans allowed on this interface. 
 