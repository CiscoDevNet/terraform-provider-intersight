# CapabilitySwitchCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFcoeVlan** | Pointer to **int64** | Default Fcoe VLAN associated with this switch. | [optional] 
**DynamicVifsSupported** | Pointer to **bool** | Dynamic VIFs support on this switch. | [optional] 
**FanModulesSupported** | Pointer to **bool** | Fan Modules support on this switch. | [optional] 
**FcEndHostModeReservedVsans** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**FcUplinkPortsAutoNegotiationSupported** | Pointer to **bool** | Fc Uplink ports auto negotiation speed support on this switch. | [optional] 
**LocatorBeaconSupported** | Pointer to **bool** | Locator Beacon LED support on this switch. | [optional] 
**MaxActiveSpanSessions** | Pointer to **int64** | Maximum allowed Traffic Monitoring (SPAN) sessions on this switch. | [optional] 
**MaxEthernetPortChannelMembers** | Pointer to **int64** | Maximum allowed Ethernet Uplink Port-channel members for each Uplink Port-channel on this switch. | [optional] 
**MaxEthernetPortChannels** | Pointer to **int64** | Maximum allowed Ethernet Uplink Port-channels on this switch. | [optional] 
**MaxEthernetUplinkPorts** | Pointer to **int64** | Maximum allowed Ethernet Uplink Ports on this switch. | [optional] 
**MaxFcFcoePortChannels** | Pointer to **int64** | Total maximum Fc and Fcoe Port-channels allowed on this switch. | [optional] 
**MaxFcPortChannelMembers** | Pointer to **int64** | Maximum allowed FC Uplink Port-channel members for each FCoE Port-channel on this switch. | [optional] 
**MaxFcoePortChannelMembers** | Pointer to **int64** | Maximum allowed FCoE Uplink Port-channel members for each FCoE Port-channel on this switch. | [optional] 
**MaxPorts** | Pointer to **int64** | Maximum allowed physical ports on this switch. | [optional] 
**MaxSlots** | Pointer to **int64** | Maximum allowed physical slots on this switch. | [optional] 
**MaxVsansSupported** | Pointer to **int64** | Maximum number of Vsans supported on this switch. | [optional] 
**MinActiveFans** | Pointer to **int64** | Minimum number of fans needed to be active/running on this switch. | [optional] 
**PortsSupporting100gSpeed** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupporting10gSpeed** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupporting1gSpeed** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupporting25gSpeed** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupporting40gSpeed** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupportingBreakout** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupportingFcoe** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**PortsSupportingServerRole** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**ReservedVsans** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**SerenoNetflowSupported** | Pointer to **bool** | Sereno Adaptor with Netflow support on this switch. | [optional] 
**UnifiedPorts** | Pointer to [**[]CapabilityPortRange**](capability.PortRange.md) |  | [optional] 
**UnifiedRule** | Pointer to **string** | The Slider rule for Unified ports on this switch. | [optional] 
**VpCompressionSupported** | Pointer to **bool** | VP Compression support on this switch. | [optional] 

## Methods

### NewCapabilitySwitchCapability

`func NewCapabilitySwitchCapability() *CapabilitySwitchCapability`

NewCapabilitySwitchCapability instantiates a new CapabilitySwitchCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchCapabilityWithDefaults

`func NewCapabilitySwitchCapabilityWithDefaults() *CapabilitySwitchCapability`

NewCapabilitySwitchCapabilityWithDefaults instantiates a new CapabilitySwitchCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFcoeVlan

`func (o *CapabilitySwitchCapability) GetDefaultFcoeVlan() int64`

GetDefaultFcoeVlan returns the DefaultFcoeVlan field if non-nil, zero value otherwise.

### GetDefaultFcoeVlanOk

`func (o *CapabilitySwitchCapability) GetDefaultFcoeVlanOk() (*int64, bool)`

GetDefaultFcoeVlanOk returns a tuple with the DefaultFcoeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFcoeVlan

`func (o *CapabilitySwitchCapability) SetDefaultFcoeVlan(v int64)`

SetDefaultFcoeVlan sets DefaultFcoeVlan field to given value.

### HasDefaultFcoeVlan

`func (o *CapabilitySwitchCapability) HasDefaultFcoeVlan() bool`

HasDefaultFcoeVlan returns a boolean if a field has been set.

### GetDynamicVifsSupported

`func (o *CapabilitySwitchCapability) GetDynamicVifsSupported() bool`

GetDynamicVifsSupported returns the DynamicVifsSupported field if non-nil, zero value otherwise.

### GetDynamicVifsSupportedOk

`func (o *CapabilitySwitchCapability) GetDynamicVifsSupportedOk() (*bool, bool)`

GetDynamicVifsSupportedOk returns a tuple with the DynamicVifsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVifsSupported

`func (o *CapabilitySwitchCapability) SetDynamicVifsSupported(v bool)`

SetDynamicVifsSupported sets DynamicVifsSupported field to given value.

### HasDynamicVifsSupported

`func (o *CapabilitySwitchCapability) HasDynamicVifsSupported() bool`

HasDynamicVifsSupported returns a boolean if a field has been set.

### GetFanModulesSupported

`func (o *CapabilitySwitchCapability) GetFanModulesSupported() bool`

GetFanModulesSupported returns the FanModulesSupported field if non-nil, zero value otherwise.

### GetFanModulesSupportedOk

`func (o *CapabilitySwitchCapability) GetFanModulesSupportedOk() (*bool, bool)`

GetFanModulesSupportedOk returns a tuple with the FanModulesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModulesSupported

`func (o *CapabilitySwitchCapability) SetFanModulesSupported(v bool)`

SetFanModulesSupported sets FanModulesSupported field to given value.

### HasFanModulesSupported

`func (o *CapabilitySwitchCapability) HasFanModulesSupported() bool`

HasFanModulesSupported returns a boolean if a field has been set.

### GetFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapability) GetFcEndHostModeReservedVsans() []CapabilityPortRange`

GetFcEndHostModeReservedVsans returns the FcEndHostModeReservedVsans field if non-nil, zero value otherwise.

### GetFcEndHostModeReservedVsansOk

`func (o *CapabilitySwitchCapability) GetFcEndHostModeReservedVsansOk() (*[]CapabilityPortRange, bool)`

GetFcEndHostModeReservedVsansOk returns a tuple with the FcEndHostModeReservedVsans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapability) SetFcEndHostModeReservedVsans(v []CapabilityPortRange)`

SetFcEndHostModeReservedVsans sets FcEndHostModeReservedVsans field to given value.

### HasFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapability) HasFcEndHostModeReservedVsans() bool`

HasFcEndHostModeReservedVsans returns a boolean if a field has been set.

### GetFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapability) GetFcUplinkPortsAutoNegotiationSupported() bool`

GetFcUplinkPortsAutoNegotiationSupported returns the FcUplinkPortsAutoNegotiationSupported field if non-nil, zero value otherwise.

### GetFcUplinkPortsAutoNegotiationSupportedOk

`func (o *CapabilitySwitchCapability) GetFcUplinkPortsAutoNegotiationSupportedOk() (*bool, bool)`

GetFcUplinkPortsAutoNegotiationSupportedOk returns a tuple with the FcUplinkPortsAutoNegotiationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapability) SetFcUplinkPortsAutoNegotiationSupported(v bool)`

SetFcUplinkPortsAutoNegotiationSupported sets FcUplinkPortsAutoNegotiationSupported field to given value.

### HasFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapability) HasFcUplinkPortsAutoNegotiationSupported() bool`

HasFcUplinkPortsAutoNegotiationSupported returns a boolean if a field has been set.

### GetLocatorBeaconSupported

`func (o *CapabilitySwitchCapability) GetLocatorBeaconSupported() bool`

GetLocatorBeaconSupported returns the LocatorBeaconSupported field if non-nil, zero value otherwise.

### GetLocatorBeaconSupportedOk

