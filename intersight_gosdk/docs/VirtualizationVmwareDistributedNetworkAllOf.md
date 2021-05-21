# VirtualizationVmwareDistributedNetworkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDistributedNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDistributedNetwork"]
**ForgedTransmits** | Pointer to **string** | If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**MacAddressChanges** | Pointer to **string** | If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**NicTeamingAndFailover** | Pointer to [**NullableVirtualizationVmwareTeamingAndFailover**](virtualization.VmwareTeamingAndFailover.md) |  | [optional] 
**NumHosts** | Pointer to **int64** | The total number of hosts connected to this distributed virtual network. | [optional] 
**NumPorts** | Pointer to **int64** | The total number of ports in the distributed virtual network. | [optional] 
**PromiscuousMode** | Pointer to **string** | If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * &#x60;Reject&#x60; - Indicates that the security policy is rejected. * &#x60;Accept&#x60; - Indicates that the security policy is accepted. | [optional] [default to "Reject"]
**UpLink** | Pointer to **bool** | Indicates if the distributed virtual network is a uplink. | [optional] 
**VlanRange** | Pointer to [**[]VirtualizationVmwareVlanRange**](VirtualizationVmwareVlanRange.md) |  | [optional] 
**VlanType** | Pointer to **string** | VLAN type of the distributed virtual network. It can be None, VLAN, VLAN Trunking or Private VLAN. * &#x60;None&#x60; - Do not tag traffic with any VLAN Id. * &#x60;VLAN&#x60; - Tag traffic with the Id from the VLAN Id field. * &#x60;VLAN trunking&#x60; - Pass VLAN traffic with Id within the VLAN trunk range to guest operating system. * &#x60;Private VLAN&#x60; - Associate the traffic with a private VLAN created on the distributed switch. | [optional] [default to "None"]
**DistributedSwitch** | Pointer to [**VirtualizationVmwareDistributedSwitchRelationship**](virtualization.VmwareDistributedSwitch.Relationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDistributedNetworkAllOf

`func NewVirtualizationVmwareDistributedNetworkAllOf(classId string, objectType string, ) *VirtualizationVmwareDistributedNetworkAllOf`

NewVirtualizationVmwareDistributedNetworkAllOf instantiates a new VirtualizationVmwareDistributedNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDistributedNetworkAllOfWithDefaults

`func NewVirtualizationVmwareDistributedNetworkAllOfWithDefaults() *VirtualizationVmwareDistributedNetworkAllOf`

NewVirtualizationVmwareDistributedNetworkAllOfWithDefaults instantiates a new VirtualizationVmwareDistributedNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetForgedTransmits

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetForgedTransmits() string`

GetForgedTransmits returns the ForgedTransmits field if non-nil, zero value otherwise.

### GetForgedTransmitsOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetForgedTransmitsOk() (*string, bool)`

GetForgedTransmitsOk returns a tuple with the ForgedTransmits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgedTransmits

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetForgedTransmits(v string)`

SetForgedTransmits sets ForgedTransmits field to given value.

### HasForgedTransmits

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasForgedTransmits() bool`

HasForgedTransmits returns a boolean if a field has been set.

### GetMacAddressChanges

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetMacAddressChanges() string`

GetMacAddressChanges returns the MacAddressChanges field if non-nil, zero value otherwise.

### GetMacAddressChangesOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetMacAddressChangesOk() (*string, bool)`

GetMacAddressChangesOk returns a tuple with the MacAddressChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressChanges

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetMacAddressChanges(v string)`

SetMacAddressChanges sets MacAddressChanges field to given value.

### HasMacAddressChanges

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasMacAddressChanges() bool`

HasMacAddressChanges returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareDistributedNetworkAllOf) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetNumHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNumHosts() int64`

GetNumHosts returns the NumHosts field if non-nil, zero value otherwise.

### GetNumHostsOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNumHostsOk() (*int64, bool)`

GetNumHostsOk returns a tuple with the NumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetNumHosts(v int64)`

SetNumHosts sets NumHosts field to given value.

### HasNumHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasNumHosts() bool`

HasNumHosts returns a boolean if a field has been set.

### GetNumPorts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNumPorts() int64`

GetNumPorts returns the NumPorts field if non-nil, zero value otherwise.

### GetNumPortsOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetNumPortsOk() (*int64, bool)`

GetNumPortsOk returns a tuple with the NumPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPorts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetNumPorts(v int64)`

SetNumPorts sets NumPorts field to given value.

### HasNumPorts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasNumPorts() bool`

HasNumPorts returns a boolean if a field has been set.

### GetPromiscuousMode

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetPromiscuousMode() string`

GetPromiscuousMode returns the PromiscuousMode field if non-nil, zero value otherwise.

### GetPromiscuousModeOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetPromiscuousModeOk() (*string, bool)`

GetPromiscuousModeOk returns a tuple with the PromiscuousMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromiscuousMode

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetPromiscuousMode(v string)`

SetPromiscuousMode sets PromiscuousMode field to given value.

### HasPromiscuousMode

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasPromiscuousMode() bool`

HasPromiscuousMode returns a boolean if a field has been set.

### GetUpLink

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetUpLink() bool`

GetUpLink returns the UpLink field if non-nil, zero value otherwise.

### GetUpLinkOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetUpLinkOk() (*bool, bool)`

GetUpLinkOk returns a tuple with the UpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLink

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetUpLink(v bool)`

SetUpLink sets UpLink field to given value.

### HasUpLink

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasUpLink() bool`

HasUpLink returns a boolean if a field has been set.

### GetVlanRange

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetVlanRange() []VirtualizationVmwareVlanRange`

GetVlanRange returns the VlanRange field if non-nil, zero value otherwise.

### GetVlanRangeOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetVlanRangeOk() (*[]VirtualizationVmwareVlanRange, bool)`

GetVlanRangeOk returns a tuple with the VlanRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRange

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetVlanRange(v []VirtualizationVmwareVlanRange)`

SetVlanRange sets VlanRange field to given value.

### HasVlanRange

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasVlanRange() bool`

HasVlanRange returns a boolean if a field has been set.

### SetVlanRangeNil

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetVlanRangeNil(b bool)`

 SetVlanRangeNil sets the value for VlanRange to be an explicit nil

### UnsetVlanRange
`func (o *VirtualizationVmwareDistributedNetworkAllOf) UnsetVlanRange()`

UnsetVlanRange ensures that no value is present for VlanRange, not even an explicit nil
### GetVlanType

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetVlanType() string`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetVlanTypeOk() (*string, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetVlanType(v string)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetDistributedSwitch

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetDistributedSwitch() VirtualizationVmwareDistributedSwitchRelationship`

GetDistributedSwitch returns the DistributedSwitch field if non-nil, zero value otherwise.

### GetDistributedSwitchOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetDistributedSwitchOk() (*VirtualizationVmwareDistributedSwitchRelationship, bool)`

GetDistributedSwitchOk returns a tuple with the DistributedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedSwitch

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetDistributedSwitch(v VirtualizationVmwareDistributedSwitchRelationship)`

SetDistributedSwitch sets DistributedSwitch field to given value.

### HasDistributedSwitch

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasDistributedSwitch() bool`

HasDistributedSwitch returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDistributedNetworkAllOf) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDistributedNetworkAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDistributedNetworkAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDistributedNetworkAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


