# CatalystsdwanInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.Interface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.Interface"]
**AdminState** | Pointer to **string** | The administrative state of the WAN Edge Device interface. * &#x60;Unknown&#x60; - Administrative state is unknown. * &#x60;Up&#x60; - Administrative state is up. * &#x60;Down&#x60; - Administrative state is down. | [optional] [readonly] [default to "Unknown"]
**Dn** | Pointer to **string** | The distinguished name of the interface. | [optional] [readonly] 
**InterfaceName** | Pointer to **string** | The name of the interface on the WAN Edge Device. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | The type of the interface (Portchannel, Sub-interfaces, tunnels, SVIs, Routed ports). * &#x60;Unknown&#x60; - Interface type is unknown. * &#x60;Loopback&#x60; - Interface type is loopback. * &#x60;Port Channel&#x60; - Interface type is port channel. * &#x60;Sub-Interface&#x60; - Interface type is sub-interface. * &#x60;Tunnel&#x60; - Interface type is tunnel. * &#x60;SVI&#x60; - Interface type is SVI (Switched Virtual Interface). * &#x60;Routed Port&#x60; - Interface type is routed port. | [optional] [readonly] [default to "Unknown"]
**IpAddress** | Pointer to **string** | The IP address assigned to the WAN Edge Device interface. | [optional] [readonly] 
**IsNativeVlan** | Pointer to **bool** | Indicates whether the sub-interface&#39;s VLAN is configured as a native VLAN. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address of the interface. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the WAN Edge Device interface. * &#x60;Unknown&#x60; - Operational state of physical port or port channel or interface is unknown. * &#x60;Up&#x60; - Operational state of physical port or port channel or interface is up. * &#x60;Down&#x60; - Operational state of physical port or port channel or interface is down. | [optional] [readonly] [default to "Unknown"]
**Subnet** | Pointer to **string** | The subnet configured on the WAN Edge Device interface. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | The system IP address of the WAN Edge Device. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID of the WAN Edge Device to which this interface belongs. | [optional] [readonly] 
**VlanId** | Pointer to **int64** | The VLAN ID configured on the sub-interface. | [optional] [readonly] [default to 1]
**VpnId** | Pointer to **int64** | The VPN ID associated with this interface. | [optional] [readonly] 
**WanEdgeDevice** | Pointer to [**NullableCatalystsdwanWanEdgeDeviceRelationship**](CatalystsdwanWanEdgeDeviceRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanInterface

`func NewCatalystsdwanInterface(classId string, objectType string, ) *CatalystsdwanInterface`

NewCatalystsdwanInterface instantiates a new CatalystsdwanInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanInterfaceWithDefaults

`func NewCatalystsdwanInterfaceWithDefaults() *CatalystsdwanInterface`

NewCatalystsdwanInterfaceWithDefaults instantiates a new CatalystsdwanInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *CatalystsdwanInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CatalystsdwanInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CatalystsdwanInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CatalystsdwanInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetDn

`func (o *CatalystsdwanInterface) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CatalystsdwanInterface) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CatalystsdwanInterface) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *CatalystsdwanInterface) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetInterfaceName

`func (o *CatalystsdwanInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *CatalystsdwanInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *CatalystsdwanInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *CatalystsdwanInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *CatalystsdwanInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *CatalystsdwanInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *CatalystsdwanInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *CatalystsdwanInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetIpAddress

`func (o *CatalystsdwanInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CatalystsdwanInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CatalystsdwanInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CatalystsdwanInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsNativeVlan

`func (o *CatalystsdwanInterface) GetIsNativeVlan() bool`

GetIsNativeVlan returns the IsNativeVlan field if non-nil, zero value otherwise.

### GetIsNativeVlanOk

`func (o *CatalystsdwanInterface) GetIsNativeVlanOk() (*bool, bool)`

GetIsNativeVlanOk returns a tuple with the IsNativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeVlan

`func (o *CatalystsdwanInterface) SetIsNativeVlan(v bool)`

SetIsNativeVlan sets IsNativeVlan field to given value.

### HasIsNativeVlan

`func (o *CatalystsdwanInterface) HasIsNativeVlan() bool`

HasIsNativeVlan returns a boolean if a field has been set.

### GetMacAddress

`func (o *CatalystsdwanInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CatalystsdwanInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CatalystsdwanInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CatalystsdwanInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetOperState

`func (o *CatalystsdwanInterface) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *CatalystsdwanInterface) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *CatalystsdwanInterface) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *CatalystsdwanInterface) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetSubnet

`func (o *CatalystsdwanInterface) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CatalystsdwanInterface) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CatalystsdwanInterface) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CatalystsdwanInterface) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanInterface) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanInterface) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanInterface) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanInterface) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetUuid

`func (o *CatalystsdwanInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CatalystsdwanInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CatalystsdwanInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CatalystsdwanInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVlanId

`func (o *CatalystsdwanInterface) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CatalystsdwanInterface) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CatalystsdwanInterface) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CatalystsdwanInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVpnId

`func (o *CatalystsdwanInterface) GetVpnId() int64`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *CatalystsdwanInterface) GetVpnIdOk() (*int64, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *CatalystsdwanInterface) SetVpnId(v int64)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *CatalystsdwanInterface) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetWanEdgeDevice

`func (o *CatalystsdwanInterface) GetWanEdgeDevice() CatalystsdwanWanEdgeDeviceRelationship`

GetWanEdgeDevice returns the WanEdgeDevice field if non-nil, zero value otherwise.

### GetWanEdgeDeviceOk

`func (o *CatalystsdwanInterface) GetWanEdgeDeviceOk() (*CatalystsdwanWanEdgeDeviceRelationship, bool)`

GetWanEdgeDeviceOk returns a tuple with the WanEdgeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEdgeDevice

`func (o *CatalystsdwanInterface) SetWanEdgeDevice(v CatalystsdwanWanEdgeDeviceRelationship)`

SetWanEdgeDevice sets WanEdgeDevice field to given value.

### HasWanEdgeDevice

`func (o *CatalystsdwanInterface) HasWanEdgeDevice() bool`

HasWanEdgeDevice returns a boolean if a field has been set.

### SetWanEdgeDeviceNil

`func (o *CatalystsdwanInterface) SetWanEdgeDeviceNil(b bool)`

 SetWanEdgeDeviceNil sets the value for WanEdgeDevice to be an explicit nil

### UnsetWanEdgeDevice
`func (o *CatalystsdwanInterface) UnsetWanEdgeDevice()`

UnsetWanEdgeDevice ensures that no value is present for WanEdgeDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


