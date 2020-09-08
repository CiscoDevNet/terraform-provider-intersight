# TelemetryDruidArithmeticPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the minimum/maximum timestamp value. | [optional] 
**Fn** | Pointer to **string** | null | [optional] 
**Fields** | Pointer to **[]string** | null | [optional] 
**Ordering** | Pointer to **string** | Arithmetic post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. This can be useful for topN queries for instance. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last. | [optional] 

## Methods

### NewTelemetryDruidArithmeticPostAggregator

`func NewTelemetryDruidArithmeticPostAggregator(type_ string, ) *TelemetryDruidArithmeticPostAggregator`

NewTelemetryDruidArithmeticPostAggregator instantiates a new TelemetryDruidArithmeticPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidArithmeticPostAggregatorWithDefaults

`func NewTelemetryDruidArithmeticPostAggregatorWithDefaults() *TelemetryDruidArithmeticPostAggregator`

NewTelemetryDruidArithmeticPostAggregatorWithDefaults instantiates a new TelemetryDruidArithmeticPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidArithmeticPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidArithmeticPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidArithmeticPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidArithmeticPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidArithmeticPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidArithmeticPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidArithmeticPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFn

`func (o *TelemetryDruidArithmeticPostAggregator) GetFn() string`

GetFn returns the Fn field if non-nil, zero value otherwise.

### GetFnOk

`func (o *TelemetryDruidArithmeticPostAggregator) GetFnOk() (*string, bool)`

GetFnOk returns a tuple with the Fn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFn

`func (o *TelemetryDruidArithmeticPostAggregator) SetFn(v string)`

SetFn sets Fn field to given value.

### HasFn

`func (o *TelemetryDruidArithmeticPostAggregator) HasFn() bool`

HasFn returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidArithmeticPostAggregator) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidArithmeticPostAggregator) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidArithmeticPostAggregator) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidArithmeticPostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidArithmeticPostAggregator) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidArithmeticPostAggregator) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidArithmeticPostAggregator) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidArithmeticPostAggregator) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


