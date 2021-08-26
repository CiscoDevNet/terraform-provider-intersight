# TelemetryDruidHavingFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The having filter type. | 
**Filter** | [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Aggregation** | **string** | aggregate metric | 
**Value** | **float64** | null | 
**Dimension** | **string** | dimension | 

## Methods

### NewTelemetryDruidHavingFilter

`func NewTelemetryDruidHavingFilter(type_ string, filter TelemetryDruidFilter, aggregation string, value float64, dimension string, ) *TelemetryDruidHavingFilter`

NewTelemetryDruidHavingFilter instantiates a new TelemetryDruidHavingFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidHavingFilterWithDefaults

`func NewTelemetryDruidHavingFilterWithDefaults() *TelemetryDruidHavingFilter`

NewTelemetryDruidHavingFilterWithDefaults instantiates a new TelemetryDruidHavingFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidHavingFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidHavingFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidHavingFilter) SetType(v string)`

SetType sets Type field to given value.


### GetFilter

`func (o *TelemetryDruidHavingFilter) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidHavingFilter) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidHavingFilter) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.


### GetAggregation

`func (o *TelemetryDruidHavingFilter) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TelemetryDruidHavingFilter) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TelemetryDruidHavingFilter) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetValue

`func (o *TelemetryDruidHavingFilter) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidHavingFilter) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidHavingFilter) SetValue(v float64)`

SetValue sets Value field to given value.


### GetDimension

`func (o *TelemetryDruidHavingFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidHavingFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidHavingFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


