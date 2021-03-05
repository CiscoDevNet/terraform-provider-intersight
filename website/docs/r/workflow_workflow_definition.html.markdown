---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_definition"
description: |-
  Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
---

# Resource: intersight_workflow_workflow_definition
Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
## Argument Reference
The following arguments are supported:
* `catalog`:(HashMap) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `default_version`:(bool) When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version. 
* `description`:(string) The description for this workflow. 
* `input_definition`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType)
[workflow.CustomDataType](#workflowCustomDataType)
[workflow.MoReferenceDataType](#workflowMoReferenceDataType)
[workflow.PrimitiveDataType](#workflowPrimitiveDataType)
[workflow.TargetDataType](#workflowTargetDataType)
  + `default`:(HashMap) - Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `override`:(bool) Override the default value provided for the data type. When true, allow the user to enter value for the data type. 
    + `value`: Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
  + `description`:(string) Provide a detailed description of the data type. 
  + `display_meta`:(HashMap) - Captures the meta data needed for displaying workflow data types in Intersight User Interface. 
This complex property has following sub-properties:
    + `inventory_selector`:(bool) Inventory selector specified for primitive data property should be used in Intersight User Interface. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `widget_type`:(string) Specify the widget type for data display.* `None` - Display none of the widget types.* `Radio` - Display the widget as a radio button.* `Dropdown` - Display the widget as a dropdown.* `GridSelector` - Display the widget as a selector.* `DrawerSelector` - Display the widget as a selector. 
  + `input_parameters`: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. 
  + `label`:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. 
  + `name`:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `required`:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. 
* `input_parameter_set`:(Array)
This complex property has following sub-properties:
  + `condition`:(string) The condition to be evaluated.* `eq` - Checks if the values of the two parameters are equal.* `ne` - Checks if the values of the two parameters are not equal.* `contains` - Checks if the second parameter string value is a substring of the first parameter string value.* `matchesPattern` - Checks if a string matches a regular expression. 
  + `control_parameter`:(string) Name of the controlling entity, whose value will be used for evaluating the parameter set. 
  + `enable_parameters`:
                (Array of schema.TypeString) -
  + `name`:(string) Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) The controlling parameter will be evaluated against this 'value'. 
* `label`:(string) A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `license_entitlement`:(string)(Computed) License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `max_task_count`:(int)(Computed) The maximum number of tasks that can be executed on this workflow. 
* `max_worker_task_count`:(int)(Computed) The maximum number of external (worker) tasks that can be executed on this workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `output_definition`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType)
[workflow.CustomDataType](#workflowCustomDataType)
[workflow.MoReferenceDataType](#workflowMoReferenceDataType)
[workflow.PrimitiveDataType](#workflowPrimitiveDataType)
[workflow.TargetDataType](#workflowTargetDataType)
  + `default`:(HashMap) - Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `override`:(bool) Override the default value provided for the data type. When true, allow the user to enter value for the data type. 
    + `value`: Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
  + `description`:(string) Provide a detailed description of the data type. 
  + `display_meta`:(HashMap) - Captures the meta data needed for displaying workflow data types in Intersight User Interface. 
This complex property has following sub-properties:
    + `inventory_selector`:(bool) Inventory selector specified for primitive data property should be used in Intersight User Interface. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `widget_type`:(string) Specify the widget type for data display.* `None` - Display none of the widget types.* `Radio` - Display the widget as a radio button.* `Dropdown` - Display the widget as a dropdown.* `GridSelector` - Display the widget as a selector.* `DrawerSelector` - Display the widget as a selector. 
  + `input_parameters`: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. 
  + `label`:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. 
  + `name`:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `required`:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. 
* `output_parameters`: The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output. 
* `properties`:(HashMap) - Type to capture the properties of a workflow definition. Some of these properties are passed to workflow execution instance. 
This complex property has following sub-properties:
  + `external_meta`:(bool) When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `retryable`:(bool) When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. 
  + `support_status`:(string) Supported status of the definition.* `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs.* `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported.* `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `tasks`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.DecisionTask](#workflowDecisionTask)
[workflow.FailureEndTask](#workflowFailureEndTask)
[workflow.ForkTask](#workflowForkTask)
[workflow.JoinTask](#workflowJoinTask)
[workflow.LoopTask](#workflowLoopTask)
[workflow.StartTask](#workflowStartTask)
[workflow.SubWorkflowTask](#workflowSubWorkflowTask)
[workflow.SuccessEndTask](#workflowSuccessEndTask)
[workflow.WaitTask](#workflowWaitTask)
[workflow.WorkerTask](#workflowWorkerTask)
  + `description`:(string) The description of this task instance in the workflow. 
  + `label`:(string) A user defined label identifier of the workflow task used for UI display. 
  + `name`:(string) The name of the task within the workflow and it must be unique among all WorkflowTasks within a workflow definition. This name serves as the internal unique identifier for the task and is used to pick input and output parameters to feed into other tasks. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `ui_input_filters`:(Array)
This complex property has following sub-properties:
  + `filters`:
                (Array of schema.TypeString) -
  + `name`:(string) Name for the input definition to which this filter applies. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. When defining the cascade filter for a sub property, use a period (.) to seperate each section of the name like \ StorageConfig.Volume\  where 'StorageConfig' is an input name and 'Volume' is a sub property defined through custom data type definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `user_help_message`:(string) Help message shown to the user about which prior input needs to be selected to enable the input mapped to this filter. 
* `ui_rendering_data`: This will hold the data needed for workflow to be rendered in the user interface. 
* `validation_information`:(HashMap) -(Computed) The current validation state and associated information for this workflow. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `state`:(string)(Computed) The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).* `NotValidated` - The state when workflow definition has not been validated.* `Valid` - The state when workflow definition is valid.* `Invalid` - The state when workflow definition is invalid. 
  + `validation_error`:(Array)
This complex property has following sub-properties:
    + `error_log`:(string)(Computed) Description of the error. 
    + `field`:(string)(Computed) When populated this refers to the input or output field within the workflow or task. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `task_name`:(string)(Computed) The task name on which the error is found, when empty the error applies to the top level workflow. 
    + `transition_name`:(string)(Computed) When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue. 
* `nr_version`:(int) The version of the workflow to support multiple versions. 
* `workflow_metadata`:(HashMap) - A reference to a workflowWorkflowMetadata resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_workflow_workflow_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_workflow_definition.example 1234567890987654321abcde
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
  + `custom_data_type_id`:(string)(Computed) The resolved custom data type definition managed object. 
  + `custom_data_type_name`:(string) Name of the custom data type for this input. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 

### [workflow.MoReferenceDataType](#argument-reference)
Data type to capture an Intersight Managed object reference.
* `properties`:(Array)
This complex property has following sub-properties:
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
  + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 

### [workflow.PrimitiveDataType](#argument-reference)
This data type is used to represent primitives like string, floats and integers.
* `properties`:(HashMap) - Primitive data type properties. 
This complex property has following sub-properties:
  + `constraints`:(HashMap) - Constraints that must be applied to the parameter value supplied for this data type. 
This complex property has following sub-properties:
    + `enum_list`:(Array)
This complex property has following sub-properties:
    + `label`:(string) Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_). 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). 
  + `max`:(float) Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced. 
  + `min`:(float) Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `regex`:(string) When the parameter is a string this regular expression is used to ensure the value is valid. 
  + `inventory_selector`:(Array)
This complex property has following sub-properties:
    + `display_attributes`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
    + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `secure`:(bool) Intersight supports secure properties as task input/output. The values ofthese properties are encrypted and stored in Intersight.This flag marks the property to be secure when it is set to true. 
  + `type`:(string) Specify the enum type for primitive data type.* `string` - Enum to specify a string data type.* `integer` - Enum to specify an integer32 data type.* `float` - Enum to specify a float64 data type.* `boolean` - Enum to specify a boolean data type.* `json` - Enum to specify a json data type.* `enum` - Enum to specify a enum data type which is a list of pre-defined strings. 

### [workflow.TargetDataType](#argument-reference)
Data type to capture a target endpoint or device.
* `custom_data_type_properties`:(HashMap) - Reference to custom data type definition. 
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this custom data type belongs. 
  + `custom_data_type_id`:(string)(Computed) The resolved custom data type definition managed object. 
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
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. 
  + `supported_objects`:
                (Array of schema.TypeString) -
  
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
  + `custom_data_type_id`:(string)(Computed) The resolved custom data type definition managed object. 
  + `custom_data_type_name`:(string) Name of the custom data type for this input. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 

### [workflow.MoReferenceDataType](#argument-reference)
Data type to capture an Intersight Managed object reference.
* `properties`:(Array)
This complex property has following sub-properties:
  + `display_attributes`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
  + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 

### [workflow.PrimitiveDataType](#argument-reference)
This data type is used to represent primitives like string, floats and integers.
* `properties`:(HashMap) - Primitive data type properties. 
This complex property has following sub-properties:
  + `constraints`:(HashMap) - Constraints that must be applied to the parameter value supplied for this data type. 
This complex property has following sub-properties:
    + `enum_list`:(Array)
This complex property has following sub-properties:
    + `label`:(string) Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_). 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). 
  + `max`:(float) Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced. 
  + `min`:(float) Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `regex`:(string) When the parameter is a string this regular expression is used to ensure the value is valid. 
  + `inventory_selector`:(Array)
This complex property has following sub-properties:
    + `display_attributes`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options. 
    + `value_attribute`:(string) A property from the Intersight object, value of which can be used as value for referenced input definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `secure`:(bool) Intersight supports secure properties as task input/output. The values ofthese properties are encrypted and stored in Intersight.This flag marks the property to be secure when it is set to true. 
  + `type`:(string) Specify the enum type for primitive data type.* `string` - Enum to specify a string data type.* `integer` - Enum to specify an integer32 data type.* `float` - Enum to specify a float64 data type.* `boolean` - Enum to specify a boolean data type.* `json` - Enum to specify a json data type.* `enum` - Enum to specify a enum data type which is a list of pre-defined strings. 

### [workflow.TargetDataType](#argument-reference)
Data type to capture a target endpoint or device.
* `custom_data_type_properties`:(HashMap) - Reference to custom data type definition. 
This complex property has following sub-properties:
  + `catalog_moid`:(string) Specify the catalog moid that this custom data type belongs. 
  + `custom_data_type_id`:(string)(Computed) The resolved custom data type definition managed object. 
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
  + `selector`:(string) Field to hold an Intersight API along with an optional filter to narrow down the search options for target device. 
  + `supported_objects`:
                (Array of schema.TypeString) -
  
### [workflow.DecisionTask](#argument-reference)
A DecisionTask is a control task that executes a sequence of WorkflowTasks based off decision provided and evaluated by this task.
* `condition`:(string) The condition to evaluate for this decision task. The condition can be a workflow or task variable or an JavaScript expression. Example value for condition could be a variable like \ ${task1.output.var1} or ${workflow.input.var2}\  which evaluates to a value matching any of the decision case values. Example value for condition if it's an expression - \ if ( ${task1.output.var1} ! = null && ${task1.output.var1} > 0 ) 'true'; else 'false'; \  which evaluates to 'true' or 'false' and will match one of the decision case values. You can also use JavaScript functions like indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator. 
* `decision_cases`:(Array)
This complex property has following sub-properties:
  + `description`:(string) Description of this decision case. 
  + `next_task`:(string) The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) Value for the decision case. 
* `default_task`:(string) The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases. 
* `input_parameters`: This field is deprecated. Decision case conditions can be added using the workflow input or task output variables in the Condition field. Refer to Condition field for more details. 

### [workflow.FailureEndTask](#argument-reference)
A FailureEndTask denotes the failed completion of a workflow.

### [workflow.ForkTask](#argument-reference)
A ForkTask is a control task that forks tasks for parallel execution in a workflow.
* `forked_tasks`:
                (Array of schema.TypeString) -
* `join_task`:(string) Task name for the join control task that must follow a fork control task. 

### [workflow.JoinTask](#argument-reference)
A JoinTask is a control task that must follow a fork task and specify all the fork tasks that must complete and join before the worfklow can proceed to the task specified in the OnSuccess transition.
* `join_on_tasks`:
                (Array of schema.TypeString) -
* `on_success`:(string) Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 

### [workflow.LoopTask](#argument-reference)
A LoopTask is a control task that runs one or more task multiple times based on a counter. The tasks can be run in serial or parallel.
* `nr_count`:(string) Count value for the loop, this can be a constant or an expression which will evaluate to an integer value. Example, Use the length of the input A which is an array. Count must be less than or equal to 500. 
* `loop_start_task`:(string) Start task where the list of tasks will be executed multiple times based on the count value. 
* `number_of_batches`:(int) When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them. 
* `on_success`:(string) This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node. 
* `parallel`:(bool) When set to true the loop will run in parallel else it will run in a serial fashion. Only one task is supported inside the loop task when the loop is run in parallel. Subworkflow can be used inside the single loop task to build complex conditions. 

### [workflow.StartTask](#argument-reference)
A StartTask is the starting point for a workflow.  There can only be one StartTask in a workflow.
* `next_task`:(string) The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node. 

### [workflow.SubWorkflowTask](#argument-reference)
A SubWorkflowTask is used to include another workflow as a task within this workflow.
* `catalog_moid`:(string) Specify the catalog moid that this task belongs. 
* `input_parameters`: JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output. 
* `on_failure`:(string) This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `on_success`:(string) This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `rollback_disabled`:(bool) The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. 
* `use_default`:(bool) UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution. 
* `nr_version`:(int) The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. 
* `workflow_definition_id`:(string)(Computed) The resolved referenced workflow definition managed object. 
* `workflow_definition_name`:(string) The qualified name of workflow that should be executed as a task. 

### [workflow.SuccessEndTask](#argument-reference)
A SuccessEndTask denotes the successful completion of a workflow.

### [workflow.WaitTask](#argument-reference)
A WaitTask will remain in progress until marked success or failed by an external trigger. The timeout for wait task is 180 days, so a workflow can wait for task status update for upto 180 days. Currently the only supported means to update the task status is through an API that provides the status of the task runtime instance. Once the wait task status has been set the workflow will continue with the execution based on the task status.
* `input_parameters`: JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output. 
* `on_failure`:(string) This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `on_success`:(string) This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `prompts`:(Array)
This complex property has following sub-properties:
  + `description`:(string) Description that give more details about what it means to pick this prompt option for the wait task. 
  + `label`:(string) User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task. 
  + `name`:(string) Name for the wait prompt. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `task_status`:(string) Task status for the wait task when this prompt option is selected.* `Scheduled` - The enum represents the status when task is in scheduled state.* `InProgress` - The enum represents the status when task is in-progress state.* `NoOp` - The enum represents the status when task is a noop.* `Timeout` - The enum represents the status when task has timed out.* `Completed` - The enum represents the status when task has completed.* `Failed` - The enum represents the status when task has failed. 
* `rollback_disabled`:(bool) The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. 
* `use_default`:(bool) UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution. 

### [workflow.WorkerTask](#argument-reference)
A WorkerTask is a simple task and the smallest granularity of work that can be defined as a task.
* `catalog_moid`:(string) Specify the catalog moid that this task belongs. 
* `input_parameters`: JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output. 
* `on_failure`:(string) This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `on_success`:(string) This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. 
* `rollback_disabled`:(bool) The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. 
* `task_definition_id`:(string)(Computed) The resolved referenced task definition managed object. 
* `task_definition_name`:(string) The qualified name of task that should be executed. 
* `use_default`:(bool) UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution. 
* `nr_version`:(int) The task definition version to use in this workflow. When no version is specified then the default version of the task at the time of creating or updating this workflow is used. 
  