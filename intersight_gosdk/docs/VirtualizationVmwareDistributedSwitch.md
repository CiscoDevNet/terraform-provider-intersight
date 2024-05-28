# VirtualizationVmwareDistributedSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDistributedSwitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDistributedSwitch"]
**Description** | Pointer to **string** | Switch description (user provided), if any. | [optional] 
**DiscoveryProtocol** | Pointer to [**NullableVirtualizationVmwareDiscoveryProtocol**](VirtualizationVmwareDiscoveryProtocol.md) |  | [optional] 
**MaxPort** | Pointer to **int64** | Maximum number of ports allowed on this distributed virtual switch. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmission unit configured on a distributed virtual switch. | [optional] 
**NetworkIoControl** | Pointer to **bool** | If network io control is enabled, will set the value as true. | [optional] 
**NicTeamingAndFailover** | Pointer to [**NullableVirtualizationVmwareTeamingAndFailover**](VirtualizationVmwareTeamingAndFailover.md) |  | [optional] 
**NumHosts** | Pointer to **int64** | The total number of hosts attached to the distributed virtual switch. | [optional] 
**NumNetworks** | Pointer to **int64** | The total number of distributed networks in the distributed virtual switch. | [optional] 
**NumStandAlonePorts** | Pointer to **int64** | Number of stand-alone ports in use. | [optional] 
**NumUplinks** | Pointer to **int64** | Number of uplinks configured in this distributed virtual switch. | [optional] 
**ResourceAllocationSystemTraffic** | Pointer to [**[]VirtualizationVmwareResourceAllocationSystemTrafficTypes**](VirtualizationVmwareResourceAllocationSystemTrafficTypes.md) |  | [optional] 
**SwitchCapacity** | Pointer to [**NullableVirtualizationStorageCapacity**](VirtualizationStorageCapacity.md) |  | [optional] 
**Uuid** | Pointer to **string** | Universally Unique Id of this distributed virtual switch. | [optional] 
**Version** | Pointer to **string** | The running config&#39;s version details are represented. | [optional] 
**Datacenter** | Pointer to [**NullableVirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDistributedSwitch

`func NewVirtualizationVmwareDistributedSwitch(classId string, objectType string, ) *VirtualizationVmwareDistributedSwitch`

NewVirtualizationVmwareDistributedSwitch instantiates a new VirtualizationVmwareDistributedSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDistributedSwitchWithDefaults

`func NewVirtualizationVmwareDistributedSwitchWithDefaults() *VirtualizationVmwareDistributedSwitch`

NewVirtualizationVmwareDistributedSwitchWithDefaults instantiates a new VirtualizationVmwareDistributedSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDistributedSwitch) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDistributedSwitch) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDistributedSwitch) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDistributedSwitch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDistributedSwitch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDistributedSwitch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VirtualizationVmwareDistributedSwitch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVmwareDistributedSwitch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVmwareDistributedSwitch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVmwareDistributedSwitch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscoveryProtocol

`func (o *VirtualizationVmwareDistributedSwitch) GetDiscoveryProtocol() VirtualizationVmwareDiscoveryProtocol`

GetDiscoveryProtocol returns the DiscoveryProtocol field if non-nil, zero value otherwise.

### GetDiscoveryProtocolOk

`func (o *VirtualizationVmwareDistributedSwitch) GetDiscoveryProtocolOk() (*VirtualizationVmwareDiscoveryProtocol, bool)`

GetDiscoveryProtocolOk returns a tuple with the DiscoveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryProtocol

`func (o *VirtualizationVmwareDistributedSwitch) SetDiscoveryProtocol(v VirtualizationVmwareDiscoveryProtocol)`

SetDiscoveryProtocol sets DiscoveryProtocol field to given value.

### HasDiscoveryProtocol

`func (o *VirtualizationVmwareDistributedSwitch) HasDiscoveryProtocol() bool`

HasDiscoveryProtocol returns a boolean if a field has been set.

### SetDiscoveryProtocolNil

`func (o *VirtualizationVmwareDistributedSwitch) SetDiscoveryProtocolNil(b bool)`

 SetDiscoveryProtocolNil sets the value for DiscoveryProtocol to be an explicit nil

