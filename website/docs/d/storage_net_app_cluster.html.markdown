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
* `default_admin_locked`:(bool) Indicates whether the default admin user is locked out. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_type`:(string) The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fips_compliant`:(bool) Indicates whether or not the software FIPS mode is enabled on the cluster. 
* `hardware_version`:(string) The hardware version of the device. 
* `insecure_ciphers`:(int) Number of SVMs on the cluster that use insecure ciphers. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `key`:(string) Unique identifier of NetApp Cluster across data center. 
* `location`:(string) Location of the storage controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `rsh_enabled`:(bool) Indicates whether or not rsh is enabled on the cluster. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `telnet_enabled`:(bool) Indicates whether or not telnet is enabled on the cluster. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Current running software version of the device. 
* `version_generation`:(int) The generation portion of the version. 
* `version_major`:(int) The major portion of the version. 
* `version_minor`:(int) The minor portion of the version. 
 
