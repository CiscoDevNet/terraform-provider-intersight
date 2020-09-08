# VnicEthIfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**VnicCdn**](vnic.Cdn.md) |  | [optional] 
**FailoverEnabled** | Pointer to **bool** | Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address that is assigned to the vnic based on the MAC pool that has been assigned to the LAN Connectivity Policy. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**Placement** | Pointer to [**VnicPlacementSettings**](vnic.PlacementSettings.md) |  | [optional] 
**StandbyVifId** | Pointer to **int64** | The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. | [optional] [readonly] 
**UsnicSettings** | Pointer to [**VnicUsnicSettings**](vnic.UsnicSettings.md) |  | [optional] 
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**VmqSettings** | Pointer to [**VnicVmqSettings**](vnic.VmqSettings.md) |  | [optional] 
**EthAdapterPolicy** | Pointer to [**VnicEthAdapterPolicyRelationship**](vnic.EthAdapterPolicy.Relationship.md) |  | [optional] 
**EthNetworkPolicy** | Pointer to [**VnicEthNetworkPolicyRelationship**](vnic.EthNetworkPolicy.Relationship.md) |  | [optional] 
**EthQosPolicy** | Pointer to [**VnicEthQosPolicyRelationship**](vnic.EthQosPolicy.Relationship.md) |  | [optional] 
**FabricEthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](fabric.EthNetworkControlPolicy.Relationship.md) |  | [optional] 
**FabricEthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyRelationship**](fabric.EthNetworkGroupPolicy.Relationship.md) | An array of relationships to fabricEthNetworkGroupPolicy resources. | [optional] 
**LanConnectivityPolicy** | Pointer to [**VnicLanConnectivityPolicyRelationship**](vnic.LanConnectivityPolicy.Relationship.md) |  | [optional] 
**LcpVnic** | Pointer to [**VnicEthIfRelationship**](vnic.EthIf.Relationship.md) |  | [optional] 
**MacLease** | Pointer to [**MacpoolLeaseRelationship**](macpool.Lease.Relationship.md) |  | [optional] 
**MacPool** | Pointer to [**MacpoolPoolRelationship**](macpool.Pool.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) |  | [optional] 
**SpVnics** | Pointer to [**[]VnicEthIfRelationship**](vnic.EthIf.Relationship.md) | An array of relationships to vnicEthIf resources. | [optional] 

## Methods

### NewVnicEthIfAllOf

`func NewVnicEthIfAllOf() *VnicEthIfAllOf`

NewVnicEthIfAllOf instantiates a new VnicEthIfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfAllOfWithDefaults

`func NewVnicEthIfAllOfWithDefaults() *VnicEthIfAllOf`

NewVnicEthIfAllOfWithDefaults instantiates a new VnicEthIfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *VnicEthIfAllOf) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicEthIfAllOf) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicEthIfAllOf) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicEthIfAllOf) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetFailoverEnabled

`func (o *VnicEthIfAllOf) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicEthIfAllOf) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicEthIfAllOf) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicEthIfAllOf) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *VnicEthIfAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VnicEthIfAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VnicEthIfAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VnicEthIfAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *VnicEthIfAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthIfAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthIfAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthIfAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicEthIfAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicEthIfAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicEthIfAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicEthIfAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicEthIfAllOf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicEthIfAllOf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicEthIfAllOf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicEthIfAllOf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetStandbyVifId

`func (o *VnicEthIfAllOf) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *VnicEthIfAllOf) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *VnicEthIfAllOf) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *VnicEthIfAllOf) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetUsnicSettings

