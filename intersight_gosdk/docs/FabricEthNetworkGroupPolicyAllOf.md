# FabricEthNetworkGroupPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.EthNetworkGroupPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.EthNetworkGroupPolicy"]
**VlanSettings** | Pointer to [**NullableFabricVlanSettings**](FabricVlanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkGroupPolicyAllOf

`func NewFabricEthNetworkGroupPolicyAllOf(classId string, objectType string, ) *FabricEthNetworkGroupPolicyAllOf`

NewFabricEthNetworkGroupPolicyAllOf instantiates a new FabricEthNetworkGroupPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkGroupPolicyAllOfWithDefaults

`func NewFabricEthNetworkGroupPolicyAllOfWithDefaults() *FabricEthNetworkGroupPolicyAllOf`

NewFabricEthNetworkGroupPolicyAllOfWithDefaults instantiates a new FabricEthNetworkGroupPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricEthNetworkGroupPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricEthNetworkGroupPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricEthNetworkGroupPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricEthNetworkGroupPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricEthNetworkGroupPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricEthNetworkGroupPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanSettings

`func (o *FabricEthNetworkGroupPolicyAllOf) GetVlanSettings() FabricVlanSettings`

GetVlanSettings returns the VlanSettings field if non-nil, zero value otherwise.

### GetVlanSettingsOk

`func (o *FabricEthNetworkGroupPolicyAllOf) GetVlanSettingsOk() (*FabricVlanSettings, bool)`

GetVlanSettingsOk returns a tuple with the VlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSettings

`func (o *FabricEthNetworkGroupPolicyAllOf) SetVlanSettings(v FabricVlanSettings)`

SetVlanSettings sets VlanSettings field to given value.

### HasVlanSettings

`func (o *FabricEthNetworkGroupPolicyAllOf) HasVlanSettings() bool`

HasVlanSettings returns a boolean if a field has been set.

### SetVlanSettingsNil

`func (o *FabricEthNetworkGroupPolicyAllOf) SetVlanSettingsNil(b bool)`

 SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil

### UnsetVlanSettings
`func (o *FabricEthNetworkGroupPolicyAllOf) UnsetVlanSettings()`

UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
### GetOrganization

`func (o *FabricEthNetworkGroupPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricEthNetworkGroupPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricEthNetworkGroupPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricEthNetworkGroupPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


