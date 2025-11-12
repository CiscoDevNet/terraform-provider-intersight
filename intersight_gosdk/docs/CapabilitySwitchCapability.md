# CapabilitySwitchCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchCapability"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchCapability"]
**AllowedUplinkPcIdRange** | Pointer to [**NullableCapabilityPcIdRange**](CapabilityPcIdRange.md) |  | [optional] 
**DefaultFcoeVlan** | Pointer to **int64** | Default Fcoe VLAN associated with this switch. | [optional] [readonly] 
**DynamicVifsSupported** | Pointer to **bool** | Dynamic VIFs support on this switch. | [optional] [readonly] 
**FanModulesSupported** | Pointer to **bool** | Fan Modules support on this switch. | [optional] [readonly] 
**FcEndHostModeReservedVsans** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**FcUplinkPortsAutoNegotiationSupported** | Pointer to **bool** | Fc Uplink ports auto negotiation speed support on this switch. | [optional] [readonly] 
**ImmControlsVpcompression** | Pointer to **bool** | VlanPort Compression is controlled by IMM. | [optional] [readonly] 
**InterClusterLinkVlanSupported** | Pointer to **bool** | Inter cluster link vlan support on this switch. | [optional] [readonly] 
**LocatorBeaconSupported** | Pointer to **bool** | Locator Beacon LED support on this switch. | [optional] [readonly] 
**MacsecSupportedPorts** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**MaxPorts** | Pointer to **int64** | Maximum allowed physical ports on this switch. | [optional] [readonly] 
**MaxSlots** | Pointer to **int64** | Maximum allowed physical slots on this switch. | [optional] [readonly] 
**MinVersionMapForSwitchFeatures** | Pointer to [**[]FirmwareFeatureVersionMap**](FirmwareFeatureVersionMap.md) |  | [optional] 
**MinVersionMapWithBreakoutSupport** | Pointer to [**NullableFirmwareVersionMap**](FirmwareVersionMap.md) |  | [optional] 
**MinVersionMapWithLocatorLedSupport** | Pointer to [**NullableFirmwareVersionMap**](FirmwareVersionMap.md) |  | [optional] 
**MinVersionMapWithNegAuto25gSupport** | Pointer to [**NullableFirmwareVersionMap**](FirmwareVersionMap.md) |  | [optional] 
**NetworkLimits** | Pointer to [**NullableCapabilitySwitchNetworkLimits**](CapabilitySwitchNetworkLimits.md) |  | [optional] 
**PortsSupporting100gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting10gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting1gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting25gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupporting40gSpeed** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingApplianceRole** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingBreakout** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingFcoe** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingInterClusterLink** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**PortsSupportingServerRole** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**ReservedVsans** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**SerenoNetflowSupported** | Pointer to **bool** | Sereno Adaptor with Netflow support on this switch. | [optional] [readonly] 
**ServerRoleSupportedOnBreakout** | Pointer to **[]string** |  | [optional] 
**StorageLimits** | Pointer to [**NullableCapabilitySwitchStorageLimits**](CapabilitySwitchStorageLimits.md) |  | [optional] 
**SwitchingModeCapabilities** | Pointer to [**[]CapabilitySwitchingModeCapability**](CapabilitySwitchingModeCapability.md) |  | [optional] 
**SystemLimits** | Pointer to [**NullableCapabilitySwitchSystemLimits**](CapabilitySwitchSystemLimits.md) |  | [optional] 
**UnSupportedEquipmentModel** | Pointer to **[]string** |  | [optional] 
**UnifiedPorts** | Pointer to [**[]CapabilityPortRange**](CapabilityPortRange.md) |  | [optional] 
**UnifiedRule** | Pointer to **string** | The Slider rule for Unified ports on this switch. | [optional] [readonly] 
**UplinkAdminPortSpeedNegAuto25GbpsSupported** | Pointer to **bool** | &#39;Negotiate Auto 25000&#39; admin speed support on this switch for port or port-channel with Ethernet Uplink/Appliance/FCoE Uplink roles. | [optional] [readonly] 

## Methods

### NewCapabilitySwitchCapability

`func NewCapabilitySwitchCapability(classId string, objectType string, ) *CapabilitySwitchCapability`

