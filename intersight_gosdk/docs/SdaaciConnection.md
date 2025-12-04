# SdaaciConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdaaci.Connection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdaaci.Connection"]
**AciL3Out** | Pointer to **string** | ACI L3Out Name User Input. | [optional] 
**AciMatchRuleName** | Pointer to **string** | Name of the Match Rule in Cisco APIC. | [optional] 
**AciTenant** | Pointer to **string** | ACI tenant Name for Selected APIC Target. | [optional] 
**CampusFabricSite** | Pointer to **string** | Campus fabric site id in which the border node has configured. | [optional] 
**Epg** | Pointer to **string** | Application EPG Name of this connection. | [optional] 
**EpgSubnet** | Pointer to **string** | EPG Subnet Ipv4Cidr which is configured on APIC. | [optional] 
**FirewallDevice** | Pointer to **string** | Device within the selected domain used to configure Firewall. | [optional] 
**FirewallDomain** | Pointer to **string** | Domain used to configure Firewall. | [optional] 
**NodeProfile** | Pointer to **string** | L3Out Node Profile in Cisco APIC. | [optional] 
**Status** | Pointer to **string** | Connection status between SDA and ACI. * &#x60;NotConnected&#x60; - Connection Status NotConnected. * &#x60;Connected&#x60; - Connection Status Connected. | [optional] [default to "NotConnected"]
**Transit** | Pointer to **string** | Transit id for given border node. | [optional] 
**VirtualNetwork** | Pointer to **string** | Virtual Network of this connection. | [optional] 
**VnEpg** | Pointer to **string** | Contains both VN and EPG of this connection. | [optional] 
**Vrf** | Pointer to **string** | APIC Tenant VRF from APIC. | [optional] 
**ApicTarget** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 
**CatalystCenterTarget** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 
**FmcTarget** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewSdaaciConnection

`func NewSdaaciConnection(classId string, objectType string, ) *SdaaciConnection`

NewSdaaciConnection instantiates a new SdaaciConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdaaciConnectionWithDefaults

`func NewSdaaciConnectionWithDefaults() *SdaaciConnection`

NewSdaaciConnectionWithDefaults instantiates a new SdaaciConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdaaciConnection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdaaciConnection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdaaciConnection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdaaciConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdaaciConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdaaciConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAciL3Out

`func (o *SdaaciConnection) GetAciL3Out() string`

GetAciL3Out returns the AciL3Out field if non-nil, zero value otherwise.

### GetAciL3OutOk

`func (o *SdaaciConnection) GetAciL3OutOk() (*string, bool)`

GetAciL3OutOk returns a tuple with the AciL3Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciL3Out

`func (o *SdaaciConnection) SetAciL3Out(v string)`

SetAciL3Out sets AciL3Out field to given value.

### HasAciL3Out

`func (o *SdaaciConnection) HasAciL3Out() bool`

HasAciL3Out returns a boolean if a field has been set.

### GetAciMatchRuleName

`func (o *SdaaciConnection) GetAciMatchRuleName() string`

GetAciMatchRuleName returns the AciMatchRuleName field if non-nil, zero value otherwise.

### GetAciMatchRuleNameOk

`func (o *SdaaciConnection) GetAciMatchRuleNameOk() (*string, bool)`

GetAciMatchRuleNameOk returns a tuple with the AciMatchRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciMatchRuleName

`func (o *SdaaciConnection) SetAciMatchRuleName(v string)`

SetAciMatchRuleName sets AciMatchRuleName field to given value.

### HasAciMatchRuleName

`func (o *SdaaciConnection) HasAciMatchRuleName() bool`

HasAciMatchRuleName returns a boolean if a field has been set.

### GetAciTenant

`func (o *SdaaciConnection) GetAciTenant() string`

GetAciTenant returns the AciTenant field if non-nil, zero value otherwise.

### GetAciTenantOk

`func (o *SdaaciConnection) GetAciTenantOk() (*string, bool)`

GetAciTenantOk returns a tuple with the AciTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciTenant

`func (o *SdaaciConnection) SetAciTenant(v string)`

SetAciTenant sets AciTenant field to given value.

### HasAciTenant

`func (o *SdaaciConnection) HasAciTenant() bool`

HasAciTenant returns a boolean if a field has been set.

### GetCampusFabricSite

`func (o *SdaaciConnection) GetCampusFabricSite() string`

