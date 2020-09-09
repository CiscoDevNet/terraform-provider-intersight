# TelemetryDruidRegexFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The filter type. | 
**ExtractionFn** | Pointer to **map[string]interface{}** | All filters except the \&quot;spatial\&quot; filter support extraction functions. An extraction function is defined by setting the \&quot;extractionFn\&quot; field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \&quot;bar_1\&quot;. | [optional] 
**Dimension** | **string** | null | 
**Pattern** | **string** | null | 

## Methods

### NewTelemetryDruidRegexFilter

`func NewTelemetryDruidRegexFilter(type_ string, dimension string, pattern string, ) *TelemetryDruidRegexFilter`

NewTelemetryDruidRegexFilter instantiates a new TelemetryDruidRegexFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidRegexFilterWithDefaults

`func NewTelemetryDruidRegexFilterWithDefaults() *TelemetryDruidRegexFilter`

NewTelemetryDruidRegexFilterWithDefaults instantiates a new TelemetryDruidRegexFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidRegexFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidRegexFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidRegexFilter) SetType(v string)`

SetType sets Type field to given value.


### GetExtractionFn

`func (o *TelemetryDruidRegexFilter) GetExtractionFn() map[string]interface{}`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidRegexFilter) GetExtractionFnOk() (*map[string]interface{}, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidRegexFilter) SetExtractionFn(v map[string]interface{})`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidRegexFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.

### GetDimension

`func (o *TelemetryDruidRegexFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidRegexFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidRegexFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPattern

`func (o *TelemetryDruidRegexFilter) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TelemetryDruidRegexFilter) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TelemetryDruidRegexFilter) SetPattern(v string)`

SetPattern sets Pattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


