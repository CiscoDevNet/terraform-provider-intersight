---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_workload_definition"
description: |-
        ### Overview
        The WorkloadDefinition object represents the blueprint for defining reusable workload configurations that can be deployed across one or more organizations. It serves as the foundational template that captures the complete specification of a workload, including its associated blueprints, validation requirements, and deployment strategies.
        #### Purpose
        A WorkloadDefinition provides a versioned, reusable specification for workloads that can be deployed multiple times across different contexts organizations. It enables administrators to define complex multi-component workloads once and deploy them consistently, supporting both iterative development through versioning and controlled rollout through deployment strategies.
        #### Key Concepts
        - **Platform Support:** - Platform types can be specified to indicate which Intersight platforms support the workload, ensuring deployments target compatible infrastructure.
        - **Versioning and Preferred Versions:** - Each workload definition supports multiple versions, with the ability to designate a preferred (default) version. This allows for safe evolution of workload specifications while maintaining backward compatibility with existing deployments.
        - **Blueprint Composition:** - Workload definitions are composed of one or more blueprint references, enabling modular design where complex workloads are built from reusable blueprint components.
        - **Validation and Status Tracking:** - Each definition maintains validation information to ensure all required inputs and dependencies are properly configured before deployment, along with deployment summary aggregations across all instances.
        - **Override Inputs:** - Each workload definition can specify input overrides for the blueprints it references, allowing for customization within deployments.

---

# Resource: intersight_workload_workload_definition
### Overview
The WorkloadDefinition object represents the blueprint for defining reusable workload configurations that can be deployed across one or more organizations. It serves as the foundational template that captures the complete specification of a workload, including its associated blueprints, validation requirements, and deployment strategies.
#### Purpose
A WorkloadDefinition provides a versioned, reusable specification for workloads that can be deployed multiple times across different contexts organizations. It enables administrators to define complex multi-component workloads once and deploy them consistently, supporting both iterative development through versioning and controlled rollout through deployment strategies.
#### Key Concepts
- **Platform Support:** - Platform types can be specified to indicate which Intersight platforms support the workload, ensuring deployments target compatible infrastructure.
- **Versioning and Preferred Versions:** - Each workload definition supports multiple versions, with the ability to designate a preferred (default) version. This allows for safe evolution of workload specifications while maintaining backward compatibility with existing deployments.
- **Blueprint Composition:** - Workload definitions are composed of one or more blueprint references, enabling modular design where complex workloads are built from reusable blueprint components.
- **Validation and Status Tracking:** - Each definition maintains validation information to ensure all required inputs and dependencies are properly configured before deployment, along with deployment summary aggregations across all instances.
- **Override Inputs:** - Each workload definition can specify input overrides for the blueprints it references, allowing for customization within deployments.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `blueprints`:(Array)
This complex property has following sub-properties:
  + `blueprint`:(HashMap) - The referred blueprint details. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `input`:(JSON as string) The input data for the referred blueprint. All required inputs of the blueprint must have a value provided. 
  + `input_override`:
                (Array of schema.TypeString) -
  + `name`:(string) The name for the referred blueprint. This name must be unique within the workload definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_name`:(string)(ReadOnly) The reference name for the blueprint which is derived by the system. 
  + `resource_constraint`:(HashMap) - The constraints that need to be applied to the resources in order to match this blueprint. 
This complex property has following sub-properties:
    + `input`:(JSON as string) The input values from the user for the resource definition of the blueprint. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `deployment_summary`:(Array)
This complex property has following sub-properties:
  + `nr_count`:(int)(ReadOnly) The total number of referenced objects included in the aggregation. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `state`:(string)(ReadOnly) The overall aggregated state as a string, summarizing the status of all referenced objects. 
* `description`:(string) The description for this workload definition. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this Workload. You can have multiple versions of the Workload with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), space, hyphen (-), or an underscore (_). A refName is generated from the given name, and that along with the version must be unique within an Organization. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
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
* `platform_type`:
                (Array of schema.TypeString) -
* `preferred_version`:(bool) The flag to indicate that this is the preferred (default) version of the workload. 
* `preferred_version_rollout_strategy`:(HashMap) - The strategy used for rolling out deployment changes when this workload version is marked as the preferred version. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workload.BatchDeployment](#workloadBatchDeployment)
[workload.CanaryDeployment](#workloadCanaryDeployment)
  + `failure_threshold`:(int)(ReadOnly) Specifies no of errors can be allowed to skip executing next set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `ref_name`:(string)(ReadOnly) A reference name is generated by the system based on the given name by replacing spaces and hyphen in name with underscore. This reference name is used internally and cannot be edited by users. It may only contain letters (a–z, A–Z), numbers (0–9), and underscores (_). 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the definition.* `Inactive` - The definition is in an inactive state and there are no workload instances associated with this workload. Changes to input parameters within the workload are allowed in this state.* `Active` - The definition is active and associated with one or more workload instances. 
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
* `validation_information`:(HashMap) -(ReadOnly) The current validation state and associated validation errors when state is invalid. When state is valid, it means all the required inputs for the blueprint are provided. 
This complex property has following sub-properties:
  + `engine_state`:(string)(ReadOnly) The state of workflow definition metadata in the workflow engine. The workflow definition must be successfully updated in the engine before workflows can be executed.* `NotUpdated` - The workflow and task definition metadata is not yet updated in the workflow engine.* `Updating` - The workflow and task definition metadata is in the processing of being updated in the workflow engine.* `UpdateFailed` - The workflow and task definition metadata failed to be updated in the workflow engine.* `Updated` - The workflow and task definition metadata was updated successfully in the workflow engine. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `state`:(string)(ReadOnly) The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).* `NotValidated` - The state when workflow definition has not been validated.* `Valid` - The state when workflow definition is valid.* `Invalid` - The state when workflow definition is invalid. 
  + `validation_error`:(Array)
This complex property has following sub-properties:
    + `error_log`:(string)(ReadOnly) Description of the error. 
    + `field`:(string)(ReadOnly) When populated this refers to the input or output field within the workflow or task. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `task_name`:(string)(ReadOnly) The task name on which the error is found, when empty the error applies to the top level workflow. 
    + `transition_name`:(string)(ReadOnly) When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue. 
* `nr_version`:(int) The version of the workload to support multiple versions. 
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
* `workload_metadata`:(HashMap) -(ReadOnly) A reference to a workloadWorkloadMetadata resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_workload_workload_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workload_workload_definition.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [workload.BatchDeployment](#argument-reference)
The number of instances to be included in each batch during the deployment..
* `batch_size`:(int)(ReadOnly) The number of instances to be included in each batch during the deployment. 

### [workload.CanaryDeployment](#argument-reference)
The percentage of instances to which the changes will be rolled out in each step of the deployment.
* `rollout_percentages`:
                (Array of schema.TypeInt) -
  
