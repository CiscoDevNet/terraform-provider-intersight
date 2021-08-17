# StorageDriveGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.DriveGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.DriveGroup"]
**AutomaticDriveGroup** | Pointer to [**NullableStorageAutomaticDriveGroup**](storage.AutomaticDriveGroup.md) |  | [optional] 
**ManualDriveGroup** | Pointer to [**NullableStorageManualDriveGroup**](storage.ManualDriveGroup.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the drive group. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. | [optional] 
**RaidLevel** | Pointer to **string** | The supported RAID level for the disk group. * &#x60;Raid0&#x60; - RAID 0 Stripe Raid Level. * &#x60;Raid1&#x60; - RAID 1 Mirror Raid Level. * &#x60;Raid5&#x60; - RAID 5 Mirror Raid Level. * &#x60;Raid6&#x60; - RAID 6 Mirror Raid Level. * &#x60;Raid10&#x60; - RAID 10 Mirror Raid Level. * &#x60;Raid50&#x60; - RAID 50 Mirror Raid Level. * &#x60;Raid60&#x60; - RAID 60 Mirror Raid Level. | [optional] [default to "Raid0"]
**Type** | Pointer to **int32** | Type of drive selection to be used for this drive group. * &#x60;0&#x60; - Drives are selected manually by the user. * &#x60;1&#x60; - Drives are selected automatically based on the RAID and virtual drive configuration. | [optional] [readonly] [default to 0]
**VirtualDrives** | Pointer to [**[]StorageVirtualDriveConfiguration**](StorageVirtualDriveConfiguration.md) |  | [optional] 
**StoragePolicy** | Pointer to [**StorageStoragePolicyRelationship**](storage.StoragePolicy.Relationship.md) |  | [optional] 

## Methods

### NewStorageDriveGroup

`func NewStorageDriveGroup(classId string, objectType string, ) *StorageDriveGroup`

NewStorageDriveGroup instantiates a new StorageDriveGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDriveGroupWithDefaults

`func NewStorageDriveGroupWithDefaults() *StorageDriveGroup`

NewStorageDriveGroupWithDefaults instantiates a new StorageDriveGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageDriveGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageDriveGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageDriveGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageDriveGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageDriveGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageDriveGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutomaticDriveGroup

`func (o *StorageDriveGroup) GetAutomaticDriveGroup() StorageAutomaticDriveGroup`

GetAutomaticDriveGroup returns the AutomaticDriveGroup field if non-nil, zero value otherwise.

### GetAutomaticDriveGroupOk

`func (o *StorageDriveGroup) GetAutomaticDriveGroupOk() (*StorageAutomaticDriveGroup, bool)`

GetAutomaticDriveGroupOk returns a tuple with the AutomaticDriveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDriveGroup

`func (o *StorageDriveGroup) SetAutomaticDriveGroup(v StorageAutomaticDriveGroup)`

SetAutomaticDriveGroup sets AutomaticDriveGroup field to given value.

### HasAutomaticDriveGroup

`func (o *StorageDriveGroup) HasAutomaticDriveGroup() bool`

HasAutomaticDriveGroup returns a boolean if a field has been set.

### SetAutomaticDriveGroupNil

`func (o *StorageDriveGroup) SetAutomaticDriveGroupNil(b bool)`

 SetAutomaticDriveGroupNil sets the value for AutomaticDriveGroup to be an explicit nil

### UnsetAutomaticDriveGroup
`func (o *StorageDriveGroup) UnsetAutomaticDriveGroup()`

UnsetAutomaticDriveGroup ensures that no value is present for AutomaticDriveGroup, not even an explicit nil
### GetManualDriveGroup

`func (o *StorageDriveGroup) GetManualDriveGroup() StorageManualDriveGroup`

GetManualDriveGroup returns the ManualDriveGroup field if non-nil, zero value otherwise.

### GetManualDriveGroupOk

`func (o *StorageDriveGroup) GetManualDriveGroupOk() (*StorageManualDriveGroup, bool)`

GetManualDriveGroupOk returns a tuple with the ManualDriveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualDriveGroup

`func (o *StorageDriveGroup) SetManualDriveGroup(v StorageManualDriveGroup)`

SetManualDriveGroup sets ManualDriveGroup field to given value.

### HasManualDriveGroup

`func (o *StorageDriveGroup) HasManualDriveGroup() bool`

HasManualDriveGroup returns a boolean if a field has been set.

### SetManualDriveGroupNil

`func (o *StorageDriveGroup) SetManualDriveGroupNil(b bool)`

 SetManualDriveGroupNil sets the value for ManualDriveGroup to be an explicit nil

### UnsetManualDriveGroup
`func (o *StorageDriveGroup) UnsetManualDriveGroup()`

UnsetManualDriveGroup ensures that no value is present for ManualDriveGroup, not even an explicit nil
### GetName

`func (o *StorageDriveGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDriveGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDriveGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageDriveGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRaidLevel

`func (o *StorageDriveGroup) GetRaidLevel() string`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *StorageDriveGroup) GetRaidLevelOk() (*string, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *StorageDriveGroup) SetRaidLevel(v string)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *StorageDriveGroup) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### GetType

`func (o *StorageDriveGroup) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageDriveGroup) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageDriveGroup) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *StorageDriveGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDrives

`func (o *StorageDriveGroup) GetVirtualDrives() []StorageVirtualDriveConfiguration`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageDriveGroup) GetVirtualDrivesOk() (*[]StorageVirtualDriveConfiguration, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageDriveGroup) SetVirtualDrives(v []StorageVirtualDriveConfiguration)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageDriveGroup) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageDriveGroup) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageDriveGroup) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil
### GetStoragePolicy

`func (o *StorageDriveGroup) GetStoragePolicy() StorageStoragePolicyRelationship`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageDriveGroup) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageDriveGroup) SetStoragePolicy(v StorageStoragePolicyRelationship)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageDriveGroup) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


