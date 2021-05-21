# VirtualizationVmwareNetworkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareNetwork"]
**ForgedTransmits** | Pointer to **string** | If forgedTransmits property value is set to reject, outbound frames with source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**MacAddressChanges** | Pointer to **string** | If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**NicTeamingAndFailover** | Pointer to [**NullableVirtualizationVmwareTeamingAndFailover**](virtualization.VmwareTeamingAndFailover.md) |  | [optional] 
**PromiscuousMode** | Pointer to **string** | If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**VlanId** | Pointer to **int64** | VLAN id with which the network is associated. A value of 0 specifies that port is not associated with a VLAN. | [optional] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 
**VirtualSwitch** | Pointer to [**VirtualizationVmwareVirtualSwitchRelationship**](virtualization.VmwareVirtualSwitch.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareNetworkAllOf

`func NewVirtualizationVmwareNetworkAllOf(classId string, objectType string, ) *VirtualizationVmwareNetworkAllOf`

NewVirtualizationVmwareNetworkAllOf instantiates a new VirtualizationVmwareNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareNetworkAllOfWithDefaults

`func NewVirtualizationVmwareNetworkAllOfWithDefaults() *VirtualizationVmwareNetworkAllOf`

NewVirtualizationVmwareNetworkAllOfWithDefaults instantiates a new VirtualizationVmwareNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetForgedTransmits

`func (o *VirtualizationVmwareNetworkAllOf) GetForgedTransmits() string`

GetForgedTransmits returns the ForgedTransmits field if non-nil, zero value otherwise.

### GetForgedTransmitsOk

`func (o *VirtualizationVmwareNetworkAllOf) GetForgedTransmitsOk() (*string, bool)`

GetForgedTransmitsOk returns a tuple with the ForgedTransmits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgedTransmits

`func (o *VirtualizationVmwareNetworkAllOf) SetForgedTransmits(v string)`

SetForgedTransmits sets ForgedTransmits field to given value.

### HasForgedTransmits

`func (o *VirtualizationVmwareNetworkAllOf) HasForgedTransmits() bool`

HasForgedTransmits returns a boolean if a field has been set.

### GetMacAddressChanges

`func (o *VirtualizationVmwareNetworkAllOf) GetMacAddressChanges() string`

GetMacAddressChanges returns the MacAddressChanges field if non-nil, zero value otherwise.

### GetMacAddressChangesOk

`func (o *VirtualizationVmwareNetworkAllOf) GetMacAddressChangesOk() (*string, bool)`

GetMacAddressChangesOk returns a tuple with the MacAddressChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressChanges

`func (o *VirtualizationVmwareNetworkAllOf) SetMacAddressChanges(v string)`

SetMacAddressChanges sets MacAddressChanges field to given value.

### HasMacAddressChanges

`func (o *VirtualizationVmwareNetworkAllOf) HasMacAddressChanges() bool`

HasMacAddressChanges returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareNetworkAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareNetworkAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareNetworkAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareNetworkAllOf) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareNetworkAllOf) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareNetworkAllOf) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetPromiscuousMode

`func (o *VirtualizationVmwareNetworkAllOf) GetPromiscuousMode() string`

GetPromiscuousMode returns the PromiscuousMode field if non-nil, zero value otherwise.

### GetPromiscuousModeOk

`func (o *VirtualizationVmwareNetworkAllOf) GetPromiscuousModeOk() (*string, bool)`

GetPromiscuousModeOk returns a tuple with the PromiscuousMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromiscuousMode

`func (o *VirtualizationVmwareNetworkAllOf) SetPromiscuousMode(v string)`

SetPromiscuousMode sets PromiscuousMode field to given value.

### HasPromiscuousMode

`func (o *VirtualizationVmwareNetworkAllOf) HasPromiscuousMode() bool`

HasPromiscuousMode returns a boolean if a field has been set.

### GetVlanId

`func (o *VirtualizationVmwareNetworkAllOf) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VirtualizationVmwareNetworkAllOf) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VirtualizationVmwareNetworkAllOf) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *VirtualizationVmwareNetworkAllOf) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareNetworkAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareNetworkAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareNetworkAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareNetworkAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetVirtualSwitch

`func (o *VirtualizationVmwareNetworkAllOf) GetVirtualSwitch() VirtualizationVmwareVirtualSwitchRelationship`

GetVirtualSwitch returns the VirtualSwitch field if non-nil, zero value otherwise.

### GetVirtualSwitchOk

`func (o *VirtualizationVmwareNetworkAllOf) GetVirtualSwitchOk() (*VirtualizationVmwareVirtualSwitchRelationship, bool)`

GetVirtualSwitchOk returns a tuple with the VirtualSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitch

`func (o *VirtualizationVmwareNetworkAllOf) SetVirtualSwitch(v VirtualizationVmwareVirtualSwitchRelationship)`

SetVirtualSwitch sets VirtualSwitch field to given value.

### HasVirtualSwitch

`func (o *VirtualizationVmwareNetworkAllOf) HasVirtualSwitch() bool`

HasVirtualSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


