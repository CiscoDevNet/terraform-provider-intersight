# TelemetryDruidInFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Values** | **[]string** | List of string value to match. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidInFilterAllOf

`func NewTelemetryDruidInFilterAllOf(type_ string, dimension string, values []string, ) *TelemetryDruidInFilterAllOf`

NewTelemetryDruidInFilterAllOf instantiates a new TelemetryDruidInFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidInFilterAllOfWithDefaults

`func NewTelemetryDruidInFilterAllOfWithDefaults() *TelemetryDruidInFilterAllOf`

NewTelemetryDruidInFilterAllOfWithDefaults instantiates a new TelemetryDruidInFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidInFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidInFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidInFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidInFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidInFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidInFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValues

`func (o *TelemetryDruidInFilterAllOf) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidInFilterAllOf) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidInFilterAllOf) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExtractionFn

`func (o *TelemetryDruidInFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidInFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidInFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidInFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


