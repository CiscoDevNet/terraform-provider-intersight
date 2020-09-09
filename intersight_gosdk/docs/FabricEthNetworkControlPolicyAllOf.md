# FabricEthNetworkControlPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdpEnabled** | Pointer to **bool** | Enables the CDP on an interface. | [optional] 
**ForgeMac** | Pointer to **string** | Determines if the MAC forging is allowed or denied on an interface. * &#x60;allow&#x60; - Allows mac forging on an interface. * &#x60;deny&#x60; - Denies mac forging on an interface. | [optional] [default to "allow"]
**LldpSettings** | Pointer to [**FabricLldpSettings**](fabric.LldpSettings.md) |  | [optional] 
**MacRegistrationMode** | Pointer to **string** | Determines the MAC addresses that have to be registered with the switch. * &#x60;nativeVlanOnly&#x60; - Register only the MAC addresses learnt on the native VLAN. * &#x60;allVlans&#x60; - Register all the MAC addresses learnt on all the allowed VLANs. | [optional] [default to "nativeVlanOnly"]
**UplinkFailAction** | Pointer to **string** | Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * &#x60;linkDown&#x60; - The vethernet will go down in case a suitable uplink is not pinned. * &#x60;warning&#x60; - The vethernet will remain up even if a suitable uplink is not pinned. | [optional] [default to "linkDown"]
**NetworkPolicy** | Pointer to [**[]VnicEthNetworkPolicyRelationship**](vnic.EthNetworkPolicy.Relationship.md) | An array of relationships to vnicEthNetworkPolicy resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkControlPolicyAllOf

`func NewFabricEthNetworkControlPolicyAllOf() *FabricEthNetworkControlPolicyAllOf`

NewFabricEthNetworkControlPolicyAllOf instantiates a new FabricEthNetworkControlPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkControlPolicyAllOfWithDefaults

`func NewFabricEthNetworkControlPolicyAllOfWithDefaults() *FabricEthNetworkControlPolicyAllOf`

NewFabricEthNetworkControlPolicyAllOfWithDefaults instantiates a new FabricEthNetworkControlPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdpEnabled

`func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *FabricEthNetworkControlPolicyAllOf) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *FabricEthNetworkControlPolicyAllOf) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### GetForgeMac

`func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMac() string`

GetForgeMac returns the ForgeMac field if non-nil, zero value otherwise.

### GetForgeMacOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMacOk() (*string, bool)`

GetForgeMacOk returns a tuple with the ForgeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgeMac

`func (o *FabricEthNetworkControlPolicyAllOf) SetForgeMac(v string)`

SetForgeMac sets ForgeMac field to given value.

### HasForgeMac

`func (o *FabricEthNetworkControlPolicyAllOf) HasForgeMac() bool`

HasForgeMac returns a boolean if a field has been set.

### GetLldpSettings

`func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettings() FabricLldpSettings`

GetLldpSettings returns the LldpSettings field if non-nil, zero value otherwise.

### GetLldpSettingsOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettingsOk() (*FabricLldpSettings, bool)`

GetLldpSettingsOk returns a tuple with the LldpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSettings

`func (o *FabricEthNetworkControlPolicyAllOf) SetLldpSettings(v FabricLldpSettings)`

SetLldpSettings sets LldpSettings field to given value.

### HasLldpSettings

`func (o *FabricEthNetworkControlPolicyAllOf) HasLldpSettings() bool`

HasLldpSettings returns a boolean if a field has been set.

### GetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegistrationMode() string`

GetMacRegistrationMode returns the MacRegistrationMode field if non-nil, zero value otherwise.

### GetMacRegistrationModeOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegistrationModeOk() (*string, bool)`

GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyAllOf) SetMacRegistrationMode(v string)`

SetMacRegistrationMode sets MacRegistrationMode field to given value.

### HasMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyAllOf) HasMacRegistrationMode() bool`

HasMacRegistrationMode returns a boolean if a field has been set.

### GetUplinkFailAction

`func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailAction() string`

GetUplinkFailAction returns the UplinkFailAction field if non-nil, zero value otherwise.

### GetUplinkFailActionOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailActionOk() (*string, bool)`

GetUplinkFailActionOk returns a tuple with the UplinkFailAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailAction

`func (o *FabricEthNetworkControlPolicyAllOf) SetUplinkFailAction(v string)`

SetUplinkFailAction sets UplinkFailAction field to given value.

### HasUplinkFailAction

`func (o *FabricEthNetworkControlPolicyAllOf) HasUplinkFailAction() bool`

HasUplinkFailAction returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicy() []VnicEthNetworkPolicyRelationship`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyRelationship, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *FabricEthNetworkControlPolicyAllOf) SetNetworkPolicy(v []VnicEthNetworkPolicyRelationship)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *FabricEthNetworkControlPolicyAllOf) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### SetNetworkPolicyNil

`func (o *FabricEthNetworkControlPolicyAllOf) SetNetworkPolicyNil(b bool)`

 SetNetworkPolicyNil sets the value for NetworkPolicy to be an explicit nil

### UnsetNetworkPolicy
`func (o *FabricEthNetworkControlPolicyAllOf) UnsetNetworkPolicy()`

UnsetNetworkPolicy ensures that no value is present for NetworkPolicy, not even an explicit nil
### GetOrganization

`func (o *FabricEthNetworkControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricEthNetworkControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricEthNetworkControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricEthNetworkControlPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


