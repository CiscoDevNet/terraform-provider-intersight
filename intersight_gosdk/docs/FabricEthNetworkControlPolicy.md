# FabricEthNetworkControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.EthNetworkControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.EthNetworkControlPolicy"]
**CdpEnabled** | Pointer to **bool** | Enables the CDP on an interface. | [optional] [default to false]
**ForgeMac** | Pointer to **string** | Determines if the MAC forging is allowed or denied on an interface. * &#x60;allow&#x60; - Allows mac forging on an interface. * &#x60;deny&#x60; - Denies mac forging on an interface. | [optional] [default to "allow"]
**LldpSettings** | Pointer to [**NullableFabricLldpSettings**](FabricLldpSettings.md) |  | [optional] 
**MacRegistrationMode** | Pointer to **string** | Determines the MAC addresses that have to be registered with the switch. * &#x60;nativeVlanOnly&#x60; - Register only the MAC addresses learnt on the native VLAN. * &#x60;allVlans&#x60; - Register all the MAC addresses learnt on all the allowed VLANs. | [optional] [default to "nativeVlanOnly"]
**UplinkFailAction** | Pointer to **string** | Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * &#x60;linkDown&#x60; - The vethernet will go down in case a suitable uplink is not pinned. * &#x60;warning&#x60; - The vethernet will remain up even if a suitable uplink is not pinned. | [optional] [default to "linkDown"]
**NetworkPolicy** | Pointer to [**[]VnicEthNetworkPolicyRelationship**](VnicEthNetworkPolicyRelationship.md) | An array of relationships to vnicEthNetworkPolicy resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkControlPolicy

`func NewFabricEthNetworkControlPolicy(classId string, objectType string, ) *FabricEthNetworkControlPolicy`

NewFabricEthNetworkControlPolicy instantiates a new FabricEthNetworkControlPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkControlPolicyWithDefaults

`func NewFabricEthNetworkControlPolicyWithDefaults() *FabricEthNetworkControlPolicy`

NewFabricEthNetworkControlPolicyWithDefaults instantiates a new FabricEthNetworkControlPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricEthNetworkControlPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricEthNetworkControlPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricEthNetworkControlPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricEthNetworkControlPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricEthNetworkControlPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricEthNetworkControlPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdpEnabled

`func (o *FabricEthNetworkControlPolicy) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *FabricEthNetworkControlPolicy) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *FabricEthNetworkControlPolicy) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *FabricEthNetworkControlPolicy) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### GetForgeMac

`func (o *FabricEthNetworkControlPolicy) GetForgeMac() string`

GetForgeMac returns the ForgeMac field if non-nil, zero value otherwise.

### GetForgeMacOk

`func (o *FabricEthNetworkControlPolicy) GetForgeMacOk() (*string, bool)`

GetForgeMacOk returns a tuple with the ForgeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgeMac

`func (o *FabricEthNetworkControlPolicy) SetForgeMac(v string)`

SetForgeMac sets ForgeMac field to given value.

### HasForgeMac

`func (o *FabricEthNetworkControlPolicy) HasForgeMac() bool`

HasForgeMac returns a boolean if a field has been set.

### GetLldpSettings

`func (o *FabricEthNetworkControlPolicy) GetLldpSettings() FabricLldpSettings`

GetLldpSettings returns the LldpSettings field if non-nil, zero value otherwise.

### GetLldpSettingsOk

`func (o *FabricEthNetworkControlPolicy) GetLldpSettingsOk() (*FabricLldpSettings, bool)`

GetLldpSettingsOk returns a tuple with the LldpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSettings

`func (o *FabricEthNetworkControlPolicy) SetLldpSettings(v FabricLldpSettings)`

SetLldpSettings sets LldpSettings field to given value.

### HasLldpSettings

`func (o *FabricEthNetworkControlPolicy) HasLldpSettings() bool`

HasLldpSettings returns a boolean if a field has been set.

### SetLldpSettingsNil

`func (o *FabricEthNetworkControlPolicy) SetLldpSettingsNil(b bool)`

 SetLldpSettingsNil sets the value for LldpSettings to be an explicit nil

### UnsetLldpSettings
`func (o *FabricEthNetworkControlPolicy) UnsetLldpSettings()`

UnsetLldpSettings ensures that no value is present for LldpSettings, not even an explicit nil
### GetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) GetMacRegistrationMode() string`

GetMacRegistrationMode returns the MacRegistrationMode field if non-nil, zero value otherwise.

### GetMacRegistrationModeOk

`func (o *FabricEthNetworkControlPolicy) GetMacRegistrationModeOk() (*string, bool)`

GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) SetMacRegistrationMode(v string)`

SetMacRegistrationMode sets MacRegistrationMode field to given value.

### HasMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) HasMacRegistrationMode() bool`

HasMacRegistrationMode returns a boolean if a field has been set.

### GetUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) GetUplinkFailAction() string`

GetUplinkFailAction returns the UplinkFailAction field if non-nil, zero value otherwise.

### GetUplinkFailActionOk

`func (o *FabricEthNetworkControlPolicy) GetUplinkFailActionOk() (*string, bool)`

GetUplinkFailActionOk returns a tuple with the UplinkFailAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) SetUplinkFailAction(v string)`

SetUplinkFailAction sets UplinkFailAction field to given value.

### HasUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) HasUplinkFailAction() bool`

HasUplinkFailAction returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) GetNetworkPolicy() []VnicEthNetworkPolicyRelationship`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *FabricEthNetworkControlPolicy) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyRelationship, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) SetNetworkPolicy(v []VnicEthNetworkPolicyRelationship)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### SetNetworkPolicyNil

`func (o *FabricEthNetworkControlPolicy) SetNetworkPolicyNil(b bool)`

 SetNetworkPolicyNil sets the value for NetworkPolicy to be an explicit nil

### UnsetNetworkPolicy
`func (o *FabricEthNetworkControlPolicy) UnsetNetworkPolicy()`

UnsetNetworkPolicy ensures that no value is present for NetworkPolicy, not even an explicit nil
### GetOrganization

`func (o *FabricEthNetworkControlPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricEthNetworkControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricEthNetworkControlPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricEthNetworkControlPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


