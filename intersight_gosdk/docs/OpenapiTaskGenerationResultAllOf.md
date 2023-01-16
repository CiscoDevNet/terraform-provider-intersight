# OpenapiTaskGenerationResultAllOf

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
**TaskGenerationRequest** | Pointer to [**OpenapiTaskGenerationRequestRelationship**](OpenapiTaskGenerationRequestRelationship.md) |  | [optional] 

## Methods

### NewOpenapiTaskGenerationResultAllOf

`func NewOpenapiTaskGenerationResultAllOf(classId string, objectType string, ) *OpenapiTaskGenerationResultAllOf`

NewOpenapiTaskGenerationResultAllOf instantiates a new OpenapiTaskGenerationResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiTaskGenerationResultAllOfWithDefaults

`func NewOpenapiTaskGenerationResultAllOfWithDefaults() *OpenapiTaskGenerationResultAllOf`

NewOpenapiTaskGenerationResultAllOfWithDefaults instantiates a new OpenapiTaskGenerationResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiTaskGenerationResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiTaskGenerationResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiTaskGenerationResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiTaskGenerationResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiTaskGenerationResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiTaskGenerationResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureReason

`func (o *OpenapiTaskGenerationResultAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OpenapiTaskGenerationResultAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OpenapiTaskGenerationResultAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OpenapiTaskGenerationResultAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetTaskDisplayName

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskDisplayName() string`

GetTaskDisplayName returns the TaskDisplayName field if non-nil, zero value otherwise.

### GetTaskDisplayNameOk

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskDisplayNameOk() (*string, bool)`

GetTaskDisplayNameOk returns a tuple with the TaskDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDisplayName

`func (o *OpenapiTaskGenerationResultAllOf) SetTaskDisplayName(v string)`

SetTaskDisplayName sets TaskDisplayName field to given value.

### HasTaskDisplayName

`func (o *OpenapiTaskGenerationResultAllOf) HasTaskDisplayName() bool`

HasTaskDisplayName returns a boolean if a field has been set.

### GetTaskName

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *OpenapiTaskGenerationResultAllOf) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *OpenapiTaskGenerationResultAllOf) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskStatus

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *OpenapiTaskGenerationResultAllOf) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *OpenapiTaskGenerationResultAllOf) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### GetTaskVersion

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskVersion() int64`

GetTaskVersion returns the TaskVersion field if non-nil, zero value otherwise.

### GetTaskVersionOk

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskVersionOk() (*int64, bool)`

GetTaskVersionOk returns a tuple with the TaskVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskVersion

`func (o *OpenapiTaskGenerationResultAllOf) SetTaskVersion(v int64)`

SetTaskVersion sets TaskVersion field to given value.

### HasTaskVersion

`func (o *OpenapiTaskGenerationResultAllOf) HasTaskVersion() bool`

HasTaskVersion returns a boolean if a field has been set.

### GetTaskGenerationRequest

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskGenerationRequest() OpenapiTaskGenerationRequestRelationship`

GetTaskGenerationRequest returns the TaskGenerationRequest field if non-nil, zero value otherwise.

### GetTaskGenerationRequestOk

`func (o *OpenapiTaskGenerationResultAllOf) GetTaskGenerationRequestOk() (*OpenapiTaskGenerationRequestRelationship, bool)`

GetTaskGenerationRequestOk returns a tuple with the TaskGenerationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGenerationRequest

`func (o *OpenapiTaskGenerationResultAllOf) SetTaskGenerationRequest(v OpenapiTaskGenerationRequestRelationship)`

SetTaskGenerationRequest sets TaskGenerationRequest field to given value.

### HasTaskGenerationRequest

`func (o *OpenapiTaskGenerationResultAllOf) HasTaskGenerationRequest() bool`

HasTaskGenerationRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


