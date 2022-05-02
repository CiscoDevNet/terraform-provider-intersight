---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cluster"
description: |-
        NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.

---

# Data Source: intersight_storage_net_app_cluster
NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_health_status`:(string) The health status of the cluster. Possible states are ok, ok-with-suppressed, degraded, and unreachable.* `Unreachable` - Cluster status is unreachable.* `OK` - Cluster status is either ok or ok-with-suppressed.* `Degraded` - Cluster status is degraded. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(string) Unique identifier of NetApp Cluster across data center. 
* `location`:(string) Location of the storage controller. 
* `management_address`:(string) FQDN or IP Address of Storage Cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `telnet_enabled`:(bool) Indicates whether or not telnet is enabled on the cluster. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Current running software version of the device. 
 
