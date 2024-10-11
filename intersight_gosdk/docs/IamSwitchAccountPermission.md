# IamSwitchAccountPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SwitchAccountPermission"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SwitchAccountPermission"]
**AccountId** | Pointer to **string** | Moid of the Account to/from which user switched the scope. | [optional] [readonly] 
**PermissionId** | Pointer to **string** | Moid of the Permission for the Account to/from which user switched the scope. | [optional] [readonly] 

## Methods

### NewIamSwitchAccountPermission

`func NewIamSwitchAccountPermission(classId string, objectType string, ) *IamSwitchAccountPermission`

NewIamSwitchAccountPermission instantiates a new IamSwitchAccountPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSwitchAccountPermissionWithDefaults

`func NewIamSwitchAccountPermissionWithDefaults() *IamSwitchAccountPermission`

NewIamSwitchAccountPermissionWithDefaults instantiates a new IamSwitchAccountPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSwitchAccountPermission) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSwitchAccountPermission) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSwitchAccountPermission) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSwitchAccountPermission) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSwitchAccountPermission) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSwitchAccountPermission) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountId

`func (o *IamSwitchAccountPermission) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IamSwitchAccountPermission) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IamSwitchAccountPermission) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IamSwitchAccountPermission) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPermissionId

`func (o *IamSwitchAccountPermission) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *IamSwitchAccountPermission) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *IamSwitchAccountPermission) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.

### HasPermissionId

`func (o *IamSwitchAccountPermission) HasPermissionId() bool`

HasPermissionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


