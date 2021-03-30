# WorkflowApiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AssetTargetMoid** | Pointer to **string** | Asset target defines the remote target endpoints which are either managed by Intersight or their service APIs are invoked from Intersight. Generic API executor service Jasmine can invoke these remote APIs via its executors. Asset targets can be accessed directly for cloud targets or via an associated Intersight Assist. Prerequisite to use asset targets is to persist the target details. Asset target MoRef can be given as task input of type TargetDataType. Fusion determines and populates the target context with the Assist if associated with. It is set internally at the API level. In case of an associated assist, it is used by Assist to fetch the target details and form the API request to send to endpoints. In case of cloud asset targets, Jasmine fetched the target details from DB, forms the API request and sends it to the target. | [optional] [readonly] 
**Body** | Pointer to **string** | The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters. | [optional] 
**ContentType** | Pointer to **string** | Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. The type of the content that gets passed as payload and response in this API. The supported values are json, xml, text. | [optional] 
**Description** | Pointer to **string** | A description that task designer can add to individual API requests that explain  what the API call is about. | [optional] 
**ErrorContentType** | Pointer to **string** | Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. Optional input to specify the content type in case of error API response. This should be used if the content type of error response is different from that of the success response. If not specified, contentType input value is used to parse the error response. | [optional] 
**Label** | Pointer to **string** | A user friendly label that task designers have given to the batch API request. | [optional] 
**Name** | Pointer to **string** | A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**ResponseSpec** | Pointer to **interface{}** | The optional grammar specification for parsing the response to extract the required values. The specification should have extraction specification specified for all the API output parameters. | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 
**StartDelay** | Pointer to **int64** | The delay in seconds after which the API needs to be executed. By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution. Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out. | [optional] 
**Timeout** | Pointer to **int64** | The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed. | [optional] 

## Methods

### NewWorkflowApiAllOf

`func NewWorkflowApiAllOf(classId string, objectType string, ) *WorkflowApiAllOf`

NewWorkflowApiAllOf instantiates a new WorkflowApiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowApiAllOfWithDefaults

`func NewWorkflowApiAllOfWithDefaults() *WorkflowApiAllOf`

NewWorkflowApiAllOfWithDefaults instantiates a new WorkflowApiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowApiAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowApiAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowApiAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowApiAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowApiAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowApiAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetTargetMoid

`func (o *WorkflowApiAllOf) GetAssetTargetMoid() string`

GetAssetTargetMoid returns the AssetTargetMoid field if non-nil, zero value otherwise.

### GetAssetTargetMoidOk

`func (o *WorkflowApiAllOf) GetAssetTargetMoidOk() (*string, bool)`

GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTargetMoid

`func (o *WorkflowApiAllOf) SetAssetTargetMoid(v string)`

SetAssetTargetMoid sets AssetTargetMoid field to given value.

### HasAssetTargetMoid

`func (o *WorkflowApiAllOf) HasAssetTargetMoid() bool`

HasAssetTargetMoid returns a boolean if a field has been set.

### GetBody

`func (o *WorkflowApiAllOf) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WorkflowApiAllOf) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WorkflowApiAllOf) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WorkflowApiAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *WorkflowApiAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkflowApiAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkflowApiAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WorkflowApiAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowApiAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowApiAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowApiAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowApiAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorContentType

`func (o *WorkflowApiAllOf) GetErrorContentType() string`

GetErrorContentType returns the ErrorContentType field if non-nil, zero value otherwise.

### GetErrorContentTypeOk

`func (o *WorkflowApiAllOf) GetErrorContentTypeOk() (*string, bool)`

GetErrorContentTypeOk returns a tuple with the ErrorContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorContentType

`func (o *WorkflowApiAllOf) SetErrorContentType(v string)`

SetErrorContentType sets ErrorContentType field to given value.

### HasErrorContentType

`func (o *WorkflowApiAllOf) HasErrorContentType() bool`

HasErrorContentType returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowApiAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowApiAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowApiAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowApiAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowApiAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowApiAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowApiAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowApiAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowApiAllOf) GetOutcomes() interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowApiAllOf) GetOutcomesOk() (*interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowApiAllOf) SetOutcomes(v interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowApiAllOf) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *WorkflowApiAllOf) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *WorkflowApiAllOf) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetResponseSpec

`func (o *WorkflowApiAllOf) GetResponseSpec() interface{}`

GetResponseSpec returns the ResponseSpec field if non-nil, zero value otherwise.

### GetResponseSpecOk

`func (o *WorkflowApiAllOf) GetResponseSpecOk() (*interface{}, bool)`

GetResponseSpecOk returns a tuple with the ResponseSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSpec

`func (o *WorkflowApiAllOf) SetResponseSpec(v interface{})`

SetResponseSpec sets ResponseSpec field to given value.

### HasResponseSpec

`func (o *WorkflowApiAllOf) HasResponseSpec() bool`

HasResponseSpec returns a boolean if a field has been set.

### SetResponseSpecNil

`func (o *WorkflowApiAllOf) SetResponseSpecNil(b bool)`

 SetResponseSpecNil sets the value for ResponseSpec to be an explicit nil

### UnsetResponseSpec
`func (o *WorkflowApiAllOf) UnsetResponseSpec()`

UnsetResponseSpec ensures that no value is present for ResponseSpec, not even an explicit nil
### GetSkipOnCondition

`func (o *WorkflowApiAllOf) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowApiAllOf) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowApiAllOf) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowApiAllOf) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.

### GetStartDelay

`func (o *WorkflowApiAllOf) GetStartDelay() int64`

GetStartDelay returns the StartDelay field if non-nil, zero value otherwise.

### GetStartDelayOk

`func (o *WorkflowApiAllOf) GetStartDelayOk() (*int64, bool)`

GetStartDelayOk returns a tuple with the StartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelay

`func (o *WorkflowApiAllOf) SetStartDelay(v int64)`

SetStartDelay sets StartDelay field to given value.

### HasStartDelay

`func (o *WorkflowApiAllOf) HasStartDelay() bool`

HasStartDelay returns a boolean if a field has been set.

### GetTimeout

`func (o *WorkflowApiAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WorkflowApiAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WorkflowApiAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WorkflowApiAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


