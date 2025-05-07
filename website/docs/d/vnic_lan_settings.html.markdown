---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_lan_settings"
description: |-
        An internal MO to track the LAN settings that are applied to a server profile (like iQNLease).

---

# Data Source: intersight_vnic_lan_settings
An internal MO to track the LAN settings that are applied to a server profile (like iQNLease).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_lan_settings.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `iqn_name`:(string) Qualified Name (IQN) assigned to iSCSI vNICs in a Fabric Interconnect domain. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
