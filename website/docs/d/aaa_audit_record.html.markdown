---
subcategory: "aaa"
layout: "intersight"
page_title: "Intersight: intersight_aaa_audit_record"
description: |-
        AuditRecord presents the configuration changes made by the user per transaction.

---

# Data Source: intersight_aaa_audit_record
AuditRecord presents the configuration changes made by the user per transaction.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_aaa_audit_record.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `affected_object_type_label`:(string) The user-friendly label for the object type that was changed. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email`:(string) The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information. 
* `event`:(string) The operation that was performed on this Managed Object.The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type. 
* `http_operation`:(string) The REST URL for the operation. 
* `inst_id`:(string) The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before. 
* `mo_type`:(string) The object type of the REST resource that was created, modified or deleted. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_moid`:(string) The Moid of the REST resource that was created, modified or deleted. 
* `session_id`:(string) The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_ip`:(string) The source IP of the client. 
* `timestamp`:(string) The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted. 
* `trace_id`:(string) The trace id of the request that was used to create, modify or delete a REST resource.A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purposeby the Intersight technical support team. 
* `user_agent_string`:(string) The raw, string representation of the user agent of the request from the user-agent http request header. 
* `user_id_or_email`:(string) The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information. 
 
