# WorkflowTaskRetryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskRetryInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskRetryInfo"]
**Status** | Pointer to **string** | Status of the retried task. | [optional] 
**TaskInstId** | Pointer to **string** | Retry instance will get a unique instance id. | [optional] 

## Methods

### NewWorkflowTaskRetryInfo

`func NewWorkflowTaskRetryInfo(classId string, objectType string, ) *WorkflowTaskRetryInfo`

NewWorkflowTaskRetryInfo instantiates a new WorkflowTaskRetryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskRetryInfoWithDefaults

`func NewWorkflowTaskRetryInfoWithDefaults() *WorkflowTaskRetryInfo`

NewWorkflowTaskRetryInfoWithDefaults instantiates a new WorkflowTaskRetryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskRetryInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskRetryInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskRetryInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskRetryInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskRetryInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskRetryInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *WorkflowTaskRetryInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTaskRetryInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTaskRetryInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTaskRetryInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskInstId

`func (o *WorkflowTaskRetryInfo) GetTaskInstId() string`

GetTaskInstId returns the TaskInstId field if non-nil, zero value otherwise.

### GetTaskInstIdOk

`func (o *WorkflowTaskRetryInfo) GetTaskInstIdOk() (*string, bool)`

GetTaskInstIdOk returns a tuple with the TaskInstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstId

`func (o *WorkflowTaskRetryInfo) SetTaskInstId(v string)`

SetTaskInstId sets TaskInstId field to given value.

### HasTaskInstId

`func (o *WorkflowTaskRetryInfo) HasTaskInstId() bool`

HasTaskInstId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


