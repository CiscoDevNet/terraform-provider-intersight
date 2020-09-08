# ForecastInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltModel** | Pointer to **[]float32** |  | [optional] 
**DeviceId** | Pointer to **string** | The Moid of the Intersight managed device instance for which regression model is derived. | [optional] [readonly] 
**FullCapDays** | Pointer to **int64** | The number of days remaining before the device reaches its full functional capacity. | [optional] [readonly] 
**MetricName** | Pointer to **string** | The name of the metric for which regression model is generated. | [optional] [readonly] 
**MinDaysForForecast** | Pointer to **int64** | The minimum number of days the HyperFlex cluster should be up for computing forecast. | [optional] [readonly] 
**Model** | Pointer to [**ForecastModel**](forecast.Model.md) |  | [optional] 
**ThresholdDays** | Pointer to **int64** | The number of days remaining before the device reaches the specified threshold for the metric as defined in definition. | [optional] [readonly] 
**ForecastDef** | Pointer to [**ForecastDefinitionRelationship**](forecast.Definition.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewForecastInstance

`func NewForecastInstance() *ForecastInstance`

NewForecastInstance instantiates a new ForecastInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastInstanceWithDefaults

`func NewForecastInstanceWithDefaults() *ForecastInstance`

NewForecastInstanceWithDefaults instantiates a new ForecastInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltModel

`func (o *ForecastInstance) GetAltModel() []float32`

GetAltModel returns the AltModel field if non-nil, zero value otherwise.

### GetAltModelOk

`func (o *ForecastInstance) GetAltModelOk() (*[]float32, bool)`

GetAltModelOk returns a tuple with the AltModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltModel

`func (o *ForecastInstance) SetAltModel(v []float32)`

SetAltModel sets AltModel field to given value.

### HasAltModel

`func (o *ForecastInstance) HasAltModel() bool`

HasAltModel returns a boolean if a field has been set.

### GetDeviceId

`func (o *ForecastInstance) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ForecastInstance) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ForecastInstance) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ForecastInstance) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFullCapDays

`func (o *ForecastInstance) GetFullCapDays() int64`

GetFullCapDays returns the FullCapDays field if non-nil, zero value otherwise.

### GetFullCapDaysOk

`func (o *ForecastInstance) GetFullCapDaysOk() (*int64, bool)`

GetFullCapDaysOk returns a tuple with the FullCapDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCapDays

`func (o *ForecastInstance) SetFullCapDays(v int64)`

SetFullCapDays sets FullCapDays field to given value.

### HasFullCapDays

`func (o *ForecastInstance) HasFullCapDays() bool`

HasFullCapDays returns a boolean if a field has been set.

### GetMetricName

`func (o *ForecastInstance) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ForecastInstance) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ForecastInstance) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ForecastInstance) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMinDaysForForecast

`func (o *ForecastInstance) GetMinDaysForForecast() int64`

GetMinDaysForForecast returns the MinDaysForForecast field if non-nil, zero value otherwise.

### GetMinDaysForForecastOk

`func (o *ForecastInstance) GetMinDaysForForecastOk() (*int64, bool)`

GetMinDaysForForecastOk returns a tuple with the MinDaysForForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysForForecast

`func (o *ForecastInstance) SetMinDaysForForecast(v int64)`

SetMinDaysForForecast sets MinDaysForForecast field to given value.

### HasMinDaysForForecast

`func (o *ForecastInstance) HasMinDaysForForecast() bool`

HasMinDaysForForecast returns a boolean if a field has been set.

### GetModel

`func (o *ForecastInstance) GetModel() ForecastModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ForecastInstance) GetModelOk() (*ForecastModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ForecastInstance) SetModel(v ForecastModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *ForecastInstance) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetThresholdDays

`func (o *ForecastInstance) GetThresholdDays() int64`

GetThresholdDays returns the ThresholdDays field if non-nil, zero value otherwise.

### GetThresholdDaysOk

`func (o *ForecastInstance) GetThresholdDaysOk() (*int64, bool)`

GetThresholdDaysOk returns a tuple with the ThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDays

`func (o *ForecastInstance) SetThresholdDays(v int64)`

SetThresholdDays sets ThresholdDays field to given value.

### HasThresholdDays

`func (o *ForecastInstance) HasThresholdDays() bool`

HasThresholdDays returns a boolean if a field has been set.

### GetForecastDef

`func (o *ForecastInstance) GetForecastDef() ForecastDefinitionRelationship`

GetForecastDef returns the ForecastDef field if non-nil, zero value otherwise.

### GetForecastDefOk

`func (o *ForecastInstance) GetForecastDefOk() (*ForecastDefinitionRelationship, bool)`

GetForecastDefOk returns a tuple with the ForecastDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastDef

`func (o *ForecastInstance) SetForecastDef(v ForecastDefinitionRelationship)`

SetForecastDef sets ForecastDef field to given value.

### HasForecastDef

`func (o *ForecastInstance) HasForecastDef() bool`

HasForecastDef returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ForecastInstance) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ForecastInstance) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ForecastInstance) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ForecastInstance) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


