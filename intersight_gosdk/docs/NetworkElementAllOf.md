# NetworkElementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Element"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Element"]
**AdminEvacState** | Pointer to **string** | Administratively configured state of Fabric Evacuation feature, for this switch. | [optional] [readonly] 
**AdminInbandInterfaceState** | Pointer to **string** | The administrative state of the network Element inband management interface. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**AvailableMemory** | Pointer to **string** | Available memory (un-used) on this switch platform. | [optional] [readonly] 
**Chassis** | Pointer to **string** | Chassis IP of the switch. | [optional] 
**ConfModTs** | Pointer to **string** | Configuration modified timestamp of the switch. | [optional] 
**ConfModTsBackup** | Pointer to **string** | Configuration modified backup timestamp of the switch. | [optional] 
**EthernetMode** | Pointer to **string** | The user configured Ethernet operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**EthernetSwitchingMode** | Pointer to **string** | The user configured Ethernet operational mode for this switch (End-Host or Switching). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**FaultSummary** | Pointer to **int64** | The fault summary of the network Element out-of-band management interface. | [optional] 
**FcMode** | Pointer to **string** | The user configured FC operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**FcSwitchingMode** | Pointer to **string** | The user configured FC operational mode for this switch (End-Host or Switching). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**InbandIpAddress** | Pointer to **string** | The IP address of the network Element inband management interface. | [optional] [readonly] 
**InbandIpGateway** | Pointer to **string** | The default gateway of the network Element inband management interface. | [optional] [readonly] 
**InbandIpMask** | Pointer to **string** | The network mask of the network Element inband management interface. | [optional] [readonly] 
**InbandVlan** | Pointer to **int64** | The VLAN ID of the network Element inband management interface. | [optional] [readonly] 
**ManagementMode** | Pointer to **string** | The management mode of the fabric interconnect. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [default to "IntersightStandalone"]
**OperEvacState** | Pointer to **string** | Operational state of the Fabric Evacuation feature, for this switch. | [optional] [readonly] 
**Operability** | Pointer to **string** | The switch&#39;s current overall operational/health state. | [optional] [readonly] 
**OutOfBandIpAddress** | Pointer to **string** | The IP address of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpGateway** | Pointer to **string** | The default gateway of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpMask** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Address** | Pointer to **string** | The IPv4 address of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Gateway** | Pointer to **string** | The default IPv4 gateway of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Mask** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv6Address** | Pointer to **string** | The IPv6 address of the network Element out-of-band management interface. | [optional] 
**OutOfBandIpv6Gateway** | Pointer to **string** | The default IPv6 gateway of the network Element out-of-band management interface. | [optional] 
**OutOfBandIpv6Prefix** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] 
**OutOfBandMac** | Pointer to **string** | The MAC address of the network Element out-of-band management interface. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number of the switch. | [optional] 
**Status** | Pointer to **string** | The status of the switch. | [optional] 
**SwitchId** | Pointer to **string** | The Switch Id of the network Element. | [optional] [readonly] 
**SwitchType** | Pointer to **string** | The Switch type that the network element is a part of. * &#x60;FabricInterconnect&#x60; - The default Switch type of UCSM and IMM mode devices. * &#x60;NexusDevice&#x60; - Switch type of Nexus devices. * &#x60;MDSDevice&#x60; - Switch type of Nexus MDS devices. | [optional] [default to "FabricInterconnect"]
**SystemUpTime** | Pointer to **string** | System up time of the switch. | [optional] 
**Thermal** | Pointer to **string** | The Thermal status of the fabric interconnect. * &#x60;unknown&#x60; - The default state of the sensor (in case no data is received). * &#x60;ok&#x60; - State of the sensor indicating the sensor&#39;s temperature range is okay. * &#x60;upper-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely high above normal range. * &#x60;upper-critical&#x60; - State of the sensor indicating that the temperature is above normal range. * &#x60;upper-non-critical&#x60; - State of the sensor indicating that the temperature is a little above the normal range. * &#x60;lower-non-critical&#x60; - State of the sensor indicating that the temperature is a little below the normal range. * &#x60;lower-critical&#x60; - State of the sensor indicating that the temperature is below normal range. * &#x60;lower-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely below normal range. | [optional] [default to "unknown"]
**TotalMemory** | Pointer to **int64** | Total available memory on this switch platform. | [optional] [readonly] 
**Version** | Pointer to **string** | Firmware version of the switch. | [optional] 
**Cards** | Pointer to [**[]EquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) | An array of relationships to equipmentSwitchCard resources. | [optional] [readonly] 
**Console** | Pointer to [**[]ConsoleConsoleConfigRelationship**](ConsoleConsoleConfigRelationship.md) | An array of relationships to consoleConsoleConfig resources. | [optional] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](EquipmentFanModuleRelationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**FeatureControl** | Pointer to [**[]NetworkFeatureControlRelationship**](NetworkFeatureControlRelationship.md) | An array of relationships to networkFeatureControl resources. | [optional] 
**InterfaceList** | Pointer to [**[]NetworkInterfaceListRelationship**](NetworkInterfaceListRelationship.md) | An array of relationships to networkInterfaceList resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LicenseFile** | Pointer to [**[]NetworkLicenseFileRelationship**](NetworkLicenseFileRelationship.md) | An array of relationships to networkLicenseFile resources. | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**ManagementEntity** | Pointer to [**ManagementEntityRelationship**](ManagementEntityRelationship.md) |  | [optional] 
**NetworkFcZoneInfo** | Pointer to [**NetworkFcZoneInfoRelationship**](NetworkFcZoneInfoRelationship.md) |  | [optional] 
**NetworkVlanPortInfo** | Pointer to [**NetworkVlanPortInfoRelationship**](NetworkVlanPortInfoRelationship.md) |  | [optional] 
**PortMacBindings** | Pointer to [**[]PortMacBindingRelationship**](PortMacBindingRelationship.md) | An array of relationships to portMacBinding resources. | [optional] 
**ProcessorUnit** | Pointer to [**[]ProcessorUnitRelationship**](ProcessorUnitRelationship.md) | An array of relationships to processorUnit resources. | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](EquipmentPsuRelationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageItems** | Pointer to [**[]StorageItemRelationship**](StorageItemRelationship.md) | An array of relationships to storageItem resources. | [optional] [readonly] 
**SupervisorCard** | Pointer to [**[]NetworkSupervisorCardRelationship**](NetworkSupervisorCardRelationship.md) | An array of relationships to networkSupervisorCard resources. | [optional] 
**TopSystem** | Pointer to [**TopSystemRelationship**](TopSystemRelationship.md) |  | [optional] 
**UcsmRunningFirmware** | Pointer to [**FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) |  | [optional] 
**Vrf** | Pointer to [**[]NetworkVrfRelationship**](NetworkVrfRelationship.md) | An array of relationships to networkVrf resources. | [optional] 

## Methods

### NewNetworkElementAllOf

`func NewNetworkElementAllOf(classId string, objectType string, ) *NetworkElementAllOf`

NewNetworkElementAllOf instantiates a new NetworkElementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkElementAllOfWithDefaults

`func NewNetworkElementAllOfWithDefaults() *NetworkElementAllOf`

NewNetworkElementAllOfWithDefaults instantiates a new NetworkElementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkElementAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkElementAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkElementAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkElementAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkElementAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkElementAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminEvacState

`func (o *NetworkElementAllOf) GetAdminEvacState() string`

GetAdminEvacState returns the AdminEvacState field if non-nil, zero value otherwise.

### GetAdminEvacStateOk

`func (o *NetworkElementAllOf) GetAdminEvacStateOk() (*string, bool)`

GetAdminEvacStateOk returns a tuple with the AdminEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEvacState

`func (o *NetworkElementAllOf) SetAdminEvacState(v string)`

SetAdminEvacState sets AdminEvacState field to given value.

### HasAdminEvacState

`func (o *NetworkElementAllOf) HasAdminEvacState() bool`

HasAdminEvacState returns a boolean if a field has been set.

### GetAdminInbandInterfaceState

`func (o *NetworkElementAllOf) GetAdminInbandInterfaceState() string`

GetAdminInbandInterfaceState returns the AdminInbandInterfaceState field if non-nil, zero value otherwise.

### GetAdminInbandInterfaceStateOk

`func (o *NetworkElementAllOf) GetAdminInbandInterfaceStateOk() (*string, bool)`

GetAdminInbandInterfaceStateOk returns a tuple with the AdminInbandInterfaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminInbandInterfaceState

`func (o *NetworkElementAllOf) SetAdminInbandInterfaceState(v string)`

SetAdminInbandInterfaceState sets AdminInbandInterfaceState field to given value.

### HasAdminInbandInterfaceState

`func (o *NetworkElementAllOf) HasAdminInbandInterfaceState() bool`

HasAdminInbandInterfaceState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *NetworkElementAllOf) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *NetworkElementAllOf) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *NetworkElementAllOf) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *NetworkElementAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *NetworkElementAllOf) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *NetworkElementAllOf) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetAvailableMemory

