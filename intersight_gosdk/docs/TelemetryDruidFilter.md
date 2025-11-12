# TelemetryDruidFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Value** | **string** | The value of a dimension. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 
**Column** | **string** | Input column or virtual column name to filter. | 
**MatchValueType** | **string** | String specifying the type of bounds to match. For example STRING, LONG, DOUBLE, FLOAT, ARRAY&lt;STRING&gt;, ARRAY&lt;LONG&gt;, or any other Druid type. The matchValueType determines how Druid interprets the matchValue to assist in converting to the type of the matched column and also defines the type of comparison used when matching values. | 
**MatchValue** | **interface{}** | Value to match, must not be null. | 
**Dimensions** | [**[]TelemetryDruidDimensionSpec**](TelemetryDruidDimensionSpec.md) | A list of DimensionSpecs, making it possible to apply an extraction function if needed. | 
**Fields** | [**[]TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Field** | [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Values** | **[]string** | List of string value to match. | 
**Lower** | Pointer to **interface{}** | Lower bound value to match. At least one of lower or upper must not be null. | [optional] 
**Upper** | Pointer to **interface{}** | Upper bound value to match. At least one of lower or upper must not be null. | [optional] 
**LowerStrict** | Pointer to **bool** | Boolean indicating whether to perform strict comparison on the lower bound (\&quot;&gt;\&quot; instead of \&quot;&gt;&#x3D;\&quot;). | [optional] 
**UpperStrict** | Pointer to **bool** | Boolean indicating whether to perform strict comparison on the upper bound (\&quot;&lt;\&quot; instead of \&quot;&lt;&#x3D;\&quot;). | [optional] 
**Ordering** | Pointer to **string** | String that specifies the sorting order to use when comparing values against the bound. lexicographic - Sorts values by converting Strings to their UTF-8 byte array representations and comparing lexicographically, byte-by-byte. alphanumeric - Suitable for strings with both numeric and non-numeric content, e.g., \&quot;file12 sorts after file2\&quot; This ordering is not suitable for numbers with decimal points or negative numbers. For example, \&quot;1.3\&quot; precedes \&quot;1.15\&quot; in this ordering because \&quot;15\&quot; has more significant digits than \&quot;3\&quot;. Negative numbers are sorted after positive numbers (because numeric characters precede the \&quot;-\&quot; in the negative numbers). numeric - Sorts values as numbers, supports integers and floating point values. Negative values are supported. This sorting order will try to parse all string values as numbers. Unparseable values are treated as nulls, and nulls precede numbers. When comparing two unparseable values (e.g., \&quot;hello\&quot; and \&quot;world\&quot;), this ordering will sort by comparing the unparsed strings lexicographically. strlen - Sorts values by their string lengths. When there is a tie, this comparator falls back to using the String compareTo method. version - Sorts values as versions, e.g., \&quot;10.0 sorts after 9.0\&quot;, \&quot;1.0.0-SNAPSHOT sorts after 1.0.0\&quot;. | [optional] 
**LowerOpen** | Pointer to **bool** | Boolean indicating if lower bound is open in the interval of values defined by the range (\&quot;&gt;\&quot; instead of \&quot;&gt;&#x3D;\&quot;). | [optional] 
**UpperOpen** | Pointer to **bool** | Boolean indicating if upper bound is open on the interval of values defined by the range (\&quot;&lt;\&quot; instead of \&quot;&lt;&#x3D;\&quot;). | [optional] 
**Pattern** | **string** | String pattern to match - any standard Java regular expression. | 
**Escape** | Pointer to **string** | A string escape character that can be used to escape special characters. | [optional] 
**Intervals** | **[]string** | A JSON array containing ISO-8601 interval strings that defines the time ranges to filter on. | 
**Query** | [**TelemetryDruidQuerySpec**](TelemetryDruidQuerySpec.md) |  | 
**Expression** | **string** | Expression string to evaluate into true or false. See the Druid expression system for more details. | 

## Methods

### NewTelemetryDruidFilter

`func NewTelemetryDruidFilter(type_ string, dimension string, value string, column string, matchValueType string, matchValue interface{}, dimensions []TelemetryDruidDimensionSpec, fields []TelemetryDruidFilter, field TelemetryDruidFilter, values []string, pattern string, intervals []string, query TelemetryDruidQuerySpec, expression string, ) *TelemetryDruidFilter`

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


### GetExtractionFn

`func (o *TelemetryDruidFilter) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidFilter) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidFilter) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.

### GetColumn

`func (o *TelemetryDruidFilter) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *TelemetryDruidFilter) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *TelemetryDruidFilter) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetMatchValueType

`func (o *TelemetryDruidFilter) GetMatchValueType() string`

GetMatchValueType returns the MatchValueType field if non-nil, zero value otherwise.

### GetMatchValueTypeOk

`func (o *TelemetryDruidFilter) GetMatchValueTypeOk() (*string, bool)`

GetMatchValueTypeOk returns a tuple with the MatchValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValueType

`func (o *TelemetryDruidFilter) SetMatchValueType(v string)`

SetMatchValueType sets MatchValueType field to given value.


### GetMatchValue

`func (o *TelemetryDruidFilter) GetMatchValue() interface{}`

GetMatchValue returns the MatchValue field if non-nil, zero value otherwise.

### GetMatchValueOk

`func (o *TelemetryDruidFilter) GetMatchValueOk() (*interface{}, bool)`

GetMatchValueOk returns a tuple with the MatchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValue

`func (o *TelemetryDruidFilter) SetMatchValue(v interface{})`

SetMatchValue sets MatchValue field to given value.


### SetMatchValueNil

`func (o *TelemetryDruidFilter) SetMatchValueNil(b bool)`

 SetMatchValueNil sets the value for MatchValue to be an explicit nil

### UnsetMatchValue
`func (o *TelemetryDruidFilter) UnsetMatchValue()`

UnsetMatchValue ensures that no value is present for MatchValue, not even an explicit nil
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


### GetValues

`func (o *TelemetryDruidFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidFilter) SetValues(v []string)`

SetValues sets Values field to given value.


### GetLower

`func (o *TelemetryDruidFilter) GetLower() interface{}`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *TelemetryDruidFilter) GetLowerOk() (*interface{}, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *TelemetryDruidFilter) SetLower(v interface{})`

SetLower sets Lower field to given value.

### HasLower

`func (o *TelemetryDruidFilter) HasLower() bool`

HasLower returns a boolean if a field has been set.

### SetLowerNil

`func (o *TelemetryDruidFilter) SetLowerNil(b bool)`

 SetLowerNil sets the value for Lower to be an explicit nil

### UnsetLower
`func (o *TelemetryDruidFilter) UnsetLower()`

UnsetLower ensures that no value is present for Lower, not even an explicit nil
### GetUpper

`func (o *TelemetryDruidFilter) GetUpper() interface{}`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *TelemetryDruidFilter) GetUpperOk() (*interface{}, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *TelemetryDruidFilter) SetUpper(v interface{})`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *TelemetryDruidFilter) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### SetUpperNil

`func (o *TelemetryDruidFilter) SetUpperNil(b bool)`

 SetUpperNil sets the value for Upper to be an explicit nil

### UnsetUpper
`func (o *TelemetryDruidFilter) UnsetUpper()`

UnsetUpper ensures that no value is present for Upper, not even an explicit nil
### GetLowerStrict

`func (o *TelemetryDruidFilter) GetLowerStrict() bool`

GetLowerStrict returns the LowerStrict field if non-nil, zero value otherwise.

### GetLowerStrictOk

`func (o *TelemetryDruidFilter) GetLowerStrictOk() (*bool, bool)`

GetLowerStrictOk returns a tuple with the LowerStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerStrict

`func (o *TelemetryDruidFilter) SetLowerStrict(v bool)`

SetLowerStrict sets LowerStrict field to given value.

### HasLowerStrict

`func (o *TelemetryDruidFilter) HasLowerStrict() bool`

HasLowerStrict returns a boolean if a field has been set.

### GetUpperStrict

`func (o *TelemetryDruidFilter) GetUpperStrict() bool`

GetUpperStrict returns the UpperStrict field if non-nil, zero value otherwise.

### GetUpperStrictOk

`func (o *TelemetryDruidFilter) GetUpperStrictOk() (*bool, bool)`

GetUpperStrictOk returns a tuple with the UpperStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperStrict

`func (o *TelemetryDruidFilter) SetUpperStrict(v bool)`

SetUpperStrict sets UpperStrict field to given value.

### HasUpperStrict

`func (o *TelemetryDruidFilter) HasUpperStrict() bool`

HasUpperStrict returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidFilter) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidFilter) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidFilter) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidFilter) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetLowerOpen

`func (o *TelemetryDruidFilter) GetLowerOpen() bool`

GetLowerOpen returns the LowerOpen field if non-nil, zero value otherwise.

### GetLowerOpenOk

`func (o *TelemetryDruidFilter) GetLowerOpenOk() (*bool, bool)`

GetLowerOpenOk returns a tuple with the LowerOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerOpen

`func (o *TelemetryDruidFilter) SetLowerOpen(v bool)`

SetLowerOpen sets LowerOpen field to given value.

### HasLowerOpen

`func (o *TelemetryDruidFilter) HasLowerOpen() bool`

HasLowerOpen returns a boolean if a field has been set.

### GetUpperOpen

`func (o *TelemetryDruidFilter) GetUpperOpen() bool`

GetUpperOpen returns the UpperOpen field if non-nil, zero value otherwise.

### GetUpperOpenOk

`func (o *TelemetryDruidFilter) GetUpperOpenOk() (*bool, bool)`

GetUpperOpenOk returns a tuple with the UpperOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperOpen

`func (o *TelemetryDruidFilter) SetUpperOpen(v bool)`

SetUpperOpen sets UpperOpen field to given value.

### HasUpperOpen

`func (o *TelemetryDruidFilter) HasUpperOpen() bool`

HasUpperOpen returns a boolean if a field has been set.

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


### GetEscape

`func (o *TelemetryDruidFilter) GetEscape() string`

GetEscape returns the Escape field if non-nil, zero value otherwise.

### GetEscapeOk

`func (o *TelemetryDruidFilter) GetEscapeOk() (*string, bool)`

GetEscapeOk returns a tuple with the Escape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscape

`func (o *TelemetryDruidFilter) SetEscape(v string)`

SetEscape sets Escape field to given value.

### HasEscape

`func (o *TelemetryDruidFilter) HasEscape() bool`

HasEscape returns a boolean if a field has been set.

### GetIntervals

`func (o *TelemetryDruidFilter) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidFilter) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidFilter) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetQuery

`func (o *TelemetryDruidFilter) GetQuery() TelemetryDruidQuerySpec`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidFilter) GetQueryOk() (*TelemetryDruidQuerySpec, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidFilter) SetQuery(v TelemetryDruidQuerySpec)`

SetQuery sets Query field to given value.


### GetExpression

`func (o *TelemetryDruidFilter) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidFilter) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidFilter) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


