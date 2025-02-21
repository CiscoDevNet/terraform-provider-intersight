# StorageHitachiNvmSubsystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiNvmSubsystem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiNvmSubsystem"]
**HostMode** | Pointer to **string** | Host mode of the NVM subsystem. | [optional] [readonly] 
**HostModeOptions** | Pointer to **[]int64** |  | [optional] 
**HostNqns** | Pointer to [**[]StorageHostNqn**](StorageHostNqn.md) |  | [optional] 
**Name** | Pointer to **string** | NVM subsystem ID. NVM subsystem is a resource to manage system components in an NVMe-oF connection. | [optional] [readonly] 
**NamespacePaths** | Pointer to [**[]StorageNamespacePath**](StorageNamespacePath.md) |  | [optional] 
**NamespaceSecuritySetting** | Pointer to **string** | Namespace security settings. | [optional] [readonly] 
**Namespaces** | Pointer to [**[]StorageNamespace**](StorageNamespace.md) |  | [optional] 
**NvmSubsystemName** | Pointer to **string** | NVM subsystem name. Can be up to 32 characters long. | [optional] [readonly] 
**NvmSubsystemNqn** | Pointer to **string** | Subsystem NQN. If the NVM subsystem is virtualized, the NQN of the virtualized NVM subsystem is output. | [optional] [readonly] 
**PortIds** | Pointer to **[]string** |  | [optional] 
**ResourceGroupId** | Pointer to **int64** | Resource group ID of the resource group to which the NVM subsystem belongs. | [optional] [readonly] 
**T10piMode** | Pointer to **string** | Status of the T10 PI mode of the port. | [optional] [readonly] 
**VirtualNvmSubsystemId** | Pointer to **int64** | The Virtual NVM Subsystem ID property is applicable for use with storage systems in the VSP 5000 series. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiNvmSubsystem

`func NewStorageHitachiNvmSubsystem(classId string, objectType string, ) *StorageHitachiNvmSubsystem`

NewStorageHitachiNvmSubsystem instantiates a new StorageHitachiNvmSubsystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiNvmSubsystemWithDefaults

`func NewStorageHitachiNvmSubsystemWithDefaults() *StorageHitachiNvmSubsystem`

NewStorageHitachiNvmSubsystemWithDefaults instantiates a new StorageHitachiNvmSubsystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiNvmSubsystem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiNvmSubsystem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiNvmSubsystem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiNvmSubsystem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiNvmSubsystem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiNvmSubsystem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostMode

`func (o *StorageHitachiNvmSubsystem) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *StorageHitachiNvmSubsystem) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *StorageHitachiNvmSubsystem) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.

### HasHostMode

`func (o *StorageHitachiNvmSubsystem) HasHostMode() bool`

HasHostMode returns a boolean if a field has been set.

### GetHostModeOptions

`func (o *StorageHitachiNvmSubsystem) GetHostModeOptions() []int64`

GetHostModeOptions returns the HostModeOptions field if non-nil, zero value otherwise.

### GetHostModeOptionsOk

`func (o *StorageHitachiNvmSubsystem) GetHostModeOptionsOk() (*[]int64, bool)`

GetHostModeOptionsOk returns a tuple with the HostModeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostModeOptions

`func (o *StorageHitachiNvmSubsystem) SetHostModeOptions(v []int64)`

SetHostModeOptions sets HostModeOptions field to given value.

### HasHostModeOptions

`func (o *StorageHitachiNvmSubsystem) HasHostModeOptions() bool`

HasHostModeOptions returns a boolean if a field has been set.

### SetHostModeOptionsNil

`func (o *StorageHitachiNvmSubsystem) SetHostModeOptionsNil(b bool)`

 SetHostModeOptionsNil sets the value for HostModeOptions to be an explicit nil

### UnsetHostModeOptions
`func (o *StorageHitachiNvmSubsystem) UnsetHostModeOptions()`

UnsetHostModeOptions ensures that no value is present for HostModeOptions, not even an explicit nil
### GetHostNqns

`func (o *StorageHitachiNvmSubsystem) GetHostNqns() []StorageHostNqn`

GetHostNqns returns the HostNqns field if non-nil, zero value otherwise.

### GetHostNqnsOk

`func (o *StorageHitachiNvmSubsystem) GetHostNqnsOk() (*[]StorageHostNqn, bool)`

GetHostNqnsOk returns a tuple with the HostNqns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNqns

`func (o *StorageHitachiNvmSubsystem) SetHostNqns(v []StorageHostNqn)`

SetHostNqns sets HostNqns field to given value.

### HasHostNqns

`func (o *StorageHitachiNvmSubsystem) HasHostNqns() bool`

HasHostNqns returns a boolean if a field has been set.

### SetHostNqnsNil

`func (o *StorageHitachiNvmSubsystem) SetHostNqnsNil(b bool)`

 SetHostNqnsNil sets the value for HostNqns to be an explicit nil

### UnsetHostNqns
`func (o *StorageHitachiNvmSubsystem) UnsetHostNqns()`

UnsetHostNqns ensures that no value is present for HostNqns, not even an explicit nil
### GetName

`func (o *StorageHitachiNvmSubsystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiNvmSubsystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiNvmSubsystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiNvmSubsystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespacePaths

`func (o *StorageHitachiNvmSubsystem) GetNamespacePaths() []StorageNamespacePath`

GetNamespacePaths returns the NamespacePaths field if non-nil, zero value otherwise.

### GetNamespacePathsOk

`func (o *StorageHitachiNvmSubsystem) GetNamespacePathsOk() (*[]StorageNamespacePath, bool)`

GetNamespacePathsOk returns a tuple with the NamespacePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacePaths

`func (o *StorageHitachiNvmSubsystem) SetNamespacePaths(v []StorageNamespacePath)`

SetNamespacePaths sets NamespacePaths field to given value.

### HasNamespacePaths

`func (o *StorageHitachiNvmSubsystem) HasNamespacePaths() bool`

HasNamespacePaths returns a boolean if a field has been set.

### SetNamespacePathsNil

`func (o *StorageHitachiNvmSubsystem) SetNamespacePathsNil(b bool)`

 SetNamespacePathsNil sets the value for NamespacePaths to be an explicit nil

### UnsetNamespacePaths
`func (o *StorageHitachiNvmSubsystem) UnsetNamespacePaths()`

UnsetNamespacePaths ensures that no value is present for NamespacePaths, not even an explicit nil
### GetNamespaceSecuritySetting

`func (o *StorageHitachiNvmSubsystem) GetNamespaceSecuritySetting() string`

GetNamespaceSecuritySetting returns the NamespaceSecuritySetting field if non-nil, zero value otherwise.

### GetNamespaceSecuritySettingOk

`func (o *StorageHitachiNvmSubsystem) GetNamespaceSecuritySettingOk() (*string, bool)`

GetNamespaceSecuritySettingOk returns a tuple with the NamespaceSecuritySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceSecuritySetting

`func (o *StorageHitachiNvmSubsystem) SetNamespaceSecuritySetting(v string)`

SetNamespaceSecuritySetting sets NamespaceSecuritySetting field to given value.

### HasNamespaceSecuritySetting

`func (o *StorageHitachiNvmSubsystem) HasNamespaceSecuritySetting() bool`

HasNamespaceSecuritySetting returns a boolean if a field has been set.

### GetNamespaces

`func (o *StorageHitachiNvmSubsystem) GetNamespaces() []StorageNamespace`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *StorageHitachiNvmSubsystem) GetNamespacesOk() (*[]StorageNamespace, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *StorageHitachiNvmSubsystem) SetNamespaces(v []StorageNamespace)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *StorageHitachiNvmSubsystem) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### SetNamespacesNil

`func (o *StorageHitachiNvmSubsystem) SetNamespacesNil(b bool)`

 SetNamespacesNil sets the value for Namespaces to be an explicit nil

### UnsetNamespaces
`func (o *StorageHitachiNvmSubsystem) UnsetNamespaces()`

UnsetNamespaces ensures that no value is present for Namespaces, not even an explicit nil
### GetNvmSubsystemName

`func (o *StorageHitachiNvmSubsystem) GetNvmSubsystemName() string`

GetNvmSubsystemName returns the NvmSubsystemName field if non-nil, zero value otherwise.

### GetNvmSubsystemNameOk

`func (o *StorageHitachiNvmSubsystem) GetNvmSubsystemNameOk() (*string, bool)`

GetNvmSubsystemNameOk returns a tuple with the NvmSubsystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmSubsystemName

`func (o *StorageHitachiNvmSubsystem) SetNvmSubsystemName(v string)`

SetNvmSubsystemName sets NvmSubsystemName field to given value.

