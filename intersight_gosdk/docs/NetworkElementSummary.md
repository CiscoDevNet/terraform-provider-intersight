# NetworkElementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEvacState** | Pointer to **string** | Administratively configured state of Fabric Evacuation feature, for this switch. | [optional] [readonly] 
**AdminInbandInterfaceState** | Pointer to **string** | The administrative state of the network Element inband management interface. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**ComputeAlarmSummary**](compute.AlarmSummary.md) |  | [optional] 
**AvailableMemory** | Pointer to **string** | Available memory (un-used) on this switch platform. | [optional] [readonly] 
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**EthernetMode** | Pointer to **string** | The user configured Ethernet operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**FaultSummary** | Pointer to **int64** | The fault summary of the network Element out-of-band management interface. | [optional] [readonly] 
**FcMode** | Pointer to **string** | The user configured FC operational mode for this switch (End-Host or Switching). | [optional] [readonly] 
**Firmware** | Pointer to **string** | Running firmware information. | [optional] [readonly] 
**InbandIpAddress** | Pointer to **string** | The IP address of the network Element inband management interface. | [optional] [readonly] 
**InbandIpGateway** | Pointer to **string** | The default gateway of the network Element inband management interface. | [optional] [readonly] 
**InbandIpMask** | Pointer to **string** | The network mask of the network Element inband management interface. | [optional] [readonly] 
**InbandVlan** | Pointer to **int64** | The VLAN ID of the network Element inband management interface. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | IP version 4 address is saved in this property. | [optional] [readonly] 
**ManagementMode** | Pointer to **string** | The management mode of the fabric interconnect. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [readonly] [default to "IntersightStandalone"]
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the ElementSummary object is saved in this property. | [optional] [readonly] 
**NumEtherPorts** | Pointer to **int64** | Total number of Ethernet ports. | [optional] [readonly] 
**NumEtherPortsConfigured** | Pointer to **int64** | Total number of configured Ethernet ports. | [optional] [readonly] 
**NumEtherPortsLinkUp** | Pointer to **int64** | Total number of Ethernet ports which are UP. | [optional] [readonly] 
**NumExpansionModules** | Pointer to **int64** | Total number of expansion modules. | [optional] [readonly] 
**NumFcPorts** | Pointer to **int64** | Total number of FC ports. | [optional] [readonly] 
**NumFcPortsConfigured** | Pointer to **int64** | Total number of configured FC ports. | [optional] [readonly] 
**NumFcPortsLinkUp** | Pointer to **int64** | Total number of FC ports which are UP. | [optional] [readonly] 
**OperEvacState** | Pointer to **string** | Operational state of the Fabric Evacuation feature, for this switch. | [optional] [readonly] 
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
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of this view MO. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | The Switch Id of the network Element. | [optional] [readonly] 
**TotalMemory** | Pointer to **int64** | Total available memory on this switch platform. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**Version** | Pointer to **string** | Version holds the firmware version related information. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNetworkElementSummary

`func NewNetworkElementSummary() *NetworkElementSummary`

NewNetworkElementSummary instantiates a new NetworkElementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkElementSummaryWithDefaults

`func NewNetworkElementSummaryWithDefaults() *NetworkElementSummary`

NewNetworkElementSummaryWithDefaults instantiates a new NetworkElementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


