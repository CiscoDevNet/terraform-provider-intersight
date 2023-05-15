# StorageRemoteKeySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.RemoteKeySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.RemoteKeySetting"]
**AuthCredentials** | Pointer to [**NullableStorageKmipAuthCredentials**](StorageKmipAuthCredentials.md) |  | [optional] 
**IsExistingKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;existingKey&#39; property has been set. | [optional] [readonly] [default to false]
**PrimaryServer** | Pointer to [**NullableStorageKmipServer**](StorageKmipServer.md) |  | [optional] 
**SecondaryServer** | Pointer to [**NullableStorageKmipServer**](StorageKmipServer.md) |  | [optional] 
**ServerCertificate** | Pointer to **string** | The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server. | [optional] 

## Methods

### NewStorageRemoteKeySetting

`func NewStorageRemoteKeySetting(classId string, objectType string, ) *StorageRemoteKeySetting`

NewStorageRemoteKeySetting instantiates a new StorageRemoteKeySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageRemoteKeySettingWithDefaults

`func NewStorageRemoteKeySettingWithDefaults() *StorageRemoteKeySetting`

NewStorageRemoteKeySettingWithDefaults instantiates a new StorageRemoteKeySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageRemoteKeySetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageRemoteKeySetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageRemoteKeySetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageRemoteKeySetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageRemoteKeySetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageRemoteKeySetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthCredentials

`func (o *StorageRemoteKeySetting) GetAuthCredentials() StorageKmipAuthCredentials`

GetAuthCredentials returns the AuthCredentials field if non-nil, zero value otherwise.

### GetAuthCredentialsOk

`func (o *StorageRemoteKeySetting) GetAuthCredentialsOk() (*StorageKmipAuthCredentials, bool)`

GetAuthCredentialsOk returns a tuple with the AuthCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCredentials

`func (o *StorageRemoteKeySetting) SetAuthCredentials(v StorageKmipAuthCredentials)`

SetAuthCredentials sets AuthCredentials field to given value.

### HasAuthCredentials

`func (o *StorageRemoteKeySetting) HasAuthCredentials() bool`

HasAuthCredentials returns a boolean if a field has been set.

### SetAuthCredentialsNil

`func (o *StorageRemoteKeySetting) SetAuthCredentialsNil(b bool)`

 SetAuthCredentialsNil sets the value for AuthCredentials to be an explicit nil

### UnsetAuthCredentials
`func (o *StorageRemoteKeySetting) UnsetAuthCredentials()`

UnsetAuthCredentials ensures that no value is present for AuthCredentials, not even an explicit nil
### GetIsExistingKeySet

`func (o *StorageRemoteKeySetting) GetIsExistingKeySet() bool`

GetIsExistingKeySet returns the IsExistingKeySet field if non-nil, zero value otherwise.

### GetIsExistingKeySetOk

`func (o *StorageRemoteKeySetting) GetIsExistingKeySetOk() (*bool, bool)`

GetIsExistingKeySetOk returns a tuple with the IsExistingKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExistingKeySet

`func (o *StorageRemoteKeySetting) SetIsExistingKeySet(v bool)`

SetIsExistingKeySet sets IsExistingKeySet field to given value.

### HasIsExistingKeySet

`func (o *StorageRemoteKeySetting) HasIsExistingKeySet() bool`

HasIsExistingKeySet returns a boolean if a field has been set.

### GetPrimaryServer

`func (o *StorageRemoteKeySetting) GetPrimaryServer() StorageKmipServer`

GetPrimaryServer returns the PrimaryServer field if non-nil, zero value otherwise.

### GetPrimaryServerOk

`func (o *StorageRemoteKeySetting) GetPrimaryServerOk() (*StorageKmipServer, bool)`

GetPrimaryServerOk returns a tuple with the PrimaryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryServer

`func (o *StorageRemoteKeySetting) SetPrimaryServer(v StorageKmipServer)`

SetPrimaryServer sets PrimaryServer field to given value.

### HasPrimaryServer

`func (o *StorageRemoteKeySetting) HasPrimaryServer() bool`

HasPrimaryServer returns a boolean if a field has been set.

### SetPrimaryServerNil

`func (o *StorageRemoteKeySetting) SetPrimaryServerNil(b bool)`

 SetPrimaryServerNil sets the value for PrimaryServer to be an explicit nil

### UnsetPrimaryServer
`func (o *StorageRemoteKeySetting) UnsetPrimaryServer()`

UnsetPrimaryServer ensures that no value is present for PrimaryServer, not even an explicit nil
### GetSecondaryServer

`func (o *StorageRemoteKeySetting) GetSecondaryServer() StorageKmipServer`

GetSecondaryServer returns the SecondaryServer field if non-nil, zero value otherwise.

### GetSecondaryServerOk

`func (o *StorageRemoteKeySetting) GetSecondaryServerOk() (*StorageKmipServer, bool)`

GetSecondaryServerOk returns a tuple with the SecondaryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryServer

`func (o *StorageRemoteKeySetting) SetSecondaryServer(v StorageKmipServer)`

SetSecondaryServer sets SecondaryServer field to given value.

### HasSecondaryServer

`func (o *StorageRemoteKeySetting) HasSecondaryServer() bool`

HasSecondaryServer returns a boolean if a field has been set.

### SetSecondaryServerNil

`func (o *StorageRemoteKeySetting) SetSecondaryServerNil(b bool)`

 SetSecondaryServerNil sets the value for SecondaryServer to be an explicit nil

### UnsetSecondaryServer
`func (o *StorageRemoteKeySetting) UnsetSecondaryServer()`

UnsetSecondaryServer ensures that no value is present for SecondaryServer, not even an explicit nil
### GetServerCertificate

`func (o *StorageRemoteKeySetting) GetServerCertificate() string`

GetServerCertificate returns the ServerCertificate field if non-nil, zero value otherwise.

### GetServerCertificateOk

`func (o *StorageRemoteKeySetting) GetServerCertificateOk() (*string, bool)`

GetServerCertificateOk returns a tuple with the ServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificate

`func (o *StorageRemoteKeySetting) SetServerCertificate(v string)`

SetServerCertificate sets ServerCertificate field to given value.

### HasServerCertificate

`func (o *StorageRemoteKeySetting) HasServerCertificate() bool`

HasServerCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


