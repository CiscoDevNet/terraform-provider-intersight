# TelemetryDruidExtractionFunctionRegexAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expr** | **string** |  | 
**Index** | Pointer to **interface{}** | group to extract, default 1. If \&quot;index\&quot; is set, it will control which group from the match to extract. Index zero extracts the string matching the entire pattern. | [optional] 
**ReplaceMissingValue** | Pointer to **bool** | If the replaceMissingValue property is true, the extraction function will transform dimension values that do not match the regex pattern to a user-specified String. Default value is false. | [optional] 
**ReplaceMissingValueWith** | Pointer to **string** | The replaceMissingValueWith property sets the String that unmatched dimension values will be replaced with, if replaceMissingValue is true. If replaceMissingValueWith is not specified, unmatched dimension values will be replaced with nulls. | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionRegexAllOf

`func NewTelemetryDruidExtractionFunctionRegexAllOf(type_ string, expr string, ) *TelemetryDruidExtractionFunctionRegexAllOf`

NewTelemetryDruidExtractionFunctionRegexAllOf instantiates a new TelemetryDruidExtractionFunctionRegexAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionRegexAllOfWithDefaults

`func NewTelemetryDruidExtractionFunctionRegexAllOfWithDefaults() *TelemetryDruidExtractionFunctionRegexAllOf`

NewTelemetryDruidExtractionFunctionRegexAllOfWithDefaults instantiates a new TelemetryDruidExtractionFunctionRegexAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetExpr

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetIndex

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetIndex() interface{}`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetIndexOk() (*interface{}, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetIndex(v interface{})`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *TelemetryDruidExtractionFunctionRegexAllOf) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetReplaceMissingValue

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetReplaceMissingValue() bool`

GetReplaceMissingValue returns the ReplaceMissingValue field if non-nil, zero value otherwise.

### GetReplaceMissingValueOk

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetReplaceMissingValueOk() (*bool, bool)`

GetReplaceMissingValueOk returns a tuple with the ReplaceMissingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingValue

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetReplaceMissingValue(v bool)`

SetReplaceMissingValue sets ReplaceMissingValue field to given value.

### HasReplaceMissingValue

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) HasReplaceMissingValue() bool`

HasReplaceMissingValue returns a boolean if a field has been set.

### GetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetReplaceMissingValueWith() string`

GetReplaceMissingValueWith returns the ReplaceMissingValueWith field if non-nil, zero value otherwise.

### GetReplaceMissingValueWithOk

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) GetReplaceMissingValueWithOk() (*string, bool)`

GetReplaceMissingValueWithOk returns a tuple with the ReplaceMissingValueWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) SetReplaceMissingValueWith(v string)`

SetReplaceMissingValueWith sets ReplaceMissingValueWith field to given value.

### HasReplaceMissingValueWith

`func (o *TelemetryDruidExtractionFunctionRegexAllOf) HasReplaceMissingValueWith() bool`

HasReplaceMissingValueWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


