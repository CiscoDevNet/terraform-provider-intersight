# VnicEthIfInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthIfInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthIfInventory"]
**Cdn** | Pointer to [**NullableVnicCdn**](VnicCdn.md) |  | [optional] 
**FailoverEnabled** | Pointer to **bool** | Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. | [optional] [readonly] [default to false]
**IscsiIpV4AddressAllocationType** | Pointer to **string** | Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type. * &#x60;None&#x60; - Type indicates that there is no IP associated to an vnic. * &#x60;DHCP&#x60; - The IP address is assigned using DHCP, if available. * &#x60;Static&#x60; - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * &#x60;Pool&#x60; - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. | [optional] [readonly] [default to "None"]
**IscsiIpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IscsiIpv4Address** | Pointer to **string** | IP address associated to the vNIC. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. | [optional] [readonly] 
**MacAddressType** | Pointer to **string** | Type of allocation selected to assign a MAC address for the vnic. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [readonly] [default to "POOL"]
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] [readonly] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two. | [optional] [readonly] 
**PinGroupName** | Pointer to **string** | Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. | [optional] [readonly] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**SriovSettings** | Pointer to [**NullableVnicSriovSettings**](VnicSriovSettings.md) |  | [optional] 
**StandbyVifId** | Pointer to **int64** | The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. | [optional] [readonly] 
**StaticMacAddress** | Pointer to **string** | The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx. To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the following MAC prefix 00:25:B5:xx:xx:xx. | [optional] [readonly] 
**UsnicSettings** | Pointer to [**NullableVnicUsnicSettings**](VnicUsnicSettings.md) |  | [optional] 
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**VmqSettings** | Pointer to [**NullableVnicVmqSettings**](VnicVmqSettings.md) |  | [optional] 
**EthAdapterPolicy** | Pointer to [**VnicEthAdapterPolicyInventoryRelationship**](VnicEthAdapterPolicyInventoryRelationship.md) |  | [optional] 
**EthNetworkPolicy** | Pointer to [**VnicEthNetworkPolicyInventoryRelationship**](VnicEthNetworkPolicyInventoryRelationship.md) |  | [optional] 
**EthQosPolicy** | Pointer to [**VnicEthQosPolicyInventoryRelationship**](VnicEthQosPolicyInventoryRelationship.md) |  | [optional] 
**FabricEthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyInventoryRelationship**](FabricEthNetworkControlPolicyInventoryRelationship.md) |  | [optional] 
**FabricEthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyInventoryRelationship**](FabricEthNetworkGroupPolicyInventoryRelationship.md) | An array of relationships to fabricEthNetworkGroupPolicyInventory resources. | [optional] [readonly] 
**IpLease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**IscsiBootPolicy** | Pointer to [**VnicIscsiBootPolicyInventoryRelationship**](VnicIscsiBootPolicyInventoryRelationship.md) |  | [optional] 
**LanConnectivityPolicy** | Pointer to [**VnicLanConnectivityPolicyInventoryRelationship**](VnicLanConnectivityPolicyInventoryRelationship.md) |  | [optional] 
**LcpVnic** | Pointer to [**VnicEthIfInventoryRelationship**](VnicEthIfInventoryRelationship.md) |  | [optional] 
**MacLease** | Pointer to [**MacpoolLeaseRelationship**](MacpoolLeaseRelationship.md) |  | [optional] 
**MacPool** | Pointer to [**MacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 
**SpVnics** | Pointer to [**[]VnicEthIfInventoryRelationship**](VnicEthIfInventoryRelationship.md) | An array of relationships to vnicEthIfInventory resources. | [optional] [readonly] 

## Methods

### NewVnicEthIfInventoryAllOf

`func NewVnicEthIfInventoryAllOf(classId string, objectType string, ) *VnicEthIfInventoryAllOf`

NewVnicEthIfInventoryAllOf instantiates a new VnicEthIfInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfInventoryAllOfWithDefaults

`func NewVnicEthIfInventoryAllOfWithDefaults() *VnicEthIfInventoryAllOf`

NewVnicEthIfInventoryAllOfWithDefaults instantiates a new VnicEthIfInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthIfInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthIfInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthIfInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthIfInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthIfInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthIfInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdn

`func (o *VnicEthIfInventoryAllOf) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicEthIfInventoryAllOf) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicEthIfInventoryAllOf) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicEthIfInventoryAllOf) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *VnicEthIfInventoryAllOf) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *VnicEthIfInventoryAllOf) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetFailoverEnabled

`func (o *VnicEthIfInventoryAllOf) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicEthIfInventoryAllOf) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicEthIfInventoryAllOf) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicEthIfInventoryAllOf) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpV4AddressAllocationType() string`

GetIscsiIpV4AddressAllocationType returns the IscsiIpV4AddressAllocationType field if non-nil, zero value otherwise.

### GetIscsiIpV4AddressAllocationTypeOk

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpV4AddressAllocationTypeOk() (*string, bool)`

