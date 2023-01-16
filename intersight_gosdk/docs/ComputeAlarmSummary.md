# ComputeAlarmSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.AlarmSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.AlarmSummary"]
**Critical** | Pointer to **int64** | The count of alarms that have severity type Critical. | [optional] [readonly] 
**Health** | Pointer to **string** | Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * &#x60;Healthy&#x60; - The Enum value represents that the entity is healthy. * &#x60;Warning&#x60; - The Enum value Warning represents that the entity has one or more active warnings on it. * &#x60;Critical&#x60; - The Enum value Critical represents that the entity is in a critical state. | [optional] [readonly] [default to "Healthy"]
**Info** | Pointer to **int64** | The count of alarms that have severity type Info. | [optional] [readonly] 
**Warning** | Pointer to **int64** | The count of alarms that have severity type Warning. | [optional] [readonly] 

## Methods

### NewComputeAlarmSummary

`func NewComputeAlarmSummary(classId string, objectType string, ) *ComputeAlarmSummary`

NewComputeAlarmSummary instantiates a new ComputeAlarmSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAlarmSummaryWithDefaults

`func NewComputeAlarmSummaryWithDefaults() *ComputeAlarmSummary`

NewComputeAlarmSummaryWithDefaults instantiates a new ComputeAlarmSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeAlarmSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeAlarmSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeAlarmSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeAlarmSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeAlarmSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeAlarmSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *ComputeAlarmSummary) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ComputeAlarmSummary) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ComputeAlarmSummary) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ComputeAlarmSummary) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHealth

`func (o *ComputeAlarmSummary) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ComputeAlarmSummary) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ComputeAlarmSummary) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ComputeAlarmSummary) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInfo

`func (o *ComputeAlarmSummary) GetInfo() int64`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ComputeAlarmSummary) GetInfoOk() (*int64, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ComputeAlarmSummary) SetInfo(v int64)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ComputeAlarmSummary) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWarning

`func (o *ComputeAlarmSummary) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ComputeAlarmSummary) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ComputeAlarmSummary) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ComputeAlarmSummary) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


