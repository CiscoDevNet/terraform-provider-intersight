# TelemetryDruidFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The filter type. | 
**ExtractionFn** | Pointer to **map[string]interface{}** | All filters except the \&quot;spatial\&quot; filter support extraction functions. An extraction function is defined by setting the \&quot;extractionFn\&quot; field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \&quot;bar_1\&quot;. | [optional] 
**Dimension** | **string** | null | 
**Value** | **string** | The value of a dimension. | 
**Dimensions** | [**[]TelemetryDruidDimensionSpec**](TelemetryDruidDimensionSpec.md) | A list of DimensionSpecs, making it possible to apply an extraction function if needed. | 
**Pattern** | **string** | null | 
**Fields** | [**[]TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Field** | [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 

## Methods

### NewTelemetryDruidFilter

`func NewTelemetryDruidFilter(type_ string, dimension string, value string, dimensions []TelemetryDruidDimensionSpec, pattern string, fields []TelemetryDruidFilter, field TelemetryDruidFilter, ) *TelemetryDruidFilter`

NewTelemetryDruidFilter instantiates a new TelemetryDruidFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidFilterWithDefaults

`func NewTelemetryDruidFilterWithDefaults() *TelemetryDruidFilter`

NewTelemetryDruidFilterWithDefaults instantiates a new TelemetryDruidFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidFilter) SetType(v string)`

SetType sets Type field to given value.


### GetExtractionFn

`func (o *TelemetryDruidFilter) GetExtractionFn() map[string]interface{}`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidFilter) GetExtractionFnOk() (*map[string]interface{}, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidFilter) SetExtractionFn(v map[string]interface{})`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.

### GetDimension

`func (o *TelemetryDruidFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValue

`func (o *TelemetryDruidFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetDimensions

`func (o *TelemetryDruidFilter) GetDimensions() []TelemetryDruidDimensionSpec`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TelemetryDruidFilter) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TelemetryDruidFilter) SetDimensions(v []TelemetryDruidDimensionSpec)`

SetDimensions sets Dimensions field to given value.


### GetPattern

`func (o *TelemetryDruidFilter) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TelemetryDruidFilter) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TelemetryDruidFilter) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetFields

`func (o *TelemetryDruidFilter) GetFields() []TelemetryDruidFilter`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidFilter) GetFieldsOk() (*[]TelemetryDruidFilter, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidFilter) SetFields(v []TelemetryDruidFilter)`

SetFields sets Fields field to given value.


### GetField

`func (o *TelemetryDruidFilter) GetField() TelemetryDruidFilter`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TelemetryDruidFilter) GetFieldOk() (*TelemetryDruidFilter, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TelemetryDruidFilter) SetField(v TelemetryDruidFilter)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


