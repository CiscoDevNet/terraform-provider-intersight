---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_cluster_node"
description: |-
        A node managed by a nexus dashboard instance.

---

# Data Source: intersight_niatelemetry_cluster_node
A node managed by a nexus dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_cluster_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `boot_time`:(string) The boot time of the node in the cluster. 
* `cluster_name`:(string) The name of the cluster this node belongs to. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) The firmware version of the node in the cluster. 
* `management_ip`:(string) The management IP address of the node in the cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the node in the cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Returns the name of the node. 
* `platform_type`:(string) The platform type of the node in the cluster. 
* `role`:(string) Role of the node in the cluster. 
* `serial_number`:(string) Serial number of the node in the cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the node in the cluster. 
 
