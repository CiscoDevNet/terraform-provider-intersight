---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_system_status"
description: |-
        The SystemStatus object provides the overall health and operational status of an Intersight Appliance. It aggregates the results of multiple system-level status checks into a single high-level operational state, offering an at-a-glance indicator of whether the appliance is operating normally or requires attention.
        #### Purpose
        SystemStatus enables administrators to monitor the appliance's overall condition through a single, read-only status object that summarizes health checks into an operational status.
        #### Key Concepts
        - **Aggregated Health State:** `operationalStatus` summarizes underlying checks into a high-level operational condition, such as Operational, AttentionNeeded, or Impaired, depending on the model's OperationalStatus semantics.
        - **Read-Only Monitoring Surface:** The appliance platform exposes this value through READ for operational visibility, and it produces and updates the value automatically.
        - **Permission Inheritance:** Permissions inherited from `systemInfo` and `registeredDevice` govern access, aligning visibility with the appliance identity and its registration context.
        - **Privileged Access:** READ access is limited to the System Administrator and Account Administrator roles.

---

# Data Source: intersight_appliance_system_status
The SystemStatus object provides the overall health and operational status of an Intersight Appliance. It aggregates the results of multiple system-level status checks into a single high-level operational state, offering an at-a-glance indicator of whether the appliance is operating normally or requires attention.
#### Purpose
SystemStatus enables administrators to monitor the appliance's overall condition through a single, read-only status object that summarizes health checks into an operational status.
#### Key Concepts
- **Aggregated Health State:** `operationalStatus` summarizes underlying checks into a high-level operational condition, such as Operational, AttentionNeeded, or Impaired, depending on the model's OperationalStatus semantics.
- **Read-Only Monitoring Surface:** The appliance platform exposes this value through READ for operational visibility, and it produces and updates the value automatically.
- **Permission Inheritance:** Permissions inherited from `systemInfo` and `registeredDevice` govern access, aligning visibility with the appliance identity and its registration context.
- **Privileged Access:** READ access is limited to the System Administrator and Account Administrator roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_system_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operational_status`:(string) Operational status of the Intersight Appliance.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
