# CapabilitySwitchCapabilityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchCapability"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchCapability"]
**DefaultFcoeVlan** | Pointer to **int64** | Default Fcoe VLAN associated with this switch. | [optional] [readonly] 
**DynamicVifsSupported** | Pointer to **bool** | Dynamic VIFs support on this switch. | [optional] [readonly] 
**FanModulesSupported** | Pointer to **bool** | Fan Modules support on this switch. | [optional] [readonly] 
**FcEndHostModeReservedVsans** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**FcUplinkPortsAutoNegotiationSupported** | Pointer to **bool** | Fc Uplink ports auto negotiation speed support on this switch. | [optional] [readonly] 
**LocatorBeaconSupported** | Pointer to **bool** | Locator Beacon LED support on this switch. | [optional] [readonly] 
**MaxPorts** | Pointer to **int64** | Maximum allowed physical ports on this switch. | [optional] [readonly] 
**MaxSlots** | Pointer to **int64** | Maximum allowed physical slots on this switch. | [optional] [readonly] 
**MinVersionWithBreakoutSupport** | Pointer to **string** | Minimum firmware version supported for breakout ports on this switch. | [optional] [readonly] 
**MinVersionWithLocatorLedSupport** | Pointer to **string** | Minimum firmware version supported for locator leds on this switch. | [optional] [readonly] 
**MinVersionWithNegAuto25g** | Pointer to **string** | Minimum firmware version supported for &#39;negotiate auto 25000&#39; port admin speed on this switch. | [optional] [readonly] 
**NetworkLimits** | Pointer to [**NullableCapabilitySwitchNetworkLimits**](CapabilitySwitchNetworkLimits.md) |  | [optional] 
**PortsSupporting100gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting10gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting1gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting25gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting40gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingBreakout** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingFcoe** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingServerRole** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**ReservedVsans** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**SerenoNetflowSupported** | Pointer to **bool** | Sereno Adaptor with Netflow support on this switch. | [optional] [readonly] 
**ServerRoleSupportedOnBreakout** | Pointer to **[]string** |  | [optional] 
**StorageLimits** | Pointer to [**NullableCapabilitySwitchStorageLimits**](CapabilitySwitchStorageLimits.md) |  | [optional] 
**SwitchingModeCapabilities** | Pointer to [**[]CapabilitySwitchingModeCapability**](CapabilitySwitchingModeCapability.md) |  | [optional] 
**SystemLimits** | Pointer to [**NullableCapabilitySwitchSystemLimits**](CapabilitySwitchSystemLimits.md) |  | [optional] 
**UnifiedPorts** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**UnifiedRule** | Pointer to **string** | The Slider rule for Unified ports on this switch. | [optional] [readonly] 
**UplinkAdminPortSpeedNegAuto25GbpsSupported** | Pointer to **bool** | &#39;Negotiate Auto 25000&#39; admin speed support on this switch for port or port-channel with Ethernet Uplink/Appliance/FCoE Uplink roles. | [optional] [readonly] 

## Methods

### NewCapabilitySwitchCapabilityAllOf

`func NewCapabilitySwitchCapabilityAllOf(classId string, objectType string, ) *CapabilitySwitchCapabilityAllOf`

NewCapabilitySwitchCapabilityAllOf instantiates a new CapabilitySwitchCapabilityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchCapabilityAllOfWithDefaults

`func NewCapabilitySwitchCapabilityAllOfWithDefaults() *CapabilitySwitchCapabilityAllOf`

NewCapabilitySwitchCapabilityAllOfWithDefaults instantiates a new CapabilitySwitchCapabilityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchCapabilityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchCapabilityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchCapabilityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchCapabilityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchCapabilityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchCapabilityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultFcoeVlan

`func (o *CapabilitySwitchCapabilityAllOf) GetDefaultFcoeVlan() int64`

GetDefaultFcoeVlan returns the DefaultFcoeVlan field if non-nil, zero value otherwise.

### GetDefaultFcoeVlanOk

`func (o *CapabilitySwitchCapabilityAllOf) GetDefaultFcoeVlanOk() (*int64, bool)`

