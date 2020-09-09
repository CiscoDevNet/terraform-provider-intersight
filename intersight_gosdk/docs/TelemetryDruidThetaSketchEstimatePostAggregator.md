# TelemetryDruidThetaSketchEstimatePostAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The post-aggregator type. | 
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**Field** | Pointer to **string** | Post aggregator of type fieldAccess that refers to a thetaSketch aggregator or that of type thetaSketchSetOp. | [optional] 

## Methods

### NewTelemetryDruidThetaSketchEstimatePostAggregator

`func NewTelemetryDruidThetaSketchEstimatePostAggregator(type_ string, ) *TelemetryDruidThetaSketchEstimatePostAggregator`

NewTelemetryDruidThetaSketchEstimatePostAggregator instantiates a new TelemetryDruidThetaSketchEstimatePostAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidThetaSketchEstimatePostAggregatorWithDefaults

`func NewTelemetryDruidThetaSketchEstimatePostAggregatorWithDefaults() *TelemetryDruidThetaSketchEstimatePostAggregator`

NewTelemetryDruidThetaSketchEstimatePostAggregatorWithDefaults instantiates a new TelemetryDruidThetaSketchEstimatePostAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetField

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TelemetryDruidThetaSketchEstimatePostAggregator) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


