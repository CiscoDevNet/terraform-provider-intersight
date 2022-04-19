# FabricEthNetworkGroupPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.EthNetworkGroupPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.EthNetworkGroupPolicyInventory"]
**VlanSettings** | Pointer to [**NullableFabricVlanSettings**](FabricVlanSettings.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkGroupPolicyInventoryAllOf

`func NewFabricEthNetworkGroupPolicyInventoryAllOf(classId string, objectType string, ) *FabricEthNetworkGroupPolicyInventoryAllOf`

NewFabricEthNetworkGroupPolicyInventoryAllOf instantiates a new FabricEthNetworkGroupPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkGroupPolicyInventoryAllOfWithDefaults

`func NewFabricEthNetworkGroupPolicyInventoryAllOfWithDefaults() *FabricEthNetworkGroupPolicyInventoryAllOf`

NewFabricEthNetworkGroupPolicyInventoryAllOfWithDefaults instantiates a new FabricEthNetworkGroupPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanSettings

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetVlanSettings() FabricVlanSettings`

GetVlanSettings returns the VlanSettings field if non-nil, zero value otherwise.

### GetVlanSettingsOk

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetVlanSettingsOk() (*FabricVlanSettings, bool)`

GetVlanSettingsOk returns a tuple with the VlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSettings

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) SetVlanSettings(v FabricVlanSettings)`

SetVlanSettings sets VlanSettings field to given value.

### HasVlanSettings

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) HasVlanSettings() bool`

HasVlanSettings returns a boolean if a field has been set.

### SetVlanSettingsNil

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) SetVlanSettingsNil(b bool)`

 SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil

### UnsetVlanSettings
`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) UnsetVlanSettings()`

UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
### GetTargetMo

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *FabricEthNetworkGroupPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


