# StorageStoragePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.StoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.StoragePolicy"]
**DefaultDriveMode** | Pointer to **string** | All unconfigured drives will move to the selected state on deployment. Newly inserted drives will move to the selected state. Select Unconfigured Good option to retain the existing configuration. Select JBOD to move the unconfigured drives to JBOD state. Select RAID0 to create a RAID0 virtual drive on each of the unconfigured drives. If JBOD is selected, unconfigured drives will move to JBOD state on host reboot. This setting is applicable only to selected set of controllers on FI attached servers. * &#x60;UnconfiguredGood&#x60; - Newly inserted drives or on reboot, drives will remain the same state. * &#x60;Jbod&#x60; - Newly inserted drives or on reboot, drives will automatically move to JBOD state if drive state was UnconfiguredGood. * &#x60;RAID0&#x60; - Newly inserted drives or on reboot, virtual drives will be created, respective drives will move to Online state. | [optional] [default to "UnconfiguredGood"]
**DirectAttachedNvmeSlots** | Pointer to **string** | Only U.3 NVMe drives has to be specified, entered slots will be moved to Direct attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range. | [optional] 
**GlobalHotSpares** | Pointer to **string** | A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen. | [optional] 
**M2VirtualDrive** | Pointer to [**NullableStorageM2VirtualDriveConfig**](StorageM2VirtualDriveConfig.md) |  | [optional] 
**Raid0Drive** | Pointer to [**NullableStorageR0Drive**](StorageR0Drive.md) |  | [optional] 
**RaidAttachedNvmeSlots** | Pointer to **string** | Only U.3 NVMe drives has to be specified, entered slots will be moved to RAID attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range. | [optional] 
**SecureJbods** | Pointer to **string** | JBOD drives specified in this slot range will be encrypted. Allowed values are &#39;ALL&#39;, or a comma or hyphen separated number range. Sample format is ALL or 1, 3 or 4-6, 8. Setting the value to &#39;ALL&#39; will encrypt all the unused UnconfigureGood/JBOD disks. | [optional] 
**UnusedDisksState** | Pointer to **string** | State to which drives, not used in this policy, are to be moved. NoChange will not change the drive state. No Change must be selected if Default Drive State is set to JBOD or RAID0. * &#x60;NoChange&#x60; - Drive state will not be modified by Storage Policy. * &#x60;UnconfiguredGood&#x60; - Unconfigured good state -ready to be added in a RAID group. * &#x60;Jbod&#x60; - JBOD state where the disks start showing up to Host OS. | [optional] [default to "NoChange"]
**UseJbodForVdCreation** | Pointer to **bool** | Disks in JBOD State are used to create virtual drives. This setting must be disabled if Default Drive State is set to JBOD. | [optional] 
**DriveGroup** | Pointer to [**[]StorageDriveGroupRelationship**](StorageDriveGroupRelationship.md) | An array of relationships to storageDriveGroup resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewStorageStoragePolicyAllOf

`func NewStorageStoragePolicyAllOf(classId string, objectType string, ) *StorageStoragePolicyAllOf`

NewStorageStoragePolicyAllOf instantiates a new StorageStoragePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStoragePolicyAllOfWithDefaults

`func NewStorageStoragePolicyAllOfWithDefaults() *StorageStoragePolicyAllOf`

NewStorageStoragePolicyAllOfWithDefaults instantiates a new StorageStoragePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageStoragePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageStoragePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageStoragePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageStoragePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageStoragePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageStoragePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultDriveMode

`func (o *StorageStoragePolicyAllOf) GetDefaultDriveMode() string`

GetDefaultDriveMode returns the DefaultDriveMode field if non-nil, zero value otherwise.

### GetDefaultDriveModeOk

`func (o *StorageStoragePolicyAllOf) GetDefaultDriveModeOk() (*string, bool)`

GetDefaultDriveModeOk returns a tuple with the DefaultDriveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDriveMode

`func (o *StorageStoragePolicyAllOf) SetDefaultDriveMode(v string)`

SetDefaultDriveMode sets DefaultDriveMode field to given value.

### HasDefaultDriveMode

`func (o *StorageStoragePolicyAllOf) HasDefaultDriveMode() bool`

HasDefaultDriveMode returns a boolean if a field has been set.

### GetDirectAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) GetDirectAttachedNvmeSlots() string`

GetDirectAttachedNvmeSlots returns the DirectAttachedNvmeSlots field if non-nil, zero value otherwise.

### GetDirectAttachedNvmeSlotsOk

`func (o *StorageStoragePolicyAllOf) GetDirectAttachedNvmeSlotsOk() (*string, bool)`

GetDirectAttachedNvmeSlotsOk returns a tuple with the DirectAttachedNvmeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) SetDirectAttachedNvmeSlots(v string)`

SetDirectAttachedNvmeSlots sets DirectAttachedNvmeSlots field to given value.

### HasDirectAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) HasDirectAttachedNvmeSlots() bool`

HasDirectAttachedNvmeSlots returns a boolean if a field has been set.

### GetGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) GetGlobalHotSpares() string`

GetGlobalHotSpares returns the GlobalHotSpares field if non-nil, zero value otherwise.

### GetGlobalHotSparesOk

`func (o *StorageStoragePolicyAllOf) GetGlobalHotSparesOk() (*string, bool)`

GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) SetGlobalHotSpares(v string)`

SetGlobalHotSpares sets GlobalHotSpares field to given value.

### HasGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) HasGlobalHotSpares() bool`

HasGlobalHotSpares returns a boolean if a field has been set.

### GetM2VirtualDrive

`func (o *StorageStoragePolicyAllOf) GetM2VirtualDrive() StorageM2VirtualDriveConfig`

GetM2VirtualDrive returns the M2VirtualDrive field if non-nil, zero value otherwise.

### GetM2VirtualDriveOk

`func (o *StorageStoragePolicyAllOf) GetM2VirtualDriveOk() (*StorageM2VirtualDriveConfig, bool)`

GetM2VirtualDriveOk returns a tuple with the M2VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2VirtualDrive

`func (o *StorageStoragePolicyAllOf) SetM2VirtualDrive(v StorageM2VirtualDriveConfig)`

SetM2VirtualDrive sets M2VirtualDrive field to given value.

### HasM2VirtualDrive

`func (o *StorageStoragePolicyAllOf) HasM2VirtualDrive() bool`

HasM2VirtualDrive returns a boolean if a field has been set.

### SetM2VirtualDriveNil

`func (o *StorageStoragePolicyAllOf) SetM2VirtualDriveNil(b bool)`

 SetM2VirtualDriveNil sets the value for M2VirtualDrive to be an explicit nil

### UnsetM2VirtualDrive
`func (o *StorageStoragePolicyAllOf) UnsetM2VirtualDrive()`

UnsetM2VirtualDrive ensures that no value is present for M2VirtualDrive, not even an explicit nil
### GetRaid0Drive

`func (o *StorageStoragePolicyAllOf) GetRaid0Drive() StorageR0Drive`

GetRaid0Drive returns the Raid0Drive field if non-nil, zero value otherwise.

### GetRaid0DriveOk

`func (o *StorageStoragePolicyAllOf) GetRaid0DriveOk() (*StorageR0Drive, bool)`

GetRaid0DriveOk returns a tuple with the Raid0Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid0Drive

`func (o *StorageStoragePolicyAllOf) SetRaid0Drive(v StorageR0Drive)`

SetRaid0Drive sets Raid0Drive field to given value.

### HasRaid0Drive

`func (o *StorageStoragePolicyAllOf) HasRaid0Drive() bool`

HasRaid0Drive returns a boolean if a field has been set.

### SetRaid0DriveNil

`func (o *StorageStoragePolicyAllOf) SetRaid0DriveNil(b bool)`

 SetRaid0DriveNil sets the value for Raid0Drive to be an explicit nil

### UnsetRaid0Drive
`func (o *StorageStoragePolicyAllOf) UnsetRaid0Drive()`

UnsetRaid0Drive ensures that no value is present for Raid0Drive, not even an explicit nil
### GetRaidAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) GetRaidAttachedNvmeSlots() string`

GetRaidAttachedNvmeSlots returns the RaidAttachedNvmeSlots field if non-nil, zero value otherwise.

### GetRaidAttachedNvmeSlotsOk

`func (o *StorageStoragePolicyAllOf) GetRaidAttachedNvmeSlotsOk() (*string, bool)`

