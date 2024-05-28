# FabricVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.Vlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.Vlan"]
**AutoAllowOnUplinks** | Pointer to **bool** | Enable to automatically allow this VLAN on all uplinks. Disable must be specified for Disjoint Layer 2 VLAN configuration. Default VLAN-1 cannot be configured as Disjoint Layer 2 VLAN. | [optional] [default to true]
**IsNative** | Pointer to **bool** | Used to define whether this VLAN is to be classified as &#39;native&#39; for traffic in this FI. | [optional] 
**Name** | Pointer to **string** | The &#39;name&#39; used to identify this VLAN. | [optional] 
**PrimaryVlanId** | Pointer to **int64** | The Primary VLAN ID of the VLAN, if the sharing type of the VLAN is Isolated or Community. | [optional] [default to 0]
**SharingType** | Pointer to **string** | The sharing type of this VLAN. * &#x60;None&#x60; - This represents a regular VLAN. * &#x60;Primary&#x60; - This represents a primary VLAN. * &#x60;Isolated&#x60; - This represents an isolated VLAN. * &#x60;Community&#x60; - This represents a community VLAN. | [optional] [default to "None"]
**VlanId** | Pointer to **int64** | The identifier for this Virtual LAN. | [optional] 
**EthNetworkPolicy** | Pointer to [**NullableFabricEthNetworkPolicyRelationship**](FabricEthNetworkPolicyRelationship.md) |  | [optional] 
**MulticastPolicy** | Pointer to [**NullableFabricMulticastPolicyRelationship**](FabricMulticastPolicyRelationship.md) |  | [optional] 
**VlanSet** | Pointer to [**NullableFabricVlanSetRelationship**](FabricVlanSetRelationship.md) |  | [optional] 

## Methods

### NewFabricVlan

`func NewFabricVlan(classId string, objectType string, ) *FabricVlan`

NewFabricVlan instantiates a new FabricVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanWithDefaults

`func NewFabricVlanWithDefaults() *FabricVlan`

NewFabricVlanWithDefaults instantiates a new FabricVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlan) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlan) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlan) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlan) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlan) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlan) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoAllowOnUplinks

`func (o *FabricVlan) GetAutoAllowOnUplinks() bool`

GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field if non-nil, zero value otherwise.

### GetAutoAllowOnUplinksOk

`func (o *FabricVlan) GetAutoAllowOnUplinksOk() (*bool, bool)`

GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllowOnUplinks

`func (o *FabricVlan) SetAutoAllowOnUplinks(v bool)`

SetAutoAllowOnUplinks sets AutoAllowOnUplinks field to given value.

### HasAutoAllowOnUplinks

`func (o *FabricVlan) HasAutoAllowOnUplinks() bool`

HasAutoAllowOnUplinks returns a boolean if a field has been set.

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

### GetPrimaryVlanId

`func (o *FabricVlan) GetPrimaryVlanId() int64`

GetPrimaryVlanId returns the PrimaryVlanId field if non-nil, zero value otherwise.

### GetPrimaryVlanIdOk

`func (o *FabricVlan) GetPrimaryVlanIdOk() (*int64, bool)`

GetPrimaryVlanIdOk returns a tuple with the PrimaryVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryVlanId

`func (o *FabricVlan) SetPrimaryVlanId(v int64)`

SetPrimaryVlanId sets PrimaryVlanId field to given value.

### HasPrimaryVlanId

`func (o *FabricVlan) HasPrimaryVlanId() bool`

HasPrimaryVlanId returns a boolean if a field has been set.

### GetSharingType

`func (o *FabricVlan) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *FabricVlan) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *FabricVlan) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.

### HasSharingType

`func (o *FabricVlan) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.

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

### SetEthNetworkPolicyNil

`func (o *FabricVlan) SetEthNetworkPolicyNil(b bool)`

 SetEthNetworkPolicyNil sets the value for EthNetworkPolicy to be an explicit nil

### UnsetEthNetworkPolicy
`func (o *FabricVlan) UnsetEthNetworkPolicy()`

UnsetEthNetworkPolicy ensures that no value is present for EthNetworkPolicy, not even an explicit nil
### GetMulticastPolicy

`func (o *FabricVlan) GetMulticastPolicy() FabricMulticastPolicyRelationship`

GetMulticastPolicy returns the MulticastPolicy field if non-nil, zero value otherwise.

### GetMulticastPolicyOk

`func (o *FabricVlan) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool)`

GetMulticastPolicyOk returns a tuple with the MulticastPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastPolicy

`func (o *FabricVlan) SetMulticastPolicy(v FabricMulticastPolicyRelationship)`

SetMulticastPolicy sets MulticastPolicy field to given value.

### HasMulticastPolicy

`func (o *FabricVlan) HasMulticastPolicy() bool`

HasMulticastPolicy returns a boolean if a field has been set.

### SetMulticastPolicyNil

`func (o *FabricVlan) SetMulticastPolicyNil(b bool)`

 SetMulticastPolicyNil sets the value for MulticastPolicy to be an explicit nil

### UnsetMulticastPolicy
`func (o *FabricVlan) UnsetMulticastPolicy()`

UnsetMulticastPolicy ensures that no value is present for MulticastPolicy, not even an explicit nil
### GetVlanSet

`func (o *FabricVlan) GetVlanSet() FabricVlanSetRelationship`

GetVlanSet returns the VlanSet field if non-nil, zero value otherwise.

### GetVlanSetOk

`func (o *FabricVlan) GetVlanSetOk() (*FabricVlanSetRelationship, bool)`

GetVlanSetOk returns a tuple with the VlanSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSet

`func (o *FabricVlan) SetVlanSet(v FabricVlanSetRelationship)`

SetVlanSet sets VlanSet field to given value.

### HasVlanSet

`func (o *FabricVlan) HasVlanSet() bool`

HasVlanSet returns a boolean if a field has been set.

### SetVlanSetNil

`func (o *FabricVlan) SetVlanSetNil(b bool)`

 SetVlanSetNil sets the value for VlanSet to be an explicit nil

### UnsetVlanSet
`func (o *FabricVlan) UnsetVlanSet()`

UnsetVlanSet ensures that no value is present for VlanSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


