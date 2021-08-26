# FabricVlanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.Vlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.Vlan"]
**AutoAllowOnUplinks** | Pointer to **bool** | Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI. | [optional] [default to true]
**IsNative** | Pointer to **bool** | Used to define whether this VLAN is to be classified as &#39;native&#39; for traffic in this FI. | [optional] 
**Name** | Pointer to **string** | The &#39;name&#39; used to identify this VLAN. | [optional] 
**VlanId** | Pointer to **int64** | The identifier for this Virtual LAN. | [optional] 
**EthNetworkPolicy** | Pointer to [**FabricEthNetworkPolicyRelationship**](FabricEthNetworkPolicyRelationship.md) |  | [optional] 
**MulticastPolicy** | Pointer to [**FabricMulticastPolicyRelationship**](FabricMulticastPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricVlanAllOf

`func NewFabricVlanAllOf(classId string, objectType string, ) *FabricVlanAllOf`

NewFabricVlanAllOf instantiates a new FabricVlanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanAllOfWithDefaults

`func NewFabricVlanAllOfWithDefaults() *FabricVlanAllOf`

NewFabricVlanAllOfWithDefaults instantiates a new FabricVlanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoAllowOnUplinks

`func (o *FabricVlanAllOf) GetAutoAllowOnUplinks() bool`

GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field if non-nil, zero value otherwise.

### GetAutoAllowOnUplinksOk

`func (o *FabricVlanAllOf) GetAutoAllowOnUplinksOk() (*bool, bool)`

GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllowOnUplinks

`func (o *FabricVlanAllOf) SetAutoAllowOnUplinks(v bool)`

SetAutoAllowOnUplinks sets AutoAllowOnUplinks field to given value.

### HasAutoAllowOnUplinks

`func (o *FabricVlanAllOf) HasAutoAllowOnUplinks() bool`

HasAutoAllowOnUplinks returns a boolean if a field has been set.

### GetIsNative

`func (o *FabricVlanAllOf) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *FabricVlanAllOf) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *FabricVlanAllOf) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.

### HasIsNative

`func (o *FabricVlanAllOf) HasIsNative() bool`

HasIsNative returns a boolean if a field has been set.

### GetName

`func (o *FabricVlanAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVlanAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVlanAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVlanAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanId

`func (o *FabricVlanAllOf) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *FabricVlanAllOf) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *FabricVlanAllOf) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *FabricVlanAllOf) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *FabricVlanAllOf) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *FabricVlanAllOf) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *FabricVlanAllOf) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *FabricVlanAllOf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetMulticastPolicy

`func (o *FabricVlanAllOf) GetMulticastPolicy() FabricMulticastPolicyRelationship`

GetMulticastPolicy returns the MulticastPolicy field if non-nil, zero value otherwise.

### GetMulticastPolicyOk

`func (o *FabricVlanAllOf) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool)`

GetMulticastPolicyOk returns a tuple with the MulticastPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastPolicy

`func (o *FabricVlanAllOf) SetMulticastPolicy(v FabricMulticastPolicyRelationship)`

SetMulticastPolicy sets MulticastPolicy field to given value.

### HasMulticastPolicy

`func (o *FabricVlanAllOf) HasMulticastPolicy() bool`

HasMulticastPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


