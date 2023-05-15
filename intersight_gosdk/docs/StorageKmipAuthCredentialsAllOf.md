# StorageKmipAuthCredentialsAllOf

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

### NewStorageKmipAuthCredentialsAllOf

`func NewStorageKmipAuthCredentialsAllOf(classId string, objectType string, ) *StorageKmipAuthCredentialsAllOf`

NewStorageKmipAuthCredentialsAllOf instantiates a new StorageKmipAuthCredentialsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKmipAuthCredentialsAllOfWithDefaults

`func NewStorageKmipAuthCredentialsAllOfWithDefaults() *StorageKmipAuthCredentialsAllOf`

NewStorageKmipAuthCredentialsAllOfWithDefaults instantiates a new StorageKmipAuthCredentialsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKmipAuthCredentialsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKmipAuthCredentialsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKmipAuthCredentialsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKmipAuthCredentialsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKmipAuthCredentialsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKmipAuthCredentialsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *StorageKmipAuthCredentialsAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *StorageKmipAuthCredentialsAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *StorageKmipAuthCredentialsAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *StorageKmipAuthCredentialsAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *StorageKmipAuthCredentialsAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageKmipAuthCredentialsAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageKmipAuthCredentialsAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageKmipAuthCredentialsAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseAuthentication

`func (o *StorageKmipAuthCredentialsAllOf) GetUseAuthentication() bool`

GetUseAuthentication returns the UseAuthentication field if non-nil, zero value otherwise.

### GetUseAuthenticationOk

`func (o *StorageKmipAuthCredentialsAllOf) GetUseAuthenticationOk() (*bool, bool)`

GetUseAuthenticationOk returns a tuple with the UseAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthentication

`func (o *StorageKmipAuthCredentialsAllOf) SetUseAuthentication(v bool)`

SetUseAuthentication sets UseAuthentication field to given value.

### HasUseAuthentication

`func (o *StorageKmipAuthCredentialsAllOf) HasUseAuthentication() bool`

HasUseAuthentication returns a boolean if a field has been set.

### GetUsername

`func (o *StorageKmipAuthCredentialsAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageKmipAuthCredentialsAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageKmipAuthCredentialsAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageKmipAuthCredentialsAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


