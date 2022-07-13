# VnicFcIfInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcIfInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcIfInventory"]
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] [readonly] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] [readonly] 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] [readonly] 
**PinGroupName** | Pointer to **string** | Pingroup name associated to vfc for static pinning. SCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vfc traffic. | [optional] [readonly] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StaticWwpnAddress** | Pointer to **string** | The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. | [optional] [readonly] 
**Type** | Pointer to **string** | VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * &#x60;fc-initiator&#x60; - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-initiator&#x60; - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-target&#x60; - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * &#x60;fc-target&#x60; - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. | [optional] [readonly] [default to "fc-initiator"]
**VifId** | Pointer to **int64** | This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy. | [optional] [readonly] 
**WwpnAddressType** | Pointer to **string** | Type of allocation selected to assign a WWPN address to the vhba. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [readonly] [default to "POOL"]
**FcAdapterPolicy** | Pointer to [**VnicFcAdapterPolicyInventoryRelationship**](VnicFcAdapterPolicyInventoryRelationship.md) |  | [optional] 
**FcNetworkPolicy** | Pointer to [**VnicFcNetworkPolicyInventoryRelationship**](VnicFcNetworkPolicyInventoryRelationship.md) |  | [optional] 
**FcQosPolicy** | Pointer to [**VnicFcQosPolicyInventoryRelationship**](VnicFcQosPolicyInventoryRelationship.md) |  | [optional] 
**FcZonePolicies** | Pointer to [**[]FabricFcZonePolicyRelationship**](FabricFcZonePolicyRelationship.md) | An array of relationships to fabricFcZonePolicy resources. | [optional] [readonly] 
**SanConnectivityPolicy** | Pointer to [**VnicSanConnectivityPolicyInventoryRelationship**](VnicSanConnectivityPolicyInventoryRelationship.md) |  | [optional] 
**ScpVhba** | Pointer to [**VnicFcIfInventoryRelationship**](VnicFcIfInventoryRelationship.md) |  | [optional] 
**SpVhbas** | Pointer to [**[]VnicFcIfInventoryRelationship**](VnicFcIfInventoryRelationship.md) | An array of relationships to vnicFcIfInventory resources. | [optional] [readonly] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 
**WwpnPool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicFcIfInventory

`func NewVnicFcIfInventory(classId string, objectType string, ) *VnicFcIfInventory`

NewVnicFcIfInventory instantiates a new VnicFcIfInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfInventoryWithDefaults

`func NewVnicFcIfInventoryWithDefaults() *VnicFcIfInventory`

NewVnicFcIfInventoryWithDefaults instantiates a new VnicFcIfInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcIfInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcIfInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcIfInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcIfInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcIfInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcIfInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VnicFcIfInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcIfInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcIfInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcIfInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicFcIfInventory) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicFcIfInventory) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicFcIfInventory) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicFcIfInventory) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPersistentBindings

`func (o *VnicFcIfInventory) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicFcIfInventory) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicFcIfInventory) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicFcIfInventory) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicFcIfInventory) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicFcIfInventory) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicFcIfInventory) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicFcIfInventory) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicFcIfInventory) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicFcIfInventory) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicFcIfInventory) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicFcIfInventory) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicFcIfInventory) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicFcIfInventory) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStaticWwpnAddress

`func (o *VnicFcIfInventory) GetStaticWwpnAddress() string`

GetStaticWwpnAddress returns the StaticWwpnAddress field if non-nil, zero value otherwise.

### GetStaticWwpnAddressOk

`func (o *VnicFcIfInventory) GetStaticWwpnAddressOk() (*string, bool)`

GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticWwpnAddress

`func (o *VnicFcIfInventory) SetStaticWwpnAddress(v string)`

SetStaticWwpnAddress sets StaticWwpnAddress field to given value.

### HasStaticWwpnAddress

`func (o *VnicFcIfInventory) HasStaticWwpnAddress() bool`

HasStaticWwpnAddress returns a boolean if a field has been set.

### GetType

`func (o *VnicFcIfInventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicFcIfInventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicFcIfInventory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicFcIfInventory) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVifId

`func (o *VnicFcIfInventory) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicFcIfInventory) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicFcIfInventory) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicFcIfInventory) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetWwpn

`func (o *VnicFcIfInventory) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *VnicFcIfInventory) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *VnicFcIfInventory) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *VnicFcIfInventory) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetWwpnAddressType

`func (o *VnicFcIfInventory) GetWwpnAddressType() string`

GetWwpnAddressType returns the WwpnAddressType field if non-nil, zero value otherwise.

### GetWwpnAddressTypeOk

`func (o *VnicFcIfInventory) GetWwpnAddressTypeOk() (*string, bool)`

GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnAddressType

`func (o *VnicFcIfInventory) SetWwpnAddressType(v string)`

SetWwpnAddressType sets WwpnAddressType field to given value.

### HasWwpnAddressType

`func (o *VnicFcIfInventory) HasWwpnAddressType() bool`

HasWwpnAddressType returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicFcIfInventory) GetFcAdapterPolicy() VnicFcAdapterPolicyInventoryRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicFcIfInventory) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyInventoryRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicFcIfInventory) SetFcAdapterPolicy(v VnicFcAdapterPolicyInventoryRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicFcIfInventory) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *VnicFcIfInventory) GetFcNetworkPolicy() VnicFcNetworkPolicyInventoryRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicFcIfInventory) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyInventoryRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicFcIfInventory) SetFcNetworkPolicy(v VnicFcNetworkPolicyInventoryRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicFcIfInventory) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### GetFcQosPolicy

