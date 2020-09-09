# TelemetryDruidThetaSketchAggregatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A String for the output (result) name of the calculation. | 
**FieldName** | **string** | A String for the name of the aggregator used at ingestion time. | 
**Size** | Pointer to **int32** | Must be a power of 2. Internally, size refers to the maximum number of entries sketch object will retain. Higher size means higher accuracy but more space to store sketches. Note that after you index with a particular size, druid will persist sketch in segments and you will use size greater or equal to that at query time. In general, We recommend just sticking to default size. | [optional] [default to 16384]

## Methods

### NewTelemetryDruidThetaSketchAggregatorAllOf

`func NewTelemetryDruidThetaSketchAggregatorAllOf(name string, fieldName string, ) *TelemetryDruidThetaSketchAggregatorAllOf`

NewTelemetryDruidThetaSketchAggregatorAllOf instantiates a new TelemetryDruidThetaSketchAggregatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidThetaSketchAggregatorAllOfWithDefaults

`func NewTelemetryDruidThetaSketchAggregatorAllOfWithDefaults() *TelemetryDruidThetaSketchAggregatorAllOf`

NewTelemetryDruidThetaSketchAggregatorAllOfWithDefaults instantiates a new TelemetryDruidThetaSketchAggregatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetSize

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidThetaSketchAggregatorAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