`func (o *NetworkElementAllOf) GetAvailableMemory() string`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *NetworkElementAllOf) GetAvailableMemoryOk() (*string, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *NetworkElementAllOf) SetAvailableMemory(v string)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *NetworkElementAllOf) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetChassis

`func (o *NetworkElementAllOf) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *NetworkElementAllOf) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *NetworkElementAllOf) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *NetworkElementAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetConfModTs

`func (o *NetworkElementAllOf) GetConfModTs() string`

GetConfModTs returns the ConfModTs field if non-nil, zero value otherwise.

### GetConfModTsOk

`func (o *NetworkElementAllOf) GetConfModTsOk() (*string, bool)`

GetConfModTsOk returns a tuple with the ConfModTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfModTs

`func (o *NetworkElementAllOf) SetConfModTs(v string)`

SetConfModTs sets ConfModTs field to given value.

### HasConfModTs

`func (o *NetworkElementAllOf) HasConfModTs() bool`

HasConfModTs returns a boolean if a field has been set.

### GetConfModTsBackup

`func (o *NetworkElementAllOf) GetConfModTsBackup() string`

GetConfModTsBackup returns the ConfModTsBackup field if non-nil, zero value otherwise.

### GetConfModTsBackupOk

`func (o *NetworkElementAllOf) GetConfModTsBackupOk() (*string, bool)`

GetConfModTsBackupOk returns a tuple with the ConfModTsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfModTsBackup

`func (o *NetworkElementAllOf) SetConfModTsBackup(v string)`

SetConfModTsBackup sets ConfModTsBackup field to given value.

### HasConfModTsBackup

`func (o *NetworkElementAllOf) HasConfModTsBackup() bool`

HasConfModTsBackup returns a boolean if a field has been set.

### GetEthernetMode

`func (o *NetworkElementAllOf) GetEthernetMode() string`

GetEthernetMode returns the EthernetMode field if non-nil, zero value otherwise.

### GetEthernetModeOk

`func (o *NetworkElementAllOf) GetEthernetModeOk() (*string, bool)`

GetEthernetModeOk returns a tuple with the EthernetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMode

`func (o *NetworkElementAllOf) SetEthernetMode(v string)`

SetEthernetMode sets EthernetMode field to given value.

### HasEthernetMode

`func (o *NetworkElementAllOf) HasEthernetMode() bool`

HasEthernetMode returns a boolean if a field has been set.

### GetEthernetSwitchingMode

`func (o *NetworkElementAllOf) GetEthernetSwitchingMode() string`

GetEthernetSwitchingMode returns the EthernetSwitchingMode field if non-nil, zero value otherwise.

### GetEthernetSwitchingModeOk

`func (o *NetworkElementAllOf) GetEthernetSwitchingModeOk() (*string, bool)`

GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetSwitchingMode

`func (o *NetworkElementAllOf) SetEthernetSwitchingMode(v string)`

SetEthernetSwitchingMode sets EthernetSwitchingMode field to given value.

### HasEthernetSwitchingMode

`func (o *NetworkElementAllOf) HasEthernetSwitchingMode() bool`

HasEthernetSwitchingMode returns a boolean if a field has been set.

### GetFaultSummary

`func (o *NetworkElementAllOf) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *NetworkElementAllOf) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *NetworkElementAllOf) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *NetworkElementAllOf) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFcMode

`func (o *NetworkElementAllOf) GetFcMode() string`

GetFcMode returns the FcMode field if non-nil, zero value otherwise.

### GetFcModeOk

`func (o *NetworkElementAllOf) GetFcModeOk() (*string, bool)`

GetFcModeOk returns a tuple with the FcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcMode

`func (o *NetworkElementAllOf) SetFcMode(v string)`

SetFcMode sets FcMode field to given value.

### HasFcMode

`func (o *NetworkElementAllOf) HasFcMode() bool`

HasFcMode returns a boolean if a field has been set.

### GetFcSwitchingMode

`func (o *NetworkElementAllOf) GetFcSwitchingMode() string`

GetFcSwitchingMode returns the FcSwitchingMode field if non-nil, zero value otherwise.

### GetFcSwitchingModeOk

`func (o *NetworkElementAllOf) GetFcSwitchingModeOk() (*string, bool)`

GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSwitchingMode

`func (o *NetworkElementAllOf) SetFcSwitchingMode(v string)`

SetFcSwitchingMode sets FcSwitchingMode field to given value.

### HasFcSwitchingMode

`func (o *NetworkElementAllOf) HasFcSwitchingMode() bool`

HasFcSwitchingMode returns a boolean if a field has been set.

### GetInbandIpAddress

`func (o *NetworkElementAllOf) GetInbandIpAddress() string`

GetInbandIpAddress returns the InbandIpAddress field if non-nil, zero value otherwise.

### GetInbandIpAddressOk

`func (o *NetworkElementAllOf) GetInbandIpAddressOk() (*string, bool)`

GetInbandIpAddressOk returns a tuple with the InbandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpAddress

`func (o *NetworkElementAllOf) SetInbandIpAddress(v string)`

SetInbandIpAddress sets InbandIpAddress field to given value.

### HasInbandIpAddress

`func (o *NetworkElementAllOf) HasInbandIpAddress() bool`

HasInbandIpAddress returns a boolean if a field has been set.

### GetInbandIpGateway

`func (o *NetworkElementAllOf) GetInbandIpGateway() string`

GetInbandIpGateway returns the InbandIpGateway field if non-nil, zero value otherwise.

### GetInbandIpGatewayOk

