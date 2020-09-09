# FabricPortPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | This field specifies the device model that this Port Policy is being configured for. * &#x60;UCS-FI-6454&#x60; - The standard 4th generation UCS Fabric Interconnect with 54 ports. * &#x60;UCS-FI-64108&#x60; - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * &#x60;unknown&#x60; - Unknown device type, usage is TBD. | [optional] [default to "UCS-FI-6454"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricPortPolicy

`func NewFabricPortPolicy() *FabricPortPolicy`

NewFabricPortPolicy instantiates a new FabricPortPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortPolicyWithDefaults

`func NewFabricPortPolicyWithDefaults() *FabricPortPolicy`

NewFabricPortPolicyWithDefaults instantiates a new FabricPortPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *FabricPortPolicy) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *FabricPortPolicy) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *FabricPortPolicy) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *FabricPortPolicy) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricPortPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricPortPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricPortPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricPortPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricPortPolicy) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricPortPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricPortPolicy) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricPortPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricPortPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricPortPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


