# TelemetryDruidRangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Column** | **string** | Input column or virtual column name to filter. | 
**MatchValueType** | **string** | String specifying the type of bounds to match. For example STRING, LONG, DOUBLE, FLOAT, ARRAY&lt;STRING&gt;, ARRAY&lt;LONG&gt;, or any other Druid type. The matchValueType determines how Druid interprets the matchValue to assist in converting to the type of the matched column and also defines the type of comparison used when matching values. | 
**Lower** | Pointer to **interface{}** | Lower bound value to match. At least one of lower or upper must not be null. | [optional] 
**Upper** | Pointer to **interface{}** | Upper bound value to match. At least one of lower or upper must not be null. | [optional] 
**LowerOpen** | Pointer to **bool** | Boolean indicating if lower bound is open in the interval of values defined by the range (\&quot;&gt;\&quot; instead of \&quot;&gt;&#x3D;\&quot;). | [optional] 
**UpperOpen** | Pointer to **bool** | Boolean indicating if upper bound is open on the interval of values defined by the range (\&quot;&lt;\&quot; instead of \&quot;&lt;&#x3D;\&quot;). | [optional] 

## Methods

### NewTelemetryDruidRangeFilter

`func NewTelemetryDruidRangeFilter(type_ string, column string, matchValueType string, ) *TelemetryDruidRangeFilter`

NewTelemetryDruidRangeFilter instantiates a new TelemetryDruidRangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidRangeFilterWithDefaults

`func NewTelemetryDruidRangeFilterWithDefaults() *TelemetryDruidRangeFilter`

NewTelemetryDruidRangeFilterWithDefaults instantiates a new TelemetryDruidRangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidRangeFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidRangeFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidRangeFilter) SetType(v string)`

SetType sets Type field to given value.


### GetColumn

`func (o *TelemetryDruidRangeFilter) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *TelemetryDruidRangeFilter) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *TelemetryDruidRangeFilter) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetMatchValueType

`func (o *TelemetryDruidRangeFilter) GetMatchValueType() string`

GetMatchValueType returns the MatchValueType field if non-nil, zero value otherwise.

### GetMatchValueTypeOk

`func (o *TelemetryDruidRangeFilter) GetMatchValueTypeOk() (*string, bool)`

GetMatchValueTypeOk returns a tuple with the MatchValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValueType

`func (o *TelemetryDruidRangeFilter) SetMatchValueType(v string)`

SetMatchValueType sets MatchValueType field to given value.


### GetLower

`func (o *TelemetryDruidRangeFilter) GetLower() interface{}`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *TelemetryDruidRangeFilter) GetLowerOk() (*interface{}, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *TelemetryDruidRangeFilter) SetLower(v interface{})`

SetLower sets Lower field to given value.

### HasLower

`func (o *TelemetryDruidRangeFilter) HasLower() bool`

HasLower returns a boolean if a field has been set.

### SetLowerNil

`func (o *TelemetryDruidRangeFilter) SetLowerNil(b bool)`

 SetLowerNil sets the value for Lower to be an explicit nil

### UnsetLower
`func (o *TelemetryDruidRangeFilter) UnsetLower()`

UnsetLower ensures that no value is present for Lower, not even an explicit nil
### GetUpper

`func (o *TelemetryDruidRangeFilter) GetUpper() interface{}`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *TelemetryDruidRangeFilter) GetUpperOk() (*interface{}, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *TelemetryDruidRangeFilter) SetUpper(v interface{})`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *TelemetryDruidRangeFilter) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### SetUpperNil

`func (o *TelemetryDruidRangeFilter) SetUpperNil(b bool)`

 SetUpperNil sets the value for Upper to be an explicit nil

### UnsetUpper
`func (o *TelemetryDruidRangeFilter) UnsetUpper()`

UnsetUpper ensures that no value is present for Upper, not even an explicit nil
### GetLowerOpen

`func (o *TelemetryDruidRangeFilter) GetLowerOpen() bool`

GetLowerOpen returns the LowerOpen field if non-nil, zero value otherwise.

### GetLowerOpenOk

`func (o *TelemetryDruidRangeFilter) GetLowerOpenOk() (*bool, bool)`

GetLowerOpenOk returns a tuple with the LowerOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerOpen

`func (o *TelemetryDruidRangeFilter) SetLowerOpen(v bool)`

SetLowerOpen sets LowerOpen field to given value.

### HasLowerOpen

`func (o *TelemetryDruidRangeFilter) HasLowerOpen() bool`

HasLowerOpen returns a boolean if a field has been set.

### GetUpperOpen

`func (o *TelemetryDruidRangeFilter) GetUpperOpen() bool`

GetUpperOpen returns the UpperOpen field if non-nil, zero value otherwise.

### GetUpperOpenOk

`func (o *TelemetryDruidRangeFilter) GetUpperOpenOk() (*bool, bool)`

GetUpperOpenOk returns a tuple with the UpperOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperOpen

`func (o *TelemetryDruidRangeFilter) SetUpperOpen(v bool)`

SetUpperOpen sets UpperOpen field to given value.

### HasUpperOpen

`func (o *TelemetryDruidRangeFilter) HasUpperOpen() bool`

HasUpperOpen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


