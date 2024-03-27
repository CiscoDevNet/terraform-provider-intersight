# TelemetryDruidExtractionFunctionStringFormatAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Format** | **string** | A sprintf expression. For example if you want to concat \&quot;[\&quot; and \&quot;]\&quot; before and after the actual dimension value, you need to specify \&quot;[%s]\&quot; as format string. | 
**NullHandling** | Pointer to **string** |  | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionStringFormatAllOf

`func NewTelemetryDruidExtractionFunctionStringFormatAllOf(type_ string, format string, ) *TelemetryDruidExtractionFunctionStringFormatAllOf`

NewTelemetryDruidExtractionFunctionStringFormatAllOf instantiates a new TelemetryDruidExtractionFunctionStringFormatAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionStringFormatAllOfWithDefaults

`func NewTelemetryDruidExtractionFunctionStringFormatAllOfWithDefaults() *TelemetryDruidExtractionFunctionStringFormatAllOf`

NewTelemetryDruidExtractionFunctionStringFormatAllOfWithDefaults instantiates a new TelemetryDruidExtractionFunctionStringFormatAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetNullHandling() string`

GetNullHandling returns the NullHandling field if non-nil, zero value otherwise.

### GetNullHandlingOk

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) GetNullHandlingOk() (*string, bool)`

GetNullHandlingOk returns a tuple with the NullHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) SetNullHandling(v string)`

SetNullHandling sets NullHandling field to given value.

### HasNullHandling

`func (o *TelemetryDruidExtractionFunctionStringFormatAllOf) HasNullHandling() bool`

HasNullHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


