# KubernetesNodeGroupProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeGroupProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeGroupProfile"]
**Currentsize** | Pointer to **int64** | Current number of nodes in this node group at any given point in time. | [optional] [readonly] 
**Desiredsize** | Pointer to **int64** | Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler. | [optional] [default to 3]
**Labels** | Pointer to [**[]KubernetesNodeGroupLabel**](KubernetesNodeGroupLabel.md) |  | [optional] 
**Maxsize** | Pointer to **int64** | Maximum number of nodes desired in this node group. | [optional] 
**Minsize** | Pointer to **int64** | Minimum number of nodes desired in this node group. | [optional] 
**NodeType** | Pointer to **string** | The node type ControlPlane, Worker or ControlPlaneWorker. * &#x60;Worker&#x60; - Node will be marked as a worker node. * &#x60;ControlPlane&#x60; - Node will be marked as a control plane node. * &#x60;ControlPlaneWorker&#x60; - Node will be both a controle plane and a worker. | [optional] [default to "Worker"]
**Taints** | Pointer to [**[]KubernetesNodeGroupTaint**](KubernetesNodeGroupTaint.md) |  | [optional] 
**ClusterProfile** | Pointer to [**KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) |  | [optional] 
**InfraProvider** | Pointer to [**KubernetesBaseInfrastructureProviderRelationship**](KubernetesBaseInfrastructureProviderRelationship.md) |  | [optional] 
**IpPools** | Pointer to [**[]IppoolPoolRelationship**](IppoolPoolRelationship.md) | An array of relationships to ippoolPool resources. | [optional] 
**KubernetesVersion** | Pointer to [**KubernetesVersionPolicyRelationship**](KubernetesVersionPolicyRelationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]KubernetesNodeProfileRelationship**](KubernetesNodeProfileRelationship.md) | An array of relationships to kubernetesNodeProfile resources. | [optional] 

## Methods

### NewKubernetesNodeGroupProfileAllOf

`func NewKubernetesNodeGroupProfileAllOf(classId string, objectType string, ) *KubernetesNodeGroupProfileAllOf`

NewKubernetesNodeGroupProfileAllOf instantiates a new KubernetesNodeGroupProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeGroupProfileAllOfWithDefaults

`func NewKubernetesNodeGroupProfileAllOfWithDefaults() *KubernetesNodeGroupProfileAllOf`

NewKubernetesNodeGroupProfileAllOfWithDefaults instantiates a new KubernetesNodeGroupProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeGroupProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeGroupProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeGroupProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeGroupProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeGroupProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentsize

`func (o *KubernetesNodeGroupProfileAllOf) GetCurrentsize() int64`

GetCurrentsize returns the Currentsize field if non-nil, zero value otherwise.

### GetCurrentsizeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetCurrentsizeOk() (*int64, bool)`

GetCurrentsizeOk returns a tuple with the Currentsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentsize

`func (o *KubernetesNodeGroupProfileAllOf) SetCurrentsize(v int64)`

SetCurrentsize sets Currentsize field to given value.

### HasCurrentsize

`func (o *KubernetesNodeGroupProfileAllOf) HasCurrentsize() bool`

HasCurrentsize returns a boolean if a field has been set.

### GetDesiredsize

`func (o *KubernetesNodeGroupProfileAllOf) GetDesiredsize() int64`

GetDesiredsize returns the Desiredsize field if non-nil, zero value otherwise.

### GetDesiredsizeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetDesiredsizeOk() (*int64, bool)`

GetDesiredsizeOk returns a tuple with the Desiredsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredsize

`func (o *KubernetesNodeGroupProfileAllOf) SetDesiredsize(v int64)`

SetDesiredsize sets Desiredsize field to given value.

### HasDesiredsize

`func (o *KubernetesNodeGroupProfileAllOf) HasDesiredsize() bool`

