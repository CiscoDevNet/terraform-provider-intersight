# NetworkElementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.ElementSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.ElementSummary"]
**AdminEvacState** | Pointer to **string** | Administratively configured state of Fabric Evacuation feature, for this switch. * &#x60;&#x60; - Evacuation state of the switch is unknown. * &#x60;enabled&#x60; - Evacuation state of the switch is enabled. * &#x60;disabled&#x60; - Evacuation state of the switch is disabled. * &#x60;applying&#x60; - Evacuation state of the switch when evacuation is in progress. * &#x60;on&#x60; - Evacuation state of the switch is enabled. * &#x60;off&#x60; - Evacuation state of the switch is disabled. * &#x60;N/A&#x60; - Evacuation state of the switch is not applicable. | [optional] [readonly] [default to ""]
**AdminInbandInterfaceState** | Pointer to **string** | The administrative state of the network Element inband management interface. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**AvailableMemory** | Pointer to **string** | Available memory (un-used) on this switch platform. | [optional] [readonly] 
**BundleVersion** | Pointer to **string** | Running firmware bundle information. | [optional] [readonly] 
**Chassis** | Pointer to **string** | Chassis IP of the switch. | [optional] [readonly] 
**ConfModTs** | Pointer to **string** | Configuration modified timestamp of the switch. | [optional] [readonly] 
**ConfModTsBackup** | Pointer to **string** | Configuration modified backup timestamp of the switch. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connection status of the switch. | [optional] [readonly] 
**DefaultDomain** | Pointer to **string** | The default domain name configured on the switch. | [optional] [readonly] 
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**EthernetMode** | Pointer to **string** | The user configured Ethernet operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**EthernetSwitchingMode** | Pointer to **string** | The user configured Ethernet operational mode for this switch (End-Host or Switching). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**FaultSummary** | Pointer to **int64** | The fault summary of the network Element out-of-band management interface. | [optional] [readonly] 
**FcMode** | Pointer to **string** | The user configured FC operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**FcSwitchingMode** | Pointer to **string** | The user configured FC operational mode for this switch (End-Host or Switching). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**Firmware** | Pointer to **string** | Running firmware information. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Running firmware information. | [optional] [readonly] 
**FpgaUpgradeNeeded** | Pointer to **bool** | The flag to check vulnerability with secure boot technology. | [optional] [readonly] 
**InbandIpAddress** | Pointer to **string** | The IP address of the network Element inband management interface. | [optional] [readonly] 
**InbandIpGateway** | Pointer to **string** | The default gateway of the network Element inband management interface. | [optional] [readonly] 
**InbandIpMask** | Pointer to **string** | The network mask of the network Element inband management interface. | [optional] [readonly] 
**InbandVlan** | Pointer to **int64** | The VLAN ID of the network Element inband management interface. | [optional] [readonly] 
**InterClusterLinkState** | Pointer to **string** | The intercluster link state of the switch. * &#x60;Unknown&#x60; - The operational state of the link is not known. * &#x60;Up&#x60; - The operational state of the link is up. * &#x60;Down&#x60; - The operational state of the link is down. * &#x60;Degraded&#x60; - The link is operational but degraded. This state is applicable to port channels when any one of the member links is down. | [optional] [readonly] [default to "Unknown"]
**Ipv4Address** | Pointer to **string** | IP version 4 address is saved in this property. | [optional] [readonly] 
**IsUpgraded** | Pointer to **bool** | This field indicates the compute status of the catalog values for the associated component or hardware. | [optional] [readonly] [default to false]
**ManagementMode** | Pointer to **string** | The management mode of the fabric interconnect. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [readonly] [default to "IntersightStandalone"]
**Model** | Pointer to **string** | This field displays the model number of the associated component or hardware. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the ElementSummary object is saved in this property. | [optional] [readonly] 
**NumEtherPorts** | Pointer to **int64** | Total number of Ethernet ports. | [optional] [readonly] 
**NumEtherPortsConfigured** | Pointer to **int64** | Total number of configured Ethernet ports. | [optional] [readonly] 
**NumEtherPortsLinkUp** | Pointer to **int64** | Total number of Ethernet ports which are UP. | [optional] [readonly] 
**NumExpansionModules** | Pointer to **int64** | Total number of expansion modules. | [optional] [readonly] 
**NumFcPorts** | Pointer to **int64** | Total number of FC ports. | [optional] [readonly] 
**NumFcPortsConfigured** | Pointer to **int64** | Total number of configured FC ports. | [optional] [readonly] 
**NumFcPortsLinkUp** | Pointer to **int64** | Total number of FC ports which are UP. | [optional] [readonly] 
**OperEvacState** | Pointer to **string** | Operational state of the Fabric Evacuation feature, for this switch. * &#x60;&#x60; - Evacuation state of the switch is unknown. * &#x60;enabled&#x60; - Evacuation state of the switch is enabled. * &#x60;disabled&#x60; - Evacuation state of the switch is disabled. * &#x60;applying&#x60; - Evacuation state of the switch when evacuation is in progress. * &#x60;on&#x60; - Evacuation state of the switch is enabled. * &#x60;off&#x60; - Evacuation state of the switch is disabled. * &#x60;N/A&#x60; - Evacuation state of the switch is not applicable. | [optional] [readonly] [default to ""]
**Operability** | Pointer to **string** | The switch&#39;s current overall operational/health state. | [optional] [readonly] 
**OutOfBandIpAddress** | Pointer to **string** | The IP address of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpGateway** | Pointer to **string** | The default gateway of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpMask** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Address** | Pointer to **string** | The IPv4 address of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Gateway** | Pointer to **string** | The default IPv4 gateway of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv4Mask** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv6Address** | Pointer to **string** | The IPv6 address of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv6Gateway** | Pointer to **string** | The default IPv6 gateway of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpv6Prefix** | Pointer to **string** | The network mask of the network Element out-of-band management interface. | [optional] [readonly] 
**OutOfBandMac** | Pointer to **string** | The MAC address of the network Element out-of-band management interface. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number of the switch. | [optional] [readonly] 
**PeerFirmwareOutOfSync** | Pointer to **bool** | The flag to indicate the firmware of peer Fabric Interconnect is out of sync. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. | [optional] [readonly] 
**ReservedVlanStartId** | Pointer to **int64** | The reserved VLAN start ID of the Network Element. A block of 128 VLANs are reserved for internal use and cannot be used for carrying network traffic. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field displays the revised version of the associated component or hardware (if any). | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field displays the serial number of the associated component or hardware. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of this view MO. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the switch. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | The Switch Id of the network Element. | [optional] [readonly] 
**SwitchProfileName** | Pointer to **string** | The name of switch profile associated with the switch. | [optional] [readonly] 
**SwitchType** | Pointer to **string** | The Switch type that the network element is a part of. * &#x60;FabricInterconnect&#x60; - The default Switch type of UCSM and IMM mode devices. * &#x60;NexusDevice&#x60; - Switch type of Nexus devices. * &#x60;MDSDevice&#x60; - Switch type of Nexus MDS devices. | [optional] [readonly] [default to "FabricInterconnect"]
**SystemUpTime** | Pointer to **string** | System up time of the switch. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The Thermal status of the fabric interconnect. * &#x60;unknown&#x60; - The default state of the sensor (in case no data is received). * &#x60;ok&#x60; - State of the sensor indicating the sensor&#39;s temperature range is okay. * &#x60;upper-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely high above normal range. * &#x60;upper-critical&#x60; - State of the sensor indicating that the temperature is above normal range. * &#x60;upper-non-critical&#x60; - State of the sensor indicating that the temperature is a little above the normal range. * &#x60;lower-non-critical&#x60; - State of the sensor indicating that the temperature is a little below the normal range. * &#x60;lower-critical&#x60; - State of the sensor indicating that the temperature is below normal range. * &#x60;lower-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely below normal range. | [optional] [readonly] [default to "unknown"]
**TotalMemory** | Pointer to **int64** | Total available memory on this switch platform. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the switch. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field displays the vendor information of the associated component or hardware. | [optional] [readonly] 
**Version** | Pointer to **string** | Version holds the firmware version related information. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkElementSummary

`func NewNetworkElementSummary(classId string, objectType string, ) *NetworkElementSummary`

NewNetworkElementSummary instantiates a new NetworkElementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkElementSummaryWithDefaults

`func NewNetworkElementSummaryWithDefaults() *NetworkElementSummary`

NewNetworkElementSummaryWithDefaults instantiates a new NetworkElementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkElementSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkElementSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkElementSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkElementSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkElementSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkElementSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminEvacState

`func (o *NetworkElementSummary) GetAdminEvacState() string`

GetAdminEvacState returns the AdminEvacState field if non-nil, zero value otherwise.

### GetAdminEvacStateOk

`func (o *NetworkElementSummary) GetAdminEvacStateOk() (*string, bool)`

GetAdminEvacStateOk returns a tuple with the AdminEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEvacState

`func (o *NetworkElementSummary) SetAdminEvacState(v string)`

SetAdminEvacState sets AdminEvacState field to given value.

### HasAdminEvacState

`func (o *NetworkElementSummary) HasAdminEvacState() bool`

HasAdminEvacState returns a boolean if a field has been set.

### GetAdminInbandInterfaceState

`func (o *NetworkElementSummary) GetAdminInbandInterfaceState() string`

GetAdminInbandInterfaceState returns the AdminInbandInterfaceState field if non-nil, zero value otherwise.

### GetAdminInbandInterfaceStateOk

`func (o *NetworkElementSummary) GetAdminInbandInterfaceStateOk() (*string, bool)`

