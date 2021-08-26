# VnicFcNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcNetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcNetworkPolicy"]
**VsanSettings** | Pointer to [**NullableVnicVsanSettings**](VnicVsanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicFcNetworkPolicy

`func NewVnicFcNetworkPolicy(classId string, objectType string, ) *VnicFcNetworkPolicy`

NewVnicFcNetworkPolicy instantiates a new VnicFcNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcNetworkPolicyWithDefaults

`func NewVnicFcNetworkPolicyWithDefaults() *VnicFcNetworkPolicy`

NewVnicFcNetworkPolicyWithDefaults instantiates a new VnicFcNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcNetworkPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcNetworkPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcNetworkPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcNetworkPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcNetworkPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcNetworkPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVsanSettings

`func (o *VnicFcNetworkPolicy) GetVsanSettings() VnicVsanSettings`

GetVsanSettings returns the VsanSettings field if non-nil, zero value otherwise.

### GetVsanSettingsOk

`func (o *VnicFcNetworkPolicy) GetVsanSettingsOk() (*VnicVsanSettings, bool)`

GetVsanSettingsOk returns a tuple with the VsanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanSettings

`func (o *VnicFcNetworkPolicy) SetVsanSettings(v VnicVsanSettings)`

SetVsanSettings sets VsanSettings field to given value.

### HasVsanSettings

`func (o *VnicFcNetworkPolicy) HasVsanSettings() bool`

HasVsanSettings returns a boolean if a field has been set.

### SetVsanSettingsNil

`func (o *VnicFcNetworkPolicy) SetVsanSettingsNil(b bool)`

 SetVsanSettingsNil sets the value for VsanSettings to be an explicit nil

### UnsetVsanSettings
`func (o *VnicFcNetworkPolicy) UnsetVsanSettings()`

UnsetVsanSettings ensures that no value is present for VsanSettings, not even an explicit nil
### GetOrganization

`func (o *VnicFcNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicFcNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicFcNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicFcNetworkPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


