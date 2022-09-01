# FabricVlanSetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanSet"]
**AutoAllowOnUplinks** | Pointer to **bool** | Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI. | [optional] [readonly] [default to true]
**IsNative** | Pointer to **bool** | Used to define whether this VLAN is to be classified as &#39;native&#39; for traffic in this FI. | [optional] [readonly] 
**Name** | Pointer to **string** | The &#39;name&#39; used to identify this VLAN. | [optional] [readonly] 
**Vlans** | Pointer to **string** | Set of VLANs defined by VLAN object with identical configuration. | [optional] [readonly] 
**EthNetworkPolicy** | Pointer to [**FabricEthNetworkPolicyRelationship**](FabricEthNetworkPolicyRelationship.md) |  | [optional] 
**MulticastPolicy** | Pointer to [**FabricMulticastPolicyRelationship**](FabricMulticastPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricVlanSetAllOf

`func NewFabricVlanSetAllOf(classId string, objectType string, ) *FabricVlanSetAllOf`

NewFabricVlanSetAllOf instantiates a new FabricVlanSetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSetAllOfWithDefaults

`func NewFabricVlanSetAllOfWithDefaults() *FabricVlanSetAllOf`

NewFabricVlanSetAllOfWithDefaults instantiates a new FabricVlanSetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanSetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanSetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanSetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanSetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanSetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanSetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoAllowOnUplinks

`func (o *FabricVlanSetAllOf) GetAutoAllowOnUplinks() bool`

GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field if non-nil, zero value otherwise.

### GetAutoAllowOnUplinksOk

`func (o *FabricVlanSetAllOf) GetAutoAllowOnUplinksOk() (*bool, bool)`

GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllowOnUplinks

`func (o *FabricVlanSetAllOf) SetAutoAllowOnUplinks(v bool)`

SetAutoAllowOnUplinks sets AutoAllowOnUplinks field to given value.

### HasAutoAllowOnUplinks

`func (o *FabricVlanSetAllOf) HasAutoAllowOnUplinks() bool`

HasAutoAllowOnUplinks returns a boolean if a field has been set.

### GetIsNative

`func (o *FabricVlanSetAllOf) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *FabricVlanSetAllOf) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *FabricVlanSetAllOf) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.

### HasIsNative

`func (o *FabricVlanSetAllOf) HasIsNative() bool`

HasIsNative returns a boolean if a field has been set.

### GetName

`func (o *FabricVlanSetAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVlanSetAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVlanSetAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVlanSetAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlans

`func (o *FabricVlanSetAllOf) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *FabricVlanSetAllOf) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *FabricVlanSetAllOf) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *FabricVlanSetAllOf) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *FabricVlanSetAllOf) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *FabricVlanSetAllOf) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *FabricVlanSetAllOf) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *FabricVlanSetAllOf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetMulticastPolicy

`func (o *FabricVlanSetAllOf) GetMulticastPolicy() FabricMulticastPolicyRelationship`

GetMulticastPolicy returns the MulticastPolicy field if non-nil, zero value otherwise.

### GetMulticastPolicyOk

`func (o *FabricVlanSetAllOf) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool)`

GetMulticastPolicyOk returns a tuple with the MulticastPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastPolicy

`func (o *FabricVlanSetAllOf) SetMulticastPolicy(v FabricMulticastPolicyRelationship)`

SetMulticastPolicy sets MulticastPolicy field to given value.

### HasMulticastPolicy

`func (o *FabricVlanSetAllOf) HasMulticastPolicy() bool`

HasMulticastPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


