# TelemetryDruidLikeFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Pattern** | **string** | String LIKE pattern, such as \&quot;foo%\&quot; or \&quot;___bar\&quot;. | 
**Escape** | Pointer to **string** | A string escape character that can be used to escape special characters. | [optional] 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidLikeFilterAllOf

`func NewTelemetryDruidLikeFilterAllOf(type_ string, dimension string, pattern string, ) *TelemetryDruidLikeFilterAllOf`

NewTelemetryDruidLikeFilterAllOf instantiates a new TelemetryDruidLikeFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidLikeFilterAllOfWithDefaults

`func NewTelemetryDruidLikeFilterAllOfWithDefaults() *TelemetryDruidLikeFilterAllOf`

NewTelemetryDruidLikeFilterAllOfWithDefaults instantiates a new TelemetryDruidLikeFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidLikeFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidLikeFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidLikeFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidLikeFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidLikeFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidLikeFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPattern

`func (o *TelemetryDruidLikeFilterAllOf) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TelemetryDruidLikeFilterAllOf) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TelemetryDruidLikeFilterAllOf) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetEscape

`func (o *TelemetryDruidLikeFilterAllOf) GetEscape() string`

GetEscape returns the Escape field if non-nil, zero value otherwise.

### GetEscapeOk

`func (o *TelemetryDruidLikeFilterAllOf) GetEscapeOk() (*string, bool)`

GetEscapeOk returns a tuple with the Escape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscape

`func (o *TelemetryDruidLikeFilterAllOf) SetEscape(v string)`

SetEscape sets Escape field to given value.

### HasEscape

`func (o *TelemetryDruidLikeFilterAllOf) HasEscape() bool`

HasEscape returns a boolean if a field has been set.

### GetExtractionFn

`func (o *TelemetryDruidLikeFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidLikeFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidLikeFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidLikeFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


