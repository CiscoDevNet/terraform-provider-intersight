---
subcategory: "convergedinfra"
layout: "intersight"
page_title: "Intersight: intersight_convergedinfra_server_compliance_details"
description: |-
        ### Overview
        The ServerComplianceDetails object provides compliance information for servers within a converged infrastructure pod, detailing hardware compatibility and operational status.
        #### Purpose
        ServerComplianceDetails is focused on evaluating server compliance, ensuring that server components meet hardware compatibility list (HCL) criteria and operational standards necessary for reliable infrastructure performance.
        #### Key Concepts
        - **HCL Compatibility:** - Provides status and reasoning related to hardware compatibility, ensuring that server components are aligned with prescribed standards.
        - **Platform Information:** - Encompasses detailed information about server platforms, processors, firmware, and operating systems, aiding in compliance evaluation.
        - **Integration with Pod Compliance:** - Connects server compliance data with pod-level compliance insights, facilitating integrated compliance management.
        - **Read-Only Access:** - Designed for secure consumption of server compliance data, ensuring consistent integration with management interfaces.

---

# Data Source: intersight_convergedinfra_server_compliance_details
### Overview
The ServerComplianceDetails object provides compliance information for servers within a converged infrastructure pod, detailing hardware compatibility and operational status.   
#### Purpose  
ServerComplianceDetails is focused on evaluating server compliance, ensuring that server components meet hardware compatibility list (HCL) criteria and operational standards necessary for reliable infrastructure performance.   
#### Key Concepts  
- **HCL Compatibility:** - Provides status and reasoning related to hardware compatibility, ensuring that server components are aligned with prescribed standards. 
- **Platform Information:** - Encompasses detailed information about server platforms, processors, firmware, and operating systems, aiding in compliance evaluation. 
- **Integration with Pod Compliance:** - Connects server compliance data with pod-level compliance insights, facilitating integrated compliance management. 
- **Read-Only Access:** - Designed for secure consumption of server compliance data, ensuring consistent integration with management interfaces.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_convergedinfra_server_compliance_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_count`:(int) The number of ethernet NIC adapters in the server. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware`:(string) The Cisco IMC firmware version of the server. 
* `hcl_status`:(string) The HCL compatibility status of the server.* `NotEvaluated` - The interoperability compliance for the component has not be checked.* `Approved` - The component is valid as per the interoperability compliance check.* `NotApproved` - The component is not valid as per the interoperability compliance check.* `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data. 
* `hcl_status_reason`:(string) The reason for server's HCL status.* `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information.* `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL.* `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model.* `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination.* `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version.* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.* `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result.* `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.* `Incompatible-Components` - The validation has failed for this server because one or more components have \ Not-Listed\  status.* `Compatible` - The validation has passed for this server's model, processor, OS vendor and version. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model information of the server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the component for which the compliance is evaluated. 
* `os`:(string) Details of name and version of the operating system running on the server. 
* `platform`:(string) Details of platform of the server, examples are B-Series, C-Series, X-Series etc. 
* `processor`:(string) The processor information of the server. 
* `reason`:(string) Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant.* `NotEvaluated` - The validation for the HCL or Interop status is yet to be performed.* `Missing-Os-Info` - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information.* `Incompatible-Storage-Os` - The validation failed because the given storage name and version combination is not found in Interop matrix.* `Incompatible-Os-Info` - The validation failed because the given OS name and version combination is not found in HCL.* `Incompatible-Processor` - The validation failed because the given processor is not found for the given server PID in HCL.* `Incompatible-Server-Platform` - The validation failed because the given server platform is not found in the Interop matrix.* `Incompatible-Server-Model` - The validation failed because the given server model is not found in the Interop matrix.* `Incompatible-Adapter-Model` - The validation failed because the given adapter model is not found in the Interop matrix.* `Incompatible-Switch-Model` - The validation failed because the given switch model is not found in the Interop matrix.* `Incompatible-Server-Firmware` - The validation failed because the given server firmware version is not found in HCL.* `Incompatible-Switch-Firmware` - The validation failed because the given switch firmware version is not found in Interop matrix.* `Incompatible-Firmware` - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix.* `Incompatible-Driver` - The validation failed because the given driver version is not found in either HCL or Interop matrix.* `Incompatible-Firmware-Driver` - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix.* `Missing-Os-Driver-Info` - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.* `Missing-Os-Or-Driver-Info` - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers.* `Incompatible-Components` - The validation failed for the given server because one or more of its components failed validation.* `Compatible` - This means the HCL status and Interop status for the component have passed all validations for all of its related components. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Compliance status for the component.* `NotEvaluated` - The interoperability compliance for the component has not be checked.* `Approved` - The component is valid as per the interoperability compliance check.* `NotApproved` - The component is not valid as per the interoperability compliance check.* `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data. 
 
