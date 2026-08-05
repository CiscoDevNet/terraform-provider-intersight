---
subcategory: "connectorpack"
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_upgrade_impact"
description: |-
        The UpgradeImpact object is designed to assess and identify the list of connector packs to be installed or upgraded on UCS Director. It identifies the requisite connector packs for installation during the next upgrade cycle, ensuring a smooth transition and enhanced functionality.
        #### Purpose
        UpgradeImpact object plays a crucial role in planning and executing upgrades for UCS Director, ensuring that all necessary connector packs are identified and managed effectively. It evaluates eligibility for upgrades and provides insights into the upgrade readiness of UCS Director.
        #### Key Concepts
        - **Upgrade Planning:** Facilitates strategic planning for upgrades by identifying necessary connector packs and evaluating upgrade readiness.
        - **Eligibility Assessment:** Assesses UCS Director's eligibility for upgrades, ensuring that all criteria for successful upgrade execution are met.
        - **Download Management:** Manages the download status of connector packs required for upgrades, ensuring readiness for the next iteration.
        - **Relationship Management:** Establishes relationship within connector packs of a UCS Director instance to ensure accurate upgrade of connector packs.

---

# Data Source: intersight_connectorpack_upgrade_impact
The UpgradeImpact object is designed to assess and identify the list of connector packs to be installed or upgraded on UCS Director. It identifies the requisite connector packs for installation during the next upgrade cycle, ensuring a smooth transition and enhanced functionality.
#### Purpose
UpgradeImpact object plays a crucial role in planning and executing upgrades for UCS Director, ensuring that all necessary connector packs are identified and managed effectively. It evaluates eligibility for upgrades and provides insights into the upgrade readiness of UCS Director.
#### Key Concepts
- **Upgrade Planning:** Facilitates strategic planning for upgrades by identifying necessary connector packs and evaluating upgrade readiness.
- **Eligibility Assessment:** Assesses UCS Director's eligibility for upgrades, ensuring that all criteria for successful upgrade execution are met.
- **Download Management:** Manages the download status of connector packs required for upgrades, ensuring readiness for the next iteration.
- **Relationship Management:** Establishes relationship within connector packs of a UCS Director instance to ensure accurate upgrade of connector packs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_connectorpack_upgrade_impact.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_eligible_for_upgrade`:(bool) States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false. 
* `is_update_downloaded`:(bool) States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
