# TelemetryDruidFieldAccessorPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**FieldName** | Pointer to **string** | Name of the metric column. | [optional] 

## Methods

### NewTelemetryDruidFieldAccessorPostAggregator

`func NewTelemetryDruidFieldAccessorPostAggregator(type_ string, ) *TelemetryDruidFieldAccessorPostAggregator`

NewTelemetryDruidFieldAccessorPostAggregator instantiates a new TelemetryDruidFieldAccessorPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidFieldAccessorPostAggregatorWithDefaults

`func NewTelemetryDruidFieldAccessorPostAggregatorWithDefaults() *TelemetryDruidFieldAccessorPostAggregator`

NewTelemetryDruidFieldAccessorPostAggregatorWithDefaults instantiates a new TelemetryDruidFieldAccessorPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidFieldAccessorPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidFieldAccessorPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidFieldAccessorPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFieldName

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidFieldAccessorPostAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidFieldAccessorPostAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *TelemetryDruidFieldAccessorPostAggregator) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


