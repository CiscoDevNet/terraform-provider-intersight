# EtherPhysicalPortBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AccessVlan** | Pointer to **string** | Access VLAN for this port. | [optional] [readonly] 
**AdminFec** | Pointer to **string** | Administratively configured FEC mode for this port. * &#x60;&#x60; - Default value for FEC state. * &#x60;Not Supported&#x60; - FEC is not supported on this port. * &#x60;Disabled&#x60; - FEC is disabled on this port. * &#x60;Auto&#x60; - FEC mode is automatically negotiated between link partners. * &#x60;Cl74&#x60; - FEC is configured to use the IEEE Clause 74 (FireCode) standard. * &#x60;RS-IEEE(Cl108)&#x60; - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard. * &#x60;KP&#x60; - FEC is configured to use the KP (Backplane Ethernet) FEC mode. | [optional] [readonly] [default to ""]
**AllowedVlans** | Pointer to **string** | Allowed VLANs on this port. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac Address of a port in the Fabric Interconnect. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port. | [optional] [readonly] 
**NativeVlan** | Pointer to **string** | Native VLAN for this port. | [optional] [readonly] 
**OperFec** | Pointer to **string** | Operational FEC mode for this port. * &#x60;&#x60; - Default value for FEC state. * &#x60;Not Supported&#x60; - FEC is not supported on this port. * &#x60;Disabled&#x60; - FEC is disabled on this port. * &#x60;Auto&#x60; - FEC mode is automatically negotiated between link partners. * &#x60;Cl74&#x60; - FEC is configured to use the IEEE Clause 74 (FireCode) standard. * &#x60;RS-IEEE(Cl108)&#x60; - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard. * &#x60;KP&#x60; - FEC is configured to use the KP (Backplane Ethernet) FEC mode. | [optional] [readonly] [default to ""]
**OperSpeed** | Pointer to **string** | Current Operational speed for this port. | [optional] [readonly] 
**OperVlans** | Pointer to **string** | Operational VLANs on this port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for ethernet physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id for port channel created on FI switch. | [optional] [readonly] 
**PortName** | Pointer to **string** | Switch physical port name. | [optional] [readonly] 
**PortType** | Pointer to **string** | Defines the transport type for this port (ethernet OR fc). | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver model attached to a port in the Fabric Interconnect. | [optional] [readonly] 
**AcknowledgedPeerInterface** | Pointer to [**NullablePortInterfaceBaseRelationship**](PortInterfaceBaseRelationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**NullablePortInterfaceBaseRelationship**](PortInterfaceBaseRelationship.md) |  | [optional] 

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


### GetAccessVlan

`func (o *EtherPhysicalPortBase) GetAccessVlan() string`

GetAccessVlan returns the AccessVlan field if non-nil, zero value otherwise.

### GetAccessVlanOk

`func (o *EtherPhysicalPortBase) GetAccessVlanOk() (*string, bool)`

GetAccessVlanOk returns a tuple with the AccessVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlan

`func (o *EtherPhysicalPortBase) SetAccessVlan(v string)`

SetAccessVlan sets AccessVlan field to given value.

### HasAccessVlan

`func (o *EtherPhysicalPortBase) HasAccessVlan() bool`

HasAccessVlan returns a boolean if a field has been set.

### GetAdminFec

`func (o *EtherPhysicalPortBase) GetAdminFec() string`

GetAdminFec returns the AdminFec field if non-nil, zero value otherwise.

### GetAdminFecOk

`func (o *EtherPhysicalPortBase) GetAdminFecOk() (*string, bool)`

GetAdminFecOk returns a tuple with the AdminFec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFec

`func (o *EtherPhysicalPortBase) SetAdminFec(v string)`

SetAdminFec sets AdminFec field to given value.

### HasAdminFec

`func (o *EtherPhysicalPortBase) HasAdminFec() bool`

HasAdminFec returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *EtherPhysicalPortBase) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *EtherPhysicalPortBase) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *EtherPhysicalPortBase) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *EtherPhysicalPortBase) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

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

### GetNativeVlan

`func (o *EtherPhysicalPortBase) GetNativeVlan() string`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *EtherPhysicalPortBase) GetNativeVlanOk() (*string, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *EtherPhysicalPortBase) SetNativeVlan(v string)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *EtherPhysicalPortBase) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetOperFec

`func (o *EtherPhysicalPortBase) GetOperFec() string`

GetOperFec returns the OperFec field if non-nil, zero value otherwise.

### GetOperFecOk

`func (o *EtherPhysicalPortBase) GetOperFecOk() (*string, bool)`

GetOperFecOk returns a tuple with the OperFec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperFec

`func (o *EtherPhysicalPortBase) SetOperFec(v string)`

SetOperFec sets OperFec field to given value.

### HasOperFec

`func (o *EtherPhysicalPortBase) HasOperFec() bool`

HasOperFec returns a boolean if a field has been set.

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

### GetOperVlans

`func (o *EtherPhysicalPortBase) GetOperVlans() string`

GetOperVlans returns the OperVlans field if non-nil, zero value otherwise.

### GetOperVlansOk

`func (o *EtherPhysicalPortBase) GetOperVlansOk() (*string, bool)`

GetOperVlansOk returns a tuple with the OperVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperVlans

`func (o *EtherPhysicalPortBase) SetOperVlans(v string)`

SetOperVlans sets OperVlans field to given value.

### HasOperVlans

`func (o *EtherPhysicalPortBase) HasOperVlans() bool`

HasOperVlans returns a boolean if a field has been set.

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

### GetPortName

`func (o *EtherPhysicalPortBase) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *EtherPhysicalPortBase) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *EtherPhysicalPortBase) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *EtherPhysicalPortBase) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

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

### SetAcknowledgedPeerInterfaceNil

`func (o *EtherPhysicalPortBase) SetAcknowledgedPeerInterfaceNil(b bool)`

 SetAcknowledgedPeerInterfaceNil sets the value for AcknowledgedPeerInterface to be an explicit nil

### UnsetAcknowledgedPeerInterface
`func (o *EtherPhysicalPortBase) UnsetAcknowledgedPeerInterface()`

UnsetAcknowledgedPeerInterface ensures that no value is present for AcknowledgedPeerInterface, not even an explicit nil
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

### SetPeerInterfaceNil

`func (o *EtherPhysicalPortBase) SetPeerInterfaceNil(b bool)`

 SetPeerInterfaceNil sets the value for PeerInterface to be an explicit nil

### UnsetPeerInterface
`func (o *EtherPhysicalPortBase) UnsetPeerInterface()`

UnsetPeerInterface ensures that no value is present for PeerInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


