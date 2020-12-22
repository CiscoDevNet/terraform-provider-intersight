# StorageStoragePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.StoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.StoragePolicy"]
**GlobalHotSpares** | Pointer to [**[]StorageLocalDisk**](StorageLocalDisk.md) |  | [optional] 
**RetainPolicyVirtualDrives** | Pointer to **bool** | Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy. | [optional] [default to true]
**UnusedDisksState** | Pointer to **string** | Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved. * &#x60;UnconfiguredGood&#x60; - Unconfigured good state -ready to be added in a RAID group. * &#x60;Jbod&#x60; - JBOD state where the disks start showing up to host os. | [optional] [default to "UnconfiguredGood"]
**VirtualDrives** | Pointer to [**[]StorageVirtualDriveConfig**](StorageVirtualDriveConfig.md) |  | [optional] 
**DiskGroupPolicies** | Pointer to [**[]StorageDiskGroupPolicyRelationship**](StorageDiskGroupPolicyRelationship.md) | An array of relationships to storageDiskGroupPolicy resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
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


### GetGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) GetGlobalHotSpares() []StorageLocalDisk`

GetGlobalHotSpares returns the GlobalHotSpares field if non-nil, zero value otherwise.

### GetGlobalHotSparesOk

`func (o *StorageStoragePolicyAllOf) GetGlobalHotSparesOk() (*[]StorageLocalDisk, bool)`

GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) SetGlobalHotSpares(v []StorageLocalDisk)`

SetGlobalHotSpares sets GlobalHotSpares field to given value.

### HasGlobalHotSpares

`func (o *StorageStoragePolicyAllOf) HasGlobalHotSpares() bool`

HasGlobalHotSpares returns a boolean if a field has been set.

### SetGlobalHotSparesNil

`func (o *StorageStoragePolicyAllOf) SetGlobalHotSparesNil(b bool)`

 SetGlobalHotSparesNil sets the value for GlobalHotSpares to be an explicit nil

### UnsetGlobalHotSpares
`func (o *StorageStoragePolicyAllOf) UnsetGlobalHotSpares()`

UnsetGlobalHotSpares ensures that no value is present for GlobalHotSpares, not even an explicit nil
### GetRetainPolicyVirtualDrives

`func (o *StorageStoragePolicyAllOf) GetRetainPolicyVirtualDrives() bool`

GetRetainPolicyVirtualDrives returns the RetainPolicyVirtualDrives field if non-nil, zero value otherwise.

### GetRetainPolicyVirtualDrivesOk

`func (o *StorageStoragePolicyAllOf) GetRetainPolicyVirtualDrivesOk() (*bool, bool)`

GetRetainPolicyVirtualDrivesOk returns a tuple with the RetainPolicyVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPolicyVirtualDrives

`func (o *StorageStoragePolicyAllOf) SetRetainPolicyVirtualDrives(v bool)`

SetRetainPolicyVirtualDrives sets RetainPolicyVirtualDrives field to given value.

### HasRetainPolicyVirtualDrives

`func (o *StorageStoragePolicyAllOf) HasRetainPolicyVirtualDrives() bool`

HasRetainPolicyVirtualDrives returns a boolean if a field has been set.

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

### GetVirtualDrives

`func (o *StorageStoragePolicyAllOf) GetVirtualDrives() []StorageVirtualDriveConfig`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageStoragePolicyAllOf) GetVirtualDrivesOk() (*[]StorageVirtualDriveConfig, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageStoragePolicyAllOf) SetVirtualDrives(v []StorageVirtualDriveConfig)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageStoragePolicyAllOf) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageStoragePolicyAllOf) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageStoragePolicyAllOf) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil
### GetDiskGroupPolicies

`func (o *StorageStoragePolicyAllOf) GetDiskGroupPolicies() []StorageDiskGroupPolicyRelationship`

GetDiskGroupPolicies returns the DiskGroupPolicies field if non-nil, zero value otherwise.

### GetDiskGroupPoliciesOk

`func (o *StorageStoragePolicyAllOf) GetDiskGroupPoliciesOk() (*[]StorageDiskGroupPolicyRelationship, bool)`

GetDiskGroupPoliciesOk returns a tuple with the DiskGroupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroupPolicies

`func (o *StorageStoragePolicyAllOf) SetDiskGroupPolicies(v []StorageDiskGroupPolicyRelationship)`

SetDiskGroupPolicies sets DiskGroupPolicies field to given value.

### HasDiskGroupPolicies

`func (o *StorageStoragePolicyAllOf) HasDiskGroupPolicies() bool`

HasDiskGroupPolicies returns a boolean if a field has been set.

### SetDiskGroupPoliciesNil

`func (o *StorageStoragePolicyAllOf) SetDiskGroupPoliciesNil(b bool)`

 SetDiskGroupPoliciesNil sets the value for DiskGroupPolicies to be an explicit nil

### UnsetDiskGroupPolicies
`func (o *StorageStoragePolicyAllOf) UnsetDiskGroupPolicies()`

UnsetDiskGroupPolicies ensures that no value is present for DiskGroupPolicies, not even an explicit nil
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


