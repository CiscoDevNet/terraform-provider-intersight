# TelemetryDruidThetaSketchOperationsPostAggregatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Output name for the post-aggregator. | [optional] 
**Func** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **[]string** | array of fieldAccess type post aggregators to access the thetaSketch aggregators or thetaSketchSetOp type post aggregators to allow arbitrary combination of set operations. | [optional] 
**Size** | Pointer to **int32** | must be max of size from sketches in fields input. | [optional] [default to 16384]

## Methods

### NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOf

`func NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOf() *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf`

NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOf instantiates a new TelemetryDruidThetaSketchOperationsPostAggregatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOfWithDefaults

`func NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOfWithDefaults() *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf`

NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidThetaSketchOperationsPostAggregatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFunc

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFunc() string`

GetFunc returns the Func field if non-nil, zero value otherwise.

### GetFuncOk

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFuncOk() (*string, bool)`

GetFuncOk returns a tuple with the Func field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunc

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetFunc(v string)`

SetFunc sets Func field to given value.

### HasFunc

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasFunc() bool`

HasFunc returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSize

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


