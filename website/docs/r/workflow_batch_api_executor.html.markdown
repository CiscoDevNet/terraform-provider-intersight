
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_batch_api_executor"
sidebar_current: "docs-intersight-resource-workflow-batch-api-executor"
description: |-
  Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Batch API associates the list of API requests to be executed as part of single
task execution. Each API request takes the request body and a response parser
specification.
---

# Resource: intersight_workflow._batch_api_executor
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Batch API associates the list of API requests to be executed as part of single
task execution. Each API request takes the request body and a response parser
specification.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `batch`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `body`:(string) The optional request body that is sent as part of this API request.The request body can contain a golang template that can be populated with task inputparameters and previous API output parameters. 
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `content_type`:(string) Intersight Orchestrator, with the support of response parser specification,can extract the values from API responses and map them to task output parameters.The value extraction is supported for response content types XML and JSON.The type of the content that gets passed as payload and response in thisAPI.* `json` - The type of content to be parsed, API response or device response, is inJSON format.* `xml` - The type of content to be parsed, API response or device response,is in XML format.* `text` - The type of content to be parsed, API response or device response, is inTEXT format. 
  + `name`:(string) A reference name for this API request within the batch API request.This name shall be used to map the API output parameters to subsequentAPI input parameters within a batch API task. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `outcomes`: All the possible outcomes of this API are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success. 
  + `response_spec`:(Array with Maximum of one item) - The optional grammar specification for parsing the response to extract therequired values.The specification should have extraction specification specified forall the API output parameters. 
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `error_parameters`:(Array)
This complex property has following sub-properties:
    + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
    + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
    + `name`:(string) The name of the parameter. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
    + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `parameters`:(Array)
This complex property has following sub-properties:
    + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
    + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
    + `name`:(string) The name of the parameter. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
    + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
  + `types`:(Array)
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `name`:(string) The unique name of this complex type within the grammar specification. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `parameters`:(Array)
This complex property has following sub-properties:
    + `accept_single_value`:(bool) The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only. 
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `complex_type`:(string) The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property. 
    + `item_type`:(string) The type of the collection item in case this is a collection parameter.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
    + `name`:(string) The name of the parameter. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `path`:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. 
    + `type`:(string) The type of the parameter. Accepted values are simple, complex,collection.* `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.* `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.* `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. 
  + `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip theapi execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
  + `start_delay`:(int) The delay in seconds after which the API needs to be executed.By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution.Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out. 
  + `timeout`:(int) The duration in seconds by which the API response is expected from the API target.If the end point does not respond for the API request within this timeoutduration, the task will be marked as failed. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `constraints`:(Array with Maximum of one item) - Enter the constraints on when this task should match against the task definition. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `target_data_type`: List of property constraints that helps to narrow down task implementations based on target device input. 
* `description`:(string) A detailed description about the batch APIs. 
* `error_response_handler`:(Array with Maximum of one item) - A reference to a workflowErrorResponseHandler resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the batch API task. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `outcomes`: All the possible outcomes of this task are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success. 
* `output`: Intersight Orchestrator allows the extraction of required values from APIresponses using the API response grammar. These extracted values can be mappedto task output parameters defined in task definition.The mapping of API output parameters to the task output parameters is providedas JSON in this property. 
* `retry_from_failed_api`:(bool) When an execution of a nth API in the Batch fails,Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry.By default the value is set to false. 
* `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip thetask execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `task_definition`:(Array with Maximum of one item) - A reference to a workflowTaskDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
