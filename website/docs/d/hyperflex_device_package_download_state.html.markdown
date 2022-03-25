---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_device_package_download_state"
description: |-
        HyperFlex Device Package Download State.

---

# Data Source: intersight_hyperflex_device_package_download_state
HyperFlex Device Package Download State.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_device_package_download_state.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `checksum`:(string) Checksum of HyperFlex health check Debian package installed on the HyperFlex Device. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hx_device_name`:(string) HyperFlex Device Name for which the package download state is tracked. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timestamp`:(string) Timestamp of the last health check Debian package installation on the HyperFlex Device. 
 
