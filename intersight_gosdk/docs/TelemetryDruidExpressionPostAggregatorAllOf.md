# TelemetryDruidExpressionPostAggregatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Output name for the return value of the expression. | [optional] 
**Expression** | Pointer to **string** | The Druid expression. | [optional] 
**Ordering** | Pointer to **string** | Expression post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last. | [optional] 

## Methods

### NewTelemetryDruidExpressionPostAggregatorAllOf

`func NewTelemetryDruidExpressionPostAggregatorAllOf() *TelemetryDruidExpressionPostAggregatorAllOf`

NewTelemetryDruidExpressionPostAggregatorAllOf instantiates a new TelemetryDruidExpressionPostAggregatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExpressionPostAggregatorAllOfWithDefaults

`func NewTelemetryDruidExpressionPostAggregatorAllOfWithDefaults() *TelemetryDruidExpressionPostAggregatorAllOf`

NewTelemetryDruidExpressionPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidExpressionPostAggregatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


