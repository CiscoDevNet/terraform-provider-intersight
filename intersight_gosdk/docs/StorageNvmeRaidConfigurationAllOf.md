# StorageNvmeRaidConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeRaidConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeRaidConfiguration"]
**ControllerDn** | Pointer to **string** | The storage controller Dn Name for which Nvme RAID is created at endpoint. | [optional] [readonly] 
**ControllerMoid** | Pointer to **string** | The storage controller Moid for which Nvme RAID creation is supported. | [optional] [readonly] 
**DiskStates** | Pointer to [**[]StorageNvmePhysicalDiskState**](StorageNvmePhysicalDiskState.md) |  | [optional] 
**DriveGroups** | Pointer to [**[]StorageNvmeRaidDriveGroup**](StorageNvmeRaidDriveGroup.md) |  | [optional] 
**ServerProfile** | Pointer to [**ServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**StoragePolicy** | Pointer to [**StorageStoragePolicyRelationship**](StorageStoragePolicyRelationship.md) |  | [optional] 

## Methods

### NewStorageNvmeRaidConfigurationAllOf

`func NewStorageNvmeRaidConfigurationAllOf(classId string, objectType string, ) *StorageNvmeRaidConfigurationAllOf`

NewStorageNvmeRaidConfigurationAllOf instantiates a new StorageNvmeRaidConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeRaidConfigurationAllOfWithDefaults

`func NewStorageNvmeRaidConfigurationAllOfWithDefaults() *StorageNvmeRaidConfigurationAllOf`

NewStorageNvmeRaidConfigurationAllOfWithDefaults instantiates a new StorageNvmeRaidConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeRaidConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeRaidConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeRaidConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeRaidConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerDn

`func (o *StorageNvmeRaidConfigurationAllOf) GetControllerDn() string`

GetControllerDn returns the ControllerDn field if non-nil, zero value otherwise.

### GetControllerDnOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetControllerDnOk() (*string, bool)`

GetControllerDnOk returns a tuple with the ControllerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDn

`func (o *StorageNvmeRaidConfigurationAllOf) SetControllerDn(v string)`

SetControllerDn sets ControllerDn field to given value.

### HasControllerDn

`func (o *StorageNvmeRaidConfigurationAllOf) HasControllerDn() bool`

HasControllerDn returns a boolean if a field has been set.

### GetControllerMoid

`func (o *StorageNvmeRaidConfigurationAllOf) GetControllerMoid() string`

GetControllerMoid returns the ControllerMoid field if non-nil, zero value otherwise.

### GetControllerMoidOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetControllerMoidOk() (*string, bool)`

GetControllerMoidOk returns a tuple with the ControllerMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMoid

`func (o *StorageNvmeRaidConfigurationAllOf) SetControllerMoid(v string)`

SetControllerMoid sets ControllerMoid field to given value.

### HasControllerMoid

`func (o *StorageNvmeRaidConfigurationAllOf) HasControllerMoid() bool`

HasControllerMoid returns a boolean if a field has been set.

### GetDiskStates

`func (o *StorageNvmeRaidConfigurationAllOf) GetDiskStates() []StorageNvmePhysicalDiskState`

GetDiskStates returns the DiskStates field if non-nil, zero value otherwise.

### GetDiskStatesOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetDiskStatesOk() (*[]StorageNvmePhysicalDiskState, bool)`

GetDiskStatesOk returns a tuple with the DiskStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStates

`func (o *StorageNvmeRaidConfigurationAllOf) SetDiskStates(v []StorageNvmePhysicalDiskState)`

SetDiskStates sets DiskStates field to given value.

### HasDiskStates

`func (o *StorageNvmeRaidConfigurationAllOf) HasDiskStates() bool`

HasDiskStates returns a boolean if a field has been set.

### SetDiskStatesNil

`func (o *StorageNvmeRaidConfigurationAllOf) SetDiskStatesNil(b bool)`

 SetDiskStatesNil sets the value for DiskStates to be an explicit nil

### UnsetDiskStates
`func (o *StorageNvmeRaidConfigurationAllOf) UnsetDiskStates()`

UnsetDiskStates ensures that no value is present for DiskStates, not even an explicit nil
### GetDriveGroups

`func (o *StorageNvmeRaidConfigurationAllOf) GetDriveGroups() []StorageNvmeRaidDriveGroup`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetDriveGroupsOk() (*[]StorageNvmeRaidDriveGroup, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *StorageNvmeRaidConfigurationAllOf) SetDriveGroups(v []StorageNvmeRaidDriveGroup)`

SetDriveGroups sets DriveGroups field to given value.

### HasDriveGroups

`func (o *StorageNvmeRaidConfigurationAllOf) HasDriveGroups() bool`

HasDriveGroups returns a boolean if a field has been set.

### SetDriveGroupsNil

`func (o *StorageNvmeRaidConfigurationAllOf) SetDriveGroupsNil(b bool)`

 SetDriveGroupsNil sets the value for DriveGroups to be an explicit nil

### UnsetDriveGroups
`func (o *StorageNvmeRaidConfigurationAllOf) UnsetDriveGroups()`

UnsetDriveGroups ensures that no value is present for DriveGroups, not even an explicit nil
### GetServerProfile

`func (o *StorageNvmeRaidConfigurationAllOf) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *StorageNvmeRaidConfigurationAllOf) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *StorageNvmeRaidConfigurationAllOf) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *StorageNvmeRaidConfigurationAllOf) GetStoragePolicy() StorageStoragePolicyRelationship`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageNvmeRaidConfigurationAllOf) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageNvmeRaidConfigurationAllOf) SetStoragePolicy(v StorageStoragePolicyRelationship)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageNvmeRaidConfigurationAllOf) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


