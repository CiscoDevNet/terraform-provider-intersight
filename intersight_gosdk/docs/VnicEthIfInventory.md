# VnicEthIfInventory

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
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] [readonly] 
**PinGroupName** | Pointer to **string** | Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. | [optional] [readonly] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
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

### NewVnicEthIfInventory

`func NewVnicEthIfInventory(classId string, objectType string, ) *VnicEthIfInventory`

NewVnicEthIfInventory instantiates a new VnicEthIfInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfInventoryWithDefaults

`func NewVnicEthIfInventoryWithDefaults() *VnicEthIfInventory`

NewVnicEthIfInventoryWithDefaults instantiates a new VnicEthIfInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthIfInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthIfInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthIfInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthIfInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthIfInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthIfInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdn

`func (o *VnicEthIfInventory) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicEthIfInventory) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicEthIfInventory) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicEthIfInventory) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *VnicEthIfInventory) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *VnicEthIfInventory) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetFailoverEnabled

`func (o *VnicEthIfInventory) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicEthIfInventory) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicEthIfInventory) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicEthIfInventory) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventory) GetIscsiIpV4AddressAllocationType() string`

GetIscsiIpV4AddressAllocationType returns the IscsiIpV4AddressAllocationType field if non-nil, zero value otherwise.

### GetIscsiIpV4AddressAllocationTypeOk

`func (o *VnicEthIfInventory) GetIscsiIpV4AddressAllocationTypeOk() (*string, bool)`

GetIscsiIpV4AddressAllocationTypeOk returns a tuple with the IscsiIpV4AddressAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventory) SetIscsiIpV4AddressAllocationType(v string)`

SetIscsiIpV4AddressAllocationType sets IscsiIpV4AddressAllocationType field to given value.

### HasIscsiIpV4AddressAllocationType

`func (o *VnicEthIfInventory) HasIscsiIpV4AddressAllocationType() bool`

HasIscsiIpV4AddressAllocationType returns a boolean if a field has been set.

### GetIscsiIpV4Config

`func (o *VnicEthIfInventory) GetIscsiIpV4Config() IppoolIpV4Config`

GetIscsiIpV4Config returns the IscsiIpV4Config field if non-nil, zero value otherwise.

### GetIscsiIpV4ConfigOk

`func (o *VnicEthIfInventory) GetIscsiIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIscsiIpV4ConfigOk returns a tuple with the IscsiIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4Config

`func (o *VnicEthIfInventory) SetIscsiIpV4Config(v IppoolIpV4Config)`

SetIscsiIpV4Config sets IscsiIpV4Config field to given value.

### HasIscsiIpV4Config

`func (o *VnicEthIfInventory) HasIscsiIpV4Config() bool`

HasIscsiIpV4Config returns a boolean if a field has been set.

### SetIscsiIpV4ConfigNil

`func (o *VnicEthIfInventory) SetIscsiIpV4ConfigNil(b bool)`

 SetIscsiIpV4ConfigNil sets the value for IscsiIpV4Config to be an explicit nil

### UnsetIscsiIpV4Config
`func (o *VnicEthIfInventory) UnsetIscsiIpV4Config()`

UnsetIscsiIpV4Config ensures that no value is present for IscsiIpV4Config, not even an explicit nil
### GetIscsiIpv4Address

`func (o *VnicEthIfInventory) GetIscsiIpv4Address() string`

GetIscsiIpv4Address returns the IscsiIpv4Address field if non-nil, zero value otherwise.

### GetIscsiIpv4AddressOk

`func (o *VnicEthIfInventory) GetIscsiIpv4AddressOk() (*string, bool)`

GetIscsiIpv4AddressOk returns a tuple with the IscsiIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpv4Address

`func (o *VnicEthIfInventory) SetIscsiIpv4Address(v string)`

SetIscsiIpv4Address sets IscsiIpv4Address field to given value.

### HasIscsiIpv4Address

`func (o *VnicEthIfInventory) HasIscsiIpv4Address() bool`

HasIscsiIpv4Address returns a boolean if a field has been set.

### GetMacAddress

`func (o *VnicEthIfInventory) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VnicEthIfInventory) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VnicEthIfInventory) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VnicEthIfInventory) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacAddressType

`func (o *VnicEthIfInventory) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *VnicEthIfInventory) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *VnicEthIfInventory) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *VnicEthIfInventory) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

### GetName

`func (o *VnicEthIfInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthIfInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthIfInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthIfInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicEthIfInventory) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicEthIfInventory) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicEthIfInventory) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicEthIfInventory) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicEthIfInventory) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicEthIfInventory) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicEthIfInventory) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicEthIfInventory) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicEthIfInventory) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicEthIfInventory) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicEthIfInventory) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicEthIfInventory) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicEthIfInventory) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicEthIfInventory) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStandbyVifId

`func (o *VnicEthIfInventory) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *VnicEthIfInventory) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *VnicEthIfInventory) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *VnicEthIfInventory) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetStaticMacAddress

`func (o *VnicEthIfInventory) GetStaticMacAddress() string`

GetStaticMacAddress returns the StaticMacAddress field if non-nil, zero value otherwise.

### GetStaticMacAddressOk

`func (o *VnicEthIfInventory) GetStaticMacAddressOk() (*string, bool)`

GetStaticMacAddressOk returns a tuple with the StaticMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddress

`func (o *VnicEthIfInventory) SetStaticMacAddress(v string)`

SetStaticMacAddress sets StaticMacAddress field to given value.

### HasStaticMacAddress

`func (o *VnicEthIfInventory) HasStaticMacAddress() bool`

HasStaticMacAddress returns a boolean if a field has been set.

### GetUsnicSettings

`func (o *VnicEthIfInventory) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicEthIfInventory) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicEthIfInventory) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicEthIfInventory) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### SetUsnicSettingsNil

`func (o *VnicEthIfInventory) SetUsnicSettingsNil(b bool)`

 SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil

### UnsetUsnicSettings
`func (o *VnicEthIfInventory) UnsetUsnicSettings()`

UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
### GetVifId

`func (o *VnicEthIfInventory) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicEthIfInventory) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicEthIfInventory) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicEthIfInventory) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVmqSettings

`func (o *VnicEthIfInventory) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicEthIfInventory) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicEthIfInventory) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicEthIfInventory) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### SetVmqSettingsNil

`func (o *VnicEthIfInventory) SetVmqSettingsNil(b bool)`

 SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil

### UnsetVmqSettings
`func (o *VnicEthIfInventory) UnsetVmqSettings()`

UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
### GetEthAdapterPolicy

`func (o *VnicEthIfInventory) GetEthAdapterPolicy() VnicEthAdapterPolicyInventoryRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicEthIfInventory) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyInventoryRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicEthIfInventory) SetEthAdapterPolicy(v VnicEthAdapterPolicyInventoryRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicEthIfInventory) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicEthIfInventory) GetEthNetworkPolicy() VnicEthNetworkPolicyInventoryRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicEthIfInventory) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyInventoryRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicEthIfInventory) SetEthNetworkPolicy(v VnicEthNetworkPolicyInventoryRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicEthIfInventory) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicEthIfInventory) GetEthQosPolicy() VnicEthQosPolicyInventoryRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicEthIfInventory) GetEthQosPolicyOk() (*VnicEthQosPolicyInventoryRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicEthIfInventory) SetEthQosPolicy(v VnicEthQosPolicyInventoryRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicEthIfInventory) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventory) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyInventoryRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicEthIfInventory) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyInventoryRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventory) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyInventoryRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicEthIfInventory) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventory) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyInventoryRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicEthIfInventory) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyInventoryRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventory) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyInventoryRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicEthIfInventory) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicEthIfInventory) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicEthIfInventory) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetIpLease

`func (o *VnicEthIfInventory) GetIpLease() IppoolIpLeaseRelationship`

GetIpLease returns the IpLease field if non-nil, zero value otherwise.

### GetIpLeaseOk

`func (o *VnicEthIfInventory) GetIpLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpLeaseOk returns a tuple with the IpLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLease

`func (o *VnicEthIfInventory) SetIpLease(v IppoolIpLeaseRelationship)`

SetIpLease sets IpLease field to given value.

### HasIpLease

`func (o *VnicEthIfInventory) HasIpLease() bool`

HasIpLease returns a boolean if a field has been set.

### GetIscsiBootPolicy

