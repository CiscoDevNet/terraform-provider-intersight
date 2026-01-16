---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade"
description: |-
        The Upgrade object facilitates the firmware upgrade process for rack and blade servers, providing a structured approach to downloading and applying firmware images from various sources, including cloud repositories and network shares.
        #### Purpose
        The Upgrade object is central to orchestrating firmware upgrades, offering a standardized method to enhance server functionality and security through timely firmware updates. It supports both direct download and network share upgrade methods.
        #### Key Concepts
        - **Flexible Upgrade Options:** Supports direct download from Cisco repositories and network share-based upgrades, accommodating diverse infrastructure setups and preferences.
        - **Policy Integration:** Seamlessly interacts with firmware policies, ensuring compliance and alignment with organizational upgrade strategies.
        - **Comprehensive Management:** Facilitates the entire upgrade lifecycle, from image selection and download to installation and verification, ensuring a reliable upgrade experience.
        - **Secure and Controlled:** Enforces strict access controls and licensing requirements to safeguard the upgrade process, maintaining system security and compliance.

---

# Data Source: intersight_firmware_upgrade
The Upgrade object facilitates the firmware upgrade process for rack and blade servers, providing a structured approach to downloading and applying firmware images from various sources, including cloud repositories and network shares.
#### Purpose
The Upgrade object is central to orchestrating firmware upgrades, offering a standardized method to enhance server functionality and security through timely firmware updates. It supports both direct download and network share upgrade methods.
#### Key Concepts
- **Flexible Upgrade Options:** Supports direct download from Cisco repositories and network share-based upgrades, accommodating diverse infrastructure setups and preferences.
- **Policy Integration:** Seamlessly interacts with firmware policies, ensuring compliance and alignment with organizational upgrade strategies.
- **Comprehensive Management:** Facilitates the entire upgrade lifecycle, from image selection and download to installation and verification, ensuring a reliable upgrade experience.
- **Secure and Controlled:** Enforces strict access controls and licensing requirements to safeguard the upgrade process, maintaining system security and compliance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_upgrade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_name`:(string) Name of the server on which the firmware upgrade operation is performed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `upgrade_trigger_method`:(string) The source that triggered the upgrade. Either via profile or traditional way.* `none` - Upgrade is invoked within the service.* `profileTrigger` - Upgrade is invoked from a profile deployment. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 
 
