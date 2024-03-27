# TelemetryDruidSelectorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | The name of a dimension. | 
**Value** | **string** | The value of a dimension. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidSelectorFilter

`func NewTelemetryDruidSelectorFilter(type_ string, dimension string, value string, ) *TelemetryDruidSelectorFilter`

NewTelemetryDruidSelectorFilter instantiates a new TelemetryDruidSelectorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSelectorFilterWithDefaults

`func NewTelemetryDruidSelectorFilterWithDefaults() *TelemetryDruidSelectorFilter`

NewTelemetryDruidSelectorFilterWithDefaults instantiates a new TelemetryDruidSelectorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidSelectorFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidSelectorFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidSelectorFilter) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidSelectorFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidSelectorFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidSelectorFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValue

`func (o *TelemetryDruidSelectorFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidSelectorFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidSelectorFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetExtractionFn

`func (o *TelemetryDruidSelectorFilter) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidSelectorFilter) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidSelectorFilter) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidSelectorFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


