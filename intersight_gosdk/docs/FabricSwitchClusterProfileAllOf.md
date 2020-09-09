# FabricSwitchClusterProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigContext** | Pointer to [**PolicyConfigContext**](policy.ConfigContext.md) |  | [optional] 
**SwitchProfilesCount** | Pointer to **int64** | Number of switch profiles that are part of this cluster profile. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**SwitchProfiles** | Pointer to [**[]FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfileAllOf

`func NewFabricSwitchClusterProfileAllOf() *FabricSwitchClusterProfileAllOf`

NewFabricSwitchClusterProfileAllOf instantiates a new FabricSwitchClusterProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileAllOfWithDefaults

`func NewFabricSwitchClusterProfileAllOfWithDefaults() *FabricSwitchClusterProfileAllOf`

NewFabricSwitchClusterProfileAllOfWithDefaults instantiates a new FabricSwitchClusterProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigContext

`func (o *FabricSwitchClusterProfileAllOf) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *FabricSwitchClusterProfileAllOf) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *FabricSwitchClusterProfileAllOf) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *FabricSwitchClusterProfileAllOf) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetSwitchProfilesCount

`func (o *FabricSwitchClusterProfileAllOf) GetSwitchProfilesCount() int64`

GetSwitchProfilesCount returns the SwitchProfilesCount field if non-nil, zero value otherwise.

### GetSwitchProfilesCountOk

`func (o *FabricSwitchClusterProfileAllOf) GetSwitchProfilesCountOk() (*int64, bool)`

GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilesCount

`func (o *FabricSwitchClusterProfileAllOf) SetSwitchProfilesCount(v int64)`

SetSwitchProfilesCount sets SwitchProfilesCount field to given value.

### HasSwitchProfilesCount

`func (o *FabricSwitchClusterProfileAllOf) HasSwitchProfilesCount() bool`

HasSwitchProfilesCount returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *FabricSwitchClusterProfileAllOf) GetSwitchProfiles() []FabricSwitchProfileRelationship`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *FabricSwitchClusterProfileAllOf) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *FabricSwitchClusterProfileAllOf) SetSwitchProfiles(v []FabricSwitchProfileRelationship)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *FabricSwitchClusterProfileAllOf) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### SetSwitchProfilesNil

`func (o *FabricSwitchClusterProfileAllOf) SetSwitchProfilesNil(b bool)`

 SetSwitchProfilesNil sets the value for SwitchProfiles to be an explicit nil

### UnsetSwitchProfiles
`func (o *FabricSwitchClusterProfileAllOf) UnsetSwitchProfiles()`

UnsetSwitchProfiles ensures that no value is present for SwitchProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


