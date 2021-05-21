# VirtualizationVmwareVirtualSwitch

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

### NewVirtualizationVmwareVirtualSwitch

`func NewVirtualizationVmwareVirtualSwitch(classId string, objectType string, ) *VirtualizationVmwareVirtualSwitch`

NewVirtualizationVmwareVirtualSwitch instantiates a new VirtualizationVmwareVirtualSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualSwitchWithDefaults

`func NewVirtualizationVmwareVirtualSwitchWithDefaults() *VirtualizationVmwareVirtualSwitch`

NewVirtualizationVmwareVirtualSwitchWithDefaults instantiates a new VirtualizationVmwareVirtualSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualSwitch) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualSwitch) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualSwitch) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualSwitch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualSwitch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualSwitch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitch) GetForgedTransmits() string`

GetForgedTransmits returns the ForgedTransmits field if non-nil, zero value otherwise.

### GetForgedTransmitsOk

`func (o *VirtualizationVmwareVirtualSwitch) GetForgedTransmitsOk() (*string, bool)`

GetForgedTransmitsOk returns a tuple with the ForgedTransmits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitch) SetForgedTransmits(v string)`

SetForgedTransmits sets ForgedTransmits field to given value.

### HasForgedTransmits

`func (o *VirtualizationVmwareVirtualSwitch) HasForgedTransmits() bool`

HasForgedTransmits returns a boolean if a field has been set.

### GetMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitch) GetMacAddressChanges() string`

GetMacAddressChanges returns the MacAddressChanges field if non-nil, zero value otherwise.

### GetMacAddressChangesOk

`func (o *VirtualizationVmwareVirtualSwitch) GetMacAddressChangesOk() (*string, bool)`

GetMacAddressChangesOk returns a tuple with the MacAddressChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitch) SetMacAddressChanges(v string)`

SetMacAddressChanges sets MacAddressChanges field to given value.

### HasMacAddressChanges

`func (o *VirtualizationVmwareVirtualSwitch) HasMacAddressChanges() bool`

HasMacAddressChanges returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareVirtualSwitch) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareVirtualSwitch) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareVirtualSwitch) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareVirtualSwitch) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitch) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareVirtualSwitch) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitch) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareVirtualSwitch) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareVirtualSwitch) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareVirtualSwitch) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetNumNetworks

`func (o *VirtualizationVmwareVirtualSwitch) GetNumNetworks() int64`

GetNumNetworks returns the NumNetworks field if non-nil, zero value otherwise.

### GetNumNetworksOk

`func (o *VirtualizationVmwareVirtualSwitch) GetNumNetworksOk() (*int64, bool)`

GetNumNetworksOk returns a tuple with the NumNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNetworks

`func (o *VirtualizationVmwareVirtualSwitch) SetNumNetworks(v int64)`

SetNumNetworks sets NumNetworks field to given value.

### HasNumNetworks

`func (o *VirtualizationVmwareVirtualSwitch) HasNumNetworks() bool`

HasNumNetworks returns a boolean if a field has been set.

### GetNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitch) GetNumPhysicalNetworkInterfaces() int64`

GetNumPhysicalNetworkInterfaces returns the NumPhysicalNetworkInterfaces field if non-nil, zero value otherwise.

### GetNumPhysicalNetworkInterfacesOk

`func (o *VirtualizationVmwareVirtualSwitch) GetNumPhysicalNetworkInterfacesOk() (*int64, bool)`

GetNumPhysicalNetworkInterfacesOk returns a tuple with the NumPhysicalNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitch) SetNumPhysicalNetworkInterfaces(v int64)`

SetNumPhysicalNetworkInterfaces sets NumPhysicalNetworkInterfaces field to given value.

### HasNumPhysicalNetworkInterfaces

`func (o *VirtualizationVmwareVirtualSwitch) HasNumPhysicalNetworkInterfaces() bool`

HasNumPhysicalNetworkInterfaces returns a boolean if a field has been set.

### GetPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitch) GetPromiscuousMode() string`

GetPromiscuousMode returns the PromiscuousMode field if non-nil, zero value otherwise.

### GetPromiscuousModeOk

`func (o *VirtualizationVmwareVirtualSwitch) GetPromiscuousModeOk() (*string, bool)`

GetPromiscuousModeOk returns a tuple with the PromiscuousMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitch) SetPromiscuousMode(v string)`

SetPromiscuousMode sets PromiscuousMode field to given value.

### HasPromiscuousMode

`func (o *VirtualizationVmwareVirtualSwitch) HasPromiscuousMode() bool`

HasPromiscuousMode returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareVirtualSwitch) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareVirtualSwitch) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareVirtualSwitch) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareVirtualSwitch) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


