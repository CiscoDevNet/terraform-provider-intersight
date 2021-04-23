---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_diagnostic_messages"
description: |-
  Gets diagnostics messages from UCSD.
---

# Data Source: intersight_iaas_diagnostic_messages
Gets diagnostics messages from UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_diagnostic_messages.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) Message category of the alerts. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `item`:(string) Message target type of the alerts. 
* `last_checked`:(string) Last checked time of the alerts. 
* `message`:(string) Detailed info about the alert. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `recommendation`:(string) Recommended fix for the alert. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the given alert. 
* `status_id`:(string) Status Id of the given alert. 
 