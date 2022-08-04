# KvmPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.PolicyInventory"]
**EnableLocalServerVideo** | Pointer to **bool** | If enabled, displays KVM session on any monitor attached to the server. | [optional] [readonly] [default to true]
**EnableVideoEncryption** | Pointer to **bool** | If enabled, encrypts all video information sent through KVM. Please note that this can no longer be disabled for servers running versions 4.2 and above. | [optional] [readonly] [default to true]
**Enabled** | Pointer to **bool** | State of the vKVM service on the endpoint. | [optional] [readonly] [default to true]
**MaximumSessions** | Pointer to **int64** | The maximum number of concurrent KVM sessions allowed. | [optional] [readonly] [default to 4]
**RemotePort** | Pointer to **int64** | The port used for KVM communication. | [optional] [readonly] [default to 2068]
**TunneledKvmEnabled** | Pointer to **bool** | Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM. | [optional] [readonly] [default to false]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewKvmPolicyInventory

`func NewKvmPolicyInventory(classId string, objectType string, ) *KvmPolicyInventory`

NewKvmPolicyInventory instantiates a new KvmPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmPolicyInventoryWithDefaults

`func NewKvmPolicyInventoryWithDefaults() *KvmPolicyInventory`

NewKvmPolicyInventoryWithDefaults instantiates a new KvmPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableLocalServerVideo

`func (o *KvmPolicyInventory) GetEnableLocalServerVideo() bool`

GetEnableLocalServerVideo returns the EnableLocalServerVideo field if non-nil, zero value otherwise.

### GetEnableLocalServerVideoOk

`func (o *KvmPolicyInventory) GetEnableLocalServerVideoOk() (*bool, bool)`

GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalServerVideo

`func (o *KvmPolicyInventory) SetEnableLocalServerVideo(v bool)`

SetEnableLocalServerVideo sets EnableLocalServerVideo field to given value.

### HasEnableLocalServerVideo

`func (o *KvmPolicyInventory) HasEnableLocalServerVideo() bool`

HasEnableLocalServerVideo returns a boolean if a field has been set.

### GetEnableVideoEncryption

`func (o *KvmPolicyInventory) GetEnableVideoEncryption() bool`

GetEnableVideoEncryption returns the EnableVideoEncryption field if non-nil, zero value otherwise.

### GetEnableVideoEncryptionOk

`func (o *KvmPolicyInventory) GetEnableVideoEncryptionOk() (*bool, bool)`

GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoEncryption

`func (o *KvmPolicyInventory) SetEnableVideoEncryption(v bool)`

SetEnableVideoEncryption sets EnableVideoEncryption field to given value.

### HasEnableVideoEncryption

`func (o *KvmPolicyInventory) HasEnableVideoEncryption() bool`

HasEnableVideoEncryption returns a boolean if a field has been set.

### GetEnabled

`func (o *KvmPolicyInventory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KvmPolicyInventory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KvmPolicyInventory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KvmPolicyInventory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumSessions

`func (o *KvmPolicyInventory) GetMaximumSessions() int64`

GetMaximumSessions returns the MaximumSessions field if non-nil, zero value otherwise.

### GetMaximumSessionsOk

`func (o *KvmPolicyInventory) GetMaximumSessionsOk() (*int64, bool)`

GetMaximumSessionsOk returns a tuple with the MaximumSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessions

`func (o *KvmPolicyInventory) SetMaximumSessions(v int64)`

SetMaximumSessions sets MaximumSessions field to given value.

### HasMaximumSessions

`func (o *KvmPolicyInventory) HasMaximumSessions() bool`

HasMaximumSessions returns a boolean if a field has been set.

### GetRemotePort

`func (o *KvmPolicyInventory) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *KvmPolicyInventory) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *KvmPolicyInventory) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *KvmPolicyInventory) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetTunneledKvmEnabled

`func (o *KvmPolicyInventory) GetTunneledKvmEnabled() bool`

GetTunneledKvmEnabled returns the TunneledKvmEnabled field if non-nil, zero value otherwise.

### GetTunneledKvmEnabledOk

`func (o *KvmPolicyInventory) GetTunneledKvmEnabledOk() (*bool, bool)`

GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmEnabled

`func (o *KvmPolicyInventory) SetTunneledKvmEnabled(v bool)`

SetTunneledKvmEnabled sets TunneledKvmEnabled field to given value.

### HasTunneledKvmEnabled

`func (o *KvmPolicyInventory) HasTunneledKvmEnabled() bool`

HasTunneledKvmEnabled returns a boolean if a field has been set.

### GetTargetMo

`func (o *KvmPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *KvmPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *KvmPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *KvmPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