### UnsetDiscoveryProtocol
`func (o *VirtualizationVmwareDistributedSwitch) UnsetDiscoveryProtocol()`

UnsetDiscoveryProtocol ensures that no value is present for DiscoveryProtocol, not even an explicit nil
### GetMaxPort

`func (o *VirtualizationVmwareDistributedSwitch) GetMaxPort() int64`

GetMaxPort returns the MaxPort field if non-nil, zero value otherwise.

### GetMaxPortOk

`func (o *VirtualizationVmwareDistributedSwitch) GetMaxPortOk() (*int64, bool)`

GetMaxPortOk returns a tuple with the MaxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPort

`func (o *VirtualizationVmwareDistributedSwitch) SetMaxPort(v int64)`

SetMaxPort sets MaxPort field to given value.

### HasMaxPort

`func (o *VirtualizationVmwareDistributedSwitch) HasMaxPort() bool`

HasMaxPort returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareDistributedSwitch) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareDistributedSwitch) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareDistributedSwitch) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareDistributedSwitch) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetworkIoControl

`func (o *VirtualizationVmwareDistributedSwitch) GetNetworkIoControl() bool`

GetNetworkIoControl returns the NetworkIoControl field if non-nil, zero value otherwise.

### GetNetworkIoControlOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNetworkIoControlOk() (*bool, bool)`

GetNetworkIoControlOk returns a tuple with the NetworkIoControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIoControl

`func (o *VirtualizationVmwareDistributedSwitch) SetNetworkIoControl(v bool)`

SetNetworkIoControl sets NetworkIoControl field to given value.

### HasNetworkIoControl

`func (o *VirtualizationVmwareDistributedSwitch) HasNetworkIoControl() bool`

HasNetworkIoControl returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitch) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitch) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitch) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareDistributedSwitch) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareDistributedSwitch) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetNumHosts

`func (o *VirtualizationVmwareDistributedSwitch) GetNumHosts() int64`

GetNumHosts returns the NumHosts field if non-nil, zero value otherwise.

### GetNumHostsOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNumHostsOk() (*int64, bool)`

GetNumHostsOk returns a tuple with the NumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHosts

`func (o *VirtualizationVmwareDistributedSwitch) SetNumHosts(v int64)`

SetNumHosts sets NumHosts field to given value.

### HasNumHosts

`func (o *VirtualizationVmwareDistributedSwitch) HasNumHosts() bool`

HasNumHosts returns a boolean if a field has been set.

### GetNumNetworks

`func (o *VirtualizationVmwareDistributedSwitch) GetNumNetworks() int64`

GetNumNetworks returns the NumNetworks field if non-nil, zero value otherwise.

### GetNumNetworksOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNumNetworksOk() (*int64, bool)`

GetNumNetworksOk returns a tuple with the NumNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNetworks

`func (o *VirtualizationVmwareDistributedSwitch) SetNumNetworks(v int64)`

SetNumNetworks sets NumNetworks field to given value.

### HasNumNetworks

`func (o *VirtualizationVmwareDistributedSwitch) HasNumNetworks() bool`

HasNumNetworks returns a boolean if a field has been set.

### GetNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitch) GetNumStandAlonePorts() int64`

GetNumStandAlonePorts returns the NumStandAlonePorts field if non-nil, zero value otherwise.

### GetNumStandAlonePortsOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNumStandAlonePortsOk() (*int64, bool)`

GetNumStandAlonePortsOk returns a tuple with the NumStandAlonePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitch) SetNumStandAlonePorts(v int64)`

SetNumStandAlonePorts sets NumStandAlonePorts field to given value.

### HasNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitch) HasNumStandAlonePorts() bool`

HasNumStandAlonePorts returns a boolean if a field has been set.

### GetNumUplinks

`func (o *VirtualizationVmwareDistributedSwitch) GetNumUplinks() int64`

GetNumUplinks returns the NumUplinks field if non-nil, zero value otherwise.

### GetNumUplinksOk

`func (o *VirtualizationVmwareDistributedSwitch) GetNumUplinksOk() (*int64, bool)`

GetNumUplinksOk returns a tuple with the NumUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUplinks

`func (o *VirtualizationVmwareDistributedSwitch) SetNumUplinks(v int64)`

