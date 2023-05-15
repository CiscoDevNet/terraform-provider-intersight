# StorageNetAppNtpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNtpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNtpServer"]
**AuthenticationEnabled** | Pointer to **bool** | Indicates that NTP symmetric authentication is enabled. | [optional] [readonly] 
**AuthenticationKeyId** | Pointer to **string** | NTP symmetric authentication key identifier or index number (ID). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**Server** | Pointer to **string** | NTP server host name, IPv4, or IPv6 address. | [optional] [readonly] 
**Version** | Pointer to **string** | NTP protocol version for server. Valid versions are 3, 4, or auto. * &#x60;none&#x60; - Default unknown NTP protocol version. * &#x60;3&#x60; - NTP protocol version is 3. * &#x60;4&#x60; - NTP protocol version is 4. * &#x60;auto&#x60; - NTP protocol version is auto. | [optional] [readonly] [default to "none"]
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNtpServer

`func NewStorageNetAppNtpServer(classId string, objectType string, ) *StorageNetAppNtpServer`

NewStorageNetAppNtpServer instantiates a new StorageNetAppNtpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNtpServerWithDefaults

`func NewStorageNetAppNtpServerWithDefaults() *StorageNetAppNtpServer`

NewStorageNetAppNtpServerWithDefaults instantiates a new StorageNetAppNtpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNtpServer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNtpServer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNtpServer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNtpServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNtpServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNtpServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationEnabled

`func (o *StorageNetAppNtpServer) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *StorageNetAppNtpServer) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *StorageNetAppNtpServer) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *StorageNetAppNtpServer) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthenticationKeyId

`func (o *StorageNetAppNtpServer) GetAuthenticationKeyId() string`

GetAuthenticationKeyId returns the AuthenticationKeyId field if non-nil, zero value otherwise.

### GetAuthenticationKeyIdOk

`func (o *StorageNetAppNtpServer) GetAuthenticationKeyIdOk() (*string, bool)`

GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKeyId

`func (o *StorageNetAppNtpServer) SetAuthenticationKeyId(v string)`

SetAuthenticationKeyId sets AuthenticationKeyId field to given value.

### HasAuthenticationKeyId

`func (o *StorageNetAppNtpServer) HasAuthenticationKeyId() bool`

HasAuthenticationKeyId returns a boolean if a field has been set.

### GetClusterUuid

`func (o *StorageNetAppNtpServer) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppNtpServer) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppNtpServer) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppNtpServer) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetServer

`func (o *StorageNetAppNtpServer) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *StorageNetAppNtpServer) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *StorageNetAppNtpServer) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *StorageNetAppNtpServer) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetVersion

`func (o *StorageNetAppNtpServer) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageNetAppNtpServer) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageNetAppNtpServer) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageNetAppNtpServer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNtpServer) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNtpServer) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNtpServer) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNtpServer) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