GetIscsiIpV4AddressAllocationTypeOk returns a tuple with the IscsiIpV4AddressAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventoryAllOf) SetIscsiIpV4AddressAllocationType(v string)`

SetIscsiIpV4AddressAllocationType sets IscsiIpV4AddressAllocationType field to given value.

### HasIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventoryAllOf) HasIscsiIpV4AddressAllocationType() bool`

HasIscsiIpV4AddressAllocationType returns a boolean if a field has been set.

### GetIscsiIpV4Config

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpV4Config() IppoolIpV4Config`

GetIscsiIpV4Config returns the IscsiIpV4Config field if non-nil, zero value otherwise.

### GetIscsiIpV4ConfigOk

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIscsiIpV4ConfigOk returns a tuple with the IscsiIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4Config

`func (o *VnicEthIfInventoryAllOf) SetIscsiIpV4Config(v IppoolIpV4Config)`

SetIscsiIpV4Config sets IscsiIpV4Config field to given value.

### HasIscsiIpV4Config

`func (o *VnicEthIfInventoryAllOf) HasIscsiIpV4Config() bool`

HasIscsiIpV4Config returns a boolean if a field has been set.

### SetIscsiIpV4ConfigNil

`func (o *VnicEthIfInventoryAllOf) SetIscsiIpV4ConfigNil(b bool)`

 SetIscsiIpV4ConfigNil sets the value for IscsiIpV4Config to be an explicit nil

### UnsetIscsiIpV4Config
`func (o *VnicEthIfInventoryAllOf) UnsetIscsiIpV4Config()`

UnsetIscsiIpV4Config ensures that no value is present for IscsiIpV4Config, not even an explicit nil
### GetIscsiIpv4Address

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpv4Address() string`

GetIscsiIpv4Address returns the IscsiIpv4Address field if non-nil, zero value otherwise.

### GetIscsiIpv4AddressOk

`func (o *VnicEthIfInventoryAllOf) GetIscsiIpv4AddressOk() (*string, bool)`

GetIscsiIpv4AddressOk returns a tuple with the IscsiIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpv4Address

`func (o *VnicEthIfInventoryAllOf) SetIscsiIpv4Address(v string)`

SetIscsiIpv4Address sets IscsiIpv4Address field to given value.

### HasIscsiIpv4Address

`func (o *VnicEthIfInventoryAllOf) HasIscsiIpv4Address() bool`

HasIscsiIpv4Address returns a boolean if a field has been set.

### GetMacAddress

`func (o *VnicEthIfInventoryAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VnicEthIfInventoryAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VnicEthIfInventoryAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VnicEthIfInventoryAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacAddressType

`func (o *VnicEthIfInventoryAllOf) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *VnicEthIfInventoryAllOf) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *VnicEthIfInventoryAllOf) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *VnicEthIfInventoryAllOf) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

### GetName

`func (o *VnicEthIfInventoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthIfInventoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthIfInventoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthIfInventoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicEthIfInventoryAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicEthIfInventoryAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicEthIfInventoryAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicEthIfInventoryAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicEthIfInventoryAllOf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicEthIfInventoryAllOf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicEthIfInventoryAllOf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicEthIfInventoryAllOf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicEthIfInventoryAllOf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicEthIfInventoryAllOf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicEthIfInventoryAllOf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicEthIfInventoryAllOf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicEthIfInventoryAllOf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicEthIfInventoryAllOf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetSriovSettings

`func (o *VnicEthIfInventoryAllOf) GetSriovSettings() VnicSriovSettings`

GetSriovSettings returns the SriovSettings field if non-nil, zero value otherwise.

### GetSriovSettingsOk

`func (o *VnicEthIfInventoryAllOf) GetSriovSettingsOk() (*VnicSriovSettings, bool)`

GetSriovSettingsOk returns a tuple with the SriovSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSriovSettings

`func (o *VnicEthIfInventoryAllOf) SetSriovSettings(v VnicSriovSettings)`

SetSriovSettings sets SriovSettings field to given value.

### HasSriovSettings

`func (o *VnicEthIfInventoryAllOf) HasSriovSettings() bool`

HasSriovSettings returns a boolean if a field has been set.

### SetSriovSettingsNil

`func (o *VnicEthIfInventoryAllOf) SetSriovSettingsNil(b bool)`

 SetSriovSettingsNil sets the value for SriovSettings to be an explicit nil

### UnsetSriovSettings
`func (o *VnicEthIfInventoryAllOf) UnsetSriovSettings()`

UnsetSriovSettings ensures that no value is present for SriovSettings, not even an explicit nil
### GetStandbyVifId

`func (o *VnicEthIfInventoryAllOf) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *VnicEthIfInventoryAllOf) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *VnicEthIfInventoryAllOf) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *VnicEthIfInventoryAllOf) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetStaticMacAddress

`func (o *VnicEthIfInventoryAllOf) GetStaticMacAddress() string`

GetStaticMacAddress returns the StaticMacAddress field if non-nil, zero value otherwise.

### GetStaticMacAddressOk

`func (o *VnicEthIfInventoryAllOf) GetStaticMacAddressOk() (*string, bool)`

GetStaticMacAddressOk returns a tuple with the StaticMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddress

`func (o *VnicEthIfInventoryAllOf) SetStaticMacAddress(v string)`

SetStaticMacAddress sets StaticMacAddress field to given value.

### HasStaticMacAddress

`func (o *VnicEthIfInventoryAllOf) HasStaticMacAddress() bool`

HasStaticMacAddress returns a boolean if a field has been set.

### GetUsnicSettings

`func (o *VnicEthIfInventoryAllOf) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicEthIfInventoryAllOf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicEthIfInventoryAllOf) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicEthIfInventoryAllOf) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### SetUsnicSettingsNil

`func (o *VnicEthIfInventoryAllOf) SetUsnicSettingsNil(b bool)`

 SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil

### UnsetUsnicSettings
`func (o *VnicEthIfInventoryAllOf) UnsetUsnicSettings()`

UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
### GetVifId

`func (o *VnicEthIfInventoryAllOf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicEthIfInventoryAllOf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicEthIfInventoryAllOf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicEthIfInventoryAllOf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVmqSettings

`func (o *VnicEthIfInventoryAllOf) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicEthIfInventoryAllOf) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicEthIfInventoryAllOf) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicEthIfInventoryAllOf) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### SetVmqSettingsNil

`func (o *VnicEthIfInventoryAllOf) SetVmqSettingsNil(b bool)`

 SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil

### UnsetVmqSettings
`func (o *VnicEthIfInventoryAllOf) UnsetVmqSettings()`

UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
### GetEthAdapterPolicy

`func (o *VnicEthIfInventoryAllOf) GetEthAdapterPolicy() VnicEthAdapterPolicyInventoryRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyInventoryRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicEthIfInventoryAllOf) SetEthAdapterPolicy(v VnicEthAdapterPolicyInventoryRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicEthIfInventoryAllOf) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicEthIfInventoryAllOf) GetEthNetworkPolicy() VnicEthNetworkPolicyInventoryRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyInventoryRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicEthIfInventoryAllOf) SetEthNetworkPolicy(v VnicEthNetworkPolicyInventoryRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicEthIfInventoryAllOf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicEthIfInventoryAllOf) GetEthQosPolicy() VnicEthQosPolicyInventoryRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetEthQosPolicyOk() (*VnicEthQosPolicyInventoryRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicEthIfInventoryAllOf) SetEthQosPolicy(v VnicEthQosPolicyInventoryRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicEthIfInventoryAllOf) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventoryAllOf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyInventoryRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyInventoryRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventoryAllOf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyInventoryRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventoryAllOf) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventoryAllOf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyInventoryRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyInventoryRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventoryAllOf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyInventoryRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventoryAllOf) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicEthIfInventoryAllOf) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicEthIfInventoryAllOf) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetIpLease

`func (o *VnicEthIfInventoryAllOf) GetIpLease() IppoolIpLeaseRelationship`

GetIpLease returns the IpLease field if non-nil, zero value otherwise.

### GetIpLeaseOk

