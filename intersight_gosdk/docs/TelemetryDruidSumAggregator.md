# TelemetryDruidSumAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The aggregator type. | 
**Name** | **string** | Output name for the summed value. | 
**FieldName** | **string** | Name of the metric column to sum over. | 

## Methods

### NewTelemetryDruidSumAggregator

`func NewTelemetryDruidSumAggregator(type_ string, name string, fieldName string, ) *TelemetryDruidSumAggregator`

NewTelemetryDruidSumAggregator instantiates a new TelemetryDruidSumAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSumAggregatorWithDefaults

`func NewTelemetryDruidSumAggregatorWithDefaults() *TelemetryDruidSumAggregator`

NewTelemetryDruidSumAggregatorWithDefaults instantiates a new TelemetryDruidSumAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidSumAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidSumAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidSumAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidSumAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidSumAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidSumAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidSumAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidSumAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidSumAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


