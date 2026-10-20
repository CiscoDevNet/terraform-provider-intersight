---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_dest_remote"
description: |-
        Configures Remote Destination details for a given Remote ERSPAN session.

---

# Data Source: intersight_fabric_span_dest_remote
Configures Remote Destination details for a given Remote ERSPAN session.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_dest_remote.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `destination_ip`:(string) IP address of the remote ERSPAN destination. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `erspan_id`:(int) ERSPAN ID identifies the ERSPAN session on the Fabric Interconnect. 
* `ip_dscp`:(int) Differentiated Services Code Point (DSCP) value for the ERSPAN tunnel IP packets. 
* `ip_ttl`:(int) Time To Live (TTL) value for the ERSPAN tunnel IP packets. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Enables or disables packet truncation for the ERSPAN session. 
* `packet_truncation`:(bool) Enables or disables packet truncation for the ERSPAN session. 
* `session_id`:(int) Session ID identifies the SPAN session on the Fabric Interconnect. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) Switch ID (A or B) corresponding to network element. 
 