GetCampusFabricSite returns the CampusFabricSite field if non-nil, zero value otherwise.

### GetCampusFabricSiteOk

`func (o *SdaaciConnection) GetCampusFabricSiteOk() (*string, bool)`

GetCampusFabricSiteOk returns a tuple with the CampusFabricSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusFabricSite

`func (o *SdaaciConnection) SetCampusFabricSite(v string)`

SetCampusFabricSite sets CampusFabricSite field to given value.

### HasCampusFabricSite

`func (o *SdaaciConnection) HasCampusFabricSite() bool`

HasCampusFabricSite returns a boolean if a field has been set.

### GetEpg

`func (o *SdaaciConnection) GetEpg() string`

GetEpg returns the Epg field if non-nil, zero value otherwise.

### GetEpgOk

`func (o *SdaaciConnection) GetEpgOk() (*string, bool)`

GetEpgOk returns a tuple with the Epg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpg

`func (o *SdaaciConnection) SetEpg(v string)`

SetEpg sets Epg field to given value.

### HasEpg

`func (o *SdaaciConnection) HasEpg() bool`

HasEpg returns a boolean if a field has been set.

### GetEpgSubnet

`func (o *SdaaciConnection) GetEpgSubnet() string`

GetEpgSubnet returns the EpgSubnet field if non-nil, zero value otherwise.

### GetEpgSubnetOk

`func (o *SdaaciConnection) GetEpgSubnetOk() (*string, bool)`

GetEpgSubnetOk returns a tuple with the EpgSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgSubnet

`func (o *SdaaciConnection) SetEpgSubnet(v string)`

SetEpgSubnet sets EpgSubnet field to given value.

### HasEpgSubnet

`func (o *SdaaciConnection) HasEpgSubnet() bool`

HasEpgSubnet returns a boolean if a field has been set.

### GetFirewallDevice

`func (o *SdaaciConnection) GetFirewallDevice() string`

GetFirewallDevice returns the FirewallDevice field if non-nil, zero value otherwise.

### GetFirewallDeviceOk

`func (o *SdaaciConnection) GetFirewallDeviceOk() (*string, bool)`

GetFirewallDeviceOk returns a tuple with the FirewallDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallDevice

`func (o *SdaaciConnection) SetFirewallDevice(v string)`

SetFirewallDevice sets FirewallDevice field to given value.

### HasFirewallDevice

`func (o *SdaaciConnection) HasFirewallDevice() bool`

HasFirewallDevice returns a boolean if a field has been set.

### GetFirewallDomain

`func (o *SdaaciConnection) GetFirewallDomain() string`

GetFirewallDomain returns the FirewallDomain field if non-nil, zero value otherwise.

### GetFirewallDomainOk

`func (o *SdaaciConnection) GetFirewallDomainOk() (*string, bool)`

GetFirewallDomainOk returns a tuple with the FirewallDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallDomain

`func (o *SdaaciConnection) SetFirewallDomain(v string)`

SetFirewallDomain sets FirewallDomain field to given value.

### HasFirewallDomain

`func (o *SdaaciConnection) HasFirewallDomain() bool`

HasFirewallDomain returns a boolean if a field has been set.

### GetNodeProfile

`func (o *SdaaciConnection) GetNodeProfile() string`

GetNodeProfile returns the NodeProfile field if non-nil, zero value otherwise.

### GetNodeProfileOk

`func (o *SdaaciConnection) GetNodeProfileOk() (*string, bool)`

GetNodeProfileOk returns a tuple with the NodeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeProfile

`func (o *SdaaciConnection) SetNodeProfile(v string)`

SetNodeProfile sets NodeProfile field to given value.

### HasNodeProfile

`func (o *SdaaciConnection) HasNodeProfile() bool`

HasNodeProfile returns a boolean if a field has been set.

### GetStatus

