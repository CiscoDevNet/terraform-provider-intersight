# AdapterHostEthInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.HostEthInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.HostEthInterface"]
**ActiveOperState** | Pointer to **string** | The operational state of the Active vNIC. vNIC operational state information is updated by events from the adapter. This operational state is applicable to primary vNIC. If a host is powered off, this property might not be accurate as we may or may not receive events from the adapter. For Intersight Managed Domains Mode domains (IMM), the vNIC&#39;s peer object Vethernet will have the current operational state of the connection when a host is powered off. | [optional] [readonly] 
**ActiveVethOperState** | Pointer to **string** | The operational state of the Active Vethernet peer of a vNIC in Intersight Managed Mode. This state is updated by events from Fabric Interconnect or by periodic updates from Fabric Interconnect. When a Fabric Interconnect is not connected to Intersight or if the Fabric Interconnect is powered down, this property will not be updated. This state is not applicable for standalone servers. * &#x60;unknown&#x60; - The operational state of the Vethernet is not known. * &#x60;adminDown&#x60; - The operational state of the Vethernet is admin down. * &#x60;up&#x60; - The operational state of the Vethernet is Up. * &#x60;down&#x60; - The operational state of the Vethernet is Down. * &#x60;noLicense&#x60; - The operational state of the Vethernet is no license. * &#x60;linkUp&#x60; - The operational state of the Vethernet is link up. * &#x60;hardwareFailure&#x60; - The operational state of the Vethernet is hardware failure. * &#x60;softwareFailure&#x60; - The operational state of the Vethernet is software failure. * &#x60;errorDisabled&#x60; - The operational state of the Vethernet is error disabled. * &#x60;linkDown&#x60; - The operational state of the Vethernet is link down. * &#x60;sfpNotPresent&#x60; - The operational state of the Vethernet is SFP not present. * &#x60;udldAggrDown&#x60; - The operational state of the Vethernet is UDLD aggregate down. | [optional] [readonly] [default to "unknown"]
**AdminState** | Pointer to **string** | Admin state of the Host Ethernet Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Ethernet Interface. | [optional] [readonly] 
**HostEthInterfaceId** | Pointer to **int64** | Unique Identifier for an Host Ethernet Interface within the adapter object. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Type of External Ethernet Interface. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Ethernet Interface. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**Operability** | Pointer to **string** | Operability status of Host Ethernet Channel Interface. | [optional] [readonly] 
**OriginalMacAddress** | Pointer to **string** | The factory default Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The PCI address of the Host Ethernet Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | The distinguished name of the peer endpoint connected to the Host Ethernet interface. | [optional] [readonly] 
**PinGroupName** | Pointer to **string** | Name given for Lan PinGroup. | [optional] 
**QinqEnabled** | Pointer to **bool** | Setting qinqEnabled to true if we have QinQ tagging enabled on the vNIC. | [optional] 
**QinqVlan** | Pointer to **int64** | The VLAN ID for VIC QinQ (802.1Q) Tunneling. | [optional] [default to 2]
**StandbyOperState** | Pointer to **string** | The operational state of the standby vNIC. vNIC operational state information is updated by events from the adapter. This operational state is applicable only to failover vNIC. If a host is powered off, this property might not be accurate as we may or may not receive events from the adapter. For Intersight Managed Mode domains (IMM), the vNIC&#39;s peer object Vethernet will have the current operational state of the connection when a host is powered off. | [optional] [readonly] 
**StandbyVethOperState** | Pointer to **string** | The operational state of the Standby Vethernet peer of a failover vNIC in Intersight Managed Mode. This state is updated by events from Fabric Interconnect or by periodic updates from Fabric Interconnect. When a Fabric Interconnect is not connected to Intersight or if the Fabric Interconnect is powered down, this property will not be updated. This state is not applicable for standalone servers. * &#x60;unknown&#x60; - The operational state of the Vethernet is not known. * &#x60;adminDown&#x60; - The operational state of the Vethernet is admin down. * &#x60;up&#x60; - The operational state of the Vethernet is Up. * &#x60;down&#x60; - The operational state of the Vethernet is Down. * &#x60;noLicense&#x60; - The operational state of the Vethernet is no license. * &#x60;linkUp&#x60; - The operational state of the Vethernet is link up. * &#x60;hardwareFailure&#x60; - The operational state of the Vethernet is hardware failure. * &#x60;softwareFailure&#x60; - The operational state of the Vethernet is software failure. * &#x60;errorDisabled&#x60; - The operational state of the Vethernet is error disabled. * &#x60;linkDown&#x60; - The operational state of the Vethernet is link down. * &#x60;sfpNotPresent&#x60; - The operational state of the Vethernet is SFP not present. * &#x60;udldAggrDown&#x60; - The operational state of the Vethernet is UDLD aggregate down. | [optional] [readonly] [default to "unknown"]
**StandbyVifId** | Pointer to **int64** | Identifier of the Standby virtual ethernet interface (Vethernet) on the networking component (e.g., Fabric Interconnect) for the corresponding Host Ethernet Interface (vNIC). | [optional] [readonly] 
**VethAction** | Pointer to **string** | The action to be performed on the vethernet corresponding to the vNIC. * &#x60;None&#x60; - Default value for vif operation. * &#x60;ResetConnectivity&#x60; - Resets connectivity on both active and passive vif. * &#x60;ResetConnectivityActive&#x60; - Resets connectivity on the active vif. * &#x60;ResetConnectivityPassive&#x60; - Resets connectivity on the passive vif. * &#x60;Enable&#x60; - Enables the vif on both the FIs. * &#x60;Disable&#x60; - Disables the vif on both the FIs. * &#x60;EnableActive&#x60; - Enables the corresponding active vif. * &#x60;EnablePassive&#x60; - Enables the corresponding standby vif. * &#x60;DisableActive&#x60; - Disables the corresponding active vif. * &#x60;DisablePassive&#x60; - Disables the corresponding standby vif. | [optional] [default to "None"]
**VifId** | Pointer to **int64** | Identifier of the virtual ethernet interface (Vethernet) on the networking component (e.g., Fabric Interconnect) for the corresponding Host Ethernet Interface (vNIC). | [optional] [readonly] 
**VirtualizationPreference** | Pointer to **string** | Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not. | [optional] [readonly] 
**VnicDn** | Pointer to **string** | The Virtual Ethernet Interface DN connected to the Host Ethernet Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**NullableAdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PinnedInterface** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StandbyVethernet** | Pointer to [**NullableNetworkVethernetRelationship**](NetworkVethernetRelationship.md) |  | [optional] 
**Vethernet** | Pointer to [**NullableNetworkVethernetRelationship**](NetworkVethernetRelationship.md) |  | [optional] 

