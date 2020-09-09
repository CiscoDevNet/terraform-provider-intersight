# HyperflexUcsmConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KvmIpRange** | Pointer to [**HyperflexIpAddrRange**](hyperflex.IpAddrRange.md) |  | [optional] 
**MacPrefixRange** | Pointer to [**HyperflexMacAddrPrefixRange**](hyperflex.MacAddrPrefixRange.md) |  | [optional] 
**ServerFirmwareVersion** | Pointer to **string** | The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexUcsmConfigPolicy

`func NewHyperflexUcsmConfigPolicy() *HyperflexUcsmConfigPolicy`

NewHyperflexUcsmConfigPolicy instantiates a new HyperflexUcsmConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexUcsmConfigPolicyWithDefaults

`func NewHyperflexUcsmConfigPolicyWithDefaults() *HyperflexUcsmConfigPolicy`

NewHyperflexUcsmConfigPolicyWithDefaults instantiates a new HyperflexUcsmConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKvmIpRange

`func (o *HyperflexUcsmConfigPolicy) GetKvmIpRange() HyperflexIpAddrRange`

GetKvmIpRange returns the KvmIpRange field if non-nil, zero value otherwise.

### GetKvmIpRangeOk

`func (o *HyperflexUcsmConfigPolicy) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetKvmIpRangeOk returns a tuple with the KvmIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpRange

`func (o *HyperflexUcsmConfigPolicy) SetKvmIpRange(v HyperflexIpAddrRange)`

SetKvmIpRange sets KvmIpRange field to given value.

### HasKvmIpRange

`func (o *HyperflexUcsmConfigPolicy) HasKvmIpRange() bool`

HasKvmIpRange returns a boolean if a field has been set.

### GetMacPrefixRange

`func (o *HyperflexUcsmConfigPolicy) GetMacPrefixRange() HyperflexMacAddrPrefixRange`

GetMacPrefixRange returns the MacPrefixRange field if non-nil, zero value otherwise.

### GetMacPrefixRangeOk

`func (o *HyperflexUcsmConfigPolicy) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool)`

GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPrefixRange

`func (o *HyperflexUcsmConfigPolicy) SetMacPrefixRange(v HyperflexMacAddrPrefixRange)`

SetMacPrefixRange sets MacPrefixRange field to given value.

### HasMacPrefixRange

`func (o *HyperflexUcsmConfigPolicy) HasMacPrefixRange() bool`

HasMacPrefixRange returns a boolean if a field has been set.

### GetServerFirmwareVersion

`func (o *HyperflexUcsmConfigPolicy) GetServerFirmwareVersion() string`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexUcsmConfigPolicy) GetServerFirmwareVersionOk() (*string, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexUcsmConfigPolicy) SetServerFirmwareVersion(v string)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexUcsmConfigPolicy) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexUcsmConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexUcsmConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexUcsmConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexUcsmConfigPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexUcsmConfigPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexUcsmConfigPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexUcsmConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexUcsmConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexUcsmConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexUcsmConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


