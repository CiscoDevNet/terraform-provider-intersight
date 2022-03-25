---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_host_lun"
description: |-
        A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.

---

# Data Source: intersight_storage_hitachi_host_lun
A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_host_lun.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(string) Port ID of the Hitachi host lun. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 
