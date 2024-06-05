---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_device_interface"
description: |-
        The device interface serves as the gateway between a network and external entities.

---

# Data Source: intersight_dnac_device_interface
The device interface serves as the gateway between a network and external entities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_device_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_status`:(string) Administrator status for network device interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Identity of network device. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interface_id`:(string) Unique identity of network device interface. 
* `interface_name`:(string) The name for the network device interface. 
* `interface_type`:(string) Interface type for network device interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_type`:(string) Port type for network device interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
