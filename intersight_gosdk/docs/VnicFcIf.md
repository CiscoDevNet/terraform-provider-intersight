# VnicFcIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcIf"]
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StaticWwpnAddress** | Pointer to **string** | The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. | [optional] 
**Type** | Pointer to **string** | VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * &#x60;fc-initiator&#x60; - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-initiator&#x60; - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-target&#x60; - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * &#x60;fc-target&#x60; - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. | [optional] [default to "fc-initiator"]
**VifId** | Pointer to **int64** | This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy. | [optional] [readonly] 
**WwpnAddressType** | Pointer to **string** | Type of allocation selected to assign a WWPN address to the vhba. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [default to "POOL"]
**FcAdapterPolicy** | Pointer to [**VnicFcAdapterPolicyRelationship**](VnicFcAdapterPolicyRelationship.md) |  | [optional] 
**FcNetworkPolicy** | Pointer to [**VnicFcNetworkPolicyRelationship**](VnicFcNetworkPolicyRelationship.md) |  | [optional] 
**FcQosPolicy** | Pointer to [**VnicFcQosPolicyRelationship**](VnicFcQosPolicyRelationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**SanConnectivityPolicy** | Pointer to [**VnicSanConnectivityPolicyRelationship**](VnicSanConnectivityPolicyRelationship.md) |  | [optional] 
**ScpVhba** | Pointer to [**VnicFcIfRelationship**](VnicFcIfRelationship.md) |  | [optional] 
**SpVhbas** | Pointer to [**[]VnicFcIfRelationship**](VnicFcIfRelationship.md) | An array of relationships to vnicFcIf resources. | [optional] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 
**WwpnPool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicFcIf

`func NewVnicFcIf(classId string, objectType string, ) *VnicFcIf`

NewVnicFcIf instantiates a new VnicFcIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfWithDefaults

`func NewVnicFcIfWithDefaults() *VnicFcIf`

NewVnicFcIfWithDefaults instantiates a new VnicFcIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VnicFcIf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcIf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcIf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcIf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicFcIf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicFcIf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicFcIf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicFcIf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPersistentBindings

`func (o *VnicFcIf) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicFcIf) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicFcIf) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicFcIf) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicFcIf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicFcIf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicFcIf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicFcIf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicFcIf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicFcIf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStaticWwpnAddress

`func (o *VnicFcIf) GetStaticWwpnAddress() string`

GetStaticWwpnAddress returns the StaticWwpnAddress field if non-nil, zero value otherwise.

### GetStaticWwpnAddressOk

`func (o *VnicFcIf) GetStaticWwpnAddressOk() (*string, bool)`

GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticWwpnAddress

`func (o *VnicFcIf) SetStaticWwpnAddress(v string)`

SetStaticWwpnAddress sets StaticWwpnAddress field to given value.

### HasStaticWwpnAddress

`func (o *VnicFcIf) HasStaticWwpnAddress() bool`

HasStaticWwpnAddress returns a boolean if a field has been set.

### GetType

`func (o *VnicFcIf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicFcIf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicFcIf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicFcIf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVifId

`func (o *VnicFcIf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicFcIf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicFcIf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicFcIf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetWwpn

`func (o *VnicFcIf) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *VnicFcIf) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *VnicFcIf) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *VnicFcIf) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetWwpnAddressType

`func (o *VnicFcIf) GetWwpnAddressType() string`

GetWwpnAddressType returns the WwpnAddressType field if non-nil, zero value otherwise.

### GetWwpnAddressTypeOk

`func (o *VnicFcIf) GetWwpnAddressTypeOk() (*string, bool)`

GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnAddressType

`func (o *VnicFcIf) SetWwpnAddressType(v string)`

SetWwpnAddressType sets WwpnAddressType field to given value.

### HasWwpnAddressType

`func (o *VnicFcIf) HasWwpnAddressType() bool`

