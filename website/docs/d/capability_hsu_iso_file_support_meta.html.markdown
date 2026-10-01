---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_hsu_iso_file_support_meta"
description: |-
        The HsuIsoFileSupportMeta object provides metadata to verify if the HSU (Host Upgrade) process supports accepting an ISO file path in an upgrade request.
        #### Purpose
        It allows the system to determine if a specific server series supports HSU-based ISO upgrades, providing the necessary symbolic links and version requirements.
        #### Key Concepts
        - **ISO Support Validation:** Checks if the HSU capability is present based on firmware versions.
        - **File Referencing:** Maps series to specific ISO file symbolic links.
        - **Model-Specific Constraints:** Provides granular control over ISO support based on model and version combinations.

---

# Data Source: intersight_capability_hsu_iso_file_support_meta
The HsuIsoFileSupportMeta object provides metadata to verify if the HSU (Host Upgrade) process supports accepting an ISO file path in an upgrade request.
#### Purpose
It allows the system to determine if a specific server series supports HSU-based ISO upgrades, providing the necessary symbolic links and version requirements.
#### Key Concepts
- **ISO Support Validation:** Checks if the HSU capability is present based on firmware versions.
- **File Referencing:** Maps series to specific ISO file symbolic links.
- **Model-Specific Constraints:** Provides granular control over ISO support based on model and version combinations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_hsu_iso_file_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_file_name`:(string) Name of the symbolic link to the actual iso file in the default case. 
* `default_min_version`:(string) Firmware version from which the HSU capability is present in the default case. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series`:(string) Series name for the capability group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
