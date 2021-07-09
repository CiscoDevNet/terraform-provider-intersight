# HyperflexClusterReplicationNetworkPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterReplicationNetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterReplicationNetworkPolicy"]
**ReplicationBandwidthMbps** | Pointer to **int64** | Bandwidth for the Replication network in Mbps. | [optional] [default to 0]
**ReplicationIpranges** | Pointer to [**[]HyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**ReplicationMtu** | Pointer to **int64** | MTU for the Replication network. | [optional] [default to 1500]
**ReplicationVlan** | Pointer to [**NullableHyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterReplicationNetworkPolicyAllOf

`func NewHyperflexClusterReplicationNetworkPolicyAllOf(classId string, objectType string, ) *HyperflexClusterReplicationNetworkPolicyAllOf`

NewHyperflexClusterReplicationNetworkPolicyAllOf instantiates a new HyperflexClusterReplicationNetworkPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterReplicationNetworkPolicyAllOfWithDefaults

`func NewHyperflexClusterReplicationNetworkPolicyAllOfWithDefaults() *HyperflexClusterReplicationNetworkPolicyAllOf`

NewHyperflexClusterReplicationNetworkPolicyAllOfWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationBandwidthMbps() int64`

GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field if non-nil, zero value otherwise.

### GetReplicationBandwidthMbpsOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationBandwidthMbpsOk() (*int64, bool)`

GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationBandwidthMbps(v int64)`

SetReplicationBandwidthMbps sets ReplicationBandwidthMbps field to given value.

### HasReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationBandwidthMbps() bool`

HasReplicationBandwidthMbps returns a boolean if a field has been set.

### GetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationIpranges() []HyperflexIpAddrRange`

GetReplicationIpranges returns the ReplicationIpranges field if non-nil, zero value otherwise.

### GetReplicationIprangesOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationIprangesOk() (*[]HyperflexIpAddrRange, bool)`

GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationIpranges(v []HyperflexIpAddrRange)`

SetReplicationIpranges sets ReplicationIpranges field to given value.

### HasReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationIpranges() bool`

HasReplicationIpranges returns a boolean if a field has been set.

### SetReplicationIprangesNil

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationIprangesNil(b bool)`

 SetReplicationIprangesNil sets the value for ReplicationIpranges to be an explicit nil

### UnsetReplicationIpranges
`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) UnsetReplicationIpranges()`

UnsetReplicationIpranges ensures that no value is present for ReplicationIpranges, not even an explicit nil
### GetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationMtu() int64`

GetReplicationMtu returns the ReplicationMtu field if non-nil, zero value otherwise.

### GetReplicationMtuOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationMtuOk() (*int64, bool)`

GetReplicationMtuOk returns a tuple with the ReplicationMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationMtu(v int64)`

SetReplicationMtu sets ReplicationMtu field to given value.

### HasReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationMtu() bool`

HasReplicationMtu returns a boolean if a field has been set.

### GetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationVlan() HyperflexNamedVlan`

GetReplicationVlan returns the ReplicationVlan field if non-nil, zero value otherwise.

### GetReplicationVlanOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetReplicationVlanOk() (*HyperflexNamedVlan, bool)`

GetReplicationVlanOk returns a tuple with the ReplicationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationVlan(v HyperflexNamedVlan)`

SetReplicationVlan sets ReplicationVlan field to given value.

### HasReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasReplicationVlan() bool`

HasReplicationVlan returns a boolean if a field has been set.

### SetReplicationVlanNil

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetReplicationVlanNil(b bool)`

 SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil

### UnsetReplicationVlan
`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) UnsetReplicationVlan()`

UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
### GetClusterProfiles

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


