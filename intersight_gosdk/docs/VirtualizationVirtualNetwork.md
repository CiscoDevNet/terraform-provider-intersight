# VirtualizationVirtualNetwork

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

### NewVirtualizationVirtualNetwork

`func NewVirtualizationVirtualNetwork(classId string, objectType string, ) *VirtualizationVirtualNetwork`

NewVirtualizationVirtualNetwork instantiates a new VirtualizationVirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualNetworkWithDefaults

`func NewVirtualizationVirtualNetworkWithDefaults() *VirtualizationVirtualNetwork`

NewVirtualizationVirtualNetworkWithDefaults instantiates a new VirtualizationVirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VirtualizationVirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualNetwork) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualNetwork) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualNetwork) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualNetwork) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetInfraNetwork

`func (o *VirtualizationVirtualNetwork) GetInfraNetwork() bool`

GetInfraNetwork returns the InfraNetwork field if non-nil, zero value otherwise.

### GetInfraNetworkOk

`func (o *VirtualizationVirtualNetwork) GetInfraNetworkOk() (*bool, bool)`

GetInfraNetworkOk returns a tuple with the InfraNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraNetwork

`func (o *VirtualizationVirtualNetwork) SetInfraNetwork(v bool)`

SetInfraNetwork sets InfraNetwork field to given value.

### HasInfraNetwork

`func (o *VirtualizationVirtualNetwork) HasInfraNetwork() bool`

HasInfraNetwork returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVirtualNetwork) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVirtualNetwork) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVirtualNetwork) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVirtualNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkType

`func (o *VirtualizationVirtualNetwork) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VirtualizationVirtualNetwork) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VirtualizationVirtualNetwork) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VirtualizationVirtualNetwork) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetTrunk

`func (o *VirtualizationVirtualNetwork) GetTrunk() []string`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *VirtualizationVirtualNetwork) GetTrunkOk() (*[]string, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *VirtualizationVirtualNetwork) SetTrunk(v []string)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *VirtualizationVirtualNetwork) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.

### SetTrunkNil

`func (o *VirtualizationVirtualNetwork) SetTrunkNil(b bool)`

 SetTrunkNil sets the value for Trunk to be an explicit nil

### UnsetTrunk
`func (o *VirtualizationVirtualNetwork) UnsetTrunk()`

UnsetTrunk ensures that no value is present for Trunk, not even an explicit nil
### GetVlan

`func (o *VirtualizationVirtualNetwork) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VirtualizationVirtualNetwork) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VirtualizationVirtualNetwork) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VirtualizationVirtualNetwork) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVswitch

`func (o *VirtualizationVirtualNetwork) GetVswitch() string`

GetVswitch returns the Vswitch field if non-nil, zero value otherwise.

### GetVswitchOk

`func (o *VirtualizationVirtualNetwork) GetVswitchOk() (*string, bool)`

GetVswitchOk returns a tuple with the Vswitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitch

`func (o *VirtualizationVirtualNetwork) SetVswitch(v string)`

SetVswitch sets Vswitch field to given value.

### HasVswitch

`func (o *VirtualizationVirtualNetwork) HasVswitch() bool`

HasVswitch returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVirtualNetwork) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualNetwork) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualNetwork) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualNetwork) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualNetwork) GetInventory() VirtualizationBaseVirtualNetworkRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualNetwork) GetInventoryOk() (*VirtualizationBaseVirtualNetworkRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualNetwork) SetInventory(v VirtualizationBaseVirtualNetworkRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualNetwork) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualNetwork) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualNetwork) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualNetwork) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualNetwork) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualNetwork) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualNetwork) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualNetwork) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualNetwork) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


