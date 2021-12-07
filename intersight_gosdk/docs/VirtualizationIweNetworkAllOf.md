# VirtualizationIweNetworkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweNetwork"]
**InfraNetwork** | Pointer to **bool** | A flag to distinguish if a network belongs to a HxAp infrastructure network or a user defined network that guest workloads can use. | [optional] 
**Mtu** | Pointer to **int64** | The MTU size of the network. | [optional] 
**NetworkType** | Pointer to **string** | Network attachment type, only \&quot;L2\&quot; is available now. * &#x60;unknown&#x60; - This network is of an unknown network type. * &#x60;L2&#x60; - A Layer 2 switching network type. | [optional] [default to "unknown"]
**Trunk** | Pointer to **[]string** |  | [optional] 
**Vlan** | Pointer to **int64** | A vlan id set on the network attachment point. | [optional] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**Dvswitch** | Pointer to [**VirtualizationIweDvswitchRelationship**](VirtualizationIweDvswitchRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweNetworkAllOf

`func NewVirtualizationIweNetworkAllOf(classId string, objectType string, ) *VirtualizationIweNetworkAllOf`

NewVirtualizationIweNetworkAllOf instantiates a new VirtualizationIweNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweNetworkAllOfWithDefaults

`func NewVirtualizationIweNetworkAllOfWithDefaults() *VirtualizationIweNetworkAllOf`

NewVirtualizationIweNetworkAllOfWithDefaults instantiates a new VirtualizationIweNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInfraNetwork

`func (o *VirtualizationIweNetworkAllOf) GetInfraNetwork() bool`

GetInfraNetwork returns the InfraNetwork field if non-nil, zero value otherwise.

### GetInfraNetworkOk

`func (o *VirtualizationIweNetworkAllOf) GetInfraNetworkOk() (*bool, bool)`

GetInfraNetworkOk returns a tuple with the InfraNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraNetwork

`func (o *VirtualizationIweNetworkAllOf) SetInfraNetwork(v bool)`

SetInfraNetwork sets InfraNetwork field to given value.

### HasInfraNetwork

`func (o *VirtualizationIweNetworkAllOf) HasInfraNetwork() bool`

HasInfraNetwork returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationIweNetworkAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationIweNetworkAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationIweNetworkAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationIweNetworkAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetworkType

`func (o *VirtualizationIweNetworkAllOf) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VirtualizationIweNetworkAllOf) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VirtualizationIweNetworkAllOf) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VirtualizationIweNetworkAllOf) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetTrunk

`func (o *VirtualizationIweNetworkAllOf) GetTrunk() []string`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *VirtualizationIweNetworkAllOf) GetTrunkOk() (*[]string, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *VirtualizationIweNetworkAllOf) SetTrunk(v []string)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *VirtualizationIweNetworkAllOf) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.

### SetTrunkNil

`func (o *VirtualizationIweNetworkAllOf) SetTrunkNil(b bool)`

 SetTrunkNil sets the value for Trunk to be an explicit nil

### UnsetTrunk
`func (o *VirtualizationIweNetworkAllOf) UnsetTrunk()`

UnsetTrunk ensures that no value is present for Trunk, not even an explicit nil
### GetVlan

`func (o *VirtualizationIweNetworkAllOf) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VirtualizationIweNetworkAllOf) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VirtualizationIweNetworkAllOf) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VirtualizationIweNetworkAllOf) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweNetworkAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweNetworkAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweNetworkAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweNetworkAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDvswitch

`func (o *VirtualizationIweNetworkAllOf) GetDvswitch() VirtualizationIweDvswitchRelationship`

GetDvswitch returns the Dvswitch field if non-nil, zero value otherwise.

### GetDvswitchOk

`func (o *VirtualizationIweNetworkAllOf) GetDvswitchOk() (*VirtualizationIweDvswitchRelationship, bool)`

GetDvswitchOk returns a tuple with the Dvswitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvswitch

`func (o *VirtualizationIweNetworkAllOf) SetDvswitch(v VirtualizationIweDvswitchRelationship)`

SetDvswitch sets Dvswitch field to given value.

### HasDvswitch

`func (o *VirtualizationIweNetworkAllOf) HasDvswitch() bool`

HasDvswitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