NewCapabilitySwitchCapability instantiates a new CapabilitySwitchCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchCapabilityWithDefaults

`func NewCapabilitySwitchCapabilityWithDefaults() *CapabilitySwitchCapability`

NewCapabilitySwitchCapabilityWithDefaults instantiates a new CapabilitySwitchCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchCapability) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchCapability) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchCapability) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchCapability) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchCapability) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchCapability) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedUplinkPcIdRange

`func (o *CapabilitySwitchCapability) GetAllowedUplinkPcIdRange() CapabilityPcIdRange`

GetAllowedUplinkPcIdRange returns the AllowedUplinkPcIdRange field if non-nil, zero value otherwise.

### GetAllowedUplinkPcIdRangeOk

`func (o *CapabilitySwitchCapability) GetAllowedUplinkPcIdRangeOk() (*CapabilityPcIdRange, bool)`

GetAllowedUplinkPcIdRangeOk returns a tuple with the AllowedUplinkPcIdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUplinkPcIdRange

`func (o *CapabilitySwitchCapability) SetAllowedUplinkPcIdRange(v CapabilityPcIdRange)`

SetAllowedUplinkPcIdRange sets AllowedUplinkPcIdRange field to given value.

### HasAllowedUplinkPcIdRange

`func (o *CapabilitySwitchCapability) HasAllowedUplinkPcIdRange() bool`

HasAllowedUplinkPcIdRange returns a boolean if a field has been set.

### SetAllowedUplinkPcIdRangeNil

`func (o *CapabilitySwitchCapability) SetAllowedUplinkPcIdRangeNil(b bool)`

 SetAllowedUplinkPcIdRangeNil sets the value for AllowedUplinkPcIdRange to be an explicit nil

### UnsetAllowedUplinkPcIdRange
`func (o *CapabilitySwitchCapability) UnsetAllowedUplinkPcIdRange()`

UnsetAllowedUplinkPcIdRange ensures that no value is present for AllowedUplinkPcIdRange, not even an explicit nil
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

### SetFcEndHostModeReservedVsansNil

`func (o *CapabilitySwitchCapability) SetFcEndHostModeReservedVsansNil(b bool)`

 SetFcEndHostModeReservedVsansNil sets the value for FcEndHostModeReservedVsans to be an explicit nil

### UnsetFcEndHostModeReservedVsans
`func (o *CapabilitySwitchCapability) UnsetFcEndHostModeReservedVsans()`

UnsetFcEndHostModeReservedVsans ensures that no value is present for FcEndHostModeReservedVsans, not even an explicit nil
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

### GetImmControlsVpcompression

`func (o *CapabilitySwitchCapability) GetImmControlsVpcompression() bool`

GetImmControlsVpcompression returns the ImmControlsVpcompression field if non-nil, zero value otherwise.

### GetImmControlsVpcompressionOk

`func (o *CapabilitySwitchCapability) GetImmControlsVpcompressionOk() (*bool, bool)`

GetImmControlsVpcompressionOk returns a tuple with the ImmControlsVpcompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmControlsVpcompression

`func (o *CapabilitySwitchCapability) SetImmControlsVpcompression(v bool)`

SetImmControlsVpcompression sets ImmControlsVpcompression field to given value.

### HasImmControlsVpcompression

`func (o *CapabilitySwitchCapability) HasImmControlsVpcompression() bool`

HasImmControlsVpcompression returns a boolean if a field has been set.

### GetInterClusterLinkVlanSupported

`func (o *CapabilitySwitchCapability) GetInterClusterLinkVlanSupported() bool`

GetInterClusterLinkVlanSupported returns the InterClusterLinkVlanSupported field if non-nil, zero value otherwise.

### GetInterClusterLinkVlanSupportedOk

`func (o *CapabilitySwitchCapability) GetInterClusterLinkVlanSupportedOk() (*bool, bool)`

GetInterClusterLinkVlanSupportedOk returns a tuple with the InterClusterLinkVlanSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterClusterLinkVlanSupported

`func (o *CapabilitySwitchCapability) SetInterClusterLinkVlanSupported(v bool)`

SetInterClusterLinkVlanSupported sets InterClusterLinkVlanSupported field to given value.

### HasInterClusterLinkVlanSupported

