# TelemetryDruidConstantPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**Value** | Pointer to **float64** | The numerical value. | [optional] 

## Methods

### NewTelemetryDruidConstantPostAggregator

`func NewTelemetryDruidConstantPostAggregator(type_ string, ) *TelemetryDruidConstantPostAggregator`

NewTelemetryDruidConstantPostAggregator instantiates a new TelemetryDruidConstantPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidConstantPostAggregatorWithDefaults

`func NewTelemetryDruidConstantPostAggregatorWithDefaults() *TelemetryDruidConstantPostAggregator`

NewTelemetryDruidConstantPostAggregatorWithDefaults instantiates a new TelemetryDruidConstantPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidConstantPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidConstantPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidConstantPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidConstantPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidConstantPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidConstantPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidConstantPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *TelemetryDruidConstantPostAggregator) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelemetryDruidConstantPostAggregator) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelemetryDruidConstantPostAggregator) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TelemetryDruidConstantPostAggregator) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


