# VnicFcIfInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcIfInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcIfInventory"]
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] [readonly] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two. | [optional] [readonly] 
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] [readonly] 
**PinGroupName** | Pointer to **string** | Pingroup name associated to vfc for static pinning. SCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vfc traffic. | [optional] [readonly] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StaticWwpnAddress** | Pointer to **string** | The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. | [optional] [readonly] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
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
**SrcTemplate** | Pointer to [**VnicVhbaTemplateRelationship**](VnicVhbaTemplateRelationship.md) |  | [optional] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 
**WwpnPool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicFcIfInventoryAllOf

`func NewVnicFcIfInventoryAllOf(classId string, objectType string, ) *VnicFcIfInventoryAllOf`

NewVnicFcIfInventoryAllOf instantiates a new VnicFcIfInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfInventoryAllOfWithDefaults

`func NewVnicFcIfInventoryAllOfWithDefaults() *VnicFcIfInventoryAllOf`

NewVnicFcIfInventoryAllOfWithDefaults instantiates a new VnicFcIfInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcIfInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcIfInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcIfInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcIfInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcIfInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcIfInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VnicFcIfInventoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcIfInventoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcIfInventoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcIfInventoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicFcIfInventoryAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicFcIfInventoryAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicFcIfInventoryAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicFcIfInventoryAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOverriddenList

`func (o *VnicFcIfInventoryAllOf) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *VnicFcIfInventoryAllOf) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *VnicFcIfInventoryAllOf) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *VnicFcIfInventoryAllOf) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *VnicFcIfInventoryAllOf) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *VnicFcIfInventoryAllOf) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
### GetPersistentBindings

`func (o *VnicFcIfInventoryAllOf) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicFcIfInventoryAllOf) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicFcIfInventoryAllOf) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicFcIfInventoryAllOf) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicFcIfInventoryAllOf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicFcIfInventoryAllOf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicFcIfInventoryAllOf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicFcIfInventoryAllOf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicFcIfInventoryAllOf) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicFcIfInventoryAllOf) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicFcIfInventoryAllOf) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicFcIfInventoryAllOf) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *VnicFcIfInventoryAllOf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicFcIfInventoryAllOf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStaticWwpnAddress

`func (o *VnicFcIfInventoryAllOf) GetStaticWwpnAddress() string`

GetStaticWwpnAddress returns the StaticWwpnAddress field if non-nil, zero value otherwise.

### GetStaticWwpnAddressOk

`func (o *VnicFcIfInventoryAllOf) GetStaticWwpnAddressOk() (*string, bool)`

GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticWwpnAddress

`func (o *VnicFcIfInventoryAllOf) SetStaticWwpnAddress(v string)`

SetStaticWwpnAddress sets StaticWwpnAddress field to given value.

### HasStaticWwpnAddress

`func (o *VnicFcIfInventoryAllOf) HasStaticWwpnAddress() bool`

HasStaticWwpnAddress returns a boolean if a field has been set.

### GetTemplateActions

`func (o *VnicFcIfInventoryAllOf) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicFcIfInventoryAllOf) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicFcIfInventoryAllOf) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicFcIfInventoryAllOf) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicFcIfInventoryAllOf) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicFcIfInventoryAllOf) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *VnicFcIfInventoryAllOf) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *VnicFcIfInventoryAllOf) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *VnicFcIfInventoryAllOf) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *VnicFcIfInventoryAllOf) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *VnicFcIfInventoryAllOf) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *VnicFcIfInventoryAllOf) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *VnicFcIfInventoryAllOf) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *VnicFcIfInventoryAllOf) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *VnicFcIfInventoryAllOf) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *VnicFcIfInventoryAllOf) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

### GetType

`func (o *VnicFcIfInventoryAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicFcIfInventoryAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicFcIfInventoryAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicFcIfInventoryAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVifId

`func (o *VnicFcIfInventoryAllOf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicFcIfInventoryAllOf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicFcIfInventoryAllOf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicFcIfInventoryAllOf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetWwpn

`func (o *VnicFcIfInventoryAllOf) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *VnicFcIfInventoryAllOf) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *VnicFcIfInventoryAllOf) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *VnicFcIfInventoryAllOf) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetWwpnAddressType

`func (o *VnicFcIfInventoryAllOf) GetWwpnAddressType() string`

GetWwpnAddressType returns the WwpnAddressType field if non-nil, zero value otherwise.

### GetWwpnAddressTypeOk

`func (o *VnicFcIfInventoryAllOf) GetWwpnAddressTypeOk() (*string, bool)`

GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnAddressType

`func (o *VnicFcIfInventoryAllOf) SetWwpnAddressType(v string)`

SetWwpnAddressType sets WwpnAddressType field to given value.

### HasWwpnAddressType

`func (o *VnicFcIfInventoryAllOf) HasWwpnAddressType() bool`

HasWwpnAddressType returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicFcIfInventoryAllOf) GetFcAdapterPolicy() VnicFcAdapterPolicyInventoryRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicFcIfInventoryAllOf) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyInventoryRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicFcIfInventoryAllOf) SetFcAdapterPolicy(v VnicFcAdapterPolicyInventoryRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicFcIfInventoryAllOf) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *VnicFcIfInventoryAllOf) GetFcNetworkPolicy() VnicFcNetworkPolicyInventoryRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicFcIfInventoryAllOf) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyInventoryRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicFcIfInventoryAllOf) SetFcNetworkPolicy(v VnicFcNetworkPolicyInventoryRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicFcIfInventoryAllOf) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### GetFcQosPolicy

