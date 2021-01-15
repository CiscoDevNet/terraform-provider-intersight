# IppoolIpLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpLease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpLease"]
**IpType** | Pointer to **string** | Type of the IP address requested. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Address** | Pointer to **string** | IPv4 Address given as a lease to an external entity like server profiles. | [optional] [readonly] 
**IpV4Config** | Pointer to [**NullableIppoolIpV4Config**](ippool.IpV4Config.md) |  | [optional] 
**IpV6Address** | Pointer to **string** | IPv6 Address given as a lease to an external entity like server profiles. | [optional] [readonly] 
**IpV6Config** | Pointer to [**NullableIppoolIpV6Config**](ippool.IpV6Config.md) |  | [optional] 
**Var0VirtualMachineNodeProfile** | Pointer to [**KubernetesVirtualMachineNodeProfileRelationship**](kubernetes.VirtualMachineNodeProfile.Relationship.md) |  | [optional] 
**Var1VirtualMachineNodeProfile** | Pointer to [**KubernetesVirtualMachineNodeProfileRelationship**](kubernetes.VirtualMachineNodeProfile.Relationship.md) |  | [optional] 
**Var2ClusterProfile** | Pointer to [**KubernetesClusterProfileRelationship**](kubernetes.ClusterProfile.Relationship.md) |  | [optional] 
**Var3ClusterProfile** | Pointer to [**KubernetesClusterProfileRelationship**](kubernetes.ClusterProfile.Relationship.md) |  | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolPoolRelationship**](ippool.Pool.Relationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IppoolPoolMemberRelationship**](ippool.PoolMember.Relationship.md) |  | [optional] 
**Universe** | Pointer to [**IppoolUniverseRelationship**](ippool.Universe.Relationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](vrf.Vrf.Relationship.md) |  | [optional] 

## Methods

### NewIppoolIpLeaseAllOf

`func NewIppoolIpLeaseAllOf(classId string, objectType string, ) *IppoolIpLeaseAllOf`

NewIppoolIpLeaseAllOf instantiates a new IppoolIpLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpLeaseAllOfWithDefaults

`func NewIppoolIpLeaseAllOfWithDefaults() *IppoolIpLeaseAllOf`

NewIppoolIpLeaseAllOfWithDefaults instantiates a new IppoolIpLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolIpLeaseAllOf) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolIpLeaseAllOf) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolIpLeaseAllOf) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolIpLeaseAllOf) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Address

`func (o *IppoolIpLeaseAllOf) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolIpLeaseAllOf) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolIpLeaseAllOf) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolIpLeaseAllOf) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolIpLeaseAllOf) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolIpLeaseAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolIpLeaseAllOf) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolIpLeaseAllOf) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### SetIpV4ConfigNil

`func (o *IppoolIpLeaseAllOf) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *IppoolIpLeaseAllOf) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
### GetIpV6Address

`func (o *IppoolIpLeaseAllOf) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *IppoolIpLeaseAllOf) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *IppoolIpLeaseAllOf) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *IppoolIpLeaseAllOf) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetIpV6Config

`func (o *IppoolIpLeaseAllOf) GetIpV6Config() IppoolIpV6Config`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *IppoolIpLeaseAllOf) GetIpV6ConfigOk() (*IppoolIpV6Config, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *IppoolIpLeaseAllOf) SetIpV6Config(v IppoolIpV6Config)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *IppoolIpLeaseAllOf) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *IppoolIpLeaseAllOf) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *IppoolIpLeaseAllOf) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
### GetVar0VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) GetVar0VirtualMachineNodeProfile() KubernetesVirtualMachineNodeProfileRelationship`

GetVar0VirtualMachineNodeProfile returns the Var0VirtualMachineNodeProfile field if non-nil, zero value otherwise.

### GetVar0VirtualMachineNodeProfileOk

`func (o *IppoolIpLeaseAllOf) GetVar0VirtualMachineNodeProfileOk() (*KubernetesVirtualMachineNodeProfileRelationship, bool)`