## Methods

### NewAdapterHostEthInterface

`func NewAdapterHostEthInterface(classId string, objectType string, ) *AdapterHostEthInterface`

NewAdapterHostEthInterface instantiates a new AdapterHostEthInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostEthInterfaceWithDefaults

`func NewAdapterHostEthInterfaceWithDefaults() *AdapterHostEthInterface`

NewAdapterHostEthInterfaceWithDefaults instantiates a new AdapterHostEthInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterHostEthInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostEthInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostEthInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterHostEthInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostEthInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostEthInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveOperState

`func (o *AdapterHostEthInterface) GetActiveOperState() string`

GetActiveOperState returns the ActiveOperState field if non-nil, zero value otherwise.

### GetActiveOperStateOk

`func (o *AdapterHostEthInterface) GetActiveOperStateOk() (*string, bool)`

GetActiveOperStateOk returns a tuple with the ActiveOperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOperState

`func (o *AdapterHostEthInterface) SetActiveOperState(v string)`

SetActiveOperState sets ActiveOperState field to given value.

### HasActiveOperState

`func (o *AdapterHostEthInterface) HasActiveOperState() bool`

HasActiveOperState returns a boolean if a field has been set.

### GetActiveVethOperState

`func (o *AdapterHostEthInterface) GetActiveVethOperState() string`

GetActiveVethOperState returns the ActiveVethOperState field if non-nil, zero value otherwise.

### GetActiveVethOperStateOk

`func (o *AdapterHostEthInterface) GetActiveVethOperStateOk() (*string, bool)`

