---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_custom_data_type_definition"
description: |-
  Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
---

# Resource: intersight_workflow_custom_data_type_definition
Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
## Argument Reference
The following arguments are supported:
* `catalog`:(HashMap) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `composite_type`:(bool) When true this data type definition is a collection of type definitions to represent composite data like JSON. 
* `description`:(string) A human-friendly description of this custom data type indicating it's domain and usage. 
* `label`:(string) A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'. 
* `parameter_set`:(Array)
This complex property has following sub-properties:
  + `condition`:(string) The condition to be evaluated.* `eq` - Checks if the values of the two parameters are equal.* `ne` - Checks if the values of the two parameters are not equal.* `contains` - Checks if the second parameter string value is a substring of the first parameter string value.* `matchesPattern` - Checks if a string matches a regular expression. 
  + `control_parameter`:(string) Name of the controlling entity, whose value will be used for evaluating the parameter set. 
  + `enable_parameters`:
                (Array of schema.TypeString) -
  + `name`:(string) Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) The controlling parameter will be evaluated against this 'value'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `type_definition`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType)
[workflow.CustomDataType](#workflowCustomDataType)
[workflow.MoReferenceDataType](#workflowMoReferenceDataType)
[workflow.PrimitiveDataType](#workflowPrimitiveDataType)
[workflow.TargetDataType](#workflowTargetDataType)
  + `default`:(HashMap) - Default value for the data type. If default value was provided and the input was required the default value will be used as the input. 
This complex property has following sub-properties:
    + `is_value_set`:(bool)(Computed) A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses. 
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


## Import
`intersight_workflow_custom_data_type_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_custom_data_type_definition.example 1234567890987654321abcde
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
  