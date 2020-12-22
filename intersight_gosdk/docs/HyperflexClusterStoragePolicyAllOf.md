# HyperflexClusterStoragePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterStoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterStoragePolicy"]
**DiskPartitionCleanup** | Pointer to **bool** | If enabled, formats existing disk partitions (destroys all user data). | [optional] 
**LogicalAvalabilityZoneConfig** | Pointer to [**NullableHyperflexLogicalAvailabilityZone**](hyperflex.LogicalAvailabilityZone.md) |  | [optional] 
**VdiOptimization** | Pointer to **bool** | Enable or disable VDI optimization (hybrid HyperFlex systems only). | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterStoragePolicyAllOf

`func NewHyperflexClusterStoragePolicyAllOf(classId string, objectType string, ) *HyperflexClusterStoragePolicyAllOf`

NewHyperflexClusterStoragePolicyAllOf instantiates a new HyperflexClusterStoragePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterStoragePolicyAllOfWithDefaults

`func NewHyperflexClusterStoragePolicyAllOfWithDefaults() *HyperflexClusterStoragePolicyAllOf`

NewHyperflexClusterStoragePolicyAllOfWithDefaults instantiates a new HyperflexClusterStoragePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterStoragePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterStoragePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterStoragePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterStoragePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicyAllOf) GetDiskPartitionCleanup() bool`

GetDiskPartitionCleanup returns the DiskPartitionCleanup field if non-nil, zero value otherwise.

### GetDiskPartitionCleanupOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetDiskPartitionCleanupOk() (*bool, bool)`

GetDiskPartitionCleanupOk returns a tuple with the DiskPartitionCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicyAllOf) SetDiskPartitionCleanup(v bool)`

SetDiskPartitionCleanup sets DiskPartitionCleanup field to given value.

### HasDiskPartitionCleanup

`func (o *HyperflexClusterStoragePolicyAllOf) HasDiskPartitionCleanup() bool`

HasDiskPartitionCleanup returns a boolean if a field has been set.

### GetLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicyAllOf) GetLogicalAvalabilityZoneConfig() HyperflexLogicalAvailabilityZone`

GetLogicalAvalabilityZoneConfig returns the LogicalAvalabilityZoneConfig field if non-nil, zero value otherwise.

### GetLogicalAvalabilityZoneConfigOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetLogicalAvalabilityZoneConfigOk() (*HyperflexLogicalAvailabilityZone, bool)`

GetLogicalAvalabilityZoneConfigOk returns a tuple with the LogicalAvalabilityZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicyAllOf) SetLogicalAvalabilityZoneConfig(v HyperflexLogicalAvailabilityZone)`

SetLogicalAvalabilityZoneConfig sets LogicalAvalabilityZoneConfig field to given value.

### HasLogicalAvalabilityZoneConfig

`func (o *HyperflexClusterStoragePolicyAllOf) HasLogicalAvalabilityZoneConfig() bool`

HasLogicalAvalabilityZoneConfig returns a boolean if a field has been set.

### SetLogicalAvalabilityZoneConfigNil

`func (o *HyperflexClusterStoragePolicyAllOf) SetLogicalAvalabilityZoneConfigNil(b bool)`

 SetLogicalAvalabilityZoneConfigNil sets the value for LogicalAvalabilityZoneConfig to be an explicit nil

### UnsetLogicalAvalabilityZoneConfig
`func (o *HyperflexClusterStoragePolicyAllOf) UnsetLogicalAvalabilityZoneConfig()`

UnsetLogicalAvalabilityZoneConfig ensures that no value is present for LogicalAvalabilityZoneConfig, not even an explicit nil
### GetVdiOptimization

`func (o *HyperflexClusterStoragePolicyAllOf) GetVdiOptimization() bool`

GetVdiOptimization returns the VdiOptimization field if non-nil, zero value otherwise.

### GetVdiOptimizationOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetVdiOptimizationOk() (*bool, bool)`

GetVdiOptimizationOk returns a tuple with the VdiOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiOptimization

`func (o *HyperflexClusterStoragePolicyAllOf) SetVdiOptimization(v bool)`

SetVdiOptimization sets VdiOptimization field to given value.

### HasVdiOptimization

`func (o *HyperflexClusterStoragePolicyAllOf) HasVdiOptimization() bool`

HasVdiOptimization returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexClusterStoragePolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterStoragePolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterStoragePolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterStoragePolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterStoragePolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterStoragePolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


