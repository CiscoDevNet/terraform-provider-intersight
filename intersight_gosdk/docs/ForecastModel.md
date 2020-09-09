# ForecastModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to **float32** | The standard error of the estimate is a measure of the accuracy of predictions from predective modeling. | [optional] 
**ModelData** | Pointer to **[]float32** |  | [optional] 
**ModelType** | Pointer to **string** | Model type indicating type of predictive model used for computing forecast. * &#x60;Linear&#x60; - The Enum value Linear represents that the predictive model type used for forecast computation is linear regression. | [optional] [default to "Linear"]

## Methods

### NewForecastModel

`func NewForecastModel() *ForecastModel`

NewForecastModel instantiates a new ForecastModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastModelWithDefaults

`func NewForecastModelWithDefaults() *ForecastModel`

NewForecastModelWithDefaults instantiates a new ForecastModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *ForecastModel) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *ForecastModel) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *ForecastModel) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *ForecastModel) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetModelData

`func (o *ForecastModel) GetModelData() []float32`

GetModelData returns the ModelData field if non-nil, zero value otherwise.

### GetModelDataOk

`func (o *ForecastModel) GetModelDataOk() (*[]float32, bool)`

GetModelDataOk returns a tuple with the ModelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelData

`func (o *ForecastModel) SetModelData(v []float32)`

SetModelData sets ModelData field to given value.

### HasModelData

`func (o *ForecastModel) HasModelData() bool`

HasModelData returns a boolean if a field has been set.

### GetModelType

`func (o *ForecastModel) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ForecastModel) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ForecastModel) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ForecastModel) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


