# TelemetryDruidExtractionFunctionSubstring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Index** | **int32** |  | 
**Length** | Pointer to **int32** | The length may be omitted for substring to return the remainder of the dimension value starting from index, or null if index greater than the length of the dimension value. | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionSubstring

`func NewTelemetryDruidExtractionFunctionSubstring(type_ string, index int32, ) *TelemetryDruidExtractionFunctionSubstring`

NewTelemetryDruidExtractionFunctionSubstring instantiates a new TelemetryDruidExtractionFunctionSubstring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionSubstringWithDefaults

`func NewTelemetryDruidExtractionFunctionSubstringWithDefaults() *TelemetryDruidExtractionFunctionSubstring`

NewTelemetryDruidExtractionFunctionSubstringWithDefaults instantiates a new TelemetryDruidExtractionFunctionSubstring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionSubstring) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionSubstring) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionSubstring) SetType(v string)`

SetType sets Type field to given value.


### GetIndex

`func (o *TelemetryDruidExtractionFunctionSubstring) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TelemetryDruidExtractionFunctionSubstring) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TelemetryDruidExtractionFunctionSubstring) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLength

`func (o *TelemetryDruidExtractionFunctionSubstring) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *TelemetryDruidExtractionFunctionSubstring) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *TelemetryDruidExtractionFunctionSubstring) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *TelemetryDruidExtractionFunctionSubstring) HasLength() bool`

HasLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


