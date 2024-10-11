# TelemetryNestedFieldVirtualColumnAllOfPathParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the path part. Use &#39;field&#39; when accessing a specific field in a nested structure. Use &#39;arrayElement&#39; when accessing a specific integer position of an array. | [optional] 
**Field** | Pointer to **string** | Name of the field in the field type path part. | [optional] 
**Index** | Pointer to **int32** | Index of the array element. | [optional] 

## Methods

### NewTelemetryNestedFieldVirtualColumnAllOfPathParts

`func NewTelemetryNestedFieldVirtualColumnAllOfPathParts() *TelemetryNestedFieldVirtualColumnAllOfPathParts`

NewTelemetryNestedFieldVirtualColumnAllOfPathParts instantiates a new TelemetryNestedFieldVirtualColumnAllOfPathParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryNestedFieldVirtualColumnAllOfPathPartsWithDefaults

`func NewTelemetryNestedFieldVirtualColumnAllOfPathPartsWithDefaults() *TelemetryNestedFieldVirtualColumnAllOfPathParts`

NewTelemetryNestedFieldVirtualColumnAllOfPathPartsWithDefaults instantiates a new TelemetryNestedFieldVirtualColumnAllOfPathParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetField

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) HasField() bool`

HasField returns a boolean if a field has been set.

### GetIndex

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TelemetryNestedFieldVirtualColumnAllOfPathParts) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


