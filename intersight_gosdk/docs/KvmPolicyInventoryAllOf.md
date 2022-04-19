# KvmPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.PolicyInventory"]
**EnableLocalServerVideo** | Pointer to **bool** | If enabled, displays KVM session on any monitor attached to the server. | [optional] [readonly] [default to true]
**EnableVideoEncryption** | Pointer to **bool** | If enabled, encrypts all video information sent through KVM. Please note that this is no longer applicable for servers running versions 4.2 and above. | [optional] [readonly] [default to true]
**Enabled** | Pointer to **bool** | State of the vKVM service on the endpoint. | [optional] [readonly] [default to true]
**MaximumSessions** | Pointer to **int64** | The maximum number of concurrent KVM sessions allowed. | [optional] [readonly] [default to 4]
**RemotePort** | Pointer to **int64** | The port used for KVM communication. | [optional] [readonly] [default to 2068]
**TunneledKvmEnabled** | Pointer to **bool** | Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM. | [optional] [readonly] [default to false]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewKvmPolicyInventoryAllOf

`func NewKvmPolicyInventoryAllOf(classId string, objectType string, ) *KvmPolicyInventoryAllOf`

NewKvmPolicyInventoryAllOf instantiates a new KvmPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmPolicyInventoryAllOfWithDefaults

`func NewKvmPolicyInventoryAllOfWithDefaults() *KvmPolicyInventoryAllOf`

NewKvmPolicyInventoryAllOfWithDefaults instantiates a new KvmPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableLocalServerVideo

`func (o *KvmPolicyInventoryAllOf) GetEnableLocalServerVideo() bool`

GetEnableLocalServerVideo returns the EnableLocalServerVideo field if non-nil, zero value otherwise.

### GetEnableLocalServerVideoOk

`func (o *KvmPolicyInventoryAllOf) GetEnableLocalServerVideoOk() (*bool, bool)`

GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalServerVideo

`func (o *KvmPolicyInventoryAllOf) SetEnableLocalServerVideo(v bool)`

SetEnableLocalServerVideo sets EnableLocalServerVideo field to given value.

### HasEnableLocalServerVideo

`func (o *KvmPolicyInventoryAllOf) HasEnableLocalServerVideo() bool`

HasEnableLocalServerVideo returns a boolean if a field has been set.

### GetEnableVideoEncryption

`func (o *KvmPolicyInventoryAllOf) GetEnableVideoEncryption() bool`

GetEnableVideoEncryption returns the EnableVideoEncryption field if non-nil, zero value otherwise.

### GetEnableVideoEncryptionOk

`func (o *KvmPolicyInventoryAllOf) GetEnableVideoEncryptionOk() (*bool, bool)`

GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoEncryption

`func (o *KvmPolicyInventoryAllOf) SetEnableVideoEncryption(v bool)`

SetEnableVideoEncryption sets EnableVideoEncryption field to given value.

### HasEnableVideoEncryption

`func (o *KvmPolicyInventoryAllOf) HasEnableVideoEncryption() bool`

HasEnableVideoEncryption returns a boolean if a field has been set.

### GetEnabled

`func (o *KvmPolicyInventoryAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KvmPolicyInventoryAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KvmPolicyInventoryAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KvmPolicyInventoryAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumSessions

`func (o *KvmPolicyInventoryAllOf) GetMaximumSessions() int64`

GetMaximumSessions returns the MaximumSessions field if non-nil, zero value otherwise.

### GetMaximumSessionsOk

`func (o *KvmPolicyInventoryAllOf) GetMaximumSessionsOk() (*int64, bool)`

GetMaximumSessionsOk returns a tuple with the MaximumSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessions

`func (o *KvmPolicyInventoryAllOf) SetMaximumSessions(v int64)`

SetMaximumSessions sets MaximumSessions field to given value.

### HasMaximumSessions

`func (o *KvmPolicyInventoryAllOf) HasMaximumSessions() bool`

HasMaximumSessions returns a boolean if a field has been set.

### GetRemotePort

`func (o *KvmPolicyInventoryAllOf) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *KvmPolicyInventoryAllOf) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *KvmPolicyInventoryAllOf) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *KvmPolicyInventoryAllOf) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetTunneledKvmEnabled

`func (o *KvmPolicyInventoryAllOf) GetTunneledKvmEnabled() bool`

GetTunneledKvmEnabled returns the TunneledKvmEnabled field if non-nil, zero value otherwise.

### GetTunneledKvmEnabledOk

`func (o *KvmPolicyInventoryAllOf) GetTunneledKvmEnabledOk() (*bool, bool)`

GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmEnabled

`func (o *KvmPolicyInventoryAllOf) SetTunneledKvmEnabled(v bool)`

SetTunneledKvmEnabled sets TunneledKvmEnabled field to given value.

### HasTunneledKvmEnabled

`func (o *KvmPolicyInventoryAllOf) HasTunneledKvmEnabled() bool`

HasTunneledKvmEnabled returns a boolean if a field has been set.

### GetTargetMo

`func (o *KvmPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *KvmPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *KvmPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *KvmPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


