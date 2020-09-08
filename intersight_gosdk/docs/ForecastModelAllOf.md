# ForecastModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to **float32** | The standard error of the estimate is a measure of the accuracy of predictions from predective modeling. | [optional] 
**ModelData** | Pointer to **[]float32** |  | [optional] 
**ModelType** | Pointer to **string** | Model type indicating type of predictive model used for computing forecast. * &#x60;Linear&#x60; - The Enum value Linear represents that the predictive model type used for forecast computation is linear regression. | [optional] [default to "Linear"]

## Methods

### NewForecastModelAllOf

`func NewForecastModelAllOf() *ForecastModelAllOf`

NewForecastModelAllOf instantiates a new ForecastModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastModelAllOfWithDefaults

`func NewForecastModelAllOfWithDefaults() *ForecastModelAllOf`

NewForecastModelAllOfWithDefaults instantiates a new ForecastModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *ForecastModelAllOf) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *ForecastModelAllOf) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *ForecastModelAllOf) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *ForecastModelAllOf) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetModelData

`func (o *ForecastModelAllOf) GetModelData() []float32`

GetModelData returns the ModelData field if non-nil, zero value otherwise.

### GetModelDataOk

`func (o *ForecastModelAllOf) GetModelDataOk() (*[]float32, bool)`

GetModelDataOk returns a tuple with the ModelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelData

`func (o *ForecastModelAllOf) SetModelData(v []float32)`

SetModelData sets ModelData field to given value.

### HasModelData

`func (o *ForecastModelAllOf) HasModelData() bool`

HasModelData returns a boolean if a field has been set.

### GetModelType

`func (o *ForecastModelAllOf) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ForecastModelAllOf) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ForecastModelAllOf) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ForecastModelAllOf) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


