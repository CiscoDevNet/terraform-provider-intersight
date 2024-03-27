# TelemetryDruidIntervalFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dimension** | **string** | Input column or virtual column name to filter. | 
**Intervals** | **[]string** | A JSON array containing ISO-8601 interval strings that defines the time ranges to filter on. | 
**ExtractionFn** | Pointer to [**TelemetryDruidExtractionFunction**](TelemetryDruidExtractionFunction.md) |  | [optional] 

## Methods

### NewTelemetryDruidIntervalFilterAllOf

`func NewTelemetryDruidIntervalFilterAllOf(type_ string, dimension string, intervals []string, ) *TelemetryDruidIntervalFilterAllOf`

NewTelemetryDruidIntervalFilterAllOf instantiates a new TelemetryDruidIntervalFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidIntervalFilterAllOfWithDefaults

`func NewTelemetryDruidIntervalFilterAllOfWithDefaults() *TelemetryDruidIntervalFilterAllOf`

NewTelemetryDruidIntervalFilterAllOfWithDefaults instantiates a new TelemetryDruidIntervalFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidIntervalFilterAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidIntervalFilterAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidIntervalFilterAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetDimension

`func (o *TelemetryDruidIntervalFilterAllOf) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidIntervalFilterAllOf) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidIntervalFilterAllOf) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetIntervals

`func (o *TelemetryDruidIntervalFilterAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidIntervalFilterAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidIntervalFilterAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetExtractionFn

`func (o *TelemetryDruidIntervalFilterAllOf) GetExtractionFn() TelemetryDruidExtractionFunction`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidIntervalFilterAllOf) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidIntervalFilterAllOf) SetExtractionFn(v TelemetryDruidExtractionFunction)`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidIntervalFilterAllOf) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


