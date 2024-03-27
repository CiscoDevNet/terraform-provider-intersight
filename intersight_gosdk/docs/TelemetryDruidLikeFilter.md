# TelemetryDruidLikeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Pattern** | **string** | String LIKE pattern, such as \&quot;foo%\&quot; or \&quot;___bar\&quot;. | 
**Escape** | Pointer to **string** | A string escape character that can be used to escape special characters. | [optional] 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidLikeFilter

`func NewTelemetryDruidLikeFilter(type_ string, dimension string, pattern string, ) *TelemetryDruidLikeFilter`

NewTelemetryDruidLikeFilter instantiates a new TelemetryDruidLikeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidLikeFilterWithDefaults

`func NewTelemetryDruidLikeFilterWithDefaults() *TelemetryDruidLikeFilter`

NewTelemetryDruidLikeFilterWithDefaults instantiates a new TelemetryDruidLikeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidLikeFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidLikeFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidLikeFilter) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidLikeFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidLikeFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidLikeFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPattern

`func (o *TelemetryDruidLikeFilter) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TelemetryDruidLikeFilter) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TelemetryDruidLikeFilter) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetEscape

`func (o *TelemetryDruidLikeFilter) GetEscape() string`

GetEscape returns the Escape field if non-nil, zero value otherwise.

### GetEscapeOk

`func (o *TelemetryDruidLikeFilter) GetEscapeOk() (*string, bool)`

GetEscapeOk returns a tuple with the Escape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscape

`func (o *TelemetryDruidLikeFilter) SetEscape(v string)`

SetEscape sets Escape field to given value.

### HasEscape

`func (o *TelemetryDruidLikeFilter) HasEscape() bool`

HasEscape returns a boolean if a field has been set.

### GetExtractionFn

`func (o *TelemetryDruidLikeFilter) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidLikeFilter) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidLikeFilter) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidLikeFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


