# VnicFcIfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcIf"]
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two. | [optional] 
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StaticWwpnAddress** | Pointer to **string** | The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**VifId** | Pointer to **int64** | This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy. | [optional] [readonly] 
**WwpnAddressType** | Pointer to **string** | Type of allocation selected to assign a WWPN address to the vhba. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [default to "POOL"]
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**SanConnectivityPolicy** | Pointer to [**VnicSanConnectivityPolicyRelationship**](VnicSanConnectivityPolicyRelationship.md) |  | [optional] 
**ScpVhba** | Pointer to [**VnicFcIfRelationship**](VnicFcIfRelationship.md) |  | [optional] 
**SpVhbas** | Pointer to [**[]VnicFcIfRelationship**](VnicFcIfRelationship.md) | An array of relationships to vnicFcIf resources. | [optional] [readonly] 
**SrcTemplate** | Pointer to [**VnicVhbaTemplateRelationship**](VnicVhbaTemplateRelationship.md) |  | [optional] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 

## Methods

### NewVnicFcIfAllOf

`func NewVnicFcIfAllOf(classId string, objectType string, ) *VnicFcIfAllOf`

NewVnicFcIfAllOf instantiates a new VnicFcIfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfAllOfWithDefaults

`func NewVnicFcIfAllOfWithDefaults() *VnicFcIfAllOf`

NewVnicFcIfAllOfWithDefaults instantiates a new VnicFcIfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcIfAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcIfAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcIfAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcIfAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcIfAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcIfAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetOverriddenList

`func (o *VnicFcIfAllOf) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *VnicFcIfAllOf) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *VnicFcIfAllOf) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *VnicFcIfAllOf) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *VnicFcIfAllOf) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *VnicFcIfAllOf) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
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

### SetPlacementNil

`func (o *VnicFcIfAllOf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicFcIfAllOf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetStaticWwpnAddress

`func (o *VnicFcIfAllOf) GetStaticWwpnAddress() string`

GetStaticWwpnAddress returns the StaticWwpnAddress field if non-nil, zero value otherwise.

### GetStaticWwpnAddressOk

`func (o *VnicFcIfAllOf) GetStaticWwpnAddressOk() (*string, bool)`

GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticWwpnAddress

`func (o *VnicFcIfAllOf) SetStaticWwpnAddress(v string)`

SetStaticWwpnAddress sets StaticWwpnAddress field to given value.

### HasStaticWwpnAddress

`func (o *VnicFcIfAllOf) HasStaticWwpnAddress() bool`

HasStaticWwpnAddress returns a boolean if a field has been set.

### GetTemplateActions

`func (o *VnicFcIfAllOf) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicFcIfAllOf) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicFcIfAllOf) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicFcIfAllOf) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicFcIfAllOf) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicFcIfAllOf) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *VnicFcIfAllOf) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *VnicFcIfAllOf) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *VnicFcIfAllOf) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *VnicFcIfAllOf) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *VnicFcIfAllOf) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *VnicFcIfAllOf) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *VnicFcIfAllOf) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *VnicFcIfAllOf) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *VnicFcIfAllOf) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *VnicFcIfAllOf) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

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

### GetWwpnAddressType

`func (o *VnicFcIfAllOf) GetWwpnAddressType() string`

GetWwpnAddressType returns the WwpnAddressType field if non-nil, zero value otherwise.

### GetWwpnAddressTypeOk

`func (o *VnicFcIfAllOf) GetWwpnAddressTypeOk() (*string, bool)`

GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnAddressType

`func (o *VnicFcIfAllOf) SetWwpnAddressType(v string)`

SetWwpnAddressType sets WwpnAddressType field to given value.

### HasWwpnAddressType

`func (o *VnicFcIfAllOf) HasWwpnAddressType() bool`

HasWwpnAddressType returns a boolean if a field has been set.

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
### GetSrcTemplate

`func (o *VnicFcIfAllOf) GetSrcTemplate() VnicVhbaTemplateRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *VnicFcIfAllOf) GetSrcTemplateOk() (*VnicVhbaTemplateRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *VnicFcIfAllOf) SetSrcTemplate(v VnicVhbaTemplateRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *VnicFcIfAllOf) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