HasWwpnAddressType returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicFcIf) GetFcAdapterPolicy() VnicFcAdapterPolicyRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicFcIf) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicFcIf) SetFcAdapterPolicy(v VnicFcAdapterPolicyRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicFcIf) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *VnicFcIf) GetFcNetworkPolicy() VnicFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicFcIf) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicFcIf) SetFcNetworkPolicy(v VnicFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicFcIf) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### GetFcQosPolicy

`func (o *VnicFcIf) GetFcQosPolicy() VnicFcQosPolicyRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicFcIf) GetFcQosPolicyOk() (*VnicFcQosPolicyRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicFcIf) SetFcQosPolicy(v VnicFcQosPolicyRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicFcIf) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### GetProfile

`func (o *VnicFcIf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicFcIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicFcIf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicFcIf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSanConnectivityPolicy

`func (o *VnicFcIf) GetSanConnectivityPolicy() VnicSanConnectivityPolicyRelationship`

GetSanConnectivityPolicy returns the SanConnectivityPolicy field if non-nil, zero value otherwise.

### GetSanConnectivityPolicyOk

`func (o *VnicFcIf) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyRelationship, bool)`

GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanConnectivityPolicy

`func (o *VnicFcIf) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyRelationship)`

SetSanConnectivityPolicy sets SanConnectivityPolicy field to given value.

### HasSanConnectivityPolicy

`func (o *VnicFcIf) HasSanConnectivityPolicy() bool`

HasSanConnectivityPolicy returns a boolean if a field has been set.

### GetScpVhba

`func (o *VnicFcIf) GetScpVhba() VnicFcIfRelationship`

GetScpVhba returns the ScpVhba field if non-nil, zero value otherwise.

### GetScpVhbaOk

`func (o *VnicFcIf) GetScpVhbaOk() (*VnicFcIfRelationship, bool)`

GetScpVhbaOk returns a tuple with the ScpVhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpVhba

`func (o *VnicFcIf) SetScpVhba(v VnicFcIfRelationship)`

SetScpVhba sets ScpVhba field to given value.

### HasScpVhba

`func (o *VnicFcIf) HasScpVhba() bool`

HasScpVhba returns a boolean if a field has been set.

### GetSpVhbas

`func (o *VnicFcIf) GetSpVhbas() []VnicFcIfRelationship`

GetSpVhbas returns the SpVhbas field if non-nil, zero value otherwise.

### GetSpVhbasOk

`func (o *VnicFcIf) GetSpVhbasOk() (*[]VnicFcIfRelationship, bool)`

GetSpVhbasOk returns a tuple with the SpVhbas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVhbas

`func (o *VnicFcIf) SetSpVhbas(v []VnicFcIfRelationship)`

SetSpVhbas sets SpVhbas field to given value.

### HasSpVhbas

`func (o *VnicFcIf) HasSpVhbas() bool`

HasSpVhbas returns a boolean if a field has been set.

### SetSpVhbasNil

`func (o *VnicFcIf) SetSpVhbasNil(b bool)`

 SetSpVhbasNil sets the value for SpVhbas to be an explicit nil

### UnsetSpVhbas
`func (o *VnicFcIf) UnsetSpVhbas()`

UnsetSpVhbas ensures that no value is present for SpVhbas, not even an explicit nil
### GetWwpnLease

`func (o *VnicFcIf) GetWwpnLease() FcpoolLeaseRelationship`

GetWwpnLease returns the WwpnLease field if non-nil, zero value otherwise.

### GetWwpnLeaseOk

`func (o *VnicFcIf) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwpnLeaseOk returns a tuple with the WwpnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnLease

`func (o *VnicFcIf) SetWwpnLease(v FcpoolLeaseRelationship)`

SetWwpnLease sets WwpnLease field to given value.

### HasWwpnLease

`func (o *VnicFcIf) HasWwpnLease() bool`

HasWwpnLease returns a boolean if a field has been set.

### GetWwpnPool

`func (o *VnicFcIf) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicFcIf) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicFcIf) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicFcIf) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


