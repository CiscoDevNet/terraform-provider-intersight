# FabricEthNetworkControlPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.EthNetworkControlPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.EthNetworkControlPolicyInventory"]
**CdpEnabled** | Pointer to **bool** | Enables the CDP on an interface. | [optional] [readonly] [default to false]
**ForgeMac** | Pointer to **string** | Determines if the MAC forging is allowed or denied on an interface. * &#x60;allow&#x60; - Allows mac forging on an interface. * &#x60;deny&#x60; - Denies mac forging on an interface. | [optional] [readonly] [default to "allow"]
**LldpSettings** | Pointer to [**NullableFabricLldpSettings**](FabricLldpSettings.md) |  | [optional] 
**MacRegistrationMode** | Pointer to **string** | Determines the MAC addresses that have to be registered with the switch. * &#x60;nativeVlanOnly&#x60; - Register only the MAC addresses learnt on the native VLAN. * &#x60;allVlans&#x60; - Register all the MAC addresses learnt on all the allowed VLANs. | [optional] [readonly] [default to "nativeVlanOnly"]
**UplinkFailAction** | Pointer to **string** | Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * &#x60;linkDown&#x60; - The vethernet will go down in case a suitable uplink is not pinned. * &#x60;warning&#x60; - The vethernet will remain up even if a suitable uplink is not pinned. | [optional] [readonly] [default to "linkDown"]
**NetworkPolicy** | Pointer to [**[]VnicEthNetworkPolicyInventoryRelationship**](VnicEthNetworkPolicyInventoryRelationship.md) | An array of relationships to vnicEthNetworkPolicyInventory resources. | [optional] [readonly] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkControlPolicyInventoryAllOf

`func NewFabricEthNetworkControlPolicyInventoryAllOf(classId string, objectType string, ) *FabricEthNetworkControlPolicyInventoryAllOf`

NewFabricEthNetworkControlPolicyInventoryAllOf instantiates a new FabricEthNetworkControlPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkControlPolicyInventoryAllOfWithDefaults

`func NewFabricEthNetworkControlPolicyInventoryAllOfWithDefaults() *FabricEthNetworkControlPolicyInventoryAllOf`

NewFabricEthNetworkControlPolicyInventoryAllOfWithDefaults instantiates a new FabricEthNetworkControlPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdpEnabled

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### GetForgeMac

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetForgeMac() string`

GetForgeMac returns the ForgeMac field if non-nil, zero value otherwise.

### GetForgeMacOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetForgeMacOk() (*string, bool)`

GetForgeMacOk returns a tuple with the ForgeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgeMac

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetForgeMac(v string)`

SetForgeMac sets ForgeMac field to given value.

### HasForgeMac

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasForgeMac() bool`

HasForgeMac returns a boolean if a field has been set.

### GetLldpSettings

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetLldpSettings() FabricLldpSettings`

GetLldpSettings returns the LldpSettings field if non-nil, zero value otherwise.

### GetLldpSettingsOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetLldpSettingsOk() (*FabricLldpSettings, bool)`

GetLldpSettingsOk returns a tuple with the LldpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSettings

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetLldpSettings(v FabricLldpSettings)`

SetLldpSettings sets LldpSettings field to given value.

### HasLldpSettings

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasLldpSettings() bool`

HasLldpSettings returns a boolean if a field has been set.

### SetLldpSettingsNil

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetLldpSettingsNil(b bool)`

 SetLldpSettingsNil sets the value for LldpSettings to be an explicit nil

### UnsetLldpSettings
`func (o *FabricEthNetworkControlPolicyInventoryAllOf) UnsetLldpSettings()`

UnsetLldpSettings ensures that no value is present for LldpSettings, not even an explicit nil
### GetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetMacRegistrationMode() string`

GetMacRegistrationMode returns the MacRegistrationMode field if non-nil, zero value otherwise.

### GetMacRegistrationModeOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetMacRegistrationModeOk() (*string, bool)`

GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetMacRegistrationMode(v string)`

SetMacRegistrationMode sets MacRegistrationMode field to given value.

### HasMacRegistrationMode

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasMacRegistrationMode() bool`

HasMacRegistrationMode returns a boolean if a field has been set.

### GetUplinkFailAction

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetUplinkFailAction() string`

GetUplinkFailAction returns the UplinkFailAction field if non-nil, zero value otherwise.

### GetUplinkFailActionOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetUplinkFailActionOk() (*string, bool)`

GetUplinkFailActionOk returns a tuple with the UplinkFailAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailAction

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetUplinkFailAction(v string)`

SetUplinkFailAction sets UplinkFailAction field to given value.

### HasUplinkFailAction

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasUplinkFailAction() bool`

HasUplinkFailAction returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetNetworkPolicy() []VnicEthNetworkPolicyInventoryRelationship`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyInventoryRelationship, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetNetworkPolicy(v []VnicEthNetworkPolicyInventoryRelationship)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### SetNetworkPolicyNil

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetNetworkPolicyNil(b bool)`

 SetNetworkPolicyNil sets the value for NetworkPolicy to be an explicit nil

### UnsetNetworkPolicy
`func (o *FabricEthNetworkControlPolicyInventoryAllOf) UnsetNetworkPolicy()`

UnsetNetworkPolicy ensures that no value is present for NetworkPolicy, not even an explicit nil
### GetTargetMo

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *FabricEthNetworkControlPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


