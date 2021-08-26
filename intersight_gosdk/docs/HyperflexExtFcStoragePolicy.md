# HyperflexExtFcStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ExtFcStoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ExtFcStoragePolicy"]
**AdminState** | Pointer to **bool** | Enables or disables external FC storage configuration. | [optional] 
**ExtaTraffic** | Pointer to [**NullableHyperflexNamedVsan**](HyperflexNamedVsan.md) |  | [optional] 
**ExtbTraffic** | Pointer to [**NullableHyperflexNamedVsan**](HyperflexNamedVsan.md) |  | [optional] 
**WwxnPrefixRange** | Pointer to [**NullableHyperflexWwxnPrefixRange**](HyperflexWwxnPrefixRange.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexExtFcStoragePolicy

`func NewHyperflexExtFcStoragePolicy(classId string, objectType string, ) *HyperflexExtFcStoragePolicy`

NewHyperflexExtFcStoragePolicy instantiates a new HyperflexExtFcStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexExtFcStoragePolicyWithDefaults

`func NewHyperflexExtFcStoragePolicyWithDefaults() *HyperflexExtFcStoragePolicy`

NewHyperflexExtFcStoragePolicyWithDefaults instantiates a new HyperflexExtFcStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexExtFcStoragePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexExtFcStoragePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexExtFcStoragePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexExtFcStoragePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexExtFcStoragePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexExtFcStoragePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *HyperflexExtFcStoragePolicy) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *HyperflexExtFcStoragePolicy) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *HyperflexExtFcStoragePolicy) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *HyperflexExtFcStoragePolicy) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetExtaTraffic

`func (o *HyperflexExtFcStoragePolicy) GetExtaTraffic() HyperflexNamedVsan`

GetExtaTraffic returns the ExtaTraffic field if non-nil, zero value otherwise.

### GetExtaTrafficOk

`func (o *HyperflexExtFcStoragePolicy) GetExtaTrafficOk() (*HyperflexNamedVsan, bool)`

GetExtaTrafficOk returns a tuple with the ExtaTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtaTraffic

`func (o *HyperflexExtFcStoragePolicy) SetExtaTraffic(v HyperflexNamedVsan)`

SetExtaTraffic sets ExtaTraffic field to given value.

### HasExtaTraffic

`func (o *HyperflexExtFcStoragePolicy) HasExtaTraffic() bool`

HasExtaTraffic returns a boolean if a field has been set.

### SetExtaTrafficNil

`func (o *HyperflexExtFcStoragePolicy) SetExtaTrafficNil(b bool)`

 SetExtaTrafficNil sets the value for ExtaTraffic to be an explicit nil

### UnsetExtaTraffic
`func (o *HyperflexExtFcStoragePolicy) UnsetExtaTraffic()`

UnsetExtaTraffic ensures that no value is present for ExtaTraffic, not even an explicit nil
### GetExtbTraffic

`func (o *HyperflexExtFcStoragePolicy) GetExtbTraffic() HyperflexNamedVsan`

GetExtbTraffic returns the ExtbTraffic field if non-nil, zero value otherwise.

### GetExtbTrafficOk

`func (o *HyperflexExtFcStoragePolicy) GetExtbTrafficOk() (*HyperflexNamedVsan, bool)`

GetExtbTrafficOk returns a tuple with the ExtbTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtbTraffic

`func (o *HyperflexExtFcStoragePolicy) SetExtbTraffic(v HyperflexNamedVsan)`

SetExtbTraffic sets ExtbTraffic field to given value.

### HasExtbTraffic

`func (o *HyperflexExtFcStoragePolicy) HasExtbTraffic() bool`

HasExtbTraffic returns a boolean if a field has been set.

### SetExtbTrafficNil

`func (o *HyperflexExtFcStoragePolicy) SetExtbTrafficNil(b bool)`

 SetExtbTrafficNil sets the value for ExtbTraffic to be an explicit nil

### UnsetExtbTraffic
`func (o *HyperflexExtFcStoragePolicy) UnsetExtbTraffic()`

UnsetExtbTraffic ensures that no value is present for ExtbTraffic, not even an explicit nil
### GetWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicy) GetWwxnPrefixRange() HyperflexWwxnPrefixRange`

GetWwxnPrefixRange returns the WwxnPrefixRange field if non-nil, zero value otherwise.

### GetWwxnPrefixRangeOk

`func (o *HyperflexExtFcStoragePolicy) GetWwxnPrefixRangeOk() (*HyperflexWwxnPrefixRange, bool)`

GetWwxnPrefixRangeOk returns a tuple with the WwxnPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicy) SetWwxnPrefixRange(v HyperflexWwxnPrefixRange)`

SetWwxnPrefixRange sets WwxnPrefixRange field to given value.

### HasWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicy) HasWwxnPrefixRange() bool`

HasWwxnPrefixRange returns a boolean if a field has been set.

### SetWwxnPrefixRangeNil

`func (o *HyperflexExtFcStoragePolicy) SetWwxnPrefixRangeNil(b bool)`

 SetWwxnPrefixRangeNil sets the value for WwxnPrefixRange to be an explicit nil

### UnsetWwxnPrefixRange
`func (o *HyperflexExtFcStoragePolicy) UnsetWwxnPrefixRange()`

UnsetWwxnPrefixRange ensures that no value is present for WwxnPrefixRange, not even an explicit nil
### GetClusterProfiles

`func (o *HyperflexExtFcStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexExtFcStoragePolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexExtFcStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexExtFcStoragePolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexExtFcStoragePolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexExtFcStoragePolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexExtFcStoragePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexExtFcStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexExtFcStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexExtFcStoragePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