GetVar0VirtualMachineNodeProfileOk returns a tuple with the Var0VirtualMachineNodeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) SetVar0VirtualMachineNodeProfile(v KubernetesVirtualMachineNodeProfileRelationship)`

SetVar0VirtualMachineNodeProfile sets Var0VirtualMachineNodeProfile field to given value.

### HasVar0VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) HasVar0VirtualMachineNodeProfile() bool`

HasVar0VirtualMachineNodeProfile returns a boolean if a field has been set.

### GetVar1VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) GetVar1VirtualMachineNodeProfile() KubernetesVirtualMachineNodeProfileRelationship`

GetVar1VirtualMachineNodeProfile returns the Var1VirtualMachineNodeProfile field if non-nil, zero value otherwise.

### GetVar1VirtualMachineNodeProfileOk

`func (o *IppoolIpLeaseAllOf) GetVar1VirtualMachineNodeProfileOk() (*KubernetesVirtualMachineNodeProfileRelationship, bool)`

GetVar1VirtualMachineNodeProfileOk returns a tuple with the Var1VirtualMachineNodeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) SetVar1VirtualMachineNodeProfile(v KubernetesVirtualMachineNodeProfileRelationship)`

SetVar1VirtualMachineNodeProfile sets Var1VirtualMachineNodeProfile field to given value.

### HasVar1VirtualMachineNodeProfile

`func (o *IppoolIpLeaseAllOf) HasVar1VirtualMachineNodeProfile() bool`

HasVar1VirtualMachineNodeProfile returns a boolean if a field has been set.

### GetVar2ClusterProfile

`func (o *IppoolIpLeaseAllOf) GetVar2ClusterProfile() KubernetesClusterProfileRelationship`

GetVar2ClusterProfile returns the Var2ClusterProfile field if non-nil, zero value otherwise.

### GetVar2ClusterProfileOk

`func (o *IppoolIpLeaseAllOf) GetVar2ClusterProfileOk() (*KubernetesClusterProfileRelationship, bool)`

GetVar2ClusterProfileOk returns a tuple with the Var2ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2ClusterProfile

`func (o *IppoolIpLeaseAllOf) SetVar2ClusterProfile(v KubernetesClusterProfileRelationship)`

SetVar2ClusterProfile sets Var2ClusterProfile field to given value.

### HasVar2ClusterProfile

`func (o *IppoolIpLeaseAllOf) HasVar2ClusterProfile() bool`

HasVar2ClusterProfile returns a boolean if a field has been set.

### GetVar3ClusterProfile

`func (o *IppoolIpLeaseAllOf) GetVar3ClusterProfile() KubernetesClusterProfileRelationship`

GetVar3ClusterProfile returns the Var3ClusterProfile field if non-nil, zero value otherwise.

### GetVar3ClusterProfileOk

`func (o *IppoolIpLeaseAllOf) GetVar3ClusterProfileOk() (*KubernetesClusterProfileRelationship, bool)`

GetVar3ClusterProfileOk returns a tuple with the Var3ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3ClusterProfile

`func (o *IppoolIpLeaseAllOf) SetVar3ClusterProfile(v KubernetesClusterProfileRelationship)`

SetVar3ClusterProfile sets Var3ClusterProfile field to given value.

### HasVar3ClusterProfile

`func (o *IppoolIpLeaseAllOf) HasVar3ClusterProfile() bool`

HasVar3ClusterProfile returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolIpLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolIpLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolIpLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolIpLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *IppoolIpLeaseAllOf) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolIpLeaseAllOf) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolIpLeaseAllOf) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolIpLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IppoolIpLeaseAllOf) GetPoolMember() IppoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IppoolIpLeaseAllOf) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IppoolIpLeaseAllOf) SetPoolMember(v IppoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IppoolIpLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolIpLeaseAllOf) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolIpLeaseAllOf) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolIpLeaseAllOf) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolIpLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolIpLeaseAllOf) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolIpLeaseAllOf) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolIpLeaseAllOf) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolIpLeaseAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


