# StorageNetAppNtpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNtpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNtpServer"]
**AuthenticationEnabled** | Pointer to **bool** | Indicates whether or not NTP symmetric authentication is enabled. | [optional] [readonly] 
**AuthenticationKeyId** | Pointer to **string** | NTP symmetric authentication key identifier or index number (ID). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**IsAuthenticationEnabled** | Pointer to **string** | Indicates whether or not NTP symmetric authentication is enabled. | [optional] [readonly] 
**Server** | Pointer to **string** | NTP server host name, IPv4, or IPv6 address. | [optional] [readonly] 
**Version** | Pointer to **string** | NTP protocol version for server. Valid versions are 3, 4, or auto. * &#x60;none&#x60; - Default unknown NTP protocol version. * &#x60;3&#x60; - NTP protocol version is 3. * &#x60;4&#x60; - NTP protocol version is 4. * &#x60;auto&#x60; - NTP protocol version is auto. | [optional] [readonly] [default to "none"]
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNtpServerAllOf

`func NewStorageNetAppNtpServerAllOf(classId string, objectType string, ) *StorageNetAppNtpServerAllOf`

NewStorageNetAppNtpServerAllOf instantiates a new StorageNetAppNtpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNtpServerAllOfWithDefaults

`func NewStorageNetAppNtpServerAllOfWithDefaults() *StorageNetAppNtpServerAllOf`

NewStorageNetAppNtpServerAllOfWithDefaults instantiates a new StorageNetAppNtpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNtpServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNtpServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNtpServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNtpServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNtpServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNtpServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *StorageNetAppNtpServerAllOf) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthenticationKeyId

`func (o *StorageNetAppNtpServerAllOf) GetAuthenticationKeyId() string`

GetAuthenticationKeyId returns the AuthenticationKeyId field if non-nil, zero value otherwise.

### GetAuthenticationKeyIdOk

`func (o *StorageNetAppNtpServerAllOf) GetAuthenticationKeyIdOk() (*string, bool)`

GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKeyId

`func (o *StorageNetAppNtpServerAllOf) SetAuthenticationKeyId(v string)`

SetAuthenticationKeyId sets AuthenticationKeyId field to given value.

### HasAuthenticationKeyId

`func (o *StorageNetAppNtpServerAllOf) HasAuthenticationKeyId() bool`

HasAuthenticationKeyId returns a boolean if a field has been set.

### GetClusterUuid

`func (o *StorageNetAppNtpServerAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppNtpServerAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppNtpServerAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppNtpServerAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetIsAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) GetIsAuthenticationEnabled() string`

GetIsAuthenticationEnabled returns the IsAuthenticationEnabled field if non-nil, zero value otherwise.

### GetIsAuthenticationEnabledOk

`func (o *StorageNetAppNtpServerAllOf) GetIsAuthenticationEnabledOk() (*string, bool)`

GetIsAuthenticationEnabledOk returns a tuple with the IsAuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) SetIsAuthenticationEnabled(v string)`

SetIsAuthenticationEnabled sets IsAuthenticationEnabled field to given value.

### HasIsAuthenticationEnabled

`func (o *StorageNetAppNtpServerAllOf) HasIsAuthenticationEnabled() bool`

HasIsAuthenticationEnabled returns a boolean if a field has been set.

### GetServer

`func (o *StorageNetAppNtpServerAllOf) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *StorageNetAppNtpServerAllOf) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *StorageNetAppNtpServerAllOf) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *StorageNetAppNtpServerAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetVersion

`func (o *StorageNetAppNtpServerAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageNetAppNtpServerAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageNetAppNtpServerAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageNetAppNtpServerAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNtpServerAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNtpServerAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNtpServerAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNtpServerAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


