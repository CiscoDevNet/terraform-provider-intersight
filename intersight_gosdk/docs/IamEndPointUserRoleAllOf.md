# IamEndPointUserRoleAllOf

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

### NewIamEndPointUserRoleAllOf

`func NewIamEndPointUserRoleAllOf(classId string, objectType string, ) *IamEndPointUserRoleAllOf`

NewIamEndPointUserRoleAllOf instantiates a new IamEndPointUserRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserRoleAllOfWithDefaults

`func NewIamEndPointUserRoleAllOfWithDefaults() *IamEndPointUserRoleAllOf`

NewIamEndPointUserRoleAllOfWithDefaults instantiates a new IamEndPointUserRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangePassword

`func (o *IamEndPointUserRoleAllOf) GetChangePassword() bool`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *IamEndPointUserRoleAllOf) GetChangePasswordOk() (*bool, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *IamEndPointUserRoleAllOf) SetChangePassword(v bool)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *IamEndPointUserRoleAllOf) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### GetEnabled

`func (o *IamEndPointUserRoleAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamEndPointUserRoleAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamEndPointUserRoleAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamEndPointUserRoleAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *IamEndPointUserRoleAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamEndPointUserRoleAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamEndPointUserRoleAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamEndPointUserRoleAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *IamEndPointUserRoleAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamEndPointUserRoleAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamEndPointUserRoleAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamEndPointUserRoleAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEndPointRole

`func (o *IamEndPointUserRoleAllOf) GetEndPointRole() []IamEndPointRoleRelationship`

GetEndPointRole returns the EndPointRole field if non-nil, zero value otherwise.

### GetEndPointRoleOk

`func (o *IamEndPointUserRoleAllOf) GetEndPointRoleOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRoleOk returns a tuple with the EndPointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRole

`func (o *IamEndPointUserRoleAllOf) SetEndPointRole(v []IamEndPointRoleRelationship)`

SetEndPointRole sets EndPointRole field to given value.

### HasEndPointRole

`func (o *IamEndPointUserRoleAllOf) HasEndPointRole() bool`

HasEndPointRole returns a boolean if a field has been set.

### SetEndPointRoleNil

`func (o *IamEndPointUserRoleAllOf) SetEndPointRoleNil(b bool)`

 SetEndPointRoleNil sets the value for EndPointRole to be an explicit nil

### UnsetEndPointRole
`func (o *IamEndPointUserRoleAllOf) UnsetEndPointRole()`

UnsetEndPointRole ensures that no value is present for EndPointRole, not even an explicit nil
### GetEndPointUser

`func (o *IamEndPointUserRoleAllOf) GetEndPointUser() IamEndPointUserRelationship`

GetEndPointUser returns the EndPointUser field if non-nil, zero value otherwise.

### GetEndPointUserOk

`func (o *IamEndPointUserRoleAllOf) GetEndPointUserOk() (*IamEndPointUserRelationship, bool)`

GetEndPointUserOk returns a tuple with the EndPointUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUser

`func (o *IamEndPointUserRoleAllOf) SetEndPointUser(v IamEndPointUserRelationship)`

SetEndPointUser sets EndPointUser field to given value.

### HasEndPointUser

`func (o *IamEndPointUserRoleAllOf) HasEndPointUser() bool`

HasEndPointUser returns a boolean if a field has been set.

### GetEndPointUserPolicy

`func (o *IamEndPointUserRoleAllOf) GetEndPointUserPolicy() IamEndPointUserPolicyRelationship`

GetEndPointUserPolicy returns the EndPointUserPolicy field if non-nil, zero value otherwise.

### GetEndPointUserPolicyOk

`func (o *IamEndPointUserRoleAllOf) GetEndPointUserPolicyOk() (*IamEndPointUserPolicyRelationship, bool)`

GetEndPointUserPolicyOk returns a tuple with the EndPointUserPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserPolicy

`func (o *IamEndPointUserRoleAllOf) SetEndPointUserPolicy(v IamEndPointUserPolicyRelationship)`

SetEndPointUserPolicy sets EndPointUserPolicy field to given value.

### HasEndPointUserPolicy

`func (o *IamEndPointUserRoleAllOf) HasEndPointUserPolicy() bool`

HasEndPointUserPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


