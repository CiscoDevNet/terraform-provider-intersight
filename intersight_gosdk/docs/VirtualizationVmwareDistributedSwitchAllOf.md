# VirtualizationVmwareDistributedSwitchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDistributedSwitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDistributedSwitch"]
**Description** | Pointer to **string** | Switch description (user provided), if any. | [optional] 
**MaxPort** | Pointer to **int64** | Maximum number of ports allowed on this distributed virtual switch. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmission unit configured on a distributed virtual switch. | [optional] 
**NicTeamingAndFailover** | Pointer to [**NullableVirtualizationVmwareTeamingAndFailover**](virtualization.VmwareTeamingAndFailover.md) |  | [optional] 
**NumHosts** | Pointer to **int64** | The total number of hosts attached to the distributed virtual switch. | [optional] 
**NumNetworks** | Pointer to **int64** | The total number of distributed networks in the distributed virtual switch. | [optional] 
**NumStandAlonePorts** | Pointer to **int64** | Number of stand-alone ports in use. | [optional] 
**NumUplinks** | Pointer to **int64** | Number of uplinks configured in this distributed virtual switch. | [optional] 
**SwitchCapacity** | Pointer to [**NullableVirtualizationStorageCapacity**](virtualization.StorageCapacity.md) |  | [optional] 
**Uuid** | Pointer to **string** | Universally Unique Id of this distributed virtual switch. | [optional] 
**Version** | Pointer to **string** | The running config&#39;s version details are represented. | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDistributedSwitchAllOf

`func NewVirtualizationVmwareDistributedSwitchAllOf(classId string, objectType string, ) *VirtualizationVmwareDistributedSwitchAllOf`

NewVirtualizationVmwareDistributedSwitchAllOf instantiates a new VirtualizationVmwareDistributedSwitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDistributedSwitchAllOfWithDefaults

`func NewVirtualizationVmwareDistributedSwitchAllOfWithDefaults() *VirtualizationVmwareDistributedSwitchAllOf`

NewVirtualizationVmwareDistributedSwitchAllOfWithDefaults instantiates a new VirtualizationVmwareDistributedSwitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxPort

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMaxPort() int64`

GetMaxPort returns the MaxPort field if non-nil, zero value otherwise.

### GetMaxPortOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMaxPortOk() (*int64, bool)`

GetMaxPortOk returns a tuple with the MaxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPort

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetMaxPort(v int64)`

SetMaxPort sets MaxPort field to given value.

### HasMaxPort

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasMaxPort() bool`

HasMaxPort returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover`

GetNicTeamingAndFailover returns the NicTeamingAndFailover field if non-nil, zero value otherwise.

### GetNicTeamingAndFailoverOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool)`

GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover)`

SetNicTeamingAndFailover sets NicTeamingAndFailover field to given value.

### HasNicTeamingAndFailover

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNicTeamingAndFailover() bool`

HasNicTeamingAndFailover returns a boolean if a field has been set.

### SetNicTeamingAndFailoverNil

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNicTeamingAndFailoverNil(b bool)`

 SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil

### UnsetNicTeamingAndFailover
`func (o *VirtualizationVmwareDistributedSwitchAllOf) UnsetNicTeamingAndFailover()`

UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
### GetNumHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumHosts() int64`

GetNumHosts returns the NumHosts field if non-nil, zero value otherwise.

### GetNumHostsOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumHostsOk() (*int64, bool)`

GetNumHostsOk returns a tuple with the NumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumHosts(v int64)`

SetNumHosts sets NumHosts field to given value.

### HasNumHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumHosts() bool`

HasNumHosts returns a boolean if a field has been set.

### GetNumNetworks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumNetworks() int64`

GetNumNetworks returns the NumNetworks field if non-nil, zero value otherwise.

### GetNumNetworksOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumNetworksOk() (*int64, bool)`

GetNumNetworksOk returns a tuple with the NumNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNetworks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumNetworks(v int64)`

SetNumNetworks sets NumNetworks field to given value.

### HasNumNetworks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumNetworks() bool`

HasNumNetworks returns a boolean if a field has been set.

### GetNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumStandAlonePorts() int64`

GetNumStandAlonePorts returns the NumStandAlonePorts field if non-nil, zero value otherwise.

### GetNumStandAlonePortsOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumStandAlonePortsOk() (*int64, bool)`

GetNumStandAlonePortsOk returns a tuple with the NumStandAlonePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumStandAlonePorts(v int64)`

SetNumStandAlonePorts sets NumStandAlonePorts field to given value.

### HasNumStandAlonePorts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumStandAlonePorts() bool`

HasNumStandAlonePorts returns a boolean if a field has been set.

### GetNumUplinks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumUplinks() int64`

GetNumUplinks returns the NumUplinks field if non-nil, zero value otherwise.

### GetNumUplinksOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumUplinksOk() (*int64, bool)`

GetNumUplinksOk returns a tuple with the NumUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUplinks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumUplinks(v int64)`

SetNumUplinks sets NumUplinks field to given value.

### HasNumUplinks

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumUplinks() bool`

HasNumUplinks returns a boolean if a field has been set.

### GetSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetSwitchCapacity() VirtualizationStorageCapacity`

GetSwitchCapacity returns the SwitchCapacity field if non-nil, zero value otherwise.

### GetSwitchCapacityOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetSwitchCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetSwitchCapacityOk returns a tuple with the SwitchCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetSwitchCapacity(v VirtualizationStorageCapacity)`

SetSwitchCapacity sets SwitchCapacity field to given value.

### HasSwitchCapacity

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasSwitchCapacity() bool`

HasSwitchCapacity returns a boolean if a field has been set.

### SetSwitchCapacityNil

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetSwitchCapacityNil(b bool)`

 SetSwitchCapacityNil sets the value for SwitchCapacity to be an explicit nil

### UnsetSwitchCapacity
`func (o *VirtualizationVmwareDistributedSwitchAllOf) UnsetSwitchCapacity()`

UnsetSwitchCapacity ensures that no value is present for SwitchCapacity, not even an explicit nil
### GetUuid

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDistributedSwitchAllOf) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDistributedSwitchAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDistributedSwitchAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDistributedSwitchAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


