---
subcategory: "servicenow"
layout: "intersight"
page_title: "Intersight: intersight_servicenow_change_request"
description: |-
        The ChangeRequest object is a core component of the ServiceNow integration framework, designed to manage and document change requests within the ServiceNow system. It ensures organized tracking of changes, providing a systematic approach to change management.
        #### Purpose
        A ChangeRequest object encapsulates the details of a change request, including its identification, status, and associated metadata. It serves as the central point of reference for managing change requests, ensuring all change-related activities are documented and tracked effectively.
        #### Key Concepts
        - **Identification and Tracking:** Provides unique identifiers for change requests, facilitating easy reference and management.
        - **Metadata Management:** Stores detailed information about change requests, including descriptions, priorities, and impacts, ensuring comprehensive documentation.
        - **Integration with ServiceNow:** Seamlessly interacts with the ServiceNow platform, supporting efficient change management workflows.
        - **Access Control:** Ensures controlled access to change request information, maintaining security and privacy standards within the system.

---

# Data Source: intersight_servicenow_change_request
The ChangeRequest object is a core component of the ServiceNow integration framework, designed to manage and document change requests within the ServiceNow system. It ensures organized tracking of changes, providing a systematic approach to change management.
#### Purpose
A ChangeRequest object encapsulates the details of a change request, including its identification, status, and associated metadata. It serves as the central point of reference for managing change requests, ensuring all change-related activities are documented and tracked effectively.
#### Key Concepts
- **Identification and Tracking:** Provides unique identifiers for change requests, facilitating easy reference and management.
- **Metadata Management:** Stores detailed information about change requests, including descriptions, priorities, and impacts, ensuring comprehensive documentation.
- **Integration with ServiceNow:** Seamlessly interacts with the ServiceNow platform, supporting efficient change management workflows.
- **Access Control:** Ensures controlled access to change request information, maintaining security and privacy standards within the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_servicenow_change_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `approval`:(string) The approver of Change Request. 
* `assigned_to`:(string) Assigned to value for Change Request. 
* `assigned_to_display_value`:(string) Assigned to display value for Change Request. 
* `change_model`:(string) Change Model for Change Request. 
* `change_model_display_value`:(string) Change Model display value for Change Request. 
* `change_request_number`:(string) Number for Change Request. 
* `comments`:(string) Comments on Change Request. 
* `create_time`:(string) The time when this managed object was created. 
* `created_by`:(string) Creator of Change Request. 
* `description`:(string) Description for Change Request. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `due_date`:(string) Due date for Change Request. 
* `end_date`:(string) End date for Change Request. 
* `impact`:(string) Impact for Change Request. 
* `impact_display_value`:(string) Impact display value for Change Request. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `priority`:(string) Priority for Change Request. 
* `priority_display_value`:(string) Priority display value for Change Request. 
* `risk`:(string) The risk for Change Request. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `short_description`:(string) Short Description for Change Request. 
* `start_date`:(string) Start date for Change Request. 
* `sys_id`:(string) Sys Id for Change Request. 
* `type`:(string) The type for Change Request. 
* `updated_by`:(string) Last update Change Request. 
 
