# StorageNvmeRaidConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeRaidConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeRaidConfiguration"]
**ControllerDn** | Pointer to **string** | The storage controller Dn Name for which Nvme RAID is created at endpoint. | [optional] [readonly] 
**ControllerMoid** | Pointer to **string** | The storage controller Moid for which Nvme RAID creation is supported. | [optional] [readonly] 
**ControllerSeries** | Pointer to **string** | Describes series of the installed controller. This will be used in the activation step after reboot to calculate the different parameters w.r.t specific controller series. | [optional] [readonly] 
**DiskStates** | Pointer to [**[]StorageNvmePhysicalDiskState**](StorageNvmePhysicalDiskState.md) |  | [optional] 
**DriveGroups** | Pointer to [**[]StorageNvmeRaidDriveGroup**](StorageNvmeRaidDriveGroup.md) |  | [optional] 
**ServerProfile** | Pointer to [**NullableServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**StoragePolicy** | Pointer to [**NullableStorageStoragePolicyRelationship**](StorageStoragePolicyRelationship.md) |  | [optional] 

## Methods

### NewStorageNvmeRaidConfiguration

`func NewStorageNvmeRaidConfiguration(classId string, objectType string, ) *StorageNvmeRaidConfiguration`

NewStorageNvmeRaidConfiguration instantiates a new StorageNvmeRaidConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeRaidConfigurationWithDefaults

`func NewStorageNvmeRaidConfigurationWithDefaults() *StorageNvmeRaidConfiguration`

NewStorageNvmeRaidConfigurationWithDefaults instantiates a new StorageNvmeRaidConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeRaidConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeRaidConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeRaidConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeRaidConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeRaidConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeRaidConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerDn

`func (o *StorageNvmeRaidConfiguration) GetControllerDn() string`

GetControllerDn returns the ControllerDn field if non-nil, zero value otherwise.

### GetControllerDnOk

`func (o *StorageNvmeRaidConfiguration) GetControllerDnOk() (*string, bool)`

GetControllerDnOk returns a tuple with the ControllerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDn

`func (o *StorageNvmeRaidConfiguration) SetControllerDn(v string)`

SetControllerDn sets ControllerDn field to given value.

### HasControllerDn

`func (o *StorageNvmeRaidConfiguration) HasControllerDn() bool`

HasControllerDn returns a boolean if a field has been set.

### GetControllerMoid

`func (o *StorageNvmeRaidConfiguration) GetControllerMoid() string`

GetControllerMoid returns the ControllerMoid field if non-nil, zero value otherwise.

### GetControllerMoidOk

`func (o *StorageNvmeRaidConfiguration) GetControllerMoidOk() (*string, bool)`

GetControllerMoidOk returns a tuple with the ControllerMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMoid

`func (o *StorageNvmeRaidConfiguration) SetControllerMoid(v string)`

SetControllerMoid sets ControllerMoid field to given value.

### HasControllerMoid

`func (o *StorageNvmeRaidConfiguration) HasControllerMoid() bool`

HasControllerMoid returns a boolean if a field has been set.

### GetControllerSeries

`func (o *StorageNvmeRaidConfiguration) GetControllerSeries() string`

GetControllerSeries returns the ControllerSeries field if non-nil, zero value otherwise.

### GetControllerSeriesOk

`func (o *StorageNvmeRaidConfiguration) GetControllerSeriesOk() (*string, bool)`

GetControllerSeriesOk returns a tuple with the ControllerSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerSeries

`func (o *StorageNvmeRaidConfiguration) SetControllerSeries(v string)`

SetControllerSeries sets ControllerSeries field to given value.

### HasControllerSeries

`func (o *StorageNvmeRaidConfiguration) HasControllerSeries() bool`

HasControllerSeries returns a boolean if a field has been set.

### GetDiskStates

`func (o *StorageNvmeRaidConfiguration) GetDiskStates() []StorageNvmePhysicalDiskState`

GetDiskStates returns the DiskStates field if non-nil, zero value otherwise.

### GetDiskStatesOk

`func (o *StorageNvmeRaidConfiguration) GetDiskStatesOk() (*[]StorageNvmePhysicalDiskState, bool)`

GetDiskStatesOk returns a tuple with the DiskStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStates

`func (o *StorageNvmeRaidConfiguration) SetDiskStates(v []StorageNvmePhysicalDiskState)`

SetDiskStates sets DiskStates field to given value.

### HasDiskStates

`func (o *StorageNvmeRaidConfiguration) HasDiskStates() bool`

HasDiskStates returns a boolean if a field has been set.

### SetDiskStatesNil

`func (o *StorageNvmeRaidConfiguration) SetDiskStatesNil(b bool)`

 SetDiskStatesNil sets the value for DiskStates to be an explicit nil

### UnsetDiskStates
`func (o *StorageNvmeRaidConfiguration) UnsetDiskStates()`

UnsetDiskStates ensures that no value is present for DiskStates, not even an explicit nil
### GetDriveGroups

`func (o *StorageNvmeRaidConfiguration) GetDriveGroups() []StorageNvmeRaidDriveGroup`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *StorageNvmeRaidConfiguration) GetDriveGroupsOk() (*[]StorageNvmeRaidDriveGroup, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *StorageNvmeRaidConfiguration) SetDriveGroups(v []StorageNvmeRaidDriveGroup)`

SetDriveGroups sets DriveGroups field to given value.

### HasDriveGroups

`func (o *StorageNvmeRaidConfiguration) HasDriveGroups() bool`

HasDriveGroups returns a boolean if a field has been set.

### SetDriveGroupsNil

`func (o *StorageNvmeRaidConfiguration) SetDriveGroupsNil(b bool)`

 SetDriveGroupsNil sets the value for DriveGroups to be an explicit nil

### UnsetDriveGroups
`func (o *StorageNvmeRaidConfiguration) UnsetDriveGroups()`

UnsetDriveGroups ensures that no value is present for DriveGroups, not even an explicit nil
### GetServerProfile

`func (o *StorageNvmeRaidConfiguration) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *StorageNvmeRaidConfiguration) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *StorageNvmeRaidConfiguration) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *StorageNvmeRaidConfiguration) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.

### SetServerProfileNil

`func (o *StorageNvmeRaidConfiguration) SetServerProfileNil(b bool)`

 SetServerProfileNil sets the value for ServerProfile to be an explicit nil

### UnsetServerProfile
`func (o *StorageNvmeRaidConfiguration) UnsetServerProfile()`

UnsetServerProfile ensures that no value is present for ServerProfile, not even an explicit nil
### GetStoragePolicy

`func (o *StorageNvmeRaidConfiguration) GetStoragePolicy() StorageStoragePolicyRelationship`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageNvmeRaidConfiguration) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageNvmeRaidConfiguration) SetStoragePolicy(v StorageStoragePolicyRelationship)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageNvmeRaidConfiguration) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### SetStoragePolicyNil

`func (o *StorageNvmeRaidConfiguration) SetStoragePolicyNil(b bool)`

 SetStoragePolicyNil sets the value for StoragePolicy to be an explicit nil

### UnsetStoragePolicy
`func (o *StorageNvmeRaidConfiguration) UnsetStoragePolicy()`

UnsetStoragePolicy ensures that no value is present for StoragePolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


