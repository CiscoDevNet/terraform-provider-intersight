# HyperflexNodeConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataIpRange** | Pointer to [**HyperflexIpAddrRange**](hyperflex.IpAddrRange.md) |  | [optional] 
**HxdpIpRange** | Pointer to [**HyperflexIpAddrRange**](hyperflex.IpAddrRange.md) |  | [optional] 
**MgmtIpRange** | Pointer to [**HyperflexIpAddrRange**](hyperflex.IpAddrRange.md) |  | [optional] 
**NodeNamePrefix** | Pointer to **string** | The node name prefix that is used to automatically generate the default hostname for each server. A dash (-) will be appended to the prefix followed by the node number to form a hostname. This default naming scheme can be manually overridden in the node configuration. The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must start with an alphanumeric character. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeConfigPolicy

`func NewHyperflexNodeConfigPolicy() *HyperflexNodeConfigPolicy`

NewHyperflexNodeConfigPolicy instantiates a new HyperflexNodeConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeConfigPolicyWithDefaults

`func NewHyperflexNodeConfigPolicyWithDefaults() *HyperflexNodeConfigPolicy`

NewHyperflexNodeConfigPolicyWithDefaults instantiates a new HyperflexNodeConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataIpRange

`func (o *HyperflexNodeConfigPolicy) GetDataIpRange() HyperflexIpAddrRange`

GetDataIpRange returns the DataIpRange field if non-nil, zero value otherwise.

### GetDataIpRangeOk

`func (o *HyperflexNodeConfigPolicy) GetDataIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetDataIpRangeOk returns a tuple with the DataIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIpRange

`func (o *HyperflexNodeConfigPolicy) SetDataIpRange(v HyperflexIpAddrRange)`

SetDataIpRange sets DataIpRange field to given value.

### HasDataIpRange

`func (o *HyperflexNodeConfigPolicy) HasDataIpRange() bool`

HasDataIpRange returns a boolean if a field has been set.

### GetHxdpIpRange

`func (o *HyperflexNodeConfigPolicy) GetHxdpIpRange() HyperflexIpAddrRange`

GetHxdpIpRange returns the HxdpIpRange field if non-nil, zero value otherwise.

### GetHxdpIpRangeOk

`func (o *HyperflexNodeConfigPolicy) GetHxdpIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetHxdpIpRangeOk returns a tuple with the HxdpIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpIpRange

`func (o *HyperflexNodeConfigPolicy) SetHxdpIpRange(v HyperflexIpAddrRange)`

SetHxdpIpRange sets HxdpIpRange field to given value.

### HasHxdpIpRange

`func (o *HyperflexNodeConfigPolicy) HasHxdpIpRange() bool`

HasHxdpIpRange returns a boolean if a field has been set.

### GetMgmtIpRange

`func (o *HyperflexNodeConfigPolicy) GetMgmtIpRange() HyperflexIpAddrRange`

GetMgmtIpRange returns the MgmtIpRange field if non-nil, zero value otherwise.

### GetMgmtIpRangeOk

`func (o *HyperflexNodeConfigPolicy) GetMgmtIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetMgmtIpRangeOk returns a tuple with the MgmtIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpRange

`func (o *HyperflexNodeConfigPolicy) SetMgmtIpRange(v HyperflexIpAddrRange)`

SetMgmtIpRange sets MgmtIpRange field to given value.

### HasMgmtIpRange

`func (o *HyperflexNodeConfigPolicy) HasMgmtIpRange() bool`

HasMgmtIpRange returns a boolean if a field has been set.

### GetNodeNamePrefix

`func (o *HyperflexNodeConfigPolicy) GetNodeNamePrefix() string`

GetNodeNamePrefix returns the NodeNamePrefix field if non-nil, zero value otherwise.

### GetNodeNamePrefixOk

`func (o *HyperflexNodeConfigPolicy) GetNodeNamePrefixOk() (*string, bool)`

GetNodeNamePrefixOk returns a tuple with the NodeNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNamePrefix

`func (o *HyperflexNodeConfigPolicy) SetNodeNamePrefix(v string)`

SetNodeNamePrefix sets NodeNamePrefix field to given value.

### HasNodeNamePrefix

`func (o *HyperflexNodeConfigPolicy) HasNodeNamePrefix() bool`

HasNodeNamePrefix returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexNodeConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexNodeConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexNodeConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexNodeConfigPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexNodeConfigPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexNodeConfigPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexNodeConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexNodeConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexNodeConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexNodeConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


