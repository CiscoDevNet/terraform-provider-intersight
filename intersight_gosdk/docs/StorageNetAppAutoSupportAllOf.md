# StorageNetAppAutoSupportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppAutoSupport"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppAutoSupport"]
**Enabled** | Pointer to **bool** | Specifies whether the AutoSupport daemon is enabled. When this setting is disabled, delivery of all AutoSupport messages is turned off. | [optional] [readonly] 
**From** | Pointer to **string** | The e-mail address from which the AutoSupport messages are sent. | [optional] [readonly] 
**ProxyUrl** | Pointer to **string** | Proxy server for AutoSupport message delivery via HTTP/S. | [optional] [readonly] 
**Transport** | Pointer to **string** | The name of the transport protocol used to deliver AutoSupport messages. * &#x60;none&#x60; - Default value of none when nothing is returned. * &#x60;smtp&#x60; - Simple Mail Transfer Protocol. * &#x60;http&#x60; - Hypertext Transfer Protocol. * &#x60;https&#x60; - Hypertext Transfer Protocol Secure. | [optional] [readonly] [default to "none"]

## Methods

### NewStorageNetAppAutoSupportAllOf

`func NewStorageNetAppAutoSupportAllOf(classId string, objectType string, ) *StorageNetAppAutoSupportAllOf`

NewStorageNetAppAutoSupportAllOf instantiates a new StorageNetAppAutoSupportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppAutoSupportAllOfWithDefaults

`func NewStorageNetAppAutoSupportAllOfWithDefaults() *StorageNetAppAutoSupportAllOf`

NewStorageNetAppAutoSupportAllOfWithDefaults instantiates a new StorageNetAppAutoSupportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppAutoSupportAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppAutoSupportAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppAutoSupportAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppAutoSupportAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppAutoSupportAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppAutoSupportAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppAutoSupportAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppAutoSupportAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppAutoSupportAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppAutoSupportAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFrom

`func (o *StorageNetAppAutoSupportAllOf) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *StorageNetAppAutoSupportAllOf) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *StorageNetAppAutoSupportAllOf) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *StorageNetAppAutoSupportAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetProxyUrl

`func (o *StorageNetAppAutoSupportAllOf) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *StorageNetAppAutoSupportAllOf) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *StorageNetAppAutoSupportAllOf) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *StorageNetAppAutoSupportAllOf) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetTransport

`func (o *StorageNetAppAutoSupportAllOf) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *StorageNetAppAutoSupportAllOf) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *StorageNetAppAutoSupportAllOf) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *StorageNetAppAutoSupportAllOf) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


