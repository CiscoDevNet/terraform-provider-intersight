# StorageKmipAuthCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.KmipAuthCredentials"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.KmipAuthCredentials"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password for the KMIP server login. | [optional] 
**UseAuthentication** | Pointer to **bool** | Enables/disables the authentication for communicating with KMIP server. This flag enables the authentication which makes authentication mandatory. | [optional] 
**Username** | Pointer to **string** | The user name for the KMIP server login. | [optional] 

## Methods

### NewStorageKmipAuthCredentials

`func NewStorageKmipAuthCredentials(classId string, objectType string, ) *StorageKmipAuthCredentials`

NewStorageKmipAuthCredentials instantiates a new StorageKmipAuthCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKmipAuthCredentialsWithDefaults

`func NewStorageKmipAuthCredentialsWithDefaults() *StorageKmipAuthCredentials`

NewStorageKmipAuthCredentialsWithDefaults instantiates a new StorageKmipAuthCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKmipAuthCredentials) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKmipAuthCredentials) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKmipAuthCredentials) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKmipAuthCredentials) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKmipAuthCredentials) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKmipAuthCredentials) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *StorageKmipAuthCredentials) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *StorageKmipAuthCredentials) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *StorageKmipAuthCredentials) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *StorageKmipAuthCredentials) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *StorageKmipAuthCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageKmipAuthCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageKmipAuthCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageKmipAuthCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseAuthentication

`func (o *StorageKmipAuthCredentials) GetUseAuthentication() bool`

GetUseAuthentication returns the UseAuthentication field if non-nil, zero value otherwise.

### GetUseAuthenticationOk

`func (o *StorageKmipAuthCredentials) GetUseAuthenticationOk() (*bool, bool)`

GetUseAuthenticationOk returns a tuple with the UseAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthentication

`func (o *StorageKmipAuthCredentials) SetUseAuthentication(v bool)`

SetUseAuthentication sets UseAuthentication field to given value.

### HasUseAuthentication

`func (o *StorageKmipAuthCredentials) HasUseAuthentication() bool`

HasUseAuthentication returns a boolean if a field has been set.

### GetUsername

`func (o *StorageKmipAuthCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageKmipAuthCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageKmipAuthCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageKmipAuthCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


