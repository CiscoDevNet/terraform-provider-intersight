---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_action_definition"
description: |-
        Definition to capture the details needed to execute an action on the service item.

---

# Resource: intersight_workflow_service_item_action_definition
Definition to capture the details needed to execute an action on the service item.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action_properties`:(HashMap) - Action properties for the service item. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `operation_type`:(string) Type of action operation to be executed on the service item.* `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item.* `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment.* `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.* `Migration` - This represents the migration action, used to migrate service item instance from one service item definition version to another service item definition version. There can be more than one migration operations associated with a service item. Once a migration action is running on a service item instance, no further operations are allowed on that service item instance during the migration process. 
  + `properties`:(HashMap) - The properties of the action. The actual structure of properties can vary based on the operationType. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `stop_on_failure`:(bool) When true, the action on the service item will be stopped when it reaches a failure by either calling the configured stop workflow or by calling the rollback workflow. By default value is set to true. 
* `action_type`:(string) Type of actionDefinition which decides on how to trigger the action.* `External` - External actions definition can be triggered by enduser to perform actions on the service item. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again.* `Internal` - Internal action definition is used to trigger periodic actions on the service item instance.* `Repetitive` - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created service item. 
* `allowed_instance_states`:
                (Array of schema.TypeString) -
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_roles`:(Array) An array of relationships to iamRole resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `attribute_parameters`:(JSON as string) The mappings from workflows in the action definition to the service item attribute definition. Any output from core or post-core workflow can be mapped to service item attribute definition. The attribute can be referred using the name of the workflow definition and the attribute name in the following format '${<ServiceItemActionWorkflowDefinition.Name>.output.<outputName>'. 
* `core_workflows`:(Array)
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. 
  + `description`:(string) The description of this workflow instance. 
  + `input_parameters`:(JSON as string) Capture the mapping of ActionDefinition inputDefinition to workflow definition. 
  + `label`:(string) A user defined label identifier of the workflow used for UI display. 
  + `name`:(string) The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
  + `workflow_definition`:(HashMap) -(ReadOnly) The moid of the workflow definition. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `workflow_definition_name`:(string) The qualified name of workflow that should be executed. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) The description for this action which provides information on what are the pre-requisites to use this action on the service item and what features are supported by this action. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `input_definition`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType)
[workflow.CustomDataType](#workflowCustomDataType)
[workflow.DynamicTemplateParserDataType](#workflowDynamicTemplateParserDataType)
[workflow.MoInventoryDataType](#workflowMoInventoryDataType)
[workflow.MoReferenceAutoDataType](#workflowMoReferenceAutoDataType)
[workflow.MoReferenceDataType](#workflowMoReferenceDataType)
[workflow.MoReferenceForCloneDataType](#workflowMoReferenceForCloneDataType)
[workflow.PrimitiveDataType](#workflowPrimitiveDataType)
[workflow.TargetDataType](#workflowTargetDataType)
  + `default`:(HashMap) - Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
This complex property has following sub-properties:
    + `is_value_set`:(bool)(ReadOnly) A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `override`:(bool) Override the default value provided for the data type. When true, allow the user to enter value for the data type. 
    + `value`:(JSON as string) Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
  + `description`:(string) Provide a detailed description of the data type. 
  + `display_meta`:(HashMap) - Captures the meta data needed for displaying workflow data types in Intersight User Interface. 
This complex property has following sub-properties:
    + `inventory_selector`:(bool) Inventory selector specified for primitive data property should be used in Intersight User Interface. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `widget_type`:(string) Specify the widget type for data display.* `None` - Display none of the widget types.* `Radio` - Display the widget as a radio button.* `Dropdown` - Display the widget as a dropdown.* `GridSelector` - Display the widget as a selector.* `DrawerSelector` - Display the widget as a selector.* `MultiSelect` - Display the widget as a multi-select. 
  + `input_parameters`:(JSON as string) JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. 
  + `label`:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ), forward slash (/) or an underscore (_). The first and last character in label must be an alphanumeric character. 
  + `name`:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `required`:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. 
* `label`:(string) A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a service item definition. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `periodicity`:(int) Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `post_core_workflows`:(Array)
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. 
  + `description`:(string) The description of this workflow instance. 
  + `input_parameters`:(JSON as string) Capture the mapping of ActionDefinition inputDefinition to workflow definition. 
  + `label`:(string) A user defined label identifier of the workflow used for UI display. 
  + `name`:(string) The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
  + `workflow_definition`:(HashMap) -(ReadOnly) The moid of the workflow definition. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `workflow_definition_name`:(string) The qualified name of workflow that should be executed. 
* `pre_core_workflows`:(Array)
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. 
  + `description`:(string) The description of this workflow instance. 
  + `input_parameters`:(JSON as string) Capture the mapping of ActionDefinition inputDefinition to workflow definition. 
  + `label`:(string) A user defined label identifier of the workflow used for UI display. 
  + `name`:(string) The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
  + `workflow_definition`:(HashMap) -(ReadOnly) The moid of the workflow definition. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `workflow_definition_name`:(string) The qualified name of workflow that should be executed. 
* `restrict_on_private_appliance`:(bool) The flag to indicate that action is restricted on a Private Virtual Appliance. 
* `service_item_definition`:(HashMap) - A reference to a workflowServiceItemDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `stop_workflows`:(Array)
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. 
  + `description`:(string) The description of this workflow instance. 
  + `input_parameters`:(JSON as string) Capture the mapping of ActionDefinition inputDefinition to workflow definition. 
  + `label`:(string) A user defined label identifier of the workflow used for UI display. 
  + `name`:(string) The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
  + `workflow_definition`:(HashMap) -(ReadOnly) The moid of the workflow definition. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `workflow_definition_name`:(string) The qualified name of workflow that should be executed. 
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
* `user_id_or_email`:(string)(ReadOnly) The user identifier who created or updated the service item action definition. 
* `validation_information`:(HashMap) -(ReadOnly) The current validation state and associated validation errors when state is invalid. 
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
* `validation_workflows`:(Array)
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. 
  + `description`:(string) The description of this workflow instance. 
  + `input_parameters`:(JSON as string) Capture the mapping of ActionDefinition inputDefinition to workflow definition. 
  + `label`:(string) A user defined label identifier of the workflow used for UI display. 
  + `name`:(string) The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
  + `workflow_definition`:(HashMap) -(ReadOnly) The moid of the workflow definition. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `workflow_definition_name`:(string) The qualified name of workflow that should be executed. 
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
* `workflow_definition`:(HashMap) -(ReadOnly) A reference to a workflowWorkflowDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_workflow_service_item_action_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_service_item_action_definition.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [workflow.ArrayDataType](#argument-reference)
This data type represents an array of a given type. It can be an array of primitive data or of custom data.
* `array_item_type`:(HashMap) - Data item within the array data type. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `max`:(int) Specify the maximum value of the array. 
* `min`:(int) Specify the minimum value of the array. 

### [workflow.CustomDataType](#argument-reference)
This data type represents a custom data object.
* `properties`:(HashMap) - Captures the custom data type properties. 
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this custom data type belongs. 
  + `custom_data_type_id`:(string)(ReadOnly) The resolved custom data type definition managed object. 
  + `custom_data_type_name`:(string) Name of the custom data type for this input. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 

### [workflow.DynamicTemplateParserDataType](#argument-reference)
Data type to fetch a generic template from given selector and parse it using an api to give an array of secure and non-secure keys for form generation. URL used to fetch the template object is based on the templateType. Final input passed to the workflow using this data type is a JSON containing {'Template':'<template string value>', 'Keys':[{'<key1>':'<val 1>'}], 'SecureKeys':[{'<key2>':'<val2>'}]}.
* `is_template_secure`:(bool) When set to true, the template is marked as secure and the content is encrypted and stored. 
* `template_type`:(string) Template type decides on the API to be used to fetch the placeholders present inside the template.* `OsInstall` - This refers to the OS configuration template MO. 

### [workflow.MoInventoryDataType](#argument-reference)
The data type to represent the selected properties of an Intersight managed object. This data type is used only in Service items to define the schema of resources and their attributes.
* `properties`:(Array)
This complex property has following sub-properties:
  + `attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `reference_object_type`:(string) ObjectType for which the attributes need to be collected. 
  + `reference_type`:(string) Defines how the reference to the shadow resource is done. Base case is via an Moid, but reference via a selector is also possible.* `Moid` - The reference to the original resource is via an Moid.* `Selector` - The reference to the original resource is via a selector query. This can potentially lead to tracking data for multiple resources. 

