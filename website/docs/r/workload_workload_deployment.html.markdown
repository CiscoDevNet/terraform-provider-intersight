---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_workload_deployment"
description: |-
        ### Overview
        The WorkloadDeployment object represents an active deployment configuration that creates and manages workload instances based on a workload definition. It defines how a workload should be deployed to specific resources within an organization, including qualification criteria, overriden input parameters, and rollout strategies.
        #### Purpose
        A WorkloadDeployment bridges the gap between a workload definition (the template) and actual workload instances (the running resources). It specifies which resources should receive the workload through qualification criteria, provides deployment-specific configuration overrides, and manages the lifecycle of all workload instances created under its scope.
        #### Key Concepts
        - **Resource Qualification:** - Uses qualification criteria to dynamically determine which resources within an organization should receive the workload deployment, enabling automated scaling and resource targeting.
        - **Instance Naming and Management:** - Automatically generates unique names for workload instances using configurable prefixes and numeric suffixes, maintaining a consistent naming convention across deployments.
        - **Deployment Actions and Status:** - Tracks deployment actions and maintains status information, enabling administrators to monitor and control deployment lifecycle operations.
        - **Change Tracking and Validation:** - Records all configuration changes and maintains validation state to ensure deployments remain properly configured before executing actions.
        - **Conformance Monitoring:** - Continuously monitors whether deployed instances remain in conformance with the deployment specification, detecting and reporting configuration drift.
        - **Rollout Strategy:** - Supports configurable rollout strategies for controlling how changes are propagated to existing workload instances, enabling controlled updates with minimal disruption.

---

# Resource: intersight_workload_workload_deployment
### Overview
The WorkloadDeployment object represents an active deployment configuration that creates and manages workload instances based on a workload definition. It defines how a workload should be deployed to specific resources within an organization, including qualification criteria, overriden input parameters, and rollout strategies.
#### Purpose
A WorkloadDeployment bridges the gap between a workload definition (the template) and actual workload instances (the running resources). It specifies which resources should receive the workload through qualification criteria, provides deployment-specific configuration overrides, and manages the lifecycle of all workload instances created under its scope.
#### Key Concepts
- **Resource Qualification:** - Uses qualification criteria to dynamically determine which resources within an organization should receive the workload deployment, enabling automated scaling and resource targeting.
- **Instance Naming and Management:** - Automatically generates unique names for workload instances using configurable prefixes and numeric suffixes, maintaining a consistent naming convention across deployments.
- **Deployment Actions and Status:** - Tracks deployment actions and maintains status information, enabling administrators to monitor and control deployment lifecycle operations.
- **Change Tracking and Validation:** - Records all configuration changes and maintains validation state to ensure deployments remain properly configured before executing actions.
- **Conformance Monitoring:** - Continuously monitors whether deployed instances remain in conformance with the deployment specification, detecting and reporting configuration drift.
- **Rollout Strategy:** - Supports configurable rollout strategies for controlling how changes are propagated to existing workload instances, enabling controlled updates with minimal disruption.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) The action is performed on the deployment.* `None` - No changes have been made.* `PrepareToDeploy` - Marks the deployment as ready once the user has completed all changes and the deployment is in a valid state. Once the deployment is marked as PrepareToDeploy, workload instances will be created, but the actual deployment will occur as part of the deploy action.* `Deploy` - Initiates the process of pushing workload configuration to the instances based on the configured schedule. Once deployed, the deployment cannot be reverted to draft status.* `Suspend` - Suspends the deployment, preventing any further actions until it is resumed. Making changes to deployment configuration will not be pushed out. The deployment will not take any changes from the attached Workload when it is suspended.* `Retry` - Retries the deployment for all instances that previously failed.* `Resume` - Resumes a suspended deployment. Any changes made to the deployment when it was suspended or any changes made to the attached Workload will now be pushed out based on configured schedules.* `Undeploy` - Undeploy cleans up the policies, templates, and leases generated from the deployment, and marks the deployment as being in the Draft state. The associated definition will be marked as Inactive if there are no other deployments linked to it.* `ReEvaluateQualifiers` - Triggers a re-evaluation of resource qualifiers for the deployment to ensure that all associated instances are aligned with the current resource constraints and qualifiers defined in the deployment. 
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
  + `name`:(string) The name for the referred blueprint. This name must be unique within the workload definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ref_name`:(string)(ReadOnly) The reference name for the blueprint which is derived by the system. 
  + `resource_constraint`:(HashMap) - The constraints that need to be applied to the resources in order to match this blueprint. 
This complex property has following sub-properties:
    + `input`:(JSON as string) The input values from the user for the resource definition of the blueprint. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `change_details`:(Array)
This complex property has following sub-properties:
  + `change_type`:(string)(ReadOnly) The type of change that was applied.* `None` - No changes have been made.* `InputChange` - A change has been made to an input parameter. For example, an update to the NTP server address.* `WorkloadDefinitionChange` - The associated workload definition has changed, such as updating to a new version.* `WorkloadPreferredVersionChange` - The deployment was created or updated with the default workload definition version, but the default version was later changed. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `conformance`:(string)(ReadOnly) The conformance status of the deployment.* `Ok` - The deployment conforms to the preferred version of the workload.* `NonConformant` - The deployment does not conform to the preferred version of the workload. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `deployment_input`:(HashMap) -(ReadOnly) A reference to a workloadDeploymentInput resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `deployment_input_history`:(Array)(ReadOnly) An array of relationships to workloadDeploymentInput resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) A brief description of the deployment. 
* `digit_count`:(int) The minimum digit count to format the instance index with leading zeros,  for example if the digit count is 4 and the start index is 1, then the  first instance will have a suffix 0001. If the number of instances created for the deployment exceeds the 9999, then the suffix will become a 5 digit number. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `instance_conformance_summary`:(Array)
This complex property has following sub-properties:
  + `nr_count`:(int)(ReadOnly) The total number of referenced objects included in the aggregation. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `state`:(string)(ReadOnly) The overall aggregated state as a string, summarizing the status of all referenced objects. 
* `instance_status_summary`:(Array)
This complex property has following sub-properties:
  + `nr_count`:(int)(ReadOnly) The total number of referenced objects included in the aggregation. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `state`:(string)(ReadOnly) The overall aggregated state as a string, summarizing the status of all referenced objects. 
* `last_action`:(string)(ReadOnly) The last action is performed on the deployment.* `None` - No changes have been made.* `PrepareToDeploy` - Marks the deployment as ready once the user has completed all changes and the deployment is in a valid state. Once the deployment is marked as PrepareToDeploy, workload instances will be created, but the actual deployment will occur as part of the deploy action.* `Deploy` - Initiates the process of pushing workload configuration to the instances based on the configured schedule. Once deployed, the deployment cannot be reverted to draft status.* `Suspend` - Suspends the deployment, preventing any further actions until it is resumed. Making changes to deployment configuration will not be pushed out. The deployment will not take any changes from the attached Workload when it is suspended.* `Retry` - Retries the deployment for all instances that previously failed.* `Resume` - Resumes a suspended deployment. Any changes made to the deployment when it was suspended or any changes made to the attached Workload will now be pushed out based on configured schedules.* `Undeploy` - Undeploy cleans up the policies, templates, and leases generated from the deployment, and marks the deployment as being in the Draft state. The associated definition will be marked as Inactive if there are no other deployments linked to it.* `ReEvaluateQualifiers` - Triggers a re-evaluation of resource qualifiers for the deployment to ensure that all associated instances are aligned with the current resource constraints and qualifiers defined in the deployment. 
* `last_instance_index`:(int)(ReadOnly) Tracks the last numeric index used for workload instances. The next instance index is derived by incrementing this value. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this Deployment. Name can only contain letters (a-z, A-Z), numbers (0-9), space, hyphen (-), or an underscore (_). The name must be unique within the organization. 
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
* `qualification_policies`:(Array)(ReadOnly) An array of relationships to resourceAbstractResourceQualificationPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `qualifiers`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `ref_name`:(string)(ReadOnly) A reference name is generated by the system based on the given name by replacing spaces and hyphen in name with underscore. This reference name is used internally and cannot be edited by users. It may only contain letters (a–z, A–Z), numbers (0–9), and underscores (_). 
* `resource_pool`:(HashMap) -(ReadOnly) A reference to a resourcepoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `rollout_strategy`:(HashMap) - The strategy used for rolling out deployment changes. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workload.BatchDeployment](#workloadBatchDeployment)
[workload.CanaryDeployment](#workloadCanaryDeployment)
  + `failure_threshold`:(int)(ReadOnly) Specifies no of errors can be allowed to skip executing next set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `schedule_policy`:(HashMap) - A reference to a schedulerSchedulePolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_index_for_suffix`:(int) The starting index used to generate the suffix for the workload instance name. 
