# VnicEthIfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthIf"]
**IscsiIpV4AddressAllocationType** | Pointer to **string** | Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type. * &#x60;None&#x60; - Type indicates that there is no IP associated to an vnic. * &#x60;DHCP&#x60; - The IP address is assigned using DHCP, if available. * &#x60;Static&#x60; - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * &#x60;Pool&#x60; - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. | [optional] [readonly] [default to "None"]
**IscsiIpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IscsiIpv4Address** | Pointer to **string** | IP address associated to the vNIC. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. | [optional] [readonly] 
**MacAddressType** | Pointer to **string** | Type of allocation selected to assign a MAC address for the vnic. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [default to "POOL"]
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two. | [optional] 
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**Placement** | Pointer to [**NullableVnicPlacementSettings**](VnicPlacementSettings.md) |  | [optional] 
**StandbyVifId** | Pointer to **int64** | The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. | [optional] [readonly] 
**StaticMacAddress** | Pointer to **string** | The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx. To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the following MAC prefix 00:25:B5:xx:xx:xx. | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**IpLease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**LanConnectivityPolicy** | Pointer to [**VnicLanConnectivityPolicyRelationship**](VnicLanConnectivityPolicyRelationship.md) |  | [optional] 
**LcpVnic** | Pointer to [**VnicEthIfRelationship**](VnicEthIfRelationship.md) |  | [optional] 
**MacLease** | Pointer to [**MacpoolLeaseRelationship**](MacpoolLeaseRelationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**SpVnics** | Pointer to [**[]VnicEthIfRelationship**](VnicEthIfRelationship.md) | An array of relationships to vnicEthIf resources. | [optional] [readonly] 
**SrcTemplate** | Pointer to [**VnicVnicTemplateRelationship**](VnicVnicTemplateRelationship.md) |  | [optional] 

## Methods

### NewVnicEthIfAllOf

`func NewVnicEthIfAllOf(classId string, objectType string, ) *VnicEthIfAllOf`

NewVnicEthIfAllOf instantiates a new VnicEthIfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfAllOfWithDefaults

`func NewVnicEthIfAllOfWithDefaults() *VnicEthIfAllOf`

NewVnicEthIfAllOfWithDefaults instantiates a new VnicEthIfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthIfAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthIfAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthIfAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthIfAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthIfAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthIfAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfAllOf) GetIscsiIpV4AddressAllocationType() string`

GetIscsiIpV4AddressAllocationType returns the IscsiIpV4AddressAllocationType field if non-nil, zero value otherwise.

### GetIscsiIpV4AddressAllocationTypeOk

`func (o *VnicEthIfAllOf) GetIscsiIpV4AddressAllocationTypeOk() (*string, bool)`

GetIscsiIpV4AddressAllocationTypeOk returns a tuple with the IscsiIpV4AddressAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4AddressAllocationType

`func (o *VnicEthIfAllOf) SetIscsiIpV4AddressAllocationType(v string)`

SetIscsiIpV4AddressAllocationType sets IscsiIpV4AddressAllocationType field to given value.

### HasIscsiIpV4AddressAllocationType

`func (o *VnicEthIfAllOf) HasIscsiIpV4AddressAllocationType() bool`

HasIscsiIpV4AddressAllocationType returns a boolean if a field has been set.

### GetIscsiIpV4Config

`func (o *VnicEthIfAllOf) GetIscsiIpV4Config() IppoolIpV4Config`

GetIscsiIpV4Config returns the IscsiIpV4Config field if non-nil, zero value otherwise.

### GetIscsiIpV4ConfigOk

`func (o *VnicEthIfAllOf) GetIscsiIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIscsiIpV4ConfigOk returns a tuple with the IscsiIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpV4Config

`func (o *VnicEthIfAllOf) SetIscsiIpV4Config(v IppoolIpV4Config)`

SetIscsiIpV4Config sets IscsiIpV4Config field to given value.

### HasIscsiIpV4Config

`func (o *VnicEthIfAllOf) HasIscsiIpV4Config() bool`

HasIscsiIpV4Config returns a boolean if a field has been set.

### SetIscsiIpV4ConfigNil

