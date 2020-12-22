# TelemetryDruidDefaultLimitSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | How many rows to return. If not specified, all rows will be returned. | 
**Columns** | [**[]TelemetryDruidOrderByColumnSpec**](TelemetryDruidOrderByColumnSpec.md) | null | 

## Methods

### NewTelemetryDruidDefaultLimitSpecAllOf

`func NewTelemetryDruidDefaultLimitSpecAllOf(limit int32, columns []TelemetryDruidOrderByColumnSpec, ) *TelemetryDruidDefaultLimitSpecAllOf`

NewTelemetryDruidDefaultLimitSpecAllOf instantiates a new TelemetryDruidDefaultLimitSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDefaultLimitSpecAllOfWithDefaults

`func NewTelemetryDruidDefaultLimitSpecAllOfWithDefaults() *TelemetryDruidDefaultLimitSpecAllOf`

NewTelemetryDruidDefaultLimitSpecAllOfWithDefaults instantiates a new TelemetryDruidDefaultLimitSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TelemetryDruidDefaultLimitSpecAllOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidDefaultLimitSpecAllOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidDefaultLimitSpecAllOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetColumns

`func (o *TelemetryDruidDefaultLimitSpecAllOf) GetColumns() []TelemetryDruidOrderByColumnSpec`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidDefaultLimitSpecAllOf) GetColumnsOk() (*[]TelemetryDruidOrderByColumnSpec, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidDefaultLimitSpecAllOf) SetColumns(v []TelemetryDruidOrderByColumnSpec)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