`func (o *CapabilitySwitchCapability) HasInterClusterLinkVlanSupported() bool`

HasInterClusterLinkVlanSupported returns a boolean if a field has been set.

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

### GetMacsecSupportedPorts

`func (o *CapabilitySwitchCapability) GetMacsecSupportedPorts() []CapabilityPortRange`

GetMacsecSupportedPorts returns the MacsecSupportedPorts field if non-nil, zero value otherwise.

### GetMacsecSupportedPortsOk

`func (o *CapabilitySwitchCapability) GetMacsecSupportedPortsOk() (*[]CapabilityPortRange, bool)`

GetMacsecSupportedPortsOk returns a tuple with the MacsecSupportedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsecSupportedPorts

`func (o *CapabilitySwitchCapability) SetMacsecSupportedPorts(v []CapabilityPortRange)`

SetMacsecSupportedPorts sets MacsecSupportedPorts field to given value.

### HasMacsecSupportedPorts

`func (o *CapabilitySwitchCapability) HasMacsecSupportedPorts() bool`

HasMacsecSupportedPorts returns a boolean if a field has been set.

### SetMacsecSupportedPortsNil

`func (o *CapabilitySwitchCapability) SetMacsecSupportedPortsNil(b bool)`

 SetMacsecSupportedPortsNil sets the value for MacsecSupportedPorts to be an explicit nil

### UnsetMacsecSupportedPorts
`func (o *CapabilitySwitchCapability) UnsetMacsecSupportedPorts()`

UnsetMacsecSupportedPorts ensures that no value is present for MacsecSupportedPorts, not even an explicit nil
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

### GetMinVersionMapForSwitchFeatures

`func (o *CapabilitySwitchCapability) GetMinVersionMapForSwitchFeatures() []FirmwareFeatureVersionMap`

GetMinVersionMapForSwitchFeatures returns the MinVersionMapForSwitchFeatures field if non-nil, zero value otherwise.

### GetMinVersionMapForSwitchFeaturesOk

`func (o *CapabilitySwitchCapability) GetMinVersionMapForSwitchFeaturesOk() (*[]FirmwareFeatureVersionMap, bool)`

GetMinVersionMapForSwitchFeaturesOk returns a tuple with the MinVersionMapForSwitchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionMapForSwitchFeatures

`func (o *CapabilitySwitchCapability) SetMinVersionMapForSwitchFeatures(v []FirmwareFeatureVersionMap)`

SetMinVersionMapForSwitchFeatures sets MinVersionMapForSwitchFeatures field to given value.

### HasMinVersionMapForSwitchFeatures

`func (o *CapabilitySwitchCapability) HasMinVersionMapForSwitchFeatures() bool`

HasMinVersionMapForSwitchFeatures returns a boolean if a field has been set.

### SetMinVersionMapForSwitchFeaturesNil

`func (o *CapabilitySwitchCapability) SetMinVersionMapForSwitchFeaturesNil(b bool)`

 SetMinVersionMapForSwitchFeaturesNil sets the value for MinVersionMapForSwitchFeatures to be an explicit nil

### UnsetMinVersionMapForSwitchFeatures
`func (o *CapabilitySwitchCapability) UnsetMinVersionMapForSwitchFeatures()`

UnsetMinVersionMapForSwitchFeatures ensures that no value is present for MinVersionMapForSwitchFeatures, not even an explicit nil
### GetMinVersionMapWithBreakoutSupport

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithBreakoutSupport() FirmwareVersionMap`

GetMinVersionMapWithBreakoutSupport returns the MinVersionMapWithBreakoutSupport field if non-nil, zero value otherwise.

### GetMinVersionMapWithBreakoutSupportOk

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithBreakoutSupportOk() (*FirmwareVersionMap, bool)`

GetMinVersionMapWithBreakoutSupportOk returns a tuple with the MinVersionMapWithBreakoutSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionMapWithBreakoutSupport

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithBreakoutSupport(v FirmwareVersionMap)`

SetMinVersionMapWithBreakoutSupport sets MinVersionMapWithBreakoutSupport field to given value.

### HasMinVersionMapWithBreakoutSupport

`func (o *CapabilitySwitchCapability) HasMinVersionMapWithBreakoutSupport() bool`

HasMinVersionMapWithBreakoutSupport returns a boolean if a field has been set.

### SetMinVersionMapWithBreakoutSupportNil

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithBreakoutSupportNil(b bool)`

 SetMinVersionMapWithBreakoutSupportNil sets the value for MinVersionMapWithBreakoutSupport to be an explicit nil