GetAdminInbandInterfaceStateOk returns a tuple with the AdminInbandInterfaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminInbandInterfaceState

`func (o *NetworkElementSummary) SetAdminInbandInterfaceState(v string)`

SetAdminInbandInterfaceState sets AdminInbandInterfaceState field to given value.

### HasAdminInbandInterfaceState

`func (o *NetworkElementSummary) HasAdminInbandInterfaceState() bool`

HasAdminInbandInterfaceState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *NetworkElementSummary) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *NetworkElementSummary) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *NetworkElementSummary) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *NetworkElementSummary) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *NetworkElementSummary) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *NetworkElementSummary) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetAvailableMemory

`func (o *NetworkElementSummary) GetAvailableMemory() string`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *NetworkElementSummary) GetAvailableMemoryOk() (*string, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *NetworkElementSummary) SetAvailableMemory(v string)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *NetworkElementSummary) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBundleVersion

`func (o *NetworkElementSummary) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *NetworkElementSummary) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *NetworkElementSummary) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *NetworkElementSummary) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetChassis

`func (o *NetworkElementSummary) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *NetworkElementSummary) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *NetworkElementSummary) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *NetworkElementSummary) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetConfModTs

`func (o *NetworkElementSummary) GetConfModTs() string`

GetConfModTs returns the ConfModTs field if non-nil, zero value otherwise.

### GetConfModTsOk

`func (o *NetworkElementSummary) GetConfModTsOk() (*string, bool)`

GetConfModTsOk returns a tuple with the ConfModTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfModTs

`func (o *NetworkElementSummary) SetConfModTs(v string)`

SetConfModTs sets ConfModTs field to given value.

### HasConfModTs

`func (o *NetworkElementSummary) HasConfModTs() bool`

HasConfModTs returns a boolean if a field has been set.

### GetConfModTsBackup

`func (o *NetworkElementSummary) GetConfModTsBackup() string`

GetConfModTsBackup returns the ConfModTsBackup field if non-nil, zero value otherwise.

### GetConfModTsBackupOk

`func (o *NetworkElementSummary) GetConfModTsBackupOk() (*string, bool)`

GetConfModTsBackupOk returns a tuple with the ConfModTsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfModTsBackup

`func (o *NetworkElementSummary) SetConfModTsBackup(v string)`

SetConfModTsBackup sets ConfModTsBackup field to given value.

### HasConfModTsBackup

`func (o *NetworkElementSummary) HasConfModTsBackup() bool`

HasConfModTsBackup returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *NetworkElementSummary) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *NetworkElementSummary) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *NetworkElementSummary) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *NetworkElementSummary) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDefaultDomain

`func (o *NetworkElementSummary) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *NetworkElementSummary) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *NetworkElementSummary) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *NetworkElementSummary) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetDeviceMoId

`func (o *NetworkElementSummary) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *NetworkElementSummary) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *NetworkElementSummary) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *NetworkElementSummary) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *NetworkElementSummary) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NetworkElementSummary) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NetworkElementSummary) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NetworkElementSummary) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEthernetMode

`func (o *NetworkElementSummary) GetEthernetMode() string`

GetEthernetMode returns the EthernetMode field if non-nil, zero value otherwise.

### GetEthernetModeOk

`func (o *NetworkElementSummary) GetEthernetModeOk() (*string, bool)`

GetEthernetModeOk returns a tuple with the EthernetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMode

`func (o *NetworkElementSummary) SetEthernetMode(v string)`

SetEthernetMode sets EthernetMode field to given value.

### HasEthernetMode

`func (o *NetworkElementSummary) HasEthernetMode() bool`

HasEthernetMode returns a boolean if a field has been set.

### GetEthernetSwitchingMode

`func (o *NetworkElementSummary) GetEthernetSwitchingMode() string`

GetEthernetSwitchingMode returns the EthernetSwitchingMode field if non-nil, zero value otherwise.

### GetEthernetSwitchingModeOk

`func (o *NetworkElementSummary) GetEthernetSwitchingModeOk() (*string, bool)`

GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetSwitchingMode

`func (o *NetworkElementSummary) SetEthernetSwitchingMode(v string)`

SetEthernetSwitchingMode sets EthernetSwitchingMode field to given value.

### HasEthernetSwitchingMode

`func (o *NetworkElementSummary) HasEthernetSwitchingMode() bool`

HasEthernetSwitchingMode returns a boolean if a field has been set.

### GetFaultSummary

