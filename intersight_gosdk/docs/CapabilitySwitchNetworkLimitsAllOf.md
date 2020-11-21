# CapabilitySwitchNetworkLimitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchNetworkLimits"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchNetworkLimits"]
**MaxCompressedPortVlanCount** | Pointer to **int64** | Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect. | [optional] 
**MaxUncompressedPortVlanCount** | Pointer to **int64** | Maximum configurable VLANs on Switch/Fabric-Interconnect. | [optional] 
**MaximumActiveTrafficMonitoringSessions** | Pointer to **int64** | Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect. | [optional] 
**MaximumEthernetPortChannels** | Pointer to **int64** | Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect. | [optional] 
**MaximumEthernetUplinkPorts** | Pointer to **int64** | Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect. | [optional] 
**MaximumFcPortChannelMembers** | Pointer to **int64** | Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect. | [optional] 
**MaximumFcPortChannels** | Pointer to **int64** | Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect. | [optional] 
**MaximumIgmpGroups** | Pointer to **int64** | Maximum configurable IGMP Groups on Switch/Fabric-Interconnect. | [optional] 
**MaximumPortChannelMembers** | Pointer to **int64** | Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect. | [optional] 
**MaximumPrimaryVlan** | Pointer to **int64** | Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect. | [optional] 
**MaximumSecondaryVlan** | Pointer to **int64** | Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect. | [optional] 
**MaximumSecondaryVlanPerPrimary** | Pointer to **int64** | Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect. | [optional] 
**MaximumVifs** | Pointer to **int64** | Maximum allowes VIFs on Switch/Fabric-Interconnect. | [optional] 
**MaximumVlans** | Pointer to **int64** | Maximum configurable VLANs on Switch/Fabric-Interconnect. | [optional] 
**MinimumActiveFans** | Pointer to **int64** | Minimum required fans in &#39;active&#39; state for this Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilitySwitchNetworkLimitsAllOf

`func NewCapabilitySwitchNetworkLimitsAllOf(classId string, objectType string, ) *CapabilitySwitchNetworkLimitsAllOf`

NewCapabilitySwitchNetworkLimitsAllOf instantiates a new CapabilitySwitchNetworkLimitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchNetworkLimitsAllOfWithDefaults

`func NewCapabilitySwitchNetworkLimitsAllOfWithDefaults() *CapabilitySwitchNetworkLimitsAllOf`

NewCapabilitySwitchNetworkLimitsAllOfWithDefaults instantiates a new CapabilitySwitchNetworkLimitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxCompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxCompressedPortVlanCount() int64`

GetMaxCompressedPortVlanCount returns the MaxCompressedPortVlanCount field if non-nil, zero value otherwise.

### GetMaxCompressedPortVlanCountOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxCompressedPortVlanCountOk() (*int64, bool)`

GetMaxCompressedPortVlanCountOk returns a tuple with the MaxCompressedPortVlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaxCompressedPortVlanCount(v int64)`

SetMaxCompressedPortVlanCount sets MaxCompressedPortVlanCount field to given value.

### HasMaxCompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaxCompressedPortVlanCount() bool`

HasMaxCompressedPortVlanCount returns a boolean if a field has been set.

### GetMaxUncompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxUncompressedPortVlanCount() int64`

GetMaxUncompressedPortVlanCount returns the MaxUncompressedPortVlanCount field if non-nil, zero value otherwise.

### GetMaxUncompressedPortVlanCountOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxUncompressedPortVlanCountOk() (*int64, bool)`

GetMaxUncompressedPortVlanCountOk returns a tuple with the MaxUncompressedPortVlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUncompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaxUncompressedPortVlanCount(v int64)`

SetMaxUncompressedPortVlanCount sets MaxUncompressedPortVlanCount field to given value.

### HasMaxUncompressedPortVlanCount

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaxUncompressedPortVlanCount() bool`

HasMaxUncompressedPortVlanCount returns a boolean if a field has been set.

### GetMaximumActiveTrafficMonitoringSessions

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumActiveTrafficMonitoringSessions() int64`

GetMaximumActiveTrafficMonitoringSessions returns the MaximumActiveTrafficMonitoringSessions field if non-nil, zero value otherwise.

### GetMaximumActiveTrafficMonitoringSessionsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumActiveTrafficMonitoringSessionsOk() (*int64, bool)`