### UnsetMinVersionMapWithBreakoutSupport
`func (o *CapabilitySwitchCapability) UnsetMinVersionMapWithBreakoutSupport()`

UnsetMinVersionMapWithBreakoutSupport ensures that no value is present for MinVersionMapWithBreakoutSupport, not even an explicit nil
### GetMinVersionMapWithLocatorLedSupport

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithLocatorLedSupport() FirmwareVersionMap`

GetMinVersionMapWithLocatorLedSupport returns the MinVersionMapWithLocatorLedSupport field if non-nil, zero value otherwise.

### GetMinVersionMapWithLocatorLedSupportOk

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithLocatorLedSupportOk() (*FirmwareVersionMap, bool)`

GetMinVersionMapWithLocatorLedSupportOk returns a tuple with the MinVersionMapWithLocatorLedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionMapWithLocatorLedSupport

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithLocatorLedSupport(v FirmwareVersionMap)`

SetMinVersionMapWithLocatorLedSupport sets MinVersionMapWithLocatorLedSupport field to given value.

### HasMinVersionMapWithLocatorLedSupport

`func (o *CapabilitySwitchCapability) HasMinVersionMapWithLocatorLedSupport() bool`

HasMinVersionMapWithLocatorLedSupport returns a boolean if a field has been set.

### SetMinVersionMapWithLocatorLedSupportNil

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithLocatorLedSupportNil(b bool)`

 SetMinVersionMapWithLocatorLedSupportNil sets the value for MinVersionMapWithLocatorLedSupport to be an explicit nil

### UnsetMinVersionMapWithLocatorLedSupport
`func (o *CapabilitySwitchCapability) UnsetMinVersionMapWithLocatorLedSupport()`

UnsetMinVersionMapWithLocatorLedSupport ensures that no value is present for MinVersionMapWithLocatorLedSupport, not even an explicit nil
### GetMinVersionMapWithNegAuto25gSupport

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithNegAuto25gSupport() FirmwareVersionMap`

GetMinVersionMapWithNegAuto25gSupport returns the MinVersionMapWithNegAuto25gSupport field if non-nil, zero value otherwise.

### GetMinVersionMapWithNegAuto25gSupportOk

`func (o *CapabilitySwitchCapability) GetMinVersionMapWithNegAuto25gSupportOk() (*FirmwareVersionMap, bool)`

GetMinVersionMapWithNegAuto25gSupportOk returns a tuple with the MinVersionMapWithNegAuto25gSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersionMapWithNegAuto25gSupport

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithNegAuto25gSupport(v FirmwareVersionMap)`

SetMinVersionMapWithNegAuto25gSupport sets MinVersionMapWithNegAuto25gSupport field to given value.

### HasMinVersionMapWithNegAuto25gSupport

`func (o *CapabilitySwitchCapability) HasMinVersionMapWithNegAuto25gSupport() bool`

HasMinVersionMapWithNegAuto25gSupport returns a boolean if a field has been set.

### SetMinVersionMapWithNegAuto25gSupportNil

`func (o *CapabilitySwitchCapability) SetMinVersionMapWithNegAuto25gSupportNil(b bool)`

 SetMinVersionMapWithNegAuto25gSupportNil sets the value for MinVersionMapWithNegAuto25gSupport to be an explicit nil

### UnsetMinVersionMapWithNegAuto25gSupport
`func (o *CapabilitySwitchCapability) UnsetMinVersionMapWithNegAuto25gSupport()`

UnsetMinVersionMapWithNegAuto25gSupport ensures that no value is present for MinVersionMapWithNegAuto25gSupport, not even an explicit nil
### GetNetworkLimits

`func (o *CapabilitySwitchCapability) GetNetworkLimits() CapabilitySwitchNetworkLimits`

GetNetworkLimits returns the NetworkLimits field if non-nil, zero value otherwise.

### GetNetworkLimitsOk

`func (o *CapabilitySwitchCapability) GetNetworkLimitsOk() (*CapabilitySwitchNetworkLimits, bool)`

