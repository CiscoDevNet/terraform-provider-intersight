# StoragePureNfsPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureNfsPolicyRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureNfsPolicyRule"]
**Access** | Pointer to **string** | Access of the policy rule. | [optional] [readonly] 
**Anongid** | Pointer to **string** | Anongid of the policy rule. | [optional] [readonly] 
**Anonuid** | Pointer to **string** | Anonuid of the policy rule. | [optional] [readonly] 
**Client** | Pointer to **string** | Client of the policy rule. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the NFS policy rule. | [optional] [readonly] 
**NfsVersion** | Pointer to **[]string** |  | [optional] 
**Permission** | Pointer to **string** | Permission of the policy rule. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Name of the NFS policy rule. | [optional] [readonly] 
**Security** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableStoragePureDirectoryPolicyRelationship**](StoragePureDirectoryPolicyRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureNfsPolicyRule

`func NewStoragePureNfsPolicyRule(classId string, objectType string, ) *StoragePureNfsPolicyRule`

NewStoragePureNfsPolicyRule instantiates a new StoragePureNfsPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureNfsPolicyRuleWithDefaults

`func NewStoragePureNfsPolicyRuleWithDefaults() *StoragePureNfsPolicyRule`

NewStoragePureNfsPolicyRuleWithDefaults instantiates a new StoragePureNfsPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureNfsPolicyRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureNfsPolicyRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureNfsPolicyRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureNfsPolicyRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureNfsPolicyRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureNfsPolicyRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccess

`func (o *StoragePureNfsPolicyRule) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *StoragePureNfsPolicyRule) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *StoragePureNfsPolicyRule) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *StoragePureNfsPolicyRule) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAnongid

`func (o *StoragePureNfsPolicyRule) GetAnongid() string`

GetAnongid returns the Anongid field if non-nil, zero value otherwise.

### GetAnongidOk

`func (o *StoragePureNfsPolicyRule) GetAnongidOk() (*string, bool)`

GetAnongidOk returns a tuple with the Anongid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnongid

`func (o *StoragePureNfsPolicyRule) SetAnongid(v string)`

SetAnongid sets Anongid field to given value.

### HasAnongid

`func (o *StoragePureNfsPolicyRule) HasAnongid() bool`

HasAnongid returns a boolean if a field has been set.

### GetAnonuid

`func (o *StoragePureNfsPolicyRule) GetAnonuid() string`

GetAnonuid returns the Anonuid field if non-nil, zero value otherwise.

### GetAnonuidOk

`func (o *StoragePureNfsPolicyRule) GetAnonuidOk() (*string, bool)`

GetAnonuidOk returns a tuple with the Anonuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonuid

`func (o *StoragePureNfsPolicyRule) SetAnonuid(v string)`

SetAnonuid sets Anonuid field to given value.

### HasAnonuid

`func (o *StoragePureNfsPolicyRule) HasAnonuid() bool`

HasAnonuid returns a boolean if a field has been set.

### GetClient

`func (o *StoragePureNfsPolicyRule) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *StoragePureNfsPolicyRule) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *StoragePureNfsPolicyRule) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *StoragePureNfsPolicyRule) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureNfsPolicyRule) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureNfsPolicyRule) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureNfsPolicyRule) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureNfsPolicyRule) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetName

`func (o *StoragePureNfsPolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureNfsPolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureNfsPolicyRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureNfsPolicyRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsVersion

`func (o *StoragePureNfsPolicyRule) GetNfsVersion() []string`

GetNfsVersion returns the NfsVersion field if non-nil, zero value otherwise.

### GetNfsVersionOk

`func (o *StoragePureNfsPolicyRule) GetNfsVersionOk() (*[]string, bool)`

GetNfsVersionOk returns a tuple with the NfsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersion

`func (o *StoragePureNfsPolicyRule) SetNfsVersion(v []string)`

SetNfsVersion sets NfsVersion field to given value.

### HasNfsVersion

`func (o *StoragePureNfsPolicyRule) HasNfsVersion() bool`

HasNfsVersion returns a boolean if a field has been set.

### SetNfsVersionNil

`func (o *StoragePureNfsPolicyRule) SetNfsVersionNil(b bool)`

 SetNfsVersionNil sets the value for NfsVersion to be an explicit nil

### UnsetNfsVersion
`func (o *StoragePureNfsPolicyRule) UnsetNfsVersion()`

UnsetNfsVersion ensures that no value is present for NfsVersion, not even an explicit nil
### GetPermission

`func (o *StoragePureNfsPolicyRule) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *StoragePureNfsPolicyRule) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *StoragePureNfsPolicyRule) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *StoragePureNfsPolicyRule) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPolicyName

`func (o *StoragePureNfsPolicyRule) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StoragePureNfsPolicyRule) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StoragePureNfsPolicyRule) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StoragePureNfsPolicyRule) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetSecurity

`func (o *StoragePureNfsPolicyRule) GetSecurity() []string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *StoragePureNfsPolicyRule) GetSecurityOk() (*[]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *StoragePureNfsPolicyRule) SetSecurity(v []string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *StoragePureNfsPolicyRule) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *StoragePureNfsPolicyRule) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *StoragePureNfsPolicyRule) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetArray

`func (o *StoragePureNfsPolicyRule) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureNfsPolicyRule) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureNfsPolicyRule) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureNfsPolicyRule) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureNfsPolicyRule) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureNfsPolicyRule) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetPolicy

`func (o *StoragePureNfsPolicyRule) GetPolicy() StoragePureDirectoryPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StoragePureNfsPolicyRule) GetPolicyOk() (*StoragePureDirectoryPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StoragePureNfsPolicyRule) SetPolicy(v StoragePureDirectoryPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StoragePureNfsPolicyRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *StoragePureNfsPolicyRule) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *StoragePureNfsPolicyRule) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureNfsPolicyRule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureNfsPolicyRule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureNfsPolicyRule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureNfsPolicyRule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureNfsPolicyRule) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureNfsPolicyRule) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


