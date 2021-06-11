# ForecastInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "forecast.Instance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "forecast.Instance"]
**Action** | Pointer to **string** | Action to be triggered on forecast instance. Default value is None. * &#x60;None&#x60; - The Enum value None represents that no action is triggered on the forecast Instance managed object. * &#x60;Evaluate&#x60; - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered. | [optional] [default to "None"]
**AltModel** | Pointer to **[]float32** |  | [optional] 
**DataInterval** | Pointer to **int64** | The time interval (in days) for the data to be used for computing forecast model. | [optional] [default to 180]
**DataStartDate** | Pointer to **time.Time** | The start date from when the data should be used for computing forecast model. | [optional] 
**DeviceId** | Pointer to **string** | The Moid of the Intersight managed device instance for which regression model is derived. | [optional] [readonly] 
**FullCapDays** | Pointer to **int64** | The number of days remaining before the device reaches its full functional capacity. | [optional] [readonly] 
**LastModelUpdateTime** | Pointer to **time.Time** | The time when the forecast model was last updated. | [optional] [readonly] 
**MetricName** | Pointer to **string** | The name of the metric for which regression model is generated. | [optional] [readonly] 
**MinDaysForForecast** | Pointer to **int64** | The minimum number of days the HyperFlex cluster should be up for computing forecast. | [optional] [readonly] 
**Model** | Pointer to [**NullableForecastModel**](forecast.Model.md) |  | [optional] 
**ThresholdDays** | Pointer to **int64** | The number of days remaining before the device reaches the specified threshold for the metric as defined in definition. | [optional] [readonly] [default to 2147483647]
**ForecastDef** | Pointer to [**ForecastDefinitionRelationship**](forecast.Definition.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewForecastInstanceAllOf

`func NewForecastInstanceAllOf(classId string, objectType string, ) *ForecastInstanceAllOf`

NewForecastInstanceAllOf instantiates a new ForecastInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastInstanceAllOfWithDefaults

`func NewForecastInstanceAllOfWithDefaults() *ForecastInstanceAllOf`

NewForecastInstanceAllOfWithDefaults instantiates a new ForecastInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ForecastInstanceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ForecastInstanceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ForecastInstanceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ForecastInstanceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ForecastInstanceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ForecastInstanceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *ForecastInstanceAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ForecastInstanceAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ForecastInstanceAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ForecastInstanceAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAltModel

`func (o *ForecastInstanceAllOf) GetAltModel() []float32`

GetAltModel returns the AltModel field if non-nil, zero value otherwise.

### GetAltModelOk

`func (o *ForecastInstanceAllOf) GetAltModelOk() (*[]float32, bool)`

GetAltModelOk returns a tuple with the AltModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltModel

`func (o *ForecastInstanceAllOf) SetAltModel(v []float32)`

SetAltModel sets AltModel field to given value.

### HasAltModel

`func (o *ForecastInstanceAllOf) HasAltModel() bool`

HasAltModel returns a boolean if a field has been set.

### SetAltModelNil

`func (o *ForecastInstanceAllOf) SetAltModelNil(b bool)`

 SetAltModelNil sets the value for AltModel to be an explicit nil

### UnsetAltModel
`func (o *ForecastInstanceAllOf) UnsetAltModel()`

UnsetAltModel ensures that no value is present for AltModel, not even an explicit nil
### GetDataInterval

`func (o *ForecastInstanceAllOf) GetDataInterval() int64`

GetDataInterval returns the DataInterval field if non-nil, zero value otherwise.

### GetDataIntervalOk

`func (o *ForecastInstanceAllOf) GetDataIntervalOk() (*int64, bool)`

GetDataIntervalOk returns a tuple with the DataInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInterval

`func (o *ForecastInstanceAllOf) SetDataInterval(v int64)`

SetDataInterval sets DataInterval field to given value.

### HasDataInterval

`func (o *ForecastInstanceAllOf) HasDataInterval() bool`

HasDataInterval returns a boolean if a field has been set.

### GetDataStartDate

`func (o *ForecastInstanceAllOf) GetDataStartDate() time.Time`

GetDataStartDate returns the DataStartDate field if non-nil, zero value otherwise.

### GetDataStartDateOk

`func (o *ForecastInstanceAllOf) GetDataStartDateOk() (*time.Time, bool)`

GetDataStartDateOk returns a tuple with the DataStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartDate

`func (o *ForecastInstanceAllOf) SetDataStartDate(v time.Time)`

SetDataStartDate sets DataStartDate field to given value.

### HasDataStartDate

`func (o *ForecastInstanceAllOf) HasDataStartDate() bool`

HasDataStartDate returns a boolean if a field has been set.

### GetDeviceId

`func (o *ForecastInstanceAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ForecastInstanceAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ForecastInstanceAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ForecastInstanceAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFullCapDays

`func (o *ForecastInstanceAllOf) GetFullCapDays() int64`

GetFullCapDays returns the FullCapDays field if non-nil, zero value otherwise.

### GetFullCapDaysOk

`func (o *ForecastInstanceAllOf) GetFullCapDaysOk() (*int64, bool)`

GetFullCapDaysOk returns a tuple with the FullCapDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCapDays

`func (o *ForecastInstanceAllOf) SetFullCapDays(v int64)`

SetFullCapDays sets FullCapDays field to given value.

### HasFullCapDays

`func (o *ForecastInstanceAllOf) HasFullCapDays() bool`

HasFullCapDays returns a boolean if a field has been set.

### GetLastModelUpdateTime

`func (o *ForecastInstanceAllOf) GetLastModelUpdateTime() time.Time`

GetLastModelUpdateTime returns the LastModelUpdateTime field if non-nil, zero value otherwise.

### GetLastModelUpdateTimeOk

`func (o *ForecastInstanceAllOf) GetLastModelUpdateTimeOk() (*time.Time, bool)`

GetLastModelUpdateTimeOk returns a tuple with the LastModelUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModelUpdateTime

`func (o *ForecastInstanceAllOf) SetLastModelUpdateTime(v time.Time)`

SetLastModelUpdateTime sets LastModelUpdateTime field to given value.

### HasLastModelUpdateTime

`func (o *ForecastInstanceAllOf) HasLastModelUpdateTime() bool`

HasLastModelUpdateTime returns a boolean if a field has been set.

### GetMetricName

`func (o *ForecastInstanceAllOf) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ForecastInstanceAllOf) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ForecastInstanceAllOf) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ForecastInstanceAllOf) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMinDaysForForecast

`func (o *ForecastInstanceAllOf) GetMinDaysForForecast() int64`

GetMinDaysForForecast returns the MinDaysForForecast field if non-nil, zero value otherwise.

### GetMinDaysForForecastOk

`func (o *ForecastInstanceAllOf) GetMinDaysForForecastOk() (*int64, bool)`

GetMinDaysForForecastOk returns a tuple with the MinDaysForForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysForForecast

`func (o *ForecastInstanceAllOf) SetMinDaysForForecast(v int64)`

SetMinDaysForForecast sets MinDaysForForecast field to given value.

### HasMinDaysForForecast

`func (o *ForecastInstanceAllOf) HasMinDaysForForecast() bool`

HasMinDaysForForecast returns a boolean if a field has been set.

### GetModel

`func (o *ForecastInstanceAllOf) GetModel() ForecastModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ForecastInstanceAllOf) GetModelOk() (*ForecastModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ForecastInstanceAllOf) SetModel(v ForecastModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *ForecastInstanceAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ForecastInstanceAllOf) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ForecastInstanceAllOf) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetThresholdDays

`func (o *ForecastInstanceAllOf) GetThresholdDays() int64`

GetThresholdDays returns the ThresholdDays field if non-nil, zero value otherwise.

### GetThresholdDaysOk

`func (o *ForecastInstanceAllOf) GetThresholdDaysOk() (*int64, bool)`

GetThresholdDaysOk returns a tuple with the ThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDays

`func (o *ForecastInstanceAllOf) SetThresholdDays(v int64)`

SetThresholdDays sets ThresholdDays field to given value.

### HasThresholdDays

`func (o *ForecastInstanceAllOf) HasThresholdDays() bool`

HasThresholdDays returns a boolean if a field has been set.

### GetForecastDef

`func (o *ForecastInstanceAllOf) GetForecastDef() ForecastDefinitionRelationship`

GetForecastDef returns the ForecastDef field if non-nil, zero value otherwise.

### GetForecastDefOk

`func (o *ForecastInstanceAllOf) GetForecastDefOk() (*ForecastDefinitionRelationship, bool)`

GetForecastDefOk returns a tuple with the ForecastDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastDef

`func (o *ForecastInstanceAllOf) SetForecastDef(v ForecastDefinitionRelationship)`

SetForecastDef sets ForecastDef field to given value.

### HasForecastDef

`func (o *ForecastInstanceAllOf) HasForecastDef() bool`

HasForecastDef returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ForecastInstanceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ForecastInstanceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ForecastInstanceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ForecastInstanceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


