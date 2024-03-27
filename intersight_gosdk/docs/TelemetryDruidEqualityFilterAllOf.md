# TelemetryDruidEqualityFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Column** | **string** | Input column or virtual column name to filter. | 
**MatchValueType** | **string** | String specifying the type of value to match. For example STRING, LONG, DOUBLE, FLOAT, ARRAY&lt;STRING&gt;, ARRAY&lt;LONG&gt;, or any other Druid type. The matchValueType determines how Druid interprets the matchValue to assist in converting to the type of the matched column. | 
**MatchValue** | **interface{}** | Value to match, must not be null. | 

## Methods

### NewTelemetryDruidEqualityFilterAllOf

`func NewTelemetryDruidEqualityFilterAllOf(type_ string, column string, matchValueType string, matchValue interface{}, ) *TelemetryDruidEqualityFilterAllOf`

NewTelemetryDruidEqualityFilterAllOf instantiates a new TelemetryDruidEqualityFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidEqualityFilterAllOfWithDefaults

`func NewTelemetryDruidEqualityFilterAllOfWithDefaults() *TelemetryDruidEqualityFilterAllOf`

NewTelemetryDruidEqualityFilterAllOfWithDefaults instantiates a new TelemetryDruidEqualityFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidEqualityFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidEqualityFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidEqualityFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetColumn

`func (o *TelemetryDruidEqualityFilterAllOf) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *TelemetryDruidEqualityFilterAllOf) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *TelemetryDruidEqualityFilterAllOf) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetMatchValueType

`func (o *TelemetryDruidEqualityFilterAllOf) GetMatchValueType() string`

GetMatchValueType returns the MatchValueType field if non-nil, zero value otherwise.

### GetMatchValueTypeOk

`func (o *TelemetryDruidEqualityFilterAllOf) GetMatchValueTypeOk() (*string, bool)`

GetMatchValueTypeOk returns a tuple with the MatchValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValueType

`func (o *TelemetryDruidEqualityFilterAllOf) SetMatchValueType(v string)`

SetMatchValueType sets MatchValueType field to given value.


### GetMatchValue

`func (o *TelemetryDruidEqualityFilterAllOf) GetMatchValue() interface{}`

GetMatchValue returns the MatchValue field if non-nil, zero value otherwise.

### GetMatchValueOk

`func (o *TelemetryDruidEqualityFilterAllOf) GetMatchValueOk() (*interface{}, bool)`

GetMatchValueOk returns a tuple with the MatchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValue

`func (o *TelemetryDruidEqualityFilterAllOf) SetMatchValue(v interface{})`

SetMatchValue sets MatchValue field to given value.


### SetMatchValueNil

`func (o *TelemetryDruidEqualityFilterAllOf) SetMatchValueNil(b bool)`

 SetMatchValueNil sets the value for MatchValue to be an explicit nil

### UnsetMatchValue
`func (o *TelemetryDruidEqualityFilterAllOf) UnsetMatchValue()`

UnsetMatchValue ensures that no value is present for MatchValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