GetDefaultFcoeVlanOk returns a tuple with the DefaultFcoeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFcoeVlan

`func (o *CapabilitySwitchCapabilityAllOf) SetDefaultFcoeVlan(v int64)`

SetDefaultFcoeVlan sets DefaultFcoeVlan field to given value.

### HasDefaultFcoeVlan

`func (o *CapabilitySwitchCapabilityAllOf) HasDefaultFcoeVlan() bool`

HasDefaultFcoeVlan returns a boolean if a field has been set.

### GetDynamicVifsSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetDynamicVifsSupported() bool`

GetDynamicVifsSupported returns the DynamicVifsSupported field if non-nil, zero value otherwise.

### GetDynamicVifsSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetDynamicVifsSupportedOk() (*bool, bool)`

GetDynamicVifsSupportedOk returns a tuple with the DynamicVifsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVifsSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetDynamicVifsSupported(v bool)`

SetDynamicVifsSupported sets DynamicVifsSupported field to given value.

### HasDynamicVifsSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasDynamicVifsSupported() bool`

HasDynamicVifsSupported returns a boolean if a field has been set.

### GetFanModulesSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetFanModulesSupported() bool`

GetFanModulesSupported returns the FanModulesSupported field if non-nil, zero value otherwise.

### GetFanModulesSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetFanModulesSupportedOk() (*bool, bool)`

GetFanModulesSupportedOk returns a tuple with the FanModulesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModulesSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetFanModulesSupported(v bool)`

SetFanModulesSupported sets FanModulesSupported field to given value.

### HasFanModulesSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasFanModulesSupported() bool`

HasFanModulesSupported returns a boolean if a field has been set.

### GetFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) GetFcEndHostModeReservedVsans() []CapabilityPortRange`

GetFcEndHostModeReservedVsans returns the FcEndHostModeReservedVsans field if non-nil, zero value otherwise.

### GetFcEndHostModeReservedVsansOk

`func (o *CapabilitySwitchCapabilityAllOf) GetFcEndHostModeReservedVsansOk() (*[]CapabilityPortRange, bool)`

GetFcEndHostModeReservedVsansOk returns a tuple with the FcEndHostModeReservedVsans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) SetFcEndHostModeReservedVsans(v []CapabilityPortRange)`

SetFcEndHostModeReservedVsans sets FcEndHostModeReservedVsans field to given value.

### HasFcEndHostModeReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) HasFcEndHostModeReservedVsans() bool`

HasFcEndHostModeReservedVsans returns a boolean if a field has been set.

### SetFcEndHostModeReservedVsansNil

`func (o *CapabilitySwitchCapabilityAllOf) SetFcEndHostModeReservedVsansNil(b bool)`

 SetFcEndHostModeReservedVsansNil sets the value for FcEndHostModeReservedVsans to be an explicit nil

### UnsetFcEndHostModeReservedVsans
`func (o *CapabilitySwitchCapabilityAllOf) UnsetFcEndHostModeReservedVsans()`

UnsetFcEndHostModeReservedVsans ensures that no value is present for FcEndHostModeReservedVsans, not even an explicit nil
### GetFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetFcUplinkPortsAutoNegotiationSupported() bool`

GetFcUplinkPortsAutoNegotiationSupported returns the FcUplinkPortsAutoNegotiationSupported field if non-nil, zero value otherwise.

### GetFcUplinkPortsAutoNegotiationSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetFcUplinkPortsAutoNegotiationSupportedOk() (*bool, bool)`

GetFcUplinkPortsAutoNegotiationSupportedOk returns a tuple with the FcUplinkPortsAutoNegotiationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetFcUplinkPortsAutoNegotiationSupported(v bool)`

SetFcUplinkPortsAutoNegotiationSupported sets FcUplinkPortsAutoNegotiationSupported field to given value.

### HasFcUplinkPortsAutoNegotiationSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasFcUplinkPortsAutoNegotiationSupported() bool`

HasFcUplinkPortsAutoNegotiationSupported returns a boolean if a field has been set.

### GetLocatorBeaconSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetLocatorBeaconSupported() bool`