`func (o *SdaaciConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdaaciConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdaaciConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdaaciConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransit

`func (o *SdaaciConnection) GetTransit() string`

GetTransit returns the Transit field if non-nil, zero value otherwise.

### GetTransitOk

`func (o *SdaaciConnection) GetTransitOk() (*string, bool)`

GetTransitOk returns a tuple with the Transit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransit

`func (o *SdaaciConnection) SetTransit(v string)`

SetTransit sets Transit field to given value.

### HasTransit

`func (o *SdaaciConnection) HasTransit() bool`

HasTransit returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *SdaaciConnection) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *SdaaciConnection) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *SdaaciConnection) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *SdaaciConnection) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVnEpg

`func (o *SdaaciConnection) GetVnEpg() string`

GetVnEpg returns the VnEpg field if non-nil, zero value otherwise.

### GetVnEpgOk

`func (o *SdaaciConnection) GetVnEpgOk() (*string, bool)`

GetVnEpgOk returns a tuple with the VnEpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnEpg

`func (o *SdaaciConnection) SetVnEpg(v string)`

SetVnEpg sets VnEpg field to given value.

### HasVnEpg

`func (o *SdaaciConnection) HasVnEpg() bool`

HasVnEpg returns a boolean if a field has been set.

### GetVrf

`func (o *SdaaciConnection) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *SdaaciConnection) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *SdaaciConnection) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *SdaaciConnection) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetApicTarget

`func (o *SdaaciConnection) GetApicTarget() AssetTargetRelationship`

GetApicTarget returns the ApicTarget field if non-nil, zero value otherwise.

### GetApicTargetOk

`func (o *SdaaciConnection) GetApicTargetOk() (*AssetTargetRelationship, bool)`

GetApicTargetOk returns a tuple with the ApicTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicTarget

`func (o *SdaaciConnection) SetApicTarget(v AssetTargetRelationship)`

SetApicTarget sets ApicTarget field to given value.

### HasApicTarget

`func (o *SdaaciConnection) HasApicTarget() bool`

HasApicTarget returns a boolean if a field has been set.

### SetApicTargetNil

`func (o *SdaaciConnection) SetApicTargetNil(b bool)`

 SetApicTargetNil sets the value for ApicTarget to be an explicit nil

### UnsetApicTarget
`func (o *SdaaciConnection) UnsetApicTarget()`

UnsetApicTarget ensures that no value is present for ApicTarget, not even an explicit nil
### GetCatalystCenterTarget

`func (o *SdaaciConnection) GetCatalystCenterTarget() AssetTargetRelationship`

GetCatalystCenterTarget returns the CatalystCenterTarget field if non-nil, zero value otherwise.

### GetCatalystCenterTargetOk

`func (o *SdaaciConnection) GetCatalystCenterTargetOk() (*AssetTargetRelationship, bool)`

GetCatalystCenterTargetOk returns a tuple with the CatalystCenterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalystCenterTarget

`func (o *SdaaciConnection) SetCatalystCenterTarget(v AssetTargetRelationship)`

SetCatalystCenterTarget sets CatalystCenterTarget field to given value.

### HasCatalystCenterTarget

`func (o *SdaaciConnection) HasCatalystCenterTarget() bool`

HasCatalystCenterTarget returns a boolean if a field has been set.

### SetCatalystCenterTargetNil

`func (o *SdaaciConnection) SetCatalystCenterTargetNil(b bool)`

 SetCatalystCenterTargetNil sets the value for CatalystCenterTarget to be an explicit nil

### UnsetCatalystCenterTarget
`func (o *SdaaciConnection) UnsetCatalystCenterTarget()`

UnsetCatalystCenterTarget ensures that no value is present for CatalystCenterTarget, not even an explicit nil
### GetFmcTarget

`func (o *SdaaciConnection) GetFmcTarget() AssetTargetRelationship`

GetFmcTarget returns the FmcTarget field if non-nil, zero value otherwise.

### GetFmcTargetOk

`func (o *SdaaciConnection) GetFmcTargetOk() (*AssetTargetRelationship, bool)`

GetFmcTargetOk returns a tuple with the FmcTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFmcTarget

`func (o *SdaaciConnection) SetFmcTarget(v AssetTargetRelationship)`

SetFmcTarget sets FmcTarget field to given value.

### HasFmcTarget

`func (o *SdaaciConnection) HasFmcTarget() bool`

HasFmcTarget returns a boolean if a field has been set.

### SetFmcTargetNil

`func (o *SdaaciConnection) SetFmcTargetNil(b bool)`

 SetFmcTargetNil sets the value for FmcTarget to be an explicit nil

### UnsetFmcTarget
`func (o *SdaaciConnection) UnsetFmcTarget()`

UnsetFmcTarget ensures that no value is present for FmcTarget, not even an explicit nil
### GetOrganization

`func (o *SdaaciConnection) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdaaciConnection) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdaaciConnection) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdaaciConnection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *SdaaciConnection) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *SdaaciConnection) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


