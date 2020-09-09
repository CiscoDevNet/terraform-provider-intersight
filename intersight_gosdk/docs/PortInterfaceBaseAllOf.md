# PortInterfaceBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperState** | Pointer to **string** | Operational state of an Interface. | [optional] 
**AcknowledgedPeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 

## Methods

### NewPortInterfaceBaseAllOf

`func NewPortInterfaceBaseAllOf() *PortInterfaceBaseAllOf`

NewPortInterfaceBaseAllOf instantiates a new PortInterfaceBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortInterfaceBaseAllOfWithDefaults

`func NewPortInterfaceBaseAllOfWithDefaults() *PortInterfaceBaseAllOf`

NewPortInterfaceBaseAllOfWithDefaults instantiates a new PortInterfaceBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperState

`func (o *PortInterfaceBaseAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PortInterfaceBaseAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PortInterfaceBaseAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PortInterfaceBaseAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *PortInterfaceBaseAllOf) GetAcknowledgedPeerInterface() EtherPhysicalPortBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *PortInterfaceBaseAllOf) GetAcknowledgedPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *PortInterfaceBaseAllOf) SetAcknowledgedPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *PortInterfaceBaseAllOf) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *PortInterfaceBaseAllOf) GetPeerInterface() EtherPhysicalPortBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *PortInterfaceBaseAllOf) GetPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *PortInterfaceBaseAllOf) SetPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *PortInterfaceBaseAllOf) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


