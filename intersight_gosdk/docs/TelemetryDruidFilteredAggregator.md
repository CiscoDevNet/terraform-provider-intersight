# TelemetryDruidFilteredAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The aggregator type. | 
**Filter** | [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 
**Aggregator** | [**TelemetryDruidAggregator**](TelemetryDruidAggregator.md) |  | 

## Methods

### NewTelemetryDruidFilteredAggregator

`func NewTelemetryDruidFilteredAggregator(type_ string, filter TelemetryDruidFilter, aggregator TelemetryDruidAggregator, ) *TelemetryDruidFilteredAggregator`

NewTelemetryDruidFilteredAggregator instantiates a new TelemetryDruidFilteredAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidFilteredAggregatorWithDefaults

`func NewTelemetryDruidFilteredAggregatorWithDefaults() *TelemetryDruidFilteredAggregator`

NewTelemetryDruidFilteredAggregatorWithDefaults instantiates a new TelemetryDruidFilteredAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidFilteredAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidFilteredAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidFilteredAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetFilter

`func (o *TelemetryDruidFilteredAggregator) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidFilteredAggregator) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidFilteredAggregator) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.


### GetAggregator

`func (o *TelemetryDruidFilteredAggregator) GetAggregator() TelemetryDruidAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *TelemetryDruidFilteredAggregator) GetAggregatorOk() (*TelemetryDruidAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *TelemetryDruidFilteredAggregator) SetAggregator(v TelemetryDruidAggregator)`

SetAggregator sets Aggregator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


