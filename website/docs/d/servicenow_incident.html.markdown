---
subcategory: "servicenow"
layout: "intersight"
page_title: "Intersight: intersight_servicenow_incident"
description: |-
        A Incident on ServiceNow.

---

# Data Source: intersight_servicenow_incident
A Incident on ServiceNow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_servicenow_incident.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `approval`:(string) The approver property of Incident. 
* `category`:(string) Category property for Incident. 
* `comments`:(string) Comments property on Incident. 
* `create_time`:(string) The time when this managed object was created. 
* `created_by`:(string) Creator property of Incident. 
* `created_on`:(string) Incident create date property. 
* `description`:(string) Description for Incident. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `due_date`:(string) Due date property for Incident. 
* `expected_start`:(string) Expected start date for Incident. 
* `impact`:(string) Impact property for Incident. 
* `incident_state`:(string) State property of the Incident. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `opened_by`:(string) Assigned to value for Incident. 
* `priority`:(string) Priority property for Incident. 
* `risk`:(string) The risk property for Incident. 
* `severity`:(string) Severity property of the Incident. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `short_description`:(string) Short Description for Incident. 
* `sys_id`:(string) System Id property for Incident. 
* `task_effective_number`:(string) Task Effective Number for Incident. 
* `updated_by`:(string) Last update by on Incident. 
* `urgency`:(string) Urgency property of the Incident. 
* `work_end`:(string) Work end date for Incident. 
* `work_start`:(string) Work start date for Incident. 
 
