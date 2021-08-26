# IamEndPointUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUser"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUser"]
**Name** | Pointer to **string** | Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters. | [optional] 
**EndPointUserRole** | Pointer to [**[]IamEndPointUserRoleRelationship**](IamEndPointUserRoleRelationship.md) | An array of relationships to iamEndPointUserRole resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewIamEndPointUserAllOf

`func NewIamEndPointUserAllOf(classId string, objectType string, ) *IamEndPointUserAllOf`

NewIamEndPointUserAllOf instantiates a new IamEndPointUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserAllOfWithDefaults

`func NewIamEndPointUserAllOfWithDefaults() *IamEndPointUserAllOf`

NewIamEndPointUserAllOfWithDefaults instantiates a new IamEndPointUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamEndPointUserAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointUserAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointUserAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointUserAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndPointUserRole

`func (o *IamEndPointUserAllOf) GetEndPointUserRole() []IamEndPointUserRoleRelationship`

GetEndPointUserRole returns the EndPointUserRole field if non-nil, zero value otherwise.

### GetEndPointUserRoleOk

`func (o *IamEndPointUserAllOf) GetEndPointUserRoleOk() (*[]IamEndPointUserRoleRelationship, bool)`

GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRole

`func (o *IamEndPointUserAllOf) SetEndPointUserRole(v []IamEndPointUserRoleRelationship)`

SetEndPointUserRole sets EndPointUserRole field to given value.

### HasEndPointUserRole

`func (o *IamEndPointUserAllOf) HasEndPointUserRole() bool`

HasEndPointUserRole returns a boolean if a field has been set.

### SetEndPointUserRoleNil

`func (o *IamEndPointUserAllOf) SetEndPointUserRoleNil(b bool)`

 SetEndPointUserRoleNil sets the value for EndPointUserRole to be an explicit nil

### UnsetEndPointUserRole
`func (o *IamEndPointUserAllOf) UnsetEndPointUserRole()`

UnsetEndPointUserRole ensures that no value is present for EndPointUserRole, not even an explicit nil
### GetOrganization

`func (o *IamEndPointUserAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamEndPointUserAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamEndPointUserAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamEndPointUserAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