`func (o *NetworkElementAllOf) GetInbandIpGatewayOk() (*string, bool)`

GetInbandIpGatewayOk returns a tuple with the InbandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpGateway

`func (o *NetworkElementAllOf) SetInbandIpGateway(v string)`

SetInbandIpGateway sets InbandIpGateway field to given value.

### HasInbandIpGateway

`func (o *NetworkElementAllOf) HasInbandIpGateway() bool`

HasInbandIpGateway returns a boolean if a field has been set.

### GetInbandIpMask

`func (o *NetworkElementAllOf) GetInbandIpMask() string`

GetInbandIpMask returns the InbandIpMask field if non-nil, zero value otherwise.

### GetInbandIpMaskOk

`func (o *NetworkElementAllOf) GetInbandIpMaskOk() (*string, bool)`

GetInbandIpMaskOk returns a tuple with the InbandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpMask

`func (o *NetworkElementAllOf) SetInbandIpMask(v string)`

SetInbandIpMask sets InbandIpMask field to given value.

### HasInbandIpMask

`func (o *NetworkElementAllOf) HasInbandIpMask() bool`

HasInbandIpMask returns a boolean if a field has been set.

### GetInbandVlan

`func (o *NetworkElementAllOf) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *NetworkElementAllOf) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *NetworkElementAllOf) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *NetworkElementAllOf) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetManagementMode

`func (o *NetworkElementAllOf) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *NetworkElementAllOf) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *NetworkElementAllOf) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *NetworkElementAllOf) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetOperEvacState

`func (o *NetworkElementAllOf) GetOperEvacState() string`

GetOperEvacState returns the OperEvacState field if non-nil, zero value otherwise.

### GetOperEvacStateOk

`func (o *NetworkElementAllOf) GetOperEvacStateOk() (*string, bool)`

GetOperEvacStateOk returns a tuple with the OperEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperEvacState

`func (o *NetworkElementAllOf) SetOperEvacState(v string)`

SetOperEvacState sets OperEvacState field to given value.

### HasOperEvacState

`func (o *NetworkElementAllOf) HasOperEvacState() bool`

HasOperEvacState returns a boolean if a field has been set.

### GetOperability

`func (o *NetworkElementAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *NetworkElementAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *NetworkElementAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *NetworkElementAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *NetworkElementAllOf) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *NetworkElementAllOf) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *NetworkElementAllOf) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *NetworkElementAllOf) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *NetworkElementAllOf) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *NetworkElementAllOf) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *NetworkElementAllOf) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *NetworkElementAllOf) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetOutOfBandIpMask

`func (o *NetworkElementAllOf) GetOutOfBandIpMask() string`

GetOutOfBandIpMask returns the OutOfBandIpMask field if non-nil, zero value otherwise.

### GetOutOfBandIpMaskOk

`func (o *NetworkElementAllOf) GetOutOfBandIpMaskOk() (*string, bool)`

GetOutOfBandIpMaskOk returns a tuple with the OutOfBandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpMask

`func (o *NetworkElementAllOf) SetOutOfBandIpMask(v string)`

SetOutOfBandIpMask sets OutOfBandIpMask field to given value.

### HasOutOfBandIpMask

`func (o *NetworkElementAllOf) HasOutOfBandIpMask() bool`

HasOutOfBandIpMask returns a boolean if a field has been set.

### GetOutOfBandIpv4Address

`func (o *NetworkElementAllOf) GetOutOfBandIpv4Address() string`

GetOutOfBandIpv4Address returns the OutOfBandIpv4Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv4AddressOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv4AddressOk() (*string, bool)`

GetOutOfBandIpv4AddressOk returns a tuple with the OutOfBandIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Address

`func (o *NetworkElementAllOf) SetOutOfBandIpv4Address(v string)`

SetOutOfBandIpv4Address sets OutOfBandIpv4Address field to given value.

### HasOutOfBandIpv4Address

`func (o *NetworkElementAllOf) HasOutOfBandIpv4Address() bool`

HasOutOfBandIpv4Address returns a boolean if a field has been set.

### GetOutOfBandIpv4Gateway

`func (o *NetworkElementAllOf) GetOutOfBandIpv4Gateway() string`

GetOutOfBandIpv4Gateway returns the OutOfBandIpv4Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv4GatewayOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv4GatewayOk() (*string, bool)`

GetOutOfBandIpv4GatewayOk returns a tuple with the OutOfBandIpv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Gateway

`func (o *NetworkElementAllOf) SetOutOfBandIpv4Gateway(v string)`

SetOutOfBandIpv4Gateway sets OutOfBandIpv4Gateway field to given value.

### HasOutOfBandIpv4Gateway

`func (o *NetworkElementAllOf) HasOutOfBandIpv4Gateway() bool`

HasOutOfBandIpv4Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv4Mask

`func (o *NetworkElementAllOf) GetOutOfBandIpv4Mask() string`

GetOutOfBandIpv4Mask returns the OutOfBandIpv4Mask field if non-nil, zero value otherwise.

### GetOutOfBandIpv4MaskOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv4MaskOk() (*string, bool)`

