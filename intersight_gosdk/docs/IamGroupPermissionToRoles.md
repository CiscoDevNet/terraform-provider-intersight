# IamGroupPermissionToRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.GroupPermissionToRoles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.GroupPermissionToRoles"]
**Group** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Orgs** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewIamGroupPermissionToRoles

`func NewIamGroupPermissionToRoles(classId string, objectType string, ) *IamGroupPermissionToRoles`

NewIamGroupPermissionToRoles instantiates a new IamGroupPermissionToRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamGroupPermissionToRolesWithDefaults

`func NewIamGroupPermissionToRolesWithDefaults() *IamGroupPermissionToRoles`

NewIamGroupPermissionToRolesWithDefaults instantiates a new IamGroupPermissionToRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamGroupPermissionToRoles) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamGroupPermissionToRoles) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamGroupPermissionToRoles) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamGroupPermissionToRoles) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamGroupPermissionToRoles) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamGroupPermissionToRoles) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGroup

`func (o *IamGroupPermissionToRoles) GetGroup() MoMoRef`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IamGroupPermissionToRoles) GetGroupOk() (*MoMoRef, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IamGroupPermissionToRoles) SetGroup(v MoMoRef)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IamGroupPermissionToRoles) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOrgs

`func (o *IamGroupPermissionToRoles) GetOrgs() []MoMoRef`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *IamGroupPermissionToRoles) GetOrgsOk() (*[]MoMoRef, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *IamGroupPermissionToRoles) SetOrgs(v []MoMoRef)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *IamGroupPermissionToRoles) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### SetOrgsNil

`func (o *IamGroupPermissionToRoles) SetOrgsNil(b bool)`

 SetOrgsNil sets the value for Orgs to be an explicit nil

### UnsetOrgs
`func (o *IamGroupPermissionToRoles) UnsetOrgs()`

UnsetOrgs ensures that no value is present for Orgs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