### [workflow.MoReferenceAutoDataType](#argument-reference)
The data type to capture an Intersight managed object reference that is automatically selected by the system based on a given selection criteria.
* `properties`:(Array)
This complex property has following sub-properties:
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `filters`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `order_by`:(string) Determines  properties that are used to sort the collection of resources. 
  + `rule`:(HashMap) - The rule can be obtained directly from a Resource Selection Criteria or provided as inline resource filter conditions. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 

### [workflow.MoReferenceDataType](#argument-reference)
Data type to capture an Intersight Managed object reference.
* `properties`:(Array)
This complex property has following sub-properties:
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
  + `selector_property`:(HashMap) - Selector properties to define HTTP method and 'body' in case of upsert operation. 
This complex property has following sub-properties:
    + `body`:(JSON as string) Content of the request body to send for POST request. 
    + `method`:(string) The HTTP method to be used.* `GET` - The HTTP GET method requests a representation of the specified resource.* `POST` - The HTTP POST method sends data to the server. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 

### [workflow.MoReferenceForCloneDataType](#argument-reference)
Data type captures a reference to an Intersight Managed Object (MO) that serves as the source for cloning and creating an input object. The fields to be passed as for this datatype will be SourceMoid for the Source MO and the ObjectType which is the qualified name for the MO. When a clone is created, the Moid field of the new object is populated with the cloned object ID. After the cloning process is complete, the SourceMoid field is removed, leaving only Moid and ObjectType in the input. The identity properties of the cloned object are assigned by the consuming module. For instance, when used in Workload definitions to clone a LDAP policy, the policy identity — such as its Name — is determined based on the Workload definition name combined with the source policy name.
* `properties`:(Array)
This complex property has following sub-properties:
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
  + `selector_property`:(HashMap) - Selector properties to define HTTP method and 'body' in case of upsert operation. 
