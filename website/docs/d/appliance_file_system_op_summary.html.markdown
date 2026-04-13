---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_system_op_summary"
description: |-
        The FileSystemOpSummary object aggregates alarm summaries for file systems on an Intersight Appliance node. It serves as an essential tool for maintaining system efficiency and reliability by summarizing critical health and capacity metrics.
        #### Purpose
        FileSystemOpSummary provides a consolidated view of file system health, supporting capacity management and ensuring that file systems remain operational and efficient. It highlights alarm counts per severity, aiding in proactive system management.
        #### Key Concepts
        - **Alarm Aggregation** Consolidates alarm data to provide a health score, facilitating quick identification of critical issues.
        - **Capacity Insights** Offers detailed views into file system usage and capacity, preventing overutilization and ensuring optimal performance.
        - **System Integration** Links file system status with broader system health metrics, contributing to comprehensive system management.

---

# Data Source: intersight_appliance_file_system_op_summary
The FileSystemOpSummary object aggregates alarm summaries for file systems on an Intersight Appliance node. It serves as an essential tool for maintaining system efficiency and reliability by summarizing critical health and capacity metrics.
#### Purpose
FileSystemOpSummary provides a consolidated view of file system health, supporting capacity management and ensuring that file systems remain operational and efficient. It highlights alarm counts per severity, aiding in proactive system management.
#### Key Concepts
- **Alarm Aggregation** Consolidates alarm data to provide a health score, facilitating quick identification of critical issues.
- **Capacity Insights** Offers detailed views into file system usage and capacity, preventing overutilization and ensuring optimal performance.
- **System Integration** Links file system status with broader system health metrics, contributing to comprehensive system management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_file_system_op_summary.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
