# ConvergedinfraAlarmSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.AlarmSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.AlarmSummary"]
**Critical** | Pointer to **int64** | The count of alarms that have severity type Critical. | [optional] 
**Warning** | Pointer to **int64** | The count of alarms that have severity type Warning. | [optional] 

## Methods

### NewConvergedinfraAlarmSummaryAllOf

`func NewConvergedinfraAlarmSummaryAllOf(classId string, objectType string, ) *ConvergedinfraAlarmSummaryAllOf`

NewConvergedinfraAlarmSummaryAllOf instantiates a new ConvergedinfraAlarmSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraAlarmSummaryAllOfWithDefaults

`func NewConvergedinfraAlarmSummaryAllOfWithDefaults() *ConvergedinfraAlarmSummaryAllOf`

NewConvergedinfraAlarmSummaryAllOfWithDefaults instantiates a new ConvergedinfraAlarmSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraAlarmSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraAlarmSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraAlarmSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraAlarmSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraAlarmSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraAlarmSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *ConvergedinfraAlarmSummaryAllOf) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ConvergedinfraAlarmSummaryAllOf) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ConvergedinfraAlarmSummaryAllOf) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ConvergedinfraAlarmSummaryAllOf) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *ConvergedinfraAlarmSummaryAllOf) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ConvergedinfraAlarmSummaryAllOf) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ConvergedinfraAlarmSummaryAllOf) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ConvergedinfraAlarmSummaryAllOf) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