`func (o *VnicEthIfInventory) GetIscsiBootPolicy() VnicIscsiBootPolicyInventoryRelationship`

GetIscsiBootPolicy returns the IscsiBootPolicy field if non-nil, zero value otherwise.

### GetIscsiBootPolicyOk

`func (o *VnicEthIfInventory) GetIscsiBootPolicyOk() (*VnicIscsiBootPolicyInventoryRelationship, bool)`

GetIscsiBootPolicyOk returns a tuple with the IscsiBootPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiBootPolicy

`func (o *VnicEthIfInventory) SetIscsiBootPolicy(v VnicIscsiBootPolicyInventoryRelationship)`

SetIscsiBootPolicy sets IscsiBootPolicy field to given value.

### HasIscsiBootPolicy

`func (o *VnicEthIfInventory) HasIscsiBootPolicy() bool`

HasIscsiBootPolicy returns a boolean if a field has been set.

### GetLanConnectivityPolicy

`func (o *VnicEthIfInventory) GetLanConnectivityPolicy() VnicLanConnectivityPolicyInventoryRelationship`

GetLanConnectivityPolicy returns the LanConnectivityPolicy field if non-nil, zero value otherwise.

### GetLanConnectivityPolicyOk

`func (o *VnicEthIfInventory) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyInventoryRelationship, bool)`

GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanConnectivityPolicy

`func (o *VnicEthIfInventory) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyInventoryRelationship)`

SetLanConnectivityPolicy sets LanConnectivityPolicy field to given value.

### HasLanConnectivityPolicy

`func (o *VnicEthIfInventory) HasLanConnectivityPolicy() bool`

HasLanConnectivityPolicy returns a boolean if a field has been set.

### GetLcpVnic

`func (o *VnicEthIfInventory) GetLcpVnic() VnicEthIfInventoryRelationship`

GetLcpVnic returns the LcpVnic field if non-nil, zero value otherwise.

### GetLcpVnicOk

`func (o *VnicEthIfInventory) GetLcpVnicOk() (*VnicEthIfInventoryRelationship, bool)`

GetLcpVnicOk returns a tuple with the LcpVnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpVnic

`func (o *VnicEthIfInventory) SetLcpVnic(v VnicEthIfInventoryRelationship)`

SetLcpVnic sets LcpVnic field to given value.

### HasLcpVnic

`func (o *VnicEthIfInventory) HasLcpVnic() bool`

HasLcpVnic returns a boolean if a field has been set.

### GetMacLease

`func (o *VnicEthIfInventory) GetMacLease() MacpoolLeaseRelationship`

GetMacLease returns the MacLease field if non-nil, zero value otherwise.

### GetMacLeaseOk

`func (o *VnicEthIfInventory) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool)`

GetMacLeaseOk returns a tuple with the MacLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLease

`func (o *VnicEthIfInventory) SetMacLease(v MacpoolLeaseRelationship)`

SetMacLease sets MacLease field to given value.

### HasMacLease

`func (o *VnicEthIfInventory) HasMacLease() bool`

HasMacLease returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicEthIfInventory) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicEthIfInventory) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicEthIfInventory) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicEthIfInventory) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetSpVnics

`func (o *VnicEthIfInventory) GetSpVnics() []VnicEthIfInventoryRelationship`

GetSpVnics returns the SpVnics field if non-nil, zero value otherwise.

### GetSpVnicsOk

`func (o *VnicEthIfInventory) GetSpVnicsOk() (*[]VnicEthIfInventoryRelationship, bool)`

GetSpVnicsOk returns a tuple with the SpVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVnics

`func (o *VnicEthIfInventory) SetSpVnics(v []VnicEthIfInventoryRelationship)`

SetSpVnics sets SpVnics field to given value.

### HasSpVnics

`func (o *VnicEthIfInventory) HasSpVnics() bool`

HasSpVnics returns a boolean if a field has been set.

### SetSpVnicsNil

`func (o *VnicEthIfInventory) SetSpVnicsNil(b bool)`

 SetSpVnicsNil sets the value for SpVnics to be an explicit nil

### UnsetSpVnics
`func (o *VnicEthIfInventory) UnsetSpVnics()`

UnsetSpVnics ensures that no value is present for SpVnics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


