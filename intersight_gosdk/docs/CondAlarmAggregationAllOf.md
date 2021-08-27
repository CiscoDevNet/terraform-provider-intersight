# CondAlarmAggregationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmAggregation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmAggregation"]
**AlarmSummary** | Pointer to [**NullableCondAlarmSummary**](CondAlarmSummary.md) |  | [optional] 
**CriticalAlarmsCount** | Pointer to **int64** | Count of all alarms with severity Critical, irrespective of acknowledgement status. | [optional] 
**Health** | Pointer to **string** | Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [default to "None"]
**InfoAlarmsCount** | Pointer to **int64** | Count of all alarms with severity Info, irrespective of acknowledgement status. | [optional] 
**MoType** | Pointer to **string** | Managed object type. For example, FI managed object type will be network.Element. | [optional] 
**WarningAlarmsCount** | Pointer to **int64** | Count of all alarms with severity Warning, irrespective of acknowledgement status. | [optional] 
**AlarmAggregationSource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewCondAlarmAggregationAllOf

`func NewCondAlarmAggregationAllOf(classId string, objectType string, ) *CondAlarmAggregationAllOf`

NewCondAlarmAggregationAllOf instantiates a new CondAlarmAggregationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmAggregationAllOfWithDefaults

`func NewCondAlarmAggregationAllOfWithDefaults() *CondAlarmAggregationAllOf`

NewCondAlarmAggregationAllOfWithDefaults instantiates a new CondAlarmAggregationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmAggregationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmAggregationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmAggregationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmAggregationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmAggregationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmAggregationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *CondAlarmAggregationAllOf) GetAlarmSummary() CondAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *CondAlarmAggregationAllOf) GetAlarmSummaryOk() (*CondAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *CondAlarmAggregationAllOf) SetAlarmSummary(v CondAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *CondAlarmAggregationAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *CondAlarmAggregationAllOf) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *CondAlarmAggregationAllOf) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetCriticalAlarmsCount

`func (o *CondAlarmAggregationAllOf) GetCriticalAlarmsCount() int64`

GetCriticalAlarmsCount returns the CriticalAlarmsCount field if non-nil, zero value otherwise.

### GetCriticalAlarmsCountOk

`func (o *CondAlarmAggregationAllOf) GetCriticalAlarmsCountOk() (*int64, bool)`

GetCriticalAlarmsCountOk returns a tuple with the CriticalAlarmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAlarmsCount

`func (o *CondAlarmAggregationAllOf) SetCriticalAlarmsCount(v int64)`

SetCriticalAlarmsCount sets CriticalAlarmsCount field to given value.

### HasCriticalAlarmsCount

`func (o *CondAlarmAggregationAllOf) HasCriticalAlarmsCount() bool`

HasCriticalAlarmsCount returns a boolean if a field has been set.

### GetHealth

`func (o *CondAlarmAggregationAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CondAlarmAggregationAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CondAlarmAggregationAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CondAlarmAggregationAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInfoAlarmsCount

`func (o *CondAlarmAggregationAllOf) GetInfoAlarmsCount() int64`

GetInfoAlarmsCount returns the InfoAlarmsCount field if non-nil, zero value otherwise.

### GetInfoAlarmsCountOk

`func (o *CondAlarmAggregationAllOf) GetInfoAlarmsCountOk() (*int64, bool)`

GetInfoAlarmsCountOk returns a tuple with the InfoAlarmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoAlarmsCount

`func (o *CondAlarmAggregationAllOf) SetInfoAlarmsCount(v int64)`

SetInfoAlarmsCount sets InfoAlarmsCount field to given value.

### HasInfoAlarmsCount

`func (o *CondAlarmAggregationAllOf) HasInfoAlarmsCount() bool`

HasInfoAlarmsCount returns a boolean if a field has been set.

### GetMoType

`func (o *CondAlarmAggregationAllOf) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *CondAlarmAggregationAllOf) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *CondAlarmAggregationAllOf) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *CondAlarmAggregationAllOf) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetWarningAlarmsCount

`func (o *CondAlarmAggregationAllOf) GetWarningAlarmsCount() int64`

GetWarningAlarmsCount returns the WarningAlarmsCount field if non-nil, zero value otherwise.

### GetWarningAlarmsCountOk

`func (o *CondAlarmAggregationAllOf) GetWarningAlarmsCountOk() (*int64, bool)`

GetWarningAlarmsCountOk returns a tuple with the WarningAlarmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningAlarmsCount

`func (o *CondAlarmAggregationAllOf) SetWarningAlarmsCount(v int64)`

SetWarningAlarmsCount sets WarningAlarmsCount field to given value.

### HasWarningAlarmsCount

`func (o *CondAlarmAggregationAllOf) HasWarningAlarmsCount() bool`

HasWarningAlarmsCount returns a boolean if a field has been set.

### GetAlarmAggregationSource

`func (o *CondAlarmAggregationAllOf) GetAlarmAggregationSource() MoBaseMoRelationship`

GetAlarmAggregationSource returns the AlarmAggregationSource field if non-nil, zero value otherwise.

### GetAlarmAggregationSourceOk

`func (o *CondAlarmAggregationAllOf) GetAlarmAggregationSourceOk() (*MoBaseMoRelationship, bool)`

GetAlarmAggregationSourceOk returns a tuple with the AlarmAggregationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmAggregationSource

`func (o *CondAlarmAggregationAllOf) SetAlarmAggregationSource(v MoBaseMoRelationship)`

SetAlarmAggregationSource sets AlarmAggregationSource field to given value.

### HasAlarmAggregationSource

`func (o *CondAlarmAggregationAllOf) HasAlarmAggregationSource() bool`

HasAlarmAggregationSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