GetOutOfBandIpv4MaskOk returns a tuple with the OutOfBandIpv4Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Mask

`func (o *NetworkElementAllOf) SetOutOfBandIpv4Mask(v string)`

SetOutOfBandIpv4Mask sets OutOfBandIpv4Mask field to given value.

### HasOutOfBandIpv4Mask

`func (o *NetworkElementAllOf) HasOutOfBandIpv4Mask() bool`

HasOutOfBandIpv4Mask returns a boolean if a field has been set.

### GetOutOfBandIpv6Address

`func (o *NetworkElementAllOf) GetOutOfBandIpv6Address() string`

GetOutOfBandIpv6Address returns the OutOfBandIpv6Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv6AddressOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv6AddressOk() (*string, bool)`

GetOutOfBandIpv6AddressOk returns a tuple with the OutOfBandIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Address

`func (o *NetworkElementAllOf) SetOutOfBandIpv6Address(v string)`

SetOutOfBandIpv6Address sets OutOfBandIpv6Address field to given value.

### HasOutOfBandIpv6Address

`func (o *NetworkElementAllOf) HasOutOfBandIpv6Address() bool`

HasOutOfBandIpv6Address returns a boolean if a field has been set.

### GetOutOfBandIpv6Gateway

`func (o *NetworkElementAllOf) GetOutOfBandIpv6Gateway() string`

GetOutOfBandIpv6Gateway returns the OutOfBandIpv6Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv6GatewayOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv6GatewayOk() (*string, bool)`

GetOutOfBandIpv6GatewayOk returns a tuple with the OutOfBandIpv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Gateway

`func (o *NetworkElementAllOf) SetOutOfBandIpv6Gateway(v string)`

SetOutOfBandIpv6Gateway sets OutOfBandIpv6Gateway field to given value.

### HasOutOfBandIpv6Gateway

`func (o *NetworkElementAllOf) HasOutOfBandIpv6Gateway() bool`

HasOutOfBandIpv6Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv6Prefix

`func (o *NetworkElementAllOf) GetOutOfBandIpv6Prefix() string`

GetOutOfBandIpv6Prefix returns the OutOfBandIpv6Prefix field if non-nil, zero value otherwise.

### GetOutOfBandIpv6PrefixOk

`func (o *NetworkElementAllOf) GetOutOfBandIpv6PrefixOk() (*string, bool)`

GetOutOfBandIpv6PrefixOk returns a tuple with the OutOfBandIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Prefix

`func (o *NetworkElementAllOf) SetOutOfBandIpv6Prefix(v string)`

SetOutOfBandIpv6Prefix sets OutOfBandIpv6Prefix field to given value.

### HasOutOfBandIpv6Prefix

`func (o *NetworkElementAllOf) HasOutOfBandIpv6Prefix() bool`

HasOutOfBandIpv6Prefix returns a boolean if a field has been set.

### GetOutOfBandMac

`func (o *NetworkElementAllOf) GetOutOfBandMac() string`

GetOutOfBandMac returns the OutOfBandMac field if non-nil, zero value otherwise.

### GetOutOfBandMacOk

`func (o *NetworkElementAllOf) GetOutOfBandMacOk() (*string, bool)`

GetOutOfBandMacOk returns a tuple with the OutOfBandMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandMac

`func (o *NetworkElementAllOf) SetOutOfBandMac(v string)`

SetOutOfBandMac sets OutOfBandMac field to given value.

### HasOutOfBandMac

`func (o *NetworkElementAllOf) HasOutOfBandMac() bool`

HasOutOfBandMac returns a boolean if a field has been set.

### GetPartNumber

`func (o *NetworkElementAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *NetworkElementAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *NetworkElementAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *NetworkElementAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkElementAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkElementAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkElementAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkElementAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *NetworkElementAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkElementAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkElementAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *NetworkElementAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchType

`func (o *NetworkElementAllOf) GetSwitchType() string`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *NetworkElementAllOf) GetSwitchTypeOk() (*string, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *NetworkElementAllOf) SetSwitchType(v string)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *NetworkElementAllOf) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.

### GetSystemUpTime

`func (o *NetworkElementAllOf) GetSystemUpTime() string`

GetSystemUpTime returns the SystemUpTime field if non-nil, zero value otherwise.

### GetSystemUpTimeOk

`func (o *NetworkElementAllOf) GetSystemUpTimeOk() (*string, bool)`

GetSystemUpTimeOk returns a tuple with the SystemUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpTime

`func (o *NetworkElementAllOf) SetSystemUpTime(v string)`

SetSystemUpTime sets SystemUpTime field to given value.

### HasSystemUpTime

`func (o *NetworkElementAllOf) HasSystemUpTime() bool`

HasSystemUpTime returns a boolean if a field has been set.

### GetThermal

`func (o *NetworkElementAllOf) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *NetworkElementAllOf) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *NetworkElementAllOf) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *NetworkElementAllOf) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetTotalMemory

`func (o *NetworkElementAllOf) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *NetworkElementAllOf) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *NetworkElementAllOf) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *NetworkElementAllOf) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkElementAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkElementAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkElementAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkElementAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCards

`func (o *NetworkElementAllOf) GetCards() []EquipmentSwitchCardRelationship`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *NetworkElementAllOf) GetCardsOk() (*[]EquipmentSwitchCardRelationship, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *NetworkElementAllOf) SetCards(v []EquipmentSwitchCardRelationship)`

SetCards sets Cards field to given value.

### HasCards

`func (o *NetworkElementAllOf) HasCards() bool`

HasCards returns a boolean if a field has been set.

### SetCardsNil

`func (o *NetworkElementAllOf) SetCardsNil(b bool)`

 SetCardsNil sets the value for Cards to be an explicit nil

### UnsetCards
`func (o *NetworkElementAllOf) UnsetCards()`

UnsetCards ensures that no value is present for Cards, not even an explicit nil
### GetConsole

`func (o *NetworkElementAllOf) GetConsole() []ConsoleConsoleConfigRelationship`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *NetworkElementAllOf) GetConsoleOk() (*[]ConsoleConsoleConfigRelationship, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *NetworkElementAllOf) SetConsole(v []ConsoleConsoleConfigRelationship)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *NetworkElementAllOf) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### SetConsoleNil

`func (o *NetworkElementAllOf) SetConsoleNil(b bool)`

 SetConsoleNil sets the value for Console to be an explicit nil

### UnsetConsole
`func (o *NetworkElementAllOf) UnsetConsole()`

UnsetConsole ensures that no value is present for Console, not even an explicit nil
### GetFanmodules

`func (o *NetworkElementAllOf) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *NetworkElementAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *NetworkElementAllOf) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *NetworkElementAllOf) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *NetworkElementAllOf) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *NetworkElementAllOf) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetFeatureControl

`func (o *NetworkElementAllOf) GetFeatureControl() []NetworkFeatureControlRelationship`

GetFeatureControl returns the FeatureControl field if non-nil, zero value otherwise.

### GetFeatureControlOk

`func (o *NetworkElementAllOf) GetFeatureControlOk() (*[]NetworkFeatureControlRelationship, bool)`

GetFeatureControlOk returns a tuple with the FeatureControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureControl

`func (o *NetworkElementAllOf) SetFeatureControl(v []NetworkFeatureControlRelationship)`

SetFeatureControl sets FeatureControl field to given value.

### HasFeatureControl

`func (o *NetworkElementAllOf) HasFeatureControl() bool`

HasFeatureControl returns a boolean if a field has been set.

### SetFeatureControlNil

`func (o *NetworkElementAllOf) SetFeatureControlNil(b bool)`

 SetFeatureControlNil sets the value for FeatureControl to be an explicit nil

### UnsetFeatureControl
`func (o *NetworkElementAllOf) UnsetFeatureControl()`

UnsetFeatureControl ensures that no value is present for FeatureControl, not even an explicit nil
### GetInterfaceList

`func (o *NetworkElementAllOf) GetInterfaceList() []NetworkInterfaceListRelationship`

GetInterfaceList returns the InterfaceList field if non-nil, zero value otherwise.

### GetInterfaceListOk

`func (o *NetworkElementAllOf) GetInterfaceListOk() (*[]NetworkInterfaceListRelationship, bool)`

GetInterfaceListOk returns a tuple with the InterfaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceList

`func (o *NetworkElementAllOf) SetInterfaceList(v []NetworkInterfaceListRelationship)`

SetInterfaceList sets InterfaceList field to given value.

### HasInterfaceList

`func (o *NetworkElementAllOf) HasInterfaceList() bool`

HasInterfaceList returns a boolean if a field has been set.

### SetInterfaceListNil

`func (o *NetworkElementAllOf) SetInterfaceListNil(b bool)`

 SetInterfaceListNil sets the value for InterfaceList to be an explicit nil

### UnsetInterfaceList
`func (o *NetworkElementAllOf) UnsetInterfaceList()`

UnsetInterfaceList ensures that no value is present for InterfaceList, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *NetworkElementAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkElementAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkElementAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkElementAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLicenseFile

`func (o *NetworkElementAllOf) GetLicenseFile() []NetworkLicenseFileRelationship`

GetLicenseFile returns the LicenseFile field if non-nil, zero value otherwise.

### GetLicenseFileOk

`func (o *NetworkElementAllOf) GetLicenseFileOk() (*[]NetworkLicenseFileRelationship, bool)`

GetLicenseFileOk returns a tuple with the LicenseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFile

`func (o *NetworkElementAllOf) SetLicenseFile(v []NetworkLicenseFileRelationship)`

SetLicenseFile sets LicenseFile field to given value.

### HasLicenseFile

`func (o *NetworkElementAllOf) HasLicenseFile() bool`

HasLicenseFile returns a boolean if a field has been set.

### SetLicenseFileNil

`func (o *NetworkElementAllOf) SetLicenseFileNil(b bool)`

 SetLicenseFileNil sets the value for LicenseFile to be an explicit nil

### UnsetLicenseFile
`func (o *NetworkElementAllOf) UnsetLicenseFile()`

UnsetLicenseFile ensures that no value is present for LicenseFile, not even an explicit nil
### GetManagementController

`func (o *NetworkElementAllOf) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *NetworkElementAllOf) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *NetworkElementAllOf) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *NetworkElementAllOf) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetManagementEntity

`func (o *NetworkElementAllOf) GetManagementEntity() ManagementEntityRelationship`

GetManagementEntity returns the ManagementEntity field if non-nil, zero value otherwise.

### GetManagementEntityOk

`func (o *NetworkElementAllOf) GetManagementEntityOk() (*ManagementEntityRelationship, bool)`

GetManagementEntityOk returns a tuple with the ManagementEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEntity

`func (o *NetworkElementAllOf) SetManagementEntity(v ManagementEntityRelationship)`

SetManagementEntity sets ManagementEntity field to given value.

### HasManagementEntity

`func (o *NetworkElementAllOf) HasManagementEntity() bool`

HasManagementEntity returns a boolean if a field has been set.

### GetNetworkFcZoneInfo

`func (o *NetworkElementAllOf) GetNetworkFcZoneInfo() NetworkFcZoneInfoRelationship`

GetNetworkFcZoneInfo returns the NetworkFcZoneInfo field if non-nil, zero value otherwise.

### GetNetworkFcZoneInfoOk

`func (o *NetworkElementAllOf) GetNetworkFcZoneInfoOk() (*NetworkFcZoneInfoRelationship, bool)`

GetNetworkFcZoneInfoOk returns a tuple with the NetworkFcZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFcZoneInfo

`func (o *NetworkElementAllOf) SetNetworkFcZoneInfo(v NetworkFcZoneInfoRelationship)`

SetNetworkFcZoneInfo sets NetworkFcZoneInfo field to given value.

### HasNetworkFcZoneInfo

`func (o *NetworkElementAllOf) HasNetworkFcZoneInfo() bool`

HasNetworkFcZoneInfo returns a boolean if a field has been set.

### GetNetworkVlanPortInfo

`func (o *NetworkElementAllOf) GetNetworkVlanPortInfo() NetworkVlanPortInfoRelationship`

GetNetworkVlanPortInfo returns the NetworkVlanPortInfo field if non-nil, zero value otherwise.

### GetNetworkVlanPortInfoOk

`func (o *NetworkElementAllOf) GetNetworkVlanPortInfoOk() (*NetworkVlanPortInfoRelationship, bool)`

GetNetworkVlanPortInfoOk returns a tuple with the NetworkVlanPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanPortInfo

`func (o *NetworkElementAllOf) SetNetworkVlanPortInfo(v NetworkVlanPortInfoRelationship)`

SetNetworkVlanPortInfo sets NetworkVlanPortInfo field to given value.

### HasNetworkVlanPortInfo

`func (o *NetworkElementAllOf) HasNetworkVlanPortInfo() bool`

HasNetworkVlanPortInfo returns a boolean if a field has been set.

### GetPortMacBindings

`func (o *NetworkElementAllOf) GetPortMacBindings() []PortMacBindingRelationship`

GetPortMacBindings returns the PortMacBindings field if non-nil, zero value otherwise.

### GetPortMacBindingsOk

`func (o *NetworkElementAllOf) GetPortMacBindingsOk() (*[]PortMacBindingRelationship, bool)`

GetPortMacBindingsOk returns a tuple with the PortMacBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMacBindings

`func (o *NetworkElementAllOf) SetPortMacBindings(v []PortMacBindingRelationship)`

SetPortMacBindings sets PortMacBindings field to given value.

### HasPortMacBindings

`func (o *NetworkElementAllOf) HasPortMacBindings() bool`

HasPortMacBindings returns a boolean if a field has been set.

### SetPortMacBindingsNil

`func (o *NetworkElementAllOf) SetPortMacBindingsNil(b bool)`

 SetPortMacBindingsNil sets the value for PortMacBindings to be an explicit nil

### UnsetPortMacBindings
`func (o *NetworkElementAllOf) UnsetPortMacBindings()`

UnsetPortMacBindings ensures that no value is present for PortMacBindings, not even an explicit nil
### GetProcessorUnit

`func (o *NetworkElementAllOf) GetProcessorUnit() []ProcessorUnitRelationship`

GetProcessorUnit returns the ProcessorUnit field if non-nil, zero value otherwise.