GetMaximumActiveTrafficMonitoringSessionsOk returns a tuple with the MaximumActiveTrafficMonitoringSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumActiveTrafficMonitoringSessions

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumActiveTrafficMonitoringSessions(v int64)`

SetMaximumActiveTrafficMonitoringSessions sets MaximumActiveTrafficMonitoringSessions field to given value.

### HasMaximumActiveTrafficMonitoringSessions

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumActiveTrafficMonitoringSessions() bool`

HasMaximumActiveTrafficMonitoringSessions returns a boolean if a field has been set.

### GetMaximumEthernetPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetPortChannels() int64`

GetMaximumEthernetPortChannels returns the MaximumEthernetPortChannels field if non-nil, zero value otherwise.

### GetMaximumEthernetPortChannelsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetPortChannelsOk() (*int64, bool)`

GetMaximumEthernetPortChannelsOk returns a tuple with the MaximumEthernetPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumEthernetPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumEthernetPortChannels(v int64)`

SetMaximumEthernetPortChannels sets MaximumEthernetPortChannels field to given value.

### HasMaximumEthernetPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumEthernetPortChannels() bool`

HasMaximumEthernetPortChannels returns a boolean if a field has been set.

### GetMaximumEthernetUplinkPorts

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetUplinkPorts() int64`

GetMaximumEthernetUplinkPorts returns the MaximumEthernetUplinkPorts field if non-nil, zero value otherwise.

### GetMaximumEthernetUplinkPortsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetUplinkPortsOk() (*int64, bool)`

GetMaximumEthernetUplinkPortsOk returns a tuple with the MaximumEthernetUplinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumEthernetUplinkPorts

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumEthernetUplinkPorts(v int64)`

SetMaximumEthernetUplinkPorts sets MaximumEthernetUplinkPorts field to given value.

### HasMaximumEthernetUplinkPorts

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumEthernetUplinkPorts() bool`

HasMaximumEthernetUplinkPorts returns a boolean if a field has been set.

### GetMaximumFcPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelMembers() int64`

GetMaximumFcPortChannelMembers returns the MaximumFcPortChannelMembers field if non-nil, zero value otherwise.

### GetMaximumFcPortChannelMembersOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelMembersOk() (*int64, bool)`

GetMaximumFcPortChannelMembersOk returns a tuple with the MaximumFcPortChannelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFcPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumFcPortChannelMembers(v int64)`

SetMaximumFcPortChannelMembers sets MaximumFcPortChannelMembers field to given value.

### HasMaximumFcPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumFcPortChannelMembers() bool`

HasMaximumFcPortChannelMembers returns a boolean if a field has been set.

### GetMaximumFcPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannels() int64`

GetMaximumFcPortChannels returns the MaximumFcPortChannels field if non-nil, zero value otherwise.

### GetMaximumFcPortChannelsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelsOk() (*int64, bool)`

GetMaximumFcPortChannelsOk returns a tuple with the MaximumFcPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFcPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumFcPortChannels(v int64)`

SetMaximumFcPortChannels sets MaximumFcPortChannels field to given value.

### HasMaximumFcPortChannels

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumFcPortChannels() bool`

HasMaximumFcPortChannels returns a boolean if a field has been set.

### GetMaximumIgmpGroups

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumIgmpGroups() int64`

GetMaximumIgmpGroups returns the MaximumIgmpGroups field if non-nil, zero value otherwise.

### GetMaximumIgmpGroupsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumIgmpGroupsOk() (*int64, bool)`

GetMaximumIgmpGroupsOk returns a tuple with the MaximumIgmpGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIgmpGroups

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumIgmpGroups(v int64)`

SetMaximumIgmpGroups sets MaximumIgmpGroups field to given value.

### HasMaximumIgmpGroups

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumIgmpGroups() bool`

HasMaximumIgmpGroups returns a boolean if a field has been set.

### GetMaximumPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPortChannelMembers() int64`

GetMaximumPortChannelMembers returns the MaximumPortChannelMembers field if non-nil, zero value otherwise.

### GetMaximumPortChannelMembersOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPortChannelMembersOk() (*int64, bool)`

