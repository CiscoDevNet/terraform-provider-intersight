# VirtualizationNetworkPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.NetworkPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.NetworkPort"]
**BondState** | Pointer to [**NullableVirtualizationBondState**](VirtualizationBondState.md) |  | [optional] 
**NetInterfaces** | Pointer to **[]string** |  | [optional] 
**PortType** | Pointer to **string** | The type of the network port. * &#x60;unknown&#x60; - This port is of an unknown port type. * &#x60;hypervisor&#x60; - This port is connected to the hypervisor. * &#x60;vm&#x60; - This port is connected to a VM. * &#x60;uplink&#x60; - This port is an uplink port. | [optional] [readonly] [default to "unknown"]
**Vlans** | Pointer to **string** | The vlan id associated with this port. | [optional] [readonly] 

## Methods

### NewVirtualizationNetworkPortAllOf

`func NewVirtualizationNetworkPortAllOf(classId string, objectType string, ) *VirtualizationNetworkPortAllOf`

NewVirtualizationNetworkPortAllOf instantiates a new VirtualizationNetworkPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationNetworkPortAllOfWithDefaults

`func NewVirtualizationNetworkPortAllOfWithDefaults() *VirtualizationNetworkPortAllOf`

NewVirtualizationNetworkPortAllOfWithDefaults instantiates a new VirtualizationNetworkPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationNetworkPortAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationNetworkPortAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationNetworkPortAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationNetworkPortAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationNetworkPortAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationNetworkPortAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *VirtualizationNetworkPortAllOf) GetBondState() VirtualizationBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *VirtualizationNetworkPortAllOf) GetBondStateOk() (*VirtualizationBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *VirtualizationNetworkPortAllOf) SetBondState(v VirtualizationBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *VirtualizationNetworkPortAllOf) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *VirtualizationNetworkPortAllOf) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *VirtualizationNetworkPortAllOf) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetNetInterfaces

`func (o *VirtualizationNetworkPortAllOf) GetNetInterfaces() []string`

GetNetInterfaces returns the NetInterfaces field if non-nil, zero value otherwise.

### GetNetInterfacesOk

`func (o *VirtualizationNetworkPortAllOf) GetNetInterfacesOk() (*[]string, bool)`

GetNetInterfacesOk returns a tuple with the NetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterfaces

`func (o *VirtualizationNetworkPortAllOf) SetNetInterfaces(v []string)`

SetNetInterfaces sets NetInterfaces field to given value.

### HasNetInterfaces

`func (o *VirtualizationNetworkPortAllOf) HasNetInterfaces() bool`

HasNetInterfaces returns a boolean if a field has been set.

### SetNetInterfacesNil

`func (o *VirtualizationNetworkPortAllOf) SetNetInterfacesNil(b bool)`

 SetNetInterfacesNil sets the value for NetInterfaces to be an explicit nil

### UnsetNetInterfaces
`func (o *VirtualizationNetworkPortAllOf) UnsetNetInterfaces()`

UnsetNetInterfaces ensures that no value is present for NetInterfaces, not even an explicit nil
### GetPortType

`func (o *VirtualizationNetworkPortAllOf) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *VirtualizationNetworkPortAllOf) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *VirtualizationNetworkPortAllOf) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *VirtualizationNetworkPortAllOf) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetVlans

`func (o *VirtualizationNetworkPortAllOf) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VirtualizationNetworkPortAllOf) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VirtualizationNetworkPortAllOf) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *VirtualizationNetworkPortAllOf) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


