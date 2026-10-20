# OpenapiTaskGenerationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.TaskGenerationResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.TaskGenerationResult"]
**FailureReason** | Pointer to **string** | An error message for when task creation fails. | [optional] [readonly] 
**TaskDisplayName** | Pointer to **string** | The display label of the task that is created. | [optional] [readonly] 
**TaskName** | Pointer to **string** | The name of the task that is created. | [optional] [readonly] 
**TaskStatus** | Pointer to **string** | Denotes the status of the task creation. * &#x60;none&#x60; - Indicates the default status * &#x60;InProgress&#x60; - Indicates that operation is in progress * &#x60;Completed&#x60; - Indicates that the operation is complete * &#x60;Failed&#x60; - Indicates that the operation has failed. Check the failureReason attribute for more details. | [optional] [readonly] [default to "none"]
**TaskVersion** | Pointer to **int64** | The version number of the created tasks. | [optional] [readonly] 
**TaskGenerationRequest** | Pointer to [**NullableOpenapiTaskGenerationRequestRelationship**](OpenapiTaskGenerationRequestRelationship.md) |  | [optional] 

## Methods

### NewOpenapiTaskGenerationResult

`func NewOpenapiTaskGenerationResult(classId string, objectType string, ) *OpenapiTaskGenerationResult`

NewOpenapiTaskGenerationResult instantiates a new OpenapiTaskGenerationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiTaskGenerationResultWithDefaults

`func NewOpenapiTaskGenerationResultWithDefaults() *OpenapiTaskGenerationResult`

NewOpenapiTaskGenerationResultWithDefaults instantiates a new OpenapiTaskGenerationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiTaskGenerationResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiTaskGenerationResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiTaskGenerationResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiTaskGenerationResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiTaskGenerationResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiTaskGenerationResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureReason

`func (o *OpenapiTaskGenerationResult) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OpenapiTaskGenerationResult) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OpenapiTaskGenerationResult) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OpenapiTaskGenerationResult) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetTaskDisplayName

`func (o *OpenapiTaskGenerationResult) GetTaskDisplayName() string`

GetTaskDisplayName returns the TaskDisplayName field if non-nil, zero value otherwise.

### GetTaskDisplayNameOk

`func (o *OpenapiTaskGenerationResult) GetTaskDisplayNameOk() (*string, bool)`

GetTaskDisplayNameOk returns a tuple with the TaskDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDisplayName

`func (o *OpenapiTaskGenerationResult) SetTaskDisplayName(v string)`

SetTaskDisplayName sets TaskDisplayName field to given value.

### HasTaskDisplayName

`func (o *OpenapiTaskGenerationResult) HasTaskDisplayName() bool`

HasTaskDisplayName returns a boolean if a field has been set.

### GetTaskName

`func (o *OpenapiTaskGenerationResult) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *OpenapiTaskGenerationResult) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *OpenapiTaskGenerationResult) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *OpenapiTaskGenerationResult) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskStatus

`func (o *OpenapiTaskGenerationResult) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *OpenapiTaskGenerationResult) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *OpenapiTaskGenerationResult) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *OpenapiTaskGenerationResult) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### GetTaskVersion

`func (o *OpenapiTaskGenerationResult) GetTaskVersion() int64`

GetTaskVersion returns the TaskVersion field if non-nil, zero value otherwise.

### GetTaskVersionOk

`func (o *OpenapiTaskGenerationResult) GetTaskVersionOk() (*int64, bool)`

GetTaskVersionOk returns a tuple with the TaskVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskVersion

`func (o *OpenapiTaskGenerationResult) SetTaskVersion(v int64)`

SetTaskVersion sets TaskVersion field to given value.

### HasTaskVersion

`func (o *OpenapiTaskGenerationResult) HasTaskVersion() bool`

HasTaskVersion returns a boolean if a field has been set.

### GetTaskGenerationRequest

`func (o *OpenapiTaskGenerationResult) GetTaskGenerationRequest() OpenapiTaskGenerationRequestRelationship`

GetTaskGenerationRequest returns the TaskGenerationRequest field if non-nil, zero value otherwise.

### GetTaskGenerationRequestOk

`func (o *OpenapiTaskGenerationResult) GetTaskGenerationRequestOk() (*OpenapiTaskGenerationRequestRelationship, bool)`

GetTaskGenerationRequestOk returns a tuple with the TaskGenerationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGenerationRequest

`func (o *OpenapiTaskGenerationResult) SetTaskGenerationRequest(v OpenapiTaskGenerationRequestRelationship)`

SetTaskGenerationRequest sets TaskGenerationRequest field to given value.

### HasTaskGenerationRequest

`func (o *OpenapiTaskGenerationResult) HasTaskGenerationRequest() bool`

HasTaskGenerationRequest returns a boolean if a field has been set.

### SetTaskGenerationRequestNil

`func (o *OpenapiTaskGenerationResult) SetTaskGenerationRequestNil(b bool)`

 SetTaskGenerationRequestNil sets the value for TaskGenerationRequest to be an explicit nil

### UnsetTaskGenerationRequest
`func (o *OpenapiTaskGenerationResult) UnsetTaskGenerationRequest()`

UnsetTaskGenerationRequest ensures that no value is present for TaskGenerationRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