GetNetworkLimitsOk returns a tuple with the NetworkLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLimits

`func (o *CapabilitySwitchCapability) SetNetworkLimits(v CapabilitySwitchNetworkLimits)`

SetNetworkLimits sets NetworkLimits field to given value.

### HasNetworkLimits

`func (o *CapabilitySwitchCapability) HasNetworkLimits() bool`

HasNetworkLimits returns a boolean if a field has been set.

### SetNetworkLimitsNil

`func (o *CapabilitySwitchCapability) SetNetworkLimitsNil(b bool)`

 SetNetworkLimitsNil sets the value for NetworkLimits to be an explicit nil

### UnsetNetworkLimits
`func (o *CapabilitySwitchCapability) UnsetNetworkLimits()`

UnsetNetworkLimits ensures that no value is present for NetworkLimits, not even an explicit nil
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

### SetPortsSupporting100gSpeedNil

`func (o *CapabilitySwitchCapability) SetPortsSupporting100gSpeedNil(b bool)`

 SetPortsSupporting100gSpeedNil sets the value for PortsSupporting100gSpeed to be an explicit nil

### UnsetPortsSupporting100gSpeed
`func (o *CapabilitySwitchCapability) UnsetPortsSupporting100gSpeed()`

UnsetPortsSupporting100gSpeed ensures that no value is present for PortsSupporting100gSpeed, not even an explicit nil
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

### SetPortsSupporting10gSpeedNil

`func (o *CapabilitySwitchCapability) SetPortsSupporting10gSpeedNil(b bool)`

 SetPortsSupporting10gSpeedNil sets the value for PortsSupporting10gSpeed to be an explicit nil

### UnsetPortsSupporting10gSpeed
`func (o *CapabilitySwitchCapability) UnsetPortsSupporting10gSpeed()`

UnsetPortsSupporting10gSpeed ensures that no value is present for PortsSupporting10gSpeed, not even an explicit nil
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

### SetPortsSupporting1gSpeedNil

`func (o *CapabilitySwitchCapability) SetPortsSupporting1gSpeedNil(b bool)`

 SetPortsSupporting1gSpeedNil sets the value for PortsSupporting1gSpeed to be an explicit nil

### UnsetPortsSupporting1gSpeed
`func (o *CapabilitySwitchCapability) UnsetPortsSupporting1gSpeed()`

UnsetPortsSupporting1gSpeed ensures that no value is present for PortsSupporting1gSpeed, not even an explicit nil
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

### SetPortsSupporting25gSpeedNil

`func (o *CapabilitySwitchCapability) SetPortsSupporting25gSpeedNil(b bool)`

 SetPortsSupporting25gSpeedNil sets the value for PortsSupporting25gSpeed to be an explicit nil

### UnsetPortsSupporting25gSpeed
`func (o *CapabilitySwitchCapability) UnsetPortsSupporting25gSpeed()`

UnsetPortsSupporting25gSpeed ensures that no value is present for PortsSupporting25gSpeed, not even an explicit nil
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

### SetPortsSupporting40gSpeedNil

`func (o *CapabilitySwitchCapability) SetPortsSupporting40gSpeedNil(b bool)`

 SetPortsSupporting40gSpeedNil sets the value for PortsSupporting40gSpeed to be an explicit nil

### UnsetPortsSupporting40gSpeed
`func (o *CapabilitySwitchCapability) UnsetPortsSupporting40gSpeed()`

UnsetPortsSupporting40gSpeed ensures that no value is present for PortsSupporting40gSpeed, not even an explicit nil
### GetPortsSupportingApplianceRole

`func (o *CapabilitySwitchCapability) GetPortsSupportingApplianceRole() []CapabilityPortRange`

GetPortsSupportingApplianceRole returns the PortsSupportingApplianceRole field if non-nil, zero value otherwise.

### GetPortsSupportingApplianceRoleOk

