---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_iom_upgrade_support_meta"
description: |-
        The IomUpgradeSupportMeta object provides internal metadata to enable I/O Module (IOM) firmware upgrade decision-making.
        #### Purpose
        It helps the system identify IOM series and models eligible for firmware upgrades, including whether they support direct upgrade requests via a Device Connector.
        #### Key Concepts
        - **Direct Upgrade Capability:** Indicates if an IOM model supports direct upgrade requests.
        - **Series Management:** Groups IOM models into series for consistent firmware handling.

---

# Data Source: intersight_capability_iom_upgrade_support_meta
The IomUpgradeSupportMeta object provides internal metadata to enable I/O Module (IOM) firmware upgrade decision-making.
#### Purpose
It helps the system identify IOM series and models eligible for firmware upgrades, including whether they support direct upgrade requests via a Device Connector.
#### Key Concepts
- **Direct Upgrade Capability:** Indicates if an IOM model supports direct upgrade requests.
- **Series Management:** Groups IOM models into series for consistent firmware handling.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_iom_upgrade_support_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Information related to the list of IOMs. Also provides additional information such as hardware name. 
* `direct_upgrade`:(bool) Indicates if the IOM models have a Device Connector, which in turn allows direct upgrade requests to be sent to the IOM DC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `series_id`:(string) Series names of IOMs which will be supported in the firmware operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
