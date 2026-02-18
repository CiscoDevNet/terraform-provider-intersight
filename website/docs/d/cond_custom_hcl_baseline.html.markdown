---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_custom_hcl_baseline"
description: |-
        The CustomHclBaseline object defines baseline configurations for Cisco servers within an organization. It allows administrators to specify hardware and software requirements that servers must meet to be considered compliant with organizational standards.
        #### Purpose
        CustomHclBaseline serves as a framework for establishing and enforcing baseline configurations, ensuring that servers adhere to defined hardware and software criteria. This helps maintain consistency, reliability, and performance across the organization's server infrastructure.
        #### Key Concepts
        - **Baseline Configuration:** Defines specific hardware and software parameters that servers must meet.
        - **Compliance Monitoring:** Tracks the compliance status of servers against the defined baseline configurations.

---

# Data Source: intersight_cond_custom_hcl_baseline
The CustomHclBaseline object defines baseline configurations for Cisco servers within an organization. It allows administrators to specify hardware and software requirements that servers must meet to be considered compliant with organizational standards.
#### Purpose
CustomHclBaseline serves as a framework for establishing and enforcing baseline configurations, ensuring that servers adhere to defined hardware and software criteria. This helps maintain consistency, reliability, and performance across the organization's server infrastructure.
#### Key Concepts
- **Baseline Configuration:** Defines specific hardware and software parameters that servers must meet.
- **Compliance Monitoring:** Tracks the compliance status of servers against the defined baseline configurations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_custom_hcl_baseline.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `compliant_server_count`:(int) The number of servers that are compliant with this custom Hcl baseline. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the custom Hcl baseline. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware`:(string) The firmware version of currently running on the server. 
* `generation`:(string) It specifies the generation of the server. like M7, M8 etc. 
* `hcl_reason`:(string) The reason of the current Cisco HCL status of the custom Hcl baseline.* `Missing-Os-Info` - This means the HclStatus for the server failed HCL validation because we have missing operating system information. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Incompatible-Components` - This means the HclStatus for the server failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails.* `Compatible` - This means the HclStatus for the server has passed HCL validation for all of its related components.* `Not-Evaluated` - This means the HclStatus for the server has not been evaluated because it is exempted.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
* `hcl_status`:(string) The Cisco HCL compatibility status of the custom Hcl baseline. The status can be one of the following \ Incomplete\  - there is not enough information to evaluate against the HCL data \ Validated\  - all components have been evaluated against the HCL and they all have \ Validated\  status \ Not-Listed\  - all components have been evaluated against the HCL and one or more have \ Not-Listed\  status. \ Not-Evaluatecustom Hcl d\  - The provided is insufficient for HCL status evaluation or the server is exempted from HCL validation \ Not-Applicable\  - the custom Hcl baseline is not applicable to the server.* `Incomplete` - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component's hardware or software profile was not found in the HCL.* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
* `management_mode`:(string) The management mode at which server is connected to intersight. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the custom Hcl baseline. 
* `non_compliant_server_count`:(int) The number of servers that are non-compliant with this custom Hcl baseline. 
* `operation_state`:(string) Operation state specifies the state of custom Hcl baseline whether the server's baseline status is evaluated or not. \ InProgress\  - Server's baseline status assessment is currently in progress \ Pending\  - Server's baseline status assessment is not yet started \ Active\  - Server's baseline status assessment is completed.* `Pending` - Server's custom hcl status assessment is not yet started.* `Active` - Server's custom hcl status assessment is completed.* `InProgress` - Server's custom hcl status assessment is currently in progress.* `Failed` - Server's custom hcl status assessment is failed. 
* `os_vendor`:(string) The operating system vendor name running on the server. 
* `os_version`:(string) Operating System version running on the server. 
* `processor_family`:(string) The processor family associated with the server for example processor model Intel (R) Xeon (R) Platinum 8454H or AMD EPYC 9654. will be validated using its corresponding processor family. Processor model Intel (R) Xeon (R) Platinum 8454H -> 4th Gen Intel Xeon Processor Scalable Family Processor model AMD EPYC 9654 -> 4th Gen AMD EPYC 9004 Series Processors. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
