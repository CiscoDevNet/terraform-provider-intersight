# ComputeAlarmSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.AlarmSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.AlarmSummary"]
**Critical** | Pointer to **int64** | The count of alarms that have severity type Critical. | [optional] [readonly] 
**Health** | Pointer to **string** | Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * &#x60;Healthy&#x60; - The Enum value represents that the entity is healthy. * &#x60;Warning&#x60; - The Enum value Warning represents that the entity has one or more active warnings on it. * &#x60;Critical&#x60; - The Enum value Critical represents that the entity is in a critical state. | [optional] [readonly] [default to "Healthy"]
**Info** | Pointer to **int64** | The count of alarms that have severity type Info. | [optional] [readonly] 
**Suppressed** | Pointer to **bool** | The flag that indicates whether suppression is enabled or not in the entity. | [optional] [readonly] 
**SuppressedCritical** | Pointer to **int64** | The count of active suppressed alarms that have severity type Critical. | [optional] [readonly] 
**SuppressedInfo** | Pointer to **int64** | The count of active suppressed alarms that have severity type Info. | [optional] [readonly] 
**SuppressedWarning** | Pointer to **int64** | The count of active suppressed alarms that have severity type Warning. | [optional] [readonly] 
**Warning** | Pointer to **int64** | The count of alarms that have severity type Warning. | [optional] [readonly] 

## Methods

### NewComputeAlarmSummaryAllOf

`func NewComputeAlarmSummaryAllOf(classId string, objectType string, ) *ComputeAlarmSummaryAllOf`

NewComputeAlarmSummaryAllOf instantiates a new ComputeAlarmSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAlarmSummaryAllOfWithDefaults

`func NewComputeAlarmSummaryAllOfWithDefaults() *ComputeAlarmSummaryAllOf`

NewComputeAlarmSummaryAllOfWithDefaults instantiates a new ComputeAlarmSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeAlarmSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeAlarmSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeAlarmSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeAlarmSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeAlarmSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeAlarmSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *ComputeAlarmSummaryAllOf) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ComputeAlarmSummaryAllOf) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ComputeAlarmSummaryAllOf) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ComputeAlarmSummaryAllOf) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHealth

`func (o *ComputeAlarmSummaryAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ComputeAlarmSummaryAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ComputeAlarmSummaryAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ComputeAlarmSummaryAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInfo

`func (o *ComputeAlarmSummaryAllOf) GetInfo() int64`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ComputeAlarmSummaryAllOf) GetInfoOk() (*int64, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ComputeAlarmSummaryAllOf) SetInfo(v int64)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ComputeAlarmSummaryAllOf) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSuppressed

`func (o *ComputeAlarmSummaryAllOf) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *ComputeAlarmSummaryAllOf) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *ComputeAlarmSummaryAllOf) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedCritical

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedCritical() int64`

GetSuppressedCritical returns the SuppressedCritical field if non-nil, zero value otherwise.

### GetSuppressedCriticalOk

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedCriticalOk() (*int64, bool)`

GetSuppressedCriticalOk returns a tuple with the SuppressedCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedCritical

`func (o *ComputeAlarmSummaryAllOf) SetSuppressedCritical(v int64)`

SetSuppressedCritical sets SuppressedCritical field to given value.

### HasSuppressedCritical

`func (o *ComputeAlarmSummaryAllOf) HasSuppressedCritical() bool`

HasSuppressedCritical returns a boolean if a field has been set.

### GetSuppressedInfo

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedInfo() int64`

GetSuppressedInfo returns the SuppressedInfo field if non-nil, zero value otherwise.

### GetSuppressedInfoOk

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedInfoOk() (*int64, bool)`

GetSuppressedInfoOk returns a tuple with the SuppressedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedInfo

`func (o *ComputeAlarmSummaryAllOf) SetSuppressedInfo(v int64)`

SetSuppressedInfo sets SuppressedInfo field to given value.

### HasSuppressedInfo

`func (o *ComputeAlarmSummaryAllOf) HasSuppressedInfo() bool`

HasSuppressedInfo returns a boolean if a field has been set.

### GetSuppressedWarning

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedWarning() int64`

GetSuppressedWarning returns the SuppressedWarning field if non-nil, zero value otherwise.

### GetSuppressedWarningOk

`func (o *ComputeAlarmSummaryAllOf) GetSuppressedWarningOk() (*int64, bool)`

GetSuppressedWarningOk returns a tuple with the SuppressedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedWarning

`func (o *ComputeAlarmSummaryAllOf) SetSuppressedWarning(v int64)`

SetSuppressedWarning sets SuppressedWarning field to given value.

### HasSuppressedWarning

`func (o *ComputeAlarmSummaryAllOf) HasSuppressedWarning() bool`

HasSuppressedWarning returns a boolean if a field has been set.

### GetWarning

`func (o *ComputeAlarmSummaryAllOf) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ComputeAlarmSummaryAllOf) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ComputeAlarmSummaryAllOf) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ComputeAlarmSummaryAllOf) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


