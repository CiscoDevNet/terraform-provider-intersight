---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_custom_hcl_status"
description: |-
        The CustomHclStatus object represents the compliance status of a Cisco server against defined Custom Hcl baseline  within an organization. It contains detailed insights into whether a server meets the specified hardware and software criteria set forth in the baseline.
        #### Purpose
        CustomHclStatus is crucial for monitoring and managing server compliance with organizational standards. It helps administrators identify non-compliant servers, understand the reasons for non-compliance, and take corrective actions to ensure that all servers adhere to the established custom Hcl baselines.
        #### Key Concepts
        - **Compliance Status:** Indicates whether a server is compliant or non-compliant with defined custom Hcl baselines.
        - **Detailed Properties:** Includes various properties such as server model, generation, firmware, operating system details, and more for comprehensive status reporting.
        - **Relationships:** Links to managed objects and custom Hcl baselines to provide context for compliance evaluations.

---

# Data Source: intersight_cond_custom_hcl_status
The CustomHclStatus object represents the compliance status of a Cisco server against defined Custom Hcl baseline  within an organization. It contains detailed insights into whether a server meets the specified hardware and software criteria set forth in the baseline.
#### Purpose
CustomHclStatus is crucial for monitoring and managing server compliance with organizational standards. It helps administrators identify non-compliant servers, understand the reasons for non-compliance, and take corrective actions to ensure that all servers adhere to the established custom Hcl baselines.
#### Key Concepts
- **Compliance Status:** Indicates whether a server is compliant or non-compliant with defined custom Hcl baselines.
- **Detailed Properties:** Includes various properties such as server model, generation, firmware, operating system details, and more for comprehensive status reporting.
- **Relationships:** Links to managed objects and custom Hcl baselines to provide context for compliance evaluations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_custom_hcl_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cisco_hcl_status`:(string) The HCL compatibility status of the server. The status can be one of the following \ Incomplete\  - there is no enough information to evaluate against the HCL data \ Validated\  - all components have been evaluated against the HCL and they all have \ Validated\  status \ Not-Listed\  - all components have been evaluated against the HCL and one or more have \ Not-Listed\  status. \ Not-Evaluated\  - The server was not evaluated against the HCL because it is exempcustom Hcl t or the provided is insufficient for HCL status assessment.* `Incomplete` - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component's hardware or software profile was not found in the HCL.* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
* `create_time`:(string) The time when this managed object was created. 
* `custom_hcl_status`:(string) The custom HCL overall status against all the defined custom Hcl baseline. The status can be one of the following \ CompliantToBaseline\  - The server is compliant to one or all the defined custom Hcl baseline. \ NonCompliantToBaseline\  - The server is non compliant to any custom Hcl baseline.* `NonCompliantToBaseline` - The server is non compliant to any custom Hcl baseline.* `CompliantToBaseline` - The server is compliant to one or all the defined custom Hcl baseline. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware`:(string) The firmware version currently running on the server. 
* `generation`:(string) It specifies the generation of the server like M6, M7, M8 etc. 
* `management_mode`:(string) The management mode through which the server is connected to Intersight, such as standalone or managed mode.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) It specifies the server name or model. 
* `os_vendor`:(string) The operating System vendor of the server. 
* `os_version`:(string) The operating System version of the server. 
* `personality`:(string) Unique identity of added software personality. 
* `processor_family`:(string) The processor family of the specified processor model associated with the server. 
* `processor_model`:(string) The processor model associated with the server. 
* `server_model`:(string) It specifies the server model or Product ID (PID). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
