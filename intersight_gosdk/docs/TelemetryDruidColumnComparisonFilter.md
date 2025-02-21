# TelemetryDruidColumnComparisonFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimensions** | [**[]TelemetryDruidDimensionSpec**](TelemetryDruidDimensionSpec.md) | A list of DimensionSpecs, making it possible to apply an extraction function if needed. | 

## Methods

### NewTelemetryDruidColumnComparisonFilter

`func NewTelemetryDruidColumnComparisonFilter(type_ string, dimensions []TelemetryDruidDimensionSpec, ) *TelemetryDruidColumnComparisonFilter`

NewTelemetryDruidColumnComparisonFilter instantiates a new TelemetryDruidColumnComparisonFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidColumnComparisonFilterWithDefaults

`func NewTelemetryDruidColumnComparisonFilterWithDefaults() *TelemetryDruidColumnComparisonFilter`

NewTelemetryDruidColumnComparisonFilterWithDefaults instantiates a new TelemetryDruidColumnComparisonFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidColumnComparisonFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidColumnComparisonFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidColumnComparisonFilter) SetType(v string)`

SetType sets Type field to given value.


### GetDimensions

`func (o *TelemetryDruidColumnComparisonFilter) GetDimensions() []TelemetryDruidDimensionSpec`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TelemetryDruidColumnComparisonFilter) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TelemetryDruidColumnComparisonFilter) SetDimensions(v []TelemetryDruidDimensionSpec)`

SetDimensions sets Dimensions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


