# TelemetryDruidBasePostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Fields** | Pointer to [**[]TelemetryDruidPostAggregator**](TelemetryDruidPostAggregator.md) | Fields processed by post aggregator | [optional] 

## Methods

### NewTelemetryDruidBasePostAggregator

`func NewTelemetryDruidBasePostAggregator(type_ string, ) *TelemetryDruidBasePostAggregator`

NewTelemetryDruidBasePostAggregator instantiates a new TelemetryDruidBasePostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidBasePostAggregatorWithDefaults

`func NewTelemetryDruidBasePostAggregatorWithDefaults() *TelemetryDruidBasePostAggregator`

NewTelemetryDruidBasePostAggregatorWithDefaults instantiates a new TelemetryDruidBasePostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidBasePostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidBasePostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidBasePostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *TelemetryDruidBasePostAggregator) GetFields() []TelemetryDruidPostAggregator`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidBasePostAggregator) GetFieldsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidBasePostAggregator) SetFields(v []TelemetryDruidPostAggregator)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidBasePostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


