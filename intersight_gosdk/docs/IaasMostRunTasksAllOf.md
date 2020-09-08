# IaasMostRunTasksAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskCategory** | Pointer to **string** | A functional area to which a task belongs to. | [optional] [readonly] 
**TaskExecutionCount** | Pointer to **int64** | Number of times this task has executed. | [optional] [readonly] 
**TaskName** | Pointer to **string** | Name of the task executed in UCSD. | [optional] [readonly] 
**TaskType** | Pointer to **string** | Type of the task whether it is system task or custom task. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasMostRunTasksAllOf

`func NewIaasMostRunTasksAllOf() *IaasMostRunTasksAllOf`

NewIaasMostRunTasksAllOf instantiates a new IaasMostRunTasksAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasMostRunTasksAllOfWithDefaults

`func NewIaasMostRunTasksAllOfWithDefaults() *IaasMostRunTasksAllOf`

NewIaasMostRunTasksAllOfWithDefaults instantiates a new IaasMostRunTasksAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskCategory

`func (o *IaasMostRunTasksAllOf) GetTaskCategory() string`

GetTaskCategory returns the TaskCategory field if non-nil, zero value otherwise.

### GetTaskCategoryOk

`func (o *IaasMostRunTasksAllOf) GetTaskCategoryOk() (*string, bool)`

GetTaskCategoryOk returns a tuple with the TaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategory

`func (o *IaasMostRunTasksAllOf) SetTaskCategory(v string)`

SetTaskCategory sets TaskCategory field to given value.

### HasTaskCategory

`func (o *IaasMostRunTasksAllOf) HasTaskCategory() bool`

HasTaskCategory returns a boolean if a field has been set.

### GetTaskExecutionCount

`func (o *IaasMostRunTasksAllOf) GetTaskExecutionCount() int64`

GetTaskExecutionCount returns the TaskExecutionCount field if non-nil, zero value otherwise.

### GetTaskExecutionCountOk

`func (o *IaasMostRunTasksAllOf) GetTaskExecutionCountOk() (*int64, bool)`

GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionCount

`func (o *IaasMostRunTasksAllOf) SetTaskExecutionCount(v int64)`

SetTaskExecutionCount sets TaskExecutionCount field to given value.

### HasTaskExecutionCount

`func (o *IaasMostRunTasksAllOf) HasTaskExecutionCount() bool`

HasTaskExecutionCount returns a boolean if a field has been set.

### GetTaskName

`func (o *IaasMostRunTasksAllOf) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *IaasMostRunTasksAllOf) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *IaasMostRunTasksAllOf) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *IaasMostRunTasksAllOf) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskType

`func (o *IaasMostRunTasksAllOf) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *IaasMostRunTasksAllOf) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *IaasMostRunTasksAllOf) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *IaasMostRunTasksAllOf) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetGuid

`func (o *IaasMostRunTasksAllOf) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasMostRunTasksAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasMostRunTasksAllOf) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasMostRunTasksAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


