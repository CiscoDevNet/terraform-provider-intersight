# WorkflowValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ValidationError"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ValidationError"]
**ErrorLog** | Pointer to **string** | Description of the error. | [optional] [readonly] 
**Field** | Pointer to **string** | When populated this refers to the input or output field within the workflow or task. | [optional] [readonly] 
**TaskName** | Pointer to **string** | The task name on which the error is found, when empty the error applies to the top level workflow. | [optional] [readonly] 
**TransitionName** | Pointer to **string** | When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue. | [optional] [readonly] 

## Methods

### NewWorkflowValidationError

`func NewWorkflowValidationError(classId string, objectType string, ) *WorkflowValidationError`

NewWorkflowValidationError instantiates a new WorkflowValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowValidationErrorWithDefaults

`func NewWorkflowValidationErrorWithDefaults() *WorkflowValidationError`

NewWorkflowValidationErrorWithDefaults instantiates a new WorkflowValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowValidationError) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowValidationError) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowValidationError) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowValidationError) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowValidationError) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowValidationError) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorLog

`func (o *WorkflowValidationError) GetErrorLog() string`

GetErrorLog returns the ErrorLog field if non-nil, zero value otherwise.

### GetErrorLogOk

`func (o *WorkflowValidationError) GetErrorLogOk() (*string, bool)`

GetErrorLogOk returns a tuple with the ErrorLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLog

`func (o *WorkflowValidationError) SetErrorLog(v string)`

SetErrorLog sets ErrorLog field to given value.

### HasErrorLog

`func (o *WorkflowValidationError) HasErrorLog() bool`

HasErrorLog returns a boolean if a field has been set.

### GetField

`func (o *WorkflowValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *WorkflowValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *WorkflowValidationError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *WorkflowValidationError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetTaskName

`func (o *WorkflowValidationError) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *WorkflowValidationError) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *WorkflowValidationError) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *WorkflowValidationError) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTransitionName

`func (o *WorkflowValidationError) GetTransitionName() string`

GetTransitionName returns the TransitionName field if non-nil, zero value otherwise.

### GetTransitionNameOk

`func (o *WorkflowValidationError) GetTransitionNameOk() (*string, bool)`

GetTransitionNameOk returns a tuple with the TransitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionName

`func (o *WorkflowValidationError) SetTransitionName(v string)`

SetTransitionName sets TransitionName field to given value.

### HasTransitionName

`func (o *WorkflowValidationError) HasTransitionName() bool`

HasTransitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