GetLocatorBeaconSupported returns the LocatorBeaconSupported field if non-nil, zero value otherwise.

### GetLocatorBeaconSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetLocatorBeaconSupportedOk() (*bool, bool)`

GetLocatorBeaconSupportedOk returns a tuple with the LocatorBeaconSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorBeaconSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetLocatorBeaconSupported(v bool)`

SetLocatorBeaconSupported sets LocatorBeaconSupported field to given value.

### HasLocatorBeaconSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasLocatorBeaconSupported() bool`

HasLocatorBeaconSupported returns a boolean if a field has been set.

### GetMaxPorts

`func (o *CapabilitySwitchCapabilityAllOf) GetMaxPorts() int64`

GetMaxPorts returns the MaxPorts field if non-nil, zero value otherwise.

### GetMaxPortsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetMaxPortsOk() (*int64, bool)`

GetMaxPortsOk returns a tuple with the MaxPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPorts

`func (o *CapabilitySwitchCapabilityAllOf) SetMaxPorts(v int64)`

SetMaxPorts sets MaxPorts field to given value.

### HasMaxPorts

`func (o *CapabilitySwitchCapabilityAllOf) HasMaxPorts() bool`

HasMaxPorts returns a boolean if a field has been set.

### GetMaxSlots

`func (o *CapabilitySwitchCapabilityAllOf) GetMaxSlots() int64`

GetMaxSlots returns the MaxSlots field if non-nil, zero value otherwise.

### GetMaxSlotsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetMaxSlotsOk() (*int64, bool)`

GetMaxSlotsOk returns a tuple with the MaxSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSlots

`func (o *CapabilitySwitchCapabilityAllOf) SetMaxSlots(v int64)`

SetMaxSlots sets MaxSlots field to given value.

### HasMaxSlots

`func (o *CapabilitySwitchCapabilityAllOf) HasMaxSlots() bool`

HasMaxSlots returns a boolean if a field has been set.

### GetMinVersionWithBreakoutSupport

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithBreakoutSupport() string`

GetMinVersionWithBreakoutSupport returns the MinVersionWithBreakoutSupport field if non-nil, zero value otherwise.

### GetMinVersionWithBreakoutSupportOk

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithBreakoutSupportOk() (*string, bool)`

GetMinVersionWithBreakoutSupportOk returns a tuple with the MinVersionWithBreakoutSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionWithBreakoutSupport

`func (o *CapabilitySwitchCapabilityAllOf) SetMinVersionWithBreakoutSupport(v string)`

SetMinVersionWithBreakoutSupport sets MinVersionWithBreakoutSupport field to given value.

### HasMinVersionWithBreakoutSupport

`func (o *CapabilitySwitchCapabilityAllOf) HasMinVersionWithBreakoutSupport() bool`

HasMinVersionWithBreakoutSupport returns a boolean if a field has been set.

### GetMinVersionWithLocatorLedSupport

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithLocatorLedSupport() string`

GetMinVersionWithLocatorLedSupport returns the MinVersionWithLocatorLedSupport field if non-nil, zero value otherwise.

### GetMinVersionWithLocatorLedSupportOk

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithLocatorLedSupportOk() (*string, bool)`

GetMinVersionWithLocatorLedSupportOk returns a tuple with the MinVersionWithLocatorLedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionWithLocatorLedSupport

`func (o *CapabilitySwitchCapabilityAllOf) SetMinVersionWithLocatorLedSupport(v string)`

SetMinVersionWithLocatorLedSupport sets MinVersionWithLocatorLedSupport field to given value.

### HasMinVersionWithLocatorLedSupport

`func (o *CapabilitySwitchCapabilityAllOf) HasMinVersionWithLocatorLedSupport() bool`

HasMinVersionWithLocatorLedSupport returns a boolean if a field has been set.

