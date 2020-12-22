# StorageRemoteKeySettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.RemoteKeySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.RemoteKeySetting"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password for the KMIP server login. | [optional] 
**Port** | Pointer to **int64** | The port to which the KMIP client should connect. | [optional] [default to 5696]
**PrimaryServer** | Pointer to **string** | The IP address of the primary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address. | [optional] 
**SecondaryServer** | Pointer to **string** | The IP address of the secondary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address. | [optional] 
**ServerCertificate** | Pointer to **string** | The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server. | [optional] 
**Username** | Pointer to **string** | The user name for the KMIP server login. | [optional] 

## Methods

### NewStorageRemoteKeySettingAllOf

`func NewStorageRemoteKeySettingAllOf(classId string, objectType string, ) *StorageRemoteKeySettingAllOf`

NewStorageRemoteKeySettingAllOf instantiates a new StorageRemoteKeySettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageRemoteKeySettingAllOfWithDefaults

`func NewStorageRemoteKeySettingAllOfWithDefaults() *StorageRemoteKeySettingAllOf`

NewStorageRemoteKeySettingAllOfWithDefaults instantiates a new StorageRemoteKeySettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageRemoteKeySettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageRemoteKeySettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageRemoteKeySettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageRemoteKeySettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageRemoteKeySettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageRemoteKeySettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *StorageRemoteKeySettingAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *StorageRemoteKeySettingAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *StorageRemoteKeySettingAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *StorageRemoteKeySettingAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *StorageRemoteKeySettingAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageRemoteKeySettingAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageRemoteKeySettingAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageRemoteKeySettingAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *StorageRemoteKeySettingAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageRemoteKeySettingAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageRemoteKeySettingAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageRemoteKeySettingAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrimaryServer

`func (o *StorageRemoteKeySettingAllOf) GetPrimaryServer() string`

GetPrimaryServer returns the PrimaryServer field if non-nil, zero value otherwise.

### GetPrimaryServerOk

`func (o *StorageRemoteKeySettingAllOf) GetPrimaryServerOk() (*string, bool)`

GetPrimaryServerOk returns a tuple with the PrimaryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryServer

`func (o *StorageRemoteKeySettingAllOf) SetPrimaryServer(v string)`

SetPrimaryServer sets PrimaryServer field to given value.

### HasPrimaryServer

`func (o *StorageRemoteKeySettingAllOf) HasPrimaryServer() bool`

HasPrimaryServer returns a boolean if a field has been set.

### GetSecondaryServer

`func (o *StorageRemoteKeySettingAllOf) GetSecondaryServer() string`

GetSecondaryServer returns the SecondaryServer field if non-nil, zero value otherwise.

### GetSecondaryServerOk

`func (o *StorageRemoteKeySettingAllOf) GetSecondaryServerOk() (*string, bool)`

GetSecondaryServerOk returns a tuple with the SecondaryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryServer

`func (o *StorageRemoteKeySettingAllOf) SetSecondaryServer(v string)`

SetSecondaryServer sets SecondaryServer field to given value.

### HasSecondaryServer

`func (o *StorageRemoteKeySettingAllOf) HasSecondaryServer() bool`

HasSecondaryServer returns a boolean if a field has been set.

### GetServerCertificate

`func (o *StorageRemoteKeySettingAllOf) GetServerCertificate() string`

GetServerCertificate returns the ServerCertificate field if non-nil, zero value otherwise.

### GetServerCertificateOk

`func (o *StorageRemoteKeySettingAllOf) GetServerCertificateOk() (*string, bool)`

GetServerCertificateOk returns a tuple with the ServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificate

`func (o *StorageRemoteKeySettingAllOf) SetServerCertificate(v string)`

SetServerCertificate sets ServerCertificate field to given value.

### HasServerCertificate

`func (o *StorageRemoteKeySettingAllOf) HasServerCertificate() bool`

HasServerCertificate returns a boolean if a field has been set.

### GetUsername

`func (o *StorageRemoteKeySettingAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageRemoteKeySettingAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageRemoteKeySettingAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageRemoteKeySettingAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


