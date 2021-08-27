# VnicEthIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthIf"]
**Cdn** | Pointer to [**NullableVnicCdn**](VnicCdn.md) |  | [optional] 
**FailoverEnabled** | Pointer to **bool** | Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover. | [optional] [default to false]
**IscsiIpV4AddressAllocationType** | Pointer to **string** | Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type. * &#x60;None&#x60; - Type indicates that there is no IP associated to an vnic. * &#x60;DHCP&#x60; - The IP address is assigned using DHCP, if available. * &#x60;Static&#x60; - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * &#x60;Pool&#x60; - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. | [optional] [readonly] [default to "None"]
**IscsiIpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IscsiIpv4Address** | Pointer to **string** | IP address associated to the vNIC. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. | [optional] [readonly] 
**MacAddressType** | Pointer to **string** | Type of allocation selected to assign a MAC address for the vnic. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [default to "POOL"]
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StandbyVifId** | Pointer to **int64** | The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. | [optional] [readonly] 
**StaticMacAddress** | Pointer to **string** | The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx. To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the following MAC prefix 00:25:B5:xx:xx:xx. | [optional] 
**UsnicSettings** | Pointer to [**NullableVnicUsnicSettings**](VnicUsnicSettings.md) |  | [optional] 
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**VmqSettings** | Pointer to [**NullableVnicVmqSettings**](VnicVmqSettings.md) |  | [optional] 
**EthAdapterPolicy** | Pointer to [**VnicEthAdapterPolicyRelationship**](VnicEthAdapterPolicyRelationship.md) |  | [optional] 
**EthNetworkPolicy** | Pointer to [**VnicEthNetworkPolicyRelationship**](VnicEthNetworkPolicyRelationship.md) |  | [optional] 
**EthQosPolicy** | Pointer to [**VnicEthQosPolicyRelationship**](VnicEthQosPolicyRelationship.md) |  | [optional] 
**FabricEthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](FabricEthNetworkControlPolicyRelationship.md) |  | [optional] 
**FabricEthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyRelationship**](FabricEthNetworkGroupPolicyRelationship.md) | An array of relationships to fabricEthNetworkGroupPolicy resources. | [optional] 
**IpLease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**IscsiBootPolicy** | Pointer to [**VnicIscsiBootPolicyRelationship**](VnicIscsiBootPolicyRelationship.md) |  | [optional] 
**LanConnectivityPolicy** | Pointer to [**VnicLanConnectivityPolicyRelationship**](VnicLanConnectivityPolicyRelationship.md) |  | [optional] 
**LcpVnic** | Pointer to [**VnicEthIfRelationship**](VnicEthIfRelationship.md) |  | [optional] 
**MacLease** | Pointer to [**MacpoolLeaseRelationship**](MacpoolLeaseRelationship.md) |  | [optional] 
**MacPool** | Pointer to [**MacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**SpVnics** | Pointer to [**[]VnicEthIfRelationship**](VnicEthIfRelationship.md) | An array of relationships to vnicEthIf resources. | [optional] [readonly] 

## Methods

### NewVnicEthIf

`func NewVnicEthIf(classId string, objectType string, ) *VnicEthIf`

NewVnicEthIf instantiates a new VnicEthIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfWithDefaults

`func NewVnicEthIfWithDefaults() *VnicEthIf`

NewVnicEthIfWithDefaults instantiates a new VnicEthIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdn

`func (o *VnicEthIf) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicEthIf) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicEthIf) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicEthIf) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *VnicEthIf) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *VnicEthIf) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetFailoverEnabled

`func (o *VnicEthIf) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicEthIf) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicEthIf) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicEthIf) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetIscsiIpV4AddressAllocationType

`func (o *VnicEthIf) GetIscsiIpV4AddressAllocationType() string`

GetIscsiIpV4AddressAllocationType returns the IscsiIpV4AddressAllocationType field if non-nil, zero value otherwise.

### GetIscsiIpV4AddressAllocationTypeOk

`func (o *VnicEthIf) GetIscsiIpV4AddressAllocationTypeOk() (*string, bool)`

GetIscsiIpV4AddressAllocationTypeOk returns a tuple with the IscsiIpV4AddressAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4AddressAllocationType

`func (o *VnicEthIf) SetIscsiIpV4AddressAllocationType(v string)`

SetIscsiIpV4AddressAllocationType sets IscsiIpV4AddressAllocationType field to given value.

### HasIscsiIpV4AddressAllocationType

`func (o *VnicEthIf) HasIscsiIpV4AddressAllocationType() bool`

