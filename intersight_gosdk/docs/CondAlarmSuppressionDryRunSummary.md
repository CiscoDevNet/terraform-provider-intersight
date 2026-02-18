# CondAlarmSuppressionDryRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmSuppressionDryRunSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmSuppressionDryRunSummary"]
**Count** | Pointer to **int64** | The count of alarms that would be impacted by the proposed suppression rule. | [optional] [readonly] 
**TopAlarms** | Pointer to [**[]CondAlarmPreview**](CondAlarmPreview.md) |  | [optional] 

## Methods

### NewCondAlarmSuppressionDryRunSummary

`func NewCondAlarmSuppressionDryRunSummary(classId string, objectType string, ) *CondAlarmSuppressionDryRunSummary`

NewCondAlarmSuppressionDryRunSummary instantiates a new CondAlarmSuppressionDryRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmSuppressionDryRunSummaryWithDefaults

`func NewCondAlarmSuppressionDryRunSummaryWithDefaults() *CondAlarmSuppressionDryRunSummary`

NewCondAlarmSuppressionDryRunSummaryWithDefaults instantiates a new CondAlarmSuppressionDryRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmSuppressionDryRunSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmSuppressionDryRunSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmSuppressionDryRunSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmSuppressionDryRunSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmSuppressionDryRunSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmSuppressionDryRunSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *CondAlarmSuppressionDryRunSummary) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CondAlarmSuppressionDryRunSummary) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CondAlarmSuppressionDryRunSummary) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CondAlarmSuppressionDryRunSummary) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTopAlarms

`func (o *CondAlarmSuppressionDryRunSummary) GetTopAlarms() []CondAlarmPreview`

GetTopAlarms returns the TopAlarms field if non-nil, zero value otherwise.

### GetTopAlarmsOk

`func (o *CondAlarmSuppressionDryRunSummary) GetTopAlarmsOk() (*[]CondAlarmPreview, bool)`

GetTopAlarmsOk returns a tuple with the TopAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAlarms

`func (o *CondAlarmSuppressionDryRunSummary) SetTopAlarms(v []CondAlarmPreview)`

SetTopAlarms sets TopAlarms field to given value.

### HasTopAlarms

`func (o *CondAlarmSuppressionDryRunSummary) HasTopAlarms() bool`

HasTopAlarms returns a boolean if a field has been set.

### SetTopAlarmsNil

`func (o *CondAlarmSuppressionDryRunSummary) SetTopAlarmsNil(b bool)`

 SetTopAlarmsNil sets the value for TopAlarms to be an explicit nil

### UnsetTopAlarms
`func (o *CondAlarmSuppressionDryRunSummary) UnsetTopAlarms()`

UnsetTopAlarms ensures that no value is present for TopAlarms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


