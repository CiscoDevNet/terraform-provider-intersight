# IaasCustomTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.CustomTaskInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.CustomTaskInfo"]
**TaskDescription** | Pointer to **string** | Task decription or Comment of the Custom task. | [optional] [readonly] 
**TaskExecutionCount** | Pointer to **int64** | Number of times this task has executed. | [optional] [readonly] 
**TaskLabel** | Pointer to **string** | Task Label in the Workflow. | [optional] [readonly] 
**TaskName** | Pointer to **string** | Name of the Custom Task in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasCustomTaskInfo

`func NewIaasCustomTaskInfo(classId string, objectType string, ) *IaasCustomTaskInfo`

NewIaasCustomTaskInfo instantiates a new IaasCustomTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasCustomTaskInfoWithDefaults

`func NewIaasCustomTaskInfoWithDefaults() *IaasCustomTaskInfo`

NewIaasCustomTaskInfoWithDefaults instantiates a new IaasCustomTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasCustomTaskInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasCustomTaskInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasCustomTaskInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasCustomTaskInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasCustomTaskInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasCustomTaskInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTaskDescription

`func (o *IaasCustomTaskInfo) GetTaskDescription() string`

GetTaskDescription returns the TaskDescription field if non-nil, zero value otherwise.

### GetTaskDescriptionOk

`func (o *IaasCustomTaskInfo) GetTaskDescriptionOk() (*string, bool)`

GetTaskDescriptionOk returns a tuple with the TaskDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDescription

`func (o *IaasCustomTaskInfo) SetTaskDescription(v string)`

SetTaskDescription sets TaskDescription field to given value.

### HasTaskDescription

`func (o *IaasCustomTaskInfo) HasTaskDescription() bool`

HasTaskDescription returns a boolean if a field has been set.

### GetTaskExecutionCount

`func (o *IaasCustomTaskInfo) GetTaskExecutionCount() int64`

GetTaskExecutionCount returns the TaskExecutionCount field if non-nil, zero value otherwise.

### GetTaskExecutionCountOk

`func (o *IaasCustomTaskInfo) GetTaskExecutionCountOk() (*int64, bool)`

GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionCount

`func (o *IaasCustomTaskInfo) SetTaskExecutionCount(v int64)`

SetTaskExecutionCount sets TaskExecutionCount field to given value.

### HasTaskExecutionCount

`func (o *IaasCustomTaskInfo) HasTaskExecutionCount() bool`

HasTaskExecutionCount returns a boolean if a field has been set.

### GetTaskLabel

`func (o *IaasCustomTaskInfo) GetTaskLabel() string`

GetTaskLabel returns the TaskLabel field if non-nil, zero value otherwise.

### GetTaskLabelOk

`func (o *IaasCustomTaskInfo) GetTaskLabelOk() (*string, bool)`

GetTaskLabelOk returns a tuple with the TaskLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskLabel

`func (o *IaasCustomTaskInfo) SetTaskLabel(v string)`

SetTaskLabel sets TaskLabel field to given value.

### HasTaskLabel

`func (o *IaasCustomTaskInfo) HasTaskLabel() bool`

HasTaskLabel returns a boolean if a field has been set.

### GetTaskName

`func (o *IaasCustomTaskInfo) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *IaasCustomTaskInfo) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *IaasCustomTaskInfo) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *IaasCustomTaskInfo) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetGuid

`func (o *IaasCustomTaskInfo) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasCustomTaskInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasCustomTaskInfo) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasCustomTaskInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


