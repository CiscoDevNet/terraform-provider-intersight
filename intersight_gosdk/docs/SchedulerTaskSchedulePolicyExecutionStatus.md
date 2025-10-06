# SchedulerTaskSchedulePolicyExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.TaskSchedulePolicyExecutionStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.TaskSchedulePolicyExecutionStatus"]
**Name** | Pointer to **string** | Name of the schedule defined in SchedulePolicy. | [optional] [readonly] 
**Status** | Pointer to [**NullableSchedulerTaskScheduleStatus**](SchedulerTaskScheduleStatus.md) |  | [optional] 

## Methods

### NewSchedulerTaskSchedulePolicyExecutionStatus

`func NewSchedulerTaskSchedulePolicyExecutionStatus(classId string, objectType string, ) *SchedulerTaskSchedulePolicyExecutionStatus`

NewSchedulerTaskSchedulePolicyExecutionStatus instantiates a new SchedulerTaskSchedulePolicyExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskSchedulePolicyExecutionStatusWithDefaults

`func NewSchedulerTaskSchedulePolicyExecutionStatusWithDefaults() *SchedulerTaskSchedulePolicyExecutionStatus`

NewSchedulerTaskSchedulePolicyExecutionStatusWithDefaults instantiates a new SchedulerTaskSchedulePolicyExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetStatus() SchedulerTaskScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) GetStatusOk() (*SchedulerTaskScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) SetStatus(v SchedulerTaskScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskSchedulePolicyExecutionStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskSchedulePolicyExecutionStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


