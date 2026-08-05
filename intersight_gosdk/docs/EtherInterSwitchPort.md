# EtherInterSwitchPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.InterSwitchPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.InterSwitchPort"]
**AccessVlan** | Pointer to **string** | Access VLAN for this port. | [optional] [readonly] 
**AdminFec** | Pointer to **string** | Administratively configured FEC mode for this port. * &#x60;&#x60; - Default value for FEC state. * &#x60;Not Supported&#x60; - FEC is not supported on this port. * &#x60;Disabled&#x60; - FEC is disabled on this port. * &#x60;Auto&#x60; - FEC mode is automatically negotiated between link partners. * &#x60;Cl74&#x60; - FEC is configured to use the IEEE Clause 74 (FireCode) standard. * &#x60;RS-IEEE(Cl108)&#x60; - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard. * &#x60;KP&#x60; - FEC is configured to use the KP (Backplane Ethernet) FEC mode. | [optional] [readonly] [default to ""]
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AllowedVlans** | Pointer to **string** | Allowed VLANs on this port. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac Address of a port in the Fabric Interconnect. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port. | [optional] [readonly] 
**NativeVlan** | Pointer to **string** | Native VLAN for this port. | [optional] [readonly] 
**OperFec** | Pointer to **string** | Operational FEC mode for this port. * &#x60;&#x60; - Default value for FEC state. * &#x60;Not Supported&#x60; - FEC is not supported on this port. * &#x60;Disabled&#x60; - FEC is disabled on this port. * &#x60;Auto&#x60; - FEC mode is automatically negotiated between link partners. * &#x60;Cl74&#x60; - FEC is configured to use the IEEE Clause 74 (FireCode) standard. * &#x60;RS-IEEE(Cl108)&#x60; - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard. * &#x60;KP&#x60; - FEC is configured to use the KP (Backplane Ethernet) FEC mode. | [optional] [readonly] [default to ""]
**OperSpeed** | Pointer to **string** | Current Operational speed for this port. | [optional] [readonly] 
**OperVlans** | Pointer to **string** | Operational VLANs on this port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for ethernet physical port. | [optional] [readonly] 
**PortName** | Pointer to **string** | Switch physical port name. | [optional] [readonly] 
**Speed** | Pointer to **string** | Speed of the inter switch link. | [optional] [readonly] 
**EquipmentSwitchCard** | Pointer to [**NullableEquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**NullableEtherInterSwitchPortRelationship**](EtherInterSwitchPortRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherInterSwitchPort

`func NewEtherInterSwitchPort(classId string, objectType string, ) *EtherInterSwitchPort`

NewEtherInterSwitchPort instantiates a new EtherInterSwitchPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherInterSwitchPortWithDefaults

`func NewEtherInterSwitchPortWithDefaults() *EtherInterSwitchPort`

NewEtherInterSwitchPortWithDefaults instantiates a new EtherInterSwitchPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherInterSwitchPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherInterSwitchPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherInterSwitchPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherInterSwitchPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherInterSwitchPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherInterSwitchPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessVlan

`func (o *EtherInterSwitchPort) GetAccessVlan() string`

GetAccessVlan returns the AccessVlan field if non-nil, zero value otherwise.

### GetAccessVlanOk

`func (o *EtherInterSwitchPort) GetAccessVlanOk() (*string, bool)`

GetAccessVlanOk returns a tuple with the AccessVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlan

`func (o *EtherInterSwitchPort) SetAccessVlan(v string)`

SetAccessVlan sets AccessVlan field to given value.

### HasAccessVlan

`func (o *EtherInterSwitchPort) HasAccessVlan() bool`

HasAccessVlan returns a boolean if a field has been set.

### GetAdminFec

`func (o *EtherInterSwitchPort) GetAdminFec() string`

GetAdminFec returns the AdminFec field if non-nil, zero value otherwise.

### GetAdminFecOk

`func (o *EtherInterSwitchPort) GetAdminFecOk() (*string, bool)`

GetAdminFecOk returns a tuple with the AdminFec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFec

`func (o *EtherInterSwitchPort) SetAdminFec(v string)`

SetAdminFec sets AdminFec field to given value.

### HasAdminFec

`func (o *EtherInterSwitchPort) HasAdminFec() bool`

HasAdminFec returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherInterSwitchPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherInterSwitchPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherInterSwitchPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherInterSwitchPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *EtherInterSwitchPort) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *EtherInterSwitchPort) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *EtherInterSwitchPort) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *EtherInterSwitchPort) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetMacAddress

