# IamPermissionReferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionIdentifier** | Pointer to **string** | MOID of the permission which user has access to. | [optional] [readonly] 
**PermissionName** | Pointer to **string** | Name of the permission which user has access to. | [optional] [readonly] 

## Methods

### NewIamPermissionReferenceAllOf

`func NewIamPermissionReferenceAllOf() *IamPermissionReferenceAllOf`

NewIamPermissionReferenceAllOf instantiates a new IamPermissionReferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionReferenceAllOfWithDefaults

`func NewIamPermissionReferenceAllOfWithDefaults() *IamPermissionReferenceAllOf`

NewIamPermissionReferenceAllOfWithDefaults instantiates a new IamPermissionReferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionIdentifier

`func (o *IamPermissionReferenceAllOf) GetPermissionIdentifier() string`

GetPermissionIdentifier returns the PermissionIdentifier field if non-nil, zero value otherwise.

### GetPermissionIdentifierOk

`func (o *IamPermissionReferenceAllOf) GetPermissionIdentifierOk() (*string, bool)`

GetPermissionIdentifierOk returns a tuple with the PermissionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionIdentifier

`func (o *IamPermissionReferenceAllOf) SetPermissionIdentifier(v string)`

SetPermissionIdentifier sets PermissionIdentifier field to given value.

### HasPermissionIdentifier

`func (o *IamPermissionReferenceAllOf) HasPermissionIdentifier() bool`

HasPermissionIdentifier returns a boolean if a field has been set.

### GetPermissionName

`func (o *IamPermissionReferenceAllOf) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *IamPermissionReferenceAllOf) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *IamPermissionReferenceAllOf) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.

### HasPermissionName

`func (o *IamPermissionReferenceAllOf) HasPermissionName() bool`

HasPermissionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


