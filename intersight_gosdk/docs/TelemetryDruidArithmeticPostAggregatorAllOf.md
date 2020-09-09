# TelemetryDruidArithmeticPostAggregatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Output name for the minimum/maximum timestamp value. | [optional] 
**Fn** | Pointer to **string** | null | [optional] 
**Fields** | Pointer to **[]string** | null | [optional] 
**Ordering** | Pointer to **string** | Arithmetic post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. This can be useful for topN queries for instance. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last. | [optional] 

## Methods

### NewTelemetryDruidArithmeticPostAggregatorAllOf

`func NewTelemetryDruidArithmeticPostAggregatorAllOf() *TelemetryDruidArithmeticPostAggregatorAllOf`

NewTelemetryDruidArithmeticPostAggregatorAllOf instantiates a new TelemetryDruidArithmeticPostAggregatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidArithmeticPostAggregatorAllOfWithDefaults

`func NewTelemetryDruidArithmeticPostAggregatorAllOfWithDefaults() *TelemetryDruidArithmeticPostAggregatorAllOf`

NewTelemetryDruidArithmeticPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidArithmeticPostAggregatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFn

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFn() string`

GetFn returns the Fn field if non-nil, zero value otherwise.

### GetFnOk

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFnOk() (*string, bool)`

GetFnOk returns a tuple with the Fn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFn

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetFn(v string)`

SetFn sets Fn field to given value.

### HasFn

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasFn() bool`

HasFn returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


