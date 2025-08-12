# StoragePureSmbPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureSmbPolicyRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureSmbPolicyRule"]
**AnonymousAccessAllowed** | Pointer to **bool** | Anonymous access allowed of the policy rule. | [optional] [readonly] 
**ClientName** | Pointer to **string** | Client of the policy rule. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the SMB policy rule. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Name of the SMB policy rule. | [optional] [readonly] 
**SmbEncryptionRequired** | Pointer to **bool** | SMB encryption required of the policy rule. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableStoragePureDirectoryPolicyRelationship**](StoragePureDirectoryPolicyRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureSmbPolicyRule

`func NewStoragePureSmbPolicyRule(classId string, objectType string, ) *StoragePureSmbPolicyRule`

NewStoragePureSmbPolicyRule instantiates a new StoragePureSmbPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureSmbPolicyRuleWithDefaults

`func NewStoragePureSmbPolicyRuleWithDefaults() *StoragePureSmbPolicyRule`

NewStoragePureSmbPolicyRuleWithDefaults instantiates a new StoragePureSmbPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureSmbPolicyRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureSmbPolicyRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureSmbPolicyRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureSmbPolicyRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureSmbPolicyRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureSmbPolicyRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnonymousAccessAllowed

`func (o *StoragePureSmbPolicyRule) GetAnonymousAccessAllowed() bool`

GetAnonymousAccessAllowed returns the AnonymousAccessAllowed field if non-nil, zero value otherwise.

### GetAnonymousAccessAllowedOk

`func (o *StoragePureSmbPolicyRule) GetAnonymousAccessAllowedOk() (*bool, bool)`

GetAnonymousAccessAllowedOk returns a tuple with the AnonymousAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessAllowed

`func (o *StoragePureSmbPolicyRule) SetAnonymousAccessAllowed(v bool)`

SetAnonymousAccessAllowed sets AnonymousAccessAllowed field to given value.

### HasAnonymousAccessAllowed

`func (o *StoragePureSmbPolicyRule) HasAnonymousAccessAllowed() bool`

HasAnonymousAccessAllowed returns a boolean if a field has been set.

### GetClientName

`func (o *StoragePureSmbPolicyRule) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *StoragePureSmbPolicyRule) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *StoragePureSmbPolicyRule) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *StoragePureSmbPolicyRule) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureSmbPolicyRule) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureSmbPolicyRule) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureSmbPolicyRule) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureSmbPolicyRule) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetName

`func (o *StoragePureSmbPolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureSmbPolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureSmbPolicyRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureSmbPolicyRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyName

`func (o *StoragePureSmbPolicyRule) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StoragePureSmbPolicyRule) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StoragePureSmbPolicyRule) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StoragePureSmbPolicyRule) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetSmbEncryptionRequired

`func (o *StoragePureSmbPolicyRule) GetSmbEncryptionRequired() bool`

GetSmbEncryptionRequired returns the SmbEncryptionRequired field if non-nil, zero value otherwise.

### GetSmbEncryptionRequiredOk

`func (o *StoragePureSmbPolicyRule) GetSmbEncryptionRequiredOk() (*bool, bool)`

GetSmbEncryptionRequiredOk returns a tuple with the SmbEncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbEncryptionRequired

`func (o *StoragePureSmbPolicyRule) SetSmbEncryptionRequired(v bool)`

SetSmbEncryptionRequired sets SmbEncryptionRequired field to given value.

### HasSmbEncryptionRequired

`func (o *StoragePureSmbPolicyRule) HasSmbEncryptionRequired() bool`

HasSmbEncryptionRequired returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureSmbPolicyRule) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureSmbPolicyRule) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureSmbPolicyRule) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureSmbPolicyRule) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureSmbPolicyRule) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureSmbPolicyRule) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetPolicy

`func (o *StoragePureSmbPolicyRule) GetPolicy() StoragePureDirectoryPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StoragePureSmbPolicyRule) GetPolicyOk() (*StoragePureDirectoryPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StoragePureSmbPolicyRule) SetPolicy(v StoragePureDirectoryPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StoragePureSmbPolicyRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *StoragePureSmbPolicyRule) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *StoragePureSmbPolicyRule) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureSmbPolicyRule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureSmbPolicyRule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureSmbPolicyRule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureSmbPolicyRule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureSmbPolicyRule) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureSmbPolicyRule) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


