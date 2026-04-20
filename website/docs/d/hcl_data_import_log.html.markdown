---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_data_import_log"
description: |-
        The DataImportLog object serves as a record-keeping tool for tracking the data import processes from the HCL tool. This provides transparency and  accountability in the data import workflow, enabling system administrators to monitor and manage data updates effectively.
        #### Purpose
        DataImportLog is essential for documenting the data import operations within the system. It offers detailed insights into the import process, helping administrators ensure data accuracy and consistency across the platform.
        ####Key Concepts
        - **Process Tracking:** Logs the details of each data import operation, providing a historical record for audit and review purposes.
        - **Status Documentation:** Captures the status and progress of data imports, enabling administrators to track and respond to any issues that arise.
        - **Administrative Access:** Access is limited to system administrators, ensuring that only authorized personnel can create or read DataImportLog entries.
        - **Relationship Management:** Integrates with account management systems, linking data import activities to specific accounts for enhanced oversight and control.

---

# Data Source: intersight_hcl_data_import_log
The DataImportLog object serves as a record-keeping tool for tracking the data import processes from the HCL tool. This provides transparency and  accountability in the data import workflow, enabling system administrators to monitor and manage data updates effectively.
#### Purpose
DataImportLog is essential for documenting the data import operations within the system. It offers detailed insights into the import process, helping administrators ensure data accuracy and consistency across the platform.
####Key Concepts
- **Process Tracking:** Logs the details of each data import operation, providing a historical record for audit and review purposes.
- **Status Documentation:** Captures the status and progress of data imports, enabling administrators to track and respond to any issues that arise.
- **Administrative Access:** Access is limited to system administrators, ensuring that only authorized personnel can create or read DataImportLog entries.
- **Relationship Management:** Integrates with account management systems, linking data import activities to specific accounts for enhanced oversight and control.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_data_import_log.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `check_sum`:(string) MD5 Checksum of the HCL Data file. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `import_type`:(string) Type of the import whether it is full or a delta.* `Delta` - Imports only changes from the HCL tool into Intersight.* `Full` - Deletes the current data and does an full import of the data. 
* `initiator_type`:(string) Type of the initiator whether it is manual or a automated periodic operation by system. The value will be set during DoPost to Manual.* `Auto` - Import is auto triggered during service startup or periodic poll.* `Manual` - Import is triggered externally by devops using API. 
* `last_hcl_data_modified_time`:(string) The timestamp of the last modified record in the HCL tool database. Used to query and get updated records. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the import process.* `None` - Default import status when no activity is taking place with respect to import.* `InProgress` - Data import is in progress.* `Success` - Data import is successful.* `Started` - Data import process has started.* `Failed` - Data import process has failed.* `NoChange` - There is no change in the data. 
* `status_details`:(string) More information on the status. 
 
