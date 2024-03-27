# TelemetryDruidExtractionFunctionInlineLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Lookup** | Pointer to [**TelemetryDruidExtractionFunctionInlineLookupAllOfLookup**](TelemetryDruidExtractionFunctionInlineLookupAllOfLookup.md) |  | [optional] 
**RetainMissingValue** | Pointer to **bool** | Provides a hint how to handle missing values. Setting retainMissingValue to true will use the dimension&#39;s original value if it is not found in the lookup. The default values are replaceMissingValueWith &#x3D; null and retainMissingValue &#x3D; false which causes missing values to be treated as missing. It is illegal to set retainMissingValue &#x3D; true and also specify a replaceMissingValueWith. | [optional] 
**ReplaceMissingValueWith** | Pointer to **string** | Provides a hint how to handle missing values. Setting replaceMissingValueWith to \&quot;\&quot; has the same effect as setting it to null or omitting the property. | [optional] 
**Injective** | Pointer to **bool** | Override the lookup&#39;s own sense of whether or not it is injective. | [optional] 
**Optimize** | Pointer to **bool** | Allow optimization of lookup based extraction filter (by default optimize &#x3D; true). The optimization layer will run on the Broker and it will rewrite the extraction filter as clause of selector filters. | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionInlineLookup

`func NewTelemetryDruidExtractionFunctionInlineLookup(type_ string, ) *TelemetryDruidExtractionFunctionInlineLookup`

NewTelemetryDruidExtractionFunctionInlineLookup instantiates a new TelemetryDruidExtractionFunctionInlineLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionInlineLookupWithDefaults

`func NewTelemetryDruidExtractionFunctionInlineLookupWithDefaults() *TelemetryDruidExtractionFunctionInlineLookup`

NewTelemetryDruidExtractionFunctionInlineLookupWithDefaults instantiates a new TelemetryDruidExtractionFunctionInlineLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetType(v string)`

SetType sets Type field to given value.


### GetLookup

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetLookup() TelemetryDruidExtractionFunctionInlineLookupAllOfLookup`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetLookupOk() (*TelemetryDruidExtractionFunctionInlineLookupAllOfLookup, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetLookup(v TelemetryDruidExtractionFunctionInlineLookupAllOfLookup)`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *TelemetryDruidExtractionFunctionInlineLookup) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### GetRetainMissingValue

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetRetainMissingValue() bool`

GetRetainMissingValue returns the RetainMissingValue field if non-nil, zero value otherwise.

### GetRetainMissingValueOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetRetainMissingValueOk() (*bool, bool)`

GetRetainMissingValueOk returns a tuple with the RetainMissingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainMissingValue

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetRetainMissingValue(v bool)`

SetRetainMissingValue sets RetainMissingValue field to given value.

### HasRetainMissingValue

`func (o *TelemetryDruidExtractionFunctionInlineLookup) HasRetainMissingValue() bool`

HasRetainMissingValue returns a boolean if a field has been set.

### GetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetReplaceMissingValueWith() string`

GetReplaceMissingValueWith returns the ReplaceMissingValueWith field if non-nil, zero value otherwise.

### GetReplaceMissingValueWithOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetReplaceMissingValueWithOk() (*string, bool)`

GetReplaceMissingValueWithOk returns a tuple with the ReplaceMissingValueWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetReplaceMissingValueWith(v string)`

SetReplaceMissingValueWith sets ReplaceMissingValueWith field to given value.

### HasReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionInlineLookup) HasReplaceMissingValueWith() bool`

HasReplaceMissingValueWith returns a boolean if a field has been set.

### GetInjective

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetInjective() bool`

GetInjective returns the Injective field if non-nil, zero value otherwise.

### GetInjectiveOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetInjectiveOk() (*bool, bool)`

GetInjectiveOk returns a tuple with the Injective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjective

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetInjective(v bool)`

SetInjective sets Injective field to given value.

### HasInjective

`func (o *TelemetryDruidExtractionFunctionInlineLookup) HasInjective() bool`

HasInjective returns a boolean if a field has been set.

### GetOptimize

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetOptimize() bool`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *TelemetryDruidExtractionFunctionInlineLookup) GetOptimizeOk() (*bool, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *TelemetryDruidExtractionFunctionInlineLookup) SetOptimize(v bool)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *TelemetryDruidExtractionFunctionInlineLookup) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


