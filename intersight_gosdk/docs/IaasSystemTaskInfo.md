# IaasSystemTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.SystemTaskInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.SystemTaskInfo"]
**TaskCategory** | Pointer to **string** | A functional area to which a task belongs to. | [optional] [readonly] 
**TaskExecutionCount** | Pointer to **int64** | Number of times this task has been executed. | [optional] [readonly] 
**TaskName** | Pointer to **string** | Task name of the UCSD Task. | [optional] [readonly] 
**Guid** | Pointer to [**NullableIaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasSystemTaskInfo

`func NewIaasSystemTaskInfo(classId string, objectType string, ) *IaasSystemTaskInfo`

NewIaasSystemTaskInfo instantiates a new IaasSystemTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasSystemTaskInfoWithDefaults

`func NewIaasSystemTaskInfoWithDefaults() *IaasSystemTaskInfo`

NewIaasSystemTaskInfoWithDefaults instantiates a new IaasSystemTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasSystemTaskInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasSystemTaskInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasSystemTaskInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasSystemTaskInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasSystemTaskInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasSystemTaskInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTaskCategory

`func (o *IaasSystemTaskInfo) GetTaskCategory() string`

GetTaskCategory returns the TaskCategory field if non-nil, zero value otherwise.

### GetTaskCategoryOk

`func (o *IaasSystemTaskInfo) GetTaskCategoryOk() (*string, bool)`

GetTaskCategoryOk returns a tuple with the TaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategory

`func (o *IaasSystemTaskInfo) SetTaskCategory(v string)`

SetTaskCategory sets TaskCategory field to given value.

### HasTaskCategory

`func (o *IaasSystemTaskInfo) HasTaskCategory() bool`

HasTaskCategory returns a boolean if a field has been set.

### GetTaskExecutionCount

`func (o *IaasSystemTaskInfo) GetTaskExecutionCount() int64`

GetTaskExecutionCount returns the TaskExecutionCount field if non-nil, zero value otherwise.

### GetTaskExecutionCountOk

`func (o *IaasSystemTaskInfo) GetTaskExecutionCountOk() (*int64, bool)`

GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionCount

`func (o *IaasSystemTaskInfo) SetTaskExecutionCount(v int64)`

SetTaskExecutionCount sets TaskExecutionCount field to given value.

### HasTaskExecutionCount

`func (o *IaasSystemTaskInfo) HasTaskExecutionCount() bool`

HasTaskExecutionCount returns a boolean if a field has been set.

### GetTaskName

`func (o *IaasSystemTaskInfo) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *IaasSystemTaskInfo) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *IaasSystemTaskInfo) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *IaasSystemTaskInfo) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetGuid

`func (o *IaasSystemTaskInfo) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasSystemTaskInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasSystemTaskInfo) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasSystemTaskInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IaasSystemTaskInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IaasSystemTaskInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


