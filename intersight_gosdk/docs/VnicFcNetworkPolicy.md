# VnicFcNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VsanSettings** | Pointer to [**VnicVsanSettings**](vnic.VsanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicFcNetworkPolicy

`func NewVnicFcNetworkPolicy() *VnicFcNetworkPolicy`

NewVnicFcNetworkPolicy instantiates a new VnicFcNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcNetworkPolicyWithDefaults

`func NewVnicFcNetworkPolicyWithDefaults() *VnicFcNetworkPolicy`

NewVnicFcNetworkPolicyWithDefaults instantiates a new VnicFcNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


