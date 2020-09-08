# FabricEthNetworkGroupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanSettings** | Pointer to [**FabricVlanSettings**](fabric.VlanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkGroupPolicy

`func NewFabricEthNetworkGroupPolicy() *FabricEthNetworkGroupPolicy`

NewFabricEthNetworkGroupPolicy instantiates a new FabricEthNetworkGroupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkGroupPolicyWithDefaults

`func NewFabricEthNetworkGroupPolicyWithDefaults() *FabricEthNetworkGroupPolicy`

NewFabricEthNetworkGroupPolicyWithDefaults instantiates a new FabricEthNetworkGroupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanSettings

`func (o *FabricEthNetworkGroupPolicy) GetVlanSettings() FabricVlanSettings`

GetVlanSettings returns the VlanSettings field if non-nil, zero value otherwise.

### GetVlanSettingsOk

`func (o *FabricEthNetworkGroupPolicy) GetVlanSettingsOk() (*FabricVlanSettings, bool)`

GetVlanSettingsOk returns a tuple with the VlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSettings

`func (o *FabricEthNetworkGroupPolicy) SetVlanSettings(v FabricVlanSettings)`

SetVlanSettings sets VlanSettings field to given value.

### HasVlanSettings

`func (o *FabricEthNetworkGroupPolicy) HasVlanSettings() bool`

HasVlanSettings returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricEthNetworkGroupPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricEthNetworkGroupPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricEthNetworkGroupPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricEthNetworkGroupPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