This complex property has following sub-properties:
    + `body`:(JSON as string) Content of the request body to send for POST request. 
    + `method`:(string) The HTTP method to be used.* `GET` - The HTTP GET method requests a representation of the specified resource.* `POST` - The HTTP POST method sends data to the server. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 

### [workflow.PrimitiveDataType](#argument-reference)
This data type is used to represent primitives like string, floats and integers.
* `properties`:(HashMap) - Primitive data type properties. 
This complex property has following sub-properties:
  + `constraints`:(HashMap) - Constraints that must be applied to the parameter value supplied for this data type. 
This complex property has following sub-properties:
    + `enum_list`:(Array)
This complex property has following sub-properties:
    + `label`:(string) Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), round parenthesis ( () ), or an underscore (_) and must have an alphanumeric character. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), round parenthesis ( () ), forward slash (/), or an underscore (_). 
  + `max`:(float) Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced. 
  + `min`:(float) Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `regex`:(string) When the parameter is a string this regular expression is used to ensure the value is valid. 
  + `data_source_selector`:(Array)
This complex property has following sub-properties:
    + `display_attributes`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `selector`:(string) This field holds mapping information used to provide suggestions to the user. The mapping should be in the '${workflow.input.property}' format. It supports workflow input mapping for workflows, and for User Actions, it supports workflow inputs, workflow outputs, workflow variables, and outputs from previous tasks. 
    + `value_attribute`:(string) A property from the mapped parameter, value of which can be used as value for referenced input definition. 
  + `inventory_selector`:(Array)
This complex property has following sub-properties:
    + `display_attributes`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
    + `selector_property`:(HashMap) - Selector properties to define HTTP method and 'body' in case of upsert operation. 
This complex property has following sub-properties:
    + `body`:(JSON as string) Content of the request body to send for POST request. 
    + `method`:(string) The HTTP method to be used.* `GET` - The HTTP GET method requests a representation of the specified resource.* `POST` - The HTTP POST method sends data to the server. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `secure`:(bool) Intersight supports secure properties as task input/output. The values ofthese properties are encrypted and stored in Intersight.This flag marks the property to be secure when it is set to true. 
  + `type`:(string) Specify the enum type for primitive data type.* `string` - Enum to specify a string data type.* `integer` - Enum to specify an integer32 data type.* `float` - Enum to specify a float64 data type.* `boolean` - Enum to specify a boolean data type.* `json` - Enum to specify a json data type.* `enum` - Enum to specify a enum data type which is a list of pre-defined strings. 

### [workflow.TargetDataType](#argument-reference)
Data type to capture a target endpoint or device.
* `custom_data_type_properties`:(HashMap) - Reference to custom data type definition. 
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this custom data type belongs. 
  + `custom_data_type_id`:(string)(ReadOnly) The resolved custom data type definition managed object. 
  + `custom_data_type_name`:(string) Name of the custom data type for this input. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `is_array`:(bool) When this property is true then an array of targets can be passed as input. 
* `max`:(int) Specify the maximum value of the array. 
* `min`:(int) Specify the minimum value of the array. 
* `properties`:(Array)
This complex property has following sub-properties:
  + `connector_attribute`:(string) A singleton value which will contain the path to connector object from the selected object. 
  + `constraint_attributes`:
                (Array of schema.TypeString) -
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `platform_type`:
                (Array of schema.TypeString) -
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. 
  + `selector_property`:(HashMap) - Selector properties to define HTTP method and 'body' in case of upsert operation. 
This complex property has following sub-properties:
    + `body`:(JSON as string) Content of the request body to send for POST request. 
    + `method`:(string) The HTTP method to be used.* `GET` - The HTTP GET method requests a representation of the specified resource.* `POST` - The HTTP POST method sends data to the server. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `supported_objects`:
                (Array of schema.TypeString) -
  
