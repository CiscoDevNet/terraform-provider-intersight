# IamEncryptedOidcAuthTokenSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EncryptedOidcAuthTokenSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EncryptedOidcAuthTokenSet"]
**IsAccessTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;accessToken&#39; property has been set. | [optional] [readonly] [default to false]
**IsIdTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;idToken&#39; property has been set. | [optional] [readonly] [default to false]
**IsRefreshTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;refreshToken&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewIamEncryptedOidcAuthTokenSet

`func NewIamEncryptedOidcAuthTokenSet(classId string, objectType string, ) *IamEncryptedOidcAuthTokenSet`

NewIamEncryptedOidcAuthTokenSet instantiates a new IamEncryptedOidcAuthTokenSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEncryptedOidcAuthTokenSetWithDefaults

`func NewIamEncryptedOidcAuthTokenSetWithDefaults() *IamEncryptedOidcAuthTokenSet`

NewIamEncryptedOidcAuthTokenSetWithDefaults instantiates a new IamEncryptedOidcAuthTokenSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEncryptedOidcAuthTokenSet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEncryptedOidcAuthTokenSet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEncryptedOidcAuthTokenSet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEncryptedOidcAuthTokenSet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEncryptedOidcAuthTokenSet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEncryptedOidcAuthTokenSet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAccessTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) GetIsAccessTokenSet() bool`

GetIsAccessTokenSet returns the IsAccessTokenSet field if non-nil, zero value otherwise.

### GetIsAccessTokenSetOk

`func (o *IamEncryptedOidcAuthTokenSet) GetIsAccessTokenSetOk() (*bool, bool)`

GetIsAccessTokenSetOk returns a tuple with the IsAccessTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) SetIsAccessTokenSet(v bool)`

SetIsAccessTokenSet sets IsAccessTokenSet field to given value.

### HasIsAccessTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) HasIsAccessTokenSet() bool`

HasIsAccessTokenSet returns a boolean if a field has been set.

### GetIsIdTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) GetIsIdTokenSet() bool`

GetIsIdTokenSet returns the IsIdTokenSet field if non-nil, zero value otherwise.

### GetIsIdTokenSetOk

`func (o *IamEncryptedOidcAuthTokenSet) GetIsIdTokenSetOk() (*bool, bool)`

GetIsIdTokenSetOk returns a tuple with the IsIdTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) SetIsIdTokenSet(v bool)`

SetIsIdTokenSet sets IsIdTokenSet field to given value.

### HasIsIdTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) HasIsIdTokenSet() bool`

HasIsIdTokenSet returns a boolean if a field has been set.

### GetIsRefreshTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) GetIsRefreshTokenSet() bool`

GetIsRefreshTokenSet returns the IsRefreshTokenSet field if non-nil, zero value otherwise.

### GetIsRefreshTokenSetOk

`func (o *IamEncryptedOidcAuthTokenSet) GetIsRefreshTokenSetOk() (*bool, bool)`

GetIsRefreshTokenSetOk returns a tuple with the IsRefreshTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefreshTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) SetIsRefreshTokenSet(v bool)`

SetIsRefreshTokenSet sets IsRefreshTokenSet field to given value.

### HasIsRefreshTokenSet

`func (o *IamEncryptedOidcAuthTokenSet) HasIsRefreshTokenSet() bool`

HasIsRefreshTokenSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