`func (o *CapabilitySwitchCapability) GetPortsSupportingApplianceRoleOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingApplianceRoleOk returns a tuple with the PortsSupportingApplianceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingApplianceRole

`func (o *CapabilitySwitchCapability) SetPortsSupportingApplianceRole(v []CapabilityPortRange)`

SetPortsSupportingApplianceRole sets PortsSupportingApplianceRole field to given value.

### HasPortsSupportingApplianceRole

`func (o *CapabilitySwitchCapability) HasPortsSupportingApplianceRole() bool`

HasPortsSupportingApplianceRole returns a boolean if a field has been set.

### SetPortsSupportingApplianceRoleNil

`func (o *CapabilitySwitchCapability) SetPortsSupportingApplianceRoleNil(b bool)`

 SetPortsSupportingApplianceRoleNil sets the value for PortsSupportingApplianceRole to be an explicit nil

### UnsetPortsSupportingApplianceRole
`func (o *CapabilitySwitchCapability) UnsetPortsSupportingApplianceRole()`

UnsetPortsSupportingApplianceRole ensures that no value is present for PortsSupportingApplianceRole, not even an explicit nil
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

### SetPortsSupportingBreakoutNil

`func (o *CapabilitySwitchCapability) SetPortsSupportingBreakoutNil(b bool)`

 SetPortsSupportingBreakoutNil sets the value for PortsSupportingBreakout to be an explicit nil

### UnsetPortsSupportingBreakout
`func (o *CapabilitySwitchCapability) UnsetPortsSupportingBreakout()`

UnsetPortsSupportingBreakout ensures that no value is present for PortsSupportingBreakout, not even an explicit nil
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

### SetPortsSupportingFcoeNil

`func (o *CapabilitySwitchCapability) SetPortsSupportingFcoeNil(b bool)`

 SetPortsSupportingFcoeNil sets the value for PortsSupportingFcoe to be an explicit nil

### UnsetPortsSupportingFcoe
`func (o *CapabilitySwitchCapability) UnsetPortsSupportingFcoe()`

UnsetPortsSupportingFcoe ensures that no value is present for PortsSupportingFcoe, not even an explicit nil
### GetPortsSupportingInterClusterLink

`func (o *CapabilitySwitchCapability) GetPortsSupportingInterClusterLink() []CapabilityPortRange`

GetPortsSupportingInterClusterLink returns the PortsSupportingInterClusterLink field if non-nil, zero value otherwise.

### GetPortsSupportingInterClusterLinkOk

`func (o *CapabilitySwitchCapability) GetPortsSupportingInterClusterLinkOk() (*[]CapabilityPortRange, bool)`

GetPortsSupportingInterClusterLinkOk returns a tuple with the PortsSupportingInterClusterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsSupportingInterClusterLink

`func (o *CapabilitySwitchCapability) SetPortsSupportingInterClusterLink(v []CapabilityPortRange)`

SetPortsSupportingInterClusterLink sets PortsSupportingInterClusterLink field to given value.

### HasPortsSupportingInterClusterLink

`func (o *CapabilitySwitchCapability) HasPortsSupportingInterClusterLink() bool`

HasPortsSupportingInterClusterLink returns a boolean if a field has been set.

### SetPortsSupportingInterClusterLinkNil

`func (o *CapabilitySwitchCapability) SetPortsSupportingInterClusterLinkNil(b bool)`

 SetPortsSupportingInterClusterLinkNil sets the value for PortsSupportingInterClusterLink to be an explicit nil

### UnsetPortsSupportingInterClusterLink
`func (o *CapabilitySwitchCapability) UnsetPortsSupportingInterClusterLink()`

UnsetPortsSupportingInterClusterLink ensures that no value is present for PortsSupportingInterClusterLink, not even an explicit nil
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

### SetPortsSupportingServerRoleNil

`func (o *CapabilitySwitchCapability) SetPortsSupportingServerRoleNil(b bool)`

 SetPortsSupportingServerRoleNil sets the value for PortsSupportingServerRole to be an explicit nil

### UnsetPortsSupportingServerRole
`func (o *CapabilitySwitchCapability) UnsetPortsSupportingServerRole()`

UnsetPortsSupportingServerRole ensures that no value is present for PortsSupportingServerRole, not even an explicit nil
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

### SetReservedVsansNil

`func (o *CapabilitySwitchCapability) SetReservedVsansNil(b bool)`

 SetReservedVsansNil sets the value for ReservedVsans to be an explicit nil

### UnsetReservedVsans
`func (o *CapabilitySwitchCapability) UnsetReservedVsans()`

UnsetReservedVsans ensures that no value is present for ReservedVsans, not even an explicit nil
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

### GetServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapability) GetServerRoleSupportedOnBreakout() []string`

