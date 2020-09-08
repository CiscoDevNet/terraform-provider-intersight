# TelemetryDruidGreatestLeastPostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**Fields** | Pointer to **string** | the fields that are used to compute the greatest or least value. | [optional] 

## Methods

### NewTelemetryDruidGreatestLeastPostAggregator

`func NewTelemetryDruidGreatestLeastPostAggregator(type_ string, ) *TelemetryDruidGreatestLeastPostAggregator`

NewTelemetryDruidGreatestLeastPostAggregator instantiates a new TelemetryDruidGreatestLeastPostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidGreatestLeastPostAggregatorWithDefaults

`func NewTelemetryDruidGreatestLeastPostAggregatorWithDefaults() *TelemetryDruidGreatestLeastPostAggregator`

NewTelemetryDruidGreatestLeastPostAggregatorWithDefaults instantiates a new TelemetryDruidGreatestLeastPostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidGreatestLeastPostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidGreatestLeastPostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidGreatestLeastPostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidGreatestLeastPostAggregator) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidGreatestLeastPostAggregator) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidGreatestLeastPostAggregator) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


