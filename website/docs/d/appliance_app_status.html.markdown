---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_app_status"
description: |-
        The AppStatus object represents the operational status of individual applications running on an Intersight Appliance. It aggregates application-level health checks into a single operational status and exposes replica and availability signals so operators can quickly assess whether an application is functioning and stable.
        #### Purpose
        AppStatus provides read-only visibility into per-application health and readiness on the appliance, enabling troubleshooting and operational monitoring of the appliance software stack.
        #### Key Concepts
        - **Application-Level Health Aggregation:** `operationalStatus` summarizes status checks into high-level states, such as Operational, AttentionNeeded, or Impaired, based on underlying check results.
        - **Status Check Details:** `statusChecks` provides the set of checks performed from the application's perspective, primarily readiness and function checks.
        - **Availability and Scale Signals:** `replicaCount`, `runningCount`, and `readyCount` expose deployment and readiness posture for the application.
        - **Stability Indicators:** Restart counters, including `restartCountTotal`, `restartCount5Mins`, `restartCount1Hour`, and `restartCount24Hours`, help detect crash loops and instability trends.
        - **API Health Visibility:** `apiStatuses` records the status of specific GET APIs, including name, response code, elapsed time, and failure reason, to aid diagnosis of control-plane and API issues.
        - **System-Scoped Governance:** Permissions inherited through `inheritpermissions: [systemStatus]` govern access, aligning visibility to the owning appliance system status context.

---

# Data Source: intersight_appliance_app_status
The AppStatus object represents the operational status of individual applications running on an Intersight Appliance. It aggregates application-level health checks into a single operational status and exposes replica and availability signals so operators can quickly assess whether an application is functioning and stable.
#### Purpose
AppStatus provides read-only visibility into per-application health and readiness on the appliance, enabling troubleshooting and operational monitoring of the appliance software stack.
#### Key Concepts
- **Application-Level Health Aggregation:** `operationalStatus` summarizes status checks into high-level states, such as Operational, AttentionNeeded, or Impaired, based on underlying check results.
- **Status Check Details:** `statusChecks` provides the set of checks performed from the application's perspective, primarily readiness and function checks.
- **Availability and Scale Signals:** `replicaCount`, `runningCount`, and `readyCount` expose deployment and readiness posture for the application.
- **Stability Indicators:** Restart counters, including `restartCountTotal`, `restartCount5Mins`, `restartCount1Hour`, and `restartCount24Hours`, help detect crash loops and instability trends.
- **API Health Visibility:** `apiStatuses` records the status of specific GET APIs, including name, response code, elapsed time, and failure reason, to aid diagnosis of control-plane and API issues.
- **System-Scoped Governance:** Permissions inherited through `inheritpermissions: [systemStatus]` govern access, aligning visibility to the owning appliance system status context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_app_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `app_label`:(string) Unique label to identify the application. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operational_status`:(string) Operational status of the application.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `ready_count`:(int) Number of replicas ready.  The number of instances ofthe application currently ready to perform its intendedfunctions. 
* `replica_count`:(int) Number of replicas provisioned. The number of instancesof the application provisioned to run on the Intersight appliance. 
* `restart_count1_hour`:(int) Number of instance restarts in the last hour. 
* `restart_count24_hours`:(int) Number of instance restarts in the last 24 hours. 
* `restart_count5_mins`:(int) Number of instance restarts in the last 5 minutes. 
* `restart_count_total`:(int) Total number of restarts since last deployment. 
* `running_count`:(int) Number of replicas running. The number of instances ofthe application currently running. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
