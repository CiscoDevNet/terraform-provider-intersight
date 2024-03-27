# TelemetryDruidRegexFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Pattern** | **string** | String pattern to match - any standard Java regular expression. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidRegexFilterAllOf

`func NewTelemetryDruidRegexFilterAllOf(type_ string, dimension string, pattern string, ) *TelemetryDruidRegexFilterAllOf`

NewTelemetryDruidRegexFilterAllOf instantiates a new TelemetryDruidRegexFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidRegexFilterAllOfWithDefaults

`func NewTelemetryDruidRegexFilterAllOfWithDefaults() *TelemetryDruidRegexFilterAllOf`

NewTelemetryDruidRegexFilterAllOfWithDefaults instantiates a new TelemetryDruidRegexFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidRegexFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidRegexFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidRegexFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidRegexFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidRegexFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidRegexFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPattern

`func (o *TelemetryDruidRegexFilterAllOf) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TelemetryDruidRegexFilterAllOf) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TelemetryDruidRegexFilterAllOf) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetExtractionFn

`func (o *TelemetryDruidRegexFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidRegexFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidRegexFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidRegexFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