### GetProcessorUnitOk

`func (o *NetworkElementAllOf) GetProcessorUnitOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorUnitOk returns a tuple with the ProcessorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorUnit

`func (o *NetworkElementAllOf) SetProcessorUnit(v []ProcessorUnitRelationship)`

SetProcessorUnit sets ProcessorUnit field to given value.

### HasProcessorUnit

`func (o *NetworkElementAllOf) HasProcessorUnit() bool`

HasProcessorUnit returns a boolean if a field has been set.

### SetProcessorUnitNil

`func (o *NetworkElementAllOf) SetProcessorUnitNil(b bool)`

 SetProcessorUnitNil sets the value for ProcessorUnit to be an explicit nil

### UnsetProcessorUnit
`func (o *NetworkElementAllOf) UnsetProcessorUnit()`

UnsetProcessorUnit ensures that no value is present for ProcessorUnit, not even an explicit nil
### GetPsus

`func (o *NetworkElementAllOf) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *NetworkElementAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *NetworkElementAllOf) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *NetworkElementAllOf) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *NetworkElementAllOf) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *NetworkElementAllOf) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkElementAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkElementAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkElementAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkElementAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageItems

`func (o *NetworkElementAllOf) GetStorageItems() []StorageItemRelationship`

GetStorageItems returns the StorageItems field if non-nil, zero value otherwise.

### GetStorageItemsOk

`func (o *NetworkElementAllOf) GetStorageItemsOk() (*[]StorageItemRelationship, bool)`

GetStorageItemsOk returns a tuple with the StorageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageItems

`func (o *NetworkElementAllOf) SetStorageItems(v []StorageItemRelationship)`

SetStorageItems sets StorageItems field to given value.

### HasStorageItems

`func (o *NetworkElementAllOf) HasStorageItems() bool`

HasStorageItems returns a boolean if a field has been set.

### SetStorageItemsNil

`func (o *NetworkElementAllOf) SetStorageItemsNil(b bool)`

 SetStorageItemsNil sets the value for StorageItems to be an explicit nil

### UnsetStorageItems
`func (o *NetworkElementAllOf) UnsetStorageItems()`

UnsetStorageItems ensures that no value is present for StorageItems, not even an explicit nil
### GetSupervisorCard

`func (o *NetworkElementAllOf) GetSupervisorCard() []NetworkSupervisorCardRelationship`

GetSupervisorCard returns the SupervisorCard field if non-nil, zero value otherwise.

### GetSupervisorCardOk

`func (o *NetworkElementAllOf) GetSupervisorCardOk() (*[]NetworkSupervisorCardRelationship, bool)`

GetSupervisorCardOk returns a tuple with the SupervisorCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorCard

`func (o *NetworkElementAllOf) SetSupervisorCard(v []NetworkSupervisorCardRelationship)`

SetSupervisorCard sets SupervisorCard field to given value.

### HasSupervisorCard

`func (o *NetworkElementAllOf) HasSupervisorCard() bool`

HasSupervisorCard returns a boolean if a field has been set.

### SetSupervisorCardNil

`func (o *NetworkElementAllOf) SetSupervisorCardNil(b bool)`

 SetSupervisorCardNil sets the value for SupervisorCard to be an explicit nil

### UnsetSupervisorCard
`func (o *NetworkElementAllOf) UnsetSupervisorCard()`

UnsetSupervisorCard ensures that no value is present for SupervisorCard, not even an explicit nil
### GetTopSystem

`func (o *NetworkElementAllOf) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *NetworkElementAllOf) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *NetworkElementAllOf) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *NetworkElementAllOf) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.

### GetUcsmRunningFirmware

`func (o *NetworkElementAllOf) GetUcsmRunningFirmware() FirmwareRunningFirmwareRelationship`

GetUcsmRunningFirmware returns the UcsmRunningFirmware field if non-nil, zero value otherwise.

### GetUcsmRunningFirmwareOk

`func (o *NetworkElementAllOf) GetUcsmRunningFirmwareOk() (*FirmwareRunningFirmwareRelationship, bool)`

GetUcsmRunningFirmwareOk returns a tuple with the UcsmRunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsmRunningFirmware

`func (o *NetworkElementAllOf) SetUcsmRunningFirmware(v FirmwareRunningFirmwareRelationship)`

SetUcsmRunningFirmware sets UcsmRunningFirmware field to given value.

### HasUcsmRunningFirmware

`func (o *NetworkElementAllOf) HasUcsmRunningFirmware() bool`

HasUcsmRunningFirmware returns a boolean if a field has been set.

### GetVrf

`func (o *NetworkElementAllOf) GetVrf() []NetworkVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *NetworkElementAllOf) GetVrfOk() (*[]NetworkVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *NetworkElementAllOf) SetVrf(v []NetworkVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *NetworkElementAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *NetworkElementAllOf) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *NetworkElementAllOf) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


