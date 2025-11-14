---
subcategory: "convergedinfra"
layout: "intersight"
page_title: "Intersight: intersight_convergedinfra_pod_compliance_info"
description: |-
        ### Overview
        The PodComplianceInfo object captures compliance information for a converged infrastructure pod, extending the base compliance details to include specific compliance assessments for the pod itself.
        #### Purpose
        PodComplianceInfo provides a detailed view of the compliance status of a pod, highlighting areas that meet or deviate from established standards. This  supports comprehensive compliance management and helps identify non-compliant components.
        #### Key Concepts
        - **Compliance Tracking:** - Offers insights into the compliance status of the entire pod, facilitating proactive management and remediation.
        - **Inheritance:** - Leverages the base compliance details, extending them to encompass pod-specific compliance metrics.
        - **Detailed Relationships:** - Establishes connections with the pod and its components, enabling a holistic view of compliance across the infrastructure.
        - **Administrative Access:** - Supports a range of privilege sets, allowing different levels of administrative interaction with compliance data.

---

# Data Source: intersight_convergedinfra_pod_compliance_info
### Overview
The PodComplianceInfo object captures compliance information for a converged infrastructure pod, extending the base compliance details to include specific compliance assessments for the pod itself.   
#### Purpose  
PodComplianceInfo provides a detailed view of the compliance status of a pod, highlighting areas that meet or deviate from established standards. This  supports comprehensive compliance management and helps identify non-compliant components.   
#### Key Concepts  
- **Compliance Tracking:** - Offers insights into the compliance status of the entire pod, facilitating proactive management and remediation. 
- **Inheritance:** - Leverages the base compliance details, extending them to encompass pod-specific compliance metrics. 
- **Detailed Relationships:** - Establishes connections with the pod and its components, enabling a holistic view of compliance across the infrastructure.
- **Administrative Access:** - Supports a range of privilege sets, allowing different levels of administrative interaction with compliance data.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_convergedinfra_pod_compliance_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the component for which the compliance is evaluated. 
* `reason`:(string) Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant.* `NotEvaluated` - The validation for the HCL or Interop status is yet to be performed.* `Missing-Os-Info` - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information.* `Incompatible-Storage-Os` - The validation failed because the given storage name and version combination is not found in Interop matrix.* `Incompatible-Os-Info` - The validation failed because the given OS name and version combination is not found in HCL.* `Incompatible-Processor` - The validation failed because the given processor is not found for the given server PID in HCL.* `Incompatible-Server-Platform` - The validation failed because the given server platform is not found in the Interop matrix.* `Incompatible-Server-Model` - The validation failed because the given server model is not found in the Interop matrix.* `Incompatible-Adapter-Model` - The validation failed because the given adapter model is not found in the Interop matrix.* `Incompatible-Switch-Model` - The validation failed because the given switch model is not found in the Interop matrix.* `Incompatible-Server-Firmware` - The validation failed because the given server firmware version is not found in HCL.* `Incompatible-Switch-Firmware` - The validation failed because the given switch firmware version is not found in Interop matrix.* `Incompatible-Firmware` - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix.* `Incompatible-Driver` - The validation failed because the given driver version is not found in either HCL or Interop matrix.* `Incompatible-Firmware-Driver` - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix.* `Missing-Os-Driver-Info` - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.* `Missing-Os-Or-Driver-Info` - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers.* `Incompatible-Components` - The validation failed for the given server because one or more of its components failed validation.* `Compatible` - This means the HCL status and Interop status for the component have passed all validations for all of its related components. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Compliance status for the component.* `NotEvaluated` - The interoperability compliance for the component has not be checked.* `Approved` - The component is valid as per the interoperability compliance check.* `NotApproved` - The component is not valid as per the interoperability compliance check.* `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data. 
 
