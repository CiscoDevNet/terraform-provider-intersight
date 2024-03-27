# TelemetryDruidExtractionFunctionStringFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Format** | **string** | A sprintf expression. For example if you want to concat \&quot;[\&quot; and \&quot;]\&quot; before and after the actual dimension value, you need to specify \&quot;[%s]\&quot; as format string. | 
**NullHandling** | Pointer to **string** |  | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionStringFormat

`func NewTelemetryDruidExtractionFunctionStringFormat(type_ string, format string, ) *TelemetryDruidExtractionFunctionStringFormat`

NewTelemetryDruidExtractionFunctionStringFormat instantiates a new TelemetryDruidExtractionFunctionStringFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionStringFormatWithDefaults

`func NewTelemetryDruidExtractionFunctionStringFormatWithDefaults() *TelemetryDruidExtractionFunctionStringFormat`

NewTelemetryDruidExtractionFunctionStringFormatWithDefaults instantiates a new TelemetryDruidExtractionFunctionStringFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionStringFormat) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TelemetryDruidExtractionFunctionStringFormat) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetNullHandling() string`

GetNullHandling returns the NullHandling field if non-nil, zero value otherwise.

### GetNullHandlingOk

`func (o *TelemetryDruidExtractionFunctionStringFormat) GetNullHandlingOk() (*string, bool)`

GetNullHandlingOk returns a tuple with the NullHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormat) SetNullHandling(v string)`

SetNullHandling sets NullHandling field to given value.

### HasNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormat) HasNullHandling() bool`

HasNullHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