`func (o *VnicEthIfInventoryAllOf) GetIpLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpLeaseOk returns a tuple with the IpLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLease

`func (o *VnicEthIfInventoryAllOf) SetIpLease(v IppoolIpLeaseRelationship)`

SetIpLease sets IpLease field to given value.

### HasIpLease

`func (o *VnicEthIfInventoryAllOf) HasIpLease() bool`

HasIpLease returns a boolean if a field has been set.

### GetIscsiBootPolicy

`func (o *VnicEthIfInventoryAllOf) GetIscsiBootPolicy() VnicIscsiBootPolicyInventoryRelationship`

GetIscsiBootPolicy returns the IscsiBootPolicy field if non-nil, zero value otherwise.

### GetIscsiBootPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetIscsiBootPolicyOk() (*VnicIscsiBootPolicyInventoryRelationship, bool)`

GetIscsiBootPolicyOk returns a tuple with the IscsiBootPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiBootPolicy

`func (o *VnicEthIfInventoryAllOf) SetIscsiBootPolicy(v VnicIscsiBootPolicyInventoryRelationship)`

SetIscsiBootPolicy sets IscsiBootPolicy field to given value.

### HasIscsiBootPolicy

`func (o *VnicEthIfInventoryAllOf) HasIscsiBootPolicy() bool`

HasIscsiBootPolicy returns a boolean if a field has been set.

### GetLanConnectivityPolicy

`func (o *VnicEthIfInventoryAllOf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyInventoryRelationship`

GetLanConnectivityPolicy returns the LanConnectivityPolicy field if non-nil, zero value otherwise.

### GetLanConnectivityPolicyOk

`func (o *VnicEthIfInventoryAllOf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyInventoryRelationship, bool)`

GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanConnectivityPolicy

`func (o *VnicEthIfInventoryAllOf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyInventoryRelationship)`

SetLanConnectivityPolicy sets LanConnectivityPolicy field to given value.

### HasLanConnectivityPolicy

`func (o *VnicEthIfInventoryAllOf) HasLanConnectivityPolicy() bool`

HasLanConnectivityPolicy returns a boolean if a field has been set.

### GetLcpVnic

`func (o *VnicEthIfInventoryAllOf) GetLcpVnic() VnicEthIfInventoryRelationship`

GetLcpVnic returns the LcpVnic field if non-nil, zero value otherwise.

### GetLcpVnicOk

`func (o *VnicEthIfInventoryAllOf) GetLcpVnicOk() (*VnicEthIfInventoryRelationship, bool)`

GetLcpVnicOk returns a tuple with the LcpVnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpVnic

`func (o *VnicEthIfInventoryAllOf) SetLcpVnic(v VnicEthIfInventoryRelationship)`

SetLcpVnic sets LcpVnic field to given value.

### HasLcpVnic

`func (o *VnicEthIfInventoryAllOf) HasLcpVnic() bool`

HasLcpVnic returns a boolean if a field has been set.

### GetMacLease

`func (o *VnicEthIfInventoryAllOf) GetMacLease() MacpoolLeaseRelationship`

GetMacLease returns the MacLease field if non-nil, zero value otherwise.

### GetMacLeaseOk

`func (o *VnicEthIfInventoryAllOf) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool)`

GetMacLeaseOk returns a tuple with the MacLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLease

`func (o *VnicEthIfInventoryAllOf) SetMacLease(v MacpoolLeaseRelationship)`

SetMacLease sets MacLease field to given value.

### HasMacLease

`func (o *VnicEthIfInventoryAllOf) HasMacLease() bool`

HasMacLease returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicEthIfInventoryAllOf) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicEthIfInventoryAllOf) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicEthIfInventoryAllOf) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicEthIfInventoryAllOf) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetSpVnics

`func (o *VnicEthIfInventoryAllOf) GetSpVnics() []VnicEthIfInventoryRelationship`

GetSpVnics returns the SpVnics field if non-nil, zero value otherwise.

### GetSpVnicsOk

`func (o *VnicEthIfInventoryAllOf) GetSpVnicsOk() (*[]VnicEthIfInventoryRelationship, bool)`

GetSpVnicsOk returns a tuple with the SpVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVnics

`func (o *VnicEthIfInventoryAllOf) SetSpVnics(v []VnicEthIfInventoryRelationship)`

SetSpVnics sets SpVnics field to given value.

### HasSpVnics

`func (o *VnicEthIfInventoryAllOf) HasSpVnics() bool`

HasSpVnics returns a boolean if a field has been set.

### SetSpVnicsNil

`func (o *VnicEthIfInventoryAllOf) SetSpVnicsNil(b bool)`

 SetSpVnicsNil sets the value for SpVnics to be an explicit nil

### UnsetSpVnics
`func (o *VnicEthIfInventoryAllOf) UnsetSpVnics()`

UnsetSpVnics ensures that no value is present for SpVnics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


