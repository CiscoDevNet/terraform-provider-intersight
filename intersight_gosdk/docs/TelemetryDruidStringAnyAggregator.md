# TelemetryDruidStringAnyAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The aggregator type. | 
**Name** | **string** | Output name for the &#39;any&#39; value. | 
**FieldName** | **string** | Name of the metric column. | 
**MaxStringBytes** | Pointer to **int32** | null | [optional] [default to 1024]

## Methods

### NewTelemetryDruidStringAnyAggregator

`func NewTelemetryDruidStringAnyAggregator(type_ string, name string, fieldName string, ) *TelemetryDruidStringAnyAggregator`

NewTelemetryDruidStringAnyAggregator instantiates a new TelemetryDruidStringAnyAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidStringAnyAggregatorWithDefaults

`func NewTelemetryDruidStringAnyAggregatorWithDefaults() *TelemetryDruidStringAnyAggregator`

NewTelemetryDruidStringAnyAggregatorWithDefaults instantiates a new TelemetryDruidStringAnyAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidStringAnyAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidStringAnyAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidStringAnyAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidStringAnyAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidStringAnyAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidStringAnyAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidStringAnyAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidStringAnyAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidStringAnyAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetMaxStringBytes

`func (o *TelemetryDruidStringAnyAggregator) GetMaxStringBytes() int32`

GetMaxStringBytes returns the MaxStringBytes field if non-nil, zero value otherwise.

### GetMaxStringBytesOk

`func (o *TelemetryDruidStringAnyAggregator) GetMaxStringBytesOk() (*int32, bool)`

GetMaxStringBytesOk returns a tuple with the MaxStringBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringBytes

`func (o *TelemetryDruidStringAnyAggregator) SetMaxStringBytes(v int32)`

SetMaxStringBytes sets MaxStringBytes field to given value.

### HasMaxStringBytes

`func (o *TelemetryDruidStringAnyAggregator) HasMaxStringBytes() bool`

HasMaxStringBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


