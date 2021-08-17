# StorageStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.StoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.StoragePolicy"]
**GlobalHotSpares** | Pointer to **string** | A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen. | [optional] 
**M2VirtualDrive** | Pointer to [**NullableStorageM2VirtualDriveConfig**](storage.M2VirtualDriveConfig.md) |  | [optional] 
**Raid0Drive** | Pointer to [**NullableStorageR0Drive**](storage.R0Drive.md) |  | [optional] 
**UnusedDisksState** | Pointer to **string** | State to which disks, not used in this policy, are to be moved. NoChange will not change the drive state. * &#x60;NoChange&#x60; - Drive state will not be modified by Storage Policy. * &#x60;UnconfiguredGood&#x60; - Unconfigured good state -ready to be added in a RAID group. * &#x60;Jbod&#x60; - JBOD state where the disks start showing up to Host OS. | [optional] [default to "NoChange"]
**UseJbodForVdCreation** | Pointer to **bool** | Disks in JBOD State are used to create virtual drives. | [optional] 
**DriveGroup** | Pointer to [**[]StorageDriveGroupRelationship**](StorageDriveGroupRelationship.md) | An array of relationships to storageDriveGroup resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewStorageStoragePolicy

`func NewStorageStoragePolicy(classId string, objectType string, ) *StorageStoragePolicy`

NewStorageStoragePolicy instantiates a new StorageStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStoragePolicyWithDefaults

`func NewStorageStoragePolicyWithDefaults() *StorageStoragePolicy`

NewStorageStoragePolicyWithDefaults instantiates a new StorageStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageStoragePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageStoragePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageStoragePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageStoragePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageStoragePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageStoragePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGlobalHotSpares

`func (o *StorageStoragePolicy) GetGlobalHotSpares() string`

GetGlobalHotSpares returns the GlobalHotSpares field if non-nil, zero value otherwise.

### GetGlobalHotSparesOk

`func (o *StorageStoragePolicy) GetGlobalHotSparesOk() (*string, bool)`

GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHotSpares

`func (o *StorageStoragePolicy) SetGlobalHotSpares(v string)`

SetGlobalHotSpares sets GlobalHotSpares field to given value.

### HasGlobalHotSpares

`func (o *StorageStoragePolicy) HasGlobalHotSpares() bool`

HasGlobalHotSpares returns a boolean if a field has been set.

### GetM2VirtualDrive

`func (o *StorageStoragePolicy) GetM2VirtualDrive() StorageM2VirtualDriveConfig`

GetM2VirtualDrive returns the M2VirtualDrive field if non-nil, zero value otherwise.

### GetM2VirtualDriveOk

`func (o *StorageStoragePolicy) GetM2VirtualDriveOk() (*StorageM2VirtualDriveConfig, bool)`

GetM2VirtualDriveOk returns a tuple with the M2VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2VirtualDrive

`func (o *StorageStoragePolicy) SetM2VirtualDrive(v StorageM2VirtualDriveConfig)`

SetM2VirtualDrive sets M2VirtualDrive field to given value.

### HasM2VirtualDrive

`func (o *StorageStoragePolicy) HasM2VirtualDrive() bool`

HasM2VirtualDrive returns a boolean if a field has been set.

### SetM2VirtualDriveNil

`func (o *StorageStoragePolicy) SetM2VirtualDriveNil(b bool)`

 SetM2VirtualDriveNil sets the value for M2VirtualDrive to be an explicit nil

### UnsetM2VirtualDrive
`func (o *StorageStoragePolicy) UnsetM2VirtualDrive()`

UnsetM2VirtualDrive ensures that no value is present for M2VirtualDrive, not even an explicit nil
### GetRaid0Drive

`func (o *StorageStoragePolicy) GetRaid0Drive() StorageR0Drive`

GetRaid0Drive returns the Raid0Drive field if non-nil, zero value otherwise.

### GetRaid0DriveOk

`func (o *StorageStoragePolicy) GetRaid0DriveOk() (*StorageR0Drive, bool)`

GetRaid0DriveOk returns a tuple with the Raid0Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid0Drive

`func (o *StorageStoragePolicy) SetRaid0Drive(v StorageR0Drive)`

SetRaid0Drive sets Raid0Drive field to given value.

### HasRaid0Drive

`func (o *StorageStoragePolicy) HasRaid0Drive() bool`

HasRaid0Drive returns a boolean if a field has been set.

### SetRaid0DriveNil

`func (o *StorageStoragePolicy) SetRaid0DriveNil(b bool)`

 SetRaid0DriveNil sets the value for Raid0Drive to be an explicit nil

### UnsetRaid0Drive
`func (o *StorageStoragePolicy) UnsetRaid0Drive()`

UnsetRaid0Drive ensures that no value is present for Raid0Drive, not even an explicit nil
### GetUnusedDisksState

`func (o *StorageStoragePolicy) GetUnusedDisksState() string`

GetUnusedDisksState returns the UnusedDisksState field if non-nil, zero value otherwise.

### GetUnusedDisksStateOk

`func (o *StorageStoragePolicy) GetUnusedDisksStateOk() (*string, bool)`

GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedDisksState

`func (o *StorageStoragePolicy) SetUnusedDisksState(v string)`

SetUnusedDisksState sets UnusedDisksState field to given value.

### HasUnusedDisksState

`func (o *StorageStoragePolicy) HasUnusedDisksState() bool`

HasUnusedDisksState returns a boolean if a field has been set.

### GetUseJbodForVdCreation

`func (o *StorageStoragePolicy) GetUseJbodForVdCreation() bool`

GetUseJbodForVdCreation returns the UseJbodForVdCreation field if non-nil, zero value otherwise.

### GetUseJbodForVdCreationOk

`func (o *StorageStoragePolicy) GetUseJbodForVdCreationOk() (*bool, bool)`

GetUseJbodForVdCreationOk returns a tuple with the UseJbodForVdCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJbodForVdCreation

`func (o *StorageStoragePolicy) SetUseJbodForVdCreation(v bool)`

SetUseJbodForVdCreation sets UseJbodForVdCreation field to given value.

### HasUseJbodForVdCreation

`func (o *StorageStoragePolicy) HasUseJbodForVdCreation() bool`

HasUseJbodForVdCreation returns a boolean if a field has been set.

### GetDriveGroup

`func (o *StorageStoragePolicy) GetDriveGroup() []StorageDriveGroupRelationship`

GetDriveGroup returns the DriveGroup field if non-nil, zero value otherwise.

### GetDriveGroupOk

`func (o *StorageStoragePolicy) GetDriveGroupOk() (*[]StorageDriveGroupRelationship, bool)`

GetDriveGroupOk returns a tuple with the DriveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroup

`func (o *StorageStoragePolicy) SetDriveGroup(v []StorageDriveGroupRelationship)`

SetDriveGroup sets DriveGroup field to given value.

### HasDriveGroup

`func (o *StorageStoragePolicy) HasDriveGroup() bool`

HasDriveGroup returns a boolean if a field has been set.

### SetDriveGroupNil

`func (o *StorageStoragePolicy) SetDriveGroupNil(b bool)`

 SetDriveGroupNil sets the value for DriveGroup to be an explicit nil

### UnsetDriveGroup
`func (o *StorageStoragePolicy) UnsetDriveGroup()`

UnsetDriveGroup ensures that no value is present for DriveGroup, not even an explicit nil
### GetOrganization

`func (o *StorageStoragePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StorageStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StorageStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StorageStoragePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *StorageStoragePolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *StorageStoragePolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *StorageStoragePolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *StorageStoragePolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *StorageStoragePolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *StorageStoragePolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