`func (o *NetworkElementSummary) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *NetworkElementSummary) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *NetworkElementSummary) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *NetworkElementSummary) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFcMode

`func (o *NetworkElementSummary) GetFcMode() string`

GetFcMode returns the FcMode field if non-nil, zero value otherwise.

### GetFcModeOk

`func (o *NetworkElementSummary) GetFcModeOk() (*string, bool)`

GetFcModeOk returns a tuple with the FcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcMode

`func (o *NetworkElementSummary) SetFcMode(v string)`

SetFcMode sets FcMode field to given value.

### HasFcMode

`func (o *NetworkElementSummary) HasFcMode() bool`

HasFcMode returns a boolean if a field has been set.

### GetFcSwitchingMode

`func (o *NetworkElementSummary) GetFcSwitchingMode() string`

GetFcSwitchingMode returns the FcSwitchingMode field if non-nil, zero value otherwise.

### GetFcSwitchingModeOk

`func (o *NetworkElementSummary) GetFcSwitchingModeOk() (*string, bool)`

GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSwitchingMode

`func (o *NetworkElementSummary) SetFcSwitchingMode(v string)`

SetFcSwitchingMode sets FcSwitchingMode field to given value.

### HasFcSwitchingMode

`func (o *NetworkElementSummary) HasFcSwitchingMode() bool`

HasFcSwitchingMode returns a boolean if a field has been set.

### GetFirmware

`func (o *NetworkElementSummary) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *NetworkElementSummary) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *NetworkElementSummary) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *NetworkElementSummary) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *NetworkElementSummary) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *NetworkElementSummary) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *NetworkElementSummary) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *NetworkElementSummary) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFpgaUpgradeNeeded

`func (o *NetworkElementSummary) GetFpgaUpgradeNeeded() bool`

GetFpgaUpgradeNeeded returns the FpgaUpgradeNeeded field if non-nil, zero value otherwise.

### GetFpgaUpgradeNeededOk

`func (o *NetworkElementSummary) GetFpgaUpgradeNeededOk() (*bool, bool)`

GetFpgaUpgradeNeededOk returns a tuple with the FpgaUpgradeNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpgaUpgradeNeeded

`func (o *NetworkElementSummary) SetFpgaUpgradeNeeded(v bool)`

SetFpgaUpgradeNeeded sets FpgaUpgradeNeeded field to given value.

### HasFpgaUpgradeNeeded

`func (o *NetworkElementSummary) HasFpgaUpgradeNeeded() bool`

HasFpgaUpgradeNeeded returns a boolean if a field has been set.

### GetInbandIpAddress

`func (o *NetworkElementSummary) GetInbandIpAddress() string`

GetInbandIpAddress returns the InbandIpAddress field if non-nil, zero value otherwise.

### GetInbandIpAddressOk

`func (o *NetworkElementSummary) GetInbandIpAddressOk() (*string, bool)`

GetInbandIpAddressOk returns a tuple with the InbandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpAddress

`func (o *NetworkElementSummary) SetInbandIpAddress(v string)`

SetInbandIpAddress sets InbandIpAddress field to given value.

### HasInbandIpAddress

`func (o *NetworkElementSummary) HasInbandIpAddress() bool`

HasInbandIpAddress returns a boolean if a field has been set.

### GetInbandIpGateway

`func (o *NetworkElementSummary) GetInbandIpGateway() string`

GetInbandIpGateway returns the InbandIpGateway field if non-nil, zero value otherwise.

### GetInbandIpGatewayOk

`func (o *NetworkElementSummary) GetInbandIpGatewayOk() (*string, bool)`

GetInbandIpGatewayOk returns a tuple with the InbandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpGateway

`func (o *NetworkElementSummary) SetInbandIpGateway(v string)`

SetInbandIpGateway sets InbandIpGateway field to given value.

### HasInbandIpGateway

`func (o *NetworkElementSummary) HasInbandIpGateway() bool`

HasInbandIpGateway returns a boolean if a field has been set.

### GetInbandIpMask

`func (o *NetworkElementSummary) GetInbandIpMask() string`

GetInbandIpMask returns the InbandIpMask field if non-nil, zero value otherwise.

### GetInbandIpMaskOk

`func (o *NetworkElementSummary) GetInbandIpMaskOk() (*string, bool)`

GetInbandIpMaskOk returns a tuple with the InbandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpMask

`func (o *NetworkElementSummary) SetInbandIpMask(v string)`

SetInbandIpMask sets InbandIpMask field to given value.

### HasInbandIpMask

`func (o *NetworkElementSummary) HasInbandIpMask() bool`

HasInbandIpMask returns a boolean if a field has been set.

### GetInbandVlan

`func (o *NetworkElementSummary) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *NetworkElementSummary) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *NetworkElementSummary) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *NetworkElementSummary) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInterClusterLinkState

`func (o *NetworkElementSummary) GetInterClusterLinkState() string`

GetInterClusterLinkState returns the InterClusterLinkState field if non-nil, zero value otherwise.

### GetInterClusterLinkStateOk

`func (o *NetworkElementSummary) GetInterClusterLinkStateOk() (*string, bool)`

GetInterClusterLinkStateOk returns a tuple with the InterClusterLinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterClusterLinkState

`func (o *NetworkElementSummary) SetInterClusterLinkState(v string)`

SetInterClusterLinkState sets InterClusterLinkState field to given value.

### HasInterClusterLinkState

`func (o *NetworkElementSummary) HasInterClusterLinkState() bool`

HasInterClusterLinkState returns a boolean if a field has been set.

### GetIpv4Address

`func (o *NetworkElementSummary) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *NetworkElementSummary) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *NetworkElementSummary) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *NetworkElementSummary) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIsUpgraded

`func (o *NetworkElementSummary) GetIsUpgraded() bool`

GetIsUpgraded returns the IsUpgraded field if non-nil, zero value otherwise.

### GetIsUpgradedOk

`func (o *NetworkElementSummary) GetIsUpgradedOk() (*bool, bool)`

GetIsUpgradedOk returns a tuple with the IsUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgraded

`func (o *NetworkElementSummary) SetIsUpgraded(v bool)`

SetIsUpgraded sets IsUpgraded field to given value.

### HasIsUpgraded

`func (o *NetworkElementSummary) HasIsUpgraded() bool`

HasIsUpgraded returns a boolean if a field has been set.

### GetManagementMode

`func (o *NetworkElementSummary) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *NetworkElementSummary) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *NetworkElementSummary) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *NetworkElementSummary) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetModel

`func (o *NetworkElementSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NetworkElementSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NetworkElementSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NetworkElementSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *NetworkElementSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkElementSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkElementSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkElementSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumEtherPorts

`func (o *NetworkElementSummary) GetNumEtherPorts() int64`

GetNumEtherPorts returns the NumEtherPorts field if non-nil, zero value otherwise.

### GetNumEtherPortsOk

`func (o *NetworkElementSummary) GetNumEtherPortsOk() (*int64, bool)`

GetNumEtherPortsOk returns a tuple with the NumEtherPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPorts

`func (o *NetworkElementSummary) SetNumEtherPorts(v int64)`

SetNumEtherPorts sets NumEtherPorts field to given value.

### HasNumEtherPorts

`func (o *NetworkElementSummary) HasNumEtherPorts() bool`

HasNumEtherPorts returns a boolean if a field has been set.

### GetNumEtherPortsConfigured

`func (o *NetworkElementSummary) GetNumEtherPortsConfigured() int64`

GetNumEtherPortsConfigured returns the NumEtherPortsConfigured field if non-nil, zero value otherwise.

### GetNumEtherPortsConfiguredOk

`func (o *NetworkElementSummary) GetNumEtherPortsConfiguredOk() (*int64, bool)`

GetNumEtherPortsConfiguredOk returns a tuple with the NumEtherPortsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPortsConfigured

`func (o *NetworkElementSummary) SetNumEtherPortsConfigured(v int64)`

SetNumEtherPortsConfigured sets NumEtherPortsConfigured field to given value.

### HasNumEtherPortsConfigured

`func (o *NetworkElementSummary) HasNumEtherPortsConfigured() bool`

HasNumEtherPortsConfigured returns a boolean if a field has been set.

### GetNumEtherPortsLinkUp

`func (o *NetworkElementSummary) GetNumEtherPortsLinkUp() int64`

GetNumEtherPortsLinkUp returns the NumEtherPortsLinkUp field if non-nil, zero value otherwise.

### GetNumEtherPortsLinkUpOk

`func (o *NetworkElementSummary) GetNumEtherPortsLinkUpOk() (*int64, bool)`

GetNumEtherPortsLinkUpOk returns a tuple with the NumEtherPortsLinkUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPortsLinkUp

`func (o *NetworkElementSummary) SetNumEtherPortsLinkUp(v int64)`

SetNumEtherPortsLinkUp sets NumEtherPortsLinkUp field to given value.

### HasNumEtherPortsLinkUp

`func (o *NetworkElementSummary) HasNumEtherPortsLinkUp() bool`

HasNumEtherPortsLinkUp returns a boolean if a field has been set.

### GetNumExpansionModules

`func (o *NetworkElementSummary) GetNumExpansionModules() int64`

GetNumExpansionModules returns the NumExpansionModules field if non-nil, zero value otherwise.

### GetNumExpansionModulesOk

`func (o *NetworkElementSummary) GetNumExpansionModulesOk() (*int64, bool)`

GetNumExpansionModulesOk returns a tuple with the NumExpansionModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExpansionModules

`func (o *NetworkElementSummary) SetNumExpansionModules(v int64)`

SetNumExpansionModules sets NumExpansionModules field to given value.

### HasNumExpansionModules

`func (o *NetworkElementSummary) HasNumExpansionModules() bool`

HasNumExpansionModules returns a boolean if a field has been set.

### GetNumFcPorts

`func (o *NetworkElementSummary) GetNumFcPorts() int64`

GetNumFcPorts returns the NumFcPorts field if non-nil, zero value otherwise.

### GetNumFcPortsOk

`func (o *NetworkElementSummary) GetNumFcPortsOk() (*int64, bool)`

GetNumFcPortsOk returns a tuple with the NumFcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPorts

`func (o *NetworkElementSummary) SetNumFcPorts(v int64)`

SetNumFcPorts sets NumFcPorts field to given value.

### HasNumFcPorts

`func (o *NetworkElementSummary) HasNumFcPorts() bool`

HasNumFcPorts returns a boolean if a field has been set.

### GetNumFcPortsConfigured

`func (o *NetworkElementSummary) GetNumFcPortsConfigured() int64`

GetNumFcPortsConfigured returns the NumFcPortsConfigured field if non-nil, zero value otherwise.

### GetNumFcPortsConfiguredOk

`func (o *NetworkElementSummary) GetNumFcPortsConfiguredOk() (*int64, bool)`

GetNumFcPortsConfiguredOk returns a tuple with the NumFcPortsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPortsConfigured

`func (o *NetworkElementSummary) SetNumFcPortsConfigured(v int64)`

SetNumFcPortsConfigured sets NumFcPortsConfigured field to given value.

### HasNumFcPortsConfigured

`func (o *NetworkElementSummary) HasNumFcPortsConfigured() bool`

HasNumFcPortsConfigured returns a boolean if a field has been set.

### GetNumFcPortsLinkUp

`func (o *NetworkElementSummary) GetNumFcPortsLinkUp() int64`

GetNumFcPortsLinkUp returns the NumFcPortsLinkUp field if non-nil, zero value otherwise.

### GetNumFcPortsLinkUpOk

`func (o *NetworkElementSummary) GetNumFcPortsLinkUpOk() (*int64, bool)`

GetNumFcPortsLinkUpOk returns a tuple with the NumFcPortsLinkUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPortsLinkUp

`func (o *NetworkElementSummary) SetNumFcPortsLinkUp(v int64)`

SetNumFcPortsLinkUp sets NumFcPortsLinkUp field to given value.

### HasNumFcPortsLinkUp

`func (o *NetworkElementSummary) HasNumFcPortsLinkUp() bool`

HasNumFcPortsLinkUp returns a boolean if a field has been set.

### GetOperEvacState

`func (o *NetworkElementSummary) GetOperEvacState() string`

GetOperEvacState returns the OperEvacState field if non-nil, zero value otherwise.

### GetOperEvacStateOk

`func (o *NetworkElementSummary) GetOperEvacStateOk() (*string, bool)`

GetOperEvacStateOk returns a tuple with the OperEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperEvacState

`func (o *NetworkElementSummary) SetOperEvacState(v string)`

SetOperEvacState sets OperEvacState field to given value.

### HasOperEvacState

`func (o *NetworkElementSummary) HasOperEvacState() bool`

HasOperEvacState returns a boolean if a field has been set.

### GetOperability

`func (o *NetworkElementSummary) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *NetworkElementSummary) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *NetworkElementSummary) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *NetworkElementSummary) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *NetworkElementSummary) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *NetworkElementSummary) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *NetworkElementSummary) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *NetworkElementSummary) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *NetworkElementSummary) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *NetworkElementSummary) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *NetworkElementSummary) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *NetworkElementSummary) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetOutOfBandIpMask

`func (o *NetworkElementSummary) GetOutOfBandIpMask() string`

GetOutOfBandIpMask returns the OutOfBandIpMask field if non-nil, zero value otherwise.

### GetOutOfBandIpMaskOk

`func (o *NetworkElementSummary) GetOutOfBandIpMaskOk() (*string, bool)`

GetOutOfBandIpMaskOk returns a tuple with the OutOfBandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpMask

`func (o *NetworkElementSummary) SetOutOfBandIpMask(v string)`

SetOutOfBandIpMask sets OutOfBandIpMask field to given value.

### HasOutOfBandIpMask

`func (o *NetworkElementSummary) HasOutOfBandIpMask() bool`

HasOutOfBandIpMask returns a boolean if a field has been set.

### GetOutOfBandIpv4Address

`func (o *NetworkElementSummary) GetOutOfBandIpv4Address() string`

GetOutOfBandIpv4Address returns the OutOfBandIpv4Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv4AddressOk

`func (o *NetworkElementSummary) GetOutOfBandIpv4AddressOk() (*string, bool)`

GetOutOfBandIpv4AddressOk returns a tuple with the OutOfBandIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Address

`func (o *NetworkElementSummary) SetOutOfBandIpv4Address(v string)`

SetOutOfBandIpv4Address sets OutOfBandIpv4Address field to given value.

### HasOutOfBandIpv4Address

`func (o *NetworkElementSummary) HasOutOfBandIpv4Address() bool`

HasOutOfBandIpv4Address returns a boolean if a field has been set.

### GetOutOfBandIpv4Gateway

`func (o *NetworkElementSummary) GetOutOfBandIpv4Gateway() string`

GetOutOfBandIpv4Gateway returns the OutOfBandIpv4Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv4GatewayOk

`func (o *NetworkElementSummary) GetOutOfBandIpv4GatewayOk() (*string, bool)`

GetOutOfBandIpv4GatewayOk returns a tuple with the OutOfBandIpv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Gateway

`func (o *NetworkElementSummary) SetOutOfBandIpv4Gateway(v string)`

SetOutOfBandIpv4Gateway sets OutOfBandIpv4Gateway field to given value.

### HasOutOfBandIpv4Gateway

`func (o *NetworkElementSummary) HasOutOfBandIpv4Gateway() bool`

HasOutOfBandIpv4Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv4Mask

`func (o *NetworkElementSummary) GetOutOfBandIpv4Mask() string`

GetOutOfBandIpv4Mask returns the OutOfBandIpv4Mask field if non-nil, zero value otherwise.

### GetOutOfBandIpv4MaskOk

`func (o *NetworkElementSummary) GetOutOfBandIpv4MaskOk() (*string, bool)`

GetOutOfBandIpv4MaskOk returns a tuple with the OutOfBandIpv4Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Mask

`func (o *NetworkElementSummary) SetOutOfBandIpv4Mask(v string)`

SetOutOfBandIpv4Mask sets OutOfBandIpv4Mask field to given value.

### HasOutOfBandIpv4Mask

`func (o *NetworkElementSummary) HasOutOfBandIpv4Mask() bool`

HasOutOfBandIpv4Mask returns a boolean if a field has been set.

### GetOutOfBandIpv6Address

`func (o *NetworkElementSummary) GetOutOfBandIpv6Address() string`

GetOutOfBandIpv6Address returns the OutOfBandIpv6Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv6AddressOk

`func (o *NetworkElementSummary) GetOutOfBandIpv6AddressOk() (*string, bool)`

GetOutOfBandIpv6AddressOk returns a tuple with the OutOfBandIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Address

`func (o *NetworkElementSummary) SetOutOfBandIpv6Address(v string)`

SetOutOfBandIpv6Address sets OutOfBandIpv6Address field to given value.

### HasOutOfBandIpv6Address

`func (o *NetworkElementSummary) HasOutOfBandIpv6Address() bool`

HasOutOfBandIpv6Address returns a boolean if a field has been set.

### GetOutOfBandIpv6Gateway

`func (o *NetworkElementSummary) GetOutOfBandIpv6Gateway() string`

GetOutOfBandIpv6Gateway returns the OutOfBandIpv6Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv6GatewayOk

`func (o *NetworkElementSummary) GetOutOfBandIpv6GatewayOk() (*string, bool)`

GetOutOfBandIpv6GatewayOk returns a tuple with the OutOfBandIpv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Gateway

`func (o *NetworkElementSummary) SetOutOfBandIpv6Gateway(v string)`

SetOutOfBandIpv6Gateway sets OutOfBandIpv6Gateway field to given value.

### HasOutOfBandIpv6Gateway

`func (o *NetworkElementSummary) HasOutOfBandIpv6Gateway() bool`

HasOutOfBandIpv6Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv6Prefix

`func (o *NetworkElementSummary) GetOutOfBandIpv6Prefix() string`

GetOutOfBandIpv6Prefix returns the OutOfBandIpv6Prefix field if non-nil, zero value otherwise.

### GetOutOfBandIpv6PrefixOk

`func (o *NetworkElementSummary) GetOutOfBandIpv6PrefixOk() (*string, bool)`

GetOutOfBandIpv6PrefixOk returns a tuple with the OutOfBandIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Prefix

`func (o *NetworkElementSummary) SetOutOfBandIpv6Prefix(v string)`

SetOutOfBandIpv6Prefix sets OutOfBandIpv6Prefix field to given value.

### HasOutOfBandIpv6Prefix

`func (o *NetworkElementSummary) HasOutOfBandIpv6Prefix() bool`

HasOutOfBandIpv6Prefix returns a boolean if a field has been set.

### GetOutOfBandMac

`func (o *NetworkElementSummary) GetOutOfBandMac() string`

GetOutOfBandMac returns the OutOfBandMac field if non-nil, zero value otherwise.

### GetOutOfBandMacOk

`func (o *NetworkElementSummary) GetOutOfBandMacOk() (*string, bool)`

GetOutOfBandMacOk returns a tuple with the OutOfBandMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandMac

`func (o *NetworkElementSummary) SetOutOfBandMac(v string)`

SetOutOfBandMac sets OutOfBandMac field to given value.

### HasOutOfBandMac

`func (o *NetworkElementSummary) HasOutOfBandMac() bool`

HasOutOfBandMac returns a boolean if a field has been set.

### GetPartNumber

`func (o *NetworkElementSummary) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *NetworkElementSummary) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *NetworkElementSummary) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *NetworkElementSummary) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPeerFirmwareOutOfSync

`func (o *NetworkElementSummary) GetPeerFirmwareOutOfSync() bool`

GetPeerFirmwareOutOfSync returns the PeerFirmwareOutOfSync field if non-nil, zero value otherwise.

### GetPeerFirmwareOutOfSyncOk

`func (o *NetworkElementSummary) GetPeerFirmwareOutOfSyncOk() (*bool, bool)`

GetPeerFirmwareOutOfSyncOk returns a tuple with the PeerFirmwareOutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerFirmwareOutOfSync

`func (o *NetworkElementSummary) SetPeerFirmwareOutOfSync(v bool)`

SetPeerFirmwareOutOfSync sets PeerFirmwareOutOfSync field to given value.

### HasPeerFirmwareOutOfSync

`func (o *NetworkElementSummary) HasPeerFirmwareOutOfSync() bool`

HasPeerFirmwareOutOfSync returns a boolean if a field has been set.

### GetPresence

`func (o *NetworkElementSummary) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *NetworkElementSummary) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *NetworkElementSummary) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *NetworkElementSummary) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetReservedVlanStartId

`func (o *NetworkElementSummary) GetReservedVlanStartId() int64`

GetReservedVlanStartId returns the ReservedVlanStartId field if non-nil, zero value otherwise.

### GetReservedVlanStartIdOk

`func (o *NetworkElementSummary) GetReservedVlanStartIdOk() (*int64, bool)`

GetReservedVlanStartIdOk returns a tuple with the ReservedVlanStartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlanStartId

`func (o *NetworkElementSummary) SetReservedVlanStartId(v int64)`

SetReservedVlanStartId sets ReservedVlanStartId field to given value.

### HasReservedVlanStartId

`func (o *NetworkElementSummary) HasReservedVlanStartId() bool`

HasReservedVlanStartId returns a boolean if a field has been set.

### GetRevision

`func (o *NetworkElementSummary) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkElementSummary) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkElementSummary) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NetworkElementSummary) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRn

`func (o *NetworkElementSummary) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *NetworkElementSummary) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *NetworkElementSummary) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *NetworkElementSummary) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetSerial

`func (o *NetworkElementSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworkElementSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworkElementSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NetworkElementSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *NetworkElementSummary) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *NetworkElementSummary) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *NetworkElementSummary) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *NetworkElementSummary) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkElementSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkElementSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkElementSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkElementSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *NetworkElementSummary) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkElementSummary) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkElementSummary) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *NetworkElementSummary) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchProfileName

`func (o *NetworkElementSummary) GetSwitchProfileName() string`

GetSwitchProfileName returns the SwitchProfileName field if non-nil, zero value otherwise.

### GetSwitchProfileNameOk

`func (o *NetworkElementSummary) GetSwitchProfileNameOk() (*string, bool)`

GetSwitchProfileNameOk returns a tuple with the SwitchProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileName

`func (o *NetworkElementSummary) SetSwitchProfileName(v string)`

SetSwitchProfileName sets SwitchProfileName field to given value.

### HasSwitchProfileName

`func (o *NetworkElementSummary) HasSwitchProfileName() bool`

HasSwitchProfileName returns a boolean if a field has been set.

### GetSwitchType

`func (o *NetworkElementSummary) GetSwitchType() string`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *NetworkElementSummary) GetSwitchTypeOk() (*string, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *NetworkElementSummary) SetSwitchType(v string)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *NetworkElementSummary) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.

