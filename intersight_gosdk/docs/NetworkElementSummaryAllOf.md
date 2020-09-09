# NetworkElementSummaryAllOf

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

### NewNetworkElementSummaryAllOf

`func NewNetworkElementSummaryAllOf() *NetworkElementSummaryAllOf`

NewNetworkElementSummaryAllOf instantiates a new NetworkElementSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkElementSummaryAllOfWithDefaults

`func NewNetworkElementSummaryAllOfWithDefaults() *NetworkElementSummaryAllOf`

NewNetworkElementSummaryAllOfWithDefaults instantiates a new NetworkElementSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEvacState

`func (o *NetworkElementSummaryAllOf) GetAdminEvacState() string`

GetAdminEvacState returns the AdminEvacState field if non-nil, zero value otherwise.

### GetAdminEvacStateOk

`func (o *NetworkElementSummaryAllOf) GetAdminEvacStateOk() (*string, bool)`

GetAdminEvacStateOk returns a tuple with the AdminEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEvacState

`func (o *NetworkElementSummaryAllOf) SetAdminEvacState(v string)`

SetAdminEvacState sets AdminEvacState field to given value.

### HasAdminEvacState

`func (o *NetworkElementSummaryAllOf) HasAdminEvacState() bool`

HasAdminEvacState returns a boolean if a field has been set.

### GetAdminInbandInterfaceState

`func (o *NetworkElementSummaryAllOf) GetAdminInbandInterfaceState() string`

GetAdminInbandInterfaceState returns the AdminInbandInterfaceState field if non-nil, zero value otherwise.

### GetAdminInbandInterfaceStateOk

`func (o *NetworkElementSummaryAllOf) GetAdminInbandInterfaceStateOk() (*string, bool)`

GetAdminInbandInterfaceStateOk returns a tuple with the AdminInbandInterfaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminInbandInterfaceState

`func (o *NetworkElementSummaryAllOf) SetAdminInbandInterfaceState(v string)`

SetAdminInbandInterfaceState sets AdminInbandInterfaceState field to given value.

### HasAdminInbandInterfaceState

`func (o *NetworkElementSummaryAllOf) HasAdminInbandInterfaceState() bool`

HasAdminInbandInterfaceState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *NetworkElementSummaryAllOf) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *NetworkElementSummaryAllOf) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *NetworkElementSummaryAllOf) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *NetworkElementSummaryAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *NetworkElementSummaryAllOf) GetAvailableMemory() string`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *NetworkElementSummaryAllOf) GetAvailableMemoryOk() (*string, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *NetworkElementSummaryAllOf) SetAvailableMemory(v string)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *NetworkElementSummaryAllOf) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetDeviceMoId

`func (o *NetworkElementSummaryAllOf) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *NetworkElementSummaryAllOf) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *NetworkElementSummaryAllOf) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *NetworkElementSummaryAllOf) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *NetworkElementSummaryAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NetworkElementSummaryAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NetworkElementSummaryAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NetworkElementSummaryAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEthernetMode

`func (o *NetworkElementSummaryAllOf) GetEthernetMode() string`

GetEthernetMode returns the EthernetMode field if non-nil, zero value otherwise.

### GetEthernetModeOk

`func (o *NetworkElementSummaryAllOf) GetEthernetModeOk() (*string, bool)`

GetEthernetModeOk returns a tuple with the EthernetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMode

`func (o *NetworkElementSummaryAllOf) SetEthernetMode(v string)`

SetEthernetMode sets EthernetMode field to given value.

### HasEthernetMode

`func (o *NetworkElementSummaryAllOf) HasEthernetMode() bool`

HasEthernetMode returns a boolean if a field has been set.

### GetFaultSummary

