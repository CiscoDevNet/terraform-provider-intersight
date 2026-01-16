---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_chassis_upgrade"
description: |-
        The ChassisUpgrade object is focused on upgrading firmware for chassis systems, providing a streamlined approach to downloading and applying firmware images from cloud repositories or network shares.
        #### Purpose
        ChassisUpgrade ensures chassis systems remain secure and performant through firmware updates. This provides options for both direct download and network share–based upgrades for UCSM devices, and supports certain IMM chassis upgrades via direct download. IMM chassis upgrade updates supported PSU and XFM components when connected to the chassis.
        #### Key Concepts
        - **Chassis System Enhancement:** Designed to optimize chassis firmware upgrades, improving system reliability and efficiency through structured update processes.
        - **Exclusion Options:** Offers the ability to exclude certain components from the upgrade, providing flexibility and customization based on specific upgrade needs.
        - **Integrated Management:** Facilitates comprehensive upgrade operations, from image selection to deployment, ensuring seamless execution and management.
        - **Security and Access Control: Enforces strict access and compliance measures to protect the upgrade process, maintaining system integrity and security.

---

# Data Source: intersight_firmware_chassis_upgrade
The ChassisUpgrade object is focused on upgrading firmware for chassis systems, providing a streamlined approach to downloading and applying firmware images from cloud repositories or network shares.
#### Purpose
ChassisUpgrade ensures chassis systems remain secure and performant through firmware updates. This provides options for both direct download and network share–based upgrades for UCSM devices, and supports certain IMM chassis upgrades via direct download. IMM chassis upgrade updates supported PSU and XFM components when connected to the chassis.
#### Key Concepts
- **Chassis System Enhancement:** Designed to optimize chassis firmware upgrades, improving system reliability and efficiency through structured update processes.
- **Exclusion Options:** Offers the ability to exclude certain components from the upgrade, providing flexibility and customization based on specific upgrade needs.
- **Integrated Management:** Facilitates comprehensive upgrade operations, from image selection to deployment, ensuring seamless execution and management.
- **Security and Access Control: Enforces strict access and compliance measures to protect the upgrade process, maintaining system integrity and security.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_chassis_upgrade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chassis_name`:(string) Name of the chassis on which the firmware upgrade operation is performed. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 
* `xfm_upgrade_option`:(string) XFM upgrade option Full or Partial Disruption.* `none` - If no option is selected for exclusion.* `full-shutdown` - PSX Switch in XFM will be upgraded in single action.* `partial-shutdown` - PSX Switch in XFM will be upgraded one after other. 
 
