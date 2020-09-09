# HyperflexClusterStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskPartitionCleanup** | Pointer to **bool** | If enabled, formats existing disk partitions (destroys all user data). | [optional] 
**LogicalAvalabilityZoneConfig** | Pointer to [**HyperflexLogicalAvailabilityZone**](hyperflex.LogicalAvailabilityZone.md) |  | [optional] 
**VdiOptimization** | Pointer to **bool** | Enable or disable VDI optimization (hybrid HyperFlex systems only). | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterStoragePolicy

`func NewHyperflexClusterStoragePolicy() *HyperflexClusterStoragePolicy`

NewHyperflexClusterStoragePolicy instantiates a new HyperflexClusterStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterStoragePolicyWithDefaults

`func NewHyperflexClusterStoragePolicyWithDefaults() *HyperflexClusterStoragePolicy`

NewHyperflexClusterStoragePolicyWithDefaults instantiates a new HyperflexClusterStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicy) GetDiskPartitionCleanup() bool`

GetDiskPartitionCleanup returns the DiskPartitionCleanup field if non-nil, zero value otherwise.

### GetDiskPartitionCleanupOk

`func (o *HyperflexClusterStoragePolicy) GetDiskPartitionCleanupOk() (*bool, bool)`

GetDiskPartitionCleanupOk returns a tuple with the DiskPartitionCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicy) SetDiskPartitionCleanup(v bool)`

SetDiskPartitionCleanup sets DiskPartitionCleanup field to given value.

### HasDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicy) HasDiskPartitionCleanup() bool`

HasDiskPartitionCleanup returns a boolean if a field has been set.

### GetLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicy) GetLogicalAvalabilityZoneConfig() HyperflexLogicalAvailabilityZone`

GetLogicalAvalabilityZoneConfig returns the LogicalAvalabilityZoneConfig field if non-nil, zero value otherwise.

### GetLogicalAvalabilityZoneConfigOk

`func (o *HyperflexClusterStoragePolicy) GetLogicalAvalabilityZoneConfigOk() (*HyperflexLogicalAvailabilityZone, bool)`

GetLogicalAvalabilityZoneConfigOk returns a tuple with the LogicalAvalabilityZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicy) SetLogicalAvalabilityZoneConfig(v HyperflexLogicalAvailabilityZone)`

SetLogicalAvalabilityZoneConfig sets LogicalAvalabilityZoneConfig field to given value.

### HasLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicy) HasLogicalAvalabilityZoneConfig() bool`

HasLogicalAvalabilityZoneConfig returns a boolean if a field has been set.

### GetVdiOptimization

`func (o *HyperflexClusterStoragePolicy) GetVdiOptimization() bool`

GetVdiOptimization returns the VdiOptimization field if non-nil, zero value otherwise.

### GetVdiOptimizationOk

`func (o *HyperflexClusterStoragePolicy) GetVdiOptimizationOk() (*bool, bool)`

GetVdiOptimizationOk returns a tuple with the VdiOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiOptimization

`func (o *HyperflexClusterStoragePolicy) SetVdiOptimization(v bool)`

SetVdiOptimization sets VdiOptimization field to given value.

### HasVdiOptimization

`func (o *HyperflexClusterStoragePolicy) HasVdiOptimization() bool`

HasVdiOptimization returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexClusterStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterStoragePolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterStoragePolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterStoragePolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterStoragePolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterStoragePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterStoragePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


