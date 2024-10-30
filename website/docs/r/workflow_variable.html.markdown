---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_variable"
description: |-
        Variables are user-defined entities that can be shared across workflows. They allow users to set a value once and then reference it from different workflows within the same scope. The variables can be of any type that is supported by the workflow system.

---

# Resource: intersight_workflow_variable
Variables are user-defined entities that can be shared across workflows. They allow users to set a value once and then reference it from different workflows within the same scope. The variables can be of any type that is supported by the workflow system.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `catalog`:(HashMap) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `create_user`:(string)(ReadOnly) The user identifier who created the environment variable. 
* `definition`:(HashMap) - This defines the schema of the variable. It can be set to any Primitive or Custom Data Types supported by workflows. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType)
[workflow.CustomDataType](#workflowCustomDataType)
[workflow.DynamicTemplateParserDataType](#workflowDynamicTemplateParserDataType)
[workflow.MoInventoryDataType](#workflowMoInventoryDataType)
[workflow.MoReferenceAutoDataType](#workflowMoReferenceAutoDataType)
[workflow.MoReferenceDataType](#workflowMoReferenceDataType)
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
    + `widget_type`:(string) Specify the widget type for data display.* `None` - Display none of the widget types.* `Radio` - Display the widget as a radio button.* `Dropdown` - Display the widget as a dropdown.* `GridSelector` - Display the widget as a selector.* `DrawerSelector` - Display the widget as a selector. 
  + `input_parameters`:(JSON as string) JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. 
  + `label`:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ), forward slash (/) or an underscore (_). The first and last character in label must be an alphanumeric character. 
  + `name`:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `required`:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `mod_user`:(string)(ReadOnly) The user identifier who last updated the environment variable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string)(ReadOnly) This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable. 
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
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `value`:(JSON as string) This defines the value of the variable and it will be validated against the datatype defined in the definition. 
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
`intersight_workflow_variable` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_variable.example 1234567890987654321abcde
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
  
