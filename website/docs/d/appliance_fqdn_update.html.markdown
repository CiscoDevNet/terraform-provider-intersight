---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_fqdn_update"
description: |-
        FqdnUpdate provides configuration parameters for updating Appliance FQDN (fully qualified domain name) on claimed Endpoints. Appliance with new FQDN should be restored using the backup. The FQDN must be unique within the network. For multinode the input FQDN should be only one management node FQDN not the metrics node FQDN. If the restored appliance is going to be a HA cluster, other management node FQDNs will be updated during add node and cluster expansion operations.

---

# Data Source: intersight_appliance_fqdn_update
FqdnUpdate provides configuration parameters for updating Appliance FQDN (fully qualified domain name) on claimed Endpoints. Appliance with new FQDN should be restored using the backup. The FQDN must be unique within the network. For multinode the input FQDN should be only one management node FQDN not the metrics node FQDN. If the restored appliance is going to be a HA cluster, other management node FQDNs will be updated during add node and cluster expansion operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_fqdn_update.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) End date of the appliance FQDN change. 
* `fqdn`:(string) The FQDN (fully qualified domain name) of the destination appliance for target migration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the appliance FQDN change. 
* `status`:(string) The status of the FQDN update operation.* `NotStarted` - Appliance FQDN update operation has not started.* `Started` - Appliance FQDN update operation has started.* `Failed` - Appliance FQDN update operation has failed.* `Completed` - Appliance FQDN update operation has completed. 
 