### HasNvmSubsystemName

`func (o *StorageHitachiNvmSubsystem) HasNvmSubsystemName() bool`

HasNvmSubsystemName returns a boolean if a field has been set.

### GetNvmSubsystemNqn

`func (o *StorageHitachiNvmSubsystem) GetNvmSubsystemNqn() string`

GetNvmSubsystemNqn returns the NvmSubsystemNqn field if non-nil, zero value otherwise.

### GetNvmSubsystemNqnOk

`func (o *StorageHitachiNvmSubsystem) GetNvmSubsystemNqnOk() (*string, bool)`

GetNvmSubsystemNqnOk returns a tuple with the NvmSubsystemNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmSubsystemNqn

`func (o *StorageHitachiNvmSubsystem) SetNvmSubsystemNqn(v string)`

SetNvmSubsystemNqn sets NvmSubsystemNqn field to given value.

### HasNvmSubsystemNqn

`func (o *StorageHitachiNvmSubsystem) HasNvmSubsystemNqn() bool`

HasNvmSubsystemNqn returns a boolean if a field has been set.

### GetPortIds

`func (o *StorageHitachiNvmSubsystem) GetPortIds() []string`

GetPortIds returns the PortIds field if non-nil, zero value otherwise.

### GetPortIdsOk

`func (o *StorageHitachiNvmSubsystem) GetPortIdsOk() (*[]string, bool)`

GetPortIdsOk returns a tuple with the PortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIds

`func (o *StorageHitachiNvmSubsystem) SetPortIds(v []string)`

SetPortIds sets PortIds field to given value.

### HasPortIds

`func (o *StorageHitachiNvmSubsystem) HasPortIds() bool`

HasPortIds returns a boolean if a field has been set.

### SetPortIdsNil

`func (o *StorageHitachiNvmSubsystem) SetPortIdsNil(b bool)`

 SetPortIdsNil sets the value for PortIds to be an explicit nil

### UnsetPortIds
`func (o *StorageHitachiNvmSubsystem) UnsetPortIds()`

UnsetPortIds ensures that no value is present for PortIds, not even an explicit nil
### GetResourceGroupId

`func (o *StorageHitachiNvmSubsystem) GetResourceGroupId() int64`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *StorageHitachiNvmSubsystem) GetResourceGroupIdOk() (*int64, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *StorageHitachiNvmSubsystem) SetResourceGroupId(v int64)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *StorageHitachiNvmSubsystem) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### GetT10piMode

`func (o *StorageHitachiNvmSubsystem) GetT10piMode() string`

GetT10piMode returns the T10piMode field if non-nil, zero value otherwise.

### GetT10piModeOk

`func (o *StorageHitachiNvmSubsystem) GetT10piModeOk() (*string, bool)`

GetT10piModeOk returns a tuple with the T10piMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT10piMode

`func (o *StorageHitachiNvmSubsystem) SetT10piMode(v string)`

SetT10piMode sets T10piMode field to given value.

### HasT10piMode

`func (o *StorageHitachiNvmSubsystem) HasT10piMode() bool`

HasT10piMode returns a boolean if a field has been set.

### GetVirtualNvmSubsystemId

`func (o *StorageHitachiNvmSubsystem) GetVirtualNvmSubsystemId() int64`

GetVirtualNvmSubsystemId returns the VirtualNvmSubsystemId field if non-nil, zero value otherwise.

### GetVirtualNvmSubsystemIdOk

`func (o *StorageHitachiNvmSubsystem) GetVirtualNvmSubsystemIdOk() (*int64, bool)`

GetVirtualNvmSubsystemIdOk returns a tuple with the VirtualNvmSubsystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNvmSubsystemId

`func (o *StorageHitachiNvmSubsystem) SetVirtualNvmSubsystemId(v int64)`

SetVirtualNvmSubsystemId sets VirtualNvmSubsystemId field to given value.

### HasVirtualNvmSubsystemId

`func (o *StorageHitachiNvmSubsystem) HasVirtualNvmSubsystemId() bool`

HasVirtualNvmSubsystemId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiNvmSubsystem) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiNvmSubsystem) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiNvmSubsystem) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiNvmSubsystem) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiNvmSubsystem) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiNvmSubsystem) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiNvmSubsystem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiNvmSubsystem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiNvmSubsystem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiNvmSubsystem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiNvmSubsystem) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiNvmSubsystem) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


