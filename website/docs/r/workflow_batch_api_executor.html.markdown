
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
  + `asset_target_moid`:(string)(Computed) Asset target defines the remote target endpoints which are either managed byIntersight or their service APIs are invoked from Intersight. Generic API executorservice Jasmine can invoke these remote APIs via its executors. Asset targets can beaccessed directly for cloud targets or via an associated Intersight Assist. Prerequisiteto use asset targets is to persist the target details.Asset target MoRef can be given as task input of type TargetDataType. Fusion determinesand populates the target context with the Assist if associated with. It is setinternally at the API level.In case of an associated assist, it is used by Assist to fetch the target detailsand form the API request to send to endpoints. In case of cloud asset targets, Jasminefetched the target details from DB, forms the API request and sends it to the target. 
  + `body`:(string) The optional request body that is sent as part of this API request.The request body can contain a golang template that can be populated with task inputparameters and previous API output parameters. 
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `content_type`:(string) Intersight Orchestrator, with the support of response parser specification,can extract the values from API responses and map them to task output parameters.The value extraction is supported for response content types XML and JSON.The type of the content that gets passed as payload and response in thisAPI. 
  + `description`:(string) A description that task designer can add to individual API requests that explain what the API call is about. 
  + `label`:(string) A user friendly label that task designers have given to the batch API request. 
  + `name`:(string) A reference name for this API request within the batch API request.This name shall be used to map the API output parameters to subsequentAPI input parameters within a batch API task. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `outcomes`: All the possible outcomes of this API are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success. 
  + `response_spec`: The optional grammar specification for parsing the response to extract therequired values.The specification should have extraction specification specified forall the API output parameters. 
  + `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip theapi execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
  + `start_delay`:(int) The delay in seconds after which the API needs to be executed.By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution.Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out. 
  + `timeout`:(int) The duration in seconds by which the API response is expected from the API target.If the end point does not respond for the API request within this timeoutduration, the task will be marked as failed. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `constraints`:(Array with Maximum of one item) - Enter the constraints on when this task should match against the task definition. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `target_data_type`: List of property constraints that helps to narrow down task implementations based on target device input. 
* `description`:(string) A detailed description about the batch APIs. 
* `error_response_handler`:(Array with Maximum of one item) - A reference to a workflowErrorResponseHandler resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the batch API task. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
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
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
