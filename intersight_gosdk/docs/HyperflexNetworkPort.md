# HyperflexNetworkPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NetworkPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NetworkPort"]
**BondState** | Pointer to [**NullableHyperflexBondState**](HyperflexBondState.md) |  | [optional] 
**NetInterfaces** | Pointer to **[]string** |  | [optional] 
**PortType** | Pointer to **string** | The type of the network port. * &#x60;unknown&#x60; - This port is of an unknown port type. * &#x60;hypervisor&#x60; - This port is connected to the hypervisor. * &#x60;vm&#x60; - This port is connected to a VM. * &#x60;uplink&#x60; - This port is an uplink port. | [optional] [readonly] [default to "unknown"]
**Vlans** | Pointer to **string** | The vlan id associated with this port. | [optional] [readonly] 

## Methods

### NewHyperflexNetworkPort

`func NewHyperflexNetworkPort(classId string, objectType string, ) *HyperflexNetworkPort`

NewHyperflexNetworkPort instantiates a new HyperflexNetworkPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNetworkPortWithDefaults

`func NewHyperflexNetworkPortWithDefaults() *HyperflexNetworkPort`

NewHyperflexNetworkPortWithDefaults instantiates a new HyperflexNetworkPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNetworkPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNetworkPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNetworkPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNetworkPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNetworkPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNetworkPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *HyperflexNetworkPort) GetBondState() HyperflexBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *HyperflexNetworkPort) GetBondStateOk() (*HyperflexBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *HyperflexNetworkPort) SetBondState(v HyperflexBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *HyperflexNetworkPort) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *HyperflexNetworkPort) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *HyperflexNetworkPort) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetNetInterfaces

`func (o *HyperflexNetworkPort) GetNetInterfaces() []string`

GetNetInterfaces returns the NetInterfaces field if non-nil, zero value otherwise.

### GetNetInterfacesOk

`func (o *HyperflexNetworkPort) GetNetInterfacesOk() (*[]string, bool)`

GetNetInterfacesOk returns a tuple with the NetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterfaces

`func (o *HyperflexNetworkPort) SetNetInterfaces(v []string)`

SetNetInterfaces sets NetInterfaces field to given value.

### HasNetInterfaces

`func (o *HyperflexNetworkPort) HasNetInterfaces() bool`

HasNetInterfaces returns a boolean if a field has been set.

### SetNetInterfacesNil

`func (o *HyperflexNetworkPort) SetNetInterfacesNil(b bool)`

 SetNetInterfacesNil sets the value for NetInterfaces to be an explicit nil

### UnsetNetInterfaces
`func (o *HyperflexNetworkPort) UnsetNetInterfaces()`

UnsetNetInterfaces ensures that no value is present for NetInterfaces, not even an explicit nil
### GetPortType

`func (o *HyperflexNetworkPort) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *HyperflexNetworkPort) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *HyperflexNetworkPort) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *HyperflexNetworkPort) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetVlans

`func (o *HyperflexNetworkPort) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *HyperflexNetworkPort) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *HyperflexNetworkPort) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *HyperflexNetworkPort) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