* `status`:(string)(ReadOnly) The status of the deployment.* `Draft` - The initial state of the deployment, indicating it is still being configured and not ready for deployment.* `ReadyToDeploy` - The deployment has passed validation checks and is ready to be deployed to instances.* `InProgress` - The initial deployment or changes in the deployment are in progress.* `Ok` - All associated instances have been successfully deployed.* `Failed` - One or more instances failed during initial deployment or changes in the deployment operations.* `ChangesScheduled` - The deployment has changes that need to be pushed to all associated instances.* `Suspended` - The deployment is suspended, preventing any further initial deployments or changes to deployment operations. 
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
* `task_schedule`:(HashMap) -(ReadOnly) A reference to a schedulerTaskSchedule resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `validation_information`:(HashMap) -(ReadOnly) Validation details and status of the deployment. 
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
* `workload_definition`:(HashMap) -(ReadOnly) A reference to a workloadWorkloadDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `workload_definition_reference`:(HashMap) - The workload definition associated with this deployment. 
This complex property has following sub-properties:
  + `definition_name`:(string) The name of the definition being referenced. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `organization`:(HashMap) - The organization to which the definition belongs. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `use_preferred_version`:(bool) Indicates whether this version is the default version of the referenced definition. If set to true, the version should be the default. 
  + `nr_version`:(int) The version number of the referenced definition. 
* `workload_instance_prefix`:(string) The prefix to be used for naming workload instances created by this deployment. Prefix can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), or an underscore (_). This prefix must be unique within the organization. 


## Import
`intersight_workload_workload_deployment` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workload_workload_deployment.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [workload.BatchDeployment](#argument-reference)
The number of instances to be included in each batch during the deployment..
* `batch_size`:(int)(ReadOnly) The number of instances to be included in each batch during the deployment. 

### [workload.CanaryDeployment](#argument-reference)
The percentage of instances to which the changes will be rolled out in each step of the deployment.
* `rollout_percentages`:
                (Array of schema.TypeInt) -
  