`func (o *CapabilitySwitchCapability) GetLocatorBeaconSupportedOk() (*bool, bool)`

GetLocatorBeaconSupportedOk returns a tuple with the LocatorBeaconSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorBeaconSupported

`func (o *CapabilitySwitchCapability) SetLocatorBeaconSupported(v bool)`

SetLocatorBeaconSupported sets LocatorBeaconSupported field to given value.

### HasLocatorBeaconSupported

`func (o *CapabilitySwitchCapability) HasLocatorBeaconSupported() bool`

HasLocatorBeaconSupported returns a boolean if a field has been set.

### GetMaxActiveSpanSessions

`func (o *CapabilitySwitchCapability) GetMaxActiveSpanSessions() int64`

GetMaxActiveSpanSessions returns the MaxActiveSpanSessions field if non-nil, zero value otherwise.

### GetMaxActiveSpanSessionsOk

`func (o *CapabilitySwitchCapability) GetMaxActiveSpanSessionsOk() (*int64, bool)`

GetMaxActiveSpanSessionsOk returns a tuple with the MaxActiveSpanSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveSpanSessions

`func (o *CapabilitySwitchCapability) SetMaxActiveSpanSessions(v int64)`

SetMaxActiveSpanSessions sets MaxActiveSpanSessions field to given value.

### HasMaxActiveSpanSessions

`func (o *CapabilitySwitchCapability) HasMaxActiveSpanSessions() bool`

HasMaxActiveSpanSessions returns a boolean if a field has been set.

### GetMaxEthernetPortChannelMembers

`func (o *CapabilitySwitchCapability) GetMaxEthernetPortChannelMembers() int64`

GetMaxEthernetPortChannelMembers returns the MaxEthernetPortChannelMembers field if non-nil, zero value otherwise.

### GetMaxEthernetPortChannelMembersOk

`func (o *CapabilitySwitchCapability) GetMaxEthernetPortChannelMembersOk() (*int64, bool)`

GetMaxEthernetPortChannelMembersOk returns a tuple with the MaxEthernetPortChannelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthernetPortChannelMembers

`func (o *CapabilitySwitchCapability) SetMaxEthernetPortChannelMembers(v int64)`

SetMaxEthernetPortChannelMembers sets MaxEthernetPortChannelMembers field to given value.

### HasMaxEthernetPortChannelMembers

`func (o *CapabilitySwitchCapability) HasMaxEthernetPortChannelMembers() bool`

HasMaxEthernetPortChannelMembers returns a boolean if a field has been set.

### GetMaxEthernetPortChannels

`func (o *CapabilitySwitchCapability) GetMaxEthernetPortChannels() int64`

GetMaxEthernetPortChannels returns the MaxEthernetPortChannels field if non-nil, zero value otherwise.

### GetMaxEthernetPortChannelsOk

`func (o *CapabilitySwitchCapability) GetMaxEthernetPortChannelsOk() (*int64, bool)`

GetMaxEthernetPortChannelsOk returns a tuple with the MaxEthernetPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthernetPortChannels

`func (o *CapabilitySwitchCapability) SetMaxEthernetPortChannels(v int64)`

SetMaxEthernetPortChannels sets MaxEthernetPortChannels field to given value.

### HasMaxEthernetPortChannels

`func (o *CapabilitySwitchCapability) HasMaxEthernetPortChannels() bool`

HasMaxEthernetPortChannels returns a boolean if a field has been set.

### GetMaxEthernetUplinkPorts

`func (o *CapabilitySwitchCapability) GetMaxEthernetUplinkPorts() int64`

GetMaxEthernetUplinkPorts returns the MaxEthernetUplinkPorts field if non-nil, zero value otherwise.

### GetMaxEthernetUplinkPortsOk

`func (o *CapabilitySwitchCapability) GetMaxEthernetUplinkPortsOk() (*int64, bool)`

GetMaxEthernetUplinkPortsOk returns a tuple with the MaxEthernetUplinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthernetUplinkPorts

