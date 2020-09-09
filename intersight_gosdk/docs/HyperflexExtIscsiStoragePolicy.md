# HyperflexExtIscsiStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **bool** | Enable or disable external FCoE storage configuration. | [optional] 
**ExtaTraffic** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**ExtbTraffic** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexExtIscsiStoragePolicy

`func NewHyperflexExtIscsiStoragePolicy() *HyperflexExtIscsiStoragePolicy`

NewHyperflexExtIscsiStoragePolicy instantiates a new HyperflexExtIscsiStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexExtIscsiStoragePolicyWithDefaults

`func NewHyperflexExtIscsiStoragePolicyWithDefaults() *HyperflexExtIscsiStoragePolicy`

NewHyperflexExtIscsiStoragePolicyWithDefaults instantiates a new HyperflexExtIscsiStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *HyperflexExtIscsiStoragePolicy) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *HyperflexExtIscsiStoragePolicy) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *HyperflexExtIscsiStoragePolicy) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *HyperflexExtIscsiStoragePolicy) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicy) GetExtaTraffic() HyperflexNamedVlan`

GetExtaTraffic returns the ExtaTraffic field if non-nil, zero value otherwise.

### GetExtaTrafficOk

`func (o *HyperflexExtIscsiStoragePolicy) GetExtaTrafficOk() (*HyperflexNamedVlan, bool)`

GetExtaTrafficOk returns a tuple with the ExtaTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicy) SetExtaTraffic(v HyperflexNamedVlan)`

SetExtaTraffic sets ExtaTraffic field to given value.

### HasExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicy) HasExtaTraffic() bool`

HasExtaTraffic returns a boolean if a field has been set.

### GetExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicy) GetExtbTraffic() HyperflexNamedVlan`

GetExtbTraffic returns the ExtbTraffic field if non-nil, zero value otherwise.

### GetExtbTrafficOk

`func (o *HyperflexExtIscsiStoragePolicy) GetExtbTrafficOk() (*HyperflexNamedVlan, bool)`

GetExtbTrafficOk returns a tuple with the ExtbTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicy) SetExtbTraffic(v HyperflexNamedVlan)`

SetExtbTraffic sets ExtbTraffic field to given value.

### HasExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicy) HasExtbTraffic() bool`

HasExtbTraffic returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexExtIscsiStoragePolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexExtIscsiStoragePolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexExtIscsiStoragePolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexExtIscsiStoragePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexExtIscsiStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexExtIscsiStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexExtIscsiStoragePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


