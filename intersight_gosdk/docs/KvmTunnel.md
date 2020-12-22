# KvmTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.Tunnel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.Tunnel"]
**Var0Session** | Pointer to [**KvmSessionRelationship**](kvm.Session.Relationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewKvmTunnel

`func NewKvmTunnel(classId string, objectType string, ) *KvmTunnel`

NewKvmTunnel instantiates a new KvmTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTunnelWithDefaults

`func NewKvmTunnelWithDefaults() *KvmTunnel`

NewKvmTunnelWithDefaults instantiates a new KvmTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmTunnel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmTunnel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmTunnel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmTunnel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmTunnel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmTunnel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVar0Session

`func (o *KvmTunnel) GetVar0Session() KvmSessionRelationship`

GetVar0Session returns the Var0Session field if non-nil, zero value otherwise.

### GetVar0SessionOk

`func (o *KvmTunnel) GetVar0SessionOk() (*KvmSessionRelationship, bool)`

GetVar0SessionOk returns a tuple with the Var0Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0Session

`func (o *KvmTunnel) SetVar0Session(v KvmSessionRelationship)`

SetVar0Session sets Var0Session field to given value.

### HasVar0Session

`func (o *KvmTunnel) HasVar0Session() bool`

HasVar0Session returns a boolean if a field has been set.

### GetDevice

`func (o *KvmTunnel) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmTunnel) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmTunnel) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmTunnel) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmTunnel) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmTunnel) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmTunnel) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmTunnel) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