`func (o *CapabilitySwitchCapability) SetMaxEthernetUplinkPorts(v int64)`

SetMaxEthernetUplinkPorts sets MaxEthernetUplinkPorts field to given value.

### HasMaxEthernetUplinkPorts

`func (o *CapabilitySwitchCapability) HasMaxEthernetUplinkPorts() bool`

HasMaxEthernetUplinkPorts returns a boolean if a field has been set.

### GetMaxFcFcoePortChannels

`func (o *CapabilitySwitchCapability) GetMaxFcFcoePortChannels() int64`

GetMaxFcFcoePortChannels returns the MaxFcFcoePortChannels field if non-nil, zero value otherwise.

### GetMaxFcFcoePortChannelsOk

`func (o *CapabilitySwitchCapability) GetMaxFcFcoePortChannelsOk() (*int64, bool)`

GetMaxFcFcoePortChannelsOk returns a tuple with the MaxFcFcoePortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFcFcoePortChannels

`func (o *CapabilitySwitchCapability) SetMaxFcFcoePortChannels(v int64)`

SetMaxFcFcoePortChannels sets MaxFcFcoePortChannels field to given value.

### HasMaxFcFcoePortChannels

`func (o *CapabilitySwitchCapability) HasMaxFcFcoePortChannels() bool`

HasMaxFcFcoePortChannels returns a boolean if a field has been set.

### GetMaxFcPortChannelMembers

`func (o *CapabilitySwitchCapability) GetMaxFcPortChannelMembers() int64`

GetMaxFcPortChannelMembers returns the MaxFcPortChannelMembers field if non-nil, zero value otherwise.

### GetMaxFcPortChannelMembersOk

`func (o *CapabilitySwitchCapability) GetMaxFcPortChannelMembersOk() (*int64, bool)`

GetMaxFcPortChannelMembersOk returns a tuple with the MaxFcPortChannelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFcPortChannelMembers

`func (o *CapabilitySwitchCapability) SetMaxFcPortChannelMembers(v int64)`

SetMaxFcPortChannelMembers sets MaxFcPortChannelMembers field to given value.

### HasMaxFcPortChannelMembers

`func (o *CapabilitySwitchCapability) HasMaxFcPortChannelMembers() bool`

HasMaxFcPortChannelMembers returns a boolean if a field has been set.

### GetMaxFcoePortChannelMembers

`func (o *CapabilitySwitchCapability) GetMaxFcoePortChannelMembers() int64`

GetMaxFcoePortChannelMembers returns the MaxFcoePortChannelMembers field if non-nil, zero value otherwise.

### GetMaxFcoePortChannelMembersOk

`func (o *CapabilitySwitchCapability) GetMaxFcoePortChannelMembersOk() (*int64, bool)`

GetMaxFcoePortChannelMembersOk returns a tuple with the MaxFcoePortChannelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFcoePortChannelMembers

`func (o *CapabilitySwitchCapability) SetMaxFcoePortChannelMembers(v int64)`

SetMaxFcoePortChannelMembers sets MaxFcoePortChannelMembers field to given value.

### HasMaxFcoePortChannelMembers

`func (o *CapabilitySwitchCapability) HasMaxFcoePortChannelMembers() bool`

HasMaxFcoePortChannelMembers returns a boolean if a field has been set.

### GetMaxPorts

`func (o *CapabilitySwitchCapability) GetMaxPorts() int64`

GetMaxPorts returns the MaxPorts field if non-nil, zero value otherwise.

### GetMaxPortsOk

`func (o *CapabilitySwitchCapability) GetMaxPortsOk() (*int64, bool)`

GetMaxPortsOk returns a tuple with the MaxPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPorts

`func (o *CapabilitySwitchCapability) SetMaxPorts(v int64)`

SetMaxPorts sets MaxPorts field to given value.

### HasMaxPorts

`func (o *CapabilitySwitchCapability) HasMaxPorts() bool`

HasMaxPorts returns a boolean if a field has been set.

### GetMaxSlots

