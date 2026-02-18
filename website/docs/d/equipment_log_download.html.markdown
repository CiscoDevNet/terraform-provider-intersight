---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_log_download"
description: |-
        The LogDownload object facilitates the download of collected logs from Cisco endpoints. It supports efficient retrieval and management of log files.
        #### Purpose
        LogDownload enables users to access and download logs collected from endpoints, aiding in troubleshooting and data management. It is a key component in log retrieval processes within Cisco environments.
        #### Key Concepts
        - **Log Retrieval:** Provides mechanisms for downloading logs from endpoints, supporting efficient access to collected data.
        - **Access Control:** Establishes privilege sets for log download operations, ensuring secure management.
        - **Integration with Devices:** Connects with servers to ensure log downloads are associated correctly within the network.
        - **Security:** Ensures log downloads are performed securely, supporting compliance and data integrity.

---

# Data Source: intersight_equipment_log_download
The LogDownload object facilitates the download of collected logs from Cisco endpoints. It supports efficient retrieval and management of log files.
#### Purpose
LogDownload enables users to access and download logs collected from endpoints, aiding in troubleshooting and data management. It is a key component in log retrieval processes within Cisco environments.
#### Key Concepts
- **Log Retrieval:** Provides mechanisms for downloading logs from endpoints, supporting efficient access to collected data.
- **Access Control:** Establishes privilege sets for log download operations, ensuring secure management.
- **Integration with Devices:** Connects with servers to ensure log downloads are associated correctly within the network.
- **Security:** Ensures log downloads are performed securely, supporting compliance and data integrity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_log_download.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