### GetMinVersionWithNegAuto25g

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithNegAuto25g() string`

GetMinVersionWithNegAuto25g returns the MinVersionWithNegAuto25g field if non-nil, zero value otherwise.

### GetMinVersionWithNegAuto25gOk

`func (o *CapabilitySwitchCapabilityAllOf) GetMinVersionWithNegAuto25gOk() (*string, bool)`

GetMinVersionWithNegAuto25gOk returns a tuple with the MinVersionWithNegAuto25g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionWithNegAuto25g

`func (o *CapabilitySwitchCapabilityAllOf) SetMinVersionWithNegAuto25g(v string)`

SetMinVersionWithNegAuto25g sets MinVersionWithNegAuto25g field to given value.

### HasMinVersionWithNegAuto25g

`func (o *CapabilitySwitchCapabilityAllOf) HasMinVersionWithNegAuto25g() bool`

HasMinVersionWithNegAuto25g returns a boolean if a field has been set.

### GetNetworkLimits

`func (o *CapabilitySwitchCapabilityAllOf) GetNetworkLimits() CapabilitySwitchNetworkLimits`

GetNetworkLimits returns the NetworkLimits field if non-nil, zero value otherwise.

### GetNetworkLimitsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetNetworkLimitsOk() (*CapabilitySwitchNetworkLimits, bool)`

GetNetworkLimitsOk returns a tuple with the NetworkLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLimits

`func (o *CapabilitySwitchCapabilityAllOf) SetNetworkLimits(v CapabilitySwitchNetworkLimits)`

SetNetworkLimits sets NetworkLimits field to given value.

### HasNetworkLimits

`func (o *CapabilitySwitchCapabilityAllOf) HasNetworkLimits() bool`

HasNetworkLimits returns a boolean if a field has been set.

### SetNetworkLimitsNil

`func (o *CapabilitySwitchCapabilityAllOf) SetNetworkLimitsNil(b bool)`

 SetNetworkLimitsNil sets the value for NetworkLimits to be an explicit nil

### UnsetNetworkLimits
`func (o *CapabilitySwitchCapabilityAllOf) UnsetNetworkLimits()`

UnsetNetworkLimits ensures that no value is present for NetworkLimits, not even an explicit nil
### GetPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting100gSpeed() []CapabilityPortRange`

GetPortsSupporting100gSpeed returns the PortsSupporting100gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting100gSpeedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting100gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting100gSpeedOk returns a tuple with the PortsSupporting100gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting100gSpeed(v []CapabilityPortRange)`

SetPortsSupporting100gSpeed sets PortsSupporting100gSpeed field to given value.

### HasPortsSupporting100gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupporting100gSpeed() bool`

HasPortsSupporting100gSpeed returns a boolean if a field has been set.

### SetPortsSupporting100gSpeedNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting100gSpeedNil(b bool)`

 SetPortsSupporting100gSpeedNil sets the value for PortsSupporting100gSpeed to be an explicit nil

### UnsetPortsSupporting100gSpeed
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupporting100gSpeed()`

UnsetPortsSupporting100gSpeed ensures that no value is present for PortsSupporting100gSpeed, not even an explicit nil
### GetPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting10gSpeed() []CapabilityPortRange`

GetPortsSupporting10gSpeed returns the PortsSupporting10gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting10gSpeedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting10gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting10gSpeedOk returns a tuple with the PortsSupporting10gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting10gSpeed(v []CapabilityPortRange)`

SetPortsSupporting10gSpeed sets PortsSupporting10gSpeed field to given value.

### HasPortsSupporting10gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupporting10gSpeed() bool`

HasPortsSupporting10gSpeed returns a boolean if a field has been set.

### SetPortsSupporting10gSpeedNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting10gSpeedNil(b bool)`

 SetPortsSupporting10gSpeedNil sets the value for PortsSupporting10gSpeed to be an explicit nil

### UnsetPortsSupporting10gSpeed
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupporting10gSpeed()`

UnsetPortsSupporting10gSpeed ensures that no value is present for PortsSupporting10gSpeed, not even an explicit nil
### GetPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting1gSpeed() []CapabilityPortRange`

GetPortsSupporting1gSpeed returns the PortsSupporting1gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting1gSpeedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting1gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting1gSpeedOk returns a tuple with the PortsSupporting1gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting1gSpeed(v []CapabilityPortRange)`

SetPortsSupporting1gSpeed sets PortsSupporting1gSpeed field to given value.

### HasPortsSupporting1gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupporting1gSpeed() bool`