`func (o *CapabilitySwitchCapability) GetMaxSlots() int64`

GetMaxSlots returns the MaxSlots field if non-nil, zero value otherwise.

### GetMaxSlotsOk

`func (o *CapabilitySwitchCapability) GetMaxSlotsOk() (*int64, bool)`

GetMaxSlotsOk returns a tuple with the MaxSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSlots

`func (o *CapabilitySwitchCapability) SetMaxSlots(v int64)`

SetMaxSlots sets MaxSlots field to given value.

### HasMaxSlots

`func (o *CapabilitySwitchCapability) HasMaxSlots() bool`

HasMaxSlots returns a boolean if a field has been set.

### GetMaxVsansSupported

`func (o *CapabilitySwitchCapability) GetMaxVsansSupported() int64`

GetMaxVsansSupported returns the MaxVsansSupported field if non-nil, zero value otherwise.

### GetMaxVsansSupportedOk

`func (o *CapabilitySwitchCapability) GetMaxVsansSupportedOk() (*int64, bool)`

GetMaxVsansSupportedOk returns a tuple with the MaxVsansSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVsansSupported

`func (o *CapabilitySwitchCapability) SetMaxVsansSupported(v int64)`

SetMaxVsansSupported sets MaxVsansSupported field to given value.

### HasMaxVsansSupported

`func (o *CapabilitySwitchCapability) HasMaxVsansSupported() bool`

HasMaxVsansSupported returns a boolean if a field has been set.

### GetMinActiveFans

`func (o *CapabilitySwitchCapability) GetMinActiveFans() int64`

GetMinActiveFans returns the MinActiveFans field if non-nil, zero value otherwise.

### GetMinActiveFansOk

`func (o *CapabilitySwitchCapability) GetMinActiveFansOk() (*int64, bool)`

GetMinActiveFansOk returns a tuple with the MinActiveFans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActiveFans

`func (o *CapabilitySwitchCapability) SetMinActiveFans(v int64)`

SetMinActiveFans sets MinActiveFans field to given value.

### HasMinActiveFans

`func (o *CapabilitySwitchCapability) HasMinActiveFans() bool`

HasMinActiveFans returns a boolean if a field has been set.

### GetPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapability) GetPortsSupporting100gSpeed() []CapabilityPortRange`

GetPortsSupporting100gSpeed returns the PortsSupporting100gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting100gSpeedOk

`func (o *CapabilitySwitchCapability) GetPortsSupporting100gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting100gSpeedOk returns a tuple with the PortsSupporting100gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapability) SetPortsSupporting100gSpeed(v []CapabilityPortRange)`

SetPortsSupporting100gSpeed sets PortsSupporting100gSpeed field to given value.

### HasPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapability) HasPortsSupporting100gSpeed() bool`

HasPortsSupporting100gSpeed returns a boolean if a field has been set.

### GetPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapability) GetPortsSupporting10gSpeed() []CapabilityPortRange`

GetPortsSupporting10gSpeed returns the PortsSupporting10gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting10gSpeedOk

`func (o *CapabilitySwitchCapability) GetPortsSupporting10gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting10gSpeedOk returns a tuple with the PortsSupporting10gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapability) SetPortsSupporting10gSpeed(v []CapabilityPortRange)`

SetPortsSupporting10gSpeed sets PortsSupporting10gSpeed field to given value.

### HasPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapability) HasPortsSupporting10gSpeed() bool`

HasPortsSupporting10gSpeed returns a boolean if a field has been set.

### GetPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapability) GetPortsSupporting1gSpeed() []CapabilityPortRange`

GetPortsSupporting1gSpeed returns the PortsSupporting1gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting1gSpeedOk

`func (o *CapabilitySwitchCapability) GetPortsSupporting1gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting1gSpeedOk returns a tuple with the PortsSupporting1gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapability) SetPortsSupporting1gSpeed(v []CapabilityPortRange)`

SetPortsSupporting1gSpeed sets PortsSupporting1gSpeed field to given value.

### HasPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapability) HasPortsSupporting1gSpeed() bool`

HasPortsSupporting1gSpeed returns a boolean if a field has been set.

### GetPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapability) GetPortsSupporting25gSpeed() []CapabilityPortRange`

GetPortsSupporting25gSpeed returns the PortsSupporting25gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting25gSpeedOk

`func (o *CapabilitySwitchCapability) GetPortsSupporting25gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting25gSpeedOk returns a tuple with the PortsSupporting25gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapability) SetPortsSupporting25gSpeed(v []CapabilityPortRange)`

SetPortsSupporting25gSpeed sets PortsSupporting25gSpeed field to given value.

### HasPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapability) HasPortsSupporting25gSpeed() bool`

HasPortsSupporting25gSpeed returns a boolean if a field has been set.

### GetPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapability) GetPortsSupporting40gSpeed() []CapabilityPortRange`

GetPortsSupporting40gSpeed returns the PortsSupporting40gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting40gSpeedOk

`func (o *CapabilitySwitchCapability) GetPortsSupporting40gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting40gSpeedOk returns a tuple with the PortsSupporting40gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapability) SetPortsSupporting40gSpeed(v []CapabilityPortRange)`

SetPortsSupporting40gSpeed sets PortsSupporting40gSpeed field to given value.

### HasPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapability) HasPortsSupporting40gSpeed() bool`

HasPortsSupporting40gSpeed returns a boolean if a field has been set.

### GetPortsSupportingBreakout

`func (o *CapabilitySwitchCapability) GetPortsSupportingBreakout() []CapabilityPortRange`

GetPortsSupportingBreakout returns the PortsSupportingBreakout field if non-nil, zero value otherwise.

### GetPortsSupportingBreakoutOk

`func (o *CapabilitySwitchCapability) GetPortsSupportingBreakoutOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingBreakoutOk returns a tuple with the PortsSupportingBreakout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingBreakout

`func (o *CapabilitySwitchCapability) SetPortsSupportingBreakout(v []CapabilityPortRange)`

SetPortsSupportingBreakout sets PortsSupportingBreakout field to given value.

### HasPortsSupportingBreakout

`func (o *CapabilitySwitchCapability) HasPortsSupportingBreakout() bool`

HasPortsSupportingBreakout returns a boolean if a field has been set.

### GetPortsSupportingFcoe

`func (o *CapabilitySwitchCapability) GetPortsSupportingFcoe() []CapabilityPortRange`

GetPortsSupportingFcoe returns the PortsSupportingFcoe field if non-nil, zero value otherwise.

### GetPortsSupportingFcoeOk

`func (o *CapabilitySwitchCapability) GetPortsSupportingFcoeOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingFcoeOk returns a tuple with the PortsSupportingFcoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingFcoe

`func (o *CapabilitySwitchCapability) SetPortsSupportingFcoe(v []CapabilityPortRange)`

SetPortsSupportingFcoe sets PortsSupportingFcoe field to given value.

### HasPortsSupportingFcoe

`func (o *CapabilitySwitchCapability) HasPortsSupportingFcoe() bool`

HasPortsSupportingFcoe returns a boolean if a field has been set.

### GetPortsSupportingServerRole

`func (o *CapabilitySwitchCapability) GetPortsSupportingServerRole() []CapabilityPortRange`

GetPortsSupportingServerRole returns the PortsSupportingServerRole field if non-nil, zero value otherwise.

### GetPortsSupportingServerRoleOk

