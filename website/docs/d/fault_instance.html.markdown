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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledged`:(string) The user acknowledgement state of the fault. 
* `affected_dn`:(string) The Distinguished Name of the Managed object which was affected. 
* `affected_mo_id`:(string) Managed object Id which was affected. 
* `affected_mo_type`:(string) Managed object type which was affected. 
* `ancestor_mo_id`:(string) Object Id of the parent of the Managed object which was affected. 
* `ancestor_mo_type`:(string) Object type of the parent of the Managed object which was affected. 
* `code`:(string) Numerical fault code of the fault found. 
* `creation_time`:(string) The time of creation of the fault instance. 
* `description`:(string) Detailed message of the fault. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `last_transition_time`:(string) Last transition time of the fault. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_occurrences`:(int) The number of times this fault has occured. 
* `original_severity`:(string) Current Severity of the fault found. 
* `previous_severity`:(string) The Severity of the fault prior to user update. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `rule`:(string) The rule that is responsible for generation of the fault. 
* `severity`:(string) Severity of the fault found. 
