# IamGroupPermissionToRolesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.GroupPermissionToRoles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.GroupPermissionToRoles"]
**Group** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Orgs** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewIamGroupPermissionToRolesAllOf

`func NewIamGroupPermissionToRolesAllOf(classId string, objectType string, ) *IamGroupPermissionToRolesAllOf`

NewIamGroupPermissionToRolesAllOf instantiates a new IamGroupPermissionToRolesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamGroupPermissionToRolesAllOfWithDefaults

`func NewIamGroupPermissionToRolesAllOfWithDefaults() *IamGroupPermissionToRolesAllOf`

NewIamGroupPermissionToRolesAllOfWithDefaults instantiates a new IamGroupPermissionToRolesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamGroupPermissionToRolesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamGroupPermissionToRolesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamGroupPermissionToRolesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamGroupPermissionToRolesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamGroupPermissionToRolesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamGroupPermissionToRolesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGroup

`func (o *IamGroupPermissionToRolesAllOf) GetGroup() MoMoRef`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IamGroupPermissionToRolesAllOf) GetGroupOk() (*MoMoRef, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IamGroupPermissionToRolesAllOf) SetGroup(v MoMoRef)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IamGroupPermissionToRolesAllOf) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOrgs

`func (o *IamGroupPermissionToRolesAllOf) GetOrgs() []MoMoRef`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *IamGroupPermissionToRolesAllOf) GetOrgsOk() (*[]MoMoRef, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *IamGroupPermissionToRolesAllOf) SetOrgs(v []MoMoRef)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *IamGroupPermissionToRolesAllOf) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### SetOrgsNil

`func (o *IamGroupPermissionToRolesAllOf) SetOrgsNil(b bool)`

 SetOrgsNil sets the value for Orgs to be an explicit nil

### UnsetOrgs
`func (o *IamGroupPermissionToRolesAllOf) UnsetOrgs()`

UnsetOrgs ensures that no value is present for Orgs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


