# TelemetryDruidInFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Values** | **[]string** | List of string value to match. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidInFilter

`func NewTelemetryDruidInFilter(type_ string, dimension string, values []string, ) *TelemetryDruidInFilter`

NewTelemetryDruidInFilter instantiates a new TelemetryDruidInFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidInFilterWithDefaults

`func NewTelemetryDruidInFilterWithDefaults() *TelemetryDruidInFilter`

NewTelemetryDruidInFilterWithDefaults instantiates a new TelemetryDruidInFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidInFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidInFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidInFilter) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidInFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidInFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidInFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValues

`func (o *TelemetryDruidInFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidInFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidInFilter) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExtractionFn

`func (o *TelemetryDruidInFilter) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidInFilter) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidInFilter) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidInFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


