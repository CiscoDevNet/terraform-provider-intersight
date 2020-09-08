# TelemetryDruidDimensionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the dimension spec type. | 
**Dimension** | **string** | null | 
**OutputName** | **string** | null | 
**OutputType** | **string** | null | [default to "STRING"]
**ExtractionFn** | **map[string]interface{}** | All filters except the \&quot;spatial\&quot; filter support extraction functions. An extraction function is defined by setting the \&quot;extractionFn\&quot; field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \&quot;bar_1\&quot;. | 

## Methods

### NewTelemetryDruidDimensionSpec

`func NewTelemetryDruidDimensionSpec(type_ string, dimension string, outputName string, outputType string, extractionFn map[string]interface{}, ) *TelemetryDruidDimensionSpec`

NewTelemetryDruidDimensionSpec instantiates a new TelemetryDruidDimensionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDimensionSpecWithDefaults

`func NewTelemetryDruidDimensionSpecWithDefaults() *TelemetryDruidDimensionSpec`

NewTelemetryDruidDimensionSpecWithDefaults instantiates a new TelemetryDruidDimensionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidDimensionSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidDimensionSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidDimensionSpec) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidDimensionSpec) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidDimensionSpec) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidDimensionSpec) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetOutputName

`func (o *TelemetryDruidDimensionSpec) GetOutputName() string`

GetOutputName returns the OutputName field if non-nil, zero value otherwise.

### GetOutputNameOk

`func (o *TelemetryDruidDimensionSpec) GetOutputNameOk() (*string, bool)`

GetOutputNameOk returns a tuple with the OutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputName

`func (o *TelemetryDruidDimensionSpec) SetOutputName(v string)`

SetOutputName sets OutputName field to given value.


### GetOutputType

`func (o *TelemetryDruidDimensionSpec) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *TelemetryDruidDimensionSpec) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *TelemetryDruidDimensionSpec) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.


### GetExtractionFn

`func (o *TelemetryDruidDimensionSpec) GetExtractionFn() map[string]interface{}`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidDimensionSpec) GetExtractionFnOk() (*map[string]interface{}, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidDimensionSpec) SetExtractionFn(v map[string]interface{})`

SetExtractionFn sets ExtractionFn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


