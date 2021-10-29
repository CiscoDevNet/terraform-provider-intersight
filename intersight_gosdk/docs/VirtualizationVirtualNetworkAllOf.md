# VirtualizationVirtualNetworkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualNetwork"]
**Description** | Pointer to **string** | Human readable description about this network. | [optional] 
**Discovered** | Pointer to **bool** | Flag to indicate whether the configuration is created from inventory object. | [optional] [readonly] 
**InfraNetwork** | Pointer to **bool** | A flag to distinguish if a network belongs to an infrastructure network or a user defined network that guest workloads can use. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmissible unit of data in bytes that can be sent across the network. | [optional] 
**Name** | Pointer to **string** | Name of the virtual network. Name must be unique. | [optional] 
**NetworkType** | Pointer to **string** | Type of network layer, either L2 or L3. * &#x60;unknown&#x60; - This network is of an unknown network type. * &#x60;L2&#x60; - A Layer 2 switching network type. | [optional] [default to "unknown"]
**Trunk** | Pointer to **[]string** |  | [optional] 
**Vlan** | Pointer to **int64** | A VLAN id set on the network attachment point. | [optional] 
**Vswitch** | Pointer to **string** | Name of the virtual switch. | [optional] 
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](VirtualizationBaseClusterRelationship.md) |  | [optional] 
**Inventory** | Pointer to [**VirtualizationBaseVirtualNetworkRelationship**](VirtualizationBaseVirtualNetworkRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVirtualNetworkAllOf

`func NewVirtualizationVirtualNetworkAllOf(classId string, objectType string, ) *VirtualizationVirtualNetworkAllOf`

NewVirtualizationVirtualNetworkAllOf instantiates a new VirtualizationVirtualNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualNetworkAllOfWithDefaults

`func NewVirtualizationVirtualNetworkAllOfWithDefaults() *VirtualizationVirtualNetworkAllOf`

NewVirtualizationVirtualNetworkAllOfWithDefaults instantiates a new VirtualizationVirtualNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VirtualizationVirtualNetworkAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVirtualNetworkAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVirtualNetworkAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVirtualNetworkAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualNetworkAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualNetworkAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualNetworkAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualNetworkAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetInfraNetwork

`func (o *VirtualizationVirtualNetworkAllOf) GetInfraNetwork() bool`

GetInfraNetwork returns the InfraNetwork field if non-nil, zero value otherwise.

### GetInfraNetworkOk

`func (o *VirtualizationVirtualNetworkAllOf) GetInfraNetworkOk() (*bool, bool)`

GetInfraNetworkOk returns a tuple with the InfraNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraNetwork

`func (o *VirtualizationVirtualNetworkAllOf) SetInfraNetwork(v bool)`

SetInfraNetwork sets InfraNetwork field to given value.

### HasInfraNetwork

`func (o *VirtualizationVirtualNetworkAllOf) HasInfraNetwork() bool`

HasInfraNetwork returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVirtualNetworkAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVirtualNetworkAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVirtualNetworkAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVirtualNetworkAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualNetworkAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualNetworkAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualNetworkAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualNetworkAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkType

`func (o *VirtualizationVirtualNetworkAllOf) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VirtualizationVirtualNetworkAllOf) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VirtualizationVirtualNetworkAllOf) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VirtualizationVirtualNetworkAllOf) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetTrunk

`func (o *VirtualizationVirtualNetworkAllOf) GetTrunk() []string`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *VirtualizationVirtualNetworkAllOf) GetTrunkOk() (*[]string, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *VirtualizationVirtualNetworkAllOf) SetTrunk(v []string)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *VirtualizationVirtualNetworkAllOf) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.

### SetTrunkNil

`func (o *VirtualizationVirtualNetworkAllOf) SetTrunkNil(b bool)`

 SetTrunkNil sets the value for Trunk to be an explicit nil

### UnsetTrunk
`func (o *VirtualizationVirtualNetworkAllOf) UnsetTrunk()`

UnsetTrunk ensures that no value is present for Trunk, not even an explicit nil
### GetVlan

`func (o *VirtualizationVirtualNetworkAllOf) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VirtualizationVirtualNetworkAllOf) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VirtualizationVirtualNetworkAllOf) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VirtualizationVirtualNetworkAllOf) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVswitch

`func (o *VirtualizationVirtualNetworkAllOf) GetVswitch() string`

GetVswitch returns the Vswitch field if non-nil, zero value otherwise.

### GetVswitchOk

`func (o *VirtualizationVirtualNetworkAllOf) GetVswitchOk() (*string, bool)`

GetVswitchOk returns a tuple with the Vswitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitch

`func (o *VirtualizationVirtualNetworkAllOf) SetVswitch(v string)`

SetVswitch sets Vswitch field to given value.

### HasVswitch

`func (o *VirtualizationVirtualNetworkAllOf) HasVswitch() bool`

HasVswitch returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVirtualNetworkAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualNetworkAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualNetworkAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualNetworkAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualNetworkAllOf) GetInventory() VirtualizationBaseVirtualNetworkRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualNetworkAllOf) GetInventoryOk() (*VirtualizationBaseVirtualNetworkRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualNetworkAllOf) SetInventory(v VirtualizationBaseVirtualNetworkRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualNetworkAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualNetworkAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualNetworkAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualNetworkAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualNetworkAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualNetworkAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualNetworkAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualNetworkAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualNetworkAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