HasIscsiIpV4AddressAllocationType returns a boolean if a field has been set.

### GetIscsiIpV4Config

`func (o *VnicEthIf) GetIscsiIpV4Config() IppoolIpV4Config`

GetIscsiIpV4Config returns the IscsiIpV4Config field if non-nil, zero value otherwise.

### GetIscsiIpV4ConfigOk

`func (o *VnicEthIf) GetIscsiIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIscsiIpV4ConfigOk returns a tuple with the IscsiIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4Config

`func (o *VnicEthIf) SetIscsiIpV4Config(v IppoolIpV4Config)`

SetIscsiIpV4Config sets IscsiIpV4Config field to given value.

### HasIscsiIpV4Config

`func (o *VnicEthIf) HasIscsiIpV4Config() bool`

HasIscsiIpV4Config returns a boolean if a field has been set.

### SetIscsiIpV4ConfigNil

`func (o *VnicEthIf) SetIscsiIpV4ConfigNil(b bool)`

 SetIscsiIpV4ConfigNil sets the value for IscsiIpV4Config to be an explicit nil

### UnsetIscsiIpV4Config
`func (o *VnicEthIf) UnsetIscsiIpV4Config()`

UnsetIscsiIpV4Config ensures that no value is present for IscsiIpV4Config, not even an explicit nil
### GetIscsiIpv4Address

`func (o *VnicEthIf) GetIscsiIpv4Address() string`

GetIscsiIpv4Address returns the IscsiIpv4Address field if non-nil, zero value otherwise.

### GetIscsiIpv4AddressOk

`func (o *VnicEthIf) GetIscsiIpv4AddressOk() (*string, bool)`

GetIscsiIpv4AddressOk returns a tuple with the IscsiIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpv4Address

`func (o *VnicEthIf) SetIscsiIpv4Address(v string)`

SetIscsiIpv4Address sets IscsiIpv4Address field to given value.

### HasIscsiIpv4Address

`func (o *VnicEthIf) HasIscsiIpv4Address() bool`

HasIscsiIpv4Address returns a boolean if a field has been set.

### GetMacAddress

`func (o *VnicEthIf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VnicEthIf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VnicEthIf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VnicEthIf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacAddressType

`func (o *VnicEthIf) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *VnicEthIf) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *VnicEthIf) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *VnicEthIf) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

### GetName

`func (o *VnicEthIf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthIf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthIf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthIf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicEthIf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicEthIf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicEthIf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicEthIf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicEthIf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicEthIf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicEthIf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicEthIf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicEthIf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicEthIf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStandbyVifId

`func (o *VnicEthIf) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *VnicEthIf) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *VnicEthIf) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *VnicEthIf) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetStaticMacAddress

`func (o *VnicEthIf) GetStaticMacAddress() string`

GetStaticMacAddress returns the StaticMacAddress field if non-nil, zero value otherwise.

### GetStaticMacAddressOk

`func (o *VnicEthIf) GetStaticMacAddressOk() (*string, bool)`

GetStaticMacAddressOk returns a tuple with the StaticMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddress

`func (o *VnicEthIf) SetStaticMacAddress(v string)`

SetStaticMacAddress sets StaticMacAddress field to given value.

### HasStaticMacAddress

`func (o *VnicEthIf) HasStaticMacAddress() bool`

HasStaticMacAddress returns a boolean if a field has been set.

### GetUsnicSettings

`func (o *VnicEthIf) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicEthIf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicEthIf) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicEthIf) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### SetUsnicSettingsNil

`func (o *VnicEthIf) SetUsnicSettingsNil(b bool)`

 SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil

### UnsetUsnicSettings
`func (o *VnicEthIf) UnsetUsnicSettings()`

UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
### GetVifId

`func (o *VnicEthIf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicEthIf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicEthIf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicEthIf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVmqSettings

`func (o *VnicEthIf) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicEthIf) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicEthIf) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicEthIf) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### SetVmqSettingsNil

`func (o *VnicEthIf) SetVmqSettingsNil(b bool)`

 SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil

### UnsetVmqSettings
`func (o *VnicEthIf) UnsetVmqSettings()`

UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
### GetEthAdapterPolicy

`func (o *VnicEthIf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicEthIf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicEthIf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicEthIf) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicEthIf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicEthIf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicEthIf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicEthIf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicEthIf) GetEthQosPolicy() VnicEthQosPolicyRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicEthIf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicEthIf) SetEthQosPolicy(v VnicEthQosPolicyRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicEthIf) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicEthIf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicEthIf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicEthIf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicEthIf) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicEthIf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicEthIf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicEthIf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicEthIf) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicEthIf) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicEthIf) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetIpLease

`func (o *VnicEthIf) GetIpLease() IppoolIpLeaseRelationship`

GetIpLease returns the IpLease field if non-nil, zero value otherwise.

### GetIpLeaseOk

`func (o *VnicEthIf) GetIpLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpLeaseOk returns a tuple with the IpLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLease

`func (o *VnicEthIf) SetIpLease(v IppoolIpLeaseRelationship)`

SetIpLease sets IpLease field to given value.

### HasIpLease

`func (o *VnicEthIf) HasIpLease() bool`

HasIpLease returns a boolean if a field has been set.

### GetIscsiBootPolicy

`func (o *VnicEthIf) GetIscsiBootPolicy() VnicIscsiBootPolicyRelationship`

GetIscsiBootPolicy returns the IscsiBootPolicy field if non-nil, zero value otherwise.

### GetIscsiBootPolicyOk

`func (o *VnicEthIf) GetIscsiBootPolicyOk() (*VnicIscsiBootPolicyRelationship, bool)`

GetIscsiBootPolicyOk returns a tuple with the IscsiBootPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiBootPolicy

`func (o *VnicEthIf) SetIscsiBootPolicy(v VnicIscsiBootPolicyRelationship)`

SetIscsiBootPolicy sets IscsiBootPolicy field to given value.

### HasIscsiBootPolicy

`func (o *VnicEthIf) HasIscsiBootPolicy() bool`

HasIscsiBootPolicy returns a boolean if a field has been set.

### GetLanConnectivityPolicy

`func (o *VnicEthIf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship`

GetLanConnectivityPolicy returns the LanConnectivityPolicy field if non-nil, zero value otherwise.

### GetLanConnectivityPolicyOk

`func (o *VnicEthIf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool)`

GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanConnectivityPolicy

`func (o *VnicEthIf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship)`

SetLanConnectivityPolicy sets LanConnectivityPolicy field to given value.

### HasLanConnectivityPolicy

`func (o *VnicEthIf) HasLanConnectivityPolicy() bool`

HasLanConnectivityPolicy returns a boolean if a field has been set.

### GetLcpVnic

`func (o *VnicEthIf) GetLcpVnic() VnicEthIfRelationship`

GetLcpVnic returns the LcpVnic field if non-nil, zero value otherwise.

### GetLcpVnicOk

`func (o *VnicEthIf) GetLcpVnicOk() (*VnicEthIfRelationship, bool)`

GetLcpVnicOk returns a tuple with the LcpVnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpVnic

`func (o *VnicEthIf) SetLcpVnic(v VnicEthIfRelationship)`

SetLcpVnic sets LcpVnic field to given value.

### HasLcpVnic

`func (o *VnicEthIf) HasLcpVnic() bool`

HasLcpVnic returns a boolean if a field has been set.

### GetMacLease

`func (o *VnicEthIf) GetMacLease() MacpoolLeaseRelationship`

GetMacLease returns the MacLease field if non-nil, zero value otherwise.

### GetMacLeaseOk

`func (o *VnicEthIf) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool)`

GetMacLeaseOk returns a tuple with the MacLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLease

`func (o *VnicEthIf) SetMacLease(v MacpoolLeaseRelationship)`

SetMacLease sets MacLease field to given value.

### HasMacLease

`func (o *VnicEthIf) HasMacLease() bool`

HasMacLease returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicEthIf) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicEthIf) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicEthIf) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicEthIf) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetProfile

`func (o *VnicEthIf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicEthIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicEthIf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicEthIf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSpVnics

`func (o *VnicEthIf) GetSpVnics() []VnicEthIfRelationship`

GetSpVnics returns the SpVnics field if non-nil, zero value otherwise.

### GetSpVnicsOk

`func (o *VnicEthIf) GetSpVnicsOk() (*[]VnicEthIfRelationship, bool)`

GetSpVnicsOk returns a tuple with the SpVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVnics

`func (o *VnicEthIf) SetSpVnics(v []VnicEthIfRelationship)`

SetSpVnics sets SpVnics field to given value.

### HasSpVnics

`func (o *VnicEthIf) HasSpVnics() bool`

HasSpVnics returns a boolean if a field has been set.

### SetSpVnicsNil

`func (o *VnicEthIf) SetSpVnicsNil(b bool)`

 SetSpVnicsNil sets the value for SpVnics to be an explicit nil

### UnsetSpVnics
`func (o *VnicEthIf) UnsetSpVnics()`

UnsetSpVnics ensures that no value is present for SpVnics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