HasDesiredsize returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesNodeGroupProfileAllOf) GetLabels() []KubernetesNodeGroupLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNodeGroupProfileAllOf) GetLabelsOk() (*[]KubernetesNodeGroupLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNodeGroupProfileAllOf) SetLabels(v []KubernetesNodeGroupLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesNodeGroupProfileAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesNodeGroupProfileAllOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesNodeGroupProfileAllOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMaxsize

`func (o *KubernetesNodeGroupProfileAllOf) GetMaxsize() int64`

GetMaxsize returns the Maxsize field if non-nil, zero value otherwise.

### GetMaxsizeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetMaxsizeOk() (*int64, bool)`

GetMaxsizeOk returns a tuple with the Maxsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsize

`func (o *KubernetesNodeGroupProfileAllOf) SetMaxsize(v int64)`

SetMaxsize sets Maxsize field to given value.

### HasMaxsize

`func (o *KubernetesNodeGroupProfileAllOf) HasMaxsize() bool`

HasMaxsize returns a boolean if a field has been set.

### GetMinsize

`func (o *KubernetesNodeGroupProfileAllOf) GetMinsize() int64`

GetMinsize returns the Minsize field if non-nil, zero value otherwise.

### GetMinsizeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetMinsizeOk() (*int64, bool)`

GetMinsizeOk returns a tuple with the Minsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsize

`func (o *KubernetesNodeGroupProfileAllOf) SetMinsize(v int64)`

SetMinsize sets Minsize field to given value.

### HasMinsize

`func (o *KubernetesNodeGroupProfileAllOf) HasMinsize() bool`

HasMinsize returns a boolean if a field has been set.

### GetNodeType

`func (o *KubernetesNodeGroupProfileAllOf) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *KubernetesNodeGroupProfileAllOf) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *KubernetesNodeGroupProfileAllOf) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *KubernetesNodeGroupProfileAllOf) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetTaints

`func (o *KubernetesNodeGroupProfileAllOf) GetTaints() []KubernetesNodeGroupTaint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *KubernetesNodeGroupProfileAllOf) GetTaintsOk() (*[]KubernetesNodeGroupTaint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *KubernetesNodeGroupProfileAllOf) SetTaints(v []KubernetesNodeGroupTaint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *KubernetesNodeGroupProfileAllOf) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *KubernetesNodeGroupProfileAllOf) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *KubernetesNodeGroupProfileAllOf) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil
### GetClusterProfile

`func (o *KubernetesNodeGroupProfileAllOf) GetClusterProfile() KubernetesClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *KubernetesNodeGroupProfileAllOf) GetClusterProfileOk() (*KubernetesClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *KubernetesNodeGroupProfileAllOf) SetClusterProfile(v KubernetesClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *KubernetesNodeGroupProfileAllOf) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### GetInfraProvider

`func (o *KubernetesNodeGroupProfileAllOf) GetInfraProvider() KubernetesBaseInfrastructureProviderRelationship`

GetInfraProvider returns the InfraProvider field if non-nil, zero value otherwise.

### GetInfraProviderOk

`func (o *KubernetesNodeGroupProfileAllOf) GetInfraProviderOk() (*KubernetesBaseInfrastructureProviderRelationship, bool)`

GetInfraProviderOk returns a tuple with the InfraProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraProvider

`func (o *KubernetesNodeGroupProfileAllOf) SetInfraProvider(v KubernetesBaseInfrastructureProviderRelationship)`

SetInfraProvider sets InfraProvider field to given value.

### HasInfraProvider

`func (o *KubernetesNodeGroupProfileAllOf) HasInfraProvider() bool`

HasInfraProvider returns a boolean if a field has been set.

### GetIpPools

`func (o *KubernetesNodeGroupProfileAllOf) GetIpPools() []IppoolPoolRelationship`

GetIpPools returns the IpPools field if non-nil, zero value otherwise.

### GetIpPoolsOk

`func (o *KubernetesNodeGroupProfileAllOf) GetIpPoolsOk() (*[]IppoolPoolRelationship, bool)`

GetIpPoolsOk returns a tuple with the IpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPools

`func (o *KubernetesNodeGroupProfileAllOf) SetIpPools(v []IppoolPoolRelationship)`

SetIpPools sets IpPools field to given value.

### HasIpPools

`func (o *KubernetesNodeGroupProfileAllOf) HasIpPools() bool`

HasIpPools returns a boolean if a field has been set.

### SetIpPoolsNil

`func (o *KubernetesNodeGroupProfileAllOf) SetIpPoolsNil(b bool)`

 SetIpPoolsNil sets the value for IpPools to be an explicit nil

### UnsetIpPools
`func (o *KubernetesNodeGroupProfileAllOf) UnsetIpPools()`

UnsetIpPools ensures that no value is present for IpPools, not even an explicit nil
### GetKubernetesVersion

`func (o *KubernetesNodeGroupProfileAllOf) GetKubernetesVersion() KubernetesVersionPolicyRelationship`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesNodeGroupProfileAllOf) GetKubernetesVersionOk() (*KubernetesVersionPolicyRelationship, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesNodeGroupProfileAllOf) SetKubernetesVersion(v KubernetesVersionPolicyRelationship)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesNodeGroupProfileAllOf) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesNodeGroupProfileAllOf) GetNodes() []KubernetesNodeProfileRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesNodeGroupProfileAllOf) GetNodesOk() (*[]KubernetesNodeProfileRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesNodeGroupProfileAllOf) SetNodes(v []KubernetesNodeProfileRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesNodeGroupProfileAllOf) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *KubernetesNodeGroupProfileAllOf) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *KubernetesNodeGroupProfileAllOf) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


