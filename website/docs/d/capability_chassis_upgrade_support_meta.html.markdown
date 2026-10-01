---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_chassis_upgrade_support_meta"
description: |-
        The ChassisUpgradeSupportMeta object provides internal metadata to enable chassis firmware upgrade decision-making.
        #### Purpose
        It allows the system to determine the eligibility of chassis models for firmware upgrades, including specific support for power supply (PSU) and XFM components.
        #### Key Concepts
        - **Chassis Classification:** Groups chassis models into series for streamlined upgrade management.
        - **Component Support:** Tracks support for specific PSU and XFM models within a chassis series.
        - **HSU Integration:** Indicates if server adapters within the chassis are upgraded via the Host Upgrade (HSU) process.

---

# Data Source: intersight_capability_chassis_upgrade_support_meta
The ChassisUpgradeSupportMeta object provides internal metadata to enable chassis firmware upgrade decision-making.
#### Purpose
It allows the system to determine the eligibility of chassis models for firmware upgrades, including specific support for power supply (PSU) and XFM components.
#### Key Concepts
- **Chassis Classification:** Groups chassis models into series for streamlined upgrade management.
- **Component Support:** Tracks support for specific PSU and XFM models within a chassis series.
- **HSU Integration:** Indicates if server adapters within the chassis are upgraded via the Host Upgrade (HSU) process.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_chassis_upgrade_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapters_upgraded_via_hsu`:(bool) If enabled, indicates that adapters in servers within this chassis would be upgraded by HSU. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Verbose description regarding this group of chassis. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series_id`:(string) Classification of a set of chassis models. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
