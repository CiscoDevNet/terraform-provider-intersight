---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_execution_snapshot"
description: |-
        Last known health check execution results of a health check Definition.

---

# Data Source: intersight_hyperflex_health_check_execution_snapshot
Last known health check execution results of a health check Definition.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_health_check_execution_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) Category that the HyperFlex health check Definition belongs to. 
* `cause`:(string) Information detailing the possible cause of the healthcheck failure, if the check fails. 
* `completion_time`:(string) Health check execution completion time. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health_check_details`:(string) Details of the health check execution result. 
* `health_check_execution_error_details`:(string) Error details of a script execution failure. 
* `health_check_execution_error_summary`:(string) Error summary of a script execution failure. 
* `health_check_execution_status`:(string) Status of the health check execution.* `UNKNOWN` - Indicates that the health heck execution results are unknown.* `SUCCEEDED` - Indicates that the health check execution succeeded.* `FAILED` - Indicates that the health check execution failed.* `TIMED_OUT` - Indicates that the health check execution timed out before completion. 
* `health_check_result`:(string) Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED.* `UNKNOWN` - Indicates that the health check results could not be determined.* `PASS` - Indicates that the health check passed.* `FAIL` - Indicates that the health check failed.* `WARN` - Indicates that the health check completed with a warning.* `NOT_APPLICABLE` - Indicates that the health check is either unsupported, or not applicable on the Cluster. 
* `health_check_summary`:(string) A brief summary of health check results. 
* `health_check_vcenter_ip`:(string) IP Address of the vCenter. 
* `hx_device_name`:(string) HyperFlex Device Name where the healthcheck is executed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `suggested_resolution`:(string) Information detailing a suggegsted resolution for the healthcheck failure, if the check fails. 
 
