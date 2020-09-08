# IamAccountPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | Pointer to **string** | MOID of the account which a user can select after authentication. | [optional] [readonly] 
**AccountName** | Pointer to **string** | Name of the account which a user can select after authentication. | [optional] [readonly] 
**AccountStatus** | Pointer to **string** | Status of the account. Account remains inactive until a device is claimed to the account. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionReference**](iam.PermissionReference.md) |  | [optional] 

## Methods

### NewIamAccountPermissions

`func NewIamAccountPermissions() *IamAccountPermissions`

NewIamAccountPermissions instantiates a new IamAccountPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountPermissionsWithDefaults

`func NewIamAccountPermissionsWithDefaults() *IamAccountPermissions`

NewIamAccountPermissionsWithDefaults instantiates a new IamAccountPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifier

`func (o *IamAccountPermissions) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *IamAccountPermissions) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *IamAccountPermissions) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *IamAccountPermissions) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### GetAccountName

`func (o *IamAccountPermissions) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IamAccountPermissions) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IamAccountPermissions) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IamAccountPermissions) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountStatus

`func (o *IamAccountPermissions) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *IamAccountPermissions) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *IamAccountPermissions) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *IamAccountPermissions) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetPermissions

`func (o *IamAccountPermissions) GetPermissions() []IamPermissionReference`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamAccountPermissions) GetPermissionsOk() (*[]IamPermissionReference, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamAccountPermissions) SetPermissions(v []IamPermissionReference)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamAccountPermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


