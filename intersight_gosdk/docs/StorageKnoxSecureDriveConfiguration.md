# StorageKnoxSecureDriveConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.KnoxSecureDriveConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.KnoxSecureDriveConfiguration"]
**ControllerDn** | Pointer to **string** | The storage controller Dn Name for which RAID is created at endpoint. | [optional] [readonly] 
**ControllerMoid** | Pointer to **string** | The storage controller Moid for which RAID creation is supported. | [optional] [readonly] 
**DiskStates** | Pointer to [**[]StorageInternalMoPhysicalDiskConfig**](StorageInternalMoPhysicalDiskConfig.md) |  | [optional] 
**DriveGroups** | Pointer to [**[]StorageNvmeRaidDriveGroup**](StorageNvmeRaidDriveGroup.md) |  | [optional] 
**ServerProfile** | Pointer to [**NullableServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**StoragePolicy** | Pointer to [**NullableStorageStoragePolicyRelationship**](StorageStoragePolicyRelationship.md) |  | [optional] 

## Methods

### NewStorageKnoxSecureDriveConfiguration

`func NewStorageKnoxSecureDriveConfiguration(classId string, objectType string, ) *StorageKnoxSecureDriveConfiguration`

NewStorageKnoxSecureDriveConfiguration instantiates a new StorageKnoxSecureDriveConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKnoxSecureDriveConfigurationWithDefaults

`func NewStorageKnoxSecureDriveConfigurationWithDefaults() *StorageKnoxSecureDriveConfiguration`

NewStorageKnoxSecureDriveConfigurationWithDefaults instantiates a new StorageKnoxSecureDriveConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKnoxSecureDriveConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKnoxSecureDriveConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKnoxSecureDriveConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKnoxSecureDriveConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKnoxSecureDriveConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKnoxSecureDriveConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerDn

`func (o *StorageKnoxSecureDriveConfiguration) GetControllerDn() string`

GetControllerDn returns the ControllerDn field if non-nil, zero value otherwise.

### GetControllerDnOk

`func (o *StorageKnoxSecureDriveConfiguration) GetControllerDnOk() (*string, bool)`

GetControllerDnOk returns a tuple with the ControllerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDn

`func (o *StorageKnoxSecureDriveConfiguration) SetControllerDn(v string)`

SetControllerDn sets ControllerDn field to given value.

### HasControllerDn

`func (o *StorageKnoxSecureDriveConfiguration) HasControllerDn() bool`

HasControllerDn returns a boolean if a field has been set.

### GetControllerMoid

`func (o *StorageKnoxSecureDriveConfiguration) GetControllerMoid() string`

GetControllerMoid returns the ControllerMoid field if non-nil, zero value otherwise.

### GetControllerMoidOk

`func (o *StorageKnoxSecureDriveConfiguration) GetControllerMoidOk() (*string, bool)`

GetControllerMoidOk returns a tuple with the ControllerMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMoid

`func (o *StorageKnoxSecureDriveConfiguration) SetControllerMoid(v string)`

SetControllerMoid sets ControllerMoid field to given value.

### HasControllerMoid

`func (o *StorageKnoxSecureDriveConfiguration) HasControllerMoid() bool`

HasControllerMoid returns a boolean if a field has been set.

### GetDiskStates

`func (o *StorageKnoxSecureDriveConfiguration) GetDiskStates() []StorageInternalMoPhysicalDiskConfig`

GetDiskStates returns the DiskStates field if non-nil, zero value otherwise.

### GetDiskStatesOk

`func (o *StorageKnoxSecureDriveConfiguration) GetDiskStatesOk() (*[]StorageInternalMoPhysicalDiskConfig, bool)`

GetDiskStatesOk returns a tuple with the DiskStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStates

`func (o *StorageKnoxSecureDriveConfiguration) SetDiskStates(v []StorageInternalMoPhysicalDiskConfig)`

SetDiskStates sets DiskStates field to given value.

### HasDiskStates

`func (o *StorageKnoxSecureDriveConfiguration) HasDiskStates() bool`

HasDiskStates returns a boolean if a field has been set.

### SetDiskStatesNil

`func (o *StorageKnoxSecureDriveConfiguration) SetDiskStatesNil(b bool)`

 SetDiskStatesNil sets the value for DiskStates to be an explicit nil

### UnsetDiskStates
`func (o *StorageKnoxSecureDriveConfiguration) UnsetDiskStates()`

UnsetDiskStates ensures that no value is present for DiskStates, not even an explicit nil
### GetDriveGroups

`func (o *StorageKnoxSecureDriveConfiguration) GetDriveGroups() []StorageNvmeRaidDriveGroup`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *StorageKnoxSecureDriveConfiguration) GetDriveGroupsOk() (*[]StorageNvmeRaidDriveGroup, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *StorageKnoxSecureDriveConfiguration) SetDriveGroups(v []StorageNvmeRaidDriveGroup)`

SetDriveGroups sets DriveGroups field to given value.

### HasDriveGroups

`func (o *StorageKnoxSecureDriveConfiguration) HasDriveGroups() bool`

HasDriveGroups returns a boolean if a field has been set.

### SetDriveGroupsNil

`func (o *StorageKnoxSecureDriveConfiguration) SetDriveGroupsNil(b bool)`

 SetDriveGroupsNil sets the value for DriveGroups to be an explicit nil

### UnsetDriveGroups
`func (o *StorageKnoxSecureDriveConfiguration) UnsetDriveGroups()`

UnsetDriveGroups ensures that no value is present for DriveGroups, not even an explicit nil
### GetServerProfile

`func (o *StorageKnoxSecureDriveConfiguration) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *StorageKnoxSecureDriveConfiguration) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *StorageKnoxSecureDriveConfiguration) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *StorageKnoxSecureDriveConfiguration) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.

### SetServerProfileNil

`func (o *StorageKnoxSecureDriveConfiguration) SetServerProfileNil(b bool)`

 SetServerProfileNil sets the value for ServerProfile to be an explicit nil

### UnsetServerProfile
`func (o *StorageKnoxSecureDriveConfiguration) UnsetServerProfile()`

UnsetServerProfile ensures that no value is present for ServerProfile, not even an explicit nil
### GetStoragePolicy

`func (o *StorageKnoxSecureDriveConfiguration) GetStoragePolicy() StorageStoragePolicyRelationship`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageKnoxSecureDriveConfiguration) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageKnoxSecureDriveConfiguration) SetStoragePolicy(v StorageStoragePolicyRelationship)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageKnoxSecureDriveConfiguration) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### SetStoragePolicyNil

`func (o *StorageKnoxSecureDriveConfiguration) SetStoragePolicyNil(b bool)`

 SetStoragePolicyNil sets the value for StoragePolicy to be an explicit nil

### UnsetStoragePolicy
`func (o *StorageKnoxSecureDriveConfiguration) UnsetStoragePolicy()`

UnsetStoragePolicy ensures that no value is present for StoragePolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


