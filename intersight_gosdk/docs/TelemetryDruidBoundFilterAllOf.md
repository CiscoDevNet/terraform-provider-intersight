# TelemetryDruidBoundFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Lower** | Pointer to **string** | The lower bound string match value for the filter. | [optional] 
**Upper** | Pointer to **string** | The upper bound string match value for the filter. | [optional] 
**LowerStrict** | Pointer to **bool** | Boolean indicating whether to perform strict comparison on the lower bound (\&quot;&gt;\&quot; instead of \&quot;&gt;&#x3D;\&quot;). | [optional] 
**UpperStrict** | Pointer to **bool** | Boolean indicating whether to perform strict comparison on the upper bound (\&quot;&lt;\&quot; instead of \&quot;&lt;&#x3D;\&quot;). | [optional] 
**Ordering** | Pointer to **string** | String that specifies the sorting order to use when comparing values against the bound. lexicographic - Sorts values by converting Strings to their UTF-8 byte array representations and comparing lexicographically, byte-by-byte. alphanumeric - Suitable for strings with both numeric and non-numeric content, e.g., \&quot;file12 sorts after file2\&quot; This ordering is not suitable for numbers with decimal points or negative numbers. For example, \&quot;1.3\&quot; precedes \&quot;1.15\&quot; in this ordering because \&quot;15\&quot; has more significant digits than \&quot;3\&quot;. Negative numbers are sorted after positive numbers (because numeric characters precede the \&quot;-\&quot; in the negative numbers). numeric - Sorts values as numbers, supports integers and floating point values. Negative values are supported. This sorting order will try to parse all string values as numbers. Unparseable values are treated as nulls, and nulls precede numbers. When comparing two unparseable values (e.g., \&quot;hello\&quot; and \&quot;world\&quot;), this ordering will sort by comparing the unparsed strings lexicographically. strlen - Sorts values by their string lengths. When there is a tie, this comparator falls back to using the String compareTo method. version - Sorts values as versions, e.g., \&quot;10.0 sorts after 9.0\&quot;, \&quot;1.0.0-SNAPSHOT sorts after 1.0.0\&quot;. | [optional] 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidBoundFilterAllOf

`func NewTelemetryDruidBoundFilterAllOf(type_ string, dimension string, ) *TelemetryDruidBoundFilterAllOf`

NewTelemetryDruidBoundFilterAllOf instantiates a new TelemetryDruidBoundFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidBoundFilterAllOfWithDefaults

`func NewTelemetryDruidBoundFilterAllOfWithDefaults() *TelemetryDruidBoundFilterAllOf`

NewTelemetryDruidBoundFilterAllOfWithDefaults instantiates a new TelemetryDruidBoundFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidBoundFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidBoundFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidBoundFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidBoundFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidBoundFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidBoundFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetLower

`func (o *TelemetryDruidBoundFilterAllOf) GetLower() string`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *TelemetryDruidBoundFilterAllOf) GetLowerOk() (*string, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *TelemetryDruidBoundFilterAllOf) SetLower(v string)`

SetLower sets Lower field to given value.

### HasLower

`func (o *TelemetryDruidBoundFilterAllOf) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetUpper

`func (o *TelemetryDruidBoundFilterAllOf) GetUpper() string`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *TelemetryDruidBoundFilterAllOf) GetUpperOk() (*string, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *TelemetryDruidBoundFilterAllOf) SetUpper(v string)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *TelemetryDruidBoundFilterAllOf) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### GetLowerStrict

`func (o *TelemetryDruidBoundFilterAllOf) GetLowerStrict() bool`

GetLowerStrict returns the LowerStrict field if non-nil, zero value otherwise.

### GetLowerStrictOk

`func (o *TelemetryDruidBoundFilterAllOf) GetLowerStrictOk() (*bool, bool)`

GetLowerStrictOk returns a tuple with the LowerStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerStrict

`func (o *TelemetryDruidBoundFilterAllOf) SetLowerStrict(v bool)`

SetLowerStrict sets LowerStrict field to given value.

### HasLowerStrict

`func (o *TelemetryDruidBoundFilterAllOf) HasLowerStrict() bool`

HasLowerStrict returns a boolean if a field has been set.

### GetUpperStrict

`func (o *TelemetryDruidBoundFilterAllOf) GetUpperStrict() bool`

GetUpperStrict returns the UpperStrict field if non-nil, zero value otherwise.

### GetUpperStrictOk

`func (o *TelemetryDruidBoundFilterAllOf) GetUpperStrictOk() (*bool, bool)`

GetUpperStrictOk returns a tuple with the UpperStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperStrict

`func (o *TelemetryDruidBoundFilterAllOf) SetUpperStrict(v bool)`

SetUpperStrict sets UpperStrict field to given value.

### HasUpperStrict

`func (o *TelemetryDruidBoundFilterAllOf) HasUpperStrict() bool`

HasUpperStrict returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidBoundFilterAllOf) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidBoundFilterAllOf) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidBoundFilterAllOf) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidBoundFilterAllOf) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetExtractionFn

`func (o *TelemetryDruidBoundFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidBoundFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidBoundFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidBoundFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


