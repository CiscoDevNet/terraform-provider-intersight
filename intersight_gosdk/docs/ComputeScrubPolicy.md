# ComputeScrubPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ScrubPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ScrubPolicy"]
**ScrubTargets** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewComputeScrubPolicy

`func NewComputeScrubPolicy(classId string, objectType string, ) *ComputeScrubPolicy`

NewComputeScrubPolicy instantiates a new ComputeScrubPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeScrubPolicyWithDefaults

`func NewComputeScrubPolicyWithDefaults() *ComputeScrubPolicy`

NewComputeScrubPolicyWithDefaults instantiates a new ComputeScrubPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeScrubPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeScrubPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeScrubPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeScrubPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeScrubPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeScrubPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetScrubTargets

`func (o *ComputeScrubPolicy) GetScrubTargets() []string`

GetScrubTargets returns the ScrubTargets field if non-nil, zero value otherwise.

### GetScrubTargetsOk

`func (o *ComputeScrubPolicy) GetScrubTargetsOk() (*[]string, bool)`

GetScrubTargetsOk returns a tuple with the ScrubTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubTargets

`func (o *ComputeScrubPolicy) SetScrubTargets(v []string)`

SetScrubTargets sets ScrubTargets field to given value.

### HasScrubTargets

`func (o *ComputeScrubPolicy) HasScrubTargets() bool`

HasScrubTargets returns a boolean if a field has been set.

### SetScrubTargetsNil

`func (o *ComputeScrubPolicy) SetScrubTargetsNil(b bool)`

 SetScrubTargetsNil sets the value for ScrubTargets to be an explicit nil

### UnsetScrubTargets
`func (o *ComputeScrubPolicy) UnsetScrubTargets()`

UnsetScrubTargets ensures that no value is present for ScrubTargets, not even an explicit nil
### GetOrganization

`func (o *ComputeScrubPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ComputeScrubPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ComputeScrubPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ComputeScrubPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ComputeScrubPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ComputeScrubPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *ComputeScrubPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ComputeScrubPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ComputeScrubPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ComputeScrubPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *ComputeScrubPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *ComputeScrubPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


