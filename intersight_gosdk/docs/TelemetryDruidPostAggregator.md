# TelemetryDruidPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**Fn** | Pointer to **string** | null | [optional] 
**Fields** | Pointer to **[]string** | array of fieldAccess type post aggregators to access the thetaSketch aggregators or thetaSketchSetOp type post aggregators to allow arbitrary combination of set operations. | [optional] 
**Ordering** | Pointer to **string** | Arithmetic post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. This can be useful for topN queries for instance. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last. | [optional] 
**FieldName** | Pointer to **string** | The name field value of the hyperUnique aggregator. | [optional] 
**Value** | Pointer to **float64** | The numerical value. | [optional] 
**Field** | Pointer to **string** | Post aggregator of type fieldAccess that refers to a thetaSketch aggregator or that of type thetaSketchSetOp. | [optional] 
**Func** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | must be max of size from sketches in fields input. | [optional] [default to 16384]

## Methods

### NewTelemetryDruidPostAggregator

`func NewTelemetryDruidPostAggregator(type_ string, ) *TelemetryDruidPostAggregator`

NewTelemetryDruidPostAggregator instantiates a new TelemetryDruidPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidPostAggregatorWithDefaults

`func NewTelemetryDruidPostAggregatorWithDefaults() *TelemetryDruidPostAggregator`

NewTelemetryDruidPostAggregatorWithDefaults instantiates a new TelemetryDruidPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFn

`func (o *TelemetryDruidPostAggregator) GetFn() string`

GetFn returns the Fn field if non-nil, zero value otherwise.

### GetFnOk

`func (o *TelemetryDruidPostAggregator) GetFnOk() (*string, bool)`

GetFnOk returns a tuple with the Fn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFn

`func (o *TelemetryDruidPostAggregator) SetFn(v string)`

SetFn sets Fn field to given value.

### HasFn

`func (o *TelemetryDruidPostAggregator) HasFn() bool`

HasFn returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidPostAggregator) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidPostAggregator) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidPostAggregator) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidPostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOrdering

`func (o *TelemetryDruidPostAggregator) GetOrdering() string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TelemetryDruidPostAggregator) GetOrderingOk() (*string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TelemetryDruidPostAggregator) SetOrdering(v string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TelemetryDruidPostAggregator) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetFieldName

`func (o *TelemetryDruidPostAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidPostAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidPostAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *TelemetryDruidPostAggregator) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetValue

`func (o *TelemetryDruidPostAggregator) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidPostAggregator) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidPostAggregator) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TelemetryDruidPostAggregator) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetField

`func (o *TelemetryDruidPostAggregator) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TelemetryDruidPostAggregator) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TelemetryDruidPostAggregator) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TelemetryDruidPostAggregator) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFunc

`func (o *TelemetryDruidPostAggregator) GetFunc() string`

GetFunc returns the Func field if non-nil, zero value otherwise.

### GetFuncOk

`func (o *TelemetryDruidPostAggregator) GetFuncOk() (*string, bool)`

GetFuncOk returns a tuple with the Func field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunc

`func (o *TelemetryDruidPostAggregator) SetFunc(v string)`

SetFunc sets Func field to given value.

### HasFunc

`func (o *TelemetryDruidPostAggregator) HasFunc() bool`

HasFunc returns a boolean if a field has been set.

### GetSize

`func (o *TelemetryDruidPostAggregator) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidPostAggregator) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidPostAggregator) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidPostAggregator) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


