# IamSwitchScopePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SwitchScopePermissions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SwitchScopePermissions"]
**AccountAccessControlId** | Pointer to **string** | Moid of the AccountAccessControl through which the access is given to switch scope. | [optional] [readonly] 
**RequestIdentifier** | Pointer to **string** | Stores the identifier of the issue for which user is trying to switch scope to another account. | [optional] [readonly] 
**SwitchedFromAccount** | Pointer to [**NullableIamSwitchAccountPermission**](IamSwitchAccountPermission.md) |  | [optional] 
**SwitchedToAccounts** | Pointer to [**[]IamSwitchAccountPermission**](IamSwitchAccountPermission.md) |  | [optional] 

## Methods

### NewIamSwitchScopePermissions

`func NewIamSwitchScopePermissions(classId string, objectType string, ) *IamSwitchScopePermissions`

NewIamSwitchScopePermissions instantiates a new IamSwitchScopePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSwitchScopePermissionsWithDefaults

`func NewIamSwitchScopePermissionsWithDefaults() *IamSwitchScopePermissions`

NewIamSwitchScopePermissionsWithDefaults instantiates a new IamSwitchScopePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSwitchScopePermissions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSwitchScopePermissions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSwitchScopePermissions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSwitchScopePermissions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSwitchScopePermissions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSwitchScopePermissions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountAccessControlId

`func (o *IamSwitchScopePermissions) GetAccountAccessControlId() string`

GetAccountAccessControlId returns the AccountAccessControlId field if non-nil, zero value otherwise.

### GetAccountAccessControlIdOk

`func (o *IamSwitchScopePermissions) GetAccountAccessControlIdOk() (*string, bool)`

GetAccountAccessControlIdOk returns a tuple with the AccountAccessControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAccessControlId

`func (o *IamSwitchScopePermissions) SetAccountAccessControlId(v string)`

SetAccountAccessControlId sets AccountAccessControlId field to given value.

### HasAccountAccessControlId

`func (o *IamSwitchScopePermissions) HasAccountAccessControlId() bool`

HasAccountAccessControlId returns a boolean if a field has been set.

### GetRequestIdentifier

`func (o *IamSwitchScopePermissions) GetRequestIdentifier() string`

GetRequestIdentifier returns the RequestIdentifier field if non-nil, zero value otherwise.

### GetRequestIdentifierOk

`func (o *IamSwitchScopePermissions) GetRequestIdentifierOk() (*string, bool)`

GetRequestIdentifierOk returns a tuple with the RequestIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIdentifier

`func (o *IamSwitchScopePermissions) SetRequestIdentifier(v string)`

SetRequestIdentifier sets RequestIdentifier field to given value.

### HasRequestIdentifier

`func (o *IamSwitchScopePermissions) HasRequestIdentifier() bool`

HasRequestIdentifier returns a boolean if a field has been set.

### GetSwitchedFromAccount

`func (o *IamSwitchScopePermissions) GetSwitchedFromAccount() IamSwitchAccountPermission`

GetSwitchedFromAccount returns the SwitchedFromAccount field if non-nil, zero value otherwise.

### GetSwitchedFromAccountOk

`func (o *IamSwitchScopePermissions) GetSwitchedFromAccountOk() (*IamSwitchAccountPermission, bool)`

GetSwitchedFromAccountOk returns a tuple with the SwitchedFromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchedFromAccount

`func (o *IamSwitchScopePermissions) SetSwitchedFromAccount(v IamSwitchAccountPermission)`

SetSwitchedFromAccount sets SwitchedFromAccount field to given value.

### HasSwitchedFromAccount

`func (o *IamSwitchScopePermissions) HasSwitchedFromAccount() bool`

HasSwitchedFromAccount returns a boolean if a field has been set.

### SetSwitchedFromAccountNil

`func (o *IamSwitchScopePermissions) SetSwitchedFromAccountNil(b bool)`

 SetSwitchedFromAccountNil sets the value for SwitchedFromAccount to be an explicit nil

### UnsetSwitchedFromAccount
`func (o *IamSwitchScopePermissions) UnsetSwitchedFromAccount()`

UnsetSwitchedFromAccount ensures that no value is present for SwitchedFromAccount, not even an explicit nil
### GetSwitchedToAccounts

`func (o *IamSwitchScopePermissions) GetSwitchedToAccounts() []IamSwitchAccountPermission`

GetSwitchedToAccounts returns the SwitchedToAccounts field if non-nil, zero value otherwise.

### GetSwitchedToAccountsOk

`func (o *IamSwitchScopePermissions) GetSwitchedToAccountsOk() (*[]IamSwitchAccountPermission, bool)`

GetSwitchedToAccountsOk returns a tuple with the SwitchedToAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchedToAccounts

`func (o *IamSwitchScopePermissions) SetSwitchedToAccounts(v []IamSwitchAccountPermission)`

SetSwitchedToAccounts sets SwitchedToAccounts field to given value.

### HasSwitchedToAccounts

`func (o *IamSwitchScopePermissions) HasSwitchedToAccounts() bool`

HasSwitchedToAccounts returns a boolean if a field has been set.

### SetSwitchedToAccountsNil

`func (o *IamSwitchScopePermissions) SetSwitchedToAccountsNil(b bool)`

 SetSwitchedToAccountsNil sets the value for SwitchedToAccounts to be an explicit nil

### UnsetSwitchedToAccounts
`func (o *IamSwitchScopePermissions) UnsetSwitchedToAccounts()`

UnsetSwitchedToAccounts ensures that no value is present for SwitchedToAccounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


