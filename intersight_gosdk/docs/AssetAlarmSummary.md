# AssetAlarmSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.AlarmSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.AlarmSummary"]
**Critical** | Pointer to **int64** | The count of active alarms that have severity type Critical. | [optional] [readonly] 
**Health** | Pointer to **string** | Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * &#x60;Healthy&#x60; - The Enum value represents that the entity is healthy. * &#x60;Warning&#x60; - The Enum value Warning represents that the entity has one or more active warnings on it. * &#x60;Critical&#x60; - The Enum value Critical represents that the entity is in a critical state. | [optional] [readonly] [default to "Healthy"]
**Info** | Pointer to **int64** | The count of active alarms that have severity type Info. | [optional] [readonly] 
**SuppressedCritical** | Pointer to **int64** | The count of active suppressed alarms that have severity type Critical. | [optional] [readonly] 
**SuppressedInfo** | Pointer to **int64** | The count of active suppressed alarms that have severity type Info. | [optional] [readonly] 
**SuppressedWarning** | Pointer to **int64** | The count of active suppressed alarms that have severity type Warning. | [optional] [readonly] 
**Warning** | Pointer to **int64** | The count of active alarms that have severity type Warning. | [optional] [readonly] 

## Methods

### NewAssetAlarmSummary

`func NewAssetAlarmSummary(classId string, objectType string, ) *AssetAlarmSummary`

NewAssetAlarmSummary instantiates a new AssetAlarmSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetAlarmSummaryWithDefaults

`func NewAssetAlarmSummaryWithDefaults() *AssetAlarmSummary`

NewAssetAlarmSummaryWithDefaults instantiates a new AssetAlarmSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetAlarmSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetAlarmSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetAlarmSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetAlarmSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetAlarmSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetAlarmSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *AssetAlarmSummary) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *AssetAlarmSummary) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *AssetAlarmSummary) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *AssetAlarmSummary) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHealth

`func (o *AssetAlarmSummary) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *AssetAlarmSummary) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *AssetAlarmSummary) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *AssetAlarmSummary) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInfo

`func (o *AssetAlarmSummary) GetInfo() int64`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AssetAlarmSummary) GetInfoOk() (*int64, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AssetAlarmSummary) SetInfo(v int64)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AssetAlarmSummary) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSuppressedCritical

`func (o *AssetAlarmSummary) GetSuppressedCritical() int64`

GetSuppressedCritical returns the SuppressedCritical field if non-nil, zero value otherwise.

### GetSuppressedCriticalOk

`func (o *AssetAlarmSummary) GetSuppressedCriticalOk() (*int64, bool)`

GetSuppressedCriticalOk returns a tuple with the SuppressedCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedCritical

`func (o *AssetAlarmSummary) SetSuppressedCritical(v int64)`

SetSuppressedCritical sets SuppressedCritical field to given value.

### HasSuppressedCritical

`func (o *AssetAlarmSummary) HasSuppressedCritical() bool`

HasSuppressedCritical returns a boolean if a field has been set.

### GetSuppressedInfo

`func (o *AssetAlarmSummary) GetSuppressedInfo() int64`

GetSuppressedInfo returns the SuppressedInfo field if non-nil, zero value otherwise.

### GetSuppressedInfoOk

`func (o *AssetAlarmSummary) GetSuppressedInfoOk() (*int64, bool)`

GetSuppressedInfoOk returns a tuple with the SuppressedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedInfo

`func (o *AssetAlarmSummary) SetSuppressedInfo(v int64)`

SetSuppressedInfo sets SuppressedInfo field to given value.

### HasSuppressedInfo

`func (o *AssetAlarmSummary) HasSuppressedInfo() bool`

HasSuppressedInfo returns a boolean if a field has been set.

### GetSuppressedWarning

`func (o *AssetAlarmSummary) GetSuppressedWarning() int64`

GetSuppressedWarning returns the SuppressedWarning field if non-nil, zero value otherwise.

### GetSuppressedWarningOk

`func (o *AssetAlarmSummary) GetSuppressedWarningOk() (*int64, bool)`

GetSuppressedWarningOk returns a tuple with the SuppressedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedWarning

`func (o *AssetAlarmSummary) SetSuppressedWarning(v int64)`

SetSuppressedWarning sets SuppressedWarning field to given value.

### HasSuppressedWarning

`func (o *AssetAlarmSummary) HasSuppressedWarning() bool`

HasSuppressedWarning returns a boolean if a field has been set.

### GetWarning

`func (o *AssetAlarmSummary) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AssetAlarmSummary) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AssetAlarmSummary) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AssetAlarmSummary) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


