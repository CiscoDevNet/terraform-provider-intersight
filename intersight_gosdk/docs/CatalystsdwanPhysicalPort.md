# CatalystsdwanPhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.PhysicalPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.PhysicalPort"]
**AdminSpeed** | Pointer to **string** | Admin speed of the physical port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administrative state of the physical port. * &#x60;Unknown&#x60; - Administrative state is unknown. * &#x60;Up&#x60; - Administrative state is up. * &#x60;Down&#x60; - Administrative state is down. | [optional] [readonly] [default to "Unknown"]
**AllowedVlanIds** | Pointer to **string** | List of configured allowed VLAN IDs. | [optional] [readonly] 
**AutoNegotiation** | Pointer to **bool** | Indicates if auto-negotiation is enabled. | [optional] [readonly] 
**DefaultVlanId** | Pointer to **int64** | Default or Native VLAN configured on the port. | [optional] [readonly] [default to 1]
**Dn** | Pointer to **string** | The distinguished name of the physical port. | [optional] [readonly] 
**FecEnabled** | Pointer to **bool** | Indicates if FEC is enabled on the physical port. | [optional] [readonly] [default to false]
**MacAddress** | Pointer to **string** | MAC address of the physical port. | [optional] [readonly] 
**Mode** | Pointer to **string** | Configured physical port mode. Value will be &#39;access&#39; or &#39;trunk&#39; when configured as a switch port; &#39;backplane&#39; for internal backplane ports and &#39;routed&#39; all other ports. * &#x60;Routed&#x60; - Port mode is routed. This mode will be set when switchport mode is disabled regardless of the port configuration. * &#x60;Access&#x60; - Port mode is access. This mode will be set when switchport mode is enabled and the port is configured as an access port. * &#x60;Trunk&#x60; - Port mode is trunk. This mode will be set when switchport mode is enabled and the port is configured as a trunk port. | [optional] [readonly] [default to "Routed"]
**Mtu** | Pointer to **int64** | MTU configured on the physical port. | [optional] [readonly] 
**NegotiatedSpeed** | Pointer to **string** | Negotiated speed of the physical port. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the physical port. * &#x60;Unknown&#x60; - Operational state of physical port or port channel or interface is unknown. * &#x60;Up&#x60; - Operational state of physical port or port channel or interface is up. * &#x60;Down&#x60; - Operational state of physical port or port channel or interface is down. | [optional] [readonly] [default to "Unknown"]
**PortChannelId** | Pointer to **int64** | Port channel ID if the physical port is part of a port channel. | [optional] [readonly] [default to 0]
**PortChannelName** | Pointer to **string** | Port channel name if the physical port is part of a port channel. | [optional] [readonly] 
**PortName** | Pointer to **string** | Physical port name of the WAN Edge Device. | [optional] [readonly] 
**PortType** | Pointer to **string** | Type of the physical port (e.g., Backplane and frontpanel). * &#x60;Front Panel&#x60; - The physical ports that are located on the front panel of Catalyst SDWAN devices. * &#x60;Backplane&#x60; - The internal backplane ports on Catalyst SDWAN devices. | [optional] [readonly] [default to "Front Panel"]
**SystemIp** | Pointer to **string** | The IP address of the WAN Edge Device. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID of the WAN Edge Device to which this physical port belongs. | [optional] [readonly] 
**WanEdgeDevice** | Pointer to [**NullableCatalystsdwanWanEdgeDeviceRelationship**](CatalystsdwanWanEdgeDeviceRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanPhysicalPort

`func NewCatalystsdwanPhysicalPort(classId string, objectType string, ) *CatalystsdwanPhysicalPort`

NewCatalystsdwanPhysicalPort instantiates a new CatalystsdwanPhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanPhysicalPortWithDefaults

`func NewCatalystsdwanPhysicalPortWithDefaults() *CatalystsdwanPhysicalPort`

NewCatalystsdwanPhysicalPortWithDefaults instantiates a new CatalystsdwanPhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanPhysicalPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanPhysicalPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanPhysicalPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanPhysicalPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanPhysicalPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanPhysicalPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *CatalystsdwanPhysicalPort) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *CatalystsdwanPhysicalPort) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *CatalystsdwanPhysicalPort) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *CatalystsdwanPhysicalPort) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *CatalystsdwanPhysicalPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CatalystsdwanPhysicalPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CatalystsdwanPhysicalPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CatalystsdwanPhysicalPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlanIds

`func (o *CatalystsdwanPhysicalPort) GetAllowedVlanIds() string`

GetAllowedVlanIds returns the AllowedVlanIds field if non-nil, zero value otherwise.

### GetAllowedVlanIdsOk

`func (o *CatalystsdwanPhysicalPort) GetAllowedVlanIdsOk() (*string, bool)`

GetAllowedVlanIdsOk returns a tuple with the AllowedVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlanIds

`func (o *CatalystsdwanPhysicalPort) SetAllowedVlanIds(v string)`

SetAllowedVlanIds sets AllowedVlanIds field to given value.

### HasAllowedVlanIds

`func (o *CatalystsdwanPhysicalPort) HasAllowedVlanIds() bool`

HasAllowedVlanIds returns a boolean if a field has been set.

### GetAutoNegotiation

`func (o *CatalystsdwanPhysicalPort) GetAutoNegotiation() bool`

GetAutoNegotiation returns the AutoNegotiation field if non-nil, zero value otherwise.

### GetAutoNegotiationOk

`func (o *CatalystsdwanPhysicalPort) GetAutoNegotiationOk() (*bool, bool)`

GetAutoNegotiationOk returns a tuple with the AutoNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiation

`func (o *CatalystsdwanPhysicalPort) SetAutoNegotiation(v bool)`

SetAutoNegotiation sets AutoNegotiation field to given value.

### HasAutoNegotiation

`func (o *CatalystsdwanPhysicalPort) HasAutoNegotiation() bool`

HasAutoNegotiation returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *CatalystsdwanPhysicalPort) GetDefaultVlanId() int64`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *CatalystsdwanPhysicalPort) GetDefaultVlanIdOk() (*int64, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *CatalystsdwanPhysicalPort) SetDefaultVlanId(v int64)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *CatalystsdwanPhysicalPort) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetDn

`func (o *CatalystsdwanPhysicalPort) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CatalystsdwanPhysicalPort) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CatalystsdwanPhysicalPort) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *CatalystsdwanPhysicalPort) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFecEnabled

`func (o *CatalystsdwanPhysicalPort) GetFecEnabled() bool`

GetFecEnabled returns the FecEnabled field if non-nil, zero value otherwise.

### GetFecEnabledOk

`func (o *CatalystsdwanPhysicalPort) GetFecEnabledOk() (*bool, bool)`

GetFecEnabledOk returns a tuple with the FecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecEnabled

`func (o *CatalystsdwanPhysicalPort) SetFecEnabled(v bool)`

SetFecEnabled sets FecEnabled field to given value.

### HasFecEnabled

`func (o *CatalystsdwanPhysicalPort) HasFecEnabled() bool`

HasFecEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *CatalystsdwanPhysicalPort) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CatalystsdwanPhysicalPort) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CatalystsdwanPhysicalPort) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CatalystsdwanPhysicalPort) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *CatalystsdwanPhysicalPort) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CatalystsdwanPhysicalPort) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CatalystsdwanPhysicalPort) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CatalystsdwanPhysicalPort) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *CatalystsdwanPhysicalPort) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CatalystsdwanPhysicalPort) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CatalystsdwanPhysicalPort) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CatalystsdwanPhysicalPort) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNegotiatedSpeed

`func (o *CatalystsdwanPhysicalPort) GetNegotiatedSpeed() string`