GetActiveVethOperStateOk returns a tuple with the ActiveVethOperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVethOperState

`func (o *AdapterHostEthInterface) SetActiveVethOperState(v string)`

SetActiveVethOperState sets ActiveVethOperState field to given value.

### HasActiveVethOperState

`func (o *AdapterHostEthInterface) HasActiveVethOperState() bool`

HasActiveVethOperState returns a boolean if a field has been set.

### GetAdminState

`func (o *AdapterHostEthInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostEthInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostEthInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostEthInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostEthInterface) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostEthInterface) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostEthInterface) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostEthInterface) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostEthInterfaceId

`func (o *AdapterHostEthInterface) GetHostEthInterfaceId() int64`

GetHostEthInterfaceId returns the HostEthInterfaceId field if non-nil, zero value otherwise.

### GetHostEthInterfaceIdOk

`func (o *AdapterHostEthInterface) GetHostEthInterfaceIdOk() (*int64, bool)`

GetHostEthInterfaceIdOk returns a tuple with the HostEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthInterfaceId

`func (o *AdapterHostEthInterface) SetHostEthInterfaceId(v int64)`

SetHostEthInterfaceId sets HostEthInterfaceId field to given value.

### HasHostEthInterfaceId

`func (o *AdapterHostEthInterface) HasHostEthInterfaceId() bool`

HasHostEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterHostEthInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterHostEthInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterHostEthInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterHostEthInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterHostEthInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterHostEthInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterHostEthInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterHostEthInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostEthInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostEthInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostEthInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostEthInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *AdapterHostEthInterface) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *AdapterHostEthInterface) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *AdapterHostEthInterface) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *AdapterHostEthInterface) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *AdapterHostEthInterface) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *AdapterHostEthInterface) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperability

`func (o *AdapterHostEthInterface) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostEthInterface) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostEthInterface) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostEthInterface) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalMacAddress

`func (o *AdapterHostEthInterface) GetOriginalMacAddress() string`

GetOriginalMacAddress returns the OriginalMacAddress field if non-nil, zero value otherwise.

### GetOriginalMacAddressOk

`func (o *AdapterHostEthInterface) GetOriginalMacAddressOk() (*string, bool)`

GetOriginalMacAddressOk returns a tuple with the OriginalMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMacAddress

`func (o *AdapterHostEthInterface) SetOriginalMacAddress(v string)`

SetOriginalMacAddress sets OriginalMacAddress field to given value.

### HasOriginalMacAddress

`func (o *AdapterHostEthInterface) HasOriginalMacAddress() bool`

HasOriginalMacAddress returns a boolean if a field has been set.

### GetPciAddr

`func (o *AdapterHostEthInterface) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *AdapterHostEthInterface) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *AdapterHostEthInterface) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *AdapterHostEthInterface) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostEthInterface) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostEthInterface) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostEthInterface) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostEthInterface) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPinGroupName

`func (o *AdapterHostEthInterface) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *AdapterHostEthInterface) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *AdapterHostEthInterface) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *AdapterHostEthInterface) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetQinqEnabled

`func (o *AdapterHostEthInterface) GetQinqEnabled() bool`

GetQinqEnabled returns the QinqEnabled field if non-nil, zero value otherwise.

### GetQinqEnabledOk

`func (o *AdapterHostEthInterface) GetQinqEnabledOk() (*bool, bool)`

GetQinqEnabledOk returns a tuple with the QinqEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqEnabled

`func (o *AdapterHostEthInterface) SetQinqEnabled(v bool)`

SetQinqEnabled sets QinqEnabled field to given value.

### HasQinqEnabled

`func (o *AdapterHostEthInterface) HasQinqEnabled() bool`

HasQinqEnabled returns a boolean if a field has been set.

### GetQinqVlan

`func (o *AdapterHostEthInterface) GetQinqVlan() int64`

GetQinqVlan returns the QinqVlan field if non-nil, zero value otherwise.

### GetQinqVlanOk

`func (o *AdapterHostEthInterface) GetQinqVlanOk() (*int64, bool)`

GetQinqVlanOk returns a tuple with the QinqVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqVlan

`func (o *AdapterHostEthInterface) SetQinqVlan(v int64)`

SetQinqVlan sets QinqVlan field to given value.

### HasQinqVlan

`func (o *AdapterHostEthInterface) HasQinqVlan() bool`

HasQinqVlan returns a boolean if a field has been set.

### GetStandbyOperState

`func (o *AdapterHostEthInterface) GetStandbyOperState() string`

GetStandbyOperState returns the StandbyOperState field if non-nil, zero value otherwise.

### GetStandbyOperStateOk

`func (o *AdapterHostEthInterface) GetStandbyOperStateOk() (*string, bool)`

GetStandbyOperStateOk returns a tuple with the StandbyOperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyOperState

`func (o *AdapterHostEthInterface) SetStandbyOperState(v string)`

SetStandbyOperState sets StandbyOperState field to given value.

### HasStandbyOperState

`func (o *AdapterHostEthInterface) HasStandbyOperState() bool`

HasStandbyOperState returns a boolean if a field has been set.

### GetStandbyVethOperState

`func (o *AdapterHostEthInterface) GetStandbyVethOperState() string`

GetStandbyVethOperState returns the StandbyVethOperState field if non-nil, zero value otherwise.

### GetStandbyVethOperStateOk

`func (o *AdapterHostEthInterface) GetStandbyVethOperStateOk() (*string, bool)`

GetStandbyVethOperStateOk returns a tuple with the StandbyVethOperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVethOperState

`func (o *AdapterHostEthInterface) SetStandbyVethOperState(v string)`

SetStandbyVethOperState sets StandbyVethOperState field to given value.

### HasStandbyVethOperState

`func (o *AdapterHostEthInterface) HasStandbyVethOperState() bool`

HasStandbyVethOperState returns a boolean if a field has been set.

### GetStandbyVifId

`func (o *AdapterHostEthInterface) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *AdapterHostEthInterface) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *AdapterHostEthInterface) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *AdapterHostEthInterface) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetVethAction

`func (o *AdapterHostEthInterface) GetVethAction() string`

GetVethAction returns the VethAction field if non-nil, zero value otherwise.

### GetVethActionOk

`func (o *AdapterHostEthInterface) GetVethActionOk() (*string, bool)`

GetVethActionOk returns a tuple with the VethAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVethAction

`func (o *AdapterHostEthInterface) SetVethAction(v string)`

SetVethAction sets VethAction field to given value.

### HasVethAction

`func (o *AdapterHostEthInterface) HasVethAction() bool`

HasVethAction returns a boolean if a field has been set.

### GetVifId

`func (o *AdapterHostEthInterface) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *AdapterHostEthInterface) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *AdapterHostEthInterface) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *AdapterHostEthInterface) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVirtualizationPreference

`func (o *AdapterHostEthInterface) GetVirtualizationPreference() string`

GetVirtualizationPreference returns the VirtualizationPreference field if non-nil, zero value otherwise.

### GetVirtualizationPreferenceOk

`func (o *AdapterHostEthInterface) GetVirtualizationPreferenceOk() (*string, bool)`

GetVirtualizationPreferenceOk returns a tuple with the VirtualizationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationPreference

`func (o *AdapterHostEthInterface) SetVirtualizationPreference(v string)`

SetVirtualizationPreference sets VirtualizationPreference field to given value.

### HasVirtualizationPreference

`func (o *AdapterHostEthInterface) HasVirtualizationPreference() bool`

HasVirtualizationPreference returns a boolean if a field has been set.

### GetVnicDn

`func (o *AdapterHostEthInterface) GetVnicDn() string`

GetVnicDn returns the VnicDn field if non-nil, zero value otherwise.

### GetVnicDnOk

`func (o *AdapterHostEthInterface) GetVnicDnOk() (*string, bool)`

GetVnicDnOk returns a tuple with the VnicDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDn

`func (o *AdapterHostEthInterface) SetVnicDn(v string)`

SetVnicDn sets VnicDn field to given value.

### HasVnicDn

`func (o *AdapterHostEthInterface) HasVnicDn() bool`

HasVnicDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostEthInterface) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostEthInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostEthInterface) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostEthInterface) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### SetAdapterUnitNil