`func (o *VnicEthIfAllOf) SetIscsiIpV4ConfigNil(b bool)`

 SetIscsiIpV4ConfigNil sets the value for IscsiIpV4Config to be an explicit nil

### UnsetIscsiIpV4Config
`func (o *VnicEthIfAllOf) UnsetIscsiIpV4Config()`

UnsetIscsiIpV4Config ensures that no value is present for IscsiIpV4Config, not even an explicit nil
### GetIscsiIpv4Address

`func (o *VnicEthIfAllOf) GetIscsiIpv4Address() string`

GetIscsiIpv4Address returns the IscsiIpv4Address field if non-nil, zero value otherwise.

### GetIscsiIpv4AddressOk

`func (o *VnicEthIfAllOf) GetIscsiIpv4AddressOk() (*string, bool)`

GetIscsiIpv4AddressOk returns a tuple with the IscsiIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpv4Address

`func (o *VnicEthIfAllOf) SetIscsiIpv4Address(v string)`

SetIscsiIpv4Address sets IscsiIpv4Address field to given value.

### HasIscsiIpv4Address

`func (o *VnicEthIfAllOf) HasIscsiIpv4Address() bool`

HasIscsiIpv4Address returns a boolean if a field has been set.

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

### GetMacAddressType

`func (o *VnicEthIfAllOf) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *VnicEthIfAllOf) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *VnicEthIfAllOf) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *VnicEthIfAllOf) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

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

### GetOverriddenList

`func (o *VnicEthIfAllOf) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *VnicEthIfAllOf) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *VnicEthIfAllOf) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *VnicEthIfAllOf) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *VnicEthIfAllOf) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *VnicEthIfAllOf) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
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

### SetPlacementNil

`func (o *VnicEthIfAllOf) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *VnicEthIfAllOf) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
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

### GetStaticMacAddress

`func (o *VnicEthIfAllOf) GetStaticMacAddress() string`

GetStaticMacAddress returns the StaticMacAddress field if non-nil, zero value otherwise.

### GetStaticMacAddressOk

`func (o *VnicEthIfAllOf) GetStaticMacAddressOk() (*string, bool)`

GetStaticMacAddressOk returns a tuple with the StaticMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddress

`func (o *VnicEthIfAllOf) SetStaticMacAddress(v string)`

SetStaticMacAddress sets StaticMacAddress field to given value.

### HasStaticMacAddress

`func (o *VnicEthIfAllOf) HasStaticMacAddress() bool`

HasStaticMacAddress returns a boolean if a field has been set.

### GetTemplateActions

`func (o *VnicEthIfAllOf) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicEthIfAllOf) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicEthIfAllOf) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicEthIfAllOf) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicEthIfAllOf) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicEthIfAllOf) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *VnicEthIfAllOf) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *VnicEthIfAllOf) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *VnicEthIfAllOf) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *VnicEthIfAllOf) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *VnicEthIfAllOf) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *VnicEthIfAllOf) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *VnicEthIfAllOf) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *VnicEthIfAllOf) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *VnicEthIfAllOf) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *VnicEthIfAllOf) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

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

### GetIpLease

`func (o *VnicEthIfAllOf) GetIpLease() IppoolIpLeaseRelationship`

GetIpLease returns the IpLease field if non-nil, zero value otherwise.

### GetIpLeaseOk

`func (o *VnicEthIfAllOf) GetIpLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpLeaseOk returns a tuple with the IpLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLease

`func (o *VnicEthIfAllOf) SetIpLease(v IppoolIpLeaseRelationship)`

SetIpLease sets IpLease field to given value.

### HasIpLease

`func (o *VnicEthIfAllOf) HasIpLease() bool`

HasIpLease returns a boolean if a field has been set.

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
### GetSrcTemplate

`func (o *VnicEthIfAllOf) GetSrcTemplate() VnicVnicTemplateRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *VnicEthIfAllOf) GetSrcTemplateOk() (*VnicVnicTemplateRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *VnicEthIfAllOf) SetSrcTemplate(v VnicVnicTemplateRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *VnicEthIfAllOf) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