GetNegotiatedSpeed returns the NegotiatedSpeed field if non-nil, zero value otherwise.

### GetNegotiatedSpeedOk

`func (o *CatalystsdwanPhysicalPort) GetNegotiatedSpeedOk() (*string, bool)`

GetNegotiatedSpeedOk returns a tuple with the NegotiatedSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiatedSpeed

`func (o *CatalystsdwanPhysicalPort) SetNegotiatedSpeed(v string)`

SetNegotiatedSpeed sets NegotiatedSpeed field to given value.

### HasNegotiatedSpeed

`func (o *CatalystsdwanPhysicalPort) HasNegotiatedSpeed() bool`

HasNegotiatedSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *CatalystsdwanPhysicalPort) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *CatalystsdwanPhysicalPort) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *CatalystsdwanPhysicalPort) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *CatalystsdwanPhysicalPort) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPortChannelId

`func (o *CatalystsdwanPhysicalPort) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *CatalystsdwanPhysicalPort) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *CatalystsdwanPhysicalPort) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *CatalystsdwanPhysicalPort) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetPortChannelName

`func (o *CatalystsdwanPhysicalPort) GetPortChannelName() string`

GetPortChannelName returns the PortChannelName field if non-nil, zero value otherwise.

### GetPortChannelNameOk

`func (o *CatalystsdwanPhysicalPort) GetPortChannelNameOk() (*string, bool)`

GetPortChannelNameOk returns a tuple with the PortChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelName

`func (o *CatalystsdwanPhysicalPort) SetPortChannelName(v string)`

SetPortChannelName sets PortChannelName field to given value.

### HasPortChannelName

`func (o *CatalystsdwanPhysicalPort) HasPortChannelName() bool`

HasPortChannelName returns a boolean if a field has been set.

### GetPortName

`func (o *CatalystsdwanPhysicalPort) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *CatalystsdwanPhysicalPort) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *CatalystsdwanPhysicalPort) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *CatalystsdwanPhysicalPort) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortType

`func (o *CatalystsdwanPhysicalPort) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *CatalystsdwanPhysicalPort) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *CatalystsdwanPhysicalPort) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *CatalystsdwanPhysicalPort) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanPhysicalPort) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanPhysicalPort) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanPhysicalPort) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanPhysicalPort) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetUuid

`func (o *CatalystsdwanPhysicalPort) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CatalystsdwanPhysicalPort) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CatalystsdwanPhysicalPort) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CatalystsdwanPhysicalPort) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWanEdgeDevice

`func (o *CatalystsdwanPhysicalPort) GetWanEdgeDevice() CatalystsdwanWanEdgeDeviceRelationship`

GetWanEdgeDevice returns the WanEdgeDevice field if non-nil, zero value otherwise.

### GetWanEdgeDeviceOk

`func (o *CatalystsdwanPhysicalPort) GetWanEdgeDeviceOk() (*CatalystsdwanWanEdgeDeviceRelationship, bool)`

GetWanEdgeDeviceOk returns a tuple with the WanEdgeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEdgeDevice

`func (o *CatalystsdwanPhysicalPort) SetWanEdgeDevice(v CatalystsdwanWanEdgeDeviceRelationship)`

SetWanEdgeDevice sets WanEdgeDevice field to given value.

### HasWanEdgeDevice

`func (o *CatalystsdwanPhysicalPort) HasWanEdgeDevice() bool`

HasWanEdgeDevice returns a boolean if a field has been set.

### SetWanEdgeDeviceNil

`func (o *CatalystsdwanPhysicalPort) SetWanEdgeDeviceNil(b bool)`

 SetWanEdgeDeviceNil sets the value for WanEdgeDevice to be an explicit nil

### UnsetWanEdgeDevice
`func (o *CatalystsdwanPhysicalPort) UnsetWanEdgeDevice()`

UnsetWanEdgeDevice ensures that no value is present for WanEdgeDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


