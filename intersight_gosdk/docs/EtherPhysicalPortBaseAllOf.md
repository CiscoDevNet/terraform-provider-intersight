# EtherPhysicalPortBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | Mac Address of a port in the Fabric Interconnect. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Current Operational speed for this port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for ethernet physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id for port channel created on FI switch. | [optional] [readonly] 
**PortType** | Pointer to **string** | Defines the transport type for this port (ethernet OR fc). | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver model attached to a port in the Fabric Interconnect. | [optional] [readonly] 
**AcknowledgedPeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](port.InterfaceBase.Relationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](port.InterfaceBase.Relationship.md) |  | [optional] 

## Methods

### NewEtherPhysicalPortBaseAllOf

`func NewEtherPhysicalPortBaseAllOf() *EtherPhysicalPortBaseAllOf`

NewEtherPhysicalPortBaseAllOf instantiates a new EtherPhysicalPortBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPhysicalPortBaseAllOfWithDefaults

`func NewEtherPhysicalPortBaseAllOfWithDefaults() *EtherPhysicalPortBaseAllOf`

NewEtherPhysicalPortBaseAllOfWithDefaults instantiates a new EtherPhysicalPortBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *EtherPhysicalPortBaseAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherPhysicalPortBaseAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherPhysicalPortBaseAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherPhysicalPortBaseAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherPhysicalPortBaseAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherPhysicalPortBaseAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherPhysicalPortBaseAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherPhysicalPortBaseAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherPhysicalPortBaseAllOf) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherPhysicalPortBaseAllOf) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherPhysicalPortBaseAllOf) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherPhysicalPortBaseAllOf) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherPhysicalPortBaseAllOf) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherPhysicalPortBaseAllOf) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherPhysicalPortBaseAllOf) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherPhysicalPortBaseAllOf) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherPhysicalPortBaseAllOf) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherPhysicalPortBaseAllOf) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherPhysicalPortBaseAllOf) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherPhysicalPortBaseAllOf) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetPortType

`func (o *EtherPhysicalPortBaseAllOf) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EtherPhysicalPortBaseAllOf) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EtherPhysicalPortBaseAllOf) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EtherPhysicalPortBaseAllOf) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetTransceiverType

`func (o *EtherPhysicalPortBaseAllOf) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *EtherPhysicalPortBaseAllOf) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *EtherPhysicalPortBaseAllOf) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *EtherPhysicalPortBaseAllOf) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *EtherPhysicalPortBaseAllOf) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) GetPeerInterface() PortInterfaceBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherPhysicalPortBaseAllOf) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) SetPeerInterface(v PortInterfaceBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherPhysicalPortBaseAllOf) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


