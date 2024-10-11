# MemoryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.Policy"]
**EnableDimmBlocklisting** | Pointer to **bool** | Enable DIMM Blocklisting on the server. | [optional] [default to false]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewMemoryPolicy

`func NewMemoryPolicy(classId string, objectType string, ) *MemoryPolicy`

NewMemoryPolicy instantiates a new MemoryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPolicyWithDefaults

`func NewMemoryPolicyWithDefaults() *MemoryPolicy`

NewMemoryPolicyWithDefaults instantiates a new MemoryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableDimmBlocklisting

`func (o *MemoryPolicy) GetEnableDimmBlocklisting() bool`

GetEnableDimmBlocklisting returns the EnableDimmBlocklisting field if non-nil, zero value otherwise.

### GetEnableDimmBlocklistingOk

`func (o *MemoryPolicy) GetEnableDimmBlocklistingOk() (*bool, bool)`

GetEnableDimmBlocklistingOk returns a tuple with the EnableDimmBlocklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDimmBlocklisting

`func (o *MemoryPolicy) SetEnableDimmBlocklisting(v bool)`

SetEnableDimmBlocklisting sets EnableDimmBlocklisting field to given value.

### HasEnableDimmBlocklisting

`func (o *MemoryPolicy) HasEnableDimmBlocklisting() bool`

HasEnableDimmBlocklisting returns a boolean if a field has been set.

### GetOrganization

`func (o *MemoryPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MemoryPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MemoryPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MemoryPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MemoryPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MemoryPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *MemoryPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MemoryPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MemoryPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MemoryPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *MemoryPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *MemoryPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


