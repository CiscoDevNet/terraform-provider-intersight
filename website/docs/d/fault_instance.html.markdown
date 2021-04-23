---
subcategory: "fault"
layout: "intersight"
page_title: "Intersight: intersight_fault_instance"
description: |-
  An endpoint anomaly is represented by this object.
---

# Data Source: intersight_fault_instance
An endpoint anomaly is represented by this object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fault_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `acknowledged`:(string) The user acknowledgement state of the fault. 
* `affected_dn`:(string) The Distinguished Name of the Managed object which was affected. 
* `affected_mo_id`:(string) Managed object Id which was affected. 
* `affected_mo_type`:(string) Managed object type which was affected. 
* `ancestor_mo_id`:(string) Object Id of the parent of the Managed object which was affected. 
* `ancestor_mo_type`:(string) Object type of the parent of the Managed object which was affected. 
* `code`:(string) Numerical fault code of the fault found. 
* `create_time`:(string) The time when this managed object was created. 
* `creation_time`:(string) The time of creation of the fault instance. 
* `description`:(string) Detailed message of the fault. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_transition_time`:(string) Last transition time of the fault. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_occurrences`:(int) The number of times this fault has occured. 
* `original_severity`:(string) Current Severity of the fault found. 
* `previous_severity`:(string) The Severity of the fault prior to user update. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `rule`:(string) The rule that is responsible for generation of the fault. 
* `severity`:(string) Severity of the fault found. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 