HasPortsSupporting1gSpeed returns a boolean if a field has been set.

### SetPortsSupporting1gSpeedNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting1gSpeedNil(b bool)`

 SetPortsSupporting1gSpeedNil sets the value for PortsSupporting1gSpeed to be an explicit nil

### UnsetPortsSupporting1gSpeed
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupporting1gSpeed()`

UnsetPortsSupporting1gSpeed ensures that no value is present for PortsSupporting1gSpeed, not even an explicit nil
### GetPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting25gSpeed() []CapabilityPortRange`

GetPortsSupporting25gSpeed returns the PortsSupporting25gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting25gSpeedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting25gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting25gSpeedOk returns a tuple with the PortsSupporting25gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting25gSpeed(v []CapabilityPortRange)`

SetPortsSupporting25gSpeed sets PortsSupporting25gSpeed field to given value.

### HasPortsSupporting25gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupporting25gSpeed() bool`

HasPortsSupporting25gSpeed returns a boolean if a field has been set.

### SetPortsSupporting25gSpeedNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting25gSpeedNil(b bool)`

 SetPortsSupporting25gSpeedNil sets the value for PortsSupporting25gSpeed to be an explicit nil

### UnsetPortsSupporting25gSpeed
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupporting25gSpeed()`

UnsetPortsSupporting25gSpeed ensures that no value is present for PortsSupporting25gSpeed, not even an explicit nil
### GetPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting40gSpeed() []CapabilityPortRange`

GetPortsSupporting40gSpeed returns the PortsSupporting40gSpeed field if non-nil, zero value otherwise.

### GetPortsSupporting40gSpeedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupporting40gSpeedOk() (*[]CapabilityPortRange, bool)`

GetPortsSupporting40gSpeedOk returns a tuple with the PortsSupporting40gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting40gSpeed(v []CapabilityPortRange)`

SetPortsSupporting40gSpeed sets PortsSupporting40gSpeed field to given value.

### HasPortsSupporting40gSpeed

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupporting40gSpeed() bool`

HasPortsSupporting40gSpeed returns a boolean if a field has been set.

### SetPortsSupporting40gSpeedNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupporting40gSpeedNil(b bool)`

 SetPortsSupporting40gSpeedNil sets the value for PortsSupporting40gSpeed to be an explicit nil

### UnsetPortsSupporting40gSpeed
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupporting40gSpeed()`

UnsetPortsSupporting40gSpeed ensures that no value is present for PortsSupporting40gSpeed, not even an explicit nil
### GetPortsSupportingBreakout

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingBreakout() []CapabilityPortRange`

GetPortsSupportingBreakout returns the PortsSupportingBreakout field if non-nil, zero value otherwise.

### GetPortsSupportingBreakoutOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingBreakoutOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingBreakoutOk returns a tuple with the PortsSupportingBreakout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingBreakout

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingBreakout(v []CapabilityPortRange)`

SetPortsSupportingBreakout sets PortsSupportingBreakout field to given value.

### HasPortsSupportingBreakout

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupportingBreakout() bool`

HasPortsSupportingBreakout returns a boolean if a field has been set.

### SetPortsSupportingBreakoutNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingBreakoutNil(b bool)`

 SetPortsSupportingBreakoutNil sets the value for PortsSupportingBreakout to be an explicit nil

### UnsetPortsSupportingBreakout
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupportingBreakout()`

UnsetPortsSupportingBreakout ensures that no value is present for PortsSupportingBreakout, not even an explicit nil
### GetPortsSupportingFcoe

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingFcoe() []CapabilityPortRange`

GetPortsSupportingFcoe returns the PortsSupportingFcoe field if non-nil, zero value otherwise.

### GetPortsSupportingFcoeOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingFcoeOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingFcoeOk returns a tuple with the PortsSupportingFcoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingFcoe

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingFcoe(v []CapabilityPortRange)`

SetPortsSupportingFcoe sets PortsSupportingFcoe field to given value.

### HasPortsSupportingFcoe

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupportingFcoe() bool`

