# ComputeAlarmSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **int64** | The count of alarms that have severity type Critical. | [optional] 
**Warning** | Pointer to **int64** | The count of alarms that have severity type Warning. | [optional] 

## Methods

### NewComputeAlarmSummaryAllOf

`func NewComputeAlarmSummaryAllOf() *ComputeAlarmSummaryAllOf`

NewComputeAlarmSummaryAllOf instantiates a new ComputeAlarmSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAlarmSummaryAllOfWithDefaults

`func NewComputeAlarmSummaryAllOfWithDefaults() *ComputeAlarmSummaryAllOf`

NewComputeAlarmSummaryAllOfWithDefaults instantiates a new ComputeAlarmSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


