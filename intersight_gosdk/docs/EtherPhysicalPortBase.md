# EtherPhysicalPortBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**MacAddress** | Pointer to **string** | Mac Address of a port in the Fabric Interconnect. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Current Operational speed for this port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for ethernet physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id for port channel created on FI switch. | [optional] [readonly] 
**PortType** | Pointer to **string** | Defines the transport type for this port (ethernet OR fc). | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver model attached to a port in the Fabric Interconnect. | [optional] [readonly] 
**AcknowledgedPeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](PortInterfaceBaseRelationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](PortInterfaceBaseRelationship.md) |  | [optional] 

## Methods

### NewEtherPhysicalPortBase

`func NewEtherPhysicalPortBase(classId string, objectType string, ) *EtherPhysicalPortBase`

NewEtherPhysicalPortBase instantiates a new EtherPhysicalPortBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPhysicalPortBaseWithDefaults

`func NewEtherPhysicalPortBaseWithDefaults() *EtherPhysicalPortBase`

NewEtherPhysicalPortBaseWithDefaults instantiates a new EtherPhysicalPortBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPhysicalPortBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPhysicalPortBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPhysicalPortBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPhysicalPortBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPhysicalPortBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPhysicalPortBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAddress

`func (o *EtherPhysicalPortBase) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherPhysicalPortBase) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherPhysicalPortBase) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherPhysicalPortBase) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherPhysicalPortBase) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherPhysicalPortBase) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherPhysicalPortBase) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherPhysicalPortBase) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherPhysicalPortBase) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherPhysicalPortBase) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherPhysicalPortBase) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherPhysicalPortBase) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherPhysicalPortBase) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherPhysicalPortBase) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherPhysicalPortBase) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherPhysicalPortBase) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherPhysicalPortBase) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherPhysicalPortBase) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherPhysicalPortBase) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherPhysicalPortBase) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetPortType

`func (o *EtherPhysicalPortBase) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EtherPhysicalPortBase) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EtherPhysicalPortBase) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EtherPhysicalPortBase) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetTransceiverType

`func (o *EtherPhysicalPortBase) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *EtherPhysicalPortBase) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *EtherPhysicalPortBase) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *EtherPhysicalPortBase) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBase) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *EtherPhysicalPortBase) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBase) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *EtherPhysicalPortBase) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherPhysicalPortBase) GetPeerInterface() PortInterfaceBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherPhysicalPortBase) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherPhysicalPortBase) SetPeerInterface(v PortInterfaceBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherPhysicalPortBase) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


