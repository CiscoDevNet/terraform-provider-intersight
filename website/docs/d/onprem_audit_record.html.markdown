---
subcategory: "onprem"
layout: "intersight"
page_title: "Intersight: intersight_onprem_audit_record"
description: |-
        AuditRecord presents the configuration changes made by the user per transaction.

---

# Data Source: intersight_onprem_audit_record
AuditRecord presents the configuration changes made by the user per transaction.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_onprem_audit_record.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `event`:(string) The type of event that occurred. Possible values are \ Login\ , \ Logout\ , \ Created\ , \ Modified\ , \ Deleted\ . 
* `http_operation`:(string) It contains the http request type and URL for the operation. In case of authentication request, this field \ POST /iam/adminlogin\ . 
* `http_response_code`:(int) The response code of the operation. If the operation is successful, this field will be 200. 
* `mo_type`:(string) The object type of the REST resource that was created, modified or deleted. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_moid`:(string) The Moid of the REST resource that was created, modified or deleted. 
* `session_id`:(string) The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_ip`:(string) The source IP of the client where the user action was performed. 
* `timestamp`:(string) The timestamp of the user action performed. 
* `trace_id`:(string) The trace id of the request that was used to create, modify or delete a REST resource.A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purposeby the Intersight technical support team. 
* `user_agent_string`:(string) The raw, string representation of the user agent of the request from the user-agent http request header. 
* `user_id_or_email`:(string) The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information. 
 
