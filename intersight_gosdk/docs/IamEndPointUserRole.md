# IamEndPointUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUserRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUserRole"]
**ChangePassword** | Pointer to **bool** | Denotes whether password has changed. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Enables the user account on the endpoint. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Valid login password of the user. | [optional] 
**EndPointRole** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] 
**EndPointUser** | Pointer to [**IamEndPointUserRelationship**](IamEndPointUserRelationship.md) |  | [optional] 
**EndPointUserPolicy** | Pointer to [**IamEndPointUserPolicyRelationship**](IamEndPointUserPolicyRelationship.md) |  | [optional] 

## Methods

### NewIamEndPointUserRole

`func NewIamEndPointUserRole(classId string, objectType string, ) *IamEndPointUserRole`

NewIamEndPointUserRole instantiates a new IamEndPointUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserRoleWithDefaults

`func NewIamEndPointUserRoleWithDefaults() *IamEndPointUserRole`

NewIamEndPointUserRoleWithDefaults instantiates a new IamEndPointUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangePassword

`func (o *IamEndPointUserRole) GetChangePassword() bool`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *IamEndPointUserRole) GetChangePasswordOk() (*bool, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *IamEndPointUserRole) SetChangePassword(v bool)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *IamEndPointUserRole) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### GetEnabled

`func (o *IamEndPointUserRole) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamEndPointUserRole) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamEndPointUserRole) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamEndPointUserRole) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *IamEndPointUserRole) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamEndPointUserRole) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamEndPointUserRole) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamEndPointUserRole) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *IamEndPointUserRole) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamEndPointUserRole) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamEndPointUserRole) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamEndPointUserRole) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEndPointRole

`func (o *IamEndPointUserRole) GetEndPointRole() []IamEndPointRoleRelationship`

GetEndPointRole returns the EndPointRole field if non-nil, zero value otherwise.

### GetEndPointRoleOk

`func (o *IamEndPointUserRole) GetEndPointRoleOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRoleOk returns a tuple with the EndPointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRole

`func (o *IamEndPointUserRole) SetEndPointRole(v []IamEndPointRoleRelationship)`

SetEndPointRole sets EndPointRole field to given value.

### HasEndPointRole

`func (o *IamEndPointUserRole) HasEndPointRole() bool`

HasEndPointRole returns a boolean if a field has been set.

### SetEndPointRoleNil

`func (o *IamEndPointUserRole) SetEndPointRoleNil(b bool)`

 SetEndPointRoleNil sets the value for EndPointRole to be an explicit nil

### UnsetEndPointRole
`func (o *IamEndPointUserRole) UnsetEndPointRole()`

UnsetEndPointRole ensures that no value is present for EndPointRole, not even an explicit nil
### GetEndPointUser

`func (o *IamEndPointUserRole) GetEndPointUser() IamEndPointUserRelationship`

GetEndPointUser returns the EndPointUser field if non-nil, zero value otherwise.

### GetEndPointUserOk

`func (o *IamEndPointUserRole) GetEndPointUserOk() (*IamEndPointUserRelationship, bool)`

GetEndPointUserOk returns a tuple with the EndPointUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUser

`func (o *IamEndPointUserRole) SetEndPointUser(v IamEndPointUserRelationship)`

SetEndPointUser sets EndPointUser field to given value.

### HasEndPointUser

`func (o *IamEndPointUserRole) HasEndPointUser() bool`

HasEndPointUser returns a boolean if a field has been set.

### GetEndPointUserPolicy

`func (o *IamEndPointUserRole) GetEndPointUserPolicy() IamEndPointUserPolicyRelationship`

GetEndPointUserPolicy returns the EndPointUserPolicy field if non-nil, zero value otherwise.

### GetEndPointUserPolicyOk

`func (o *IamEndPointUserRole) GetEndPointUserPolicyOk() (*IamEndPointUserPolicyRelationship, bool)`

GetEndPointUserPolicyOk returns a tuple with the EndPointUserPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserPolicy

`func (o *IamEndPointUserRole) SetEndPointUserPolicy(v IamEndPointUserPolicyRelationship)`

SetEndPointUserPolicy sets EndPointUserPolicy field to given value.

### HasEndPointUserPolicy

`func (o *IamEndPointUserRole) HasEndPointUserPolicy() bool`

HasEndPointUserPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


