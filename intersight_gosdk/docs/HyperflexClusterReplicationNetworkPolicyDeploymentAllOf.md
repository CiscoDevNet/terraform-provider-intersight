# HyperflexClusterReplicationNetworkPolicyDeploymentAllOf

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

### NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOf

`func NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOf(classId string, objectType string, ) *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf`

NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOf instantiates a new HyperflexClusterReplicationNetworkPolicyDeploymentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOfWithDefaults

`func NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOfWithDefaults() *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf`

NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOfWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicyDeploymentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetName

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetPolicyMoid() string`

GetPolicyMoid returns the PolicyMoid field if non-nil, zero value otherwise.

### GetPolicyMoidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetPolicyMoidOk() (*string, bool)`

GetPolicyMoidOk returns a tuple with the PolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetPolicyMoid(v string)`

SetPolicyMoid sets PolicyMoid field to given value.

### HasPolicyMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasPolicyMoid() bool`

HasPolicyMoid returns a boolean if a field has been set.

### GetProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetProfileMoid() string`

GetProfileMoid returns the ProfileMoid field if non-nil, zero value otherwise.

### GetProfileMoidOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetProfileMoidOk() (*string, bool)`

GetProfileMoidOk returns a tuple with the ProfileMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetProfileMoid(v string)`

SetProfileMoid sets ProfileMoid field to given value.

### HasProfileMoid

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasProfileMoid() bool`

HasProfileMoid returns a boolean if a field has been set.

### GetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationBandwidthMbps() int64`

GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field if non-nil, zero value otherwise.

### GetReplicationBandwidthMbpsOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationBandwidthMbpsOk() (*int64, bool)`

GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationBandwidthMbps(v int64)`

SetReplicationBandwidthMbps sets ReplicationBandwidthMbps field to given value.

### HasReplicationBandwidthMbps

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationBandwidthMbps() bool`

HasReplicationBandwidthMbps returns a boolean if a field has been set.

### GetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationIpranges() []HyperflexIpAddrRange`

GetReplicationIpranges returns the ReplicationIpranges field if non-nil, zero value otherwise.

### GetReplicationIprangesOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationIprangesOk() (*[]HyperflexIpAddrRange, bool)`

GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationIpranges(v []HyperflexIpAddrRange)`

SetReplicationIpranges sets ReplicationIpranges field to given value.

### HasReplicationIpranges

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationIpranges() bool`

HasReplicationIpranges returns a boolean if a field has been set.

### SetReplicationIprangesNil

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationIprangesNil(b bool)`

 SetReplicationIprangesNil sets the value for ReplicationIpranges to be an explicit nil

### UnsetReplicationIpranges
`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) UnsetReplicationIpranges()`

UnsetReplicationIpranges ensures that no value is present for ReplicationIpranges, not even an explicit nil
### GetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationMtu() int64`

GetReplicationMtu returns the ReplicationMtu field if non-nil, zero value otherwise.

### GetReplicationMtuOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationMtuOk() (*int64, bool)`

GetReplicationMtuOk returns a tuple with the ReplicationMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationMtu(v int64)`

SetReplicationMtu sets ReplicationMtu field to given value.

### HasReplicationMtu

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationMtu() bool`

HasReplicationMtu returns a boolean if a field has been set.

### GetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationVlan() HyperflexNamedVlan`

GetReplicationVlan returns the ReplicationVlan field if non-nil, zero value otherwise.

### GetReplicationVlanOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationVlanOk() (*HyperflexNamedVlan, bool)`

GetReplicationVlanOk returns a tuple with the ReplicationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationVlan(v HyperflexNamedVlan)`

SetReplicationVlan sets ReplicationVlan field to given value.

### HasReplicationVlan

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationVlan() bool`

HasReplicationVlan returns a boolean if a field has been set.

### SetReplicationVlanNil

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationVlanNil(b bool)`

 SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil

### UnsetReplicationVlan
`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) UnsetReplicationVlan()`

UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
### GetRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


