# KvmSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.Session"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.Session"]
**OneTimePassword** | Pointer to **string** | Temporary one-time password for vKVM access. | [optional] 
**SsoSupported** | Pointer to **bool** | Indicates if vKVM SSO is supported on the server. | [optional] [readonly] 
**Username** | Pointer to **string** | Username used for vKVM access. | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 
**Tunnel** | Pointer to [**KvmTunnelRelationship**](kvm.Tunnel.Relationship.md) |  | [optional] 

## Methods

### NewKvmSessionAllOf

`func NewKvmSessionAllOf(classId string, objectType string, ) *KvmSessionAllOf`

NewKvmSessionAllOf instantiates a new KvmSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmSessionAllOfWithDefaults

`func NewKvmSessionAllOfWithDefaults() *KvmSessionAllOf`

NewKvmSessionAllOfWithDefaults instantiates a new KvmSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmSessionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmSessionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmSessionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmSessionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmSessionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmSessionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOneTimePassword

`func (o *KvmSessionAllOf) GetOneTimePassword() string`

GetOneTimePassword returns the OneTimePassword field if non-nil, zero value otherwise.

### GetOneTimePasswordOk

`func (o *KvmSessionAllOf) GetOneTimePasswordOk() (*string, bool)`

GetOneTimePasswordOk returns a tuple with the OneTimePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePassword

`func (o *KvmSessionAllOf) SetOneTimePassword(v string)`

SetOneTimePassword sets OneTimePassword field to given value.

### HasOneTimePassword

`func (o *KvmSessionAllOf) HasOneTimePassword() bool`

HasOneTimePassword returns a boolean if a field has been set.

### GetSsoSupported

`func (o *KvmSessionAllOf) GetSsoSupported() bool`

GetSsoSupported returns the SsoSupported field if non-nil, zero value otherwise.

### GetSsoSupportedOk

`func (o *KvmSessionAllOf) GetSsoSupportedOk() (*bool, bool)`

GetSsoSupportedOk returns a tuple with the SsoSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSupported

`func (o *KvmSessionAllOf) SetSsoSupported(v bool)`

SetSsoSupported sets SsoSupported field to given value.

### HasSsoSupported

`func (o *KvmSessionAllOf) HasSsoSupported() bool`

HasSsoSupported returns a boolean if a field has been set.

### GetUsername

`func (o *KvmSessionAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KvmSessionAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KvmSessionAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KvmSessionAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDevice

`func (o *KvmSessionAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmSessionAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmSessionAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmSessionAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmSessionAllOf) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmSessionAllOf) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmSessionAllOf) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmSessionAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetTunnel

`func (o *KvmSessionAllOf) GetTunnel() KvmTunnelRelationship`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *KvmSessionAllOf) GetTunnelOk() (*KvmTunnelRelationship, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *KvmSessionAllOf) SetTunnel(v KvmTunnelRelationship)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *KvmSessionAllOf) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


