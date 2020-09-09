# HyperflexAlarmSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **int64** | The count of alarms that have severity type Critical. | [optional] 
**Warning** | Pointer to **int64** | The count of alarms that have severity type Warning. | [optional] 

## Methods

### NewHyperflexAlarmSummary

`func NewHyperflexAlarmSummary() *HyperflexAlarmSummary`

NewHyperflexAlarmSummary instantiates a new HyperflexAlarmSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAlarmSummaryWithDefaults

`func NewHyperflexAlarmSummaryWithDefaults() *HyperflexAlarmSummary`

NewHyperflexAlarmSummaryWithDefaults instantiates a new HyperflexAlarmSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *HyperflexAlarmSummary) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *HyperflexAlarmSummary) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *HyperflexAlarmSummary) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *HyperflexAlarmSummary) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *HyperflexAlarmSummary) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *HyperflexAlarmSummary) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *HyperflexAlarmSummary) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *HyperflexAlarmSummary) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


