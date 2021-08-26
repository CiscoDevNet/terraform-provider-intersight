# FabricLinkControlPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LinkControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LinkControlPolicy"]
**UdldSettings** | Pointer to [**NullableFabricUdldSettings**](FabricUdldSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricLinkControlPolicyAllOf

`func NewFabricLinkControlPolicyAllOf(classId string, objectType string, ) *FabricLinkControlPolicyAllOf`

NewFabricLinkControlPolicyAllOf instantiates a new FabricLinkControlPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLinkControlPolicyAllOfWithDefaults

`func NewFabricLinkControlPolicyAllOfWithDefaults() *FabricLinkControlPolicyAllOf`

NewFabricLinkControlPolicyAllOfWithDefaults instantiates a new FabricLinkControlPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLinkControlPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLinkControlPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLinkControlPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLinkControlPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLinkControlPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLinkControlPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUdldSettings

`func (o *FabricLinkControlPolicyAllOf) GetUdldSettings() FabricUdldSettings`

GetUdldSettings returns the UdldSettings field if non-nil, zero value otherwise.

### GetUdldSettingsOk

`func (o *FabricLinkControlPolicyAllOf) GetUdldSettingsOk() (*FabricUdldSettings, bool)`

GetUdldSettingsOk returns a tuple with the UdldSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldSettings

`func (o *FabricLinkControlPolicyAllOf) SetUdldSettings(v FabricUdldSettings)`

SetUdldSettings sets UdldSettings field to given value.

### HasUdldSettings

`func (o *FabricLinkControlPolicyAllOf) HasUdldSettings() bool`

HasUdldSettings returns a boolean if a field has been set.

### SetUdldSettingsNil

`func (o *FabricLinkControlPolicyAllOf) SetUdldSettingsNil(b bool)`

 SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil

### UnsetUdldSettings
`func (o *FabricLinkControlPolicyAllOf) UnsetUdldSettings()`

UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
### GetOrganization

`func (o *FabricLinkControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricLinkControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricLinkControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricLinkControlPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


