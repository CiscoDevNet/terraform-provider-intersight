---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_upgrade_support_meta"
description: |-
        The ServerUpgradeSupportMeta object provides internal metadata to map server family classifications, which are used in firmware policy enforcement.
        #### Purpose
        It categorizes server models into families, allowing firmware policies to be applied consistently across similar hardware platforms.
        #### Key Concepts
        - **Family Classification:** Maps various server models to a unified server family.
        - **Platform Mapping:** Associates server families with their target platforms for policy enforcement.

---

# Data Source: intersight_capability_server_upgrade_support_meta
The ServerUpgradeSupportMeta object provides internal metadata to map server family classifications, which are used in firmware policy enforcement.
#### Purpose
It categorizes server models into families, allowing firmware policies to be applied consistently across similar hardware platforms.
#### Key Concepts
- **Family Classification:** Maps various server models to a unified server family.
- **Platform Mapping:** Associates server families with their target platforms for policy enforcement.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_upgrade_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Verbose description regarding this group of server. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `platform`:(string) The target platform for which the mapping is defined. 
* `server_family`:(string) Classification of a set of server models. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