HasPortsSupportingFcoe returns a boolean if a field has been set.

### SetPortsSupportingFcoeNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingFcoeNil(b bool)`

 SetPortsSupportingFcoeNil sets the value for PortsSupportingFcoe to be an explicit nil

### UnsetPortsSupportingFcoe
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupportingFcoe()`

UnsetPortsSupportingFcoe ensures that no value is present for PortsSupportingFcoe, not even an explicit nil
### GetPortsSupportingServerRole

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingServerRole() []CapabilityPortRange`

GetPortsSupportingServerRole returns the PortsSupportingServerRole field if non-nil, zero value otherwise.

### GetPortsSupportingServerRoleOk

`func (o *CapabilitySwitchCapabilityAllOf) GetPortsSupportingServerRoleOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingServerRoleOk returns a tuple with the PortsSupportingServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingServerRole

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingServerRole(v []CapabilityPortRange)`

SetPortsSupportingServerRole sets PortsSupportingServerRole field to given value.

### HasPortsSupportingServerRole

`func (o *CapabilitySwitchCapabilityAllOf) HasPortsSupportingServerRole() bool`

HasPortsSupportingServerRole returns a boolean if a field has been set.

### SetPortsSupportingServerRoleNil

`func (o *CapabilitySwitchCapabilityAllOf) SetPortsSupportingServerRoleNil(b bool)`

 SetPortsSupportingServerRoleNil sets the value for PortsSupportingServerRole to be an explicit nil

### UnsetPortsSupportingServerRole
`func (o *CapabilitySwitchCapabilityAllOf) UnsetPortsSupportingServerRole()`

UnsetPortsSupportingServerRole ensures that no value is present for PortsSupportingServerRole, not even an explicit nil
### GetReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) GetReservedVsans() []CapabilityPortRange`

GetReservedVsans returns the ReservedVsans field if non-nil, zero value otherwise.

### GetReservedVsansOk

`func (o *CapabilitySwitchCapabilityAllOf) GetReservedVsansOk() (*[]CapabilityPortRange, bool)`

GetReservedVsansOk returns a tuple with the ReservedVsans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) SetReservedVsans(v []CapabilityPortRange)`

SetReservedVsans sets ReservedVsans field to given value.

### HasReservedVsans

`func (o *CapabilitySwitchCapabilityAllOf) HasReservedVsans() bool`

HasReservedVsans returns a boolean if a field has been set.

### SetReservedVsansNil

`func (o *CapabilitySwitchCapabilityAllOf) SetReservedVsansNil(b bool)`

 SetReservedVsansNil sets the value for ReservedVsans to be an explicit nil

### UnsetReservedVsans
`func (o *CapabilitySwitchCapabilityAllOf) UnsetReservedVsans()`

UnsetReservedVsans ensures that no value is present for ReservedVsans, not even an explicit nil
### GetSerenoNetflowSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetSerenoNetflowSupported() bool`

GetSerenoNetflowSupported returns the SerenoNetflowSupported field if non-nil, zero value otherwise.

### GetSerenoNetflowSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetSerenoNetflowSupportedOk() (*bool, bool)`

GetSerenoNetflowSupportedOk returns a tuple with the SerenoNetflowSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerenoNetflowSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetSerenoNetflowSupported(v bool)`

SetSerenoNetflowSupported sets SerenoNetflowSupported field to given value.

### HasSerenoNetflowSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasSerenoNetflowSupported() bool`

HasSerenoNetflowSupported returns a boolean if a field has been set.

### GetServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapabilityAllOf) GetServerRoleSupportedOnBreakout() []string`

GetServerRoleSupportedOnBreakout returns the ServerRoleSupportedOnBreakout field if non-nil, zero value otherwise.

### GetServerRoleSupportedOnBreakoutOk

`func (o *CapabilitySwitchCapabilityAllOf) GetServerRoleSupportedOnBreakoutOk() (*[]string, bool)`