GetRaidAttachedNvmeSlotsOk returns a tuple with the RaidAttachedNvmeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) SetRaidAttachedNvmeSlots(v string)`

SetRaidAttachedNvmeSlots sets RaidAttachedNvmeSlots field to given value.

### HasRaidAttachedNvmeSlots

`func (o *StorageStoragePolicyAllOf) HasRaidAttachedNvmeSlots() bool`

HasRaidAttachedNvmeSlots returns a boolean if a field has been set.

### GetSecureJbods

`func (o *StorageStoragePolicyAllOf) GetSecureJbods() string`

GetSecureJbods returns the SecureJbods field if non-nil, zero value otherwise.

### GetSecureJbodsOk

`func (o *StorageStoragePolicyAllOf) GetSecureJbodsOk() (*string, bool)`

GetSecureJbodsOk returns a tuple with the SecureJbods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureJbods

`func (o *StorageStoragePolicyAllOf) SetSecureJbods(v string)`

SetSecureJbods sets SecureJbods field to given value.

### HasSecureJbods

`func (o *StorageStoragePolicyAllOf) HasSecureJbods() bool`

HasSecureJbods returns a boolean if a field has been set.

### GetUnusedDisksState

`func (o *StorageStoragePolicyAllOf) GetUnusedDisksState() string`

GetUnusedDisksState returns the UnusedDisksState field if non-nil, zero value otherwise.

### GetUnusedDisksStateOk

`func (o *StorageStoragePolicyAllOf) GetUnusedDisksStateOk() (*string, bool)`

GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedDisksState

`func (o *StorageStoragePolicyAllOf) SetUnusedDisksState(v string)`

SetUnusedDisksState sets UnusedDisksState field to given value.

### HasUnusedDisksState

`func (o *StorageStoragePolicyAllOf) HasUnusedDisksState() bool`

HasUnusedDisksState returns a boolean if a field has been set.

### GetUseJbodForVdCreation

`func (o *StorageStoragePolicyAllOf) GetUseJbodForVdCreation() bool`

GetUseJbodForVdCreation returns the UseJbodForVdCreation field if non-nil, zero value otherwise.

### GetUseJbodForVdCreationOk

`func (o *StorageStoragePolicyAllOf) GetUseJbodForVdCreationOk() (*bool, bool)`

GetUseJbodForVdCreationOk returns a tuple with the UseJbodForVdCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJbodForVdCreation

`func (o *StorageStoragePolicyAllOf) SetUseJbodForVdCreation(v bool)`

SetUseJbodForVdCreation sets UseJbodForVdCreation field to given value.

### HasUseJbodForVdCreation

`func (o *StorageStoragePolicyAllOf) HasUseJbodForVdCreation() bool`

HasUseJbodForVdCreation returns a boolean if a field has been set.

### GetDriveGroup

`func (o *StorageStoragePolicyAllOf) GetDriveGroup() []StorageDriveGroupRelationship`

GetDriveGroup returns the DriveGroup field if non-nil, zero value otherwise.

### GetDriveGroupOk

`func (o *StorageStoragePolicyAllOf) GetDriveGroupOk() (*[]StorageDriveGroupRelationship, bool)`

GetDriveGroupOk returns a tuple with the DriveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroup

`func (o *StorageStoragePolicyAllOf) SetDriveGroup(v []StorageDriveGroupRelationship)`

SetDriveGroup sets DriveGroup field to given value.

### HasDriveGroup

`func (o *StorageStoragePolicyAllOf) HasDriveGroup() bool`

HasDriveGroup returns a boolean if a field has been set.

### SetDriveGroupNil

`func (o *StorageStoragePolicyAllOf) SetDriveGroupNil(b bool)`

 SetDriveGroupNil sets the value for DriveGroup to be an explicit nil

### UnsetDriveGroup
`func (o *StorageStoragePolicyAllOf) UnsetDriveGroup()`

UnsetDriveGroup ensures that no value is present for DriveGroup, not even an explicit nil
### GetOrganization

`func (o *StorageStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StorageStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StorageStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StorageStoragePolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *StorageStoragePolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *StorageStoragePolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *StorageStoragePolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *StorageStoragePolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *StorageStoragePolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *StorageStoragePolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


