# IaasMostRunTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.MostRunTasks"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.MostRunTasks"]
**TaskCategory** | Pointer to **string** | A functional area to which a task belongs to. | [optional] [readonly] 
**TaskExecutionCount** | Pointer to **int64** | Number of times this task has executed. | [optional] [readonly] 
**TaskName** | Pointer to **string** | Name of the task executed in UCSD. | [optional] [readonly] 
**TaskType** | Pointer to **string** | Type of the task whether it is system task or custom task. | [optional] [readonly] 
**Guid** | Pointer to [**NullableIaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasMostRunTasks

`func NewIaasMostRunTasks(classId string, objectType string, ) *IaasMostRunTasks`

NewIaasMostRunTasks instantiates a new IaasMostRunTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasMostRunTasksWithDefaults

`func NewIaasMostRunTasksWithDefaults() *IaasMostRunTasks`

NewIaasMostRunTasksWithDefaults instantiates a new IaasMostRunTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasMostRunTasks) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasMostRunTasks) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasMostRunTasks) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasMostRunTasks) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasMostRunTasks) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasMostRunTasks) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTaskCategory

`func (o *IaasMostRunTasks) GetTaskCategory() string`

GetTaskCategory returns the TaskCategory field if non-nil, zero value otherwise.

### GetTaskCategoryOk

`func (o *IaasMostRunTasks) GetTaskCategoryOk() (*string, bool)`

GetTaskCategoryOk returns a tuple with the TaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategory

`func (o *IaasMostRunTasks) SetTaskCategory(v string)`

SetTaskCategory sets TaskCategory field to given value.

### HasTaskCategory

`func (o *IaasMostRunTasks) HasTaskCategory() bool`

HasTaskCategory returns a boolean if a field has been set.

### GetTaskExecutionCount

`func (o *IaasMostRunTasks) GetTaskExecutionCount() int64`

GetTaskExecutionCount returns the TaskExecutionCount field if non-nil, zero value otherwise.

### GetTaskExecutionCountOk

`func (o *IaasMostRunTasks) GetTaskExecutionCountOk() (*int64, bool)`

GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionCount

`func (o *IaasMostRunTasks) SetTaskExecutionCount(v int64)`

SetTaskExecutionCount sets TaskExecutionCount field to given value.

### HasTaskExecutionCount

`func (o *IaasMostRunTasks) HasTaskExecutionCount() bool`

HasTaskExecutionCount returns a boolean if a field has been set.

### GetTaskName

`func (o *IaasMostRunTasks) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *IaasMostRunTasks) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *IaasMostRunTasks) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *IaasMostRunTasks) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskType

`func (o *IaasMostRunTasks) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *IaasMostRunTasks) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *IaasMostRunTasks) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *IaasMostRunTasks) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetGuid

`func (o *IaasMostRunTasks) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasMostRunTasks) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasMostRunTasks) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasMostRunTasks) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IaasMostRunTasks) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IaasMostRunTasks) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


