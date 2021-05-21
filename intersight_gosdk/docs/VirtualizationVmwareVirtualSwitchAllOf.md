# VirtualizationVmwareVirtualSwitchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualSwitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualSwitch"]
**ForgedTransmits** | Pointer to **string** | If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, the switch does not perform filtering and permits all outbound frames. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**MacAddressChanges** | Pointer to **string** | If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, the switch drops all inbound frames to the adapter. If property value is set to accept and the MAC address is changed, the switch allows frames to the new MAC address to pass. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**Mtu** | Pointer to **int64** | Maximum transmission unit configured on a virtual switch. | [optional] 
**NicTeamingAndFailover** | Pointer to [**NullableVirtualizationVmwareTeamingAndFailover**](virtualization.VmwareTeamingAndFailover.md) |  | [optional] 
**NumNetworks** | Pointer to **int64** | Number of networks available on this virtual switch. | [optional] 
**NumPhysicalNetworkInterfaces** | Pointer to **int64** | Number of physical network interfaces connected with this virtual switch. | [optional] 
**PromiscuousMode** | Pointer to **string** | If promiscuousMode property value is set to reject, the virtual switch forwards only frames that are addressed to the adapter. If property value is set to accept, the virtual switch forwards all frames to the adapter in compliance with the active VLAN policy for the port to which it is connected. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualSwitchAllOf

`func NewVirtualizationVmwareVirtualSwitchAllOf(classId string, objectType string, ) *VirtualizationVmwareVirtualSwitchAllOf`

NewVirtualizationVmwareVirtualSwitchAllOf instantiates a new VirtualizationVmwareVirtualSwitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualSwitchAllOfWithDefaults

`func NewVirtualizationVmwareVirtualSwitchAllOfWithDefaults() *VirtualizationVmwareVirtualSwitchAllOf`

NewVirtualizationVmwareVirtualSwitchAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualSwitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetForgedTransmits() string`

GetForgedTransmits returns the ForgedTransmits field if non-nil, zero value otherwise.

### GetForgedTransmitsOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetForgedTransmitsOk() (*string, bool)`

GetForgedTransmitsOk returns a tuple with the ForgedTransmits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetForgedTransmits(v string)`

SetForgedTransmits sets ForgedTransmits field to given value.

### HasForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasForgedTransmits() bool`

HasForgedTransmits returns a boolean if a field has been set.

### GetMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMacAddressChanges() string`

GetMacAddressChanges returns the MacAddressChanges field if non-nil, zero value otherwise.

### GetMacAddressChangesOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMacAddressChangesOk() (*string, bool)`

GetMacAddressChangesOk returns a tuple with the MacAddressChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetMacAddressChanges(v string)`

SetMacAddressChanges sets MacAddressChanges field to given value.

### HasMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasMacAddressChanges() bool`

HasMacAddressChanges returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareVirtualSwitchAllOf) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetNumNetworks

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumNetworks() int64`

GetNumNetworks returns the NumNetworks field if non-nil, zero value otherwise.

### GetNumNetworksOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumNetworksOk() (*int64, bool)`

GetNumNetworksOk returns a tuple with the NumNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNetworks

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNumNetworks(v int64)`

SetNumNetworks sets NumNetworks field to given value.

### HasNumNetworks

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNumNetworks() bool`

HasNumNetworks returns a boolean if a field has been set.

### GetNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumPhysicalNetworkInterfaces() int64`

GetNumPhysicalNetworkInterfaces returns the NumPhysicalNetworkInterfaces field if non-nil, zero value otherwise.

### GetNumPhysicalNetworkInterfacesOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumPhysicalNetworkInterfacesOk() (*int64, bool)`

GetNumPhysicalNetworkInterfacesOk returns a tuple with the NumPhysicalNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNumPhysicalNetworkInterfaces(v int64)`

SetNumPhysicalNetworkInterfaces sets NumPhysicalNetworkInterfaces field to given value.

### HasNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNumPhysicalNetworkInterfaces() bool`

HasNumPhysicalNetworkInterfaces returns a boolean if a field has been set.

### GetPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetPromiscuousMode() string`

GetPromiscuousMode returns the PromiscuousMode field if non-nil, zero value otherwise.

### GetPromiscuousModeOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetPromiscuousModeOk() (*string, bool)`

GetPromiscuousModeOk returns a tuple with the PromiscuousMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetPromiscuousMode(v string)`

SetPromiscuousMode sets PromiscuousMode field to given value.

### HasPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasPromiscuousMode() bool`

HasPromiscuousMode returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareVirtualSwitchAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareVirtualSwitchAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareVirtualSwitchAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


