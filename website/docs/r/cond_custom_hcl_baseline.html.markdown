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

# Resource: intersight_cond_custom_hcl_baseline
The CustomHclBaseline object defines baseline configurations for Cisco servers within an organization. It allows administrators to specify hardware and software requirements that servers must meet to be considered compliant with organizational standards.
#### Purpose
CustomHclBaseline serves as a framework for establishing and enforcing baseline configurations, ensuring that servers adhere to defined hardware and software criteria. This helps maintain consistency, reliability, and performance across the organization's server infrastructure.
#### Key Concepts
- **Baseline Configuration:** Defines specific hardware and software parameters that servers must meet.
- **Compliance Monitoring:** Tracks the compliance status of servers against the defined baseline configurations.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `adapters`:(Array)
This complex property has following sub-properties:
  + `driver_name`:(string) It specifies the name of the driver supported by the adapter. 
  + `driver_version`:(string) It specifies the version of the driver installed for the adapter. 
  + `firmware`:(string) It specifies the firmware version installed on the adapter. 
  + `hcl_reason`:(string)(ReadOnly) The reason for the Cisco HCL validation status, more useful when status is NotListed.* `Missing-Os-Driver-Info` - The validation failed because the given server has no operating system driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL.* `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID.* `Incompatible-Os-Info` - The validation failed because the given operating system vendor and version was not found in HCL for the server PID and processor combination.* `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and operating system vendor and version.* `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, operating system vendor and version and component model.* `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, operating system vendor and version, server firmware and component firmware.* `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, operating system vendor and version and server firmware.* `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.* `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result.* `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.* `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.* `Compatible` - The validation has passed for this server PID, processor, operating system vendor and version, component model, component firmware and driver version. 
  + `hcl_status`:(string)(ReadOnly) The Cisco HCL validation status of the adapter.* `Incomplete` - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component's hardware or software profile was not found in the HCL.* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
  + `model`:(string) It specifies the model name or Product ID (PID) of the adapter. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `type`:(string) It specifies the type of the adapter, such as Network Adapter, Storage Controller, or GPU.* `Unknown` - It indicates no adapter type. it is used when component type enum is not defined.* `NetworkAdapter` - Indicates network adapter.* `StorageController` - Indicates Storage controller.* `GPU` - It refers to Graphics Processing Unit (GPU) adapters, which are used for rendering graphics and performing parallel processing tasks.* `SSD` - It refers NVME solid state drives (SSD) drives. 
  + `vendor`:(string) It specifies the vendor or manufacturer of the adapter. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `compliant_server_count`:(int)(ReadOnly) The number of servers that are compliant with this custom Hcl baseline. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the custom Hcl baseline. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `firmware`:(string) The firmware version of currently running on the server. 
* `generation`:(string) It specifies the generation of the server. like M7, M8 etc. 
* `hcl_reason`:(string)(ReadOnly) The reason of the current Cisco HCL status of the custom Hcl baseline.* `Missing-Os-Info` - This means the HclStatus for the server failed HCL validation because we have missing operating system information. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Incompatible-Components` - This means the HclStatus for the server failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails.* `Compatible` - This means the HclStatus for the server has passed HCL validation for all of its related components.* `Not-Evaluated` - This means the HclStatus for the server has not been evaluated because it is exempted.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
* `hcl_status`:(string)(ReadOnly) The Cisco HCL compatibility status of the custom Hcl baseline. The status can be one of the following \ Incomplete\  - there is not enough information to evaluate against the HCL data \ Validated\  - all components have been evaluated against the HCL and they all have \ Validated\  status \ Not-Listed\  - all components have been evaluated against the HCL and one or more have \ Not-Listed\  status. \ Not-Evaluatecustom Hcl d\  - The provided is insufficient for HCL status evaluation or the server is exempted from HCL validation \ Not-Applicable\  - the custom Hcl baseline is not applicable to the server.* `Incomplete` - This means we do not have operating system information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper operating system information.* `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that in component's hardware or software profile was not found in the HCL.* `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.* `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.* `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet or the configurations provided are insufficient for HCL status assessment. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.* `Not-Applicable` - At the HclStatus level this means that the custom Hcl provided is not applicable to the server. 
* `management_mode`:(string) The management mode at which server is connected to intersight. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the custom Hcl baseline. 
* `non_compliant_server_count`:(int)(ReadOnly) The number of servers that are non-compliant with this custom Hcl baseline. 
* `operation_state`:(string)(ReadOnly) Operation state specifies the state of custom Hcl baseline whether the server's baseline status is evaluated or not. \ InProgress\  - Server's baseline status assessment is currently in progress \ Pending\  - Server's baseline status assessment is not yet started \ Active\  - Server's baseline status assessment is completed.* `Pending` - Server's custom hcl status assessment is not yet started.* `Active` - Server's custom hcl status assessment is completed.* `InProgress` - Server's custom hcl status assessment is currently in progress.* `Failed` - Server's custom hcl status assessment is failed. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `os_vendor`:(string) The operating system vendor name running on the server. 
* `os_version`:(string) Operating System version running on the server. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `processor_family`:(string) The processor family associated with the server for example processor model Intel (R) Xeon (R) Platinum 8454H or AMD EPYC 9654. will be validated using its corresponding processor family. Processor model Intel (R) Xeon (R) Platinum 8454H -> 4th Gen Intel Xeon Processor Scalable Family Processor model AMD EPYC 9654 -> 4th Gen AMD EPYC 9004 Series Processors. 
* `server_model`:
                (Array of schema.TypeString) -
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_cond_custom_hcl_baseline` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_cond_custom_hcl_baseline.example 1234567890987654321abcde
``` 
