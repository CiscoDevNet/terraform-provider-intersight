# HyperflexClusterReplicationNetworkPolicyDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterReplicationNetworkPolicyDeployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterReplicationNetworkPolicyDeployment"]
**ClusterUuid** | Pointer to **string** | Uuid of the HyperFlex cluster. | [optional] [readonly] 
**Description** | Pointer to **string** | Description from corresponding ClusterReplicationNetworkPolicy. | [optional] [readonly] 
**Discovered** | Pointer to **bool** | True if record created by discovery on HyperFlex cluster. | [optional] 
**Name** | Pointer to **string** | Name from corresponding ClusterReplicationNetworkPolicy. | [optional] [readonly] 
**PolicyMoid** | Pointer to **string** | Deployed network policy moid. | [optional] [readonly] 
**ProfileMoid** | Pointer to **string** | Deployed cluster profile moid. | [optional] [readonly] 
**ReplicationBandwidthMbps** | Pointer to **int64** | Bandwidth for the Replication network in Mbps. | [optional] [readonly] [default to 0]
**ReplicationIpranges** | Pointer to [**[]HyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**ReplicationMtu** | Pointer to **int64** | MTU for the Replication network. | [optional] [readonly] [default to 1500]
**ReplicationVlan** | Pointer to [**NullableHyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique request ID allowing retry of the same logical request following a transient communication failure. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterReplicationNetworkPolicyDeployment

`func NewHyperflexClusterReplicationNetworkPolicyDeployment(classId string, objectType string, ) *HyperflexClusterReplicationNetworkPolicyDeployment`

NewHyperflexClusterReplicationNetworkPolicyDeployment instantiates a new HyperflexClusterReplicationNetworkPolicyDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterReplicationNetworkPolicyDeploymentWithDefaults

`func NewHyperflexClusterReplicationNetworkPolicyDeploymentWithDefaults() *HyperflexClusterReplicationNetworkPolicyDeployment`

NewHyperflexClusterReplicationNetworkPolicyDeploymentWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicyDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetName

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetPolicyMoid() string`

GetPolicyMoid returns the PolicyMoid field if non-nil, zero value otherwise.

### GetPolicyMoidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetPolicyMoidOk() (*string, bool)`

GetPolicyMoidOk returns a tuple with the PolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetPolicyMoid(v string)`

SetPolicyMoid sets PolicyMoid field to given value.

### HasPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasPolicyMoid() bool`

HasPolicyMoid returns a boolean if a field has been set.

### GetProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetProfileMoid() string`

GetProfileMoid returns the ProfileMoid field if non-nil, zero value otherwise.

### GetProfileMoidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetProfileMoidOk() (*string, bool)`

GetProfileMoidOk returns a tuple with the ProfileMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetProfileMoid(v string)`

SetProfileMoid sets ProfileMoid field to given value.

### HasProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasProfileMoid() bool`

HasProfileMoid returns a boolean if a field has been set.

### GetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationBandwidthMbps() int64`

GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field if non-nil, zero value otherwise.

### GetReplicationBandwidthMbpsOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationBandwidthMbpsOk() (*int64, bool)`

GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationBandwidthMbps(v int64)`

SetReplicationBandwidthMbps sets ReplicationBandwidthMbps field to given value.

### HasReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasReplicationBandwidthMbps() bool`

HasReplicationBandwidthMbps returns a boolean if a field has been set.

### GetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationIpranges() []HyperflexIpAddrRange`

GetReplicationIpranges returns the ReplicationIpranges field if non-nil, zero value otherwise.

### GetReplicationIprangesOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationIprangesOk() (*[]HyperflexIpAddrRange, bool)`

GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationIpranges(v []HyperflexIpAddrRange)`

SetReplicationIpranges sets ReplicationIpranges field to given value.

### HasReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasReplicationIpranges() bool`

HasReplicationIpranges returns a boolean if a field has been set.

### SetReplicationIprangesNil

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationIprangesNil(b bool)`

 SetReplicationIprangesNil sets the value for ReplicationIpranges to be an explicit nil

### UnsetReplicationIpranges
`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) UnsetReplicationIpranges()`

UnsetReplicationIpranges ensures that no value is present for ReplicationIpranges, not even an explicit nil
### GetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationMtu() int64`

GetReplicationMtu returns the ReplicationMtu field if non-nil, zero value otherwise.

### GetReplicationMtuOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationMtuOk() (*int64, bool)`

GetReplicationMtuOk returns a tuple with the ReplicationMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationMtu(v int64)`

SetReplicationMtu sets ReplicationMtu field to given value.

### HasReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasReplicationMtu() bool`

HasReplicationMtu returns a boolean if a field has been set.

### GetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationVlan() HyperflexNamedVlan`

GetReplicationVlan returns the ReplicationVlan field if non-nil, zero value otherwise.

### GetReplicationVlanOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetReplicationVlanOk() (*HyperflexNamedVlan, bool)`

GetReplicationVlanOk returns a tuple with the ReplicationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationVlan(v HyperflexNamedVlan)`

SetReplicationVlan sets ReplicationVlan field to given value.

### HasReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasReplicationVlan() bool`

HasReplicationVlan returns a boolean if a field has been set.

### SetReplicationVlanNil

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetReplicationVlanNil(b bool)`

 SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil

### UnsetReplicationVlan
`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) UnsetReplicationVlan()`

UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
### GetRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeployment) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


