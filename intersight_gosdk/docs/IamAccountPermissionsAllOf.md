# IamAccountPermissionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.AccountPermissions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.AccountPermissions"]
**AccountIdentifier** | Pointer to **string** | MOID of the account which a user can select after authentication. | [optional] [readonly] 
**AccountName** | Pointer to **string** | Name of the account which a user can select after authentication. | [optional] [readonly] 
**AccountStatus** | Pointer to **string** | Status of the account. Account remains inactive until a device is claimed to the account. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionReference**](IamPermissionReference.md) |  | [optional] 

## Methods

### NewIamAccountPermissionsAllOf

`func NewIamAccountPermissionsAllOf(classId string, objectType string, ) *IamAccountPermissionsAllOf`

NewIamAccountPermissionsAllOf instantiates a new IamAccountPermissionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountPermissionsAllOfWithDefaults

`func NewIamAccountPermissionsAllOfWithDefaults() *IamAccountPermissionsAllOf`

NewIamAccountPermissionsAllOfWithDefaults instantiates a new IamAccountPermissionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAccountPermissionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAccountPermissionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAccountPermissionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAccountPermissionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAccountPermissionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAccountPermissionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountIdentifier

`func (o *IamAccountPermissionsAllOf) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *IamAccountPermissionsAllOf) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *IamAccountPermissionsAllOf) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *IamAccountPermissionsAllOf) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### GetAccountName

`func (o *IamAccountPermissionsAllOf) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IamAccountPermissionsAllOf) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IamAccountPermissionsAllOf) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IamAccountPermissionsAllOf) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountStatus

`func (o *IamAccountPermissionsAllOf) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *IamAccountPermissionsAllOf) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *IamAccountPermissionsAllOf) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *IamAccountPermissionsAllOf) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetPermissions

`func (o *IamAccountPermissionsAllOf) GetPermissions() []IamPermissionReference`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamAccountPermissionsAllOf) GetPermissionsOk() (*[]IamPermissionReference, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamAccountPermissionsAllOf) SetPermissions(v []IamPermissionReference)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamAccountPermissionsAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamAccountPermissionsAllOf) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamAccountPermissionsAllOf) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


