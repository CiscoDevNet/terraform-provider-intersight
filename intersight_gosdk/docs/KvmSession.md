# KvmSession

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

### NewKvmSession

`func NewKvmSession(classId string, objectType string, ) *KvmSession`

NewKvmSession instantiates a new KvmSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmSessionWithDefaults

`func NewKvmSessionWithDefaults() *KvmSession`

NewKvmSessionWithDefaults instantiates a new KvmSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOneTimePassword

`func (o *KvmSession) GetOneTimePassword() string`

GetOneTimePassword returns the OneTimePassword field if non-nil, zero value otherwise.

### GetOneTimePasswordOk

`func (o *KvmSession) GetOneTimePasswordOk() (*string, bool)`

GetOneTimePasswordOk returns a tuple with the OneTimePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePassword

`func (o *KvmSession) SetOneTimePassword(v string)`

SetOneTimePassword sets OneTimePassword field to given value.

### HasOneTimePassword

`func (o *KvmSession) HasOneTimePassword() bool`

HasOneTimePassword returns a boolean if a field has been set.

### GetSsoSupported

`func (o *KvmSession) GetSsoSupported() bool`

GetSsoSupported returns the SsoSupported field if non-nil, zero value otherwise.

### GetSsoSupportedOk

`func (o *KvmSession) GetSsoSupportedOk() (*bool, bool)`

GetSsoSupportedOk returns a tuple with the SsoSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSupported

`func (o *KvmSession) SetSsoSupported(v bool)`

SetSsoSupported sets SsoSupported field to given value.

### HasSsoSupported

`func (o *KvmSession) HasSsoSupported() bool`

HasSsoSupported returns a boolean if a field has been set.

### GetUsername

`func (o *KvmSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KvmSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KvmSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KvmSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDevice

`func (o *KvmSession) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *KvmSession) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *KvmSession) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *KvmSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *KvmSession) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KvmSession) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KvmSession) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *KvmSession) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetTunnel

`func (o *KvmSession) GetTunnel() KvmTunnelRelationship`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *KvmSession) GetTunnelOk() (*KvmTunnelRelationship, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *KvmSession) SetTunnel(v KvmTunnelRelationship)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *KvmSession) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


