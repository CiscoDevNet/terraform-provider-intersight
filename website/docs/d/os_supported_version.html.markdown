---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_supported_version"
description: |-
        The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
        operating system configuration will be created by providing the vendor and version. The version will be filtered
        based on vendor.

---

# Data Source: intersight_os_supported_version
The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_supported_version.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filter_options`:(string) The OsInstall Supported Operating System Version Filter Option.* `None` - No filtering is applied, allowing all available OS versions.* `SupportedInBlueprint` - Restricts the OS version specific to blueprint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `version_name`:(string) The OsInstall Supported Operating System Version Name. 
 