SetNumUplinks sets NumUplinks field to given value.

### HasNumUplinks

`func (o *VirtualizationVmwareDistributedSwitch) HasNumUplinks() bool`

HasNumUplinks returns a boolean if a field has been set.

### GetResourceAllocationSystemTraffic

`func (o *VirtualizationVmwareDistributedSwitch) GetResourceAllocationSystemTraffic() []VirtualizationVmwareResourceAllocationSystemTrafficTypes`

GetResourceAllocationSystemTraffic returns the ResourceAllocationSystemTraffic field if non-nil, zero value otherwise.

### GetResourceAllocationSystemTrafficOk

`func (o *VirtualizationVmwareDistributedSwitch) GetResourceAllocationSystemTrafficOk() (*[]VirtualizationVmwareResourceAllocationSystemTrafficTypes, bool)`

GetResourceAllocationSystemTrafficOk returns a tuple with the ResourceAllocationSystemTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAllocationSystemTraffic

`func (o *VirtualizationVmwareDistributedSwitch) SetResourceAllocationSystemTraffic(v []VirtualizationVmwareResourceAllocationSystemTrafficTypes)`

SetResourceAllocationSystemTraffic sets ResourceAllocationSystemTraffic field to given value.

### HasResourceAllocationSystemTraffic

`func (o *VirtualizationVmwareDistributedSwitch) HasResourceAllocationSystemTraffic() bool`

HasResourceAllocationSystemTraffic returns a boolean if a field has been set.

### SetResourceAllocationSystemTrafficNil

`func (o *VirtualizationVmwareDistributedSwitch) SetResourceAllocationSystemTrafficNil(b bool)`

 SetResourceAllocationSystemTrafficNil sets the value for ResourceAllocationSystemTraffic to be an explicit nil

### UnsetResourceAllocationSystemTraffic
`func (o *VirtualizationVmwareDistributedSwitch) UnsetResourceAllocationSystemTraffic()`

UnsetResourceAllocationSystemTraffic ensures that no value is present for ResourceAllocationSystemTraffic, not even an explicit nil
### GetSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitch) GetSwitchCapacity() VirtualizationStorageCapacity`

GetSwitchCapacity returns the SwitchCapacity field if non-nil, zero value otherwise.

### GetSwitchCapacityOk

`func (o *VirtualizationVmwareDistributedSwitch) GetSwitchCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetSwitchCapacityOk returns a tuple with the SwitchCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitch) SetSwitchCapacity(v VirtualizationStorageCapacity)`

SetSwitchCapacity sets SwitchCapacity field to given value.

### HasSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitch) HasSwitchCapacity() bool`

HasSwitchCapacity returns a boolean if a field has been set.

### SetSwitchCapacityNil

`func (o *VirtualizationVmwareDistributedSwitch) SetSwitchCapacityNil(b bool)`

 SetSwitchCapacityNil sets the value for SwitchCapacity to be an explicit nil

### UnsetSwitchCapacity
`func (o *VirtualizationVmwareDistributedSwitch) UnsetSwitchCapacity()`

UnsetSwitchCapacity ensures that no value is present for SwitchCapacity, not even an explicit nil
### GetUuid

`func (o *VirtualizationVmwareDistributedSwitch) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVmwareDistributedSwitch) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVmwareDistributedSwitch) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVmwareDistributedSwitch) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationVmwareDistributedSwitch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationVmwareDistributedSwitch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationVmwareDistributedSwitch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationVmwareDistributedSwitch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDistributedSwitch) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDistributedSwitch) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDistributedSwitch) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDistributedSwitch) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### SetDatacenterNil

`func (o *VirtualizationVmwareDistributedSwitch) SetDatacenterNil(b bool)`

 SetDatacenterNil sets the value for Datacenter to be an explicit nil

### UnsetDatacenter
`func (o *VirtualizationVmwareDistributedSwitch) UnsetDatacenter()`

UnsetDatacenter ensures that no value is present for Datacenter, not even an explicit nil
### GetHosts

`func (o *VirtualizationVmwareDistributedSwitch) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDistributedSwitch) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDistributedSwitch) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDistributedSwitch) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDistributedSwitch) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDistributedSwitch) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


