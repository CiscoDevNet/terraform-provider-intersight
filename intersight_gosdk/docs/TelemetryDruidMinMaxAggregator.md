# TelemetryDruidMinMaxAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The aggregator type. | 
**Name** | **string** | Output name for the min/max value. | 
**FieldName** | **string** | Name of the metric column. | 

## Methods

### NewTelemetryDruidMinMaxAggregator

`func NewTelemetryDruidMinMaxAggregator(type_ string, name string, fieldName string, ) *TelemetryDruidMinMaxAggregator`

NewTelemetryDruidMinMaxAggregator instantiates a new TelemetryDruidMinMaxAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidMinMaxAggregatorWithDefaults

`func NewTelemetryDruidMinMaxAggregatorWithDefaults() *TelemetryDruidMinMaxAggregator`

NewTelemetryDruidMinMaxAggregatorWithDefaults instantiates a new TelemetryDruidMinMaxAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidMinMaxAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidMinMaxAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidMinMaxAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidMinMaxAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidMinMaxAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidMinMaxAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidMinMaxAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidMinMaxAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidMinMaxAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


