# StorageDiskGroupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedHotSpares** | Pointer to [**[]StorageLocalDisk**](storage.LocalDisk.md) |  | [optional] 
**RaidLevel** | Pointer to **string** | The supported RAID level for the disk group. * &#x60;Raid0&#x60; - RAID 0 Stripe Raid Level. * &#x60;Raid1&#x60; - RAID 1 Mirror Raid Level. * &#x60;Raid5&#x60; - RAID 5 Mirror Raid Level. * &#x60;Raid6&#x60; - RAID 6 Mirror Raid Level. * &#x60;Raid10&#x60; - RAID 10 Mirror Raid Level. * &#x60;Raid50&#x60; - RAID 50 Mirror Raid Level. * &#x60;Raid60&#x60; - RAID 60 Mirror Raid Level. | [optional] [default to "Raid0"]
**SpanGroups** | Pointer to [**[]StorageSpanGroup**](storage.SpanGroup.md) |  | [optional] 
**UseJbods** | Pointer to **bool** | Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**StoragePolicies** | Pointer to [**[]StorageStoragePolicyRelationship**](storage.StoragePolicy.Relationship.md) | An array of relationships to storageStoragePolicy resources. | [optional] 

## Methods

### NewStorageDiskGroupPolicy

`func NewStorageDiskGroupPolicy() *StorageDiskGroupPolicy`

NewStorageDiskGroupPolicy instantiates a new StorageDiskGroupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDiskGroupPolicyWithDefaults

`func NewStorageDiskGroupPolicyWithDefaults() *StorageDiskGroupPolicy`

NewStorageDiskGroupPolicyWithDefaults instantiates a new StorageDiskGroupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedHotSpares

`func (o *StorageDiskGroupPolicy) GetDedicatedHotSpares() []StorageLocalDisk`

GetDedicatedHotSpares returns the DedicatedHotSpares field if non-nil, zero value otherwise.

### GetDedicatedHotSparesOk

`func (o *StorageDiskGroupPolicy) GetDedicatedHotSparesOk() (*[]StorageLocalDisk, bool)`

GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSpares

`func (o *StorageDiskGroupPolicy) SetDedicatedHotSpares(v []StorageLocalDisk)`

SetDedicatedHotSpares sets DedicatedHotSpares field to given value.

### HasDedicatedHotSpares

`func (o *StorageDiskGroupPolicy) HasDedicatedHotSpares() bool`

HasDedicatedHotSpares returns a boolean if a field has been set.

### GetRaidLevel

`func (o *StorageDiskGroupPolicy) GetRaidLevel() string`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *StorageDiskGroupPolicy) GetRaidLevelOk() (*string, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *StorageDiskGroupPolicy) SetRaidLevel(v string)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *StorageDiskGroupPolicy) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### GetSpanGroups

`func (o *StorageDiskGroupPolicy) GetSpanGroups() []StorageSpanGroup`

GetSpanGroups returns the SpanGroups field if non-nil, zero value otherwise.

### GetSpanGroupsOk

`func (o *StorageDiskGroupPolicy) GetSpanGroupsOk() (*[]StorageSpanGroup, bool)`

GetSpanGroupsOk returns a tuple with the SpanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanGroups

`func (o *StorageDiskGroupPolicy) SetSpanGroups(v []StorageSpanGroup)`

SetSpanGroups sets SpanGroups field to given value.

### HasSpanGroups

`func (o *StorageDiskGroupPolicy) HasSpanGroups() bool`

HasSpanGroups returns a boolean if a field has been set.

### GetUseJbods

`func (o *StorageDiskGroupPolicy) GetUseJbods() bool`

GetUseJbods returns the UseJbods field if non-nil, zero value otherwise.

### GetUseJbodsOk

`func (o *StorageDiskGroupPolicy) GetUseJbodsOk() (*bool, bool)`

GetUseJbodsOk returns a tuple with the UseJbods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJbods

`func (o *StorageDiskGroupPolicy) SetUseJbods(v bool)`

SetUseJbods sets UseJbods field to given value.

### HasUseJbods

`func (o *StorageDiskGroupPolicy) HasUseJbods() bool`

HasUseJbods returns a boolean if a field has been set.

### GetOrganization

`func (o *StorageDiskGroupPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StorageDiskGroupPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StorageDiskGroupPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StorageDiskGroupPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetStoragePolicies

`func (o *StorageDiskGroupPolicy) GetStoragePolicies() []StorageStoragePolicyRelationship`

GetStoragePolicies returns the StoragePolicies field if non-nil, zero value otherwise.

### GetStoragePoliciesOk

`func (o *StorageDiskGroupPolicy) GetStoragePoliciesOk() (*[]StorageStoragePolicyRelationship, bool)`

GetStoragePoliciesOk returns a tuple with the StoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicies

`func (o *StorageDiskGroupPolicy) SetStoragePolicies(v []StorageStoragePolicyRelationship)`

SetStoragePolicies sets StoragePolicies field to given value.

### HasStoragePolicies

`func (o *StorageDiskGroupPolicy) HasStoragePolicies() bool`

HasStoragePolicies returns a boolean if a field has been set.

### SetStoragePoliciesNil

`func (o *StorageDiskGroupPolicy) SetStoragePoliciesNil(b bool)`

 SetStoragePoliciesNil sets the value for StoragePolicies to be an explicit nil

### UnsetStoragePolicies
`func (o *StorageDiskGroupPolicy) UnsetStoragePolicies()`

UnsetStoragePolicies ensures that no value is present for StoragePolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


