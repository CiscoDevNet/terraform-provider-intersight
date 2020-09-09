# VnicFcIfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] 
**Placement** | Pointer to [**VnicPlacementSettings**](vnic.PlacementSettings.md) |  | [optional] 
**Type** | Pointer to **string** | VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * &#x60;fc-initiator&#x60; - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-initiator&#x60; - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-target&#x60; - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * &#x60;fc-target&#x60; - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. | [optional] [default to "fc-initiator"]
**VifId** | Pointer to **int64** | This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The WWPN address that is assigned to the vhba based on the wwn pool that has been assigned to the SAN Connectivity Policy. | [optional] [readonly] 
**FcAdapterPolicy** | Pointer to [**VnicFcAdapterPolicyRelationship**](vnic.FcAdapterPolicy.Relationship.md) |  | [optional] 
**FcNetworkPolicy** | Pointer to [**VnicFcNetworkPolicyRelationship**](vnic.FcNetworkPolicy.Relationship.md) |  | [optional] 
**FcQosPolicy** | Pointer to [**VnicFcQosPolicyRelationship**](vnic.FcQosPolicy.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) |  | [optional] 
**SanConnectivityPolicy** | Pointer to [**VnicSanConnectivityPolicyRelationship**](vnic.SanConnectivityPolicy.Relationship.md) |  | [optional] 
**ScpVhba** | Pointer to [**VnicFcIfRelationship**](vnic.FcIf.Relationship.md) |  | [optional] 
**SpVhbas** | Pointer to [**[]VnicFcIfRelationship**](vnic.FcIf.Relationship.md) | An array of relationships to vnicFcIf resources. | [optional] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](fcpool.Lease.Relationship.md) |  | [optional] 
**WwpnPool** | Pointer to [**FcpoolPoolRelationship**](fcpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewVnicFcIfAllOf

`func NewVnicFcIfAllOf() *VnicFcIfAllOf`

NewVnicFcIfAllOf instantiates a new VnicFcIfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfAllOfWithDefaults

`func NewVnicFcIfAllOfWithDefaults() *VnicFcIfAllOf`

NewVnicFcIfAllOfWithDefaults instantiates a new VnicFcIfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VnicFcIfAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcIfAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcIfAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcIfAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicFcIfAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicFcIfAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicFcIfAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicFcIfAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPersistentBindings

`func (o *VnicFcIfAllOf) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicFcIfAllOf) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicFcIfAllOf) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicFcIfAllOf) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicFcIfAllOf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicFcIfAllOf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicFcIfAllOf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicFcIfAllOf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetType

`func (o *VnicFcIfAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicFcIfAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicFcIfAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicFcIfAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVifId

`func (o *VnicFcIfAllOf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicFcIfAllOf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicFcIfAllOf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicFcIfAllOf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetWwpn

`func (o *VnicFcIfAllOf) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *VnicFcIfAllOf) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *VnicFcIfAllOf) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *VnicFcIfAllOf) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicFcIfAllOf) GetFcAdapterPolicy() VnicFcAdapterPolicyRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicFcIfAllOf) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicFcIfAllOf) SetFcAdapterPolicy(v VnicFcAdapterPolicyRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicFcIfAllOf) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *VnicFcIfAllOf) GetFcNetworkPolicy() VnicFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicFcIfAllOf) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicFcIfAllOf) SetFcNetworkPolicy(v VnicFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicFcIfAllOf) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### GetFcQosPolicy

`func (o *VnicFcIfAllOf) GetFcQosPolicy() VnicFcQosPolicyRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicFcIfAllOf) GetFcQosPolicyOk() (*VnicFcQosPolicyRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicFcIfAllOf) SetFcQosPolicy(v VnicFcQosPolicyRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicFcIfAllOf) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### GetProfile

`func (o *VnicFcIfAllOf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicFcIfAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicFcIfAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicFcIfAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSanConnectivityPolicy

`func (o *VnicFcIfAllOf) GetSanConnectivityPolicy() VnicSanConnectivityPolicyRelationship`

GetSanConnectivityPolicy returns the SanConnectivityPolicy field if non-nil, zero value otherwise.

### GetSanConnectivityPolicyOk

`func (o *VnicFcIfAllOf) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyRelationship, bool)`

GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanConnectivityPolicy

`func (o *VnicFcIfAllOf) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyRelationship)`

SetSanConnectivityPolicy sets SanConnectivityPolicy field to given value.

### HasSanConnectivityPolicy

`func (o *VnicFcIfAllOf) HasSanConnectivityPolicy() bool`

HasSanConnectivityPolicy returns a boolean if a field has been set.

### GetScpVhba

`func (o *VnicFcIfAllOf) GetScpVhba() VnicFcIfRelationship`

GetScpVhba returns the ScpVhba field if non-nil, zero value otherwise.

### GetScpVhbaOk

`func (o *VnicFcIfAllOf) GetScpVhbaOk() (*VnicFcIfRelationship, bool)`

GetScpVhbaOk returns a tuple with the ScpVhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpVhba

`func (o *VnicFcIfAllOf) SetScpVhba(v VnicFcIfRelationship)`

SetScpVhba sets ScpVhba field to given value.

### HasScpVhba

`func (o *VnicFcIfAllOf) HasScpVhba() bool`

HasScpVhba returns a boolean if a field has been set.

### GetSpVhbas

`func (o *VnicFcIfAllOf) GetSpVhbas() []VnicFcIfRelationship`

GetSpVhbas returns the SpVhbas field if non-nil, zero value otherwise.

### GetSpVhbasOk

`func (o *VnicFcIfAllOf) GetSpVhbasOk() (*[]VnicFcIfRelationship, bool)`

GetSpVhbasOk returns a tuple with the SpVhbas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVhbas

`func (o *VnicFcIfAllOf) SetSpVhbas(v []VnicFcIfRelationship)`

SetSpVhbas sets SpVhbas field to given value.

### HasSpVhbas

`func (o *VnicFcIfAllOf) HasSpVhbas() bool`

HasSpVhbas returns a boolean if a field has been set.

### SetSpVhbasNil

`func (o *VnicFcIfAllOf) SetSpVhbasNil(b bool)`

 SetSpVhbasNil sets the value for SpVhbas to be an explicit nil

### UnsetSpVhbas
`func (o *VnicFcIfAllOf) UnsetSpVhbas()`

UnsetSpVhbas ensures that no value is present for SpVhbas, not even an explicit nil
### GetWwpnLease

`func (o *VnicFcIfAllOf) GetWwpnLease() FcpoolLeaseRelationship`

GetWwpnLease returns the WwpnLease field if non-nil, zero value otherwise.

### GetWwpnLeaseOk

`func (o *VnicFcIfAllOf) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwpnLeaseOk returns a tuple with the WwpnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnLease

`func (o *VnicFcIfAllOf) SetWwpnLease(v FcpoolLeaseRelationship)`

SetWwpnLease sets WwpnLease field to given value.

### HasWwpnLease

`func (o *VnicFcIfAllOf) HasWwpnLease() bool`

HasWwpnLease returns a boolean if a field has been set.

### GetWwpnPool

`func (o *VnicFcIfAllOf) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicFcIfAllOf) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicFcIfAllOf) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicFcIfAllOf) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


