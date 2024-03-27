# TelemetryDruidEqualityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Column** | **string** | Input column or virtual column name to filter. | 
**MatchValueType** | **string** | String specifying the type of value to match. For example STRING, LONG, DOUBLE, FLOAT, ARRAY&lt;STRING&gt;, ARRAY&lt;LONG&gt;, or any other Druid type. The matchValueType determines how Druid interprets the matchValue to assist in converting to the type of the matched column. | 
**MatchValue** | **interface{}** | Value to match, must not be null. | 

## Methods

### NewTelemetryDruidEqualityFilter

`func NewTelemetryDruidEqualityFilter(type_ string, column string, matchValueType string, matchValue interface{}, ) *TelemetryDruidEqualityFilter`

NewTelemetryDruidEqualityFilter instantiates a new TelemetryDruidEqualityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidEqualityFilterWithDefaults

`func NewTelemetryDruidEqualityFilterWithDefaults() *TelemetryDruidEqualityFilter`

NewTelemetryDruidEqualityFilterWithDefaults instantiates a new TelemetryDruidEqualityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidEqualityFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidEqualityFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidEqualityFilter) SetType(v string)`

SetType sets Type field to given value.


### GetColumn

`func (o *TelemetryDruidEqualityFilter) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *TelemetryDruidEqualityFilter) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *TelemetryDruidEqualityFilter) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetMatchValueType

`func (o *TelemetryDruidEqualityFilter) GetMatchValueType() string`

GetMatchValueType returns the MatchValueType field if non-nil, zero value otherwise.

### GetMatchValueTypeOk

`func (o *TelemetryDruidEqualityFilter) GetMatchValueTypeOk() (*string, bool)`

GetMatchValueTypeOk returns a tuple with the MatchValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValueType

`func (o *TelemetryDruidEqualityFilter) SetMatchValueType(v string)`

SetMatchValueType sets MatchValueType field to given value.


### GetMatchValue

`func (o *TelemetryDruidEqualityFilter) GetMatchValue() interface{}`

GetMatchValue returns the MatchValue field if non-nil, zero value otherwise.

### GetMatchValueOk

`func (o *TelemetryDruidEqualityFilter) GetMatchValueOk() (*interface{}, bool)`

GetMatchValueOk returns a tuple with the MatchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValue

`func (o *TelemetryDruidEqualityFilter) SetMatchValue(v interface{})`

SetMatchValue sets MatchValue field to given value.


### SetMatchValueNil

`func (o *TelemetryDruidEqualityFilter) SetMatchValueNil(b bool)`

 SetMatchValueNil sets the value for MatchValue to be an explicit nil

### UnsetMatchValue
`func (o *TelemetryDruidEqualityFilter) UnsetMatchValue()`

UnsetMatchValue ensures that no value is present for MatchValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


