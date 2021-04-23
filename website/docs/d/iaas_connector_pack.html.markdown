---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_connector_pack"
description: |-
  Describes about all the connector pack versions running currently in UCSD.
---

# Data Source: intersight_iaas_connector_pack
Describes about all the connector pack versions running currently in UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_connector_pack.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `complete_version`:(string) Complete version of the connector pack including build number. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `downloaded_version`:(string) Version of the connector pack that is last downloaded successfully to UCSD. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the connector pack running on the UCSD. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) State of the connector pack whether it is enabled or disabled. 
* `nr_version`:(string) Version of the connector pack. 
 