# WorkflowTaskDebugLogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskDebugLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskDebugLog"]
**RetryCount** | Pointer to **int64** | A counter for number of retries. | [optional] [readonly] 
**TaskDebugLogEntries** | Pointer to **interface{}** | Holds information helpful in isolating task failures. | [optional] [readonly] 
**TaskInstId** | Pointer to **string** | The unique identifier for task instance. | [optional] [readonly] 
**TaskInfo** | Pointer to [**WorkflowTaskInfoRelationship**](WorkflowTaskInfoRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowTaskDebugLogAllOf

`func NewWorkflowTaskDebugLogAllOf(classId string, objectType string, ) *WorkflowTaskDebugLogAllOf`

NewWorkflowTaskDebugLogAllOf instantiates a new WorkflowTaskDebugLogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskDebugLogAllOfWithDefaults

`func NewWorkflowTaskDebugLogAllOfWithDefaults() *WorkflowTaskDebugLogAllOf`

NewWorkflowTaskDebugLogAllOfWithDefaults instantiates a new WorkflowTaskDebugLogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskDebugLogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskDebugLogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskDebugLogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskDebugLogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskDebugLogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskDebugLogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRetryCount

`func (o *WorkflowTaskDebugLogAllOf) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTaskDebugLogAllOf) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTaskDebugLogAllOf) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTaskDebugLogAllOf) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetTaskDebugLogEntries

`func (o *WorkflowTaskDebugLogAllOf) GetTaskDebugLogEntries() interface{}`

GetTaskDebugLogEntries returns the TaskDebugLogEntries field if non-nil, zero value otherwise.

### GetTaskDebugLogEntriesOk

`func (o *WorkflowTaskDebugLogAllOf) GetTaskDebugLogEntriesOk() (*interface{}, bool)`

GetTaskDebugLogEntriesOk returns a tuple with the TaskDebugLogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDebugLogEntries

`func (o *WorkflowTaskDebugLogAllOf) SetTaskDebugLogEntries(v interface{})`

SetTaskDebugLogEntries sets TaskDebugLogEntries field to given value.

### HasTaskDebugLogEntries

`func (o *WorkflowTaskDebugLogAllOf) HasTaskDebugLogEntries() bool`

HasTaskDebugLogEntries returns a boolean if a field has been set.

### SetTaskDebugLogEntriesNil

`func (o *WorkflowTaskDebugLogAllOf) SetTaskDebugLogEntriesNil(b bool)`

 SetTaskDebugLogEntriesNil sets the value for TaskDebugLogEntries to be an explicit nil

### UnsetTaskDebugLogEntries
`func (o *WorkflowTaskDebugLogAllOf) UnsetTaskDebugLogEntries()`

UnsetTaskDebugLogEntries ensures that no value is present for TaskDebugLogEntries, not even an explicit nil
### GetTaskInstId

`func (o *WorkflowTaskDebugLogAllOf) GetTaskInstId() string`

GetTaskInstId returns the TaskInstId field if non-nil, zero value otherwise.

### GetTaskInstIdOk

`func (o *WorkflowTaskDebugLogAllOf) GetTaskInstIdOk() (*string, bool)`

GetTaskInstIdOk returns a tuple with the TaskInstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstId

`func (o *WorkflowTaskDebugLogAllOf) SetTaskInstId(v string)`

SetTaskInstId sets TaskInstId field to given value.

### HasTaskInstId

`func (o *WorkflowTaskDebugLogAllOf) HasTaskInstId() bool`

HasTaskInstId returns a boolean if a field has been set.

### GetTaskInfo

`func (o *WorkflowTaskDebugLogAllOf) GetTaskInfo() WorkflowTaskInfoRelationship`

GetTaskInfo returns the TaskInfo field if non-nil, zero value otherwise.

### GetTaskInfoOk

`func (o *WorkflowTaskDebugLogAllOf) GetTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetTaskInfoOk returns a tuple with the TaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfo

`func (o *WorkflowTaskDebugLogAllOf) SetTaskInfo(v WorkflowTaskInfoRelationship)`

SetTaskInfo sets TaskInfo field to given value.

### HasTaskInfo

`func (o *WorkflowTaskDebugLogAllOf) HasTaskInfo() bool`

HasTaskInfo returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *WorkflowTaskDebugLogAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowTaskDebugLogAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowTaskDebugLogAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowTaskDebugLogAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


