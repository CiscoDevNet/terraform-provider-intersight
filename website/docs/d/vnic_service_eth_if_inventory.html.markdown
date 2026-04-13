---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_service_eth_if_inventory"
description: |-
        Service Virtual Ethernet Interface is created  on each adapter to export netflow records to the collector via the Fabric Interconnect.

---

# Data Source: intersight_vnic_service_eth_if_inventory
Service Virtual Ethernet Interface is created  on each adapter to export netflow records to the collector via the Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_service_eth_if_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_id`:(string) Adapter Information for server. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the service virtual ethernet interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `template_sync_status`:(string) The sync status of the current MO wrt the attached Template MO.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `vif_id`:(int) The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for Fabric Interconnect attached servers where a vethernet is created on the switch for every vNIC. 
 
