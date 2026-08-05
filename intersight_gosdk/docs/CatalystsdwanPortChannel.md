# CatalystsdwanPortChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.PortChannel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.PortChannel"]
**AdminState** | Pointer to **string** | The administrative state of the WAN Edge Device port channel. * &#x60;Unknown&#x60; - Administrative state is unknown. * &#x60;Up&#x60; - Administrative state is up. * &#x60;Down&#x60; - Administrative state is down. | [optional] [readonly] [default to "Unknown"]
**AllowedVlanIds** | Pointer to **string** | The allowed VLAN IDs configured on the WAN Edge Device port channel. | [optional] [readonly] 
**DefaultVlanId** | Pointer to **int64** | The default VLAN ID configured on the WAN Edge Device port channel. | [optional] [readonly] [default to 1]
**Dn** | Pointer to **string** | The distinguished name of the port channel. | [optional] [readonly] 
**Mode** | Pointer to **string** | Portchannel mode on WAN Edge Device will be &#39;access&#39; or &#39;trunk&#39; when switchport mode is enabled; otherwise, the mode is &#39;routed&#39;. * &#x60;Routed&#x60; - Port mode is routed. This mode will be set when switchport mode is disabled regardless of the port configuration. * &#x60;Access&#x60; - Port mode is access. This mode will be set when switchport mode is enabled and the port is configured as an access port. * &#x60;Trunk&#x60; - Port mode is trunk. This mode will be set when switchport mode is enabled and the port is configured as a trunk port. | [optional] [readonly] [default to "Routed"]
**Mtu** | Pointer to **int64** | The MTU of the WAN Edge Device port channel. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | The operational speed (in Mbps) of the WAN Edge Device port channel. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the WAN Edge Device port channel. * &#x60;Unknown&#x60; - Operational state of physical port or port channel or interface is unknown. * &#x60;Up&#x60; - Operational state of physical port or port channel or interface is up. * &#x60;Down&#x60; - Operational state of physical port or port channel or interface is down. | [optional] [readonly] [default to "Unknown"]
**PortChannelName** | Pointer to **string** | Port channel name of the WAN Edge Device. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | The IP address of the WAN Edge Device. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID of the WAN Edge Device to which this port channel belongs. | [optional] [readonly] 
**WanEdgeDevice** | Pointer to [**NullableCatalystsdwanWanEdgeDeviceRelationship**](CatalystsdwanWanEdgeDeviceRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanPortChannel

`func NewCatalystsdwanPortChannel(classId string, objectType string, ) *CatalystsdwanPortChannel`

NewCatalystsdwanPortChannel instantiates a new CatalystsdwanPortChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanPortChannelWithDefaults

`func NewCatalystsdwanPortChannelWithDefaults() *CatalystsdwanPortChannel`

NewCatalystsdwanPortChannelWithDefaults instantiates a new CatalystsdwanPortChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanPortChannel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanPortChannel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanPortChannel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanPortChannel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanPortChannel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanPortChannel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *CatalystsdwanPortChannel) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CatalystsdwanPortChannel) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CatalystsdwanPortChannel) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CatalystsdwanPortChannel) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlanIds

`func (o *CatalystsdwanPortChannel) GetAllowedVlanIds() string`

GetAllowedVlanIds returns the AllowedVlanIds field if non-nil, zero value otherwise.

### GetAllowedVlanIdsOk

`func (o *CatalystsdwanPortChannel) GetAllowedVlanIdsOk() (*string, bool)`

GetAllowedVlanIdsOk returns a tuple with the AllowedVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlanIds

`func (o *CatalystsdwanPortChannel) SetAllowedVlanIds(v string)`

SetAllowedVlanIds sets AllowedVlanIds field to given value.

### HasAllowedVlanIds

`func (o *CatalystsdwanPortChannel) HasAllowedVlanIds() bool`

HasAllowedVlanIds returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *CatalystsdwanPortChannel) GetDefaultVlanId() int64`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *CatalystsdwanPortChannel) GetDefaultVlanIdOk() (*int64, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *CatalystsdwanPortChannel) SetDefaultVlanId(v int64)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *CatalystsdwanPortChannel) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetDn

`func (o *CatalystsdwanPortChannel) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CatalystsdwanPortChannel) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CatalystsdwanPortChannel) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *CatalystsdwanPortChannel) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetMode

`func (o *CatalystsdwanPortChannel) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CatalystsdwanPortChannel) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CatalystsdwanPortChannel) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CatalystsdwanPortChannel) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *CatalystsdwanPortChannel) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CatalystsdwanPortChannel) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CatalystsdwanPortChannel) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CatalystsdwanPortChannel) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetOperSpeed

`func (o *CatalystsdwanPortChannel) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *CatalystsdwanPortChannel) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *CatalystsdwanPortChannel) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *CatalystsdwanPortChannel) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *CatalystsdwanPortChannel) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *CatalystsdwanPortChannel) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *CatalystsdwanPortChannel) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *CatalystsdwanPortChannel) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPortChannelName

`func (o *CatalystsdwanPortChannel) GetPortChannelName() string`

GetPortChannelName returns the PortChannelName field if non-nil, zero value otherwise.

### GetPortChannelNameOk

`func (o *CatalystsdwanPortChannel) GetPortChannelNameOk() (*string, bool)`

GetPortChannelNameOk returns a tuple with the PortChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelName

`func (o *CatalystsdwanPortChannel) SetPortChannelName(v string)`

SetPortChannelName sets PortChannelName field to given value.

### HasPortChannelName

`func (o *CatalystsdwanPortChannel) HasPortChannelName() bool`

HasPortChannelName returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanPortChannel) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanPortChannel) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanPortChannel) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanPortChannel) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetUuid

`func (o *CatalystsdwanPortChannel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CatalystsdwanPortChannel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CatalystsdwanPortChannel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CatalystsdwanPortChannel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWanEdgeDevice

`func (o *CatalystsdwanPortChannel) GetWanEdgeDevice() CatalystsdwanWanEdgeDeviceRelationship`

GetWanEdgeDevice returns the WanEdgeDevice field if non-nil, zero value otherwise.

### GetWanEdgeDeviceOk

`func (o *CatalystsdwanPortChannel) GetWanEdgeDeviceOk() (*CatalystsdwanWanEdgeDeviceRelationship, bool)`

GetWanEdgeDeviceOk returns a tuple with the WanEdgeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEdgeDevice

`func (o *CatalystsdwanPortChannel) SetWanEdgeDevice(v CatalystsdwanWanEdgeDeviceRelationship)`

SetWanEdgeDevice sets WanEdgeDevice field to given value.

### HasWanEdgeDevice

`func (o *CatalystsdwanPortChannel) HasWanEdgeDevice() bool`

HasWanEdgeDevice returns a boolean if a field has been set.

### SetWanEdgeDeviceNil

`func (o *CatalystsdwanPortChannel) SetWanEdgeDeviceNil(b bool)`

 SetWanEdgeDeviceNil sets the value for WanEdgeDevice to be an explicit nil

### UnsetWanEdgeDevice
`func (o *CatalystsdwanPortChannel) UnsetWanEdgeDevice()`

UnsetWanEdgeDevice ensures that no value is present for WanEdgeDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