`func (o *EtherInterSwitchPort) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherInterSwitchPort) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherInterSwitchPort) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherInterSwitchPort) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherInterSwitchPort) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherInterSwitchPort) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherInterSwitchPort) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherInterSwitchPort) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNativeVlan

`func (o *EtherInterSwitchPort) GetNativeVlan() string`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *EtherInterSwitchPort) GetNativeVlanOk() (*string, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *EtherInterSwitchPort) SetNativeVlan(v string)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *EtherInterSwitchPort) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetOperFec

`func (o *EtherInterSwitchPort) GetOperFec() string`

GetOperFec returns the OperFec field if non-nil, zero value otherwise.

### GetOperFecOk

`func (o *EtherInterSwitchPort) GetOperFecOk() (*string, bool)`

GetOperFecOk returns a tuple with the OperFec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperFec

`func (o *EtherInterSwitchPort) SetOperFec(v string)`

SetOperFec sets OperFec field to given value.

### HasOperFec

`func (o *EtherInterSwitchPort) HasOperFec() bool`

HasOperFec returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherInterSwitchPort) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherInterSwitchPort) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherInterSwitchPort) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherInterSwitchPort) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperVlans

`func (o *EtherInterSwitchPort) GetOperVlans() string`

GetOperVlans returns the OperVlans field if non-nil, zero value otherwise.

### GetOperVlansOk

`func (o *EtherInterSwitchPort) GetOperVlansOk() (*string, bool)`

GetOperVlansOk returns a tuple with the OperVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperVlans

`func (o *EtherInterSwitchPort) SetOperVlans(v string)`

SetOperVlans sets OperVlans field to given value.

### HasOperVlans

`func (o *EtherInterSwitchPort) HasOperVlans() bool`

HasOperVlans returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherInterSwitchPort) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherInterSwitchPort) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherInterSwitchPort) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherInterSwitchPort) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortName

`func (o *EtherInterSwitchPort) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *EtherInterSwitchPort) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *EtherInterSwitchPort) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *EtherInterSwitchPort) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetSpeed

`func (o *EtherInterSwitchPort) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EtherInterSwitchPort) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EtherInterSwitchPort) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EtherInterSwitchPort) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *EtherInterSwitchPort) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *EtherInterSwitchPort) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *EtherInterSwitchPort) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *EtherInterSwitchPort) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### SetEquipmentSwitchCardNil

`func (o *EtherInterSwitchPort) SetEquipmentSwitchCardNil(b bool)`

 SetEquipmentSwitchCardNil sets the value for EquipmentSwitchCard to be an explicit nil

### UnsetEquipmentSwitchCard
`func (o *EtherInterSwitchPort) UnsetEquipmentSwitchCard()`

UnsetEquipmentSwitchCard ensures that no value is present for EquipmentSwitchCard, not even an explicit nil
### GetPeerInterface

`func (o *EtherInterSwitchPort) GetPeerInterface() EtherInterSwitchPortRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherInterSwitchPort) GetPeerInterfaceOk() (*EtherInterSwitchPortRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherInterSwitchPort) SetPeerInterface(v EtherInterSwitchPortRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherInterSwitchPort) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### SetPeerInterfaceNil

`func (o *EtherInterSwitchPort) SetPeerInterfaceNil(b bool)`

 SetPeerInterfaceNil sets the value for PeerInterface to be an explicit nil

### UnsetPeerInterface
`func (o *EtherInterSwitchPort) UnsetPeerInterface()`

UnsetPeerInterface ensures that no value is present for PeerInterface, not even an explicit nil
### GetRegisteredDevice

`func (o *EtherInterSwitchPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherInterSwitchPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherInterSwitchPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherInterSwitchPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EtherInterSwitchPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EtherInterSwitchPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