GetServerRoleSupportedOnBreakout returns the ServerRoleSupportedOnBreakout field if non-nil, zero value otherwise.

### GetServerRoleSupportedOnBreakoutOk

`func (o *CapabilitySwitchCapability) GetServerRoleSupportedOnBreakoutOk() (*[]string, bool)`

GetServerRoleSupportedOnBreakoutOk returns a tuple with the ServerRoleSupportedOnBreakout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapability) SetServerRoleSupportedOnBreakout(v []string)`

SetServerRoleSupportedOnBreakout sets ServerRoleSupportedOnBreakout field to given value.

### HasServerRoleSupportedOnBreakout

`func (o *CapabilitySwitchCapability) HasServerRoleSupportedOnBreakout() bool`

HasServerRoleSupportedOnBreakout returns a boolean if a field has been set.

### SetServerRoleSupportedOnBreakoutNil

`func (o *CapabilitySwitchCapability) SetServerRoleSupportedOnBreakoutNil(b bool)`

 SetServerRoleSupportedOnBreakoutNil sets the value for ServerRoleSupportedOnBreakout to be an explicit nil

### UnsetServerRoleSupportedOnBreakout
`func (o *CapabilitySwitchCapability) UnsetServerRoleSupportedOnBreakout()`

UnsetServerRoleSupportedOnBreakout ensures that no value is present for ServerRoleSupportedOnBreakout, not even an explicit nil
### GetStorageLimits

`func (o *CapabilitySwitchCapability) GetStorageLimits() CapabilitySwitchStorageLimits`

GetStorageLimits returns the StorageLimits field if non-nil, zero value otherwise.

### GetStorageLimitsOk

`func (o *CapabilitySwitchCapability) GetStorageLimitsOk() (*CapabilitySwitchStorageLimits, bool)`

GetStorageLimitsOk returns a tuple with the StorageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimits

`func (o *CapabilitySwitchCapability) SetStorageLimits(v CapabilitySwitchStorageLimits)`

SetStorageLimits sets StorageLimits field to given value.

### HasStorageLimits

`func (o *CapabilitySwitchCapability) HasStorageLimits() bool`

HasStorageLimits returns a boolean if a field has been set.

### SetStorageLimitsNil

`func (o *CapabilitySwitchCapability) SetStorageLimitsNil(b bool)`

 SetStorageLimitsNil sets the value for StorageLimits to be an explicit nil

### UnsetStorageLimits
`func (o *CapabilitySwitchCapability) UnsetStorageLimits()`

UnsetStorageLimits ensures that no value is present for StorageLimits, not even an explicit nil
### GetSwitchingModeCapabilities

`func (o *CapabilitySwitchCapability) GetSwitchingModeCapabilities() []CapabilitySwitchingModeCapability`

GetSwitchingModeCapabilities returns the SwitchingModeCapabilities field if non-nil, zero value otherwise.

### GetSwitchingModeCapabilitiesOk

`func (o *CapabilitySwitchCapability) GetSwitchingModeCapabilitiesOk() (*[]CapabilitySwitchingModeCapability, bool)`

GetSwitchingModeCapabilitiesOk returns a tuple with the SwitchingModeCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchingModeCapabilities

`func (o *CapabilitySwitchCapability) SetSwitchingModeCapabilities(v []CapabilitySwitchingModeCapability)`

SetSwitchingModeCapabilities sets SwitchingModeCapabilities field to given value.

### HasSwitchingModeCapabilities

`func (o *CapabilitySwitchCapability) HasSwitchingModeCapabilities() bool`

HasSwitchingModeCapabilities returns a boolean if a field has been set.

### SetSwitchingModeCapabilitiesNil

`func (o *CapabilitySwitchCapability) SetSwitchingModeCapabilitiesNil(b bool)`

 SetSwitchingModeCapabilitiesNil sets the value for SwitchingModeCapabilities to be an explicit nil

### UnsetSwitchingModeCapabilities
`func (o *CapabilitySwitchCapability) UnsetSwitchingModeCapabilities()`

UnsetSwitchingModeCapabilities ensures that no value is present for SwitchingModeCapabilities, not even an explicit nil
### GetSystemLimits

`func (o *CapabilitySwitchCapability) GetSystemLimits() CapabilitySwitchSystemLimits`

GetSystemLimits returns the SystemLimits field if non-nil, zero value otherwise.

### GetSystemLimitsOk

`func (o *CapabilitySwitchCapability) GetSystemLimitsOk() (*CapabilitySwitchSystemLimits, bool)`

GetSystemLimitsOk returns a tuple with the SystemLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLimits

`func (o *CapabilitySwitchCapability) SetSystemLimits(v CapabilitySwitchSystemLimits)`

SetSystemLimits sets SystemLimits field to given value.

### HasSystemLimits

`func (o *CapabilitySwitchCapability) HasSystemLimits() bool`

HasSystemLimits returns a boolean if a field has been set.

### SetSystemLimitsNil

`func (o *CapabilitySwitchCapability) SetSystemLimitsNil(b bool)`

 SetSystemLimitsNil sets the value for SystemLimits to be an explicit nil

### UnsetSystemLimits
`func (o *CapabilitySwitchCapability) UnsetSystemLimits()`

UnsetSystemLimits ensures that no value is present for SystemLimits, not even an explicit nil
### GetUnSupportedEquipmentModel

`func (o *CapabilitySwitchCapability) GetUnSupportedEquipmentModel() []string`

GetUnSupportedEquipmentModel returns the UnSupportedEquipmentModel field if non-nil, zero value otherwise.

### GetUnSupportedEquipmentModelOk

`func (o *CapabilitySwitchCapability) GetUnSupportedEquipmentModelOk() (*[]string, bool)`

GetUnSupportedEquipmentModelOk returns a tuple with the UnSupportedEquipmentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSupportedEquipmentModel

`func (o *CapabilitySwitchCapability) SetUnSupportedEquipmentModel(v []string)`

SetUnSupportedEquipmentModel sets UnSupportedEquipmentModel field to given value.

### HasUnSupportedEquipmentModel

`func (o *CapabilitySwitchCapability) HasUnSupportedEquipmentModel() bool`

HasUnSupportedEquipmentModel returns a boolean if a field has been set.

### SetUnSupportedEquipmentModelNil

`func (o *CapabilitySwitchCapability) SetUnSupportedEquipmentModelNil(b bool)`

 SetUnSupportedEquipmentModelNil sets the value for UnSupportedEquipmentModel to be an explicit nil

### UnsetUnSupportedEquipmentModel
`func (o *CapabilitySwitchCapability) UnsetUnSupportedEquipmentModel()`

UnsetUnSupportedEquipmentModel ensures that no value is present for UnSupportedEquipmentModel, not even an explicit nil
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

### SetUnifiedPortsNil

`func (o *CapabilitySwitchCapability) SetUnifiedPortsNil(b bool)`

 SetUnifiedPortsNil sets the value for UnifiedPorts to be an explicit nil

### UnsetUnifiedPorts
`func (o *CapabilitySwitchCapability) UnsetUnifiedPorts()`

UnsetUnifiedPorts ensures that no value is present for UnifiedPorts, not even an explicit nil
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

### GetUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapability) GetUplinkAdminPortSpeedNegAuto25GbpsSupported() bool`

GetUplinkAdminPortSpeedNegAuto25GbpsSupported returns the UplinkAdminPortSpeedNegAuto25GbpsSupported field if non-nil, zero value otherwise.

### GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk

`func (o *CapabilitySwitchCapability) GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk() (*bool, bool)`

GetUplinkAdminPortSpeedNegAuto25GbpsSupportedOk returns a tuple with the UplinkAdminPortSpeedNegAuto25GbpsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapability) SetUplinkAdminPortSpeedNegAuto25GbpsSupported(v bool)`

SetUplinkAdminPortSpeedNegAuto25GbpsSupported sets UplinkAdminPortSpeedNegAuto25GbpsSupported field to given value.

### HasUplinkAdminPortSpeedNegAuto25GbpsSupported

`func (o *CapabilitySwitchCapability) HasUplinkAdminPortSpeedNegAuto25GbpsSupported() bool`

HasUplinkAdminPortSpeedNegAuto25GbpsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


