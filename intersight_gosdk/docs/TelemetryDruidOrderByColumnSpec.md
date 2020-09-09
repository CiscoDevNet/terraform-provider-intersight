# TelemetryDruidOrderByColumnSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | Any dimension or metric name. | [optional] 
**Direction** | Pointer to **string** | null | [optional] 
**DimensionOrder** | Pointer to **string** | null | [optional] 

## Methods

### NewTelemetryDruidOrderByColumnSpec

`func NewTelemetryDruidOrderByColumnSpec() *TelemetryDruidOrderByColumnSpec`

NewTelemetryDruidOrderByColumnSpec instantiates a new TelemetryDruidOrderByColumnSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidOrderByColumnSpecWithDefaults

`func NewTelemetryDruidOrderByColumnSpecWithDefaults() *TelemetryDruidOrderByColumnSpec`

NewTelemetryDruidOrderByColumnSpecWithDefaults instantiates a new TelemetryDruidOrderByColumnSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *TelemetryDruidOrderByColumnSpec) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidOrderByColumnSpec) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *TelemetryDruidOrderByColumnSpec) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetDirection

`func (o *TelemetryDruidOrderByColumnSpec) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TelemetryDruidOrderByColumnSpec) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TelemetryDruidOrderByColumnSpec) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *TelemetryDruidOrderByColumnSpec) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDimensionOrder

`func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOrder() string`

GetDimensionOrder returns the DimensionOrder field if non-nil, zero value otherwise.

### GetDimensionOrderOk

`func (o *TelemetryDruidOrderByColumnSpec) GetDimensionOrderOk() (*string, bool)`

GetDimensionOrderOk returns a tuple with the DimensionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionOrder

`func (o *TelemetryDruidOrderByColumnSpec) SetDimensionOrder(v string)`

SetDimensionOrder sets DimensionOrder field to given value.

### HasDimensionOrder

`func (o *TelemetryDruidOrderByColumnSpec) HasDimensionOrder() bool`

HasDimensionOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