`func (o *AdapterHostEthInterface) SetAdapterUnitNil(b bool)`

 SetAdapterUnitNil sets the value for AdapterUnit to be an explicit nil

### UnsetAdapterUnit
`func (o *AdapterHostEthInterface) UnsetAdapterUnit()`

UnsetAdapterUnit ensures that no value is present for AdapterUnit, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *AdapterHostEthInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostEthInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostEthInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostEthInterface) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *AdapterHostEthInterface) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *AdapterHostEthInterface) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPinnedInterface

`func (o *AdapterHostEthInterface) GetPinnedInterface() InventoryInterfaceRelationship`

GetPinnedInterface returns the PinnedInterface field if non-nil, zero value otherwise.

### GetPinnedInterfaceOk

`func (o *AdapterHostEthInterface) GetPinnedInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetPinnedInterfaceOk returns a tuple with the PinnedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterface

`func (o *AdapterHostEthInterface) SetPinnedInterface(v InventoryInterfaceRelationship)`

SetPinnedInterface sets PinnedInterface field to given value.

### HasPinnedInterface

`func (o *AdapterHostEthInterface) HasPinnedInterface() bool`

HasPinnedInterface returns a boolean if a field has been set.

### SetPinnedInterfaceNil

`func (o *AdapterHostEthInterface) SetPinnedInterfaceNil(b bool)`

 SetPinnedInterfaceNil sets the value for PinnedInterface to be an explicit nil

### UnsetPinnedInterface
`func (o *AdapterHostEthInterface) UnsetPinnedInterface()`

UnsetPinnedInterface ensures that no value is present for PinnedInterface, not even an explicit nil
### GetRegisteredDevice

`func (o *AdapterHostEthInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostEthInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostEthInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostEthInterface) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *AdapterHostEthInterface) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *AdapterHostEthInterface) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetStandbyVethernet

`func (o *AdapterHostEthInterface) GetStandbyVethernet() NetworkVethernetRelationship`

GetStandbyVethernet returns the StandbyVethernet field if non-nil, zero value otherwise.

### GetStandbyVethernetOk

`func (o *AdapterHostEthInterface) GetStandbyVethernetOk() (*NetworkVethernetRelationship, bool)`

GetStandbyVethernetOk returns a tuple with the StandbyVethernet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVethernet

`func (o *AdapterHostEthInterface) SetStandbyVethernet(v NetworkVethernetRelationship)`

SetStandbyVethernet sets StandbyVethernet field to given value.

### HasStandbyVethernet

`func (o *AdapterHostEthInterface) HasStandbyVethernet() bool`

HasStandbyVethernet returns a boolean if a field has been set.

### SetStandbyVethernetNil

`func (o *AdapterHostEthInterface) SetStandbyVethernetNil(b bool)`

 SetStandbyVethernetNil sets the value for StandbyVethernet to be an explicit nil

### UnsetStandbyVethernet
`func (o *AdapterHostEthInterface) UnsetStandbyVethernet()`

UnsetStandbyVethernet ensures that no value is present for StandbyVethernet, not even an explicit nil
### GetVethernet

`func (o *AdapterHostEthInterface) GetVethernet() NetworkVethernetRelationship`

GetVethernet returns the Vethernet field if non-nil, zero value otherwise.

### GetVethernetOk

`func (o *AdapterHostEthInterface) GetVethernetOk() (*NetworkVethernetRelationship, bool)`

GetVethernetOk returns a tuple with the Vethernet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVethernet

`func (o *AdapterHostEthInterface) SetVethernet(v NetworkVethernetRelationship)`

SetVethernet sets Vethernet field to given value.

### HasVethernet

`func (o *AdapterHostEthInterface) HasVethernet() bool`

HasVethernet returns a boolean if a field has been set.

### SetVethernetNil

`func (o *AdapterHostEthInterface) SetVethernetNil(b bool)`

 SetVethernetNil sets the value for Vethernet to be an explicit nil

### UnsetVethernet
`func (o *AdapterHostEthInterface) UnsetVethernet()`

UnsetVethernet ensures that no value is present for Vethernet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


