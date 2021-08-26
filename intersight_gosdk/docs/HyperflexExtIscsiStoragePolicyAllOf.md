# HyperflexExtIscsiStoragePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ExtIscsiStoragePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ExtIscsiStoragePolicy"]
**AdminState** | Pointer to **bool** | Enable or disable external FCoE storage configuration. | [optional] 
**ExtaTraffic** | Pointer to [**NullableHyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**ExtbTraffic** | Pointer to [**NullableHyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexExtIscsiStoragePolicyAllOf

`func NewHyperflexExtIscsiStoragePolicyAllOf(classId string, objectType string, ) *HyperflexExtIscsiStoragePolicyAllOf`

NewHyperflexExtIscsiStoragePolicyAllOf instantiates a new HyperflexExtIscsiStoragePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexExtIscsiStoragePolicyAllOfWithDefaults

`func NewHyperflexExtIscsiStoragePolicyAllOfWithDefaults() *HyperflexExtIscsiStoragePolicyAllOf`

NewHyperflexExtIscsiStoragePolicyAllOfWithDefaults instantiates a new HyperflexExtIscsiStoragePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *HyperflexExtIscsiStoragePolicyAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetExtaTraffic() HyperflexNamedVlan`

GetExtaTraffic returns the ExtaTraffic field if non-nil, zero value otherwise.

### GetExtaTrafficOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetExtaTrafficOk() (*HyperflexNamedVlan, bool)`

GetExtaTrafficOk returns a tuple with the ExtaTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetExtaTraffic(v HyperflexNamedVlan)`

SetExtaTraffic sets ExtaTraffic field to given value.

### HasExtaTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) HasExtaTraffic() bool`

HasExtaTraffic returns a boolean if a field has been set.

### SetExtaTrafficNil

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetExtaTrafficNil(b bool)`

 SetExtaTrafficNil sets the value for ExtaTraffic to be an explicit nil

### UnsetExtaTraffic
`func (o *HyperflexExtIscsiStoragePolicyAllOf) UnsetExtaTraffic()`

UnsetExtaTraffic ensures that no value is present for ExtaTraffic, not even an explicit nil
### GetExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetExtbTraffic() HyperflexNamedVlan`

GetExtbTraffic returns the ExtbTraffic field if non-nil, zero value otherwise.

### GetExtbTrafficOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetExtbTrafficOk() (*HyperflexNamedVlan, bool)`

GetExtbTrafficOk returns a tuple with the ExtbTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetExtbTraffic(v HyperflexNamedVlan)`

SetExtbTraffic sets ExtbTraffic field to given value.

### HasExtbTraffic

`func (o *HyperflexExtIscsiStoragePolicyAllOf) HasExtbTraffic() bool`

HasExtbTraffic returns a boolean if a field has been set.

### SetExtbTrafficNil

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetExtbTrafficNil(b bool)`

 SetExtbTrafficNil sets the value for ExtbTraffic to be an explicit nil

### UnsetExtbTraffic
`func (o *HyperflexExtIscsiStoragePolicyAllOf) UnsetExtbTraffic()`

UnsetExtbTraffic ensures that no value is present for ExtbTraffic, not even an explicit nil
### GetClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexExtIscsiStoragePolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexExtIscsiStoragePolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexExtIscsiStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexExtIscsiStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexExtIscsiStoragePolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


