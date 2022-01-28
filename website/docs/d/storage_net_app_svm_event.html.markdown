---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_svm_event"
description: |-
  An event where the impacted resource type is a storage vm.
---

# Data Source: intersight_storage_net_app_svm_event
An event where the impacted resource type is a storage vm.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_svm_event.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cause`:(string) A message describing the cause for the event. 
* `cluster_uuid`:(string) Unique identifier of the cluster across the datacenter. 
* `create_time`:(string) The time when this managed object was created. 
* `current_state`:(string) The current state of the event.* `unknown` - Default unknown current state.* `new` - The current state of the event is new.* `acknowledged` - The current state of the event is acknowledged.* `resolved` - The current state of the event is resolved.* `obsolete` - The current state of the event is obsolete. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `duration`:(string) Time since the event was created, in ISO8601 standard. 
* `impact_area`:(string) Impact area of the event (availability, capacity, configuration, performance, protection, or security).* `unknown` - Default unknown impact area.* `availability` - The impact area of the event is availability.* `capacity` - The impact area of the event is capacity.* `configuration` - The impact area of the event is configuration.* `performance` - The impact area of the event is performance.* `protection` - The impact area of the event is protection.* `security` - The impact area of the event is security. 
* `impact_level`:(string) Impact level of the event (event, risk, incident, or upgrade).* `unknown` - Default unknown impact level.* `event` - The impact level of the event is event.* `risk` - The impact level of the event is risk.* `incident` - The impact level of the event is incident.* `upgrade` - The impact level of the event is upgrade. 
* `impact_resource_name`:(string) The full name of the source of the event. 
* `impact_resource_type`:(string) Type of resource with which the event is associated.* `unknown` - Default unknown resource type.* `aggregate` - The type of resource impacted by the event is an aggregate.* `cluster` - The type of resource impacted by the event is a cluster.* `cluster_node` - The type of resource impacted by the event is a node.* `disk` - The type of resource impacted by the event is a disk.* `fcp_lif` - The type of resource impacted by the event is a FC interface.* `fcp_port` - The type of resource impacted by the event is a FC port.* `lun` - The type of resource impacted by the event is a lun.* `network_lif` - The type of resource impacted by the event is an ethernet interface.* `network_port` - The type of resource impacted by the event is an ethernet port.* `volume` - The type of resource impacted by the event is a volume.* `vserver` - The type of resource impacted by the event is a storage VM. 
* `impact_resource_uuid`:(string) The unique identifier of the impacted resource. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the event that occurred. 
* `node_uuid`:(string) Unique identifier of the node across the cluster. 
* `severity`:(string) The severity of the event.* `unknown` - Default unknown severity.* `normal` - The severity of the event is normal.* `information` - The severity of the event is information.* `warning` - The severity of the event is warning.* `error` - The severity of the event is error.* `critical` - The severity of the event is critical. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_uuid`:(string) Unique identifier of the storage VM. 
* `uuid`:(string) Unique identifier of the event. 
 