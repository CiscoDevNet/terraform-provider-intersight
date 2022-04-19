# IpmioverlanPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ipmioverlan.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ipmioverlan.PolicyInventory"]
**Enabled** | Pointer to **bool** | State of the IPMI Over LAN service on the endpoint. | [optional] [readonly] [default to true]
**EncryptionKey** | Pointer to **string** | The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. | [optional] 
**IsEncryptionKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;encryptionKey&#39; property has been set. | [optional] [readonly] [default to false]
**Privilege** | Pointer to **string** | The highest privilege level that can be assigned to an IPMI session on a server. * &#x60;admin&#x60; - Privilege to perform all actions available through IPMI. * &#x60;user&#x60; - Privilege to perform some functions through IPMI but restriction on performing administrative tasks. * &#x60;read-only&#x60; - Privilege to view information throught IPMI but restriction on making any changes. | [optional] [readonly] [default to "admin"]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIpmioverlanPolicyInventoryAllOf

`func NewIpmioverlanPolicyInventoryAllOf(classId string, objectType string, ) *IpmioverlanPolicyInventoryAllOf`

NewIpmioverlanPolicyInventoryAllOf instantiates a new IpmioverlanPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmioverlanPolicyInventoryAllOfWithDefaults

`func NewIpmioverlanPolicyInventoryAllOfWithDefaults() *IpmioverlanPolicyInventoryAllOf`

NewIpmioverlanPolicyInventoryAllOfWithDefaults instantiates a new IpmioverlanPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IpmioverlanPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IpmioverlanPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IpmioverlanPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IpmioverlanPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *IpmioverlanPolicyInventoryAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IpmioverlanPolicyInventoryAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IpmioverlanPolicyInventoryAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *IpmioverlanPolicyInventoryAllOf) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *IpmioverlanPolicyInventoryAllOf) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *IpmioverlanPolicyInventoryAllOf) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetIsEncryptionKeySet

`func (o *IpmioverlanPolicyInventoryAllOf) GetIsEncryptionKeySet() bool`

GetIsEncryptionKeySet returns the IsEncryptionKeySet field if non-nil, zero value otherwise.

### GetIsEncryptionKeySetOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetIsEncryptionKeySetOk() (*bool, bool)`

GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptionKeySet

`func (o *IpmioverlanPolicyInventoryAllOf) SetIsEncryptionKeySet(v bool)`

SetIsEncryptionKeySet sets IsEncryptionKeySet field to given value.

### HasIsEncryptionKeySet

`func (o *IpmioverlanPolicyInventoryAllOf) HasIsEncryptionKeySet() bool`

HasIsEncryptionKeySet returns a boolean if a field has been set.

### GetPrivilege

`func (o *IpmioverlanPolicyInventoryAllOf) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *IpmioverlanPolicyInventoryAllOf) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *IpmioverlanPolicyInventoryAllOf) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetTargetMo

`func (o *IpmioverlanPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *IpmioverlanPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *IpmioverlanPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *IpmioverlanPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