GetServerRoleSupportedOnBreakoutOk returns a tuple with the ServerRoleSupportedOnBreakout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapabilityAllOf) SetServerRoleSupportedOnBreakout(v []string)`

SetServerRoleSupportedOnBreakout sets ServerRoleSupportedOnBreakout field to given value.

### HasServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapabilityAllOf) HasServerRoleSupportedOnBreakout() bool`

HasServerRoleSupportedOnBreakout returns a boolean if a field has been set.

### SetServerRoleSupportedOnBreakoutNil

`func (o *CapabilitySwitchCapabilityAllOf) SetServerRoleSupportedOnBreakoutNil(b bool)`

 SetServerRoleSupportedOnBreakoutNil sets the value for ServerRoleSupportedOnBreakout to be an explicit nil

### UnsetServerRoleSupportedOnBreakout
`func (o *CapabilitySwitchCapabilityAllOf) UnsetServerRoleSupportedOnBreakout()`

UnsetServerRoleSupportedOnBreakout ensures that no value is present for ServerRoleSupportedOnBreakout, not even an explicit nil
### GetStorageLimits

`func (o *CapabilitySwitchCapabilityAllOf) GetStorageLimits() CapabilitySwitchStorageLimits`

GetStorageLimits returns the StorageLimits field if non-nil, zero value otherwise.

### GetStorageLimitsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetStorageLimitsOk() (*CapabilitySwitchStorageLimits, bool)`

GetStorageLimitsOk returns a tuple with the StorageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimits

`func (o *CapabilitySwitchCapabilityAllOf) SetStorageLimits(v CapabilitySwitchStorageLimits)`

SetStorageLimits sets StorageLimits field to given value.

### HasStorageLimits

`func (o *CapabilitySwitchCapabilityAllOf) HasStorageLimits() bool`

HasStorageLimits returns a boolean if a field has been set.

### SetStorageLimitsNil

`func (o *CapabilitySwitchCapabilityAllOf) SetStorageLimitsNil(b bool)`

 SetStorageLimitsNil sets the value for StorageLimits to be an explicit nil

### UnsetStorageLimits
`func (o *CapabilitySwitchCapabilityAllOf) UnsetStorageLimits()`

UnsetStorageLimits ensures that no value is present for StorageLimits, not even an explicit nil
### GetSwitchingModeCapabilities

`func (o *CapabilitySwitchCapabilityAllOf) GetSwitchingModeCapabilities() []CapabilitySwitchingModeCapability`

GetSwitchingModeCapabilities returns the SwitchingModeCapabilities field if non-nil, zero value otherwise.

### GetSwitchingModeCapabilitiesOk

`func (o *CapabilitySwitchCapabilityAllOf) GetSwitchingModeCapabilitiesOk() (*[]CapabilitySwitchingModeCapability, bool)`

GetSwitchingModeCapabilitiesOk returns a tuple with the SwitchingModeCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchingModeCapabilities

`func (o *CapabilitySwitchCapabilityAllOf) SetSwitchingModeCapabilities(v []CapabilitySwitchingModeCapability)`

SetSwitchingModeCapabilities sets SwitchingModeCapabilities field to given value.

### HasSwitchingModeCapabilities

`func (o *CapabilitySwitchCapabilityAllOf) HasSwitchingModeCapabilities() bool`

HasSwitchingModeCapabilities returns a boolean if a field has been set.

### SetSwitchingModeCapabilitiesNil

`func (o *CapabilitySwitchCapabilityAllOf) SetSwitchingModeCapabilitiesNil(b bool)`

 SetSwitchingModeCapabilitiesNil sets the value for SwitchingModeCapabilities to be an explicit nil

### UnsetSwitchingModeCapabilities
`func (o *CapabilitySwitchCapabilityAllOf) UnsetSwitchingModeCapabilities()`

UnsetSwitchingModeCapabilities ensures that no value is present for SwitchingModeCapabilities, not even an explicit nil
### GetSystemLimits

`func (o *CapabilitySwitchCapabilityAllOf) GetSystemLimits() CapabilitySwitchSystemLimits`

