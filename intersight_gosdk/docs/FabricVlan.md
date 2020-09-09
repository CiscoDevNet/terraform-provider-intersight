# FabricVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNative** | Pointer to **bool** | Used to define whether this VLAN is to be classified as &#39;native&#39; for traffic in this FI. | [optional] 
**Name** | Pointer to **string** | The &#39;name&#39; used to identify this VLAN. | [optional] 
**VlanId** | Pointer to **int64** | The identifier for this Virtual LAN. | [optional] 
**EthNetworkPolicy** | Pointer to [**FabricEthNetworkPolicyRelationship**](fabric.EthNetworkPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricVlan

`func NewFabricVlan() *FabricVlan`

NewFabricVlan instantiates a new FabricVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanWithDefaults

`func NewFabricVlanWithDefaults() *FabricVlan`

NewFabricVlanWithDefaults instantiates a new FabricVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNative

`func (o *FabricVlan) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *FabricVlan) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *FabricVlan) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.

### HasIsNative

`func (o *FabricVlan) HasIsNative() bool`

HasIsNative returns a boolean if a field has been set.

### GetName

`func (o *FabricVlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanId

`func (o *FabricVlan) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *FabricVlan) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *FabricVlan) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *FabricVlan) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *FabricVlan) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *FabricVlan) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *FabricVlan) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *FabricVlan) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


