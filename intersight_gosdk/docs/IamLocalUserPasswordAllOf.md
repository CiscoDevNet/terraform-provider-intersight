# IamLocalUserPasswordAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | User-entered passsord to be compared to password for change password function. | [optional] 
**NewPassword** | Pointer to **string** | New password that the user&#39;s password should be changed to. | [optional] 
**Password** | Pointer to **string** | User&#39;s current valid passsord. | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamLocalUserPasswordAllOf

`func NewIamLocalUserPasswordAllOf() *IamLocalUserPasswordAllOf`

NewIamLocalUserPasswordAllOf instantiates a new IamLocalUserPasswordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLocalUserPasswordAllOfWithDefaults

`func NewIamLocalUserPasswordAllOfWithDefaults() *IamLocalUserPasswordAllOf`

NewIamLocalUserPasswordAllOfWithDefaults instantiates a new IamLocalUserPasswordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *IamLocalUserPasswordAllOf) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *IamLocalUserPasswordAllOf) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *IamLocalUserPasswordAllOf) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *IamLocalUserPasswordAllOf) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *IamLocalUserPasswordAllOf) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *IamLocalUserPasswordAllOf) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *IamLocalUserPasswordAllOf) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *IamLocalUserPasswordAllOf) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetPassword

`func (o *IamLocalUserPasswordAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamLocalUserPasswordAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamLocalUserPasswordAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamLocalUserPasswordAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *IamLocalUserPasswordAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamLocalUserPasswordAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamLocalUserPasswordAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamLocalUserPasswordAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


