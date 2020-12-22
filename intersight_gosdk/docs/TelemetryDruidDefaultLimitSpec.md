# TelemetryDruidDefaultLimitSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The limit spec type. | 
**Limit** | **int32** | How many rows to return. If not specified, all rows will be returned. | 
**Columns** | [**[]TelemetryDruidOrderByColumnSpec**](TelemetryDruidOrderByColumnSpec.md) | null | 

## Methods

### NewTelemetryDruidDefaultLimitSpec

`func NewTelemetryDruidDefaultLimitSpec(type_ string, limit int32, columns []TelemetryDruidOrderByColumnSpec, ) *TelemetryDruidDefaultLimitSpec`

NewTelemetryDruidDefaultLimitSpec instantiates a new TelemetryDruidDefaultLimitSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDefaultLimitSpecWithDefaults

`func NewTelemetryDruidDefaultLimitSpecWithDefaults() *TelemetryDruidDefaultLimitSpec`

NewTelemetryDruidDefaultLimitSpecWithDefaults instantiates a new TelemetryDruidDefaultLimitSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidDefaultLimitSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidDefaultLimitSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidDefaultLimitSpec) SetType(v string)`

SetType sets Type field to given value.


### GetLimit

`func (o *TelemetryDruidDefaultLimitSpec) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidDefaultLimitSpec) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidDefaultLimitSpec) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetColumns

`func (o *TelemetryDruidDefaultLimitSpec) GetColumns() []TelemetryDruidOrderByColumnSpec`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidDefaultLimitSpec) GetColumnsOk() (*[]TelemetryDruidOrderByColumnSpec, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidDefaultLimitSpec) SetColumns(v []TelemetryDruidOrderByColumnSpec)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