GetSystemLimits returns the SystemLimits field if non-nil, zero value otherwise.

### GetSystemLimitsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetSystemLimitsOk() (*CapabilitySwitchSystemLimits, bool)`

GetSystemLimitsOk returns a tuple with the SystemLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLimits

`func (o *CapabilitySwitchCapabilityAllOf) SetSystemLimits(v CapabilitySwitchSystemLimits)`

SetSystemLimits sets SystemLimits field to given value.

### HasSystemLimits

`func (o *CapabilitySwitchCapabilityAllOf) HasSystemLimits() bool`

HasSystemLimits returns a boolean if a field has been set.

### SetSystemLimitsNil

`func (o *CapabilitySwitchCapabilityAllOf) SetSystemLimitsNil(b bool)`

 SetSystemLimitsNil sets the value for SystemLimits to be an explicit nil

### UnsetSystemLimits
`func (o *CapabilitySwitchCapabilityAllOf) UnsetSystemLimits()`

UnsetSystemLimits ensures that no value is present for SystemLimits, not even an explicit nil
### GetUnifiedPorts

`func (o *CapabilitySwitchCapabilityAllOf) GetUnifiedPorts() []CapabilityPortRange`

GetUnifiedPorts returns the UnifiedPorts field if non-nil, zero value otherwise.

### GetUnifiedPortsOk

`func (o *CapabilitySwitchCapabilityAllOf) GetUnifiedPortsOk() (*[]CapabilityPortRange, bool)`

GetUnifiedPortsOk returns a tuple with the UnifiedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedPorts

`func (o *CapabilitySwitchCapabilityAllOf) SetUnifiedPorts(v []CapabilityPortRange)`

SetUnifiedPorts sets UnifiedPorts field to given value.

### HasUnifiedPorts

`func (o *CapabilitySwitchCapabilityAllOf) HasUnifiedPorts() bool`

HasUnifiedPorts returns a boolean if a field has been set.

### SetUnifiedPortsNil

`func (o *CapabilitySwitchCapabilityAllOf) SetUnifiedPortsNil(b bool)`

 SetUnifiedPortsNil sets the value for UnifiedPorts to be an explicit nil

### UnsetUnifiedPorts
`func (o *CapabilitySwitchCapabilityAllOf) UnsetUnifiedPorts()`

UnsetUnifiedPorts ensures that no value is present for UnifiedPorts, not even an explicit nil
### GetUnifiedRule

`func (o *CapabilitySwitchCapabilityAllOf) GetUnifiedRule() string`

GetUnifiedRule returns the UnifiedRule field if non-nil, zero value otherwise.

### GetUnifiedRuleOk

`func (o *CapabilitySwitchCapabilityAllOf) GetUnifiedRuleOk() (*string, bool)`

GetUnifiedRuleOk returns a tuple with the UnifiedRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedRule

`func (o *CapabilitySwitchCapabilityAllOf) SetUnifiedRule(v string)`

SetUnifiedRule sets UnifiedRule field to given value.

### HasUnifiedRule

`func (o *CapabilitySwitchCapabilityAllOf) HasUnifiedRule() bool`

HasUnifiedRule returns a boolean if a field has been set.

### GetUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapabilityAllOf) GetUplinkAdminPortSpeedNegAuto25GbpsSupported() bool`

GetUplinkAdminPortSpeedNegAuto25GbpsSupported returns the UplinkAdminPortSpeedNegAuto25GbpsSupported field if non-nil, zero value otherwise.

### GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk

`func (o *CapabilitySwitchCapabilityAllOf) GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk() (*bool, bool)`

GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk returns a tuple with the UplinkAdminPortSpeedNegAuto25GbpsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapabilityAllOf) SetUplinkAdminPortSpeedNegAuto25GbpsSupported(v bool)`

SetUplinkAdminPortSpeedNegAuto25GbpsSupported sets UplinkAdminPortSpeedNegAuto25GbpsSupported field to given value.

### HasUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapabilityAllOf) HasUplinkAdminPortSpeedNegAuto25GbpsSupported() bool`

HasUplinkAdminPortSpeedNegAuto25GbpsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