`func (o *VnicFcIfInventoryAllOf) GetFcQosPolicy() VnicFcQosPolicyInventoryRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicFcIfInventoryAllOf) GetFcQosPolicyOk() (*VnicFcQosPolicyInventoryRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicFcIfInventoryAllOf) SetFcQosPolicy(v VnicFcQosPolicyInventoryRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicFcIfInventoryAllOf) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### GetFcZonePolicies

`func (o *VnicFcIfInventoryAllOf) GetFcZonePolicies() []FabricFcZonePolicyRelationship`

GetFcZonePolicies returns the FcZonePolicies field if non-nil, zero value otherwise.

### GetFcZonePoliciesOk

`func (o *VnicFcIfInventoryAllOf) GetFcZonePoliciesOk() (*[]FabricFcZonePolicyRelationship, bool)`

GetFcZonePoliciesOk returns a tuple with the FcZonePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcZonePolicies

`func (o *VnicFcIfInventoryAllOf) SetFcZonePolicies(v []FabricFcZonePolicyRelationship)`

SetFcZonePolicies sets FcZonePolicies field to given value.

### HasFcZonePolicies

`func (o *VnicFcIfInventoryAllOf) HasFcZonePolicies() bool`

HasFcZonePolicies returns a boolean if a field has been set.

### SetFcZonePoliciesNil

`func (o *VnicFcIfInventoryAllOf) SetFcZonePoliciesNil(b bool)`

 SetFcZonePoliciesNil sets the value for FcZonePolicies to be an explicit nil

### UnsetFcZonePolicies
`func (o *VnicFcIfInventoryAllOf) UnsetFcZonePolicies()`

UnsetFcZonePolicies ensures that no value is present for FcZonePolicies, not even an explicit nil
### GetSanConnectivityPolicy

`func (o *VnicFcIfInventoryAllOf) GetSanConnectivityPolicy() VnicSanConnectivityPolicyInventoryRelationship`

GetSanConnectivityPolicy returns the SanConnectivityPolicy field if non-nil, zero value otherwise.

### GetSanConnectivityPolicyOk

`func (o *VnicFcIfInventoryAllOf) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyInventoryRelationship, bool)`

GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanConnectivityPolicy

`func (o *VnicFcIfInventoryAllOf) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyInventoryRelationship)`

SetSanConnectivityPolicy sets SanConnectivityPolicy field to given value.

### HasSanConnectivityPolicy

`func (o *VnicFcIfInventoryAllOf) HasSanConnectivityPolicy() bool`

HasSanConnectivityPolicy returns a boolean if a field has been set.

### GetScpVhba

`func (o *VnicFcIfInventoryAllOf) GetScpVhba() VnicFcIfInventoryRelationship`

GetScpVhba returns the ScpVhba field if non-nil, zero value otherwise.

### GetScpVhbaOk

`func (o *VnicFcIfInventoryAllOf) GetScpVhbaOk() (*VnicFcIfInventoryRelationship, bool)`

GetScpVhbaOk returns a tuple with the ScpVhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpVhba

`func (o *VnicFcIfInventoryAllOf) SetScpVhba(v VnicFcIfInventoryRelationship)`

SetScpVhba sets ScpVhba field to given value.

### HasScpVhba

`func (o *VnicFcIfInventoryAllOf) HasScpVhba() bool`

HasScpVhba returns a boolean if a field has been set.

### GetSpVhbas

`func (o *VnicFcIfInventoryAllOf) GetSpVhbas() []VnicFcIfInventoryRelationship`

GetSpVhbas returns the SpVhbas field if non-nil, zero value otherwise.

### GetSpVhbasOk

`func (o *VnicFcIfInventoryAllOf) GetSpVhbasOk() (*[]VnicFcIfInventoryRelationship, bool)`

GetSpVhbasOk returns a tuple with the SpVhbas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVhbas

`func (o *VnicFcIfInventoryAllOf) SetSpVhbas(v []VnicFcIfInventoryRelationship)`

SetSpVhbas sets SpVhbas field to given value.

### HasSpVhbas

`func (o *VnicFcIfInventoryAllOf) HasSpVhbas() bool`

HasSpVhbas returns a boolean if a field has been set.

### SetSpVhbasNil

`func (o *VnicFcIfInventoryAllOf) SetSpVhbasNil(b bool)`

 SetSpVhbasNil sets the value for SpVhbas to be an explicit nil

### UnsetSpVhbas
`func (o *VnicFcIfInventoryAllOf) UnsetSpVhbas()`

UnsetSpVhbas ensures that no value is present for SpVhbas, not even an explicit nil
### GetSrcTemplate

`func (o *VnicFcIfInventoryAllOf) GetSrcTemplate() VnicVhbaTemplateRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *VnicFcIfInventoryAllOf) GetSrcTemplateOk() (*VnicVhbaTemplateRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *VnicFcIfInventoryAllOf) SetSrcTemplate(v VnicVhbaTemplateRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *VnicFcIfInventoryAllOf) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.

### GetWwpnLease

`func (o *VnicFcIfInventoryAllOf) GetWwpnLease() FcpoolLeaseRelationship`

GetWwpnLease returns the WwpnLease field if non-nil, zero value otherwise.

### GetWwpnLeaseOk

`func (o *VnicFcIfInventoryAllOf) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwpnLeaseOk returns a tuple with the WwpnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnLease

`func (o *VnicFcIfInventoryAllOf) SetWwpnLease(v FcpoolLeaseRelationship)`

SetWwpnLease sets WwpnLease field to given value.

### HasWwpnLease

`func (o *VnicFcIfInventoryAllOf) HasWwpnLease() bool`

HasWwpnLease returns a boolean if a field has been set.

### GetWwpnPool

`func (o *VnicFcIfInventoryAllOf) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicFcIfInventoryAllOf) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicFcIfInventoryAllOf) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicFcIfInventoryAllOf) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


