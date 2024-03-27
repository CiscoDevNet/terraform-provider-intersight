# TelemetryDruidExtractionFunctionBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Size** | Pointer to **int32** | The size of the buckets (optional, default 1). | [optional] 
**Offset** | Pointer to **int32** | The offset for the buckets (optional, default 0). | [optional] 

## Methods

### NewTelemetryDruidExtractionFunctionBucket

`func NewTelemetryDruidExtractionFunctionBucket(type_ string, ) *TelemetryDruidExtractionFunctionBucket`

NewTelemetryDruidExtractionFunctionBucket instantiates a new TelemetryDruidExtractionFunctionBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExtractionFunctionBucketWithDefaults

`func NewTelemetryDruidExtractionFunctionBucketWithDefaults() *TelemetryDruidExtractionFunctionBucket`

NewTelemetryDruidExtractionFunctionBucketWithDefaults instantiates a new TelemetryDruidExtractionFunctionBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExtractionFunctionBucket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExtractionFunctionBucket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExtractionFunctionBucket) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *TelemetryDruidExtractionFunctionBucket) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidExtractionFunctionBucket) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidExtractionFunctionBucket) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidExtractionFunctionBucket) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetOffset

`func (o *TelemetryDruidExtractionFunctionBucket) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TelemetryDruidExtractionFunctionBucket) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TelemetryDruidExtractionFunctionBucket) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TelemetryDruidExtractionFunctionBucket) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


