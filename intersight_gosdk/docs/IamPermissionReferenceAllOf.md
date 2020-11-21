# IamPermissionReferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PermissionReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PermissionReference"]
**PermissionIdentifier** | Pointer to **string** | MOID of the permission which user has access to. | [optional] [readonly] 
**PermissionName** | Pointer to **string** | Name of the permission which user has access to. | [optional] [readonly] 

## Methods

### NewIamPermissionReferenceAllOf

`func NewIamPermissionReferenceAllOf(classId string, objectType string, ) *IamPermissionReferenceAllOf`

NewIamPermissionReferenceAllOf instantiates a new IamPermissionReferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionReferenceAllOfWithDefaults

`func NewIamPermissionReferenceAllOfWithDefaults() *IamPermissionReferenceAllOf`

NewIamPermissionReferenceAllOfWithDefaults instantiates a new IamPermissionReferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPermissionReferenceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPermissionReferenceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPermissionReferenceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPermissionReferenceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPermissionReferenceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPermissionReferenceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