GetMaximumPortChannelMembersOk returns a tuple with the MaximumPortChannelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumPortChannelMembers(v int64)`

SetMaximumPortChannelMembers sets MaximumPortChannelMembers field to given value.

### HasMaximumPortChannelMembers

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumPortChannelMembers() bool`

HasMaximumPortChannelMembers returns a boolean if a field has been set.

### GetMaximumPrimaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPrimaryVlan() int64`

GetMaximumPrimaryVlan returns the MaximumPrimaryVlan field if non-nil, zero value otherwise.

### GetMaximumPrimaryVlanOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPrimaryVlanOk() (*int64, bool)`

GetMaximumPrimaryVlanOk returns a tuple with the MaximumPrimaryVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPrimaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumPrimaryVlan(v int64)`

SetMaximumPrimaryVlan sets MaximumPrimaryVlan field to given value.

### HasMaximumPrimaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumPrimaryVlan() bool`

HasMaximumPrimaryVlan returns a boolean if a field has been set.

### GetMaximumSecondaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlan() int64`

GetMaximumSecondaryVlan returns the MaximumSecondaryVlan field if non-nil, zero value otherwise.

### GetMaximumSecondaryVlanOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanOk() (*int64, bool)`

GetMaximumSecondaryVlanOk returns a tuple with the MaximumSecondaryVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSecondaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumSecondaryVlan(v int64)`

SetMaximumSecondaryVlan sets MaximumSecondaryVlan field to given value.

### HasMaximumSecondaryVlan

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumSecondaryVlan() bool`

HasMaximumSecondaryVlan returns a boolean if a field has been set.

### GetMaximumSecondaryVlanPerPrimary

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanPerPrimary() int64`

GetMaximumSecondaryVlanPerPrimary returns the MaximumSecondaryVlanPerPrimary field if non-nil, zero value otherwise.

### GetMaximumSecondaryVlanPerPrimaryOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanPerPrimaryOk() (*int64, bool)`

GetMaximumSecondaryVlanPerPrimaryOk returns a tuple with the MaximumSecondaryVlanPerPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSecondaryVlanPerPrimary

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumSecondaryVlanPerPrimary(v int64)`

SetMaximumSecondaryVlanPerPrimary sets MaximumSecondaryVlanPerPrimary field to given value.

### HasMaximumSecondaryVlanPerPrimary

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumSecondaryVlanPerPrimary() bool`

HasMaximumSecondaryVlanPerPrimary returns a boolean if a field has been set.

### GetMaximumVifs

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVifs() int64`

GetMaximumVifs returns the MaximumVifs field if non-nil, zero value otherwise.

### GetMaximumVifsOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVifsOk() (*int64, bool)`

GetMaximumVifsOk returns a tuple with the MaximumVifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVifs

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumVifs(v int64)`

SetMaximumVifs sets MaximumVifs field to given value.

### HasMaximumVifs

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumVifs() bool`

HasMaximumVifs returns a boolean if a field has been set.

### GetMaximumVlans

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVlans() int64`

GetMaximumVlans returns the MaximumVlans field if non-nil, zero value otherwise.

### GetMaximumVlansOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVlansOk() (*int64, bool)`

GetMaximumVlansOk returns a tuple with the MaximumVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVlans

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumVlans(v int64)`

SetMaximumVlans sets MaximumVlans field to given value.

### HasMaximumVlans

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumVlans() bool`

HasMaximumVlans returns a boolean if a field has been set.

### GetMinimumActiveFans

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMinimumActiveFans() int64`

GetMinimumActiveFans returns the MinimumActiveFans field if non-nil, zero value otherwise.

### GetMinimumActiveFansOk

`func (o *CapabilitySwitchNetworkLimitsAllOf) GetMinimumActiveFansOk() (*int64, bool)`

GetMinimumActiveFansOk returns a tuple with the MinimumActiveFans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumActiveFans

`func (o *CapabilitySwitchNetworkLimitsAllOf) SetMinimumActiveFans(v int64)`

SetMinimumActiveFans sets MinimumActiveFans field to given value.

### HasMinimumActiveFans

`func (o *CapabilitySwitchNetworkLimitsAllOf) HasMinimumActiveFans() bool`

HasMinimumActiveFans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