`func (o *CapabilitySwitchCapability) GetPortsSupportingServerRoleOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingServerRoleOk returns a tuple with the PortsSupportingServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingServerRole

`func (o *CapabilitySwitchCapability) SetPortsSupportingServerRole(v []CapabilityPortRange)`

SetPortsSupportingServerRole sets PortsSupportingServerRole field to given value.

### HasPortsSupportingServerRole

`func (o *CapabilitySwitchCapability) HasPortsSupportingServerRole() bool`

HasPortsSupportingServerRole returns a boolean if a field has been set.

### GetReservedVsans

`func (o *CapabilitySwitchCapability) GetReservedVsans() []CapabilityPortRange`

GetReservedVsans returns the ReservedVsans field if non-nil, zero value otherwise.

### GetReservedVsansOk

`func (o *CapabilitySwitchCapability) GetReservedVsansOk() (*[]CapabilityPortRange, bool)`

GetReservedVsansOk returns a tuple with the ReservedVsans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVsans

`func (o *CapabilitySwitchCapability) SetReservedVsans(v []CapabilityPortRange)`

SetReservedVsans sets ReservedVsans field to given value.

### HasReservedVsans

`func (o *CapabilitySwitchCapability) HasReservedVsans() bool`

HasReservedVsans returns a boolean if a field has been set.

### GetSerenoNetflowSupported

`func (o *CapabilitySwitchCapability) GetSerenoNetflowSupported() bool`

GetSerenoNetflowSupported returns the SerenoNetflowSupported field if non-nil, zero value otherwise.

### GetSerenoNetflowSupportedOk

`func (o *CapabilitySwitchCapability) GetSerenoNetflowSupportedOk() (*bool, bool)`

GetSerenoNetflowSupportedOk returns a tuple with the SerenoNetflowSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerenoNetflowSupported

`func (o *CapabilitySwitchCapability) SetSerenoNetflowSupported(v bool)`

SetSerenoNetflowSupported sets SerenoNetflowSupported field to given value.

### HasSerenoNetflowSupported

`func (o *CapabilitySwitchCapability) HasSerenoNetflowSupported() bool`

HasSerenoNetflowSupported returns a boolean if a field has been set.

### GetUnifiedPorts

`func (o *CapabilitySwitchCapability) GetUnifiedPorts() []CapabilityPortRange`

GetUnifiedPorts returns the UnifiedPorts field if non-nil, zero value otherwise.

### GetUnifiedPortsOk

`func (o *CapabilitySwitchCapability) GetUnifiedPortsOk() (*[]CapabilityPortRange, bool)`

GetUnifiedPortsOk returns a tuple with the UnifiedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedPorts

`func (o *CapabilitySwitchCapability) SetUnifiedPorts(v []CapabilityPortRange)`

SetUnifiedPorts sets UnifiedPorts field to given value.

### HasUnifiedPorts

`func (o *CapabilitySwitchCapability) HasUnifiedPorts() bool`

HasUnifiedPorts returns a boolean if a field has been set.

### GetUnifiedRule

`func (o *CapabilitySwitchCapability) GetUnifiedRule() string`

GetUnifiedRule returns the UnifiedRule field if non-nil, zero value otherwise.

### GetUnifiedRuleOk

`func (o *CapabilitySwitchCapability) GetUnifiedRuleOk() (*string, bool)`

GetUnifiedRuleOk returns a tuple with the UnifiedRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedRule

`func (o *CapabilitySwitchCapability) SetUnifiedRule(v string)`

SetUnifiedRule sets UnifiedRule field to given value.

### HasUnifiedRule

`func (o *CapabilitySwitchCapability) HasUnifiedRule() bool`

HasUnifiedRule returns a boolean if a field has been set.

### GetVpCompressionSupported

`func (o *CapabilitySwitchCapability) GetVpCompressionSupported() bool`

GetVpCompressionSupported returns the VpCompressionSupported field if non-nil, zero value otherwise.

### GetVpCompressionSupportedOk

`func (o *CapabilitySwitchCapability) GetVpCompressionSupportedOk() (*bool, bool)`

GetVpCompressionSupportedOk returns a tuple with the VpCompressionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpCompressionSupported

`func (o *CapabilitySwitchCapability) SetVpCompressionSupported(v bool)`

SetVpCompressionSupported sets VpCompressionSupported field to given value.

### HasVpCompressionSupported

`func (o *CapabilitySwitchCapability) HasVpCompressionSupported() bool`

HasVpCompressionSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