`func (o *VnicFcIfInventory) GetFcQosPolicy() VnicFcQosPolicyInventoryRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicFcIfInventory) GetFcQosPolicyOk() (*VnicFcQosPolicyInventoryRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicFcIfInventory) SetFcQosPolicy(v VnicFcQosPolicyInventoryRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicFcIfInventory) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### GetFcZonePolicies

`func (o *VnicFcIfInventory) GetFcZonePolicies() []FabricFcZonePolicyRelationship`

GetFcZonePolicies returns the FcZonePolicies field if non-nil, zero value otherwise.

### GetFcZonePoliciesOk

`func (o *VnicFcIfInventory) GetFcZonePoliciesOk() (*[]FabricFcZonePolicyRelationship, bool)`

GetFcZonePoliciesOk returns a tuple with the FcZonePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcZonePolicies

`func (o *VnicFcIfInventory) SetFcZonePolicies(v []FabricFcZonePolicyRelationship)`

SetFcZonePolicies sets FcZonePolicies field to given value.

### HasFcZonePolicies

`func (o *VnicFcIfInventory) HasFcZonePolicies() bool`

HasFcZonePolicies returns a boolean if a field has been set.

### SetFcZonePoliciesNil

`func (o *VnicFcIfInventory) SetFcZonePoliciesNil(b bool)`

 SetFcZonePoliciesNil sets the value for FcZonePolicies to be an explicit nil

### UnsetFcZonePolicies
`func (o *VnicFcIfInventory) UnsetFcZonePolicies()`

UnsetFcZonePolicies ensures that no value is present for FcZonePolicies, not even an explicit nil
### GetSanConnectivityPolicy

`func (o *VnicFcIfInventory) GetSanConnectivityPolicy() VnicSanConnectivityPolicyInventoryRelationship`

GetSanConnectivityPolicy returns the SanConnectivityPolicy field if non-nil, zero value otherwise.

### GetSanConnectivityPolicyOk

`func (o *VnicFcIfInventory) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyInventoryRelationship, bool)`

GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanConnectivityPolicy

`func (o *VnicFcIfInventory) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyInventoryRelationship)`

SetSanConnectivityPolicy sets SanConnectivityPolicy field to given value.

### HasSanConnectivityPolicy

`func (o *VnicFcIfInventory) HasSanConnectivityPolicy() bool`

HasSanConnectivityPolicy returns a boolean if a field has been set.

### GetScpVhba

`func (o *VnicFcIfInventory) GetScpVhba() VnicFcIfInventoryRelationship`

GetScpVhba returns the ScpVhba field if non-nil, zero value otherwise.

### GetScpVhbaOk

`func (o *VnicFcIfInventory) GetScpVhbaOk() (*VnicFcIfInventoryRelationship, bool)`

GetScpVhbaOk returns a tuple with the ScpVhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpVhba

`func (o *VnicFcIfInventory) SetScpVhba(v VnicFcIfInventoryRelationship)`

SetScpVhba sets ScpVhba field to given value.

### HasScpVhba

`func (o *VnicFcIfInventory) HasScpVhba() bool`

HasScpVhba returns a boolean if a field has been set.

### GetSpVhbas

`func (o *VnicFcIfInventory) GetSpVhbas() []VnicFcIfInventoryRelationship`

GetSpVhbas returns the SpVhbas field if non-nil, zero value otherwise.

### GetSpVhbasOk

`func (o *VnicFcIfInventory) GetSpVhbasOk() (*[]VnicFcIfInventoryRelationship, bool)`

GetSpVhbasOk returns a tuple with the SpVhbas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVhbas

`func (o *VnicFcIfInventory) SetSpVhbas(v []VnicFcIfInventoryRelationship)`

SetSpVhbas sets SpVhbas field to given value.

### HasSpVhbas

`func (o *VnicFcIfInventory) HasSpVhbas() bool`

HasSpVhbas returns a boolean if a field has been set.

### SetSpVhbasNil

`func (o *VnicFcIfInventory) SetSpVhbasNil(b bool)`

 SetSpVhbasNil sets the value for SpVhbas to be an explicit nil

### UnsetSpVhbas
`func (o *VnicFcIfInventory) UnsetSpVhbas()`

UnsetSpVhbas ensures that no value is present for SpVhbas, not even an explicit nil
### GetWwpnLease

`func (o *VnicFcIfInventory) GetWwpnLease() FcpoolLeaseRelationship`

GetWwpnLease returns the WwpnLease field if non-nil, zero value otherwise.

### GetWwpnLeaseOk

`func (o *VnicFcIfInventory) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwpnLeaseOk returns a tuple with the WwpnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnLease

`func (o *VnicFcIfInventory) SetWwpnLease(v FcpoolLeaseRelationship)`

SetWwpnLease sets WwpnLease field to given value.

### HasWwpnLease

`func (o *VnicFcIfInventory) HasWwpnLease() bool`

HasWwpnLease returns a boolean if a field has been set.

### GetWwpnPool

`func (o *VnicFcIfInventory) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicFcIfInventory) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicFcIfInventory) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicFcIfInventory) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


