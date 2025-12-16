---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_system_op_status"
description: |-
        The SystemOpStatus object represents the status of the Intersight Appliance within a system. This provides a comprehensive view of the appliance's operational state, including the health and performance of its nodes and applications.
        #### Purpose
        SystemOpStatus serves as the primary indicator of the appliance's operational health. It encompasses the results of status checks across various components, determining the overall status as Operational, AttentionNeeded, or Impaired based on these checks.
        #### Key Concepts
        - **Operational Status:** Defines the health of the appliance, influenced by checks on nodes, applications, and system components.
        - **Node and Application Monitoring:** Tracks the number of nodes and applications that require attention or are impaired.
        - **Alert System:** Utilizes certificate expiry dates to raise alarms, ensuring timely updates and security compliance.
        - **Relationships:** Links to system information, applications, groups, and accounts to provide a holistic overview.

---

# Data Source: intersight_appliance_system_op_status
The SystemOpStatus object represents the status of the Intersight Appliance within a system. This provides a comprehensive view of the appliance's operational state, including the health and performance of its nodes and applications.
#### Purpose
SystemOpStatus serves as the primary indicator of the appliance's operational health. It encompasses the results of status checks across various components, determining the overall status as Operational, AttentionNeeded, or Impaired based on these checks.
#### Key Concepts
- **Operational Status:** Defines the health of the appliance, influenced by checks on nodes, applications, and system components.
- **Node and Application Monitoring:** Tracks the number of nodes and applications that require attention or are impaired.
- **Alert System:** Utilizes certificate expiry dates to raise alarms, ensuring timely updates and security compliance.
- **Relationships:** Links to system information, applications, groups, and accounts to provide a holistic overview.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_system_op_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operational_status`:(string) Operational status of the Intersight Appliance.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
