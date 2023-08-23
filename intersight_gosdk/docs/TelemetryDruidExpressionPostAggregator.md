# TelemetryDruidExpressionPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Fields** | Pointer to [**[]TelemetryDruidPostAggregator**](TelemetryDruidPostAggregator.md) | Fields processed by post aggregator | [optional] 
**Name** | Pointer to **string** | Output name for the return value of the expression. | [optional] 
**Expression** | Pointer to **string** | The Druid expression. | [optional] 
**Ordering** | Pointer to **string** | Expression post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last. | [optional] 

## Methods

### NewTelemetryDruidExpressionPostAggregator

`func NewTelemetryDruidExpressionPostAggregator(type_ string, ) *TelemetryDruidExpressionPostAggregator`

NewTelemetryDruidExpressionPostAggregator instantiates a new TelemetryDruidExpressionPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExpressionPostAggregatorWithDefaults

`func NewTelemetryDruidExpressionPostAggregatorWithDefaults() *TelemetryDruidExpressionPostAggregator`

NewTelemetryDruidExpressionPostAggregatorWithDefaults instantiates a new TelemetryDruidExpressionPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExpressionPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExpressionPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExpressionPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *TelemetryDruidExpressionPostAggregator) GetFields() []TelemetryDruidPostAggregator`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidExpressionPostAggregator) GetFieldsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidExpressionPostAggregator) SetFields(v []TelemetryDruidPostAggregator)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidExpressionPostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *TelemetryDruidExpressionPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidExpressionPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidExpressionPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidExpressionPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *TelemetryDruidExpressionPostAggregator) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidExpressionPostAggregator) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidExpressionPostAggregator) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *TelemetryDruidExpressionPostAggregator) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidExpressionPostAggregator) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidExpressionPostAggregator) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidExpressionPostAggregator) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidExpressionPostAggregator) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