### GetSystemUpTime

`func (o *NetworkElementSummary) GetSystemUpTime() string`

GetSystemUpTime returns the SystemUpTime field if non-nil, zero value otherwise.

### GetSystemUpTimeOk

`func (o *NetworkElementSummary) GetSystemUpTimeOk() (*string, bool)`

GetSystemUpTimeOk returns a tuple with the SystemUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpTime

`func (o *NetworkElementSummary) SetSystemUpTime(v string)`

SetSystemUpTime sets SystemUpTime field to given value.

### HasSystemUpTime

`func (o *NetworkElementSummary) HasSystemUpTime() bool`

HasSystemUpTime returns a boolean if a field has been set.

### GetThermal

`func (o *NetworkElementSummary) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *NetworkElementSummary) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *NetworkElementSummary) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *NetworkElementSummary) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetTotalMemory

`func (o *NetworkElementSummary) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *NetworkElementSummary) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *NetworkElementSummary) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *NetworkElementSummary) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetUserLabel

`func (o *NetworkElementSummary) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *NetworkElementSummary) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *NetworkElementSummary) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *NetworkElementSummary) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetVendor

`func (o *NetworkElementSummary) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *NetworkElementSummary) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *NetworkElementSummary) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *NetworkElementSummary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkElementSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkElementSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkElementSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkElementSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkElementSummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkElementSummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkElementSummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkElementSummary) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkElementSummary) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkElementSummary) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