`func (o *VnicEthIfAllOf) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicEthIfAllOf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicEthIfAllOf) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicEthIfAllOf) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### GetVifId

`func (o *VnicEthIfAllOf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicEthIfAllOf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicEthIfAllOf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicEthIfAllOf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVmqSettings

`func (o *VnicEthIfAllOf) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicEthIfAllOf) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicEthIfAllOf) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicEthIfAllOf) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### GetEthAdapterPolicy

`func (o *VnicEthIfAllOf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicEthIfAllOf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicEthIfAllOf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicEthIfAllOf) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicEthIfAllOf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicEthIfAllOf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicEthIfAllOf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicEthIfAllOf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicEthIfAllOf) GetEthQosPolicy() VnicEthQosPolicyRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicEthIfAllOf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicEthIfAllOf) SetEthQosPolicy(v VnicEthQosPolicyRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicEthIfAllOf) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicEthIfAllOf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicEthIfAllOf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicEthIfAllOf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicEthIfAllOf) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfAllOf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicEthIfAllOf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfAllOf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicEthIfAllOf) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicEthIfAllOf) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicEthIfAllOf) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetLanConnectivityPolicy

`func (o *VnicEthIfAllOf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship`

GetLanConnectivityPolicy returns the LanConnectivityPolicy field if non-nil, zero value otherwise.

### GetLanConnectivityPolicyOk

`func (o *VnicEthIfAllOf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool)`

GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanConnectivityPolicy

`func (o *VnicEthIfAllOf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship)`

SetLanConnectivityPolicy sets LanConnectivityPolicy field to given value.

### HasLanConnectivityPolicy

`func (o *VnicEthIfAllOf) HasLanConnectivityPolicy() bool`

HasLanConnectivityPolicy returns a boolean if a field has been set.

### GetLcpVnic

`func (o *VnicEthIfAllOf) GetLcpVnic() VnicEthIfRelationship`

GetLcpVnic returns the LcpVnic field if non-nil, zero value otherwise.

### GetLcpVnicOk

`func (o *VnicEthIfAllOf) GetLcpVnicOk() (*VnicEthIfRelationship, bool)`

GetLcpVnicOk returns a tuple with the LcpVnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpVnic

`func (o *VnicEthIfAllOf) SetLcpVnic(v VnicEthIfRelationship)`

SetLcpVnic sets LcpVnic field to given value.

### HasLcpVnic

`func (o *VnicEthIfAllOf) HasLcpVnic() bool`

HasLcpVnic returns a boolean if a field has been set.

### GetMacLease

`func (o *VnicEthIfAllOf) GetMacLease() MacpoolLeaseRelationship`

GetMacLease returns the MacLease field if non-nil, zero value otherwise.

### GetMacLeaseOk

`func (o *VnicEthIfAllOf) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool)`

GetMacLeaseOk returns a tuple with the MacLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLease

`func (o *VnicEthIfAllOf) SetMacLease(v MacpoolLeaseRelationship)`

SetMacLease sets MacLease field to given value.

### HasMacLease

`func (o *VnicEthIfAllOf) HasMacLease() bool`

HasMacLease returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicEthIfAllOf) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicEthIfAllOf) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicEthIfAllOf) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicEthIfAllOf) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetProfile

`func (o *VnicEthIfAllOf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicEthIfAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicEthIfAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicEthIfAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSpVnics

`func (o *VnicEthIfAllOf) GetSpVnics() []VnicEthIfRelationship`

GetSpVnics returns the SpVnics field if non-nil, zero value otherwise.

### GetSpVnicsOk

`func (o *VnicEthIfAllOf) GetSpVnicsOk() (*[]VnicEthIfRelationship, bool)`

GetSpVnicsOk returns a tuple with the SpVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVnics

`func (o *VnicEthIfAllOf) SetSpVnics(v []VnicEthIfRelationship)`

SetSpVnics sets SpVnics field to given value.

### HasSpVnics

`func (o *VnicEthIfAllOf) HasSpVnics() bool`

HasSpVnics returns a boolean if a field has been set.

### SetSpVnicsNil

`func (o *VnicEthIfAllOf) SetSpVnicsNil(b bool)`

 SetSpVnicsNil sets the value for SpVnics to be an explicit nil

### UnsetSpVnics
`func (o *VnicEthIfAllOf) UnsetSpVnics()`

UnsetSpVnics ensures that no value is present for SpVnics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


