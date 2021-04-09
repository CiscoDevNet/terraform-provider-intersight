---
subcategory: "connectorpack"
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_connector_pack_upgrade"
description: |-
  Used to download or install connector packs on the target device.
---

# Data Source: intersight_connectorpack_connector_pack_upgrade
Used to download or install connector packs on the target device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_connectorpack_connector_pack_upgrade.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 