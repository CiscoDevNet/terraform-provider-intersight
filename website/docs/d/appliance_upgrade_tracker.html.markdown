---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_upgrade_tracker"
description: |-
        The UpgradeTracker object is designed to monitor the history of software upgrades within the Intersight Appliance. It acts as a singleton object to facilitate alarm conditions and raise alarms in the event of failed upgrades.
        #### Purpose
        UpgradeTracker serves as a historical record of upgrade processes, focusing on the outcome and status of the last upgrade. It supports diagnostic and alert functions by providing clear indicators of upgrade failures.
        #### Key Concepts
        - **Failure Indication:** Tracks whether the last software upgrade was unsuccessful, serving as a basis for raising alarms and initiating corrective actions.
        - **Singleton Nature:** Operates as a single instance within the system, ensuring centralized management and consistency in tracking upgrade history.
        - **Alarm Integration:** Interfaces with alarm conditions to facilitate timely alerts and responses to upgrade issues.
        - **Version Documentation:** Records the version and descriptive details of failed upgrades, aiding in troubleshooting and resolution efforts.

---

# Data Source: intersight_appliance_upgrade_tracker
The UpgradeTracker object is designed to monitor the history of software upgrades within the Intersight Appliance. It acts as a singleton object to facilitate alarm conditions and raise alarms in the event of failed upgrades.
#### Purpose
UpgradeTracker serves as a historical record of upgrade processes, focusing on the outcome and status of the last upgrade. It supports diagnostic and alert functions by providing clear indicators of upgrade failures. 
#### Key Concepts 
- **Failure Indication:** Tracks whether the last software upgrade was unsuccessful, serving as a basis for raising alarms and initiating corrective actions.
- **Singleton Nature:** Operates as a single instance within the system, ensuring centralized management and consistency in tracking upgrade history.
- **Alarm Integration:** Interfaces with alarm conditions to facilitate timely alerts and responses to upgrade issues.
- **Version Documentation:** Records the version and descriptive details of failed upgrades, aiding in troubleshooting and resolution efforts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_upgrade_tracker.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) A description of the upgrade or upgrade failure. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_upgrade_failed`:(bool) Indicates if the last upgrade has failed or not. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) Version of the failed upgrade. 
 
