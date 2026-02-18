---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_end_point_log"
description: |-
        The EndPointLog object encapsulates log file information collected from Cisco endpoints. This provides insights into log collection and management processes.
        #### Purpose
        EndPointLog serves to document and manage log files collected from endpoints, supporting troubleshooting and operational insights. It aids in maintaining transparency and accountability in log management.
        #### Key Concepts
        - **Log Management:** Handles log file types and collection statuses, offering visibility into endpoint operations.
        - **Access Control:** Defines privilege sets for reading log information, ensuring secure and controlled access.
        - **Integration with Devices:** Links with registered devices and servers to ensure logs are associated correctly within the network.
        - **Relationship Management:** Facilitates connections with related objects, supporting cohesive log management.

---

# Data Source: intersight_equipment_end_point_log
The EndPointLog object encapsulates log file information collected from Cisco endpoints. This provides insights into log collection and management processes.
#### Purpose
EndPointLog serves to document and manage log files collected from endpoints, supporting troubleshooting and operational insights. It aids in maintaining transparency and accountability in log management.
#### Key Concepts
- **Log Management:** Handles log file types and collection statuses, offering visibility into endpoint operations.
- **Access Control:** Defines privilege sets for reading log information, ensuring secure and controlled access.
- **Integration with Devices:** Links with registered devices and servers to ensure logs are associated correctly within the network.
- **Relationship Management:** Facilitates connections with related objects, supporting cohesive log management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_end_point_log.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `collection_time`:(string) The time at which the log was collected. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `download_url`:(string) The Url to download the end point log. 
* `file_name`:(string) The end point log file name. 
* `log_type`:(string) The end point log file type.* `None` - End point log file type None.* `SEL` - End point log file type SEL. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The end point log collection status.* `None` - Log collection not started.* `CollectionInProgress` - Log file collection is in progress.* `CollectionCompleted` - Log file collection completed.* `CollectionFailed` - Log file collection failed.* `UploadInProgress` - Log file upload is in progress.* `UploadCompleted` - Log file upload completed.* `UploadFailed` - Log file upload failed to complete.* `DownloadUrlCreationFailed` - Download Url creation failed.* `Completed` - Log collection and upload completed. 
 
