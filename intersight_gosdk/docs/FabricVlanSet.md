# FabricVlanSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanSet"]
**AutoAllowOnUplinks** | Pointer to **bool** | Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI. | [optional] [readonly] [default to true]
**IsNative** | Pointer to **bool** | Used to define whether this VLAN is to be classified as &#39;native&#39; for traffic in this FI. | [optional] [readonly] 
**Name** | Pointer to **string** | The &#39;name&#39; used to identify this VLAN. | [optional] [readonly] 
**PrimaryVlanId** | Pointer to **int64** | The Primary VLAN ID of the VLAN, if the sharing type of the VLAN is Isolated or Community. | [optional] [readonly] [default to 0]
**SharingType** | Pointer to **string** | The sharing type of this VLAN. * &#x60;None&#x60; - This represents a regular VLAN. * &#x60;Primary&#x60; - This represents a primary VLAN. * &#x60;Isolated&#x60; - This represents an isolated VLAN. * &#x60;Community&#x60; - This represents a community VLAN. | [optional] [readonly] [default to "None"]
**Vlans** | Pointer to **string** | Set of VLANs defined by VLAN object with identical configuration. | [optional] [readonly] 
**EthNetworkPolicy** | Pointer to [**FabricEthNetworkPolicyRelationship**](FabricEthNetworkPolicyRelationship.md) |  | [optional] 
**MulticastPolicy** | Pointer to [**FabricMulticastPolicyRelationship**](FabricMulticastPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricVlanSet

`func NewFabricVlanSet(classId string, objectType string, ) *FabricVlanSet`

NewFabricVlanSet instantiates a new FabricVlanSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSetWithDefaults

`func NewFabricVlanSetWithDefaults() *FabricVlanSet`

NewFabricVlanSetWithDefaults instantiates a new FabricVlanSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanSet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanSet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanSet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanSet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanSet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanSet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoAllowOnUplinks

`func (o *FabricVlanSet) GetAutoAllowOnUplinks() bool`

GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field if non-nil, zero value otherwise.

### GetAutoAllowOnUplinksOk

`func (o *FabricVlanSet) GetAutoAllowOnUplinksOk() (*bool, bool)`

GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllowOnUplinks

`func (o *FabricVlanSet) SetAutoAllowOnUplinks(v bool)`

SetAutoAllowOnUplinks sets AutoAllowOnUplinks field to given value.

### HasAutoAllowOnUplinks

`func (o *FabricVlanSet) HasAutoAllowOnUplinks() bool`

HasAutoAllowOnUplinks returns a boolean if a field has been set.

### GetIsNative

`func (o *FabricVlanSet) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *FabricVlanSet) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *FabricVlanSet) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.

### HasIsNative

`func (o *FabricVlanSet) HasIsNative() bool`

HasIsNative returns a boolean if a field has been set.

### GetName

`func (o *FabricVlanSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVlanSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVlanSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVlanSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryVlanId

`func (o *FabricVlanSet) GetPrimaryVlanId() int64`

GetPrimaryVlanId returns the PrimaryVlanId field if non-nil, zero value otherwise.

### GetPrimaryVlanIdOk

`func (o *FabricVlanSet) GetPrimaryVlanIdOk() (*int64, bool)`

GetPrimaryVlanIdOk returns a tuple with the PrimaryVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryVlanId

`func (o *FabricVlanSet) SetPrimaryVlanId(v int64)`

SetPrimaryVlanId sets PrimaryVlanId field to given value.

### HasPrimaryVlanId

`func (o *FabricVlanSet) HasPrimaryVlanId() bool`

HasPrimaryVlanId returns a boolean if a field has been set.

### GetSharingType

`func (o *FabricVlanSet) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *FabricVlanSet) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *FabricVlanSet) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.

### HasSharingType

`func (o *FabricVlanSet) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.

### GetVlans

`func (o *FabricVlanSet) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *FabricVlanSet) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *FabricVlanSet) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *FabricVlanSet) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *FabricVlanSet) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *FabricVlanSet) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *FabricVlanSet) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *FabricVlanSet) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetMulticastPolicy

`func (o *FabricVlanSet) GetMulticastPolicy() FabricMulticastPolicyRelationship`

GetMulticastPolicy returns the MulticastPolicy field if non-nil, zero value otherwise.

### GetMulticastPolicyOk

`func (o *FabricVlanSet) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool)`

GetMulticastPolicyOk returns a tuple with the MulticastPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastPolicy

`func (o *FabricVlanSet) SetMulticastPolicy(v FabricMulticastPolicyRelationship)`

SetMulticastPolicy sets MulticastPolicy field to given value.

### HasMulticastPolicy

`func (o *FabricVlanSet) HasMulticastPolicy() bool`

HasMulticastPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


