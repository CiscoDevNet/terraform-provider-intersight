# StoragePureDirectoryQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureDirectoryQuota"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureDirectoryQuota"]
**Destroyed** | Pointer to **bool** | Returns a value of true if the managed directory of the export has been destroyed and is pending eradication. The export can be recovered by recovering the destroyed managed directory. | [optional] [readonly] 
**DirectoryName** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**DirectoryResourceType** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Returns a value of true if the export policy that manages this export is enabled. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | The export policy that manages this export. An export can be managed by at most one export policy. | [optional] [readonly] 
**PolicyResourceType** | Pointer to **string** | The export policy that manages this export. An export can be managed by at most one export policy. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Directory** | Pointer to [**NullableStoragePureDirectoryRelationship**](StoragePureDirectoryRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableStoragePureDirectoryPolicyRelationship**](StoragePureDirectoryPolicyRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureDirectoryQuota

`func NewStoragePureDirectoryQuota(classId string, objectType string, ) *StoragePureDirectoryQuota`

NewStoragePureDirectoryQuota instantiates a new StoragePureDirectoryQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureDirectoryQuotaWithDefaults

`func NewStoragePureDirectoryQuotaWithDefaults() *StoragePureDirectoryQuota`

NewStoragePureDirectoryQuotaWithDefaults instantiates a new StoragePureDirectoryQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureDirectoryQuota) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureDirectoryQuota) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureDirectoryQuota) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureDirectoryQuota) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureDirectoryQuota) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureDirectoryQuota) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestroyed

`func (o *StoragePureDirectoryQuota) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureDirectoryQuota) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureDirectoryQuota) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureDirectoryQuota) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetDirectoryName

`func (o *StoragePureDirectoryQuota) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *StoragePureDirectoryQuota) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *StoragePureDirectoryQuota) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *StoragePureDirectoryQuota) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### GetDirectoryResourceType

`func (o *StoragePureDirectoryQuota) GetDirectoryResourceType() string`

GetDirectoryResourceType returns the DirectoryResourceType field if non-nil, zero value otherwise.

### GetDirectoryResourceTypeOk

`func (o *StoragePureDirectoryQuota) GetDirectoryResourceTypeOk() (*string, bool)`

GetDirectoryResourceTypeOk returns a tuple with the DirectoryResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResourceType

`func (o *StoragePureDirectoryQuota) SetDirectoryResourceType(v string)`

SetDirectoryResourceType sets DirectoryResourceType field to given value.

### HasDirectoryResourceType

`func (o *StoragePureDirectoryQuota) HasDirectoryResourceType() bool`

HasDirectoryResourceType returns a boolean if a field has been set.

### GetEnabled

`func (o *StoragePureDirectoryQuota) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoragePureDirectoryQuota) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoragePureDirectoryQuota) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StoragePureDirectoryQuota) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicyName

`func (o *StoragePureDirectoryQuota) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StoragePureDirectoryQuota) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StoragePureDirectoryQuota) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StoragePureDirectoryQuota) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyResourceType

`func (o *StoragePureDirectoryQuota) GetPolicyResourceType() string`

GetPolicyResourceType returns the PolicyResourceType field if non-nil, zero value otherwise.

### GetPolicyResourceTypeOk

`func (o *StoragePureDirectoryQuota) GetPolicyResourceTypeOk() (*string, bool)`

GetPolicyResourceTypeOk returns a tuple with the PolicyResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyResourceType

`func (o *StoragePureDirectoryQuota) SetPolicyResourceType(v string)`

SetPolicyResourceType sets PolicyResourceType field to given value.

### HasPolicyResourceType

`func (o *StoragePureDirectoryQuota) HasPolicyResourceType() bool`

HasPolicyResourceType returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureDirectoryQuota) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureDirectoryQuota) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureDirectoryQuota) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureDirectoryQuota) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureDirectoryQuota) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureDirectoryQuota) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetDirectory

`func (o *StoragePureDirectoryQuota) GetDirectory() StoragePureDirectoryRelationship`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *StoragePureDirectoryQuota) GetDirectoryOk() (*StoragePureDirectoryRelationship, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *StoragePureDirectoryQuota) SetDirectory(v StoragePureDirectoryRelationship)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *StoragePureDirectoryQuota) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### SetDirectoryNil

`func (o *StoragePureDirectoryQuota) SetDirectoryNil(b bool)`

 SetDirectoryNil sets the value for Directory to be an explicit nil

### UnsetDirectory
`func (o *StoragePureDirectoryQuota) UnsetDirectory()`

UnsetDirectory ensures that no value is present for Directory, not even an explicit nil
### GetPolicy

`func (o *StoragePureDirectoryQuota) GetPolicy() StoragePureDirectoryPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StoragePureDirectoryQuota) GetPolicyOk() (*StoragePureDirectoryPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StoragePureDirectoryQuota) SetPolicy(v StoragePureDirectoryPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StoragePureDirectoryQuota) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *StoragePureDirectoryQuota) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *StoragePureDirectoryQuota) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureDirectoryQuota) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureDirectoryQuota) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureDirectoryQuota) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureDirectoryQuota) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureDirectoryQuota) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureDirectoryQuota) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


