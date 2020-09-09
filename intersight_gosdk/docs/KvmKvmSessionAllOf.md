# KvmKvmSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KvmMuxUrl** | Pointer to **string** | The URL of the KVM MUX to connect to. | [optional] [readonly] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 

## Methods

### NewKvmKvmSessionAllOf

`func NewKvmKvmSessionAllOf() *KvmKvmSessionAllOf`

NewKvmKvmSessionAllOf instantiates a new KvmKvmSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmKvmSessionAllOfWithDefaults

`func NewKvmKvmSessionAllOfWithDefaults() *KvmKvmSessionAllOf`

NewKvmKvmSessionAllOfWithDefaults instantiates a new KvmKvmSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKvmMuxUrl

`func (o *KvmKvmSessionAllOf) GetKvmMuxUrl() string`

GetKvmMuxUrl returns the KvmMuxUrl field if non-nil, zero value otherwise.

### GetKvmMuxUrlOk

`func (o *KvmKvmSessionAllOf) GetKvmMuxUrlOk() (*string, bool)`

GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmMuxUrl

`func (o *KvmKvmSessionAllOf) SetKvmMuxUrl(v string)`

SetKvmMuxUrl sets KvmMuxUrl field to given value.

### HasKvmMuxUrl

`func (o *KvmKvmSessionAllOf) HasKvmMuxUrl() bool`

HasKvmMuxUrl returns a boolean if a field has been set.

### GetDevice

`func (o *KvmKvmSessionAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmKvmSessionAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmKvmSessionAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmKvmSessionAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmKvmSessionAllOf) GetServer() ComputeRackUnitRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmKvmSessionAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmKvmSessionAllOf) SetServer(v ComputeRackUnitRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmKvmSessionAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


