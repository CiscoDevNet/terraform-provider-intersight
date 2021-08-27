# HyperflexNodeConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NodeConfigPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NodeConfigPolicy"]
**DataIpRange** | Pointer to [**NullableHyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**HxdpIpRange** | Pointer to [**NullableHyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**HypervisorControlIpRange** | Pointer to [**NullableHyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**MgmtIpRange** | Pointer to [**NullableHyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**NodeNamePrefix** | Pointer to **string** | The node name prefix that is used to automatically generate the default hostname for each server. A dash (-) will be appended to the prefix followed by the node number to form a hostname. This default naming scheme can be manually overridden in the node configuration. The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must start with an alphanumeric character. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeConfigPolicy

`func NewHyperflexNodeConfigPolicy(classId string, objectType string, ) *HyperflexNodeConfigPolicy`

NewHyperflexNodeConfigPolicy instantiates a new HyperflexNodeConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeConfigPolicyWithDefaults

`func NewHyperflexNodeConfigPolicyWithDefaults() *HyperflexNodeConfigPolicy`

NewHyperflexNodeConfigPolicyWithDefaults instantiates a new HyperflexNodeConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNodeConfigPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeConfigPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeConfigPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNodeConfigPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeConfigPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeConfigPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetDataIpRangeNil

`func (o *HyperflexNodeConfigPolicy) SetDataIpRangeNil(b bool)`

 SetDataIpRangeNil sets the value for DataIpRange to be an explicit nil

### UnsetDataIpRange
`func (o *HyperflexNodeConfigPolicy) UnsetDataIpRange()`

UnsetDataIpRange ensures that no value is present for DataIpRange, not even an explicit nil
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

### SetHxdpIpRangeNil

`func (o *HyperflexNodeConfigPolicy) SetHxdpIpRangeNil(b bool)`

 SetHxdpIpRangeNil sets the value for HxdpIpRange to be an explicit nil

### UnsetHxdpIpRange
`func (o *HyperflexNodeConfigPolicy) UnsetHxdpIpRange()`

UnsetHxdpIpRange ensures that no value is present for HxdpIpRange, not even an explicit nil
### GetHypervisorControlIpRange

`func (o *HyperflexNodeConfigPolicy) GetHypervisorControlIpRange() HyperflexIpAddrRange`

GetHypervisorControlIpRange returns the HypervisorControlIpRange field if non-nil, zero value otherwise.

### GetHypervisorControlIpRangeOk

`func (o *HyperflexNodeConfigPolicy) GetHypervisorControlIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetHypervisorControlIpRangeOk returns a tuple with the HypervisorControlIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorControlIpRange

`func (o *HyperflexNodeConfigPolicy) SetHypervisorControlIpRange(v HyperflexIpAddrRange)`

SetHypervisorControlIpRange sets HypervisorControlIpRange field to given value.

### HasHypervisorControlIpRange

`func (o *HyperflexNodeConfigPolicy) HasHypervisorControlIpRange() bool`

HasHypervisorControlIpRange returns a boolean if a field has been set.

### SetHypervisorControlIpRangeNil

`func (o *HyperflexNodeConfigPolicy) SetHypervisorControlIpRangeNil(b bool)`

 SetHypervisorControlIpRangeNil sets the value for HypervisorControlIpRange to be an explicit nil

### UnsetHypervisorControlIpRange
`func (o *HyperflexNodeConfigPolicy) UnsetHypervisorControlIpRange()`

UnsetHypervisorControlIpRange ensures that no value is present for HypervisorControlIpRange, not even an explicit nil
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

### SetMgmtIpRangeNil

`func (o *HyperflexNodeConfigPolicy) SetMgmtIpRangeNil(b bool)`

 SetMgmtIpRangeNil sets the value for MgmtIpRange to be an explicit nil

### UnsetMgmtIpRange
`func (o *HyperflexNodeConfigPolicy) UnsetMgmtIpRange()`

UnsetMgmtIpRange ensures that no value is present for MgmtIpRange, not even an explicit nil
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


