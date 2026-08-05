---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_execution"
description: |-
        The HealthCheckExecution object is a pivotal element in the HyperFlex system, detailing the execution process and outcomes for health checks performed on HyperFlex devices. It encapsulates the execution lifecycle, from initiation through completion, providing visibility into the health check's effect on device stability and performance.
        #### Purpose
        HealthCheckExecution objects document the execution of health checks, offering a detailed account of how these checks were conducted, their results, and any issues encountered. This supports informed decision-making for system improvements and resource management.
        #### Key Concepts
        - **Lifecycle Management:** Tracks the execution phase of health checks, including initiation, progress, and final results, enabling thorough analysis of operational health.
        - **Outcome Documentation:** Records the result of health check executions, highlighting successful checks and identifying areas of concern.
        - **Error Reporting:** Provides detailed error information for any execution failures, facilitating targeted troubleshooting and remediation efforts.
        - **Integration with HyperFlex Devices:** Directly associated with specific HyperFlex devices, enabling device-specific health assessments and optimizations.

---

# Data Source: intersight_hyperflex_health_check_execution
The HealthCheckExecution object is a pivotal element in the HyperFlex system, detailing the execution process and outcomes for health checks performed on HyperFlex devices. It encapsulates the execution lifecycle, from initiation through completion, providing visibility into the health check's effect on device stability and performance.
#### Purpose
HealthCheckExecution objects document the execution of health checks, offering a detailed account of how these checks were conducted, their results, and any issues encountered. This supports informed decision-making for system improvements and resource management.
#### Key Concepts
- **Lifecycle Management:** Tracks the execution phase of health checks, including initiation, progress, and final results, enabling thorough analysis of operational health.
- **Outcome Documentation:** Records the result of health check executions, highlighting successful checks and identifying areas of concern.
- **Error Reporting:** Provides detailed error information for any execution failures, facilitating targeted troubleshooting and remediation efforts.
- **Integration with HyperFlex Devices:** Directly associated with specific HyperFlex devices, enabling device-specific health assessments and optimizations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_health_check_execution.<custom_name>.results[i].<propertyname>`.
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
* `suggested_resolution`:(string) Information detailing a suggested resolution for the healthcheck failure, if the check fails. 
* `uuid`:(string) UUID of an instance of health check execution. 
 
