# FabricSwitchClusterProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchClusterProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchClusterProfile"]
**ConfigContext** | Pointer to [**NullablePolicyConfigContext**](PolicyConfigContext.md) |  | [optional] 
**SwitchProfilesCount** | Pointer to **int64** | Number of switch profiles that are part of this cluster profile. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SwitchProfiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfile

`func NewFabricSwitchClusterProfile(classId string, objectType string, ) *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfile instantiates a new FabricSwitchClusterProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileWithDefaults

`func NewFabricSwitchClusterProfileWithDefaults() *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfileWithDefaults instantiates a new FabricSwitchClusterProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchClusterProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchClusterProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigContext

`func (o *FabricSwitchClusterProfile) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *FabricSwitchClusterProfile) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *FabricSwitchClusterProfile) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *FabricSwitchClusterProfile) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### SetConfigContextNil

`func (o *FabricSwitchClusterProfile) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *FabricSwitchClusterProfile) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetSwitchProfilesCount

`func (o *FabricSwitchClusterProfile) GetSwitchProfilesCount() int64`

GetSwitchProfilesCount returns the SwitchProfilesCount field if non-nil, zero value otherwise.

### GetSwitchProfilesCountOk

`func (o *FabricSwitchClusterProfile) GetSwitchProfilesCountOk() (*int64, bool)`

GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilesCount

`func (o *FabricSwitchClusterProfile) SetSwitchProfilesCount(v int64)`

SetSwitchProfilesCount sets SwitchProfilesCount field to given value.

### HasSwitchProfilesCount

`func (o *FabricSwitchClusterProfile) HasSwitchProfilesCount() bool`

HasSwitchProfilesCount returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *FabricSwitchClusterProfile) GetSwitchProfiles() []FabricSwitchProfileRelationship`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *FabricSwitchClusterProfile) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *FabricSwitchClusterProfile) SetSwitchProfiles(v []FabricSwitchProfileRelationship)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *FabricSwitchClusterProfile) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### SetSwitchProfilesNil

`func (o *FabricSwitchClusterProfile) SetSwitchProfilesNil(b bool)`

 SetSwitchProfilesNil sets the value for SwitchProfiles to be an explicit nil

### UnsetSwitchProfiles
`func (o *FabricSwitchClusterProfile) UnsetSwitchProfiles()`

UnsetSwitchProfiles ensures that no value is present for SwitchProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


