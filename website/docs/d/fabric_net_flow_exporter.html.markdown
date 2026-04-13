---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_net_flow_exporter"
description: |-
        Netflow Exporter configuration, it exports netflow information in the form of flow record to netflow collector.

---

# Data Source: intersight_fabric_net_flow_exporter
Netflow Exporter configuration, it exports netflow information in the form of flow record to netflow collector.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_net_flow_exporter.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `destination`:(string) Netflow collector IP address, The Netflow collector receives flow records from one or more exporters.  It processes the received export packets and stores the flow record information. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dscp`:(int) DSCP value for export packets to ensure they receive proper QoS treatment. By default, NetFlow export packets may use the default DSCP value (usually 0, equivalent to Best Effort). 
* `gateway_ip`:(string) Gateway IP address for the export interface network. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of netflow exporter. Must be a maximum of 31 characters, without spacing. 
* `option_exporter_stats_timeout`:(int) The time interval, in seconds, during which a NetFlow collector maintains an option template after it has been received from an exporter. An Option Template Record is a special type of template in NetFlow used to export metadata or control information, rather than flow data such as sampling parameters or exporter statistics. 
* `option_interface_table_timeout`:(int) The time interval, in seconds, during which a NetFlow collector maintains an interface option template after it has been received from an exporter. The optionInterfaceTable refers to an option data record exported by NetFlow exporters that provides metadata about network interfaces such as interface names, types, and speeds. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `template_data_timeout`:(int) The time interval, in seconds, during which a NetFlow collector maintains a template after it has been received from an exporter. templateData refers to the actual flow record data that is exported from a exporter to a collector, using a previously defined template. The template specifies the structure and the templateData provides the corresponding values for those fields for each flow. 
* `udp_port`:(int) NetFlow export packets are encapsulated within UDP datagrams for transmission to the NetFlow collector. 
* `nr_version`:(int) Version of flow record format exported in exporter packet.  The value of this field is 9 for the current version. 
 