`func (o *NetworkElementSummaryAllOf) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *NetworkElementSummaryAllOf) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *NetworkElementSummaryAllOf) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *NetworkElementSummaryAllOf) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFcMode

`func (o *NetworkElementSummaryAllOf) GetFcMode() string`

GetFcMode returns the FcMode field if non-nil, zero value otherwise.

### GetFcModeOk

`func (o *NetworkElementSummaryAllOf) GetFcModeOk() (*string, bool)`

GetFcModeOk returns a tuple with the FcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcMode

`func (o *NetworkElementSummaryAllOf) SetFcMode(v string)`

SetFcMode sets FcMode field to given value.

### HasFcMode

`func (o *NetworkElementSummaryAllOf) HasFcMode() bool`

HasFcMode returns a boolean if a field has been set.

### GetFirmware

`func (o *NetworkElementSummaryAllOf) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *NetworkElementSummaryAllOf) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *NetworkElementSummaryAllOf) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *NetworkElementSummaryAllOf) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetInbandIpAddress

`func (o *NetworkElementSummaryAllOf) GetInbandIpAddress() string`

GetInbandIpAddress returns the InbandIpAddress field if non-nil, zero value otherwise.

### GetInbandIpAddressOk

`func (o *NetworkElementSummaryAllOf) GetInbandIpAddressOk() (*string, bool)`

GetInbandIpAddressOk returns a tuple with the InbandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpAddress

`func (o *NetworkElementSummaryAllOf) SetInbandIpAddress(v string)`

SetInbandIpAddress sets InbandIpAddress field to given value.

### HasInbandIpAddress

`func (o *NetworkElementSummaryAllOf) HasInbandIpAddress() bool`

HasInbandIpAddress returns a boolean if a field has been set.

### GetInbandIpGateway

`func (o *NetworkElementSummaryAllOf) GetInbandIpGateway() string`

GetInbandIpGateway returns the InbandIpGateway field if non-nil, zero value otherwise.

### GetInbandIpGatewayOk

`func (o *NetworkElementSummaryAllOf) GetInbandIpGatewayOk() (*string, bool)`

GetInbandIpGatewayOk returns a tuple with the InbandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpGateway

`func (o *NetworkElementSummaryAllOf) SetInbandIpGateway(v string)`

SetInbandIpGateway sets InbandIpGateway field to given value.

### HasInbandIpGateway

`func (o *NetworkElementSummaryAllOf) HasInbandIpGateway() bool`

HasInbandIpGateway returns a boolean if a field has been set.

### GetInbandIpMask

`func (o *NetworkElementSummaryAllOf) GetInbandIpMask() string`

GetInbandIpMask returns the InbandIpMask field if non-nil, zero value otherwise.

### GetInbandIpMaskOk

`func (o *NetworkElementSummaryAllOf) GetInbandIpMaskOk() (*string, bool)`

GetInbandIpMaskOk returns a tuple with the InbandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpMask

`func (o *NetworkElementSummaryAllOf) SetInbandIpMask(v string)`

SetInbandIpMask sets InbandIpMask field to given value.

### HasInbandIpMask

`func (o *NetworkElementSummaryAllOf) HasInbandIpMask() bool`

HasInbandIpMask returns a boolean if a field has been set.

### GetInbandVlan

`func (o *NetworkElementSummaryAllOf) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *NetworkElementSummaryAllOf) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *NetworkElementSummaryAllOf) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *NetworkElementSummaryAllOf) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetIpv4Address

`func (o *NetworkElementSummaryAllOf) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *NetworkElementSummaryAllOf) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *NetworkElementSummaryAllOf) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *NetworkElementSummaryAllOf) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetManagementMode

`func (o *NetworkElementSummaryAllOf) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *NetworkElementSummaryAllOf) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *NetworkElementSummaryAllOf) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *NetworkElementSummaryAllOf) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetModel

`func (o *NetworkElementSummaryAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NetworkElementSummaryAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NetworkElementSummaryAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NetworkElementSummaryAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *NetworkElementSummaryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkElementSummaryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkElementSummaryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkElementSummaryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumEtherPorts

`func (o *NetworkElementSummaryAllOf) GetNumEtherPorts() int64`

GetNumEtherPorts returns the NumEtherPorts field if non-nil, zero value otherwise.

### GetNumEtherPortsOk

`func (o *NetworkElementSummaryAllOf) GetNumEtherPortsOk() (*int64, bool)`

GetNumEtherPortsOk returns a tuple with the NumEtherPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPorts

`func (o *NetworkElementSummaryAllOf) SetNumEtherPorts(v int64)`

SetNumEtherPorts sets NumEtherPorts field to given value.

### HasNumEtherPorts

`func (o *NetworkElementSummaryAllOf) HasNumEtherPorts() bool`

HasNumEtherPorts returns a boolean if a field has been set.

### GetNumEtherPortsConfigured

`func (o *NetworkElementSummaryAllOf) GetNumEtherPortsConfigured() int64`

GetNumEtherPortsConfigured returns the NumEtherPortsConfigured field if non-nil, zero value otherwise.

### GetNumEtherPortsConfiguredOk

`func (o *NetworkElementSummaryAllOf) GetNumEtherPortsConfiguredOk() (*int64, bool)`

GetNumEtherPortsConfiguredOk returns a tuple with the NumEtherPortsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPortsConfigured

`func (o *NetworkElementSummaryAllOf) SetNumEtherPortsConfigured(v int64)`

SetNumEtherPortsConfigured sets NumEtherPortsConfigured field to given value.

### HasNumEtherPortsConfigured

`func (o *NetworkElementSummaryAllOf) HasNumEtherPortsConfigured() bool`

HasNumEtherPortsConfigured returns a boolean if a field has been set.

### GetNumEtherPortsLinkUp

`func (o *NetworkElementSummaryAllOf) GetNumEtherPortsLinkUp() int64`

GetNumEtherPortsLinkUp returns the NumEtherPortsLinkUp field if non-nil, zero value otherwise.

### GetNumEtherPortsLinkUpOk

`func (o *NetworkElementSummaryAllOf) GetNumEtherPortsLinkUpOk() (*int64, bool)`

GetNumEtherPortsLinkUpOk returns a tuple with the NumEtherPortsLinkUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEtherPortsLinkUp

`func (o *NetworkElementSummaryAllOf) SetNumEtherPortsLinkUp(v int64)`

SetNumEtherPortsLinkUp sets NumEtherPortsLinkUp field to given value.

### HasNumEtherPortsLinkUp

`func (o *NetworkElementSummaryAllOf) HasNumEtherPortsLinkUp() bool`

HasNumEtherPortsLinkUp returns a boolean if a field has been set.

### GetNumExpansionModules

`func (o *NetworkElementSummaryAllOf) GetNumExpansionModules() int64`

GetNumExpansionModules returns the NumExpansionModules field if non-nil, zero value otherwise.

### GetNumExpansionModulesOk

`func (o *NetworkElementSummaryAllOf) GetNumExpansionModulesOk() (*int64, bool)`

GetNumExpansionModulesOk returns a tuple with the NumExpansionModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExpansionModules

`func (o *NetworkElementSummaryAllOf) SetNumExpansionModules(v int64)`

SetNumExpansionModules sets NumExpansionModules field to given value.

### HasNumExpansionModules

`func (o *NetworkElementSummaryAllOf) HasNumExpansionModules() bool`

HasNumExpansionModules returns a boolean if a field has been set.

### GetNumFcPorts

`func (o *NetworkElementSummaryAllOf) GetNumFcPorts() int64`

GetNumFcPorts returns the NumFcPorts field if non-nil, zero value otherwise.

### GetNumFcPortsOk

`func (o *NetworkElementSummaryAllOf) GetNumFcPortsOk() (*int64, bool)`

GetNumFcPortsOk returns a tuple with the NumFcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPorts

`func (o *NetworkElementSummaryAllOf) SetNumFcPorts(v int64)`

SetNumFcPorts sets NumFcPorts field to given value.

### HasNumFcPorts

`func (o *NetworkElementSummaryAllOf) HasNumFcPorts() bool`

HasNumFcPorts returns a boolean if a field has been set.

### GetNumFcPortsConfigured

`func (o *NetworkElementSummaryAllOf) GetNumFcPortsConfigured() int64`

GetNumFcPortsConfigured returns the NumFcPortsConfigured field if non-nil, zero value otherwise.

### GetNumFcPortsConfiguredOk

`func (o *NetworkElementSummaryAllOf) GetNumFcPortsConfiguredOk() (*int64, bool)`

GetNumFcPortsConfiguredOk returns a tuple with the NumFcPortsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPortsConfigured

`func (o *NetworkElementSummaryAllOf) SetNumFcPortsConfigured(v int64)`

SetNumFcPortsConfigured sets NumFcPortsConfigured field to given value.

### HasNumFcPortsConfigured

`func (o *NetworkElementSummaryAllOf) HasNumFcPortsConfigured() bool`

HasNumFcPortsConfigured returns a boolean if a field has been set.

### GetNumFcPortsLinkUp

`func (o *NetworkElementSummaryAllOf) GetNumFcPortsLinkUp() int64`

GetNumFcPortsLinkUp returns the NumFcPortsLinkUp field if non-nil, zero value otherwise.

### GetNumFcPortsLinkUpOk

`func (o *NetworkElementSummaryAllOf) GetNumFcPortsLinkUpOk() (*int64, bool)`

GetNumFcPortsLinkUpOk returns a tuple with the NumFcPortsLinkUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcPortsLinkUp

`func (o *NetworkElementSummaryAllOf) SetNumFcPortsLinkUp(v int64)`

SetNumFcPortsLinkUp sets NumFcPortsLinkUp field to given value.

### HasNumFcPortsLinkUp

`func (o *NetworkElementSummaryAllOf) HasNumFcPortsLinkUp() bool`

HasNumFcPortsLinkUp returns a boolean if a field has been set.

### GetOperEvacState

`func (o *NetworkElementSummaryAllOf) GetOperEvacState() string`

GetOperEvacState returns the OperEvacState field if non-nil, zero value otherwise.

### GetOperEvacStateOk

`func (o *NetworkElementSummaryAllOf) GetOperEvacStateOk() (*string, bool)`

GetOperEvacStateOk returns a tuple with the OperEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperEvacState

`func (o *NetworkElementSummaryAllOf) SetOperEvacState(v string)`

SetOperEvacState sets OperEvacState field to given value.

### HasOperEvacState

`func (o *NetworkElementSummaryAllOf) HasOperEvacState() bool`

HasOperEvacState returns a boolean if a field has been set.

### GetOperability

`func (o *NetworkElementSummaryAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *NetworkElementSummaryAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *NetworkElementSummaryAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *NetworkElementSummaryAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetOutOfBandIpMask

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpMask() string`

GetOutOfBandIpMask returns the OutOfBandIpMask field if non-nil, zero value otherwise.

### GetOutOfBandIpMaskOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpMaskOk() (*string, bool)`

GetOutOfBandIpMaskOk returns a tuple with the OutOfBandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpMask

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpMask(v string)`

SetOutOfBandIpMask sets OutOfBandIpMask field to given value.

### HasOutOfBandIpMask

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpMask() bool`

HasOutOfBandIpMask returns a boolean if a field has been set.

### GetOutOfBandIpv4Address

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4Address() string`

GetOutOfBandIpv4Address returns the OutOfBandIpv4Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv4AddressOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4AddressOk() (*string, bool)`

GetOutOfBandIpv4AddressOk returns a tuple with the OutOfBandIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Address

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv4Address(v string)`

SetOutOfBandIpv4Address sets OutOfBandIpv4Address field to given value.

### HasOutOfBandIpv4Address

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv4Address() bool`

HasOutOfBandIpv4Address returns a boolean if a field has been set.

### GetOutOfBandIpv4Gateway

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4Gateway() string`

GetOutOfBandIpv4Gateway returns the OutOfBandIpv4Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv4GatewayOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4GatewayOk() (*string, bool)`

GetOutOfBandIpv4GatewayOk returns a tuple with the OutOfBandIpv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Gateway

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv4Gateway(v string)`

SetOutOfBandIpv4Gateway sets OutOfBandIpv4Gateway field to given value.

### HasOutOfBandIpv4Gateway

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv4Gateway() bool`

HasOutOfBandIpv4Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv4Mask

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4Mask() string`

GetOutOfBandIpv4Mask returns the OutOfBandIpv4Mask field if non-nil, zero value otherwise.

### GetOutOfBandIpv4MaskOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv4MaskOk() (*string, bool)`

GetOutOfBandIpv4MaskOk returns a tuple with the OutOfBandIpv4Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv4Mask

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv4Mask(v string)`

SetOutOfBandIpv4Mask sets OutOfBandIpv4Mask field to given value.

### HasOutOfBandIpv4Mask

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv4Mask() bool`

HasOutOfBandIpv4Mask returns a boolean if a field has been set.

### GetOutOfBandIpv6Address

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6Address() string`

GetOutOfBandIpv6Address returns the OutOfBandIpv6Address field if non-nil, zero value otherwise.

### GetOutOfBandIpv6AddressOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6AddressOk() (*string, bool)`

GetOutOfBandIpv6AddressOk returns a tuple with the OutOfBandIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Address

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv6Address(v string)`

SetOutOfBandIpv6Address sets OutOfBandIpv6Address field to given value.

### HasOutOfBandIpv6Address

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv6Address() bool`

HasOutOfBandIpv6Address returns a boolean if a field has been set.

### GetOutOfBandIpv6Gateway

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6Gateway() string`

GetOutOfBandIpv6Gateway returns the OutOfBandIpv6Gateway field if non-nil, zero value otherwise.

### GetOutOfBandIpv6GatewayOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6GatewayOk() (*string, bool)`

GetOutOfBandIpv6GatewayOk returns a tuple with the OutOfBandIpv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Gateway

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv6Gateway(v string)`

SetOutOfBandIpv6Gateway sets OutOfBandIpv6Gateway field to given value.

### HasOutOfBandIpv6Gateway

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv6Gateway() bool`

HasOutOfBandIpv6Gateway returns a boolean if a field has been set.

### GetOutOfBandIpv6Prefix

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6Prefix() string`

GetOutOfBandIpv6Prefix returns the OutOfBandIpv6Prefix field if non-nil, zero value otherwise.

### GetOutOfBandIpv6PrefixOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandIpv6PrefixOk() (*string, bool)`

GetOutOfBandIpv6PrefixOk returns a tuple with the OutOfBandIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpv6Prefix

`func (o *NetworkElementSummaryAllOf) SetOutOfBandIpv6Prefix(v string)`

SetOutOfBandIpv6Prefix sets OutOfBandIpv6Prefix field to given value.

### HasOutOfBandIpv6Prefix

`func (o *NetworkElementSummaryAllOf) HasOutOfBandIpv6Prefix() bool`

HasOutOfBandIpv6Prefix returns a boolean if a field has been set.

### GetOutOfBandMac

`func (o *NetworkElementSummaryAllOf) GetOutOfBandMac() string`

GetOutOfBandMac returns the OutOfBandMac field if non-nil, zero value otherwise.

### GetOutOfBandMacOk

`func (o *NetworkElementSummaryAllOf) GetOutOfBandMacOk() (*string, bool)`

GetOutOfBandMacOk returns a tuple with the OutOfBandMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandMac

`func (o *NetworkElementSummaryAllOf) SetOutOfBandMac(v string)`

SetOutOfBandMac sets OutOfBandMac field to given value.

### HasOutOfBandMac

`func (o *NetworkElementSummaryAllOf) HasOutOfBandMac() bool`

HasOutOfBandMac returns a boolean if a field has been set.

### GetRevision

`func (o *NetworkElementSummaryAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkElementSummaryAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkElementSummaryAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NetworkElementSummaryAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRn

`func (o *NetworkElementSummaryAllOf) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *NetworkElementSummaryAllOf) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *NetworkElementSummaryAllOf) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *NetworkElementSummaryAllOf) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetSerial

`func (o *NetworkElementSummaryAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworkElementSummaryAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworkElementSummaryAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NetworkElementSummaryAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *NetworkElementSummaryAllOf) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *NetworkElementSummaryAllOf) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *NetworkElementSummaryAllOf) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *NetworkElementSummaryAllOf) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetSwitchId

`func (o *NetworkElementSummaryAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkElementSummaryAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkElementSummaryAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *NetworkElementSummaryAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetTotalMemory

`func (o *NetworkElementSummaryAllOf) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *NetworkElementSummaryAllOf) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *NetworkElementSummaryAllOf) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *NetworkElementSummaryAllOf) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetVendor

`func (o *NetworkElementSummaryAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *NetworkElementSummaryAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *NetworkElementSummaryAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *NetworkElementSummaryAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkElementSummaryAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkElementSummaryAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkElementSummaryAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkElementSummaryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkElementSummaryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkElementSummaryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkElementSummaryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkElementSummaryAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


