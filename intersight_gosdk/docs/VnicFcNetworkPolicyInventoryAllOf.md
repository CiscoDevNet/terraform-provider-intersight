# VnicFcNetworkPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcNetworkPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcNetworkPolicyInventory"]
**VsanSettings** | Pointer to [**NullableVnicVsanSettings**](VnicVsanSettings.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicFcNetworkPolicyInventoryAllOf

`func NewVnicFcNetworkPolicyInventoryAllOf(classId string, objectType string, ) *VnicFcNetworkPolicyInventoryAllOf`

NewVnicFcNetworkPolicyInventoryAllOf instantiates a new VnicFcNetworkPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcNetworkPolicyInventoryAllOfWithDefaults

`func NewVnicFcNetworkPolicyInventoryAllOfWithDefaults() *VnicFcNetworkPolicyInventoryAllOf`

NewVnicFcNetworkPolicyInventoryAllOfWithDefaults instantiates a new VnicFcNetworkPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcNetworkPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcNetworkPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVsanSettings

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetVsanSettings() VnicVsanSettings`

GetVsanSettings returns the VsanSettings field if non-nil, zero value otherwise.

### GetVsanSettingsOk

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetVsanSettingsOk() (*VnicVsanSettings, bool)`

GetVsanSettingsOk returns a tuple with the VsanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanSettings

`func (o *VnicFcNetworkPolicyInventoryAllOf) SetVsanSettings(v VnicVsanSettings)`

SetVsanSettings sets VsanSettings field to given value.

### HasVsanSettings

`func (o *VnicFcNetworkPolicyInventoryAllOf) HasVsanSettings() bool`

HasVsanSettings returns a boolean if a field has been set.

### SetVsanSettingsNil

`func (o *VnicFcNetworkPolicyInventoryAllOf) SetVsanSettingsNil(b bool)`

 SetVsanSettingsNil sets the value for VsanSettings to be an explicit nil

### UnsetVsanSettings
`func (o *VnicFcNetworkPolicyInventoryAllOf) UnsetVsanSettings()`

UnsetVsanSettings ensures that no value is present for VsanSettings, not even an explicit nil
### GetTargetMo

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicFcNetworkPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicFcNetworkPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicFcNetworkPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


